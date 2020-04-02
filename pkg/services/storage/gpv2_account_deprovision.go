package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (gpv2m *generalPurposeV2Manager) GetDeprovisioner(
	service.Plan,
) (service.Deprovisioner, error) {
	return service.NewDeprovisioner(
		service.NewDeprovisioningStep("deleteARMDeployment", gpv2m.deleteARMDeployment),
		service.NewDeprovisioningStep(
			"deleteStorageAccount",
			gpv2m.deleteStorageAccount,
		),
	)
}

func (gpv2m *generalPurposeV2Manager) deleteARMDeployment(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	if err := gpv2m.armDeployer.Delete(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
	); err != nil {
		return nil, fmt.Errorf("error deleting ARM deployment: %s", err)
	}
	return instance.Details, nil
}

func (gpv2m *generalPurposeV2Manager) deleteStorageAccount(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	dt := instance.Details.(*instanceDetails)
	_, err := gpv2m.accountsClient.Delete(
		ctx,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		dt.StorageAccountName,
	)
	if err != nil {
		return nil, fmt.Errorf("error deleting storage account: %s", err)
	}
	return instance.Details, nil
}
