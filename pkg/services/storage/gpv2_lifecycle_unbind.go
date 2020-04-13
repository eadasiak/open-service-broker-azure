package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (l *lifecyclePolicyManager) Unbind(service.Instance, service.Binding) error { // nolint: lll
	return nil
}
