package storage

import "github.com/Azure/open-service-broker-azure/pkg/service"

func (l *lifecyclePolicyManager) GetEmptyInstanceDetails() service.InstanceDetails {
	return &instanceDetails{}
}

func (l *lifecyclePolicyManager) GetEmptyBindingDetails() service.BindingDetails {
	return nil
}

func (b *blobServicesManager) GetEmptyInstanceDetails() service.InstanceDetails {
	return &instanceDetails{}
}

func (b *blobServicesManager) GetEmptyBindingDetails() service.BindingDetails {
	return nil
}

func (b *gpv2BlobContainerManager) GetEmptyInstanceDetails() service.InstanceDetails {
	return &instanceDetails{}
}

func (b *gpv2BlobContainerManager) GetEmptyBindingDetails() service.BindingDetails {
	return nil
}
