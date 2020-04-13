package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (l *lifecyclePolicyManager) ValidateUpdatingParameters(
	instance service.Instance,
) error {
	return nil
}

func (l *lifecyclePolicyManager) GetUpdater(
	service.Plan,
) (service.Updater, error) {
	return service.NewUpdater(
		service.NewUpdatingStep("updateARMTemplate", l.updateARMTemplate),
	)
}

func (l *lifecyclePolicyManager) updateARMTemplate(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)
	up := instance.UpdatingParameters
	tagsObj := up.GetObject("tags")
	tags := make(map[string]string, len(tagsObj.Data))
	for k := range tagsObj.Data {
		tags[k] = tagsObj.GetString(k)
	}
	_, err := l.armDeployer.Update(
		dt.ARMDeploymentName,
		up.GetString("resourceGroup"),
		up.GetString("location"),
		armLifecyclePolicyTemplateBytes,
		buildLifecycleGoTemplate(instance, *up),
		map[string]interface{}{},
		tags,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating lifecycle management policy %s", err)
	}
	return dt, nil
}
