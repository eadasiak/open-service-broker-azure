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
		service.NewDeprovisioningStep("deleteBlobContainer", b.deleteGpv2Container),
		service.NewDeprovisioningStep("deleteARMDeployment", b.deleteARMDeployment),
	)
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

func (b *gpv2BlobContainerManager) deleteGpv2Container(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	dt := instance.Details.(*instanceDetails)
	pdt := instance.Parent.Details.(*instanceDetails)
	_, err := b.blobContainersClient.Delete(
		ctx,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		pdt.StorageAccountName,
		dt.ContainerName,
	)
	if err != nil {
		return nil, fmt.Errorf("error deleting container: %s", err)
	}
	return instance.Details, nil
}
