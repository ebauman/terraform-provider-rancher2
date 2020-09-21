package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspAllowedFlexVolumesConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"driver": &schema.Schema {
			Type: schema.TypeString,
			Required: true,
			Description: "Flex volume driver name",
		},
	}

	return s
}
