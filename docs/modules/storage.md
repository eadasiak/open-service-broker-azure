# [Azure Storage](https://azure.microsoft.com/en-us/services/storage/)

_Note: This module is Preview._

_This module involves the Parent-Child Model concept in OSBA, please refer to the [Parent-Child Model doc](../parent-child-model-for-multiple-layers-services.md)._

## Services & Plans

### Service: azure-storage-general-purpose-v2-storage-account

| Plan Name | Description                                                  |
| --------- | ------------------------------------------------------------ |
| `account` | This plan provisions a general purpose v2 account. General-purpose v2 storage accounts support the latest Azure Storage features and incorporate all of the functionality of general-purpose v1 and Blob storage accounts. General-purpose v2 accounts deliver the lowest per-gigabyte capacity prices for Azure Storage, as well as industry-competitive transaction prices. |

#### Behaviors

##### Provision

Provisions a general purpose v2 storage account.

###### Provisioning Parameters

| Parameter Name          | Type                | Description                                                  | Required | Default Value                                                |
| ----------------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |                                                              |
| `storageAccountName`    | `string`            | The name of the account which will be created. This name may only contain lowercase letters and numbers.  | Y | |
| `enableNonHttpsTraffic` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        | If not provided, "disabled" will be used as the default value. That is, only https traffic is allowed. |
| `accessTier`            | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. **Note** : `accountType` "Premium_LRS" only supports "Hot" in this field | N        | If not provided, "Hot" will be used as the default value.    |
| `accountType`           | `string`            | A combination of account kind and   replication strategy. All possible values: ["Standard_LRS", "Standard_GRS", "Standard_RAGRS", "Standard_ZRS", "Premium_LRS"]. **Note**: ZRS is only available in several regions, check [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-redundancy-zrs#support-coverage-and-regional-availability) for allowed regions to use ZRS. | N        | If not provided, "Standard_LRS" will be used as the default value for all plans. |
| `tags`                  | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                    | Type     | Description                                         |
| ----------------------------- | -------- | --------------------------------------------------- |
| `storageAccountName`          | `string` | The storage account name.                           |
| `accessKey`                   | `string` | A key (password) for accessing the storage account. |
| `primaryBlobServiceEndPoint`  | `string` | Primary blob service end point.                     |
| `primaryTableServiceEndPoint` | `string` | Primary table service end point.                    |
| `primaryFileServiceEndPoint`  | `string` | Primary file service end point.                     |
| `primaryQueueServiceEndPoint` | `string` | Primary queue service end point.                    |

##### Unbind

Does nothing.

##### Update

Updates an existing storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| ------------------------- | ------------------- | ------------------------------------------------------------ | -------- |
| ` enableNonHttpsTraffic ` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        |
| `accessTier`              | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. **Note** : `accountType` "Premium_LRS" only supports "Hot" in this field. | N        |
| `accountType`             | `string`            | A combination of account kind and   replication strategy. You can only update ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"] accounts to one of ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"]. For "Standard_ZRS" and "Premium_LRS" accounts, they are not updatable. | N        |
| `tags`                    | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        |

##### Deprovision

Deletes the storage account.



### Service: azure-storage-general-purpose-v1-storage-account

| Plan Name | Description                                                  |
| --------- | ------------------------------------------------------------ |
| `account` | This plan provisions a general purpose v1 account. General-purpose v1 accounts provide access to all Azure Storage services, but may not have the latest features or the lowest per gigabyte pricing. |

#### Behaviors

##### Provision

Provisions a general purpose v1 storage account.

###### Provisioning Parameters

| Parameter Name          | Type                | Description                                                  | Required | Default Value                                                |
| ----------------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |                                                              |
| `enableNonHttpsTraffic` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        | If not provided, "disabled" will be used as the default value. That is, only https traffic is allowed. |
| `accountType`           | `string`            | A combination of account kind and   replication strategy. All possible values: ["Standard_LRS", "Standard_GRS", "Standard_RAGRS", "Premium_LRS"]. | N        | If not provided, "Standard_LRS" will be used as the default value for all plans. |
| `tags`                  | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                    | Type     | Description                                         |
| ----------------------------- | -------- | --------------------------------------------------- |
| `storageAccountName`          | `string` | The storage account name.                           |
| `accessKey`                   | `string` | A key (password) for accessing the storage account. |
| `primaryBlobServiceEndPoint`  | `string` | Primary blob service end point.                     |
| `primaryTableServiceEndPoint` | `string` | Primary table service end point.                    |
| `primaryFileServiceEndPoint`  | `string` | Primary file service end point.                     |
| `primaryQueueServiceEndPoint` | `string` | Primary queue service end point.                    |

##### Unbind

Does nothing.

##### Update

