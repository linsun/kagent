//go:build !ignore_autogenerated

/*
Copyright 2025.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenAgent) DeepCopyInto(out *AutogenAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenAgent.
func (in *AutogenAgent) DeepCopy() *AutogenAgent {
	if in == nil {
		return nil
	}
	out := new(AutogenAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenAgentList) DeepCopyInto(out *AutogenAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutogenAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenAgentList.
func (in *AutogenAgentList) DeepCopy() *AutogenAgentList {
	if in == nil {
		return nil
	}
	out := new(AutogenAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenAgentSpec) DeepCopyInto(out *AutogenAgentSpec) {
	*out = *in
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenAgentSpec.
func (in *AutogenAgentSpec) DeepCopy() *AutogenAgentSpec {
	if in == nil {
		return nil
	}
	out := new(AutogenAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenAgentStatus) DeepCopyInto(out *AutogenAgentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenAgentStatus.
func (in *AutogenAgentStatus) DeepCopy() *AutogenAgentStatus {
	if in == nil {
		return nil
	}
	out := new(AutogenAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenModelConfig) DeepCopyInto(out *AutogenModelConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenModelConfig.
func (in *AutogenModelConfig) DeepCopy() *AutogenModelConfig {
	if in == nil {
		return nil
	}
	out := new(AutogenModelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenModelConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenModelConfigList) DeepCopyInto(out *AutogenModelConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutogenModelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenModelConfigList.
func (in *AutogenModelConfigList) DeepCopy() *AutogenModelConfigList {
	if in == nil {
		return nil
	}
	out := new(AutogenModelConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenModelConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenModelConfigSpec) DeepCopyInto(out *AutogenModelConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenModelConfigSpec.
func (in *AutogenModelConfigSpec) DeepCopy() *AutogenModelConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AutogenModelConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenModelConfigStatus) DeepCopyInto(out *AutogenModelConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenModelConfigStatus.
func (in *AutogenModelConfigStatus) DeepCopy() *AutogenModelConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AutogenModelConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenTeam) DeepCopyInto(out *AutogenTeam) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenTeam.
func (in *AutogenTeam) DeepCopy() *AutogenTeam {
	if in == nil {
		return nil
	}
	out := new(AutogenTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenTeam) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenTeamList) DeepCopyInto(out *AutogenTeamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutogenTeam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenTeamList.
func (in *AutogenTeamList) DeepCopy() *AutogenTeamList {
	if in == nil {
		return nil
	}
	out := new(AutogenTeamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenTeamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenTeamSpec) DeepCopyInto(out *AutogenTeamSpec) {
	*out = *in
	if in.Participants != nil {
		in, out := &in.Participants, &out.Participants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SelectorTeamConfig = in.SelectorTeamConfig
	in.TerminationCondition.DeepCopyInto(&out.TerminationCondition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenTeamSpec.
func (in *AutogenTeamSpec) DeepCopy() *AutogenTeamSpec {
	if in == nil {
		return nil
	}
	out := new(AutogenTeamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenTeamStatus) DeepCopyInto(out *AutogenTeamStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenTeamStatus.
func (in *AutogenTeamStatus) DeepCopy() *AutogenTeamStatus {
	if in == nil {
		return nil
	}
	out := new(AutogenTeamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenTool) DeepCopyInto(out *AutogenTool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenTool.
func (in *AutogenTool) DeepCopy() *AutogenTool {
	if in == nil {
		return nil
	}
	out := new(AutogenTool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenTool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenToolList) DeepCopyInto(out *AutogenToolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutogenTool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenToolList.
func (in *AutogenToolList) DeepCopy() *AutogenToolList {
	if in == nil {
		return nil
	}
	out := new(AutogenToolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutogenToolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenToolSpec) DeepCopyInto(out *AutogenToolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenToolSpec.
func (in *AutogenToolSpec) DeepCopy() *AutogenToolSpec {
	if in == nil {
		return nil
	}
	out := new(AutogenToolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutogenToolStatus) DeepCopyInto(out *AutogenToolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutogenToolStatus.
func (in *AutogenToolStatus) DeepCopy() *AutogenToolStatus {
	if in == nil {
		return nil
	}
	out := new(AutogenToolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaxMessageTermination) DeepCopyInto(out *MaxMessageTermination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaxMessageTermination.
func (in *MaxMessageTermination) DeepCopy() *MaxMessageTermination {
	if in == nil {
		return nil
	}
	out := new(MaxMessageTermination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrTermination) DeepCopyInto(out *OrTermination) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]OrTerminationCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrTermination.
func (in *OrTermination) DeepCopy() *OrTermination {
	if in == nil {
		return nil
	}
	out := new(OrTermination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrTerminationCondition) DeepCopyInto(out *OrTerminationCondition) {
	*out = *in
	if in.MaxMessageTermination != nil {
		in, out := &in.MaxMessageTermination, &out.MaxMessageTermination
		*out = new(MaxMessageTermination)
		**out = **in
	}
	if in.TextMentionTermination != nil {
		in, out := &in.TextMentionTermination, &out.TextMentionTermination
		*out = new(TextMentionTermination)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrTerminationCondition.
func (in *OrTerminationCondition) DeepCopy() *OrTerminationCondition {
	if in == nil {
		return nil
	}
	out := new(OrTerminationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorTeamConfig) DeepCopyInto(out *SelectorTeamConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorTeamConfig.
func (in *SelectorTeamConfig) DeepCopy() *SelectorTeamConfig {
	if in == nil {
		return nil
	}
	out := new(SelectorTeamConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminationCondition) DeepCopyInto(out *TerminationCondition) {
	*out = *in
	if in.MaxMessageTermination != nil {
		in, out := &in.MaxMessageTermination, &out.MaxMessageTermination
		*out = new(MaxMessageTermination)
		**out = **in
	}
	if in.TextMentionTermination != nil {
		in, out := &in.TextMentionTermination, &out.TextMentionTermination
		*out = new(TextMentionTermination)
		**out = **in
	}
	if in.OrTermination != nil {
		in, out := &in.OrTermination, &out.OrTermination
		*out = new(OrTermination)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminationCondition.
func (in *TerminationCondition) DeepCopy() *TerminationCondition {
	if in == nil {
		return nil
	}
	out := new(TerminationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TextMentionTermination) DeepCopyInto(out *TextMentionTermination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TextMentionTermination.
func (in *TextMentionTermination) DeepCopy() *TextMentionTermination {
	if in == nil {
		return nil
	}
	out := new(TextMentionTermination)
	in.DeepCopyInto(out)
	return out
}
