package storage

// nolint: lll
var armLifecyclePolicyTemplateBytes = []byte(`
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
			"type": "Microsoft.Storage/storageAccounts/managementPolicies",
			"apiVersion": "2019-04-01",
			"properties": {
				"policy": {
					"rules": [
						{{ range $ruleIndex,$rule := .rules }}
						{{if $ruleIndex}},{{ end }}
						{
							"enabled": "{{ $rule.enabled }}",
							"name": "{{ $rule.name }}",
							"type": "Lifecycle",
							"definition": {
								{{ if (index $rule.definition "filters") }}
								"filters": {
									{{ if (index $rule.definition.filters "prefixMatch") }}
									"prefixMatch": [
										{{ range $prefixIndex, $prefixMatch := $rule.definition.filters.prefixMatch }}
										{{ if $prefixIndex }},{{ end }}"{{ $prefixMatch }}"{{ end }}
									],
									{{ end }}
									"blobTypes": [
										{{ range $blobIndex, $blobType := $rule.definition.filters.blobTypes }}
										{{ if $blobIndex }},{{ end }}"{{ $blobType }}"{{ end }}
									]
								},
								{{ end }}
								"actions": {
									{{ if (index $rule.definition.actions "baseBlob") }}
									"baseBlob": {
										{{ if (index $rule.definition.actions.baseBlob "tierToCool") }}
										"tierToCool": {
											"daysAfterModificationGreaterThan": "{{ $rule.definition.actions.baseBlob.tierToCool.daysAfterModificationGreaterThan }}"
										}{{ if or (index $rule.definition.actions.baseBlob "tierToArchive") (index $rule.definition.actions.baseBlob "delete")}},{{end}}
										{{ end }}
										{{ if (index $rule.definition.actions.baseBlob "tierToArchive") }}
										"tierToArchive": {
											"daysAfterModificationGreaterThan": "{{ $rule.definition.actions.baseBlob.tierToArchive.daysAfterModificationGreaterThan }}"
										}{{ if (index $rule.definition.actions.baseBlob "delete")}},{{end}}
										{{ end }}
										{{ if (index $rule.definition.actions.baseBlob "delete") }}
										"delete": {
											"daysAfterModificationGreaterThan": "{{ $rule.definition.actions.baseBlob.delete.daysAfterModificationGreaterThan }}"
										}
										{{ end }}
									}{{ if (index $rule.definition.actions "snapshot") }},{{end}}
									{{ end }}
									{{ if (index $rule.definition.actions "snapshot") }}
									"snapshot": {
										{{ if (index $rule.definition.actions.snapshot "delete") }}
										"delete": {
											"daysAfterCreationGreaterThan": "{{ $rule.definition.actions.snapshot.delete.daysAfterCreationGreaterThan }}"
										}
										{{ end }}	
									}							
									{{ end }}
								}
							}
						}{{ end }}
					]
				}
			}
		}
	],
	"outputs": {
	}
}
`)
