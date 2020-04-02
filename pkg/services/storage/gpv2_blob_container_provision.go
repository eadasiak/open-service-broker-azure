package storage

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/open-service-broker-azure/pkg/service"
	uuid "github.com/satori/go.uuid"
)

func (b *gpv2BlobContainerManager) GetProvisioner(
	_ service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("checkNameAvailability", b.CheckNameAvailability),
		service.NewProvisioningStep("preProvision", b.PreProvision),
		service.NewProvisioningStep("deployARMTemplate", b.DeployARMTemplate),
	)
}

func (b *gpv2BlobContainerManager) CheckNameAvailability(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	containerName := instance.ProvisioningParameters.GetString("containerName")
	if containerName == "" {
		return nil, nil
	}

	pdt := instance.Parent.Details.(*instanceDetails)
	client, _ := storage.NewBasicClient(
		pdt.StorageAccountName,
		pdt.AccessKey,
	)
	blobCli := client.GetBlobService()
	response, err := blobCli.ListContainers(storage.ListContainersParameters{})
	if err != nil {
		return nil, fmt.Errorf("error checking name availability %s", err)
	}
	containers := response.Containers
	for _, container := range containers {
		if containerName == container.Name {
			return nil, fmt.Errorf(
				"container having name %s already exists in the storage account",
				containerName,
			)
		}
	}
	return nil, nil
}

func (b *gpv2BlobContainerManager) PreProvision(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pdt := instance.Parent.Details.(*instanceDetails)
	dt := instanceDetails{
		ARMDeploymentName:  uuid.NewV4().String(),
		StorageAccountName: pdt.StorageAccountName,
	}
	if instance.ProvisioningParameters.GetString("containerName") != "" {
		dt.ContainerName = instance.ProvisioningParameters.GetString("containerName")
	} else {
		dt.ContainerName = uuid.NewV4().String()
	}
	return &dt, nil
}

func (b *gpv2BlobContainerManager) DeployARMTemplate(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	goTemplateParams := buildBlobContainerGoTemplate(instance, *instance.ProvisioningParameters)
	tagsObj := instance.ProvisioningParameters.GetObject("tags")
	tags := make(map[string]string, len(tagsObj.Data))
	for k := range tagsObj.Data {
		tags[k] = tagsObj.GetString(k)
	}

	_, err := b.armDeployer.Deploy(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		instance.ProvisioningParameters.GetString("location"),
		armBlobContainerTemplateBytes,
		goTemplateParams,         // Go template params
		map[string]interface{}{}, // ARM template params
		tags,
	)
	if err != nil {
		return nil, fmt.Errorf("error deploying ARM template: %s", err)
	}

	return dt, nil
}

func buildBlobContainerGoTemplate(
	instance service.Instance,
	parameter service.ProvisioningParameters,
) map[string]interface{} {
	dt := instance.Details.(*instanceDetails)

	p := map[string]interface{}{}
	p["storageAccountName"] = dt.StorageAccountName
	p["containerName"] = parameter.GetString("containerName")
	p["publicAccess"] = parameter.GetString("publicAccess")
	return p
}
