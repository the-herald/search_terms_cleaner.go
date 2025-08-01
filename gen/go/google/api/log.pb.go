// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: google/api/log.proto

package serviceconfig

import (
	label "google.golang.org/genproto/googleapis/api/label"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A description of a log type. Example in YAML format:
//
//     - name: library.googleapis.com/activity_history
//       description: The history of borrowing and returning library items.
//       display_name: Activity
//       labels:
//       - key: /customer_id
//         description: Identifier of a library customer
type LogDescriptor struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the log. It must be less than 512 characters long and can
	// include the following characters: upper- and lower-case alphanumeric
	// characters [A-Za-z0-9], and punctuation characters including
	// slash, underscore, hyphen, period [/_-.].
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The set of labels that are available to describe a specific log entry.
	// Runtime requests that contain labels not specified here are
	// considered invalid.
	Labels []*label.LabelDescriptor `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// A human-readable description of this log. This information appears in
	// the documentation and can contain details.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The human-readable name for this log. This information appears on
	// the user interface and should be concise.
	DisplayName   string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogDescriptor) Reset() {
	*x = LogDescriptor{}
	mi := &file_google_api_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDescriptor) ProtoMessage() {}

func (x *LogDescriptor) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDescriptor.ProtoReflect.Descriptor instead.
func (*LogDescriptor) Descriptor() ([]byte, []int) {
	return file_google_api_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogDescriptor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogDescriptor) GetLabels() []*label.LabelDescriptor {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LogDescriptor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LogDescriptor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_api_log_proto protoreflect.FileDescriptor

const file_google_api_log_proto_rawDesc = "" +
	"\n" +
	"\x14google/api/log.proto\x12\n" +
	"google.api\x1a\x16google/api/label.proto\"\x9d\x01\n" +
	"\rLogDescriptor\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x123\n" +
	"\x06labels\x18\x02 \x03(\v2\x1b.google.api.LabelDescriptorR\x06labels\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12!\n" +
	"\fdisplay_name\x18\x04 \x01(\tR\vdisplayNameBj\n" +
	"\x0ecom.google.apiB\bLogProtoP\x01ZEgoogle.golang.org/genproto/googleapis/api/serviceconfig;serviceconfig\xa2\x02\x04GAPIb\x06proto3"

var (
	file_google_api_log_proto_rawDescOnce sync.Once
	file_google_api_log_proto_rawDescData []byte
)

func file_google_api_log_proto_rawDescGZIP() []byte {
	file_google_api_log_proto_rawDescOnce.Do(func() {
		file_google_api_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_api_log_proto_rawDesc), len(file_google_api_log_proto_rawDesc)))
	})
	return file_google_api_log_proto_rawDescData
}

var file_google_api_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_api_log_proto_goTypes = []any{
	(*LogDescriptor)(nil),         // 0: google.api.LogDescriptor
	(*label.LabelDescriptor)(nil), // 1: google.api.LabelDescriptor
}
var file_google_api_log_proto_depIdxs = []int32{
	1, // 0: google.api.LogDescriptor.labels:type_name -> google.api.LabelDescriptor
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_log_proto_init() }
func file_google_api_log_proto_init() {
	if File_google_api_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_api_log_proto_rawDesc), len(file_google_api_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_log_proto_goTypes,
		DependencyIndexes: file_google_api_log_proto_depIdxs,
		MessageInfos:      file_google_api_log_proto_msgTypes,
	}.Build()
	File_google_api_log_proto = out.File
	file_google_api_log_proto_goTypes = nil
	file_google_api_log_proto_depIdxs = nil
}
