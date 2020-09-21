package rancher2

import (
	managementClient "github.com/rancher/types/client/management/v3"
)

func flattenFlexVolumes(p []managementClient.AllowedFlexVolume) []interface{} {
	if len(p) == 0 {
		return []interface{}{}
	}

	out := make([]interface{}, len(p))
	for i, in := range p {
		obj := make(map[string]interface{})

		if len(in.Driver) > 0 {
			obj["driver"] = in.Driver
		}

		out[i] = obj
	}

	return out
}