/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	condition "github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/tls"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NovaComputeTemplate defines the input parameters specified by the user to
// create a NovaCompute via higher level CRDs.
type NovaComputeTemplate struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=0
	// Replicas of the service to run. For ironic.IronicDriver the max replica is 1
	Replicas *int32 `json:"replicas"`

	// +kubebuilder:validation:Optional
	// NodeSelector to target subset of worker nodes running this service. Setting here overrides
	// any global NodeSelector settings within the Nova CR.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	// CustomServiceConfig - customize the service config using this parameter to change service defaults,
	// or overwrite rendered information using raw OpenStack config format. The content gets added to
	// to /etc/<service>/<service>.conf.d directory as custom.conf file.
	CustomServiceConfig string `json:"customServiceConfig"`

	// +kubebuilder:validation:Optional
	// DefaultConfigOverwrite - interface to overwrite default config files like e.g. provider.yaml
	DefaultConfigOverwrite map[string]string `json:"defaultConfigOverwrite,omitempty"`

	// +kubebuilder:validation:Optional
	// Resources - Compute Resources required by this service (Limits/Requests).
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// NetworkAttachments is a list of NetworkAttachment resource names to expose the services to the given network
	NetworkAttachments []string `json:"networkAttachments,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ironic.IronicDriver;fake.FakeDriver
	// ComputeDriver - defines which driver to use for controlling virtualization
	ComputeDriver string `json:"computeDriver"`
}

// NovaComputeSpec defines the desired state of NovaCompute
type NovaComputeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Required
	// CellName is the name of the Nova Cell this NovaCompute belongs to.
	CellName string `json:"cellName"`

	// +kubebuilder:validation:Required
	// ComputeName - compute name.
	ComputeName string `json:"computeName"`

	// +kubebuilder:validation:Required
	// Secret is the name of the Secret instance containing password
	// information for the NovaCompute service. This secret is expected to be
	// generated by the nova-operator based on the information passed to the
	// Nova CR.
	Secret string `json:"secret"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default="nova"
	// ServiceUser - optional username used for this service to register in
	// keystone
	ServiceUser string `json:"serviceUser"`

	// +kubebuilder:validation:Required
	KeystoneAuthURL string `json:"keystoneAuthURL"`

	// +kubebuilder:validation:Required
	// NovaServiceBase specifies the generic fields of the service
	NovaServiceBase `json:",inline"`

	// +kubebuilder:validation:Required
	// ServiceAccount - service account name used internally to provide Nova services the default SA name
	ServiceAccount string `json:"serviceAccount"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ironic.IronicDriver;fake.FakeDriver
	// ComputeDriver defines which driver to use for controlling virtualization
	ComputeDriver string `json:"computeDriver"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// TLS - Parameters related to the TLS
	TLS tls.Ca `json:"tls,omitempty"`

	// +kubebuilder:validation:Optional
	// DefaultConfigOverwrite - interface to overwrite default config files like e.g. provider.yaml
	DefaultConfigOverwrite map[string]string `json:"defaultConfigOverwrite,omitempty"`
}

// NovaComputeStatus defines the observed state of NovaCompute
type NovaComputeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Map of hashes to track e.g. job status
	Hash map[string]string `json:"hash,omitempty"`

	// Conditions
	Conditions condition.Conditions `json:"conditions,omitempty" optional:"true"`

	// ReadyCount defines the number of replicas ready from NovaCompute
	ReadyCount int32 `json:"readyCount,omitempty"`

	// NetworkAttachments status of the deployment pods
	NetworkAttachments map[string][]string `json:"networkAttachments,omitempty"`
}

// NovaComputeCellStatus defines state of NovaCompute in cell
type NovaComputeCellStatus struct {
	// Deployed value: true means that the compute is deployed but can still be undiscovered
	Deployed bool `json:"deployed"`
	// Errors value True means that during deployment, errors appear, and the user needs to check the compute for problems
	Errors bool `json:"errors"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="NetworkAttachments",type="string",JSONPath=".spec.networkAttachments",description="NetworkAttachments"
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[0].status",description="Status"
//+kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[0].message",description="Message"

// NovaCompute is the Schema for the NovaCompute
type NovaCompute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NovaComputeSpec   `json:"spec,omitempty"`
	Status NovaComputeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NovaComputeList contains a list of NovaCompute
type NovaComputeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NovaCompute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NovaCompute{}, &NovaComputeList{})
}

// GetConditions returns the list of conditions from the status
func (s NovaComputeStatus) GetConditions() condition.Conditions {
	return s.Conditions
}

// GetSecret returns the value of the NovaCompute.Spec.Secret
func (n NovaCompute) GetSecret() string {
	return n.Spec.Secret
}

// IsReady returns true if the Cell reconciled successfully
func (instance NovaCompute) IsReady() bool {
	return instance.Status.Conditions.IsTrue(condition.ReadyCondition)
}

// NewNovaComputeSpec constructs a NewNovaComputeSpec
func NewNovaComputeSpec(
	novaCell NovaCellSpec,
	computeTemplate NovaComputeTemplate,
	novaComputeName string,
) NovaComputeSpec {
	novaComputeSpec := NovaComputeSpec{
		CellName:    novaCell.CellName,
		ComputeName: novaComputeName,
		Secret:      novaCell.Secret,
		NovaServiceBase: NovaServiceBase{
			ContainerImage:      novaCell.NovaComputeContainerImageURL,
			Replicas:            computeTemplate.Replicas,
			NodeSelector:        computeTemplate.NodeSelector,
			CustomServiceConfig: computeTemplate.CustomServiceConfig,
			Resources:           computeTemplate.Resources,
			NetworkAttachments:  computeTemplate.NetworkAttachments,
		},
		KeystoneAuthURL:        novaCell.KeystoneAuthURL,
		ServiceUser:            novaCell.ServiceUser,
		ServiceAccount:         novaCell.ServiceAccount,
		ComputeDriver:          computeTemplate.ComputeDriver,
		TLS:                    novaCell.TLS,
		DefaultConfigOverwrite: computeTemplate.DefaultConfigOverwrite,
	}
	return novaComputeSpec
}
