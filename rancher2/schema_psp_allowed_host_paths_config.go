package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspAllowedHostPathsConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"path_prefix": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
			Description: "Path prefix",
		},
		"read_only": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
			Description: "Read only",
		},
	}

	return s
}
