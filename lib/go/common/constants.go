
// Code generated by make build. DO NOT EDIT

package common

// Constants for node attributes
const (
	ARCH                = "si/arch"
	HostName            = "si/hostname"
	RackName            = "si/rackname"
	OS                  = "si/os"
	InstanceType        = "si/instance-type"
	FailureDomainZone   = "si/zone"
	FailureDomainRegion = "si/region"
	LocalImages         = "si/local-images"
	NodePartition       = "si/node-partition"
)

// Constants for allocation attributes
const (
	ApplicationID  = "si/application-id"
	ContainerImage = "si/container-image"
	ContainerPorts = "si/container-ports"
)
// Constants for allocation tags
const (
	// Domains
	DomainK8s      = "kubernetes.io/"
	DomainYuniKorn = "yunikorn.apache.org/"

	// Groups
	GroupMeta       = "meta/"
	GroupLabel      = "label/"
	GroupAnnotation = "annotation/"

	// Keys
	KeyPodName   = "podName"
	KeyNamespace = "namespace"
)