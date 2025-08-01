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
// source: google/ads/googleads/v20/services/customer_sk_ad_network_conversion_value_schema_service.proto

package services

import (
	resources "google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A single update operation for a CustomerSkAdNetworkConversionValueSchema.
type CustomerSkAdNetworkConversionValueSchemaOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Update operation: The schema is expected to have a valid resource name.
	Update        *resources.CustomerSkAdNetworkConversionValueSchema `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) Reset() {
	*x = CustomerSkAdNetworkConversionValueSchemaOperation{}
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSkAdNetworkConversionValueSchemaOperation) ProtoMessage() {}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSkAdNetworkConversionValueSchemaOperation.ProtoReflect.Descriptor instead.
func (*CustomerSkAdNetworkConversionValueSchemaOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) GetUpdate() *resources.CustomerSkAdNetworkConversionValueSchema {
	if x != nil {
		return x.Update
	}
	return nil
}

// Request message for
// [CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema][google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema].
type MutateCustomerSkAdNetworkConversionValueSchemaRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the customer whose shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform.
	Operation *CustomerSkAdNetworkConversionValueSchemaOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Optional. If true, enables returning warnings. Warnings return error
	// messages and error codes without blocking the execution of the mutate
	// operation.
	EnableWarnings bool `protobuf:"varint,4,opt,name=enable_warnings,json=enableWarnings,proto3" json:"enable_warnings,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaRequest{}
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetOperation() *CustomerSkAdNetworkConversionValueSchemaOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetEnableWarnings() bool {
	if x != nil {
		return x.EnableWarnings
	}
	return false
}

// The result for the CustomerSkAdNetworkConversionValueSchema mutate.
type MutateCustomerSkAdNetworkConversionValueSchemaResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Resource name of the customer that was modified.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// App ID of the SkanConversionValue modified.
	AppId         string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResult{}
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// Response message for MutateCustomerSkAdNetworkConversionValueSchema.
type MutateCustomerSkAdNetworkConversionValueSchemaResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// All results for the mutate.
	Result *MutateCustomerSkAdNetworkConversionValueSchemaResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// Non blocking errors that provides schema validation failure details.
	// Returned only when enable_warnings = true.
	Warning       *status.Status `protobuf:"bytes,2,opt,name=warning,proto3" json:"warning,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResponse{}
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) GetResult() *MutateCustomerSkAdNetworkConversionValueSchemaResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) GetWarning() *status.Status {
	if x != nil {
		return x.Warning
	}
	return nil
}

var File_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc = "" +
	"\n" +
	"^google/ads/googleads/v20/services/customer_sk_ad_network_conversion_value_schema_service.proto\x12!google.ads.googleads.v20.services\x1aWgoogle/ads/googleads/v20/resources/customer_sk_ad_network_conversion_value_schema.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x17google/rpc/status.proto\"\x99\x01\n" +
	"1CustomerSkAdNetworkConversionValueSchemaOperation\x12d\n" +
	"\x06update\x18\x01 \x01(\v2L.google.ads.googleads.v20.resources.CustomerSkAdNetworkConversionValueSchemaR\x06update\"\x9f\x02\n" +
	"5MutateCustomerSkAdNetworkConversionValueSchemaRequest\x12\x1f\n" +
	"\vcustomer_id\x18\x01 \x01(\tR\n" +
	"customerId\x12r\n" +
	"\toperation\x18\x02 \x01(\v2T.google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaOperationR\toperation\x12#\n" +
	"\rvalidate_only\x18\x03 \x01(\bR\fvalidateOnly\x12,\n" +
	"\x0fenable_warnings\x18\x04 \x01(\bB\x03\xe0A\x01R\x0eenableWarnings\"\xba\x01\n" +
	"4MutateCustomerSkAdNetworkConversionValueSchemaResult\x12k\n" +
	"\rresource_name\x18\x01 \x01(\tBF\xfaAC\n" +
	"Agoogleads.googleapis.com/CustomerSkAdNetworkConversionValueSchemaR\fresourceName\x12\x15\n" +
	"\x06app_id\x18\x02 \x01(\tR\x05appId\"\xd7\x01\n" +
	"6MutateCustomerSkAdNetworkConversionValueSchemaResponse\x12o\n" +
	"\x06result\x18\x01 \x01(\v2W.google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResultR\x06result\x12,\n" +
	"\awarning\x18\x02 \x01(\v2\x12.google.rpc.StatusR\awarning2\xbc\x03\n" +
	"/CustomerSkAdNetworkConversionValueSchemaService\x12\xc1\x02\n" +
	".MutateCustomerSkAdNetworkConversionValueSchema\x12X.google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest\x1aY.google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse\"Z\x82\xd3\xe4\x93\x02T:\x01*\"O/v20/customers/{customer_id=*}/customerSkAdNetworkConversionValueSchemas:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\xa0\x02\n" +
	"%com.google.ads.googleads.v20.servicesB4CustomerSkAdNetworkConversionValueSchemaServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = []any{
	(*CustomerSkAdNetworkConversionValueSchemaOperation)(nil),      // 0: google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaOperation
	(*MutateCustomerSkAdNetworkConversionValueSchemaRequest)(nil),  // 1: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	(*MutateCustomerSkAdNetworkConversionValueSchemaResult)(nil),   // 2: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	(*MutateCustomerSkAdNetworkConversionValueSchemaResponse)(nil), // 3: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	(*resources.CustomerSkAdNetworkConversionValueSchema)(nil),     // 4: google.ads.googleads.v20.resources.CustomerSkAdNetworkConversionValueSchema
	(*status.Status)(nil), // 5: google.rpc.Status
}
var file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaOperation.update:type_name -> google.ads.googleads.v20.resources.CustomerSkAdNetworkConversionValueSchema
	0, // 1: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest.operation:type_name -> google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaOperation
	2, // 2: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse.result:type_name -> google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	5, // 3: google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse.warning:type_name -> google.rpc.Status
	1, // 4: google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:input_type -> google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	3, // 5: google.ads.googleads.v20.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:output_type -> google.ads.googleads.v20.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_init()
}
func file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_init() {
	if File_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto = out.File
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = nil
}
