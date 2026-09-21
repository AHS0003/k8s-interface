package cloudsupport

import (
	"testing"

	cloudsupportv1 "github.com/kubescape/k8s-interface/cloudsupport/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/version"
)

func TestSetApiServerVersion(t *testing.T) {
	tests := []struct {
		name             string
		version          *version.Info
		expectedProvider string
	}{
		{
			name: "EKS version",
			version: &version.Info{
				GitVersion: "v1.22.11-eks-400",
			},
			expectedProvider: cloudsupportv1.EKS,
		},
		{
			name: "GKE version",
			version: &version.Info{
				GitVersion: "v1.22.11-gke.400",
			},
			expectedProvider: cloudsupportv1.GKE,
		},
		{
			name: "Vanilla version (AKS)",
			version: &version.Info{
				GitVersion: "v1.22.11",
			},
			expectedProvider: "",
		},
		{
			name:             "Nil version",
			version:          nil,
			expectedProvider: "",
		},
		{
			name: "Empty version",
			version: &version.Info{
				GitVersion: "",
			},
			expectedProvider: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiServerInfo := NewApiServerInfo()
			apiServerInfo.SetApiServerVersion(tt.version)
			assert.Equal(t, tt.expectedProvider, apiServerInfo.Metadata.GetProvider())
		})
	}
}

func TestSetApiServerVersionReset(t *testing.T) {
	apiServerInfo := NewApiServerInfo()

	// Set to EKS first
	apiServerInfo.SetApiServerVersion(&version.Info{GitVersion: "v1.22.11-eks-400"})
	assert.Equal(t, cloudsupportv1.EKS, apiServerInfo.Metadata.GetProvider())

	// Then set to a vanilla version, it should reset to empty
	apiServerInfo.SetApiServerVersion(&version.Info{GitVersion: "v1.22.11"})
	assert.Equal(t, "", apiServerInfo.Metadata.GetProvider())
}
