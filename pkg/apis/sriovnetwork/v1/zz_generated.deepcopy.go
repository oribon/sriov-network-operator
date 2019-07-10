// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ByPriority) DeepCopyInto(out *ByPriority) {
	{
		in := &in
		*out = make(ByPriority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByPriority.
func (in ByPriority) DeepCopy() ByPriority {
	if in == nil {
		return nil
	}
	out := new(ByPriority)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interface) DeepCopyInto(out *Interface) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Interface) DeepCopy() *Interface {
	if in == nil {
		return nil
	}
	out := new(Interface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceExt) DeepCopyInto(out *InterfaceExt) {
	*out = *in
	out.InterfaceProperty = in.InterfaceProperty
	if in.VFs != nil {
		in, out := &in.VFs, &out.VFs
		*out = make([]VirtualFunction, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceExt.
func (in *InterfaceExt) DeepCopy() *InterfaceExt {
	if in == nil {
		return nil
	}
	out := new(InterfaceExt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceProperty) DeepCopyInto(out *InterfaceProperty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceProperty.
func (in *InterfaceProperty) DeepCopy() *InterfaceProperty {
	if in == nil {
		return nil
	}
	out := new(InterfaceProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetwork) DeepCopyInto(out *SriovNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetwork.
func (in *SriovNetwork) DeepCopy() *SriovNetwork {
	if in == nil {
		return nil
	}
	out := new(SriovNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkList) DeepCopyInto(out *SriovNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkList.
func (in *SriovNetworkList) DeepCopy() *SriovNetworkList {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNicSelector) DeepCopyInto(out *SriovNetworkNicSelector) {
	*out = *in
	if in.RootDevices != nil {
		in, out := &in.RootDevices, &out.RootDevices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PfNames != nil {
		in, out := &in.PfNames, &out.PfNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNicSelector.
func (in *SriovNetworkNicSelector) DeepCopy() *SriovNetworkNicSelector {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNicSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicy) DeepCopyInto(out *SriovNetworkNodePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicy.
func (in *SriovNetworkNodePolicy) DeepCopy() *SriovNetworkNodePolicy {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicyList) DeepCopyInto(out *SriovNetworkNodePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNetworkNodePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicyList.
func (in *SriovNetworkNodePolicyList) DeepCopy() *SriovNetworkNodePolicyList {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicySpec) DeepCopyInto(out *SriovNetworkNodePolicySpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.NicSelector.DeepCopyInto(&out.NicSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicySpec.
func (in *SriovNetworkNodePolicySpec) DeepCopy() *SriovNetworkNodePolicySpec {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicyStatus) DeepCopyInto(out *SriovNetworkNodePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicyStatus.
func (in *SriovNetworkNodePolicyStatus) DeepCopy() *SriovNetworkNodePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodeState) DeepCopyInto(out *SriovNetworkNodeState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodeState.
func (in *SriovNetworkNodeState) DeepCopy() *SriovNetworkNodeState {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodeState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodeState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodeStateList) DeepCopyInto(out *SriovNetworkNodeStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNetworkNodeState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodeStateList.
func (in *SriovNetworkNodeStateList) DeepCopy() *SriovNetworkNodeStateList {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodeStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodeStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodeStateSpec) DeepCopyInto(out *SriovNetworkNodeStateSpec) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]Interface, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodeStateSpec.
func (in *SriovNetworkNodeStateSpec) DeepCopy() *SriovNetworkNodeStateSpec {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodeStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodeStateStatus) DeepCopyInto(out *SriovNetworkNodeStateStatus) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]InterfaceExt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodeStateStatus.
func (in *SriovNetworkNodeStateStatus) DeepCopy() *SriovNetworkNodeStateStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodeStateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkSpec) DeepCopyInto(out *SriovNetworkSpec) {
	*out = *in
	if in.SpoofChk != nil {
		in, out := &in.SpoofChk, &out.SpoofChk
		*out = new(bool)
		**out = **in
	}
	if in.Trust != nil {
		in, out := &in.Trust, &out.Trust
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkSpec.
func (in *SriovNetworkSpec) DeepCopy() *SriovNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkStatus) DeepCopyInto(out *SriovNetworkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkStatus.
func (in *SriovNetworkStatus) DeepCopy() *SriovNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualFunction) DeepCopyInto(out *VirtualFunction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualFunction.
func (in *VirtualFunction) DeepCopy() *VirtualFunction {
	if in == nil {
		return nil
	}
	out := new(VirtualFunction)
	in.DeepCopyInto(out)
	return out
}
