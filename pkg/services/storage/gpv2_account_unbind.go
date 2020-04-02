package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (gpv2m *generalPurposeV2Manager) Unbind(service.Instance, service.Binding) error {
	return nil
}
