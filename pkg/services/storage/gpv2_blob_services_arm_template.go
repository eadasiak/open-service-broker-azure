package storage

// nolint: lll
var armBlobServicesTemplateBytes = []byte(`
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
			"name": "{{ .storageAccountName }}/default",
			"type": "Microsoft.Storage/storageAccounts/blobServices",
			"apiVersion": "2019-04-01",
			"properties": {
				{{ if .corsRules }}
				"cors": {
					"corsRules": [
						{{ range $ruleIndex,$rule := .corsRules }}
						{{if $ruleIndex}},{{ end }}
						{
							"allowedOrigins": [
								{{ range $originIndex, $origin := $rule.allowedOrigins }}
								{{ if $originIndex }},{{ end }}"{{ $origin }}"
								{{ end }}
							],
							"allowedMethods": [
								{{ range $methodIndex, $method := $rule.allowedMethods }}
								{{ if $methodIndex }},{{ end }}"{{ $method }}"
								{{ end }}
							],
							"maxAgeInSeconds": "{{ $rule.maxAgeInSeconds }}",
							"exposedHeaders": [
								{{ range $exposedIndex, $exposed := $rule.exposedHeaders }}
								{{ if $exposedIndex }},{{ end }}"{{ $exposed }}"
								{{ end }}
							],
							"allowedHeaders": [
								{{ range $allowedIndex, $allowed := $rule.allowedHeaders }}
								{{ if $allowedIndex }},{{ end }}"{{ $allowed }}"
								{{ end }}
							]
						}{{ end }}
					]
				},
				{{ end }}
				{{ if index . "defaultServiceVersion" }}
				"defaultServiceVersion": "{{ .defaultServiceVersion }}",
				{{ end }}
				{{ if index . "deleteRetentionPolicy" }}
				"deleteRetentionPolicy": {{ .deleteRetentionPolicy }},
				{{ end }}
				{{ if index . "automaticSnapshotPolicyEnabled" }}
				"automaticSnapshotPolicyEnabled": "{{ .automaticSnapshotPolicyEnabled }}",
				{{ end }}
				"resources": []
			}
		}
	],
	"outputs": {
	}
}
`)
