//go:build !ignore_autogenerated
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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KThreesControlPlane) DeepCopyInto(out *KThreesControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KThreesControlPlane.
func (in *KThreesControlPlane) DeepCopy() *KThreesControlPlane {
	if in == nil {
		return nil
	}
	out := new(KThreesControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KThreesControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KThreesControlPlaneList) DeepCopyInto(out *KThreesControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KThreesControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KThreesControlPlaneList.
func (in *KThreesControlPlaneList) DeepCopy() *KThreesControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(KThreesControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KThreesControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KThreesControlPlaneMachineTemplate) DeepCopyInto(out *KThreesControlPlaneMachineTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KThreesControlPlaneMachineTemplate.
func (in *KThreesControlPlaneMachineTemplate) DeepCopy() *KThreesControlPlaneMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(KThreesControlPlaneMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KThreesControlPlaneSpec) DeepCopyInto(out *KThreesControlPlaneSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.InfrastructureTemplate = in.InfrastructureTemplate
	in.KThreesConfigSpec.DeepCopyInto(&out.KThreesConfigSpec)
	if in.UpgradeAfter != nil {
		in, out := &in.UpgradeAfter, &out.UpgradeAfter
		*out = (*in).DeepCopy()
	}
	if in.NodeDrainTimeout != nil {
		in, out := &in.NodeDrainTimeout, &out.NodeDrainTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	in.MachineTemplate.DeepCopyInto(&out.MachineTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KThreesControlPlaneSpec.
func (in *KThreesControlPlaneSpec) DeepCopy() *KThreesControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(KThreesControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KThreesControlPlaneStatus) DeepCopyInto(out *KThreesControlPlaneStatus) {
	*out = *in
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KThreesControlPlaneStatus.
func (in *KThreesControlPlaneStatus) DeepCopy() *KThreesControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(KThreesControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}
