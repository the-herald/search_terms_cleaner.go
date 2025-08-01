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
// source: google/ads/googleads/v20/errors/internal_error.proto

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

// Enum describing possible internal errors.
type InternalErrorEnum_InternalError int32

const (
	// Enum unspecified.
	InternalErrorEnum_UNSPECIFIED InternalErrorEnum_InternalError = 0
	// The received error code is not known in this version.
	InternalErrorEnum_UNKNOWN InternalErrorEnum_InternalError = 1
	// Google Ads API encountered unexpected internal error.
	InternalErrorEnum_INTERNAL_ERROR InternalErrorEnum_InternalError = 2
	// The intended error code doesn't exist in specified API version. It will
	// be released in a future API version.
	InternalErrorEnum_ERROR_CODE_NOT_PUBLISHED InternalErrorEnum_InternalError = 3
	// Google Ads API encountered an unexpected transient error. The user
	// should retry their request in these cases.
	InternalErrorEnum_TRANSIENT_ERROR InternalErrorEnum_InternalError = 4
	// The request took longer than a deadline.
	InternalErrorEnum_DEADLINE_EXCEEDED InternalErrorEnum_InternalError = 5
)

// Enum value maps for InternalErrorEnum_InternalError.
var (
	InternalErrorEnum_InternalError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INTERNAL_ERROR",
		3: "ERROR_CODE_NOT_PUBLISHED",
		4: "TRANSIENT_ERROR",
		5: "DEADLINE_EXCEEDED",
	}
	InternalErrorEnum_InternalError_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"INTERNAL_ERROR":           2,
		"ERROR_CODE_NOT_PUBLISHED": 3,
		"TRANSIENT_ERROR":          4,
		"DEADLINE_EXCEEDED":        5,
	}
)

func (x InternalErrorEnum_InternalError) Enum() *InternalErrorEnum_InternalError {
	p := new(InternalErrorEnum_InternalError)
	*p = x
	return p
}

func (x InternalErrorEnum_InternalError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InternalErrorEnum_InternalError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_internal_error_proto_enumTypes[0].Descriptor()
}

func (InternalErrorEnum_InternalError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_internal_error_proto_enumTypes[0]
}

func (x InternalErrorEnum_InternalError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InternalErrorEnum_InternalError.Descriptor instead.
func (InternalErrorEnum_InternalError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_internal_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible internal errors.
type InternalErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalErrorEnum) Reset() {
	*x = InternalErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_internal_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalErrorEnum) ProtoMessage() {}

func (x *InternalErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_internal_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalErrorEnum.ProtoReflect.Descriptor instead.
func (*InternalErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_internal_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_internal_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_internal_error_proto_rawDesc = "" +
	"\n" +
	"4google/ads/googleads/v20/errors/internal_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xa1\x01\n" +
	"\x11InternalErrorEnum\"\x8b\x01\n" +
	"\rInternalError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x12\n" +
	"\x0eINTERNAL_ERROR\x10\x02\x12\x1c\n" +
	"\x18ERROR_CODE_NOT_PUBLISHED\x10\x03\x12\x13\n" +
	"\x0fTRANSIENT_ERROR\x10\x04\x12\x15\n" +
	"\x11DEADLINE_EXCEEDED\x10\x05B\xf2\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x12InternalErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_internal_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_internal_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_internal_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_internal_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_internal_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_internal_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_internal_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_internal_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_internal_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_internal_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_internal_error_proto_goTypes = []any{
	(InternalErrorEnum_InternalError)(0), // 0: google.ads.googleads.v20.errors.InternalErrorEnum.InternalError
	(*InternalErrorEnum)(nil),            // 1: google.ads.googleads.v20.errors.InternalErrorEnum
}
var file_google_ads_googleads_v20_errors_internal_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_internal_error_proto_init() }
func file_google_ads_googleads_v20_errors_internal_error_proto_init() {
	if File_google_ads_googleads_v20_errors_internal_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_internal_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_internal_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_internal_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_internal_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_internal_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_internal_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_internal_error_proto = out.File
	file_google_ads_googleads_v20_errors_internal_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_internal_error_proto_depIdxs = nil
}
