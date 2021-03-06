/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package federation

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	federationcommon "github.com/marun/fnord/pkg/apis/federation/common"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalFederatedCluster = builders.NewInternalResource(
		"federatedclusters",
		"FederatedCluster",
		func() runtime.Object { return &FederatedCluster{} },
		func() runtime.Object { return &FederatedClusterList{} },
	)
	InternalFederatedClusterStatus = builders.NewInternalResourceStatus(
		"federatedclusters",
		"FederatedClusterStatus",
		func() runtime.Object { return &FederatedCluster{} },
		func() runtime.Object { return &FederatedClusterList{} },
	)
	InternalFederatedReplicaSet = builders.NewInternalResource(
		"federatedreplicasets",
		"FederatedReplicaSet",
		func() runtime.Object { return &FederatedReplicaSet{} },
		func() runtime.Object { return &FederatedReplicaSetList{} },
	)
	InternalFederatedReplicaSetStatus = builders.NewInternalResourceStatus(
		"federatedreplicasets",
		"FederatedReplicaSetStatus",
		func() runtime.Object { return &FederatedReplicaSet{} },
		func() runtime.Object { return &FederatedReplicaSetList{} },
	)
	InternalFederatedReplicaSetOverride = builders.NewInternalResource(
		"federatedreplicasetoverrides",
		"FederatedReplicaSetOverride",
		func() runtime.Object { return &FederatedReplicaSetOverride{} },
		func() runtime.Object { return &FederatedReplicaSetOverrideList{} },
	)
	InternalFederatedReplicaSetOverrideStatus = builders.NewInternalResourceStatus(
		"federatedreplicasetoverrides",
		"FederatedReplicaSetOverrideStatus",
		func() runtime.Object { return &FederatedReplicaSetOverride{} },
		func() runtime.Object { return &FederatedReplicaSetOverrideList{} },
	)
	InternalFederatedSecret = builders.NewInternalResource(
		"federatedsecrets",
		"FederatedSecret",
		func() runtime.Object { return &FederatedSecret{} },
		func() runtime.Object { return &FederatedSecretList{} },
	)
	InternalFederatedSecretStatus = builders.NewInternalResourceStatus(
		"federatedsecrets",
		"FederatedSecretStatus",
		func() runtime.Object { return &FederatedSecret{} },
		func() runtime.Object { return &FederatedSecretList{} },
	)
	InternalFederatedSecretOverride = builders.NewInternalResource(
		"federatedsecretoverrides",
		"FederatedSecretOverride",
		func() runtime.Object { return &FederatedSecretOverride{} },
		func() runtime.Object { return &FederatedSecretOverrideList{} },
	)
	InternalFederatedSecretOverrideStatus = builders.NewInternalResourceStatus(
		"federatedsecretoverrides",
		"FederatedSecretOverrideStatus",
		func() runtime.Object { return &FederatedSecretOverride{} },
		func() runtime.Object { return &FederatedSecretOverrideList{} },
	)
	InternalFederationPlacement = builders.NewInternalResource(
		"federationplacements",
		"FederationPlacement",
		func() runtime.Object { return &FederationPlacement{} },
		func() runtime.Object { return &FederationPlacementList{} },
	)
	InternalFederationPlacementStatus = builders.NewInternalResourceStatus(
		"federationplacements",
		"FederationPlacementStatus",
		func() runtime.Object { return &FederationPlacement{} },
		func() runtime.Object { return &FederationPlacementList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("federation.k8s.io").WithKinds(
		InternalFederatedCluster,
		InternalFederatedClusterStatus,
		InternalFederatedReplicaSet,
		InternalFederatedReplicaSetStatus,
		InternalFederatedReplicaSetOverride,
		InternalFederatedReplicaSetOverrideStatus,
		InternalFederatedSecret,
		InternalFederatedSecretStatus,
		InternalFederatedSecretOverride,
		InternalFederatedSecretOverrideStatus,
		InternalFederationPlacement,
		InternalFederationPlacementStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretOverride struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedSecretOverrideSpec
	Status FederatedSecretOverrideStatus
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecret struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedSecretSpec
	Status FederatedSecretStatus
}

type FederatedSecretOverrideStatus struct {
}

type FederatedSecretStatus struct {
}

type FederatedSecretSpec struct {
	Template corev1.Secret
}

type FederatedSecretOverrideSpec struct {
	Overrides []FederatedSecretClusterOverride
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetOverride struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedReplicaSetOverrideSpec
	Status FederatedReplicaSetOverrideStatus
}

type FederatedSecretClusterOverride struct {
	ClusterName string
	Override    corev1.Secret
}

type FederatedReplicaSetOverrideStatus struct {
}

type FederatedReplicaSetOverrideSpec struct {
	Overrides []FederatedReplicaSetClusterOverride
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSet struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedReplicaSetSpec
	Status FederatedReplicaSetStatus
}

type FederatedReplicaSetClusterOverride struct {
	ClusterName string
	Override    appsv1.ReplicaSet
}

type FederatedReplicaSetStatus struct {
}

type FederatedReplicaSetSpec struct {
	Template appsv1.ReplicaSet
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedCluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederatedClusterSpec
	Status FederatedClusterStatus
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederationPlacement struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   FederationPlacementSpec
	Status FederationPlacementStatus
}

type FederatedClusterStatus struct {
	Conditions []ClusterCondition
	Zones      []string
	Region     string
}

type FederationPlacementStatus struct {
}

type ClusterCondition struct {
	Type               federationcommon.ClusterConditionType
	Status             corev1.ConditionStatus
	LastProbeTime      metav1.Time
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}

type FederationPlacementSpec struct {
	ResourceSelector metav1.LabelSelector
	ClusterSelector  metav1.LabelSelector
}

type FederatedClusterSpec struct {
	ClusterRef corev1.LocalObjectReference
	SecretRef  *corev1.LocalObjectReference
}

//
// FederatedCluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedClusterList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedCluster
}

func (FederatedCluster) NewStatus() interface{} {
	return FederatedClusterStatus{}
}

func (pc *FederatedCluster) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedCluster) SetStatus(s interface{}) {
	pc.Status = s.(FederatedClusterStatus)
}

