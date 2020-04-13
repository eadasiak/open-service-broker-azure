package storage

import (
	"github.com/Azure/open-service-broker-azure/pkg/ptr"
	"github.com/Azure/open-service-broker-azure/pkg/schemas"
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

const (
	hot  = "Hot"
	cool = "Cool"
)

// nolint: lll
var accountTypeMap = map[string][]string{
	"update":                {"Standard_LRS", "Standard_GRS", "Standard_RAGRS"},
	serviceBlobAccount:      {"Standard_LRS", "Standard_GRS", "Standard_RAGRS"},
	serviceBlobAllInOne:     {"Standard_LRS", "Standard_GRS", "Standard_RAGRS"},
	serviceGeneralPurposeV1: {"Standard_LRS", "Standard_GRS", "Standard_RAGRS", "Premium_LRS"},
	serviceGeneralPurposeV2: {"Standard_LRS", "Standard_GRS", "Standard_RAGRS", "Premium_LRS", "Standard_ZRS"},
}

// nolint: lll
func generateProvisioningParamsSchema(serviceName string) service.InputParametersSchema {
	ips := service.InputParametersSchema{
		RequiredProperties: []string{"location", "resourceGroup"},
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"storageAccountName": &service.StringPropertySchema{
				Title: "Storage Account Name",
				Description: "Name of the storage account.  Will be automatically " +
					"generated if not supplied",
				AllowedPattern: `^[a-z0-9]+$`,
				MinLength:      ptr.ToInt(3),
				MaxLength:      ptr.ToInt(24),
			},
			"enableNonHttpsTraffic": &service.StringPropertySchema{
				Title:        "Enable non-https traffic",
				Description:  "Specify whether non-https traffic is enabled",
				DefaultValue: schemas.DisabledParamString,
				OneOf:        schemas.EnabledDisabledValues(),
			},
			"tags": &service.ObjectPropertySchema{
				Title: "Tags",
				Description: "Tags to be applied to new resources," +
					" specified as key/value pairs.",
				Additional: &service.StringPropertySchema{},
			},
		},
	}

	ips.PropertySchemas["accountType"] = &service.StringPropertySchema{
		Title: "Account Type",
		Description: "This field is a combination of account kind and " +
			" replication strategy",
		DefaultValue:  "Standard_LRS",
		AllowedValues: accountTypeMap[serviceName],
	}

	if serviceName != serviceGeneralPurposeV1 {
		ips.PropertySchemas["accessTier"] = &service.StringPropertySchema{
			Title:         "Access Tier",
			Description:   "The access tier used for billing.",
			DefaultValue:  hot,
			AllowedValues: []string{hot, cool},
		}
	}

	if serviceName == serviceBlobAllInOne {
		ips.PropertySchemas["containerName"] = &service.StringPropertySchema{
			Title: "Container Name",
			Description: "The name of the container which will be created inside" +
				"the blob stroage account",
			AllowedPattern: `^[a-z0-9]+(?:-[a-z0-9]+)*$`,
			MinLength:      ptr.ToInt(3),
			MaxLength:      ptr.ToInt(63),
		}
	}

	return ips
}

// nolint: lll
func generateUpdatingParamsSchema(serviceName string) service.InputParametersSchema {
	ips := service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"enableNonHttpsTraffic": &service.StringPropertySchema{
				Title:       "Enable non-https traffic",
				Description: "Specify whether non-https traffic is enabled",
				OneOf:       schemas.EnabledDisabledValues(),
			},
			"tags": &service.ObjectPropertySchema{
				Title: "Tags",
				Description: "Tags to be applied to new resources," +
					" specified as key/value pairs.",
				Additional: &service.StringPropertySchema{},
			},
		},
	}

	ips.PropertySchemas["accountType"] = &service.StringPropertySchema{
		Title: "Account Type",
		Description: "This field is a combination of account kind and " +
			" replication strategy",
		DefaultValue:  "Standard_LRS",
		AllowedValues: accountTypeMap["update"],
	}

	if serviceName != serviceGeneralPurposeV1 {
		ips.PropertySchemas["accessTier"] = &service.StringPropertySchema{
			Title:         "Access Tier",
			Description:   "The access tier used for billing.",
			AllowedValues: []string{hot, cool},
		}
	}

	return ips
}

