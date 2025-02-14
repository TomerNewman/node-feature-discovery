/*
Copyright 2025 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "sigs.k8s.io/node-feature-discovery/api/generated/clientset/versioned/scheme"
	nfdv1alpha1 "sigs.k8s.io/node-feature-discovery/api/nfd/v1alpha1"
)

// NodeFeaturesGetter has a method to return a NodeFeatureInterface.
// A group's client should implement this interface.
type NodeFeaturesGetter interface {
	NodeFeatures(namespace string) NodeFeatureInterface
}

// NodeFeatureInterface has methods to work with NodeFeature resources.
type NodeFeatureInterface interface {
	Create(ctx context.Context, nodeFeature *nfdv1alpha1.NodeFeature, opts v1.CreateOptions) (*nfdv1alpha1.NodeFeature, error)
	Update(ctx context.Context, nodeFeature *nfdv1alpha1.NodeFeature, opts v1.UpdateOptions) (*nfdv1alpha1.NodeFeature, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*nfdv1alpha1.NodeFeature, error)
	List(ctx context.Context, opts v1.ListOptions) (*nfdv1alpha1.NodeFeatureList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *nfdv1alpha1.NodeFeature, err error)
	NodeFeatureExpansion
}

// nodeFeatures implements NodeFeatureInterface
type nodeFeatures struct {
	*gentype.ClientWithList[*nfdv1alpha1.NodeFeature, *nfdv1alpha1.NodeFeatureList]
}

// newNodeFeatures returns a NodeFeatures
func newNodeFeatures(c *NfdV1alpha1Client, namespace string) *nodeFeatures {
	return &nodeFeatures{
		gentype.NewClientWithList[*nfdv1alpha1.NodeFeature, *nfdv1alpha1.NodeFeatureList](
			"nodefeatures",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *nfdv1alpha1.NodeFeature { return &nfdv1alpha1.NodeFeature{} },
			func() *nfdv1alpha1.NodeFeatureList { return &nfdv1alpha1.NodeFeatureList{} },
		),
	}
}
