package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
	uuid "github.com/satori/go.uuid"
)

func (b *gpv2BlobContainerManager) GetProvisioner(
	_ service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("preProvision", b.PreProvision),
		service.NewProvisioningStep("deployARMTemplate", b.DeployARMTemplate),
	)
}

func (b *gpv2BlobContainerManager) PreProvision(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pdt := instance.Parent.Details.(*instanceDetails)
	dt := instanceDetails{
		ARMDeploymentName:  uuid.NewV4().String(),
		StorageAccountName: pdt.StorageAccountName,
	}
	requestedName := instance.ProvisioningParameters.GetString(
		"containerName",
	)
	// If a specific name was requested, check availability
	if requestedName != "" {
		// First, retrieve the requested container name in the storage account
		container, err := b.blobContainersClient.Get(
			ctx,
			instance.ProvisioningParameters.GetString("resourceGroup"),
			pdt.StorageAccountName,
			requestedName,
		)
		// If the name wasn't found (404), set the container name to requested
		if container.Response.Response.StatusCode == 404 {
			dt.ContainerName = requestedName
		} else if container.ID != nil { // If the name was found, error out
			return nil, fmt.Errorf(
				"error with container name validation.  %s is already taken",
				requestedName,
			)	
		} else if err != nil { // Otherwise, if we got an error, return it
			return nil, fmt.Errorf("Error checking container name availability: %s", err)
		} 
	// A name wasn't requested, so assign one
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
	goTemplateParams := buildBlobContainerGoTemplate(
		instance,
		*instance.ProvisioningParameters,
	)
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
