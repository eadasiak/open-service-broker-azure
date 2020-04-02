package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
	uuid "github.com/satori/go.uuid"
)

func (l *lifecyclePolicyManager) GetProvisioner(
	_ service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep("preProvision", l.preProvision),
		service.NewProvisioningStep("deployARMTemplate", l.deployARMTemplate),
	)
}

func (l *lifecyclePolicyManager) preProvision(
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

func (l *lifecyclePolicyManager) deployARMTemplate(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	goTemplateParams := buildLifecycleGoTemplate(instance, *instance.ProvisioningParameters)
	tagsObj := instance.ProvisioningParameters.GetObject("tags")
	tags := make(map[string]string, len(tagsObj.Data))
	for k := range tagsObj.Data {
		tags[k] = tagsObj.GetString(k)
	}

	_, err := l.armDeployer.Deploy(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		instance.ProvisioningParameters.GetString("location"),
		armLifecyclePolicyTemplateBytes,
		goTemplateParams,         // Go template params
		map[string]interface{}{}, // ARM template params
		tags,
	)
	if err != nil {
		return nil, fmt.Errorf("error deploying ARM template: %s", err)
	}

	return dt, nil
}

func buildLifecycleGoTemplate(
	instance service.Instance,
	parameter service.ProvisioningParameters,
) map[string]interface{} {
	dt := instance.Details.(*instanceDetails)

	p := map[string]interface{}{}
	p["storageAccountName"] = dt.StorageAccountName
	policyRulesParams := parameter.GetObjectArray("rules")
	policyRules := make([]map[string]interface{}, len(policyRulesParams))
	for i, policyRulesParams := range policyRulesParams {
		policyRules[i] = policyRulesParams.Data
	}
	p["rules"] = policyRules

	return p
}