Updates an existing storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| ------------------------- | ------------------- | ------------------------------------------------------------ | -------- |
| ` enableNonHttpsTraffic ` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        |
| `accountType`             | `string`            | A combination of account kind and   replication strategy. You can only update ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"] accounts to one of ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"]. For "Premium_LRS" accounts, they are not updatable. | N        |
| `tags`                    | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        |

##### Deprovision

Deletes the storage account.



### Service: azure-storage-blob-storage-account-and-container

| Plan Name    | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| `all-in-one` | This plan provisions a a specialized Azure storage account for storing block blobs and append blobs, and automatically provisions a blob container within the account. |

#### Behaviors

##### Provision

Provisions a blob storage account and create a container within the account.

###### Provisioning Parameters

| Parameter Name          | Type                | Description                                                  | Required | Default Value                                                |
| ----------------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |                                                              |
| `enableNonHttpsTraffic` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        | If not provided, "disabled" will be used as the default value. That is, only https traffic is allowed. |
| `accessTier`            | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. | N        | If not provided, "Hot" will be used as the default value.    |
| `accountType`           | `string`            | A combination of account kind and   replication strategy. All possible values: ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"]. | N        | If not provided, "Standard_LRS" will be used as the default value for all plans. |
| `containerName`         | `string`            | The name of the container which will be created inside the storage account. This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The length of the name must between 3 and 63. | N        | If not provided, a random name will be generated as the container name. |
| `tags`                  | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                   | Type     | Description                                           |
| ---------------------------- | -------- | ----------------------------------------------------- |
| `storageAccountName`         | `string` | The storage account name.                             |
| `accessKey`                  | `string` | A key (password) for accessing the storage account.   |
| `primaryBlobServiceEndPoint` | `string` | Primary blob service end point.                       |
| `containerName`              | `string` | The name of the container within the storage account. |

##### Unbind

Does nothing.

##### Update

Updates an existing storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| ------------------------- | ------------------- | ------------------------------------------------------------ | -------- |
| ` enableNonHttpsTraffic ` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        |
| `accessTier`              | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. | N        |
| `accountType`             | `string`            | A combination of account kind and   replication strategy.    | N        |
| `tags`                    | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        |

##### Deprovision

Deletes the storage account and the blob container inside it.



### Service: azure-storage-blob-storage-account

| Plan Name | Description                                                  |
| --------- | ------------------------------------------------------------ |
| `account` | This plan provisions a a specialized Azure storage account for storing block blobs and append blobs. |

#### Behaviors

##### Provision

Provisions a blob storage account.

###### Provisioning Parameters

| Parameter Name          | Type                | Description                                                  | Required | Default Value                                                |
| ----------------------- | ------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |                                                              |
| `alias`                 | `string`            | Specifies an alias that can be used by later provision actions to create containers in this storage account. | Y        |                                                              |
| `enableNonHttpsTraffic` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        | If not provided, "disabled" will be used as the default value. That is, only https traffic is allowed. |
| `accessTier`            | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. | N        | If not provided, "Hot" will be used as the default value.    |
| `accountType`           | `string`            | A combination of account kind and   replication strategy. All possible values: ["Standard_LRS", "Standard_GRS", "Standard_RAGRS"]. | N        | If not provided, "Standard_LRS" will be used as the default value for all plans. |
| `tags`                  | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        | Tags (even if none are specified) are automatically supplemented with `heritage: open-service-broker-azure`. |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                   | Type     | Description                                         |
| ---------------------------- | -------- | --------------------------------------------------- |
| `storageAccountName`         | `string` | The storage account name.                           |
| `accessKey`                  | `string` | A key (password) for accessing the storage account. |
| `primaryBlobServiceEndPoint` | `string` | Primary blob service end point.                     |

##### Unbind

Does nothing.

##### Update

Updates an existing storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| ------------------------- | ------------------- | ------------------------------------------------------------ | -------- |
| ` enableNonHttpsTraffic ` | `string`            | Specify whether non-https traffic is enabled. Allowed values:["enabled", "disabled"]. | N        |
| `accessTier`              | `string`            | The access tier used for billing.    Allowed values: ["Hot", "Cool"]. Hot storage is optimized for storing data that is accessed frequently ,and cool storage is optimized for storing data that is infrequently accessed and stored for at least 30 days. | N        |
| `accountType`             | `string`            | A combination of account kind and   replication strategy.    | N        |
| `tags`                    | `map[string]string` | Tags to be applied to new resources, specified as key/value pairs. | N        |

##### Deprovision

Deletes the storage account and the blob container inside it.

### Service: azure-storage-blob-container

| Plan Name   | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| `container` | This plan creates a container inside an existing blob storage account. |

#### Behaviors

##### Provision

Create a blob container inside an blob storage account.

###### Provisioning Parameters

