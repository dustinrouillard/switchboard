package v1

import (
	v1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

func init() {
	SchemeBuilder.Register(&Config{})
}

//+kubebuilder:object:root=true

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta                        `json:",inline"`
	cfg.ControllerManagerConfigurationSpec `json:",inline"`

	Selector     IngressSelector    `json:"selector"`
	Integrations IntegrationConfigs `json:"integrations"`
}

// IngressSelector can be used to limit operations to ingresses with a specific class.
type IngressSelector struct {
	IngressClass *string `json:"ingressClass,omitempty"`
}

// IntegrationConfigs describes the configurations for all integrations.
type IntegrationConfigs struct {
	ExternalDNS *ExternalDNSIntegrationConfig `json:"externalDNS"`
	CertManager *CertManagerIntegrationConfig `json:"certManager"`
}

// ExternalDNSIntegrationConfig describes the configuration for the external-dns integration.
// Exactly one of target and target IPs should be set.
type ExternalDNSIntegrationConfig struct {
	TargetService *ServiceRef `json:"targetService,omitempty"`
	TargetIPs     []string    `json:"targetIPs,omitempty"`
}

// CertManagerIntegrationConfig describes the configuration for the cert-manager integration.
type CertManagerIntegrationConfig struct {
	Template v1.Certificate `json:"certificateTemplate"`
}

// ServiceRef uniquely describes a Kubernetes service.
type ServiceRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// IssuerRef uniquely references a cert-manager issuer.
type IssuerRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
