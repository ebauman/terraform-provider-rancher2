package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspSelinuxOptionsConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"level": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
		"role": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
		"user": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
	}

	return s
}
