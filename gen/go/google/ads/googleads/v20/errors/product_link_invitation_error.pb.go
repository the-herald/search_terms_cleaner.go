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
// source: google/ads/googleads/v20/errors/product_link_invitation_error.proto

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

// Enum describing possible product link invitation errors.
type ProductLinkInvitationErrorEnum_ProductLinkInvitationError int32

const (
	// Enum unspecified.
	ProductLinkInvitationErrorEnum_UNSPECIFIED ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 0
	// The received error code is not known in the version.
	ProductLinkInvitationErrorEnum_UNKNOWN ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 1
	// The invitation status is invalid.
	ProductLinkInvitationErrorEnum_INVALID_STATUS ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 2
	// The customer doesn't have the permission to perform this action
	ProductLinkInvitationErrorEnum_PERMISSION_DENIED ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 3
	// An invitation could not be created, since the user already has admin
	// access to the invited account. Use the ProductLinkService to directly
	// create an active link.
	ProductLinkInvitationErrorEnum_NO_INVITATION_REQUIRED ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 4
	// The customer is not permitted to create the invitation.
	ProductLinkInvitationErrorEnum_CUSTOMER_NOT_PERMITTED_TO_CREATE_INVITATION ProductLinkInvitationErrorEnum_ProductLinkInvitationError = 5
)

// Enum value maps for ProductLinkInvitationErrorEnum_ProductLinkInvitationError.
var (
	ProductLinkInvitationErrorEnum_ProductLinkInvitationError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_STATUS",
		3: "PERMISSION_DENIED",
		4: "NO_INVITATION_REQUIRED",
		5: "CUSTOMER_NOT_PERMITTED_TO_CREATE_INVITATION",
	}
	ProductLinkInvitationErrorEnum_ProductLinkInvitationError_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"INVALID_STATUS":         2,
		"PERMISSION_DENIED":      3,
		"NO_INVITATION_REQUIRED": 4,
		"CUSTOMER_NOT_PERMITTED_TO_CREATE_INVITATION": 5,
	}
)

func (x ProductLinkInvitationErrorEnum_ProductLinkInvitationError) Enum() *ProductLinkInvitationErrorEnum_ProductLinkInvitationError {
	p := new(ProductLinkInvitationErrorEnum_ProductLinkInvitationError)
	*p = x
	return p
}

func (x ProductLinkInvitationErrorEnum_ProductLinkInvitationError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductLinkInvitationErrorEnum_ProductLinkInvitationError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_enumTypes[0].Descriptor()
}

func (ProductLinkInvitationErrorEnum_ProductLinkInvitationError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_enumTypes[0]
}

func (x ProductLinkInvitationErrorEnum_ProductLinkInvitationError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductLinkInvitationErrorEnum_ProductLinkInvitationError.Descriptor instead.
func (ProductLinkInvitationErrorEnum_ProductLinkInvitationError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible product link invitation errors.
type ProductLinkInvitationErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductLinkInvitationErrorEnum) Reset() {
	*x = ProductLinkInvitationErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductLinkInvitationErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLinkInvitationErrorEnum) ProtoMessage() {}

func (x *ProductLinkInvitationErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLinkInvitationErrorEnum.ProtoReflect.Descriptor instead.
func (*ProductLinkInvitationErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_product_link_invitation_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDesc = "" +
	"\n" +
	"Cgoogle/ads/googleads/v20/errors/product_link_invitation_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xd5\x01\n" +
	"\x1eProductLinkInvitationErrorEnum\"\xb2\x01\n" +
	"\x1aProductLinkInvitationError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x12\n" +
	"\x0eINVALID_STATUS\x10\x02\x12\x15\n" +
	"\x11PERMISSION_DENIED\x10\x03\x12\x1a\n" +
	"\x16NO_INVITATION_REQUIRED\x10\x04\x12/\n" +
	"+CUSTOMER_NOT_PERMITTED_TO_CREATE_INVITATION\x10\x05B\xff\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x1fProductLinkInvitationErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_goTypes = []any{
	(ProductLinkInvitationErrorEnum_ProductLinkInvitationError)(0), // 0: google.ads.googleads.v20.errors.ProductLinkInvitationErrorEnum.ProductLinkInvitationError
	(*ProductLinkInvitationErrorEnum)(nil),                         // 1: google.ads.googleads.v20.errors.ProductLinkInvitationErrorEnum
}
var file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_init() }
func file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_init() {
	if File_google_ads_googleads_v20_errors_product_link_invitation_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_product_link_invitation_error_proto = out.File
	file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_product_link_invitation_error_proto_depIdxs = nil
}
