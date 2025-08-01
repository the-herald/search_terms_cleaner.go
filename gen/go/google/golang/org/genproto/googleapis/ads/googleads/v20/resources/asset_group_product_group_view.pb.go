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
// source: google/ads/googleads/v20/resources/asset_group_product_group_view.proto

package resources

import (
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

// An asset group product group view.
type AssetGroupProductGroupView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the asset group product group view.
	// Asset group product group view resource names have the form:
	//
	// `customers/{customer_id}/assetGroupProductGroupViews/{asset_group_id}~{listing_group_filter_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The asset group associated with the listing group filter.
	AssetGroup string `protobuf:"bytes,2,opt,name=asset_group,json=assetGroup,proto3" json:"asset_group,omitempty"`
	// Output only. The resource name of the asset group listing group filter.
	AssetGroupListingGroupFilter string `protobuf:"bytes,4,opt,name=asset_group_listing_group_filter,json=assetGroupListingGroupFilter,proto3" json:"asset_group_listing_group_filter,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *AssetGroupProductGroupView) Reset() {
	*x = AssetGroupProductGroupView{}
	mi := &file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetGroupProductGroupView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupProductGroupView) ProtoMessage() {}

func (x *AssetGroupProductGroupView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupProductGroupView.ProtoReflect.Descriptor instead.
func (*AssetGroupProductGroupView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescGZIP(), []int{0}
}

func (x *AssetGroupProductGroupView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetGroupProductGroupView) GetAssetGroup() string {
	if x != nil {
		return x.AssetGroup
	}
	return ""
}

func (x *AssetGroupProductGroupView) GetAssetGroupListingGroupFilter() string {
	if x != nil {
		return x.AssetGroupListingGroupFilter
	}
	return ""
}

var File_google_ads_googleads_v20_resources_asset_group_product_group_view_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDesc = "" +
	"\n" +
	"Ggoogle/ads/googleads/v20/resources/asset_group_product_group_view.proto\x12\"google.ads.googleads.v20.resources\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xf0\x03\n" +
	"\x1aAssetGroupProductGroupView\x12`\n" +
	"\rresource_name\x18\x01 \x01(\tB;\xe0A\x03\xfaA5\n" +
	"3googleads.googleapis.com/AssetGroupProductGroupViewR\fresourceName\x12L\n" +
	"\vasset_group\x18\x02 \x01(\tB+\xe0A\x03\xfaA%\n" +
	"#googleads.googleapis.com/AssetGroupR\n" +
	"assetGroup\x12\x85\x01\n" +
	" asset_group_listing_group_filter\x18\x04 \x01(\tB=\xe0A\x03\xfaA7\n" +
	"5googleads.googleapis.com/AssetGroupListingGroupFilterR\x1cassetGroupListingGroupFilter:\x99\x01\xeaA\x95\x01\n" +
	"3googleads.googleapis.com/AssetGroupProductGroupView\x12^customers/{customer_id}/assetGroupProductGroupViews/{asset_group_id}~{listing_group_filter_id}B\x91\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x1fAssetGroupProductGroupViewProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_goTypes = []any{
	(*AssetGroupProductGroupView)(nil), // 0: google.ads.googleads.v20.resources.AssetGroupProductGroupView
}
var file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_init() }
func file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_init() {
	if File_google_ads_googleads_v20_resources_asset_group_product_group_view_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_asset_group_product_group_view_proto = out.File
	file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_asset_group_product_group_view_proto_depIdxs = nil
}
