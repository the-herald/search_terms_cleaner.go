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
// source: google/ads/googleads/v20/enums/media_type.proto

package enums

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

// The type of media.
type MediaTypeEnum_MediaType int32

const (
	// The media type has not been specified.
	MediaTypeEnum_UNSPECIFIED MediaTypeEnum_MediaType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	MediaTypeEnum_UNKNOWN MediaTypeEnum_MediaType = 1
	// Static image, used for image ad.
	MediaTypeEnum_IMAGE MediaTypeEnum_MediaType = 2
	// Small image, used for map ad.
	MediaTypeEnum_ICON MediaTypeEnum_MediaType = 3
	// ZIP file, used in fields of template ads.
	MediaTypeEnum_MEDIA_BUNDLE MediaTypeEnum_MediaType = 4
	// Audio file.
	MediaTypeEnum_AUDIO MediaTypeEnum_MediaType = 5
	// Video file.
	MediaTypeEnum_VIDEO MediaTypeEnum_MediaType = 6
	// Animated image, such as animated GIF.
	MediaTypeEnum_DYNAMIC_IMAGE MediaTypeEnum_MediaType = 7
)

// Enum value maps for MediaTypeEnum_MediaType.
var (
	MediaTypeEnum_MediaType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "IMAGE",
		3: "ICON",
		4: "MEDIA_BUNDLE",
		5: "AUDIO",
		6: "VIDEO",
		7: "DYNAMIC_IMAGE",
	}
	MediaTypeEnum_MediaType_value = map[string]int32{
		"UNSPECIFIED":   0,
		"UNKNOWN":       1,
		"IMAGE":         2,
		"ICON":          3,
		"MEDIA_BUNDLE":  4,
		"AUDIO":         5,
		"VIDEO":         6,
		"DYNAMIC_IMAGE": 7,
	}
)

func (x MediaTypeEnum_MediaType) Enum() *MediaTypeEnum_MediaType {
	p := new(MediaTypeEnum_MediaType)
	*p = x
	return p
}

func (x MediaTypeEnum_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaTypeEnum_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_media_type_proto_enumTypes[0].Descriptor()
}

func (MediaTypeEnum_MediaType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_media_type_proto_enumTypes[0]
}

func (x MediaTypeEnum_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaTypeEnum_MediaType.Descriptor instead.
func (MediaTypeEnum_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_media_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the types of media.
type MediaTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaTypeEnum) Reset() {
	*x = MediaTypeEnum{}
	mi := &file_google_ads_googleads_v20_enums_media_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaTypeEnum) ProtoMessage() {}

func (x *MediaTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_media_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaTypeEnum.ProtoReflect.Descriptor instead.
func (*MediaTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_media_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_media_type_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_media_type_proto_rawDesc = "" +
	"\n" +
	"/google/ads/googleads/v20/enums/media_type.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x8a\x01\n" +
	"\rMediaTypeEnum\"y\n" +
	"\tMediaType\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\t\n" +
	"\x05IMAGE\x10\x02\x12\b\n" +
	"\x04ICON\x10\x03\x12\x10\n" +
	"\fMEDIA_BUNDLE\x10\x04\x12\t\n" +
	"\x05AUDIO\x10\x05\x12\t\n" +
	"\x05VIDEO\x10\x06\x12\x11\n" +
	"\rDYNAMIC_IMAGE\x10\aB\xe8\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x0eMediaTypeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_media_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_media_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_media_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_media_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_media_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_media_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_media_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_media_type_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_media_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_media_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_media_type_proto_goTypes = []any{
	(MediaTypeEnum_MediaType)(0), // 0: google.ads.googleads.v20.enums.MediaTypeEnum.MediaType
	(*MediaTypeEnum)(nil),        // 1: google.ads.googleads.v20.enums.MediaTypeEnum
}
var file_google_ads_googleads_v20_enums_media_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_media_type_proto_init() }
func file_google_ads_googleads_v20_enums_media_type_proto_init() {
	if File_google_ads_googleads_v20_enums_media_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_media_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_media_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_media_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_media_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_media_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_media_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_media_type_proto = out.File
	file_google_ads_googleads_v20_enums_media_type_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_media_type_proto_depIdxs = nil
}