func (pc *FederatedCluster) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedCluster) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedClusterSpec)
}

func (pc *FederatedCluster) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedCluster) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedCluster) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedCluster.
// +k8s:deepcopy-gen=false
type FederatedClusterRegistry interface {
	ListFederatedClusters(ctx request.Context, options *internalversion.ListOptions) (*FederatedClusterList, error)
	GetFederatedCluster(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedCluster, error)
	CreateFederatedCluster(ctx request.Context, id *FederatedCluster) (*FederatedCluster, error)
	UpdateFederatedCluster(ctx request.Context, id *FederatedCluster) (*FederatedCluster, error)
	DeleteFederatedCluster(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedClusterRegistry(sp builders.StandardStorageProvider) FederatedClusterRegistry {
	return &storageFederatedCluster{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedCluster struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedCluster) ListFederatedClusters(ctx request.Context, options *internalversion.ListOptions) (*FederatedClusterList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedClusterList), err
}

func (s *storageFederatedCluster) GetFederatedCluster(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedCluster), nil
}

func (s *storageFederatedCluster) CreateFederatedCluster(ctx request.Context, object *FederatedCluster) (*FederatedCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedCluster), nil
}

func (s *storageFederatedCluster) UpdateFederatedCluster(ctx request.Context, object *FederatedCluster) (*FederatedCluster, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedCluster), nil
}

func (s *storageFederatedCluster) DeleteFederatedCluster(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// FederatedReplicaSet Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedReplicaSet
}

func (FederatedReplicaSet) NewStatus() interface{} {
	return FederatedReplicaSetStatus{}
}

func (pc *FederatedReplicaSet) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedReplicaSet) SetStatus(s interface{}) {
	pc.Status = s.(FederatedReplicaSetStatus)
}

func (pc *FederatedReplicaSet) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedReplicaSet) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedReplicaSetSpec)
}

func (pc *FederatedReplicaSet) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedReplicaSet) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedReplicaSet) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedReplicaSet.
// +k8s:deepcopy-gen=false
type FederatedReplicaSetRegistry interface {
	ListFederatedReplicaSets(ctx request.Context, options *internalversion.ListOptions) (*FederatedReplicaSetList, error)
	GetFederatedReplicaSet(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedReplicaSet, error)
	CreateFederatedReplicaSet(ctx request.Context, id *FederatedReplicaSet) (*FederatedReplicaSet, error)
	UpdateFederatedReplicaSet(ctx request.Context, id *FederatedReplicaSet) (*FederatedReplicaSet, error)
	DeleteFederatedReplicaSet(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedReplicaSetRegistry(sp builders.StandardStorageProvider) FederatedReplicaSetRegistry {
	return &storageFederatedReplicaSet{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedReplicaSet struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedReplicaSet) ListFederatedReplicaSets(ctx request.Context, options *internalversion.ListOptions) (*FederatedReplicaSetList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSetList), err
}

func (s *storageFederatedReplicaSet) GetFederatedReplicaSet(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedReplicaSet, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSet), nil
}

func (s *storageFederatedReplicaSet) CreateFederatedReplicaSet(ctx request.Context, object *FederatedReplicaSet) (*FederatedReplicaSet, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSet), nil
}

