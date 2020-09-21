package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func pspFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema {
		"allow_privilege_escalation": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"allowed_csi_drivers": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,

			Elem: &schema.Resource {
				Schema: pspAllowedCSIDriversConfigFields(),
			},
		},
		"allowed_capabilities": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"allowed_flex_volumes": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Resource {
				Schema: pspAllowedFlexVolumesConfigFields(),
			},
		},
		"allowed_host_paths": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Resource {
				Schema: pspAllowedHostPathsConfigFields(),
			},
		},
		"allowed_proc_mount_types": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allowed_unsafe_sysctls": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"annotations": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
		},
		"default_add_capabilities": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"default_allow_privilege_escalation": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"description": &schema.Schema {
			Type: schema.TypeString,
			Optional: true,
			Description: "PSP template description",
		},
		"fs_group": &schema.Schema {
			Type: schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource {
				Schema: pspFsGroupConfigFields(),
			},
		},
		"forbidden_sysctls": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"host_ipc": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"host_network": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"host_pid": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"host_ports": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Resource {
				Schema: pspHostPortsConfigFields(),
			},
		},
		"labels": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
		},
		"privileged": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"read_only_root_filesystem": &schema.Schema {
			Type: schema.TypeBool,
			Optional: true,
		},
		"required_drop_capabilities": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
		"run_as_group": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspRunAsGroupConfigFields(),
			},
		},
		"run_as_user": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspRunAsUserConfigFields(),
			},
		},
		"runtime_class": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspRuntimeClassConfigFields(),
			},
		},
		"selinux": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspSelinuxConfigFields(),
			},
		},
		"supplemental_groups": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource {
				Schema: pspSupplementalGroupsConfigFields(),
			},
		},
		"volumes": &schema.Schema {
			Type: schema.TypeList,
			Optional: true,
			Elem: &schema.Schema {
				Type: schema.TypeString,
			},
		},
	}

	return s
}