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
// source: google/ads/googleads/v20/errors/request_error.proto

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

// Enum describing possible request errors.
type RequestErrorEnum_RequestError int32

const (
	// Enum unspecified.
	RequestErrorEnum_UNSPECIFIED RequestErrorEnum_RequestError = 0
	// The received error code is not known in this version.
	RequestErrorEnum_UNKNOWN RequestErrorEnum_RequestError = 1
	// Resource name is required for this request.
	RequestErrorEnum_RESOURCE_NAME_MISSING RequestErrorEnum_RequestError = 3
	// Resource name provided is malformed.
	RequestErrorEnum_RESOURCE_NAME_MALFORMED RequestErrorEnum_RequestError = 4
	// Resource name provided is malformed.
	RequestErrorEnum_BAD_RESOURCE_ID RequestErrorEnum_RequestError = 17
	// Customer ID is invalid.
	RequestErrorEnum_INVALID_CUSTOMER_ID RequestErrorEnum_RequestError = 16
	// Mutate operation should have either create, update, or remove specified.
	RequestErrorEnum_OPERATION_REQUIRED RequestErrorEnum_RequestError = 5
	// Requested resource not found.
	RequestErrorEnum_RESOURCE_NOT_FOUND RequestErrorEnum_RequestError = 6
	// Next page token specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_TOKEN RequestErrorEnum_RequestError = 7
	// Next page token specified in user request has expired.
	RequestErrorEnum_EXPIRED_PAGE_TOKEN RequestErrorEnum_RequestError = 8
	// Page size specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_SIZE RequestErrorEnum_RequestError = 22
	// Setting the page size is not supported, and will be unavailable in a
	// future version.
	RequestErrorEnum_PAGE_SIZE_NOT_SUPPORTED RequestErrorEnum_RequestError = 40
	// Required field is missing.
	RequestErrorEnum_REQUIRED_FIELD_MISSING RequestErrorEnum_RequestError = 9
	// The field cannot be modified because it's immutable. It's also possible
	// that the field can be modified using 'create' operation but not 'update'.
	RequestErrorEnum_IMMUTABLE_FIELD RequestErrorEnum_RequestError = 11
	// Received too many entries in request.
	RequestErrorEnum_TOO_MANY_MUTATE_OPERATIONS RequestErrorEnum_RequestError = 13
	// Request cannot be executed by a manager account.
	RequestErrorEnum_CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT RequestErrorEnum_RequestError = 14
	// Mutate request was attempting to modify a readonly field.
	// For instance, Budget fields can be requested for Ad Group,
	// but are read-only for adGroups:mutate.
	RequestErrorEnum_CANNOT_MODIFY_FOREIGN_FIELD RequestErrorEnum_RequestError = 15
	// Enum value is not permitted.
	RequestErrorEnum_INVALID_ENUM_VALUE RequestErrorEnum_RequestError = 18
	// The developer-token parameter is required for all requests.
	RequestErrorEnum_DEVELOPER_TOKEN_PARAMETER_MISSING RequestErrorEnum_RequestError = 19
	// The login-customer-id parameter is required for this request.
	RequestErrorEnum_LOGIN_CUSTOMER_ID_PARAMETER_MISSING RequestErrorEnum_RequestError = 20
	// page_token is set in the validate only request
	RequestErrorEnum_VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN RequestErrorEnum_RequestError = 21
	// return_summary_row cannot be enabled if request did not select any
	// metrics field.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS RequestErrorEnum_RequestError = 29
	// return_summary_row should not be enabled for validate only requests.
	RequestErrorEnum_CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS RequestErrorEnum_RequestError = 30
	// return_summary_row parameter value should be the same between requests
	// with page_token field set and their original request.
	RequestErrorEnum_INCONSISTENT_RETURN_SUMMARY_ROW_VALUE RequestErrorEnum_RequestError = 31
	// The total results count cannot be returned if it was not requested in the
	// original request.
	RequestErrorEnum_TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED RequestErrorEnum_RequestError = 32
	// Deadline specified by the client was too short.
	RequestErrorEnum_RPC_DEADLINE_TOO_SHORT RequestErrorEnum_RequestError = 33
	// This API version has been sunset and is no longer supported.
	RequestErrorEnum_UNSUPPORTED_VERSION RequestErrorEnum_RequestError = 38
	// The Google Cloud project in the request was not found.
	RequestErrorEnum_CLOUD_PROJECT_NOT_FOUND RequestErrorEnum_RequestError = 39
)

