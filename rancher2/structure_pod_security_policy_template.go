package rancher2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

func flattenPodSecurityPolicyTemplate(d *schema.ResourceData, in *managementClient.PodSecurityPolicyTemplate) error {
	if in == nil {
		return fmt.Errorf("[ERROR] flattening pod security policy template: Input psp is nil")
	}

	d.SetId(in.ID)
	d.Set("allow_privilege_escalation", in.AllowPrivilegeEscalation)
	if len(in.AllowedCSIDrivers) > 0 {
		d.Set("allowed_csi_drivers", flattenCsiDrivers(in.AllowedCSIDrivers))
	}

	if len(in.AllowedCapabilities) > 0 {
		d.Set("allowed_capabilities", in.AllowedCapabilities)
	}

	if len(in.AllowedFlexVolumes) > 0 {
		d.Set("allowed_flex_volumes", flattenFlexVolumes(in.AllowedFlexVolumes))
	}

	if len(in.AllowedHostPaths) > 0 {
		d.Set("allowed_host_paths", flattenHostPaths(in.AllowedHostPaths))
	}

	if len(in.AllowedProcMountTypes) > 0 {
		d.Set("allowed_proc_mount_types", in.AllowedProcMountTypes)
	}

	if len(in.AllowedUnsafeSysctls) > 0 {
		d.Set("allowed_unsafe_sysctls", in.AllowedUnsafeSysctls)
	}

	if len(in.DefaultAddCapabilities) > 0 {
		d.Set("default_add_capabilities", in.DefaultAddCapabilities)
	}

	if in.DefaultAllowPrivilegeEscalation != nil {
		d.Set("default_allow_privileges_escalation", in.DefaultAllowPrivilegeEscalation)
	}

	if len(in.Description) > 0 {
		d.Set("description", in.Description)
	}

	if in.FSGroup != nil {
		d.Set("fs_group", flattenFsGroup(in.FSGroup))
	}
}