| Parameter Name  | Type     | Description                                                  | Required | Default Value                                                |
| --------------- | -------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `parentAlias`   | `string` | Specifies the alias of the blob storage account upon which the  should be provisioned. | Y        |                                                              |
| `containerName` | `string` | The name of the container which will be created inside the storage account. This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The length of the name must between 3 and 63. | N        | If not provided, a random name will be generated as the container name. |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                   | Type     | Description                                           |
| ---------------------------- | -------- | ----------------------------------------------------- |
| `storageAccountName`         | `string` | The storage account name.                             |
| `accessKey`                  | `string` | A key (password) for accessing the storage account.   |
| `primaryBlobServiceEndPoint` | `string` | Primary blob service end point.                       |
| `containerName`              | `string` | The name of the container within the storage account. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the blob container.

### Service: azure-storage-gpv2-blob-container

| Plan Name   | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| `container` | This plan creates a container inside an existing blob storage account. |

#### Behaviors

##### Provision

Create a blob container inside a GPv2 storage account using an ARM template deployment.

###### Provisioning Parameters

| Parameter Name  | Type     | Description                                                  | Required | Default Value                                                |
| --------------- | -------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `parentAlias`   | `string` | Specifies the alias of the blob storage account upon which the  should be provisioned. | Y        |                                                              |
| `containerName` | `string` | The name of the container which will be created inside the storage account. This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The length of the name must between 3 and 63. | N        | If not provided, a random name will be generated as the container name. |
| `publicAccess` | `string` | Specifies whether data in the container may be accessed publicly and the level of access.  Allowed values are `Container`, `Blob`, or `None` | N | `None` |

##### Bind

Returns a copy of one shared set of credentials.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

Binding returns the following connection details and shared credentials:

| Field Name                   | Type     | Description                                           |
| ---------------------------- | -------- | ----------------------------------------------------- |
| `storageAccountName`         | `string` | The storage account name.                             |
| `accessKey`                  | `string` | A key (password) for accessing the storage account.   |
| `primaryBlobServiceEndPoint` | `string` | Primary blob service end point.                       |
| `containerName`              | `string` | The name of the container within the storage account. |

##### Unbind

Does nothing.

##### Deprovision

Deletes the blob container.

### Service: azure-storage-blob-services

| Plan Name   | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| `blob-services` | This plan provisions CORS rules and retention policies on blob containers in a GPv2 storage account. |

#### Behaviors

##### Provision

Sets blob services properties in a given GPv2 storage account.

###### Provisioning Parameters

| Parameter Name  | Type     | Description                                                  | Required | Default Value                                                |
| --------------- | -------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `parentAlias`   | `string` | Specifies the alias of the blob storage account upon which the  should be provisioned. | Y        |                                                              |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |
| `corsRules`   | `array` | A list of CORS rules. | N        |                                                              |
| `corsRules[n].allowedOrigins` | `array` | A list of origin domains that will be allowed via CORS, or '*' to allow all domains | N        |  |
| `corsRules[n].allowedMethods` | `string` | A list of HTTP methods that are allowed to be executed by the origin.  Allowed values are `DELETE`,`GET`,`HEAD`,`MERGE`,`POST`,`OPTIONS`,`PUT`. | N | |
| `corsRules[n].maxAgeInSeconds` | `integer` | The number of seconds that the client/browser should cache a preflight response | N | |
| `corsRules[n].exposedHeaders` | `array` | A list of response headers to expose to CORS clients | N | |
| `corsRules[n].allowedHeaders` | `array` | A list of headers allowed to be part of the cross-origin request | N | |
| `deleteRetentionPolicy` | `object` | The blob service properties for soft delete | N | |
| `deleteRetentionPolicy.Enabled` | `string` | Indicates whether DeleteRetentionPolicy is enabled.  Allowed values are `true` or `false` | N | |
| `deleteRetentionPolicy.Days` | `integer` | Indicates the number of days that the deleted blob should be retained (1 to 365) | N | |

##### Bind

Does nothing.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

No credentials are returned.

##### Unbind

Does nothing.

##### Update

Updates blob services settings within a GPv2 storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| --------------- | -------- | ------------------------------------------------------------ | -------- |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y       |
| `corsRules`   | `array` | A list of CORS rules. | N        |                                                              |
| `corsRules[n].allowedOrigins` | `array` | A list of origin domains that will be allowed via CORS, or '*' to allow all domains | N        |
| `corsRules[n].allowedMethods` | `string` | A list of HTTP methods that are allowed to be executed by the origin.  Allowed values are `DELETE`,`GET`,`HEAD`,`MERGE`,`POST`,`OPTIONS`,`PUT`. | N |
| `corsRules[n].maxAgeInSeconds` | `integer` | The number of seconds that the client/browser should cache a preflight response | N |
| `corsRules[n].exposedHeaders` | `array` | A list of response headers to expose to CORS clients | N |
| `corsRules[n].allowedHeaders` | `array` | A list of headers allowed to be part of the cross-origin request | N |
| `deleteRetentionPolicy` | `object` | The blob service properties for soft delete | N |
| `deleteRetentionPolicy.Enabled` | `string` | Indicates whether DeleteRetentionPolicy is enabled.  Allowed values are `true` or `false` | N |
| `deleteRetentionPolicy.Days` | `integer` | Indicates the number of days that the deleted blob should be retained (1 to 365) | N |

