package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspRuntimeClassConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"allowed_runtime_class_names": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"default_runtime_class_name": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
	}

	return s
}