package storage

import (
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

// Bind is not valid for Lifecycle Mangement Policies
func (l *lifecyclePolicyManager) Bind(
	service.Instance,
	service.BindingParameters,
) (service.BindingDetails, error) {
	return nil, fmt.Errorf("service is not bindable")
}

func (l *lifecyclePolicyManager) GetCredentials(
	instance service.Instance,
	binding service.Binding,
) (service.Credentials, error) {
	return nil, fmt.Errorf("service is not bindable")
}