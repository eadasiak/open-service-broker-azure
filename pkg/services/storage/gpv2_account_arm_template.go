package storage

// nolint: lll
var armAccountTemplateBytes = []byte(`
{
	"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	"contentVersion": "1.0.0.0",
	"parameters": {
		"accountType": {
			"type": "string",
			"defaultValue": "Standard_LRS",
			"allowedValues": [
				"Standard_LRS",
				"Standard_GRS",
				"Standard_RAGRS",
				"Standard_ZRS",
				"Premium_LRS"
			]
		},
		"tags": {
			"type": "object"
		}
	},
	"resources": [
		{
			"type": "Microsoft.Storage/storageAccounts",
			"name": "{{ .name }}",
			"apiVersion": "2019-04-01",
			"location": "{{ .location }}",
			"sku": {
				"name": "{{.accountType}}"
			},
			"kind": "{{.kind}}",
			"properties": {
				{{ if .accessTier }}
				"accessTier": "{{.accessTier}}",
				{{ end }}
				"supportsHttpsTrafficOnly": {{.supportHttpsTrafficOnly}}
			},
			"tags": "[parameters('tags')]"
		}
	],
	"outputs": {
	}
}
`)
