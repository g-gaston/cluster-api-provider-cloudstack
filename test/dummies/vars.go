package dummies

import (
	capcv1 "github.com/aws/cluster-api-provider-cloudstack/api/v1beta1"
	"github.com/aws/cluster-api-provider-cloudstack/pkg/cloud"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

var ( // Declare exported dummy vars.
	AffinityGroup      *cloud.AffinityGroup
	CSCluster          *capcv1.CloudStackCluster
	CAPIMachine        *capiv1.Machine
	CSMachine1         *capcv1.CloudStackMachine
	CAPICluster        *clusterv1.Cluster
	CSMachineTemplate1 *capcv1.CloudStackMachineTemplate
	Zone1              capcv1.Zone
	Zone2              capcv1.Zone
	Net1               capcv1.Network
	Net2               capcv1.Network
	DomainID           string
	Tags               map[string]string
	Tag1Key            string
	Tag1Val            string
	CSApiVersion       string
	CSClusterKind      string
	CSClusterName      string
	CSlusterNamespace  string
	TestTags           map[string]string
	CSClusterTagKey    string
	CSClusterTagVal    string
	CSClusterTag       map[string]string
	CreatedByCapcKey   string
	CreatedByCapcVal   string
)

// SetDummyVars sets/resets tag related dummy vars.
func SetTestTags() {
	TestTags = map[string]string{"TestTagKey": "TestTagValue"}
	CSClusterTagKey = "CAPC_cluster_" + string(CSCluster.ObjectMeta.UID)
	CSClusterTagVal = "1"
	CSClusterTag = map[string]string{CSClusterTagVal: CSClusterTagVal}
	CreatedByCapcKey = "create_by_CAPC"
	CreatedByCapcVal = ""
}

// SetDummyVars sets/resets all dummy vars.
func SetDummyVars() {
	// These need to be in order as they build upon eachother.
	SetDummyCAPCClusterVars()
	SetDummyCAPIClusterVars()
	SetDummyCAPIMachineVars()
	SetDummyCSMachineTemplateVars()
	SetDummyCSMachineVars()
	SetDummyTagVars()
}

// SetDummyClusterSpecVars resets the values in each of the exported CloudStackMachines related dummy variables.
func SetDummyCSMachineTemplateVars() {
	DomainID = "FakeDomainId"
	CSMachineTemplate1 = &capcv1.CloudStackMachineTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
			Kind:       "CloudStackMachineTemplate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machinetemplate-1",
			Namespace: "default",
		},
		Spec: capcv1.CloudStackMachineTemplateSpec{
			Spec: capcv1.CloudStackMachineTemplateResource{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-machinetemplateresource",
					Namespace: "default",
				},
				Spec: capcv1.CloudStackMachineSpec{
					IdentityRef: &capcv1.CloudStackIdentityReference{
						Kind: "Secret",
						Name: "IdentitySecret",
					},
					Template: "Template",
					Offering: "Offering",
					Details: map[string]string{
						"memoryOvercommitRatio": "1.2",
					},
				},
			},
		},
	}
}

// SetDummyClusterSpecVars resets the values in each of the exported CloudStackMachines related dummy variables.
func SetDummyCSMachineVars() {
	DomainID = "FakeDomainId"
	CSMachine1 = &capcv1.CloudStackMachine{
		TypeMeta: metav1.TypeMeta{
			APIVersion: CSApiVersion,
			Kind:       "CloudStackMachine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-machine-2",
			Namespace: "default",
		},
		Spec: capcv1.CloudStackMachineSpec{
			IdentityRef: &capcv1.CloudStackIdentityReference{
				Kind: "Secret",
				Name: "IdentitySecret",
			},
			InstanceID:       pointer.String("Instance1"),
			Template:         "Template",
			Offering:         "Offering",
			AffinityGroupIDs: []string{"41eeb6e4-946f-4a18-b543-b2184815f1e4"},
			Details: map[string]string{
				"memoryOvercommitRatio": "1.2",
			},
		},
	}
	CSMachine1.ObjectMeta.SetName("test-vm")
}

// SetDummyClusterSpecVars resets the values in each of the exported CloudStackCluster related dummy variables.
// It is intended to be called in BeforeEach( functions.
func SetDummyCAPCClusterVars() {
	CSApiVersion = "infrastructure.cluster.x-k8s.io/v1beta1"
	CSClusterKind = "CloudStackCluster"
	CSClusterName = "test-cluster"
	CSlusterNamespace = "default"
	AffinityGroup = &cloud.AffinityGroup{
		Name: "FakeAffinityGroup",
		Type: cloud.AffinityGroupType}
	Net1 = capcv1.Network{Name: "SharedGuestNet1"}
	Net2 = capcv1.Network{Name: "SharedGuestNet2"}
	Zone1 = capcv1.Zone{Name: "Zone1", Network: Net1}
	Zone2 = capcv1.Zone{Name: "Zone2", Network: Net2}

	CSCluster = &capcv1.CloudStackCluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: CSApiVersion,
			Kind:       CSClusterKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      CSClusterName,
			Namespace: "default",
			UID:       "0",
		},
		Spec: capcv1.CloudStackClusterSpec{
			IdentityRef: &capcv1.CloudStackIdentityReference{
				Kind: "Secret",
				Name: "IdentitySecret",
			},
			ControlPlaneEndpoint: clusterv1.APIEndpoint{Host: "EndpointHost", Port: int32(8675309)},
			Zones:                []capcv1.Zone{Zone1, Zone2},
		},
		Status: capcv1.CloudStackClusterStatus{Zones: map[string]capcv1.Zone{Zone1.ID: Zone1}},
	}
}

// SetDummyCapiCluster resets the values in each of the exported CAPICluster related dummy variables.
func SetDummyCAPIClusterVars() {
	CAPICluster = &clusterv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "capi-cluster-test-",
			Namespace:    "default",
		},
		Spec: clusterv1.ClusterSpec{
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: capcv1.GroupVersion.String(),
				Kind:       "CloudStackCluster",
				Name:       "somename",
			},
		},
	}
}

func SetDummyCAPIMachineVars() {
	CAPIMachine = &capiv1.Machine{
		Spec: capiv1.MachineSpec{FailureDomain: pointer.String(Zone1.ID)},
	}
}

// SetDummyTagVars resets the values in each of the exported Tag related dummy variables.
func SetDummyTagVars() {
	Tag1Key = "test_tag"
	Tag1Val = "arbitrary_value"
	Tags = map[string]string{Tag1Key: Tag1Val}
}
