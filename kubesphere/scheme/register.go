/*
Copyright 2020 The KubeSphere Authors.

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

package scheme

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"

	appv2 "kubesphere.io/api/application/v2"
	clusterv1alpha1 "kubesphere.io/api/cluster/v1alpha1"
	corev1alpha1 "kubesphere.io/api/core/v1alpha1"
	extensionsv1alpha1 "kubesphere.io/api/extensions/v1alpha1"
	gatewayv1alpha2 "kubesphere.io/api/gateway/v1alpha2"
	iamv1beta1 "kubesphere.io/api/iam/v1beta1"
	marketplacev1alpha1 "kubesphere.io/api/marketplace/v1alpha1"
	quotav1alpha2 "kubesphere.io/api/quota/v1alpha2"
	storagev1alpha1 "kubesphere.io/api/storage/v1alpha1"
	telemetryv1alpha1 "kubesphere.io/api/telemetry/v1alpha1"
	tenantv1alpha1 "kubesphere.io/api/tenant/v1alpha1"
	tenantv1alpha2 "kubesphere.io/api/tenant/v1alpha2"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	appv2.AddToScheme,
	clusterv1alpha1.AddToScheme,
	corev1alpha1.AddToScheme,
	extensionsv1alpha1.AddToScheme,
	iamv1beta1.AddToScheme,
	quotav1alpha2.AddToScheme,
	storagev1alpha1.AddToScheme,
	tenantv1alpha1.AddToScheme,
	tenantv1alpha2.AddToScheme,
	gatewayv1alpha2.AddToScheme,
	marketplacev1alpha1.AddToScheme,
	telemetryv1alpha1.AddToScheme,
}

var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
	utilruntime.Must(k8sscheme.AddToScheme(Scheme))
}
