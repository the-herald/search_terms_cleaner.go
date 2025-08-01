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
// source: google/ads/googleads/v20/resources/carrier_constant.proto

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

// A carrier criterion that can be used in campaign targeting.
type CarrierConstant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the carrier criterion.
	// Carrier criterion resource names have the form:
	//
	// `carrierConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the carrier criterion.
	Id *int64 `protobuf:"varint,5,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The full name of the carrier in English.
	Name *string `protobuf:"bytes,6,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The country code of the country where the carrier is located,
	// for example, "AR", "FR", etc.
	CountryCode   *string `protobuf:"bytes,7,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CarrierConstant) Reset() {
	*x = CarrierConstant{}
	mi := &file_google_ads_googleads_v20_resources_carrier_constant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CarrierConstant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarrierConstant) ProtoMessage() {}

func (x *CarrierConstant) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_carrier_constant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarrierConstant.ProtoReflect.Descriptor instead.
func (*CarrierConstant) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescGZIP(), []int{0}
}

func (x *CarrierConstant) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CarrierConstant) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *CarrierConstant) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CarrierConstant) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

var File_google_ads_googleads_v20_resources_carrier_constant_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDesc = "" +
	"\n" +
	"9google/ads/googleads/v20/resources/carrier_constant.proto\x12\"google.ads.googleads.v20.resources\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xbe\x02\n" +
	"\x0fCarrierConstant\x12U\n" +
	"\rresource_name\x18\x01 \x01(\tB0\xe0A\x03\xfaA*\n" +
	"(googleads.googleapis.com/CarrierConstantR\fresourceName\x12\x18\n" +
	"\x02id\x18\x05 \x01(\x03B\x03\xe0A\x03H\x00R\x02id\x88\x01\x01\x12\x1c\n" +
	"\x04name\x18\x06 \x01(\tB\x03\xe0A\x03H\x01R\x04name\x88\x01\x01\x12+\n" +
	"\fcountry_code\x18\a \x01(\tB\x03\xe0A\x03H\x02R\vcountryCode\x88\x01\x01:N\xeaAK\n" +
	"(googleads.googleapis.com/CarrierConstant\x12\x1fcarrierConstants/{criterion_id}B\x05\n" +
	"\x03_idB\a\n" +
	"\x05_nameB\x0f\n" +
	"\r_country_codeB\x86\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x14CarrierConstantProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDesc), len(file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_carrier_constant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_carrier_constant_proto_goTypes = []any{
	(*CarrierConstant)(nil), // 0: google.ads.googleads.v20.resources.CarrierConstant
}
var file_google_ads_googleads_v20_resources_carrier_constant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_carrier_constant_proto_init() }
func file_google_ads_googleads_v20_resources_carrier_constant_proto_init() {
	if File_google_ads_googleads_v20_resources_carrier_constant_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_carrier_constant_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDesc), len(file_google_ads_googleads_v20_resources_carrier_constant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_carrier_constant_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_carrier_constant_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_carrier_constant_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_carrier_constant_proto = out.File
	file_google_ads_googleads_v20_resources_carrier_constant_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_carrier_constant_proto_depIdxs = nil
}
