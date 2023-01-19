//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartsd) DeepCopyInto(out *AzureVMSchedulerStartsd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartsd.
func (in *AzureVMSchedulerStartsd) DeepCopy() *AzureVMSchedulerStartsd {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartsd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStartsd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartsdList) DeepCopyInto(out *AzureVMSchedulerStartsdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureVMSchedulerStartsd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartsdList.
func (in *AzureVMSchedulerStartsdList) DeepCopy() *AzureVMSchedulerStartsdList {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartsdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStartsdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartsdSpec) DeepCopyInto(out *AzureVMSchedulerStartsdSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartsdSpec.
func (in *AzureVMSchedulerStartsdSpec) DeepCopy() *AzureVMSchedulerStartsdSpec {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartsdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartsdStatus) DeepCopyInto(out *AzureVMSchedulerStartsdStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartsdStatus.
func (in *AzureVMSchedulerStartsdStatus) DeepCopy() *AzureVMSchedulerStartsdStatus {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartsdStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStopsd) DeepCopyInto(out *AzureVMSchedulerStopsd) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStopsd.
func (in *AzureVMSchedulerStopsd) DeepCopy() *AzureVMSchedulerStopsd {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStopsd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStopsd) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStopsdList) DeepCopyInto(out *AzureVMSchedulerStopsdList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureVMSchedulerStopsd, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStopsdList.
func (in *AzureVMSchedulerStopsdList) DeepCopy() *AzureVMSchedulerStopsdList {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStopsdList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStopsdList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStopsdSpec) DeepCopyInto(out *AzureVMSchedulerStopsdSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStopsdSpec.
func (in *AzureVMSchedulerStopsdSpec) DeepCopy() *AzureVMSchedulerStopsdSpec {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStopsdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStopsdStatus) DeepCopyInto(out *AzureVMSchedulerStopsdStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStopsdStatus.
func (in *AzureVMSchedulerStopsdStatus) DeepCopy() *AzureVMSchedulerStopsdStatus {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStopsdStatus)
	in.DeepCopyInto(out)
	return out
}