// nolint: lll
func generateBlobContainerProvisioningParamsSchema() service.InputParametersSchema {
	return service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"containerName": &service.StringPropertySchema{
				Title: "Container Name",
				Description: "The name of the container which will be created inside" +
					"the blob storage account",
				AllowedPattern: `^[a-z0-9]+(?:-[a-z0-9]+)*$`,
				MinLength:      ptr.ToInt(3),
				MaxLength:      ptr.ToInt(63),
			},
		},
	}
}

func generateGPv2AccountProvisioningParamsSchema(
	serviceName string,
) service.InputParametersSchema {
	ips := service.InputParametersSchema{
		RequiredProperties: []string{"location", "resourceGroup"},
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"storageAccountName": &service.StringPropertySchema{
				Title: "Storage Account Name",
				Description: "Name of the storage account.  Will be automatically " +
					"generated if not supplied",
				AllowedPattern: `^[a-z0-9]+$`,
			},
			"enableNonHttpsTraffic": &service.StringPropertySchema{
				Title:        "Enable non-https traffic",
				Description:  "Specify whether non-https traffic is enabled",
				DefaultValue: schemas.DisabledParamString,
				OneOf:        schemas.EnabledDisabledValues(),
			},
			"tags": &service.ObjectPropertySchema{
				Title: "Tags",
				Description: "Tags to be applied to new resources," +
					" specified as key/value pairs.",
				Additional: &service.StringPropertySchema{},
			},
		},
	}
	ips.PropertySchemas["accountType"] = &service.StringPropertySchema{
		Title: "Account Type",
		Description: "This field is a combination of account kind and " +
			" replication strategy",
		DefaultValue:  "Standard_LRS",
		AllowedValues: accountTypeMap[serviceName],
	}
	ips.PropertySchemas["accessTier"] = &service.StringPropertySchema{
		Title:         "Access Tier",
		Description:   "The access tier used for billing.",
		DefaultValue:  hot,
		AllowedValues: []string{hot, cool},
	}
	return ips
}

func generateGPv2AccountUpdatingParamsSchema(
	serviceName string,
) service.InputParametersSchema {
	ips := service.InputParametersSchema{
		RequiredProperties: []string{"location", "resourceGroup"},
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"storageAccountName": &service.StringPropertySchema{
				Title: "Storage Account Name",
				Description: "Name of the storage account.  Will be automatically " +
					"generated if not supplied",
				AllowedPattern: `^[a-z0-9]+$`,
			},
			"enableNonHttpsTraffic": &service.StringPropertySchema{
				Title:        "Enable non-https traffic",
				Description:  "Specify whether non-https traffic is enabled",
				DefaultValue: schemas.DisabledParamString,
				OneOf:        schemas.EnabledDisabledValues(),
			},
			"tags": &service.ObjectPropertySchema{
				Title: "Tags",
				Description: "Tags to be applied to new resources," +
					" specified as key/value pairs.",
				Additional: &service.StringPropertySchema{},
			},
		},
	}
	ips.PropertySchemas["accountType"] = &service.StringPropertySchema{
		Title: "Account Type",
		Description: "This field is a combination of account kind and " +
			" replication strategy",
		DefaultValue:  "Standard_LRS",
		AllowedValues: accountTypeMap[serviceName],
	}
	ips.PropertySchemas["accessTier"] = &service.StringPropertySchema{
		Title:         "Access Tier",
		Description:   "The access tier used for billing.",
		DefaultValue:  hot,
		AllowedValues: []string{hot, cool},
	}
	return ips
}

