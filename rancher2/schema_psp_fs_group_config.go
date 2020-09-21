package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspFsGroupConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"ranges": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Resource {
				Schema: pspIdRangeConfigFields(),
			},
		},
		"rule": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
	}

	return s
}