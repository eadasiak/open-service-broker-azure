package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (b *gpv2BlobContainerManager) Unbind(service.Instance, service.Binding) error {
	return nil
}
