package storage

import "github.com/Azure/open-service-broker-azure/pkg/service"

func (l *lifecyclePolicyManager) GetEmptyInstanceDetails() service.InstanceDetails { // nolint: lll
	return &instanceDetails{}
}

func (l *lifecyclePolicyManager) GetEmptyBindingDetails() service.BindingDetails { // nolint: lll
	return nil
}

func (b *blobServicesManager) GetEmptyInstanceDetails() service.InstanceDetails { // nolint: lll
	return &instanceDetails{}
}

func (b *blobServicesManager) GetEmptyBindingDetails() service.BindingDetails { // nolint: lll
	return nil
}

func (b *gpv2BlobContainerManager) GetEmptyInstanceDetails() service.InstanceDetails { // nolint: lll
	return &instanceDetails{}
}

func (b *gpv2BlobContainerManager) GetEmptyBindingDetails() service.BindingDetails { // nolint: lll
	return nil
}

func (a *generalPurposeV2Manager) GetEmptyInstanceDetails() service.InstanceDetails {
	return &instanceDetails{}
}

func (a *generalPurposeV2Manager) GetEmptyBindingDetails() service.BindingDetails {
	return nil
}
