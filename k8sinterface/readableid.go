package k8sinterface

import (
	"fmt"

	"github.com/kubescape/k8s-interface/workloadinterface"
)

// GetReadableID returns a readable string ID from an object's API Version, Namespace, Kind, and Name.
func GetReadableID(obj workloadinterface.IMetadata) string {
	var ID string
	if obj.GetApiVersion() != "" {
		ID += fmt.Sprintf("%s/", JoinGroupVersion(SplitApiVersion(obj.GetApiVersion())))
	}

	if obj.GetNamespace() != "" {
		ID += fmt.Sprintf("%s/", obj.GetNamespace())
	}

	ID += fmt.Sprintf("%s/%s", obj.GetKind(), obj.GetName())

	return ID
}