func (s *storageFederatedReplicaSet) UpdateFederatedReplicaSet(ctx request.Context, object *FederatedReplicaSet) (*FederatedReplicaSet, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSet), nil
}

func (s *storageFederatedReplicaSet) DeleteFederatedReplicaSet(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// FederatedReplicaSetOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedReplicaSetOverrideList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedReplicaSetOverride
}

func (FederatedReplicaSetOverride) NewStatus() interface{} {
	return FederatedReplicaSetOverrideStatus{}
}

func (pc *FederatedReplicaSetOverride) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedReplicaSetOverride) SetStatus(s interface{}) {
	pc.Status = s.(FederatedReplicaSetOverrideStatus)
}

func (pc *FederatedReplicaSetOverride) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedReplicaSetOverride) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedReplicaSetOverrideSpec)
}

func (pc *FederatedReplicaSetOverride) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedReplicaSetOverride) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedReplicaSetOverride) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedReplicaSetOverride.
// +k8s:deepcopy-gen=false
type FederatedReplicaSetOverrideRegistry interface {
	ListFederatedReplicaSetOverrides(ctx request.Context, options *internalversion.ListOptions) (*FederatedReplicaSetOverrideList, error)
	GetFederatedReplicaSetOverride(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedReplicaSetOverride, error)
	CreateFederatedReplicaSetOverride(ctx request.Context, id *FederatedReplicaSetOverride) (*FederatedReplicaSetOverride, error)
	UpdateFederatedReplicaSetOverride(ctx request.Context, id *FederatedReplicaSetOverride) (*FederatedReplicaSetOverride, error)
	DeleteFederatedReplicaSetOverride(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedReplicaSetOverrideRegistry(sp builders.StandardStorageProvider) FederatedReplicaSetOverrideRegistry {
	return &storageFederatedReplicaSetOverride{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedReplicaSetOverride struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedReplicaSetOverride) ListFederatedReplicaSetOverrides(ctx request.Context, options *internalversion.ListOptions) (*FederatedReplicaSetOverrideList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSetOverrideList), err
}

func (s *storageFederatedReplicaSetOverride) GetFederatedReplicaSetOverride(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedReplicaSetOverride, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSetOverride), nil
}

func (s *storageFederatedReplicaSetOverride) CreateFederatedReplicaSetOverride(ctx request.Context, object *FederatedReplicaSetOverride) (*FederatedReplicaSetOverride, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSetOverride), nil
}

func (s *storageFederatedReplicaSetOverride) UpdateFederatedReplicaSetOverride(ctx request.Context, object *FederatedReplicaSetOverride) (*FederatedReplicaSetOverride, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedReplicaSetOverride), nil
}

func (s *storageFederatedReplicaSetOverride) DeleteFederatedReplicaSetOverride(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// FederatedSecret Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedSecret
}

func (FederatedSecret) NewStatus() interface{} {
	return FederatedSecretStatus{}
}

func (pc *FederatedSecret) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedSecret) SetStatus(s interface{}) {
	pc.Status = s.(FederatedSecretStatus)
}

func (pc *FederatedSecret) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedSecret) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedSecretSpec)
}

func (pc *FederatedSecret) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedSecret) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedSecret) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedSecret.
// +k8s:deepcopy-gen=false
type FederatedSecretRegistry interface {
	ListFederatedSecrets(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretList, error)
	GetFederatedSecret(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecret, error)
	CreateFederatedSecret(ctx request.Context, id *FederatedSecret) (*FederatedSecret, error)
	UpdateFederatedSecret(ctx request.Context, id *FederatedSecret) (*FederatedSecret, error)
	DeleteFederatedSecret(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedSecretRegistry(sp builders.StandardStorageProvider) FederatedSecretRegistry {
	return &storageFederatedSecret{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedSecret struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedSecret) ListFederatedSecrets(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretList), err
}

func (s *storageFederatedSecret) GetFederatedSecret(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) CreateFederatedSecret(ctx request.Context, object *FederatedSecret) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) UpdateFederatedSecret(ctx request.Context, object *FederatedSecret) (*FederatedSecret, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecret), nil
}

func (s *storageFederatedSecret) DeleteFederatedSecret(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// FederatedSecretOverride Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederatedSecretOverrideStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederatedSecretOverrideList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederatedSecretOverride
}

