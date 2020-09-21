package rancher2

import (
	managementClient "github.com/rancher/types/client/management/v3"
)

func flattenFsGroup(p managementClient.FSGroupStrategyOptions) interface{} {
	out := make([]interface{}, len(p))

	// todo - figure out how to flatten fs group strategy options
	// it is not a list
}
