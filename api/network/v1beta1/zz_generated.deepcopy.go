//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The sealos Authors.

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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpoint) DeepCopyInto(out *ClusterEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpoint.
func (in *ClusterEndpoint) DeepCopy() *ClusterEndpoint {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointList) DeepCopyInto(out *ClusterEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointList.
func (in *ClusterEndpointList) DeepCopy() *ClusterEndpointList {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointSpec) DeepCopyInto(out *ClusterEndpointSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointSpec.
func (in *ClusterEndpointSpec) DeepCopy() *ClusterEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEndpointStatus) DeepCopyInto(out *ClusterEndpointStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEndpointStatus.
func (in *ClusterEndpointStatus) DeepCopy() *ClusterEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GRPCAction) DeepCopyInto(out *GRPCAction) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GRPCAction.
func (in *GRPCAction) DeepCopy() *GRPCAction {
	if in == nil {
		return nil
	}
	out := new(GRPCAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPGetAction) DeepCopyInto(out *HTTPGetAction) {
	*out = *in
	if in.HTTPHeaders != nil {
		in, out := &in.HTTPHeaders, &out.HTTPHeaders
		*out = make([]v1.HTTPHeader, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPGetAction.
func (in *HTTPGetAction) DeepCopy() *HTTPGetAction {
	if in == nil {
		return nil
	}
	out := new(HTTPGetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPSocket != nil {
		in, out := &in.TCPSocket, &out.TCPSocket
		*out = new(TCPSocketAction)
		**out = **in
	}
	if in.UDPSocket != nil {
		in, out := &in.UDPSocket, &out.UDPSocket
		*out = new(UDPSocketAction)
		(*in).DeepCopyInto(*out)
	}
	if in.GRPC != nil {
		in, out := &in.GRPC, &out.GRPC
		*out = new(GRPCAction)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePort) DeepCopyInto(out *ServicePort) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Handler.DeepCopyInto(&out.Handler)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort.
func (in *ServicePort) DeepCopy() *ServicePort {
	if in == nil {
		return nil
	}
	out := new(ServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSocketAction) DeepCopyInto(out *TCPSocketAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSocketAction.
func (in *TCPSocketAction) DeepCopy() *TCPSocketAction {
	if in == nil {
		return nil
	}
	out := new(TCPSocketAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPSocketAction) DeepCopyInto(out *UDPSocketAction) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]uint8, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPSocketAction.
func (in *UDPSocketAction) DeepCopy() *UDPSocketAction {
	if in == nil {
		return nil
	}
	out := new(UDPSocketAction)
	in.DeepCopyInto(out)
	return out
}
