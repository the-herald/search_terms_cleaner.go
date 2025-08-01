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
// source: google/ads/googleads/v20/errors/date_range_error.proto

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

// Enum describing possible date range errors.
type DateRangeErrorEnum_DateRangeError int32

const (
	// Enum unspecified.
	DateRangeErrorEnum_UNSPECIFIED DateRangeErrorEnum_DateRangeError = 0
	// The received error code is not known in this version.
	DateRangeErrorEnum_UNKNOWN DateRangeErrorEnum_DateRangeError = 1
	// Invalid date.
	DateRangeErrorEnum_INVALID_DATE DateRangeErrorEnum_DateRangeError = 2
	// The start date was after the end date.
	DateRangeErrorEnum_START_DATE_AFTER_END_DATE DateRangeErrorEnum_DateRangeError = 3
	// Cannot set date to past time
	DateRangeErrorEnum_CANNOT_SET_DATE_TO_PAST DateRangeErrorEnum_DateRangeError = 4
	// A date was used that is past the system "last" date.
	DateRangeErrorEnum_AFTER_MAXIMUM_ALLOWABLE_DATE DateRangeErrorEnum_DateRangeError = 5
	// Trying to change start date on a resource that has started.
	DateRangeErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED DateRangeErrorEnum_DateRangeError = 6
)

// Enum value maps for DateRangeErrorEnum_DateRangeError.
var (
	DateRangeErrorEnum_DateRangeError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_DATE",
		3: "START_DATE_AFTER_END_DATE",
		4: "CANNOT_SET_DATE_TO_PAST",
		5: "AFTER_MAXIMUM_ALLOWABLE_DATE",
		6: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
	}
	DateRangeErrorEnum_DateRangeError_value = map[string]int32{
		"UNSPECIFIED":                  0,
		"UNKNOWN":                      1,
		"INVALID_DATE":                 2,
		"START_DATE_AFTER_END_DATE":    3,
		"CANNOT_SET_DATE_TO_PAST":      4,
		"AFTER_MAXIMUM_ALLOWABLE_DATE": 5,
		"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED": 6,
	}
)

func (x DateRangeErrorEnum_DateRangeError) Enum() *DateRangeErrorEnum_DateRangeError {
	p := new(DateRangeErrorEnum_DateRangeError)
	*p = x
	return p
}

func (x DateRangeErrorEnum_DateRangeError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DateRangeErrorEnum_DateRangeError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_date_range_error_proto_enumTypes[0].Descriptor()
}

func (DateRangeErrorEnum_DateRangeError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_date_range_error_proto_enumTypes[0]
}

func (x DateRangeErrorEnum_DateRangeError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DateRangeErrorEnum_DateRangeError.Descriptor instead.
func (DateRangeErrorEnum_DateRangeError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible date range errors.
type DateRangeErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRangeErrorEnum) Reset() {
	*x = DateRangeErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_date_range_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRangeErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRangeErrorEnum) ProtoMessage() {}

func (x *DateRangeErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_date_range_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRangeErrorEnum.ProtoReflect.Descriptor instead.
func (*DateRangeErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_date_range_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_date_range_error_proto_rawDesc = "" +
	"\n" +
	"6google/ads/googleads/v20/errors/date_range_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xe6\x01\n" +
	"\x12DateRangeErrorEnum\"\xcf\x01\n" +
	"\x0eDateRangeError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x10\n" +
	"\fINVALID_DATE\x10\x02\x12\x1d\n" +
	"\x19START_DATE_AFTER_END_DATE\x10\x03\x12\x1b\n" +
	"\x17CANNOT_SET_DATE_TO_PAST\x10\x04\x12 \n" +
	"\x1cAFTER_MAXIMUM_ALLOWABLE_DATE\x10\x05\x12/\n" +
	"+CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED\x10\x06B\xf3\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x13DateRangeErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_date_range_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_date_range_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_date_range_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_date_range_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_date_range_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_date_range_error_proto_goTypes = []any{
	(DateRangeErrorEnum_DateRangeError)(0), // 0: google.ads.googleads.v20.errors.DateRangeErrorEnum.DateRangeError
	(*DateRangeErrorEnum)(nil),             // 1: google.ads.googleads.v20.errors.DateRangeErrorEnum
}
var file_google_ads_googleads_v20_errors_date_range_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_date_range_error_proto_init() }
func file_google_ads_googleads_v20_errors_date_range_error_proto_init() {
	if File_google_ads_googleads_v20_errors_date_range_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_date_range_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_date_range_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_date_range_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_date_range_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_date_range_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_date_range_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_date_range_error_proto = out.File
	file_google_ads_googleads_v20_errors_date_range_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_date_range_error_proto_depIdxs = nil
}
