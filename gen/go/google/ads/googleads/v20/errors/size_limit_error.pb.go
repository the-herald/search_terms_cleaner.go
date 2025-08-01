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
// source: google/ads/googleads/v20/errors/size_limit_error.proto

package errors

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

// Enum describing possible size limit errors.
type SizeLimitErrorEnum_SizeLimitError int32

const (
	// Enum unspecified.
	SizeLimitErrorEnum_UNSPECIFIED SizeLimitErrorEnum_SizeLimitError = 0
	// The received error code is not known in this version.
	SizeLimitErrorEnum_UNKNOWN SizeLimitErrorEnum_SizeLimitError = 1
	// The number of entries in the request exceeds the system limit, or the
	// contents of the operations exceed transaction limits due to their size
	// or complexity. Try reducing the number of entries per request.
	SizeLimitErrorEnum_REQUEST_SIZE_LIMIT_EXCEEDED SizeLimitErrorEnum_SizeLimitError = 2
	// The number of entries in the response exceeds the system limit.
	SizeLimitErrorEnum_RESPONSE_SIZE_LIMIT_EXCEEDED SizeLimitErrorEnum_SizeLimitError = 3
)

// Enum value maps for SizeLimitErrorEnum_SizeLimitError.
var (
	SizeLimitErrorEnum_SizeLimitError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "REQUEST_SIZE_LIMIT_EXCEEDED",
		3: "RESPONSE_SIZE_LIMIT_EXCEEDED",
	}
	SizeLimitErrorEnum_SizeLimitError_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"REQUEST_SIZE_LIMIT_EXCEEDED":  2,
		"RESPONSE_SIZE_LIMIT_EXCEEDED": 3,
	}
)

func (x SizeLimitErrorEnum_SizeLimitError) Enum() *SizeLimitErrorEnum_SizeLimitError {
	p := new(SizeLimitErrorEnum_SizeLimitError)
	*p = x
	return p
}

func (x SizeLimitErrorEnum_SizeLimitError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SizeLimitErrorEnum_SizeLimitError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_size_limit_error_proto_enumTypes[0].Descriptor()
}

func (SizeLimitErrorEnum_SizeLimitError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_size_limit_error_proto_enumTypes[0]
}

func (x SizeLimitErrorEnum_SizeLimitError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SizeLimitErrorEnum_SizeLimitError.Descriptor instead.
func (SizeLimitErrorEnum_SizeLimitError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible size limit errors.
type SizeLimitErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeLimitErrorEnum) Reset() {
	*x = SizeLimitErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_size_limit_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeLimitErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeLimitErrorEnum) ProtoMessage() {}

func (x *SizeLimitErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_size_limit_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeLimitErrorEnum.ProtoReflect.Descriptor instead.
func (*SizeLimitErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_size_limit_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDesc = "" +
	"\n" +
	"6google/ads/googleads/v20/errors/size_limit_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x87\x01\n" +
	"\x12SizeLimitErrorEnum\"q\n" +
	"\x0eSizeLimitError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x1f\n" +
	"\x1bREQUEST_SIZE_LIMIT_EXCEEDED\x10\x02\x12 \n" +
	"\x1cRESPONSE_SIZE_LIMIT_EXCEEDED\x10\x03B\xf3\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x13SizeLimitErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_size_limit_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_size_limit_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_size_limit_error_proto_goTypes = []any{
	(SizeLimitErrorEnum_SizeLimitError)(0), // 0: google.ads.googleads.v20.errors.SizeLimitErrorEnum.SizeLimitError
	(*SizeLimitErrorEnum)(nil),             // 1: google.ads.googleads.v20.errors.SizeLimitErrorEnum
}
var file_google_ads_googleads_v20_errors_size_limit_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_size_limit_error_proto_init() }
func file_google_ads_googleads_v20_errors_size_limit_error_proto_init() {
	if File_google_ads_googleads_v20_errors_size_limit_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_size_limit_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_size_limit_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_size_limit_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_size_limit_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_size_limit_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_size_limit_error_proto = out.File
	file_google_ads_googleads_v20_errors_size_limit_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_size_limit_error_proto_depIdxs = nil
}
