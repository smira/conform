// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

// Deprecated: GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Example).DeepCopyInto(out.(*Example))
			return nil
		}, InType: reflect.TypeOf(&Example{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExampleList).DeepCopyInto(out.(*ExampleList))
			return nil
		}, InType: reflect.TypeOf(&ExampleList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExampleSpec).DeepCopyInto(out.(*ExampleSpec))
			return nil
		}, InType: reflect.TypeOf(&ExampleSpec{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExampleStatus).DeepCopyInto(out.(*ExampleStatus))
			return nil
		}, InType: reflect.TypeOf(&ExampleStatus{})},
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Example) DeepCopyInto(out *Example) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Example.
func (x *Example) DeepCopy() *Example {
	if x == nil {
		return nil
	}
	out := new(Example)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Example) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExampleList) DeepCopyInto(out *ExampleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Example, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExampleList.
func (x *ExampleList) DeepCopy() *ExampleList {
	if x == nil {
		return nil
	}
	out := new(ExampleList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ExampleList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExampleSpec) DeepCopyInto(out *ExampleSpec) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExampleSpec.
func (x *ExampleSpec) DeepCopy() *ExampleSpec {
	if x == nil {
		return nil
	}
	out := new(ExampleSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExampleStatus) DeepCopyInto(out *ExampleStatus) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ExampleStatus.
func (x *ExampleStatus) DeepCopy() *ExampleStatus {
	if x == nil {
		return nil
	}
	out := new(ExampleStatus)
	x.DeepCopyInto(out)
	return out
}