func (FederatedSecretOverride) NewStatus() interface{} {
	return FederatedSecretOverrideStatus{}
}

func (pc *FederatedSecretOverride) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederatedSecretOverride) SetStatus(s interface{}) {
	pc.Status = s.(FederatedSecretOverrideStatus)
}

func (pc *FederatedSecretOverride) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederatedSecretOverride) SetSpec(s interface{}) {
	pc.Spec = s.(FederatedSecretOverrideSpec)
}

func (pc *FederatedSecretOverride) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederatedSecretOverride) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederatedSecretOverride) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederatedSecretOverride.
// +k8s:deepcopy-gen=false
type FederatedSecretOverrideRegistry interface {
	ListFederatedSecretOverrides(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretOverrideList, error)
	GetFederatedSecretOverride(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecretOverride, error)
	CreateFederatedSecretOverride(ctx request.Context, id *FederatedSecretOverride) (*FederatedSecretOverride, error)
	UpdateFederatedSecretOverride(ctx request.Context, id *FederatedSecretOverride) (*FederatedSecretOverride, error)
	DeleteFederatedSecretOverride(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederatedSecretOverrideRegistry(sp builders.StandardStorageProvider) FederatedSecretOverrideRegistry {
	return &storageFederatedSecretOverride{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederatedSecretOverride struct {
	builders.StandardStorageProvider
}

func (s *storageFederatedSecretOverride) ListFederatedSecretOverrides(ctx request.Context, options *internalversion.ListOptions) (*FederatedSecretOverrideList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretOverrideList), err
}

func (s *storageFederatedSecretOverride) GetFederatedSecretOverride(ctx request.Context, id string, options *metav1.GetOptions) (*FederatedSecretOverride, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretOverride), nil
}

func (s *storageFederatedSecretOverride) CreateFederatedSecretOverride(ctx request.Context, object *FederatedSecretOverride) (*FederatedSecretOverride, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretOverride), nil
}

func (s *storageFederatedSecretOverride) UpdateFederatedSecretOverride(ctx request.Context, object *FederatedSecretOverride) (*FederatedSecretOverride, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederatedSecretOverride), nil
}

func (s *storageFederatedSecretOverride) DeleteFederatedSecretOverride(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// FederationPlacement Functions and Structs
//
// +k8s:deepcopy-gen=false
type FederationPlacementStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type FederationPlacementStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FederationPlacementList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []FederationPlacement
}

func (FederationPlacement) NewStatus() interface{} {
	return FederationPlacementStatus{}
}

func (pc *FederationPlacement) GetStatus() interface{} {
	return pc.Status
}

func (pc *FederationPlacement) SetStatus(s interface{}) {
	pc.Status = s.(FederationPlacementStatus)
}

func (pc *FederationPlacement) GetSpec() interface{} {
	return pc.Spec
}

func (pc *FederationPlacement) SetSpec(s interface{}) {
	pc.Spec = s.(FederationPlacementSpec)
}

func (pc *FederationPlacement) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *FederationPlacement) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc FederationPlacement) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store FederationPlacement.
// +k8s:deepcopy-gen=false
type FederationPlacementRegistry interface {
	ListFederationPlacements(ctx request.Context, options *internalversion.ListOptions) (*FederationPlacementList, error)
	GetFederationPlacement(ctx request.Context, id string, options *metav1.GetOptions) (*FederationPlacement, error)
	CreateFederationPlacement(ctx request.Context, id *FederationPlacement) (*FederationPlacement, error)
	UpdateFederationPlacement(ctx request.Context, id *FederationPlacement) (*FederationPlacement, error)
	DeleteFederationPlacement(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewFederationPlacementRegistry(sp builders.StandardStorageProvider) FederationPlacementRegistry {
	return &storageFederationPlacement{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageFederationPlacement struct {
	builders.StandardStorageProvider
}

func (s *storageFederationPlacement) ListFederationPlacements(ctx request.Context, options *internalversion.ListOptions) (*FederationPlacementList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederationPlacementList), err
}

func (s *storageFederationPlacement) GetFederationPlacement(ctx request.Context, id string, options *metav1.GetOptions) (*FederationPlacement, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*FederationPlacement), nil
}

func (s *storageFederationPlacement) CreateFederationPlacement(ctx request.Context, object *FederationPlacement) (*FederationPlacement, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*FederationPlacement), nil
}

func (s *storageFederationPlacement) UpdateFederationPlacement(ctx request.Context, object *FederationPlacement) (*FederationPlacement, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*FederationPlacement), nil
}

func (s *storageFederationPlacement) DeleteFederationPlacement(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}
