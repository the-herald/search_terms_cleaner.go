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
// source: google/ads/googleads/v20/resources/asset_set_asset.proto

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

// AssetSetAsset is the link between an asset and an asset set.
// Adding an AssetSetAsset links an asset with an asset set.
type AssetSetAsset struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the asset set asset.
	// Asset set asset resource names have the form:
	//
	// `customers/{customer_id}/assetSetAssets/{asset_set_id}~{asset_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The asset set which this asset set asset is linking to.
	AssetSet string `protobuf:"bytes,2,opt,name=asset_set,json=assetSet,proto3" json:"asset_set,omitempty"`
	// Immutable. The asset which this asset set asset is linking to.
	Asset string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	// Output only. The status of the asset set asset. Read-only.
	Status        enums.AssetSetAssetStatusEnum_AssetSetAssetStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.AssetSetAssetStatusEnum_AssetSetAssetStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetSetAsset) Reset() {
	*x = AssetSetAsset{}
	mi := &file_google_ads_googleads_v20_resources_asset_set_asset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetSetAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetAsset) ProtoMessage() {}

func (x *AssetSetAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_asset_set_asset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetAsset.ProtoReflect.Descriptor instead.
func (*AssetSetAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescGZIP(), []int{0}
}

func (x *AssetSetAsset) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetSetAsset) GetAssetSet() string {
	if x != nil {
		return x.AssetSet
	}
	return ""
}

func (x *AssetSetAsset) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *AssetSetAsset) GetStatus() enums.AssetSetAssetStatusEnum_AssetSetAssetStatus {
	if x != nil {
		return x.Status
	}
	return enums.AssetSetAssetStatusEnum_AssetSetAssetStatus(0)
}

var File_google_ads_googleads_v20_resources_asset_set_asset_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDesc = "" +
	"\n" +
	"8google/ads/googleads/v20/resources/asset_set_asset.proto\x12\"google.ads.googleads.v20.resources\x1a;google/ads/googleads/v20/enums/asset_set_asset_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xc3\x03\n" +
	"\rAssetSetAsset\x12S\n" +
	"\rresource_name\x18\x01 \x01(\tB.\xe0A\x05\xfaA(\n" +
	"&googleads.googleapis.com/AssetSetAssetR\fresourceName\x12F\n" +
	"\tasset_set\x18\x02 \x01(\tB)\xe0A\x05\xfaA#\n" +
	"!googleads.googleapis.com/AssetSetR\bassetSet\x12<\n" +
	"\x05asset\x18\x03 \x01(\tB&\xe0A\x05\xfaA \n" +
	"\x1egoogleads.googleapis.com/AssetR\x05asset\x12h\n" +
	"\x06status\x18\x04 \x01(\x0e2K.google.ads.googleads.v20.enums.AssetSetAssetStatusEnum.AssetSetAssetStatusB\x03\xe0A\x03R\x06status:m\xeaAj\n" +
	"&googleads.googleapis.com/AssetSetAsset\x12@customers/{customer_id}/assetSetAssets/{asset_set_id}~{asset_id}B\x84\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x12AssetSetAssetProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_asset_set_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_asset_set_asset_proto_goTypes = []any{
	(*AssetSetAsset)(nil), // 0: google.ads.googleads.v20.resources.AssetSetAsset
	(enums.AssetSetAssetStatusEnum_AssetSetAssetStatus)(0), // 1: google.ads.googleads.v20.enums.AssetSetAssetStatusEnum.AssetSetAssetStatus
}
var file_google_ads_googleads_v20_resources_asset_set_asset_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.AssetSetAsset.status:type_name -> google.ads.googleads.v20.enums.AssetSetAssetStatusEnum.AssetSetAssetStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_asset_set_asset_proto_init() }
func file_google_ads_googleads_v20_resources_asset_set_asset_proto_init() {
	if File_google_ads_googleads_v20_resources_asset_set_asset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDesc), len(file_google_ads_googleads_v20_resources_asset_set_asset_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_asset_set_asset_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_asset_set_asset_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_asset_set_asset_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_asset_set_asset_proto = out.File
	file_google_ads_googleads_v20_resources_asset_set_asset_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_asset_set_asset_proto_depIdxs = nil
}
