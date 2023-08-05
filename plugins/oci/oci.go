package oci

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func OCICLI() schema.Executable {
	return schema.Executable{
		Name:      "OCI CLI", // TODO: Check if this is correct
		Runs:      []string{"oci"},
		DocsURL:   sdk.URL("https://oci.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIPrivateKey,
			},
		},
	}
}
