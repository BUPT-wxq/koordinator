//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Koordinator Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "k8s.io/kubernetes/pkg/scheduler/apis/config"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoschedulingArgs) DeepCopyInto(out *CoschedulingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultTimeout != nil {
		in, out := &in.DefaultTimeout, &out.DefaultTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ControllerWorkers != nil {
		in, out := &in.ControllerWorkers, &out.ControllerWorkers
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoschedulingArgs.
func (in *CoschedulingArgs) DeepCopy() *CoschedulingArgs {
	if in == nil {
		return nil
	}
	out := new(CoschedulingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoschedulingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticQuotaArgs) DeepCopyInto(out *ElasticQuotaArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DelayEvictTime != nil {
		in, out := &in.DelayEvictTime, &out.DelayEvictTime
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RevokePodInterval != nil {
		in, out := &in.RevokePodInterval, &out.RevokePodInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.DefaultQuotaGroupMax != nil {
		in, out := &in.DefaultQuotaGroupMax, &out.DefaultQuotaGroupMax
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.SystemQuotaGroupMax != nil {
		in, out := &in.SystemQuotaGroupMax, &out.SystemQuotaGroupMax
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.MonitorAllQuotas != nil {
		in, out := &in.MonitorAllQuotas, &out.MonitorAllQuotas
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticQuotaArgs.
func (in *ElasticQuotaArgs) DeepCopy() *ElasticQuotaArgs {
	if in == nil {
		return nil
	}
	out := new(ElasticQuotaArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticQuotaArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadAwareSchedulingArgs) DeepCopyInto(out *LoadAwareSchedulingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FilterExpiredNodeMetrics != nil {
		in, out := &in.FilterExpiredNodeMetrics, &out.FilterExpiredNodeMetrics
		*out = new(bool)
		**out = **in
	}
	if in.NodeMetricExpirationSeconds != nil {
		in, out := &in.NodeMetricExpirationSeconds, &out.NodeMetricExpirationSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ResourceWeights != nil {
		in, out := &in.ResourceWeights, &out.ResourceWeights
		*out = make(map[corev1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UsageThresholds != nil {
		in, out := &in.UsageThresholds, &out.UsageThresholds
		*out = make(map[corev1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProdUsageThresholds != nil {
		in, out := &in.ProdUsageThresholds, &out.ProdUsageThresholds
		*out = make(map[corev1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EstimatedScalingFactors != nil {
		in, out := &in.EstimatedScalingFactors, &out.EstimatedScalingFactors
		*out = make(map[corev1.ResourceName]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadAwareSchedulingArgs.
func (in *LoadAwareSchedulingArgs) DeepCopy() *LoadAwareSchedulingArgs {
	if in == nil {
		return nil
	}
	out := new(LoadAwareSchedulingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadAwareSchedulingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNUMAResourceArgs) DeepCopyInto(out *NodeNUMAResourceArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNUMAResourceArgs.
func (in *NodeNUMAResourceArgs) DeepCopy() *NodeNUMAResourceArgs {
	if in == nil {
		return nil
	}
	out := new(NodeNUMAResourceArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeNUMAResourceArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservationArgs) DeepCopyInto(out *ReservationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EnablePreemption != nil {
		in, out := &in.EnablePreemption, &out.EnablePreemption
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservationArgs.
func (in *ReservationArgs) DeepCopy() *ReservationArgs {
	if in == nil {
		return nil
	}
	out := new(ReservationArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReservationArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringStrategy) DeepCopyInto(out *ScoringStrategy) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]config.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringStrategy.
func (in *ScoringStrategy) DeepCopy() *ScoringStrategy {
	if in == nil {
		return nil
	}
	out := new(ScoringStrategy)
	in.DeepCopyInto(out)
	return out
}