// Enum value maps for RequestErrorEnum_RequestError.
var (
	RequestErrorEnum_RequestError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "RESOURCE_NAME_MISSING",
		4:  "RESOURCE_NAME_MALFORMED",
		17: "BAD_RESOURCE_ID",
		16: "INVALID_CUSTOMER_ID",
		5:  "OPERATION_REQUIRED",
		6:  "RESOURCE_NOT_FOUND",
		7:  "INVALID_PAGE_TOKEN",
		8:  "EXPIRED_PAGE_TOKEN",
		22: "INVALID_PAGE_SIZE",
		40: "PAGE_SIZE_NOT_SUPPORTED",
		9:  "REQUIRED_FIELD_MISSING",
		11: "IMMUTABLE_FIELD",
		13: "TOO_MANY_MUTATE_OPERATIONS",
		14: "CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT",
		15: "CANNOT_MODIFY_FOREIGN_FIELD",
		18: "INVALID_ENUM_VALUE",
		19: "DEVELOPER_TOKEN_PARAMETER_MISSING",
		20: "LOGIN_CUSTOMER_ID_PARAMETER_MISSING",
		21: "VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN",
		29: "CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS",
		30: "CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS",
		31: "INCONSISTENT_RETURN_SUMMARY_ROW_VALUE",
		32: "TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED",
		33: "RPC_DEADLINE_TOO_SHORT",
		38: "UNSUPPORTED_VERSION",
		39: "CLOUD_PROJECT_NOT_FOUND",
	}
	RequestErrorEnum_RequestError_value = map[string]int32{
		"UNSPECIFIED":                                           0,
		"UNKNOWN":                                               1,
		"RESOURCE_NAME_MISSING":                                 3,
		"RESOURCE_NAME_MALFORMED":                               4,
		"BAD_RESOURCE_ID":                                       17,
		"INVALID_CUSTOMER_ID":                                   16,
		"OPERATION_REQUIRED":                                    5,
		"RESOURCE_NOT_FOUND":                                    6,
		"INVALID_PAGE_TOKEN":                                    7,
		"EXPIRED_PAGE_TOKEN":                                    8,
		"INVALID_PAGE_SIZE":                                     22,
		"PAGE_SIZE_NOT_SUPPORTED":                               40,
		"REQUIRED_FIELD_MISSING":                                9,
		"IMMUTABLE_FIELD":                                       11,
		"TOO_MANY_MUTATE_OPERATIONS":                            13,
		"CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT":                 14,
		"CANNOT_MODIFY_FOREIGN_FIELD":                           15,
		"INVALID_ENUM_VALUE":                                    18,
		"DEVELOPER_TOKEN_PARAMETER_MISSING":                     19,
		"LOGIN_CUSTOMER_ID_PARAMETER_MISSING":                   20,
		"VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN":                  21,
		"CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS": 29,
		"CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS":  30,
		"INCONSISTENT_RETURN_SUMMARY_ROW_VALUE":                 31,
		"TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED":          32,
		"RPC_DEADLINE_TOO_SHORT":                                33,
		"UNSUPPORTED_VERSION":                                   38,
		"CLOUD_PROJECT_NOT_FOUND":                               39,
	}
)

func (x RequestErrorEnum_RequestError) Enum() *RequestErrorEnum_RequestError {
	p := new(RequestErrorEnum_RequestError)
	*p = x
	return p
}

func (x RequestErrorEnum_RequestError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestErrorEnum_RequestError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_request_error_proto_enumTypes[0].Descriptor()
}

func (RequestErrorEnum_RequestError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_request_error_proto_enumTypes[0]
}

