package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspSelinuxConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"rule": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
		},
		"selinux_options": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspSelinuxOptionsConfigFields(),
			},
		},
	}

	return s
}