##### Deprovision

Resets all the blob services properties back to their default settings.

### Service: azure-storage-lifecycle-management-policy

| Plan Name   | Description                                                  |
| ----------- | ------------------------------------------------------------ |
| `policy` | This plan provisions rule-based policies in a GPv2 storage account for transitioning data to the appropriate access tiers or exire at the end of the data's lifecycle. |

#### Behaviors

##### Provision

Creates a lifecycle management policy in a given GPv2 storage account.

###### Provisioning Parameters

| Parameter Name  | Type     | Description                                                  | Required | Default Value                                                |
| --------------- | -------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| `parentAlias`   | `string` | Specifies the alias of the blob storage account upon which the  should be provisioned. | Y        |                                                              |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y        |                                                              |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y        |
| `rules` | `array` | A filter and action set for the lifecycle policy | N | |
| `rules[n].name` | `string` | Name of the policy rule. | N | |
| `rules[n].enabled` | `string` | Whether the rule is enabled.  Allowed values are `true` or `false` | N | `true` |
| `rules[n].definition` | `object` | Lifecycle filters and actions object | N | |
| `rules[n].definition.actions` | `object` | Actions applied when the run conditions are met | N | |
| `rules[n].definition.actions.baseBlob` | `object` | Blob actions | N | |
| `rules[n].definition.actions.baseBlob.delete` | `object` | Delete blob | N | |
| `rules[n].definition.actions.baseBlob.delete.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N | |
| `rules[n].definition.actions.baseBlob.tierToCool` | `object` | Move blob to Cool tier | N | |
| `rules[n].definition.actions.baseBlob.tierToCool.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N | |
| `rules[n].definition.actions.baseBlob.tierToArchive` | `object` | Archive blob | N | |
| `rules[n].definition.actions.baseBlob.tierToArchive.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N | |
| `rules[n].definition.actions.snapshot` | `object` | Blob snapshot actions | N | |
| `rules[n].definition.actions.snapshot.delete` | `object` | Delete blob snapshot | N | |
| `rules[n].definition.actions.snapshot.delete.daysAfterCreationGreaterThan` | `integer` | Integer value indicating the age in days | N | |

##### Bind

Does nothing.

###### Binding Parameters

This binding operation does not support any parameters.

###### Credentials

No credentials are returned.

##### Unbind

Does nothing.

##### Update

Updates life cycle management policies within a GPv2 storage account.

###### Updating parameters

| Parameter Name            | Type                | Description                                                  | Required |
| --------------- | -------- | ------------------------------------------------------------ | -------- |
| `location`              | `string`            | The Azure region in which to provision applicable resources. | Y |
| `resourceGroup`         | `string`            | The (new or existing) resource group with which to associate new resources. | Y       |
| `rules` | `array` | A filter and action set for the lifecycle policy | N |
| `rules[n].name` | `string` | Name of the policy rule. | N |
| `rules[n].enabled` | `string` | Whether the rule is enabled.  Allowed values are `true` or `false` | N |
| `rules[n].definition` | `object` | Lifecycle filters and actions object | N |
| `rules[n].definition.actions` | `object` | Actions applied when the run conditions are met | N |
| `rules[n].definition.actions.baseBlob` | `object` | Blob actions | N |
| `rules[n].definition.actions.baseBlob.delete` | `object` | Delete blob | N |
| `rules[n].definition.actions.baseBlob.delete.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N |
| `rules[n].definition.actions.baseBlob.tierToCool` | `object` | Move blob to Cool tier | N |
| `rules[n].definition.actions.baseBlob.tierToCool.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N |
| `rules[n].definition.actions.baseBlob.tierToArchive` | `object` | Archive blob | N |
| `rules[n].definition.actions.baseBlob.tierToArchive.daysAfterModificationGreaterThan` | `integer` | Integer value indicating the age in days | N |
| `rules[n].definition.actions.snapshot` | `object` | Blob snapshot actions | N |
| `rules[n].definition.actions.snapshot.delete` | `object` | Delete blob snapshot | N |
| `rules[n].definition.actions.snapshot.delete.daysAfterCreationGreaterThan` | `integer` | Integer value indicating the age in days | N |
##### Deprovision

Deletes the lifecycle policy from the storage account.
