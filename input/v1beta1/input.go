// Package v1beta1 contains the input type for this Function
// +kubebuilder:object:generate=true
// +groupName=protection.fn.crossplane.io
// +versionName=v1beta1
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This isn't a custom resource, in the sense that we never install its CRD.
// It is a KRM-like object, so we generate a CRD to describe its schema.

// TODO: Add your input type here! It doesn't need to be called 'Input', you can
// rename it to anything you like.

// Input can be used to provide input to this Function.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:categories=crossplane
type Input struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// CacheTTL sets the time-to-live for the Function Response caching is an
	// alpha feature in Crossplane and can be deprecated or changed
	// in the future.
	// +optional
	// +kubebuilder:default:="1m"
	CacheTTL string `json:"cacheTTL,omitempty"`

	// EnableV1Mode if enabled generate v1 Crossplane Usages
	// By default v2 Usages and Cluster Usages are generated
	// Support for v1 Usages will be removed in a future version.
	// +optional
	// +kubebuilder:default:=false
	EnableV1Mode bool `json:"enableV1Mode,omitempty"`
}
