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
// source: google/ads/googleads/v20/resources/asset_field_type_view.proto

package resources

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An asset field type view.
// This view reports non-overcounted metrics for each asset field type when the
// asset is used as extension.
type AssetFieldTypeView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the asset field type view.
	// Asset field type view resource names have the form:
	//
	// `customers/{customer_id}/assetFieldTypeViews/{field_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The asset field type of the asset field type view.
	FieldType     enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,3,opt,name=field_type,json=fieldType,proto3,enum=google.ads.googleads.v20.enums.AssetFieldTypeEnum_AssetFieldType" json:"field_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetFieldTypeView) Reset() {
	*x = AssetFieldTypeView{}
	mi := &file_google_ads_googleads_v20_resources_asset_field_type_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetFieldTypeView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetFieldTypeView) ProtoMessage() {}

func (x *AssetFieldTypeView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_asset_field_type_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetFieldTypeView.ProtoReflect.Descriptor instead.
func (*AssetFieldTypeView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescGZIP(), []int{0}
}

func (x *AssetFieldTypeView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetFieldTypeView) GetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.FieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

var File_google_ads_googleads_v20_resources_asset_field_type_view_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/resources/asset_field_type_view.proto\x12\"google.ads.googleads.v20.resources\x1a5google/ads/googleads/v20/enums/asset_field_type.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xc1\x02\n" +
	"\x12AssetFieldTypeView\x12X\n" +
	"\rresource_name\x18\x01 \x01(\tB3\xe0A\x03\xfaA-\n" +
	"+googleads.googleapis.com/AssetFieldTypeViewR\fresourceName\x12e\n" +
	"\n" +
	"field_type\x18\x03 \x01(\x0e2A.google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldTypeB\x03\xe0A\x03R\tfieldType:j\xeaAg\n" +
	"+googleads.googleapis.com/AssetFieldTypeView\x128customers/{customer_id}/assetFieldTypeViews/{field_type}B\x89\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x17AssetFieldTypeViewProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_asset_field_type_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_asset_field_type_view_proto_goTypes = []any{
	(*AssetFieldTypeView)(nil),                   // 0: google.ads.googleads.v20.resources.AssetFieldTypeView
	(enums.AssetFieldTypeEnum_AssetFieldType)(0), // 1: google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
}
var file_google_ads_googleads_v20_resources_asset_field_type_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.AssetFieldTypeView.field_type:type_name -> google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_asset_field_type_view_proto_init() }
func file_google_ads_googleads_v20_resources_asset_field_type_view_proto_init() {
	if File_google_ads_googleads_v20_resources_asset_field_type_view_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_field_type_view_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_asset_field_type_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_asset_field_type_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_asset_field_type_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_asset_field_type_view_proto = out.File
	file_google_ads_googleads_v20_resources_asset_field_type_view_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_asset_field_type_view_proto_depIdxs = nil
}
