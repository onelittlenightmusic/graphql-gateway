// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMesh) DeepCopyInto(out *GraphqlMesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMesh.
func (in *GraphqlMesh) DeepCopy() *GraphqlMesh {
	if in == nil {
		return nil
	}
	out := new(GraphqlMesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlMesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMeshList) DeepCopyInto(out *GraphqlMeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GraphqlMesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMeshList.
func (in *GraphqlMeshList) DeepCopy() *GraphqlMeshList {
	if in == nil {
		return nil
	}
	out := new(GraphqlMeshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GraphqlMeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMeshRcConfigMap) DeepCopyInto(out *GraphqlMeshRcConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMeshRcConfigMap.
func (in *GraphqlMeshRcConfigMap) DeepCopy() *GraphqlMeshRcConfigMap {
	if in == nil {
		return nil
	}
	out := new(GraphqlMeshRcConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMeshRcSecret) DeepCopyInto(out *GraphqlMeshRcSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMeshRcSecret.
func (in *GraphqlMeshRcSecret) DeepCopy() *GraphqlMeshRcSecret {
	if in == nil {
		return nil
	}
	out := new(GraphqlMeshRcSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMeshSpec) DeepCopyInto(out *GraphqlMeshSpec) {
	*out = *in
	if in.Rc != nil {
		in, out := &in.Rc, &out.Rc
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	out.RcConfigMap = in.RcConfigMap
	out.RcSecret = in.RcSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMeshSpec.
func (in *GraphqlMeshSpec) DeepCopy() *GraphqlMeshSpec {
	if in == nil {
		return nil
	}
	out := new(GraphqlMeshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphqlMeshStatus) DeepCopyInto(out *GraphqlMeshStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphqlMeshStatus.
func (in *GraphqlMeshStatus) DeepCopy() *GraphqlMeshStatus {
	if in == nil {
		return nil
	}
	out := new(GraphqlMeshStatus)
	in.DeepCopyInto(out)
	return out
}