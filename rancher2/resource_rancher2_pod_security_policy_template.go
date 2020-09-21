package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceRancher2PodSecurityPolicyTemplate() *schema.Resource {
	//return &schema.Resource {
	//
	//}
}

func resourceRancher2PSPTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}


}