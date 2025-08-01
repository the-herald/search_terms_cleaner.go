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
// source: google/ads/googleads/v20/errors/customer_user_access_error.proto

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

// Enum describing possible customer user access errors.
type CustomerUserAccessErrorEnum_CustomerUserAccessError int32

const (
	// Enum unspecified.
	CustomerUserAccessErrorEnum_UNSPECIFIED CustomerUserAccessErrorEnum_CustomerUserAccessError = 0
	// The received error code is not known in this version.
	CustomerUserAccessErrorEnum_UNKNOWN CustomerUserAccessErrorEnum_CustomerUserAccessError = 1
	// There is no user associated with the user id specified.
	CustomerUserAccessErrorEnum_INVALID_USER_ID CustomerUserAccessErrorEnum_CustomerUserAccessError = 2
	// Unable to remove the access between the user and customer.
	CustomerUserAccessErrorEnum_REMOVAL_DISALLOWED CustomerUserAccessErrorEnum_CustomerUserAccessError = 3
	// Unable to add or update the access role as specified.
	CustomerUserAccessErrorEnum_DISALLOWED_ACCESS_ROLE CustomerUserAccessErrorEnum_CustomerUserAccessError = 4
	// The user can't remove itself from an active serving customer if it's the
	// last admin user and the customer doesn't have any owner manager
	CustomerUserAccessErrorEnum_LAST_ADMIN_USER_OF_SERVING_CUSTOMER CustomerUserAccessErrorEnum_CustomerUserAccessError = 5
	// Last admin user cannot be removed from a manager.
	CustomerUserAccessErrorEnum_LAST_ADMIN_USER_OF_MANAGER CustomerUserAccessErrorEnum_CustomerUserAccessError = 6
)

// Enum value maps for CustomerUserAccessErrorEnum_CustomerUserAccessError.
var (
	CustomerUserAccessErrorEnum_CustomerUserAccessError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_USER_ID",
		3: "REMOVAL_DISALLOWED",
		4: "DISALLOWED_ACCESS_ROLE",
		5: "LAST_ADMIN_USER_OF_SERVING_CUSTOMER",
		6: "LAST_ADMIN_USER_OF_MANAGER",
	}
	CustomerUserAccessErrorEnum_CustomerUserAccessError_value = map[string]int32{
		"UNSPECIFIED":                         0,
		"UNKNOWN":                             1,
		"INVALID_USER_ID":                     2,
		"REMOVAL_DISALLOWED":                  3,
		"DISALLOWED_ACCESS_ROLE":              4,
		"LAST_ADMIN_USER_OF_SERVING_CUSTOMER": 5,
		"LAST_ADMIN_USER_OF_MANAGER":          6,
	}
)

func (x CustomerUserAccessErrorEnum_CustomerUserAccessError) Enum() *CustomerUserAccessErrorEnum_CustomerUserAccessError {
	p := new(CustomerUserAccessErrorEnum_CustomerUserAccessError)
	*p = x
	return p
}

func (x CustomerUserAccessErrorEnum_CustomerUserAccessError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerUserAccessErrorEnum_CustomerUserAccessError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_customer_user_access_error_proto_enumTypes[0].Descriptor()
}

func (CustomerUserAccessErrorEnum_CustomerUserAccessError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_customer_user_access_error_proto_enumTypes[0]
}

func (x CustomerUserAccessErrorEnum_CustomerUserAccessError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerUserAccessErrorEnum_CustomerUserAccessError.Descriptor instead.
func (CustomerUserAccessErrorEnum_CustomerUserAccessError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible CustomerUserAccess errors.
type CustomerUserAccessErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerUserAccessErrorEnum) Reset() {
	*x = CustomerUserAccessErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_customer_user_access_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerUserAccessErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUserAccessErrorEnum) ProtoMessage() {}

func (x *CustomerUserAccessErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_customer_user_access_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUserAccessErrorEnum.ProtoReflect.Descriptor instead.
func (*CustomerUserAccessErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_customer_user_access_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDesc = "" +
	"\n" +
	"@google/ads/googleads/v20/errors/customer_user_access_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xe9\x01\n" +
	"\x1bCustomerUserAccessErrorEnum\"\xc9\x01\n" +
	"\x17CustomerUserAccessError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x13\n" +
	"\x0fINVALID_USER_ID\x10\x02\x12\x16\n" +
	"\x12REMOVAL_DISALLOWED\x10\x03\x12\x1a\n" +
	"\x16DISALLOWED_ACCESS_ROLE\x10\x04\x12'\n" +
	"#LAST_ADMIN_USER_OF_SERVING_CUSTOMER\x10\x05\x12\x1e\n" +
	"\x1aLAST_ADMIN_USER_OF_MANAGER\x10\x06B\xfc\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x1cCustomerUserAccessErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_customer_user_access_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_customer_user_access_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_customer_user_access_error_proto_goTypes = []any{
	(CustomerUserAccessErrorEnum_CustomerUserAccessError)(0), // 0: google.ads.googleads.v20.errors.CustomerUserAccessErrorEnum.CustomerUserAccessError
	(*CustomerUserAccessErrorEnum)(nil),                      // 1: google.ads.googleads.v20.errors.CustomerUserAccessErrorEnum
}
var file_google_ads_googleads_v20_errors_customer_user_access_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_customer_user_access_error_proto_init() }
func file_google_ads_googleads_v20_errors_customer_user_access_error_proto_init() {
	if File_google_ads_googleads_v20_errors_customer_user_access_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_customer_user_access_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_customer_user_access_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_customer_user_access_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_customer_user_access_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_customer_user_access_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_customer_user_access_error_proto = out.File
	file_google_ads_googleads_v20_errors_customer_user_access_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_customer_user_access_error_proto_depIdxs = nil
}
