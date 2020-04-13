package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (b *gpv2BlobContainerManager) ValidateUpdatingParameters(
	service.Instance,
) error {
	return nil
}

func (b *gpv2BlobContainerManager) GetUpdater(
	service.Plan,
) (service.Updater, error) {
	return service.NewUpdater()
}
