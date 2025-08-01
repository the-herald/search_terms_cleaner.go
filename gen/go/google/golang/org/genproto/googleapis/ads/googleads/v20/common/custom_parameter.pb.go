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
// source: google/ads/googleads/v20/common/custom_parameter.proto

package common

import (
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

// A mapping that can be used by custom parameter tags in a
// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
type CustomParameter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The key matching the parameter tag name.
	Key *string `protobuf:"bytes,3,opt,name=key,proto3,oneof" json:"key,omitempty"`
	// The value to be substituted.
	Value         *string `protobuf:"bytes,4,opt,name=value,proto3,oneof" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomParameter) Reset() {
	*x = CustomParameter{}
	mi := &file_google_ads_googleads_v20_common_custom_parameter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomParameter) ProtoMessage() {}

func (x *CustomParameter) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_custom_parameter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomParameter.ProtoReflect.Descriptor instead.
func (*CustomParameter) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescGZIP(), []int{0}
}

func (x *CustomParameter) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *CustomParameter) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_google_ads_googleads_v20_common_custom_parameter_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_common_custom_parameter_proto_rawDesc = "" +
	"\n" +
	"6google/ads/googleads/v20/common/custom_parameter.proto\x12\x1fgoogle.ads.googleads.v20.common\"U\n" +
	"\x0fCustomParameter\x12\x15\n" +
	"\x03key\x18\x03 \x01(\tH\x00R\x03key\x88\x01\x01\x12\x19\n" +
	"\x05value\x18\x04 \x01(\tH\x01R\x05value\x88\x01\x01B\x06\n" +
	"\x04_keyB\b\n" +
	"\x06_valueB\xf4\x01\n" +
	"#com.google.ads.googleads.v20.commonB\x14CustomParameterProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/common;common\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Common\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Common\xea\x02#Google::Ads::GoogleAds::V20::Commonb\x06proto3"

var (
	file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_custom_parameter_proto_rawDesc), len(file_google_ads_googleads_v20_common_custom_parameter_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_common_custom_parameter_proto_rawDescData
}

var file_google_ads_googleads_v20_common_custom_parameter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_common_custom_parameter_proto_goTypes = []any{
	(*CustomParameter)(nil), // 0: google.ads.googleads.v20.common.CustomParameter
}
var file_google_ads_googleads_v20_common_custom_parameter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_common_custom_parameter_proto_init() }
func file_google_ads_googleads_v20_common_custom_parameter_proto_init() {
	if File_google_ads_googleads_v20_common_custom_parameter_proto != nil {
		return
	}
	file_google_ads_googleads_v20_common_custom_parameter_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_custom_parameter_proto_rawDesc), len(file_google_ads_googleads_v20_common_custom_parameter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_common_custom_parameter_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_common_custom_parameter_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_common_custom_parameter_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_common_custom_parameter_proto = out.File
	file_google_ads_googleads_v20_common_custom_parameter_proto_goTypes = nil
	file_google_ads_googleads_v20_common_custom_parameter_proto_depIdxs = nil
}
