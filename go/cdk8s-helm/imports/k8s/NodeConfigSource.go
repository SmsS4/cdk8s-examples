// k8s
package k8s


// NodeConfigSource specifies a source of node configuration.
//
// Exactly one subfield (excluding metadata) must be non-nil. This API is deprecated since 1.22
type NodeConfigSource struct {
	// ConfigMap is a reference to a Node's ConfigMap.
	ConfigMap *ConfigMapNodeConfigSource `field:"optional" json:"configMap" yaml:"configMap"`
}

