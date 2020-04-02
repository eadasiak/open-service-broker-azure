package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
	uuid "github.com/satori/go.uuid"
)

func (b *blobServicesManager) GetProvisioner(
	_ service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("preProvision", b.preProvision),
		service.NewProvisioningStep("deployARMTemplate", b.deployARMTemplate),
	)
}

func (b *blobServicesManager) preProvision(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	pdt := instance.Parent.Details.(*instanceDetails)
	dt := instanceDetails{
		ARMDeploymentName:  uuid.NewV4().String(),
		StorageAccountName: pdt.StorageAccountName,
	}
	return &dt, nil
}

// TODO: add storage name verification

func (b *blobServicesManager) deployARMTemplate(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	goTemplateParams := buildBlobServicesGoTemplate(instance, *instance.ProvisioningParameters)
	tagsObj := instance.ProvisioningParameters.GetObject("tags")
	tags := make(map[string]string, len(tagsObj.Data))
	for k := range tagsObj.Data {
		tags[k] = tagsObj.GetString(k)
	}

	_, err := b.armDeployer.Deploy(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		instance.ProvisioningParameters.GetString("location"),
		armBlobServicesTemplateBytes,
		goTemplateParams,         // Go template params
		map[string]interface{}{}, // ARM template params
		tags,
	)
	if err != nil {
		return nil, fmt.Errorf("error deploying ARM template: %s", err)
	}

	return dt, nil
}

func buildBlobServicesGoTemplate(
	instance service.Instance,
	parameter service.ProvisioningParameters,
) map[string]interface{} {
	dt := instance.Details.(*instanceDetails)

	p := map[string]interface{}{}
	p["storageAccountName"] = dt.StorageAccountName
	corsRulesParams := parameter.GetObjectArray("corsRules")
	corsRules := make([]map[string]interface{}, len(corsRulesParams))
	for i, corsRulesParams := range corsRulesParams {
		corsRules[i] = corsRulesParams.Data
	}
	p["corsRules"] = corsRules
	deleteRetentionPolicy := parameter.GetObject("deleteRetentionPolicy").Data
	deleteRetentionPolicyBytes, _ := json.Marshal(deleteRetentionPolicy)
	p["deleteRetentionPolicy"] = string(deleteRetentionPolicyBytes)
	p["automaticSnapshotPolicyEnabled"] = parameter.GetString("automaticSnapshotPolicyEnabled")
	fmt.Printf("template params:\n%+v\n", p)
	return p
}
