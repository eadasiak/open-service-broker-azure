package storage

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (l *lifecyclePolicyManager) GetDeprovisioner(
	service.Plan,
) (service.Deprovisioner, error) {
	return service.NewDeprovisioner(
		service.NewDeprovisioningStep(
			"deleteARMDeployment",
			l.deleteARMDeployment,
		),
		service.NewDeprovisioningStep(
			"deleteLifecyclePolicy",
			l.deleteLifecyclePolicy,
		),
	)
}

func (l *lifecyclePolicyManager) deleteARMDeployment(
	_ context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	dt := instance.Details.(*instanceDetails)

	if err := l.armDeployer.Delete(
		dt.ARMDeploymentName,
		instance.ProvisioningParameters.GetString("resourceGroup"),
	); err != nil {
		return nil, fmt.Errorf("error deleting ARM deployment: %s", err)
	}
	return instance.Details, nil
}

func (l *lifecyclePolicyManager) deleteLifecyclePolicy(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	//dt := instance.Details.(*instanceDetails)
	pdt := instance.Parent.Details.(*instanceDetails)
	_, err := l.policyClient.Delete(
		ctx,
		instance.ProvisioningParameters.GetString("resourceGroup"),
		pdt.StorageAccountName,
	)
	if err != nil {
		return nil, fmt.Errorf("error deleting lifecycle management policy: %s", err)
	}
	return instance.Details, nil
}
