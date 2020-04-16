// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.Iscsid":                    schema_pkg_apis_nova_v1_Iscsid(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidSpec":                schema_pkg_apis_nova_v1_IscsidSpec(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidStatus":              schema_pkg_apis_nova_v1_IscsidStatus(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.Libvirtd":                  schema_pkg_apis_nova_v1_Libvirtd(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdSpec":              schema_pkg_apis_nova_v1_LibvirtdSpec(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdStatus":            schema_pkg_apis_nova_v1_LibvirtdStatus(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaCompute":               schema_pkg_apis_nova_v1_NovaCompute(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeSpec":           schema_pkg_apis_nova_v1_NovaComputeSpec(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeStatus":         schema_pkg_apis_nova_v1_NovaComputeStatus(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTarget":       schema_pkg_apis_nova_v1_NovaMigrationTarget(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetSpec":   schema_pkg_apis_nova_v1_NovaMigrationTargetSpec(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetStatus": schema_pkg_apis_nova_v1_NovaMigrationTargetStatus(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.Virtlogd":                  schema_pkg_apis_nova_v1_Virtlogd(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdSpec":              schema_pkg_apis_nova_v1_VirtlogdSpec(ref),
		"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdStatus":            schema_pkg_apis_nova_v1_VirtlogdStatus(ref),
	}
}

func schema_pkg_apis_nova_v1_Iscsid(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Iscsid is the Schema for the iscsids API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidSpec", "github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.IscsidStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nova_v1_IscsidSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IscsidSpec defines the desired state of Iscsid",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Description: "Label is the value of the 'daemon=' label to set on a node that should run the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"iscsidImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the Docker image to run for the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "service account used to create pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"label", "iscsidImage", "serviceAccount"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_IscsidStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IscsidStatus defines the observed state of Iscsid",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Description: "Count is the number of nodes the daemon is deployed to",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"daemonsetHash": {
						SchemaProps: spec.SchemaProps{
							Description: "Daemonset hash used to detect changes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"count", "daemonsetHash"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_Libvirtd(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Libvirtd is the Schema for the libvirtds API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdSpec", "github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.LibvirtdStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nova_v1_LibvirtdSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LibvirtdSpec defines the desired state of Libvirtd",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Description: "Label is the value of the 'daemon=' label to set on a node that should run the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaLibvirtImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the Docker image to run for the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "service account used to create pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"label", "novaLibvirtImage", "serviceAccount"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_LibvirtdStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LibvirtdStatus defines the observed state of Libvirtd",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Description: "Count is the number of nodes the daemon is deployed to",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"daemonsetHash": {
						SchemaProps: spec.SchemaProps{
							Description: "Daemonset hash used to detect changes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"count", "daemonsetHash"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_NovaCompute(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaCompute is the Schema for the novacomputes API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeSpec", "github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaComputeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nova_v1_NovaComputeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaComputeSpec defines the desired state of NovaCompute",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Description: "Label is the value of the 'daemon=' label to set on a node that should run the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaComputeImage": {
						SchemaProps: spec.SchemaProps{
							Description: "container image to run for the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"publicVip": {
						SchemaProps: spec.SchemaProps{
							Description: "Control Plane public VIP String",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"internalApiVip": {
						SchemaProps: spec.SchemaProps{
							Description: "Control Plane internalAPI VIP String",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"memcacheServers": {
						SchemaProps: spec.SchemaProps{
							Description: "Memcache Servers String",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"rabbitTransportUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "RabbitMQ transport URL String",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cinderPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "Cinder API Admin Password",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "Nova API Admin Password",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaComputeCpuSharedSet": {
						SchemaProps: spec.SchemaProps{
							Description: "Mask of host CPUs that can be used for ``VCPU`` resources and offloaded emulator threads. For more information, refer to the documentation.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaComputeCpuDedicatedSet": {
						SchemaProps: spec.SchemaProps{
							Description: "A list or range of host CPU cores to which processes for pinned instance CPUs (PCPUs) can be scheduled. Ex. NovaComputeCpuDedicatedSet: [4-12,^8,15] will reserve cores from 4-12 and 15, excluding 8.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"neutronPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "Neutron API Admin Password",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"placementPassword": {
						SchemaProps: spec.SchemaProps{
							Description: "Placement API Admin Password",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "service account used to create pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"label", "novaComputeImage", "publicVip", "internalApiVip", "rabbitTransportUrl", "cinderPassword", "novaPassword", "neutronPassword", "placementPassword", "serviceAccount"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_NovaComputeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaComputeStatus defines the observed state of NovaCompute",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Description: "Count is the number of nodes the daemon is deployed to",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"daemonsetHash": {
						SchemaProps: spec.SchemaProps{
							Description: "Daemonset hash used to detect changes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"count", "daemonsetHash"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_NovaMigrationTarget(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaMigrationTarget is the Schema for the novamigrationtargets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetSpec", "github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.NovaMigrationTargetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nova_v1_NovaMigrationTargetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaMigrationTargetSpec defines the desired state of NovaMigrationTarget",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Description: "Label is the value of the 'daemon=' label to set on a node that should run the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaComputeImage": {
						SchemaProps: spec.SchemaProps{
							Description: "container image to run for the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sshdPort": {
						SchemaProps: spec.SchemaProps{
							Description: "SSHD port",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "service account used to create pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"label", "novaComputeImage", "sshdPort", "serviceAccount"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_NovaMigrationTargetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NovaMigrationTargetStatus defines the observed state of NovaMigrationTarget",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Description: "Count is the number of nodes the daemon is deployed to",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"daemonsetHash": {
						SchemaProps: spec.SchemaProps{
							Description: "Daemonset hash used to detect changes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"count", "daemonsetHash"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_Virtlogd(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Virtlogd is the Schema for the virtlogds API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdSpec", "github.com/openstack-k8s-operators/nova-operator/pkg/apis/nova/v1.VirtlogdStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_nova_v1_VirtlogdSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtlogdSpec defines the desired state of Virtlogd",
				Properties: map[string]spec.Schema{
					"label": {
						SchemaProps: spec.SchemaProps{
							Description: "Label is the value of the 'daemon=' label to set on a node that should run the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"novaLibvirtImage": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the Docker image to run for the daemon",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "service account used to create pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"label", "novaLibvirtImage", "serviceAccount"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_nova_v1_VirtlogdStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VirtlogdStatus defines the observed state of Virtlogd",
				Properties: map[string]spec.Schema{
					"count": {
						SchemaProps: spec.SchemaProps{
							Description: "Count is the number of nodes the daemon is deployed to",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"count"},
			},
		},
		Dependencies: []string{},
	}
}