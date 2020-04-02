package storage

import (
	storageSDK "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage" // nolint: lll
	"github.com/Azure/open-service-broker-azure/pkg/azure/arm"
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

type module struct {
	generalPurposeV1Manager  *generalPurposeV1Manager
	generalPurposeV2Manager  *generalPurposeV2Manager
	blobAccountManager       *blobAccountManager
	blobContainerManager     *blobContainerManager
	blobAllInOneManager      *blobAllInOneManager
	lifecyclePolicyManager   *lifecyclePolicyManager
	blobServicesManager      *blobServicesManager
	gpv2BlobContainerManager *gpv2BlobContainerManager
}

type storageManager struct {
	armDeployer    arm.Deployer
	accountsClient storageSDK.AccountsClient
}

type generalPurposeV1Manager struct {
	storageManager
}

type generalPurposeV2Manager struct {
	storageManager
}

type blobAccountManager struct {
	storageManager
}

type blobContainerManager struct {
	storageManager
}

type gpv2BlobContainerManager struct {
	storageManager
}

type blobAllInOneManager struct {
	storageManager
}

type lifecyclePolicyManager struct {
	armDeployer  arm.Deployer
	policyClient storageSDK.ManagementPoliciesClient
}

type blobServicesManager struct {
	armDeployer        arm.Deployer
	blobServicesClient storageSDK.BlobServicesClient
}

// New returns a new instance of a type that fulfills the service.Module
// interface and is capable of provisioning Storage using "Azure Storage"
func New(
	armDeployer arm.Deployer,
	accountsClient storageSDK.AccountsClient,
	blobServicesClient storageSDK.BlobServicesClient,
	policyClient storageSDK.ManagementPoliciesClient,
) service.Module {
	storageMgr := storageManager{
		armDeployer:    armDeployer,
		accountsClient: accountsClient,
	}
	return &module{
		generalPurposeV1Manager:  &generalPurposeV1Manager{storageMgr},
		generalPurposeV2Manager:  &generalPurposeV2Manager{storageMgr},
		blobAccountManager:       &blobAccountManager{storageMgr},
		blobContainerManager:     &blobContainerManager{storageMgr},
		gpv2BlobContainerManager: &gpv2BlobContainerManager{storageMgr},
		blobAllInOneManager:      &blobAllInOneManager{storageMgr},
		blobServicesManager: &blobServicesManager{
			armDeployer:        armDeployer,
			blobServicesClient: blobServicesClient,
		},
		lifecyclePolicyManager: &lifecyclePolicyManager{
			armDeployer:  armDeployer,
			policyClient: policyClient,
		},
	}
}

func (m *module) GetName() string {
	return "storage"
}
