package storage

import (
	"context"
	"fmt"

	storageSDK "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage" // nolint: lll
	"github.com/Azure/open-service-broker-azure/pkg/generate"
	"github.com/Azure/open-service-broker-azure/pkg/ptr"
	"github.com/Azure/open-service-broker-azure/pkg/service"
	uuid "github.com/satori/go.uuid"
)

func (gpv2m *generalPurposeV2Manager) GetProvisioner(
	_ service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("preProvision", gpv2m.preProvision),
		service.NewProvisioningStep("deployARMTemplate", gpv2m.deployARMTemplate),
		service.NewProvisioningStep("getAccessKey", gpv2m.getAccessKey),
	)
}

func (gpv2m *generalPurposeV2Manager) preProvision(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	dt := instanceDetails{
		ARMDeploymentName:  uuid.NewV4().String(),
		StorageAccountName: generate.NewIdentifier(),
	}
	requestedName := instance.ProvisioningParameters.GetString(
		"storageAccountName",
	)
	if requestedName != "" {
		nameAvailability, err := gpv2m.accountsClient.CheckNameAvailability(
			ctx,
			storageSDK.AccountCheckNameAvailabilityParameters{
				Name: &requestedName,
				Type: ptr.ToString("Microsoft.Storage/storageAccounts"),
			},
		)
		if err != nil {
			return nil, fmt.Errorf(
				"error with storage account name validation: %s",
				err,
			)
		}
		if *nameAvailability.NameAvailable {
			dt.StorageAccountName = requestedName
		} else {
			return nil, fmt.Errorf(
				"error with storage account name validation: %s",
				*nameAvailability.Message,
			)
		}
	}
	return &dt, nil
}

func (gpv2m *generalPurposeV2Manager) deployARMTemplate(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	goTemplateParams := buildGoTemplate(instance, *instance.ProvisioningParameters)
	tagsObj := instance.ProvisioningParameters.GetObject("tags")
	tags := make(map[string]string, len(tagsObj.Data))
	for k := range tagsObj.Data {
		tags[k] = tagsObj.GetString(k)
	}

	// Remove access key output and retrieve via API
	_, err := gpv2m.armDeployer.Deploy(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		instance.ProvisioningParameters.GetString("location"),
		armAccountTemplateBytes,
		goTemplateParams,         // Go template params
		map[string]interface{}{}, // ARM template params
		tags,
	)
	if err != nil {
		return nil, fmt.Errorf("error deploying ARM template: %s", err)
	}

	return dt, nil
}

func (gpv2m *generalPurposeV2Manager) getAccessKey(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	dt := instance.Details.(*instanceDetails)
	accessKeys, err := gpv2m.accountsClient.ListKeys(
		ctx,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		dt.StorageAccountName,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"error retrieving primary access key for account: %s",
			err,
		)
	}
	if accessKeys.Keys != nil {
		for _, key := range *accessKeys.Keys {
			if *key.KeyName == "key1" {
				dt.AccessKey = *key.Value
			}
		}
	}
	return dt, nil
}
