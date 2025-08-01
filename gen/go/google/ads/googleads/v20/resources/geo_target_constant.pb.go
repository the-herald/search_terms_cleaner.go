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
// source: google/ads/googleads/v20/resources/geo_target_constant.proto

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

// A geo target constant.
type GeoTargetConstant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the geo target constant.
	// Geo target constant resource names have the form:
	//
	// `geoTargetConstants/{geo_target_constant_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the geo target constant.
	Id *int64 `protobuf:"varint,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. Geo target constant English name.
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The ISO-3166-1 alpha-2 country code that is associated with
	// the target.
	CountryCode *string `protobuf:"bytes,12,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
	// Output only. Geo target constant target type.
	TargetType *string `protobuf:"bytes,13,opt,name=target_type,json=targetType,proto3,oneof" json:"target_type,omitempty"`
	// Output only. Geo target constant status.
	Status enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus `protobuf:"varint,7,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus" json:"status,omitempty"`
	// Output only. The fully qualified English name, consisting of the target's
	// name and that of its parent and country.
	CanonicalName *string `protobuf:"bytes,14,opt,name=canonical_name,json=canonicalName,proto3,oneof" json:"canonical_name,omitempty"`
	// Output only. The resource name of the parent geo target constant.
	// Geo target constant resource names have the form:
	//
	// `geoTargetConstants/{parent_geo_target_constant_id}`
	ParentGeoTarget *string `protobuf:"bytes,9,opt,name=parent_geo_target,json=parentGeoTarget,proto3,oneof" json:"parent_geo_target,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GeoTargetConstant) Reset() {
	*x = GeoTargetConstant{}
	mi := &file_google_ads_googleads_v20_resources_geo_target_constant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeoTargetConstant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoTargetConstant) ProtoMessage() {}

func (x *GeoTargetConstant) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_geo_target_constant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoTargetConstant.ProtoReflect.Descriptor instead.
func (*GeoTargetConstant) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescGZIP(), []int{0}
}

func (x *GeoTargetConstant) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *GeoTargetConstant) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GeoTargetConstant) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GeoTargetConstant) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

func (x *GeoTargetConstant) GetTargetType() string {
	if x != nil && x.TargetType != nil {
		return *x.TargetType
	}
	return ""
}

func (x *GeoTargetConstant) GetStatus() enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus {
	if x != nil {
		return x.Status
	}
	return enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus(0)
}

func (x *GeoTargetConstant) GetCanonicalName() string {
	if x != nil && x.CanonicalName != nil {
		return *x.CanonicalName
	}
	return ""
}

func (x *GeoTargetConstant) GetParentGeoTarget() string {
	if x != nil && x.ParentGeoTarget != nil {
		return *x.ParentGeoTarget
	}
	return ""
}

var File_google_ads_googleads_v20_resources_geo_target_constant_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDesc = "" +
	"\n" +
	"<google/ads/googleads/v20/resources/geo_target_constant.proto\x12\"google.ads.googleads.v20.resources\x1a?google/ads/googleads/v20/enums/geo_target_constant_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xb2\x05\n" +
	"\x11GeoTargetConstant\x12W\n" +
	"\rresource_name\x18\x01 \x01(\tB2\xe0A\x03\xfaA,\n" +
	"*googleads.googleapis.com/GeoTargetConstantR\fresourceName\x12\x18\n" +
	"\x02id\x18\n" +
	" \x01(\x03B\x03\xe0A\x03H\x00R\x02id\x88\x01\x01\x12\x1c\n" +
	"\x04name\x18\v \x01(\tB\x03\xe0A\x03H\x01R\x04name\x88\x01\x01\x12+\n" +
	"\fcountry_code\x18\f \x01(\tB\x03\xe0A\x03H\x02R\vcountryCode\x88\x01\x01\x12)\n" +
	"\vtarget_type\x18\r \x01(\tB\x03\xe0A\x03H\x03R\n" +
	"targetType\x88\x01\x01\x12p\n" +
	"\x06status\x18\a \x01(\x0e2S.google.ads.googleads.v20.enums.GeoTargetConstantStatusEnum.GeoTargetConstantStatusB\x03\xe0A\x03R\x06status\x12/\n" +
	"\x0ecanonical_name\x18\x0e \x01(\tB\x03\xe0A\x03H\x04R\rcanonicalName\x88\x01\x01\x12c\n" +
	"\x11parent_geo_target\x18\t \x01(\tB2\xe0A\x03\xfaA,\n" +
	"*googleads.googleapis.com/GeoTargetConstantH\x05R\x0fparentGeoTarget\x88\x01\x01:R\xeaAO\n" +
	"*googleads.googleapis.com/GeoTargetConstant\x12!geoTargetConstants/{criterion_id}B\x05\n" +
	"\x03_idB\a\n" +
	"\x05_nameB\x0f\n" +
	"\r_country_codeB\x0e\n" +
	"\f_target_typeB\x11\n" +
	"\x0f_canonical_nameB\x14\n" +
	"\x12_parent_geo_targetB\x88\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x16GeoTargetConstantProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDesc), len(file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_geo_target_constant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_geo_target_constant_proto_goTypes = []any{
	(*GeoTargetConstant)(nil),                                      // 0: google.ads.googleads.v20.resources.GeoTargetConstant
	(enums.GeoTargetConstantStatusEnum_GeoTargetConstantStatus)(0), // 1: google.ads.googleads.v20.enums.GeoTargetConstantStatusEnum.GeoTargetConstantStatus
}
var file_google_ads_googleads_v20_resources_geo_target_constant_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.GeoTargetConstant.status:type_name -> google.ads.googleads.v20.enums.GeoTargetConstantStatusEnum.GeoTargetConstantStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_geo_target_constant_proto_init() }
func file_google_ads_googleads_v20_resources_geo_target_constant_proto_init() {
	if File_google_ads_googleads_v20_resources_geo_target_constant_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDesc), len(file_google_ads_googleads_v20_resources_geo_target_constant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_geo_target_constant_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_geo_target_constant_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_geo_target_constant_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_geo_target_constant_proto = out.File
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_geo_target_constant_proto_depIdxs = nil
}