func (x RequestErrorEnum_RequestError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestErrorEnum_RequestError.Descriptor instead.
func (RequestErrorEnum_RequestError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_request_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible request errors.
type RequestErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestErrorEnum) Reset() {
	*x = RequestErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_request_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestErrorEnum) ProtoMessage() {}

func (x *RequestErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_request_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestErrorEnum.ProtoReflect.Descriptor instead.
func (*RequestErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_request_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_request_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_request_error_proto_rawDesc = "" +
	"\n" +
	"3google/ads/googleads/v20/errors/request_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x8e\a\n" +
	"\x10RequestErrorEnum\"\xf9\x06\n" +
	"\fRequestError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x19\n" +
	"\x15RESOURCE_NAME_MISSING\x10\x03\x12\x1b\n" +
	"\x17RESOURCE_NAME_MALFORMED\x10\x04\x12\x13\n" +
	"\x0fBAD_RESOURCE_ID\x10\x11\x12\x17\n" +
	"\x13INVALID_CUSTOMER_ID\x10\x10\x12\x16\n" +
	"\x12OPERATION_REQUIRED\x10\x05\x12\x16\n" +
	"\x12RESOURCE_NOT_FOUND\x10\x06\x12\x16\n" +
	"\x12INVALID_PAGE_TOKEN\x10\a\x12\x16\n" +
	"\x12EXPIRED_PAGE_TOKEN\x10\b\x12\x15\n" +
	"\x11INVALID_PAGE_SIZE\x10\x16\x12\x1b\n" +
	"\x17PAGE_SIZE_NOT_SUPPORTED\x10(\x12\x1a\n" +
	"\x16REQUIRED_FIELD_MISSING\x10\t\x12\x13\n" +
	"\x0fIMMUTABLE_FIELD\x10\v\x12\x1e\n" +
	"\x1aTOO_MANY_MUTATE_OPERATIONS\x10\r\x12)\n" +
	"%CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT\x10\x0e\x12\x1f\n" +
	"\x1bCANNOT_MODIFY_FOREIGN_FIELD\x10\x0f\x12\x16\n" +
	"\x12INVALID_ENUM_VALUE\x10\x12\x12%\n" +
	"!DEVELOPER_TOKEN_PARAMETER_MISSING\x10\x13\x12'\n" +
	"#LOGIN_CUSTOMER_ID_PARAMETER_MISSING\x10\x14\x12(\n" +
	"$VALIDATE_ONLY_REQUEST_HAS_PAGE_TOKEN\x10\x15\x129\n" +
	"5CANNOT_RETURN_SUMMARY_ROW_FOR_REQUEST_WITHOUT_METRICS\x10\x1d\x128\n" +
	"4CANNOT_RETURN_SUMMARY_ROW_FOR_VALIDATE_ONLY_REQUESTS\x10\x1e\x12)\n" +
	"%INCONSISTENT_RETURN_SUMMARY_ROW_VALUE\x10\x1f\x120\n" +
	",TOTAL_RESULTS_COUNT_NOT_ORIGINALLY_REQUESTED\x10 \x12\x1a\n" +
	"\x16RPC_DEADLINE_TOO_SHORT\x10!\x12\x17\n" +
	"\x13UNSUPPORTED_VERSION\x10&\x12\x1b\n" +
	"\x17CLOUD_PROJECT_NOT_FOUND\x10'B\xf1\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x11RequestErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_request_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_request_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_request_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_request_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_request_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_request_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_request_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_request_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_request_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_request_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_request_error_proto_goTypes = []any{
	(RequestErrorEnum_RequestError)(0), // 0: google.ads.googleads.v20.errors.RequestErrorEnum.RequestError
	(*RequestErrorEnum)(nil),           // 1: google.ads.googleads.v20.errors.RequestErrorEnum
}
var file_google_ads_googleads_v20_errors_request_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_request_error_proto_init() }
func file_google_ads_googleads_v20_errors_request_error_proto_init() {
	if File_google_ads_googleads_v20_errors_request_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_request_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_request_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_request_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_request_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_request_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_request_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_request_error_proto = out.File
	file_google_ads_googleads_v20_errors_request_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_request_error_proto_depIdxs = nil
}
