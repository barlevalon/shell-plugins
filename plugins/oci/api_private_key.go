package oci

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIPrivateKey() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIPrivateKey, // TODO: Register name in project://sdk/schema/credname/names.go
		DocsURL:       sdk.URL("https://oci.com/docs/api_private_key"), // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.oci.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Key,
				MarkdownDescription: "Key used to authenticate to OCI.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 31,
					Charset: schema.Charset{
						Uppercase: true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TryOCIConfigFile(),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"OCI_KEY": fieldname.Key, // TODO: Check if this is correct
}

// TODO: Check if the platform stores the API Private Key in a local config file, and if so,
// implement the function below to add support for importing it.
func TryOCIConfigFile() sdk.Importer {
	return importer.TryFile("~/path/to/config/file.yml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		// var config Config
		// if err := contents.ToYAML(&config); err != nil {
		// 	out.AddError(err)
		// 	return
		// }

		// if config.Key == "" {
		// 	return
		// }

		// out.AddCandidate(sdk.ImportCandidate{
		// 	Fields: map[sdk.FieldName]string{
		// 		fieldname.Key: config.Key,
		// 	},
		// })
	})
}

// TODO: Implement the config file schema
// type Config struct {
//	Key string
// }
