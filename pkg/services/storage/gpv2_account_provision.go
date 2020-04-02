package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/generate"
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
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instanceDetails{
		ARMDeploymentName:  uuid.NewV4().String(),
		StorageAccountName: generate.NewIdentifier(),
	}
	if instance.ProvisioningParameters.GetString("storageAccountName") != "" {
		dt.StorageAccountName = instance.ProvisioningParameters.GetString("storageAccountName")
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

	// accessKey, ok := outputs["accessKey"].(string)
	// if !ok {
	// 	return nil, fmt.Errorf(
	// 		"error retrieving primary access key from deployment: %s",
	// 		err,
	// 	)
	// }
	// dt.AccessKey = accessKey
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
