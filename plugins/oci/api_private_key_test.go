package oci

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestAPIPrivateKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIPrivateKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Key: "AXAMENJQTVQZUMEOTLKGLPRCEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"OCI_KEY": "AXAMENJQTVQZUMEOTLKGLPRCEXAMPLE",
				},
			},
		},
	})
}

func TestAPIPrivateKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIPrivateKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"OCI_KEY": "AXAMENJQTVQZUMEOTLKGLPRCEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Key: "AXAMENJQTVQZUMEOTLKGLPRCEXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in oci/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "AXAMENJQTVQZUMEOTLKGLPRCEXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
