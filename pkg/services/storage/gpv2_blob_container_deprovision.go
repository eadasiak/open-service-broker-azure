package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (b *gpv2BlobContainerManager) GetDeprovisioner(
	service.Plan,
) (service.Deprovisioner, error) {
	return service.NewDeprovisioner(
		service.NewDeprovisioningStep("deleteBlobContainer", b.deleteBlobContainer),
		service.NewDeprovisioningStep("deleteARMDeployment", b.deleteARMDeployment),
	)
}

func (b *gpv2BlobContainerManager) deleteBlobContainer(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)
	if err := deleteBlobContainer(
		dt.StorageAccountName,
		dt.AccessKey,
		dt.ContainerName,
	); err != nil {
		return nil, err
	}
	return instance.Details, nil
}

func (b *gpv2BlobContainerManager) deleteARMDeployment(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	if err := b.armDeployer.Delete(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
	); err != nil {
		return nil, fmt.Errorf("error deleting ARM deployment: %s", err)
	}
	return instance.Details, nil
}
