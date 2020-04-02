package storage

// nolint: lll
var armBlobContainerTemplateBytes = []byte(`
{
	"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	"contentVersion": "1.0.0.0",
	"parameters": {
		"tags": {
			"type": "object"
		}
	},
	"resources": [
		{
			"name": "{{ .storageAccountName }}/default/{{ .containerName }}",
			"type": "Microsoft.Storage/storageAccounts/blobServices/containers",
			"apiVersion": "2019-04-01",
			"properties": {
				{{ if .publicAccess }}
				"publicAccess": "{{ .publicAccess }}"
				{{ end }}
			}
		}
	],
	"outputs": {
	}
}
`)
