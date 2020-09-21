package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspAllowedCSIDriversConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": &schema.Schema{
			Type: schema.TypeString,
			Required: true,
			Description: "CSI driver name",
		},
	}

	return s
}