// nolint: lll
func generateLifecyclePolicyProvisioningParamsSchema() service.InputParametersSchema {
	ips := service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"rules": &service.ArrayPropertySchema{
				Title:       "Rules",
				Description: "A filter and action set for the lifecycle policy",
				ItemsSchema: &service.ObjectPropertySchema{
					Title:       "Lifecycle policy rule",
					Description: "Individual Lifecycle Management Policy Rule",
					PropertySchemas: map[string]service.PropertySchema{
						"name": &service.StringPropertySchema{
							Title:       "Name",
							Description: "Name of the policy rule",
						},
						"enabled": &service.StringPropertySchema{
							Title:        "Enabled",
							Description:  "Is the rule enabled?",
							DefaultValue: "true",
						},
						"definition": &service.ObjectPropertySchema{
							Title:       "Definition",
							Description: "Lifecycle filters and actions object",
							RequiredProperties: []string{
								"actions",
							},
							PropertySchemas: map[string]service.PropertySchema{
								"actions": &service.ObjectPropertySchema{
									Title:       "Actions",
									Description: "Actions applied when the run conditions are met",
									PropertySchemas: map[string]service.PropertySchema{
										"baseBlob": &service.ObjectPropertySchema{
											Title:       "Base Blob",
											Description: "Blob actions",
											PropertySchemas: map[string]service.PropertySchema{
												"delete": &service.ObjectPropertySchema{
													Title:       "Delete",
													Description: "Delete blob",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
												"tierToCool": &service.ObjectPropertySchema{
													Title:       "Tier to Cool",
													Description: "Move blob to Cool tier",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
												"tierToArchive": &service.ObjectPropertySchema{
													Title:       "Tier to Archive",
													Description: "Archive blob",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
											},
										},
										"snapshot": &service.ObjectPropertySchema{
											Title:       "Snapshot",
											Description: "Blob snapshot actions",
											PropertySchemas: map[string]service.PropertySchema{
												"delete": &service.ObjectPropertySchema{
													Title:       "Delete",
													Description: "Delete blob snapshot",
													RequiredProperties: []string{
														"daysAfterCreationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterCreationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Creation Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
											},
										},
									},
								},
								"filters": &service.ObjectPropertySchema{
									Title:       "Filters",
									Description: "Limit rule actions to a subset of blobs",
									RequiredProperties: []string{
										"blobTypes",
									},
									PropertySchemas: map[string]service.PropertySchema{
										"blobTypes": &service.ArrayPropertySchema{
											Title:       "Blob Types",
											Description: "Restrict actions to this type of blob",
											ItemsSchema: &service.StringPropertySchema{
												Description: "blockBlob is the only supported type",
											},
											DefaultValue: []interface{}{
												"blockBlob",
											},
										},
										"prefixMatch": &service.ArrayPropertySchema{
											Title:       "Prefix Match",
											Description: "An array of strings of prefixes to match",
											MaxItems:    ptr.ToInt(10),
											ItemsSchema: &service.StringPropertySchema{
												Description:    "prefixMatch must start with container name",
												AllowedPattern: `^[a-z0-9]+[-a-z0-9]*\/.*$`,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return ips
}

// nolint: lll
func generateLifecyclePolicyUpdatingParamsSchema() service.InputParametersSchema {
	ips := service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"rules": &service.ArrayPropertySchema{
				Title:       "Rules",
				Description: "A filter and action set for the lifecycle policy",
				ItemsSchema: &service.ObjectPropertySchema{
					Title:       "Lifecycle policy rule",
					Description: "Individual Lifecycle Management Policy Rule",
					PropertySchemas: map[string]service.PropertySchema{
						"name": &service.StringPropertySchema{
							Title:       "Name",
							Description: "Name of the policy rule",
						},
						"enabled": &service.StringPropertySchema{
							Title:        "Enabled",
							Description:  "Is the rule enabled?",
							DefaultValue: "true",
						},
						"definition": &service.ObjectPropertySchema{
							Title:       "Definition",
							Description: "Lifecycle filters and actions object",
							RequiredProperties: []string{
								"actions",
							},
							PropertySchemas: map[string]service.PropertySchema{
								"actions": &service.ObjectPropertySchema{
									Title:       "Actions",
									Description: "Actions applied when the run conditions are met",
									PropertySchemas: map[string]service.PropertySchema{
										"baseBlob": &service.ObjectPropertySchema{
											Title:       "Base Blob",
											Description: "Blob actions",
											PropertySchemas: map[string]service.PropertySchema{
												"delete": &service.ObjectPropertySchema{
													Title:       "Delete",
													Description: "Delete blob",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
												"tierToCool": &service.ObjectPropertySchema{
													Title:       "Tier to Cool",
													Description: "Move blob to Cool tier",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
												"tierToArchive": &service.ObjectPropertySchema{
													Title:       "Tier to Archive",
													Description: "Archive blob",
													RequiredProperties: []string{
														"daysAfterModificationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterModificationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Modification Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
											},
										},
										"snapshot": &service.ObjectPropertySchema{
											Title:       "Snapshot",
											Description: "Blob snapshot actions",
											PropertySchemas: map[string]service.PropertySchema{
												"delete": &service.ObjectPropertySchema{
													Title:       "Delete",
													Description: "Delete blob snapshot",
													RequiredProperties: []string{
														"daysAfterCreationGreaterThan",
													},
													PropertySchemas: map[string]service.PropertySchema{
														"daysAfterCreationGreaterThan": &service.IntPropertySchema{
															Title:       "Days After Creation Greater Than",
															Description: "Integer value indicating the age in days",
														},
													},
												},
											},
										},
									},
								},
								"filters": &service.ObjectPropertySchema{
									Title:       "Filters",
									Description: "Limit rule actions to a subset of blobs",
									RequiredProperties: []string{
										"blobTypes",
									},
									PropertySchemas: map[string]service.PropertySchema{
										"blobTypes": &service.ArrayPropertySchema{
											Title:       "Blob Types",
											Description: "Restrict actions to this type of blob",
											ItemsSchema: &service.StringPropertySchema{
												Description: "blockBlob is the only supported type",
											},
											DefaultValue: []interface{}{
												"blockBlob",
											},
										},
										"prefixMatch": &service.ArrayPropertySchema{
											Title:       "Prefix Match",
											Description: "An array of strings of prefixes to match",
											MaxItems:    ptr.ToInt(10),
											ItemsSchema: &service.StringPropertySchema{
												Description:    "prefixMatch must start with container name",
												AllowedPattern: `^[a-z0-9]+[-a-z0-9]*\/.+$`,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return ips
}

// nolint: lll
func generateGPv2BlobServicesProvisioningParamsSchema() service.InputParametersSchema {
	ips := service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"corsRules": &service.ArrayPropertySchema{
				Title:       "CORS Rules",
				Description: "The List of CORS rules",
				MaxItems:    ptr.ToInt(5),
				ItemsSchema: &service.ObjectPropertySchema{
					Title:       "CORS Rule",
					Description: "Individual CORS Rule",
					RequiredProperties: []string{
						"allowedOrigins",
						"allowedMethods",
						"maxAgeInSeconds",
						"exposedHeaders",
						"allowedHeaders",
					},
					PropertySchemas: map[string]service.PropertySchema{
						"allowedOrigins": &service.ArrayPropertySchema{
							Title:       "Allowed Origins",
							Description: "A list of origin domains that will be allowed via CORS, or '*' to allow all domains",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual domain that will be allowed via CORS",
							},
						},
						"allowedMethods": &service.ArrayPropertySchema{
							Title:       "Allowed Methods",
							Description: "A list of HTTP methods that are allowed to be executed by the origin",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual HTTP method that will be allowed via CORS (uppercase only)",
								AllowedValues: []string{
									"DELETE",
									"GET",
									"HEAD",
									"MERGE",
									"POST",
									"OPTIONS",
									"PUT",
								},
							},
						},
						"maxAgeInSeconds": &service.IntPropertySchema{
							Title:       "Max Age in Seconds",
							Description: "The number of seconds that the client/browser should cache a preflight response",
						},
						"exposedHeaders": &service.ArrayPropertySchema{
							Title:       "Exposed Headers",
							Description: "A list of response headers to expose to CORS clients",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual response header that will be exposed via CORS",
							},
						},
						"allowedHeaders": &service.ArrayPropertySchema{
							Title:       "Allowed Headers",
							Description: "A list of headers allowed to be part of the cross-origin request",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual response header that will be included in the CORS request",
							},
						},
					},
				},
			},
			"deleteRetentionPolicy": &service.ObjectPropertySchema{
				Title:       "Delete Retention Policy",
				Description: "The blob service properties for soft delete",
				PropertySchemas: map[string]service.PropertySchema{
					"enabled": &service.StringPropertySchema{
						Title:         "Enabled",
						Description:   "Indicates whether DeleteRetentionPolicy is enabled for the Blob service",
						AllowedValues: []string{"true", "false"},
						DefaultValue:  "false",
					},
					"days": &service.IntPropertySchema{
						Title:       "Days",
						Description: "Indicates the number of days that the deleted blob should be retained (1 to 365)",
						MinValue:    ptr.ToInt64(1),
						MaxValue:    ptr.ToInt64(365),
					},
				},
			},
			"automaticSnapshotPolicyEnabled": &service.StringPropertySchema{
				Title:         "Automatic Snapshot Policy Enabled",
				Description:   "Automatic Snapshot is enabled if set to true",
				AllowedValues: []string{"true", "false"},
				DefaultValue:  "false",
			},
		},
	}
	return ips
}

func generateGPv2BlobServicesUpdatingParamsSchema() service.InputParametersSchema { // nolint: lll
	ips := service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"corsRules": &service.ArrayPropertySchema{
				Title:       "CORS Rules",
				Description: "The List of CORS rules",
				MaxItems:    ptr.ToInt(5),
				ItemsSchema: &service.ObjectPropertySchema{
					Title:       "CORS Rule",
					Description: "Individual CORS Rule",
					RequiredProperties: []string{
						"allowedOrigins",
						"allowedMethods",
						"maxAgeInSeconds",
						"exposedHeaders",
						"allowedHeaders",
					},
					PropertySchemas: map[string]service.PropertySchema{
						"allowedOrigins": &service.ArrayPropertySchema{
							Title: "Allowed Origins",
							Description: "A list of origin domains that will be allowed via CORS, " +
								"or '*' to allow all domains",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual domain that will be allowed via CORS",
							},
						},
						"allowedMethods": &service.ArrayPropertySchema{
							Title: "Allowed Methods",
							Description: "A list of HTTP methods that are allowed to be executed " +
								"by the origin",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual HTTP method that will be allowed " +
									"via CORS (uppercase only)",
								AllowedValues: []string{
									"DELETE",
									"GET",
									"HEAD",
									"MERGE",
									"POST",
									"OPTIONS",
									"PUT",
								},
							},
						},
						"maxAgeInSeconds": &service.IntPropertySchema{
							Title: "Max Age in Seconds",
							Description: "The number of seconds that the client/browser " +
								"should cache a preflight response",
						},
						"exposedHeaders": &service.ArrayPropertySchema{
							Title:       "Exposed Headers",
							Description: "A list of response headers to expose to CORS clients",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual response header that will be " +
									"exposed via CORS",
							},
						},
						"allowedHeaders": &service.ArrayPropertySchema{
							Title: "Allowed Headers",
							Description: "A list of headers allowed to be part of the " +
								"cross-origin request",
							ItemsSchema: &service.StringPropertySchema{
								Description: "An individual response header that will be " +
									"included in the CORS request",
							},
						},
					},
				},
			},
			"deleteRetentionPolicy": &service.ObjectPropertySchema{
				Title:       "Delete Retention Policy",
				Description: "The blob service properties for soft delete",
				PropertySchemas: map[string]service.PropertySchema{
					"enabled": &service.StringPropertySchema{
						Title: "Enabled",
						Description: "Indicates whether DeleteRetentionPolicy is enabled " +
							"for the Blob service",
						AllowedValues: []string{"true", "false"},
					},
					"days": &service.IntPropertySchema{
						Title: "Days",
						Description: "Indicates the number of days that the deleted blob " +
							"should be retained (1 to 365)",
						MinValue: ptr.ToInt64(1),
						MaxValue: ptr.ToInt64(365),
					},
				},
			},
			"automaticSnapshotPolicyEnabled": &service.StringPropertySchema{
				Title:         "Automatic Snapshot Policy Enabled",
				Description:   "Automatic Snapshot is enabled if set to true",
				AllowedValues: []string{"true", "false"},
			},
		},
	}
	return ips
}

func generateGPv2BlobContainerProvisioningParamsSchema() service.InputParametersSchema { // nolint: lll
	return service.InputParametersSchema{
		PropertySchemas: map[string]service.PropertySchema{
			"resourceGroup": schemas.GetResourceGroupSchema(),
			"location":      schemas.GetLocationSchema(),
			"storageAccountName": &service.StringPropertySchema{
				Title:          "Storage Account Name",
				Description:    "Name of the storage account",
				AllowedPattern: `^[a-z0-9]+$`,
			},
			"containerName": &service.StringPropertySchema{
				Title: "Container Name",
				Description: "The name of the container which will be created inside " +
					"the blob storage account",
				AllowedPattern: `^[a-z0-9]+(?:-[a-z0-9]+)*$`,
				MinLength:      ptr.ToInt(3),
				MaxLength:      ptr.ToInt(63),
			},
			"publicAccess": &service.StringPropertySchema{
				Title: "Public Access",
				Description: "pecifies whether data in the container may be accessed " +
					"publicly and the level of access.",
				AllowedValues: []string{
					"Container",
					"Blob",
					"None",
				},
				DefaultValue: "None",
			},
		},
	}
}
