package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

const serviceGeneralPurposeV2 = "azure-storage-general-purpose-v2-storage-account" // nolint: lll
const serviceGeneralPurposeV1 = "azure-storage-general-purpose-v1-storage-account" // nolint: lll
const serviceBlobAllInOne = "azure-storage-blob-storage-account-and-container"
const serviceBlobAccount = "azure-storage-blob-storage-account"
const serviceBlobContainer = "azure-storage-blob-container"
const serviceLifecyclePolicy = "azure-storage-lifecycle-management-policy"
const serviceBlobServices = "azure-storage-blob-services"

// nolint: lll
func (m *module) GetCatalog() (service.Catalog, error) {
	return service.NewCatalog([]service.Service{
		service.NewService(
			service.ServiceProperties{
				ID:   "9a3e28fe-8c02-49da-9b35-1b054eb06c95",
				Name: serviceGeneralPurposeV2,
				Description: "Azure general purpose v2 storage account; create your " +
					"own containers, files, and tables within this account",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage General Purpose V2 Storage Account",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "Azure general purpose v2 storage account; create your " +
						"own containers, files, and tables within this account (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				ChildServiceID: "0df2a5e4-d0fb-4873-a419-278d7b9af004",
				Bindable:       true,
				Tags:           []string{"Azure", "Storage"},
			},
			m.generalPurposeV2Manager,
			service.NewPlan(service.PlanProperties{
				ID:   "bc4f766a-c372-479c-b0b4-bd9d0546b3ef",
				Name: "account",
				Description: "Azure general purpose v2 storage account; create your " +
					"own containers, files, and tables within this account",
				Free:      false,
				Stability: service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "General Purpose V2 Storage Account",
					Bullets: []string{"Azure general-purpose v2 storage account",
						"Create your own containers, files, and tables within this account",
					},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateGPv2AccountProvisioningParamsSchema(serviceGeneralPurposeV2),
						UpdatingParametersSchema:     generateGPv2AccountUpdatingParamsSchema(serviceGeneralPurposeV2),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:   "d10ea062-b627-41e8-a240-543b60030694",
				Name: serviceGeneralPurposeV1,
				Description: "Azure general purpose v1 storage account; create your " +
					"own containers, files, and tables within this account",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage General Purpose V1 Storage Account",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "Azure general purpose v1 storage account; create your " +
						"own containers, files, and tables within this account (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "Storage"},
			},
			m.generalPurposeV1Manager,
			service.NewPlan(service.PlanProperties{
				ID:   "9364d013-3690-4ce5-b0a2-b43d9b970b02",
				Name: "account",
				Description: "General-purpose v1 accounts provide access to all " +
					"Azure Storage services, but may not have the latest features" +
					"or the lowest per gigabyte pricing",
				Free:      false,
				Stability: service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "General Purpose V1 Storage Account",
					Bullets: []string{"Azure general-purpose v1 storage account",
						"Create your own containers, files, and tables within this account",
					},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateProvisioningParamsSchema(serviceGeneralPurposeV1),
						UpdatingParametersSchema:     generateUpdatingParamsSchema(serviceGeneralPurposeV1),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:   "1a5b4582-29a3-48c5-9cac-511fd8c52756",
				Name: serviceBlobAccount,
				Description: "Specialized Azure storage account for storing block " +
					"blobs and append blobs",
				ChildServiceID: "fb6ce656-c16d-4b48-aff9-286714298af8",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Blob Storage Account",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "Specialized Azure storage account for storing block " +
						"blobs and append blobs (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "Storage"},
			},
			m.blobAccountManager,
			service.NewPlan(service.PlanProperties{
				ID:   "98ae02ec-da21-4b09-b5e0-e2f9583d565c",
				Name: "account",
				Description: "Specialized Azure storage account for storing block " +
					"blobs and append blobs; create your own blob containers within " +
					"this account",
				Free:      false,
				Stability: service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Blob Storage Account",
					Bullets: []string{"Specialized Azure storage account for storing " +
						"block blobs and append blobs",
						"Create your own containers, files, and tables within this account",
					},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateProvisioningParamsSchema(serviceBlobAccount),
						UpdatingParametersSchema:     generateUpdatingParamsSchema(serviceBlobAccount),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:   "d799916e-3faf-4bdf-a48b-bf5012a2d38c",
				Name: serviceBlobAllInOne,
				Description: "A specialized Azure storage account for storing block " +
					"blobs and append blobs; automatically provisions a blob container " +
					"within the account",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Blob Storage Account And Container",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "A specialized Azure storage account for storing block " +
						"blobs and append blobs; automatically provisions a blob container " +
						"within the account (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "Storage"},
			},
			m.blobAllInOneManager,
			service.NewPlan(service.PlanProperties{
				ID:   "6c3b587d-0f88-4112-982a-dbe541f30669",
				Name: "all-in-one",
				Description: "A specialized Azure storage account for storing block " +
					"blobs and append blobs; automatically provisions a blob container " +
					"within the account",
				Free:      false,
				Stability: service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Blob Storage Account And Container",
					Bullets: []string{"A specialized Azure storage account for storing " +
						"block blobs and append blobs; " +
						"automatically provisions a blob container within the account",
					},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateProvisioningParamsSchema(serviceBlobAllInOne),
						UpdatingParametersSchema:     generateUpdatingParamsSchema(serviceBlobAllInOne),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:              "fb6ce656-c16d-4b48-aff9-286714298af8",
				Name:            serviceBlobContainer,
				Description:     "A blob container inside an existing blob storage account",
				ParentServiceID: "1a5b4582-29a3-48c5-9cac-511fd8c52756",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Blob Container",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "A blob container inside an existing blob storage account" +
						" (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "Storage"},
			},
			m.blobContainerManager,
			service.NewPlan(service.PlanProperties{
				ID:          "6b120780-c1f1-49ba-83c1-ffbd6b81df5e",
				Name:        "container",
				Description: "A blob container inside an existing blob storage account",
				Free:        false,
				Stability:   service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Blob Container",
					Bullets:     []string{"A blob container inside an existing blob storage account"},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateBlobContainerProvisioningParamsSchema(),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:              "a9394a78-659d-4321-95e9-b1e0391497bd",
				Name:            serviceLifecyclePolicy,
				Description:     "A storage lifecycle management policy",
				ParentServiceID: "9a3e28fe-8c02-49da-9b35-1b054eb06c95",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Lifecycle Management Policy",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "Lifecycle Management Policies for Azure Blob Storage " +
						"offers rule-based policies for transitioning data to the appropriate " +
						"access tiers or exire at the end of the data's lifecycle.",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: false,
				Tags:     []string{"Azure", "Storage"},
			},
			m.lifecyclePolicyManager,
			service.NewPlan(service.PlanProperties{
				ID:          "5ae13555-84d3-4c5d-a561-8aa6eebba95d",
				Name:        "policy",
				Description: "A lifecycle management policy inside an existing GPv2 storage account",
				Free:        false,
				Stability:   service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Lifecycle Management Policy",
					Bullets:     []string{"A lifecycle management policy inside an existing GPv2 storage account"},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateLifecyclePolicyProvisioningParamsSchema(),
						UpdatingParametersSchema:     generateLifecyclePolicyUpdatingParamsSchema(),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:              "0df2a5e4-d0fb-4873-a419-278d7b9af004",
				Name:            serviceBlobServices,
				Description:     "A blob services resource",
				ParentServiceID: "9a3e28fe-8c02-49da-9b35-1b054eb06c95",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Blob Services Resource",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "Allows you to set CORS rules and retention policies " +
						"on blob containers in the storage account. ",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: false,
				Tags:     []string{"Azure", "Storage"},
			},
			m.blobServicesManager,
			service.NewPlan(service.PlanProperties{
				ID:          "2767deb4-d1df-4dd6-a459-0c70591d3476",
				Name:        "blob-services",
				Description: "A blob services resource inside an existing GPv2 storage account",
				Free:        false,
				Stability:   service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Blob Services",
					Bullets:     []string{"A blob services resource inside an existing GPv2 storage account"},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateGPv2BlobServicesProvisioningParamsSchema(),
						UpdatingParametersSchema:     generateGPv2BlobServicesUpdatingParamsSchema(),
					},
				},
			}),
		),
		service.NewService(
			service.ServiceProperties{
				ID:              "102c3624-dbb3-4f33-a4aa-e371bd158b4f",
				Name:            "azure-storage-gpv2-blob-container",
				Description:     "A blob container inside an existing GPv2 blob storage account",
				ParentServiceID: "9a3e28fe-8c02-49da-9b35-1b054eb06c95",
				Metadata: service.ServiceMetadata{
					DisplayName: "Azure Storage Blob Container",
					ImageURL: "https://raw.githubusercontent.com/MicrosoftDocs/" +
						"azure-docs/9eb1f875f3823af85e41ebc97e31c5b7202bf419/articles/media" +
						"/index/Storage.svg?sanitize=true",
					LongDescription: "A blob container inside an existing GPv2 blob storage account" +
						" (Preview)",
					DocumentationURL: "https://docs.microsoft.com/en-us/azure/storage/",
					SupportURL:       "https://azure.microsoft.com/en-us/support/",
				},
				Bindable: true,
				Tags:     []string{"Azure", "Storage"},
			},
			m.gpv2BlobContainerManager,
			service.NewPlan(service.PlanProperties{
				ID:          "6280a925-cdc8-4deb-82e4-45cf7a792919",
				Name:        "container",
				Description: "A blob container inside an existing GPv2 blob storage account",
				Free:        false,
				Stability:   service.StabilityPreview,
				Metadata: service.ServicePlanMetadata{
					DisplayName: "Blob Container",
					Bullets:     []string{"A blob container inside an existing blob storage account"},
				},
				Schemas: service.PlanSchemas{
					ServiceInstances: service.InstanceSchemas{
						ProvisioningParametersSchema: generateGPv2BlobContainerProvisioningParamsSchema(),
					},
				},
			}),
		),
	}), nil
}
