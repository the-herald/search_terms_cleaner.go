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
// source: google/ads/googleads/v20/services/user_data_service.proto

package services

import (
	common "google.golang.org/genproto/googleapis/ads/googleads/v20/common"
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

// Request message for
// [UserDataService.UploadUserData][google.ads.googleads.v20.services.UserDataService.UploadUserData]
type UploadUserDataRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer for which to update the user data.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to be done.
	Operations []*UserDataOperation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// Metadata of the request.
	//
	// Types that are valid to be assigned to Metadata:
	//
	//	*UploadUserDataRequest_CustomerMatchUserListMetadata
	Metadata      isUploadUserDataRequest_Metadata `protobuf_oneof:"metadata"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadUserDataRequest) Reset() {
	*x = UploadUserDataRequest{}
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUserDataRequest) ProtoMessage() {}

func (x *UploadUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUserDataRequest.ProtoReflect.Descriptor instead.
func (*UploadUserDataRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_user_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadUserDataRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *UploadUserDataRequest) GetOperations() []*UserDataOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *UploadUserDataRequest) GetMetadata() isUploadUserDataRequest_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UploadUserDataRequest) GetCustomerMatchUserListMetadata() *common.CustomerMatchUserListMetadata {
	if x != nil {
		if x, ok := x.Metadata.(*UploadUserDataRequest_CustomerMatchUserListMetadata); ok {
			return x.CustomerMatchUserListMetadata
		}
	}
	return nil
}

type isUploadUserDataRequest_Metadata interface {
	isUploadUserDataRequest_Metadata()
}

type UploadUserDataRequest_CustomerMatchUserListMetadata struct {
	// Metadata for data updates to a Customer Match user list.
	CustomerMatchUserListMetadata *common.CustomerMatchUserListMetadata `protobuf:"bytes,2,opt,name=customer_match_user_list_metadata,json=customerMatchUserListMetadata,proto3,oneof"`
}

func (*UploadUserDataRequest_CustomerMatchUserListMetadata) isUploadUserDataRequest_Metadata() {}

// Operation to be made for the UploadUserDataRequest.
type UserDataOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Operation to be made for the UploadUserDataRequest.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*UserDataOperation_Create
	//	*UserDataOperation_Remove
	Operation     isUserDataOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserDataOperation) Reset() {
	*x = UserDataOperation{}
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserDataOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataOperation) ProtoMessage() {}

func (x *UserDataOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataOperation.ProtoReflect.Descriptor instead.
func (*UserDataOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_user_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserDataOperation) GetOperation() isUserDataOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *UserDataOperation) GetCreate() *common.UserData {
	if x != nil {
		if x, ok := x.Operation.(*UserDataOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *UserDataOperation) GetRemove() *common.UserData {
	if x != nil {
		if x, ok := x.Operation.(*UserDataOperation_Remove); ok {
			return x.Remove
		}
	}
	return nil
}

type isUserDataOperation_Operation interface {
	isUserDataOperation_Operation()
}

type UserDataOperation_Create struct {
	// The list of user data to be appended to the user list.
	Create *common.UserData `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type UserDataOperation_Remove struct {
	// The list of user data to be removed from the user list.
	Remove *common.UserData `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*UserDataOperation_Create) isUserDataOperation_Operation() {}

func (*UserDataOperation_Remove) isUserDataOperation_Operation() {}

// Response message for
// [UserDataService.UploadUserData][google.ads.googleads.v20.services.UserDataService.UploadUserData]
// Uploads made through this service will not be visible under the 'Segment
// members' section for the Customer Match List in the Google Ads UI.
type UploadUserDataResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The date time at which the request was received by API, formatted as
	// "yyyy-mm-dd hh:mm:ss+|-hh:mm", for example, "2019-01-01 12:32:45-08:00".
	UploadDateTime *string `protobuf:"bytes,3,opt,name=upload_date_time,json=uploadDateTime,proto3,oneof" json:"upload_date_time,omitempty"`
	// Number of upload data operations received by API.
	ReceivedOperationsCount *int32 `protobuf:"varint,4,opt,name=received_operations_count,json=receivedOperationsCount,proto3,oneof" json:"received_operations_count,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *UploadUserDataResponse) Reset() {
	*x = UploadUserDataResponse{}
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadUserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUserDataResponse) ProtoMessage() {}

func (x *UploadUserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUserDataResponse.ProtoReflect.Descriptor instead.
func (*UploadUserDataResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_user_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadUserDataResponse) GetUploadDateTime() string {
	if x != nil && x.UploadDateTime != nil {
		return *x.UploadDateTime
	}
	return ""
}

func (x *UploadUserDataResponse) GetReceivedOperationsCount() int32 {
	if x != nil && x.ReceivedOperationsCount != nil {
		return *x.ReceivedOperationsCount
	}
	return 0
}

var File_google_ads_googleads_v20_services_user_data_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_user_data_service_proto_rawDesc = "" +
	"\n" +
	"9google/ads/googleads/v20/services/user_data_service.proto\x12!google.ads.googleads.v20.services\x1a7google/ads/googleads/v20/common/offline_user_data.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\"\xb1\x02\n" +
	"\x15UploadUserDataRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12Y\n" +
	"\n" +
	"operations\x18\x03 \x03(\v24.google.ads.googleads.v20.services.UserDataOperationB\x03\xe0A\x02R\n" +
	"operations\x12\x8a\x01\n" +
	"!customer_match_user_list_metadata\x18\x02 \x01(\v2>.google.ads.googleads.v20.common.CustomerMatchUserListMetadataH\x00R\x1dcustomerMatchUserListMetadataB\n" +
	"\n" +
	"\bmetadata\"\xaa\x01\n" +
	"\x11UserDataOperation\x12C\n" +
	"\x06create\x18\x01 \x01(\v2).google.ads.googleads.v20.common.UserDataH\x00R\x06create\x12C\n" +
	"\x06remove\x18\x02 \x01(\v2).google.ads.googleads.v20.common.UserDataH\x00R\x06removeB\v\n" +
	"\toperation\"\xbb\x01\n" +
	"\x16UploadUserDataResponse\x12-\n" +
	"\x10upload_date_time\x18\x03 \x01(\tH\x00R\x0euploadDateTime\x88\x01\x01\x12?\n" +
	"\x19received_operations_count\x18\x04 \x01(\x05H\x01R\x17receivedOperationsCount\x88\x01\x01B\x13\n" +
	"\x11_upload_date_timeB\x1c\n" +
	"\x1a_received_operations_count2\x9a\x02\n" +
	"\x0fUserDataService\x12\xbf\x01\n" +
	"\x0eUploadUserData\x128.google.ads.googleads.v20.services.UploadUserDataRequest\x1a9.google.ads.googleads.v20.services.UploadUserDataResponse\"8\x82\xd3\xe4\x93\x022:\x01*\"-/v20/customers/{customer_id=*}:uploadUserData\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x80\x02\n" +
	"%com.google.ads.googleads.v20.servicesB\x14UserDataServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_user_data_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_user_data_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_user_data_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_user_data_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_user_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_user_data_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_user_data_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_user_data_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v20_services_user_data_service_proto_goTypes = []any{
	(*UploadUserDataRequest)(nil),                // 0: google.ads.googleads.v20.services.UploadUserDataRequest
	(*UserDataOperation)(nil),                    // 1: google.ads.googleads.v20.services.UserDataOperation
	(*UploadUserDataResponse)(nil),               // 2: google.ads.googleads.v20.services.UploadUserDataResponse
	(*common.CustomerMatchUserListMetadata)(nil), // 3: google.ads.googleads.v20.common.CustomerMatchUserListMetadata
	(*common.UserData)(nil),                      // 4: google.ads.googleads.v20.common.UserData
}
var file_google_ads_googleads_v20_services_user_data_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.UploadUserDataRequest.operations:type_name -> google.ads.googleads.v20.services.UserDataOperation
	3, // 1: google.ads.googleads.v20.services.UploadUserDataRequest.customer_match_user_list_metadata:type_name -> google.ads.googleads.v20.common.CustomerMatchUserListMetadata
	4, // 2: google.ads.googleads.v20.services.UserDataOperation.create:type_name -> google.ads.googleads.v20.common.UserData
	4, // 3: google.ads.googleads.v20.services.UserDataOperation.remove:type_name -> google.ads.googleads.v20.common.UserData
	0, // 4: google.ads.googleads.v20.services.UserDataService.UploadUserData:input_type -> google.ads.googleads.v20.services.UploadUserDataRequest
	2, // 5: google.ads.googleads.v20.services.UserDataService.UploadUserData:output_type -> google.ads.googleads.v20.services.UploadUserDataResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_user_data_service_proto_init() }
func file_google_ads_googleads_v20_services_user_data_service_proto_init() {
	if File_google_ads_googleads_v20_services_user_data_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[0].OneofWrappers = []any{
		(*UploadUserDataRequest_CustomerMatchUserListMetadata)(nil),
	}
	file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[1].OneofWrappers = []any{
		(*UserDataOperation_Create)(nil),
		(*UserDataOperation_Remove)(nil),
	}
	file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_user_data_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_user_data_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_user_data_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_user_data_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_user_data_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_user_data_service_proto = out.File
	file_google_ads_googleads_v20_services_user_data_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_user_data_service_proto_depIdxs = nil
}
