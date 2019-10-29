package initializing

// TODO generate

import (
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	trafficsplitv1alpha2 "github.com/deislabs/smi-sdk-go/pkg/apis/split/v1alpha2"
)

type Inputs struct {
	Deployments appsv1.DeploymentList
}

type Outputs struct {
	Deployments   []*appsv1.Deployment
	Services      []*corev1.Service
	TrafficSplits []*trafficsplitv1alpha2.TrafficSplit
}
