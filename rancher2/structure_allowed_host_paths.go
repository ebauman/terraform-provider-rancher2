package rancher2

import (
	managementClient "github.com/rancher/types/client/management/v3"
)

func flattenHostPaths(p []managementClient.AllowedHostPath) []interface{} {
	if len(p) == 0 {
		return []interface{}{}
	}

	out := make([]interface{}, len(p))
	for i, in := range p {
		obj := make(map[string]interface{})

		if len(in.PathPrefix) > 0 {
			obj["path_prefix"] = in.PathPrefix
		}

		obj["read_only"] = in.ReadOnly

		out[i] = obj
	}

	return out
}
