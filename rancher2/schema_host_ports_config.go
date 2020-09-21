package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspHostPortsConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"max": &schema.Schema {
			Type: schema.TypeInt,
			Optional: true,
		},
		"min": &schema.Schema {
			Type: schema.TypeInt,
			Optional: true,
		},
	}

	return s
}
