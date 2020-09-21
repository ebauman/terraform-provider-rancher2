package rancher2

import (
	managementClient "github.com/rancher/types/client/management/v3"
)

func flattenCsiDrivers(p []managementClient.AllowedCSIDriver) []interface{} {
	if len(p) == 0 {
		return []interface{}{}
	}

	out := make([]interface{}, len(p))
	for i, in := range p {
		obj := make(map[string]interface{})

		if len(in.Name) > 0 {
			obj["name"] = in.Name
		}

		out[i] = obj
	}

	return out
}
