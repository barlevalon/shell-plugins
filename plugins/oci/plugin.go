package oci

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "oci",
		Platform: schema.PlatformInfo{
			Name:     "OCI",
			Homepage: sdk.URL("https://oci.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			APIPrivateKey(),
		},
		Executables: []schema.Executable{
			OCICLI(),
		},
	}
}
