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
// source: google/ads/googleads/v20/services/conversion_goal_campaign_config_service.proto

package services

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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
// [ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs][google.ads.googleads.v20.services.ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs].
type MutateConversionGoalCampaignConfigsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer whose custom conversion goals are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual conversion goal
	// campaign config.
	Operations []*ConversionGoalCampaignConfigOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,4,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v20.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MutateConversionGoalCampaignConfigsRequest) Reset() {
	*x = MutateConversionGoalCampaignConfigsRequest{}
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateConversionGoalCampaignConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigsRequest) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigsRequest.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetOperations() []*ConversionGoalCampaignConfigOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateConversionGoalCampaignConfigsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (update) on a conversion goal campaign config.
type ConversionGoalCampaignConfigOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*ConversionGoalCampaignConfigOperation_Update
	Operation     isConversionGoalCampaignConfigOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionGoalCampaignConfigOperation) Reset() {
	*x = ConversionGoalCampaignConfigOperation{}
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionGoalCampaignConfigOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionGoalCampaignConfigOperation) ProtoMessage() {}

func (x *ConversionGoalCampaignConfigOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionGoalCampaignConfigOperation.ProtoReflect.Descriptor instead.
func (*ConversionGoalCampaignConfigOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConversionGoalCampaignConfigOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *ConversionGoalCampaignConfigOperation) GetOperation() isConversionGoalCampaignConfigOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *ConversionGoalCampaignConfigOperation) GetUpdate() *resources.ConversionGoalCampaignConfig {
	if x != nil {
		if x, ok := x.Operation.(*ConversionGoalCampaignConfigOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

type isConversionGoalCampaignConfigOperation_Operation interface {
	isConversionGoalCampaignConfigOperation_Operation()
}

type ConversionGoalCampaignConfigOperation_Update struct {
	// Update operation: The conversion goal campaign config is expected to have
	// a valid resource name.
	Update *resources.ConversionGoalCampaignConfig `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

func (*ConversionGoalCampaignConfigOperation_Update) isConversionGoalCampaignConfigOperation_Operation() {
}

// Response message for a conversion goal campaign config mutate.
type MutateConversionGoalCampaignConfigsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// All results for the mutate.
	Results       []*MutateConversionGoalCampaignConfigResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateConversionGoalCampaignConfigsResponse) Reset() {
	*x = MutateConversionGoalCampaignConfigsResponse{}
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateConversionGoalCampaignConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigsResponse) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigsResponse.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateConversionGoalCampaignConfigsResponse) GetResults() []*MutateConversionGoalCampaignConfigResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the conversion goal campaign config mutate.
type MutateConversionGoalCampaignConfigResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated ConversionGoalCampaignConfig with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	ConversionGoalCampaignConfig *resources.ConversionGoalCampaignConfig `protobuf:"bytes,2,opt,name=conversion_goal_campaign_config,json=conversionGoalCampaignConfig,proto3" json:"conversion_goal_campaign_config,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *MutateConversionGoalCampaignConfigResult) Reset() {
	*x = MutateConversionGoalCampaignConfigResult{}
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateConversionGoalCampaignConfigResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateConversionGoalCampaignConfigResult) ProtoMessage() {}

func (x *MutateConversionGoalCampaignConfigResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateConversionGoalCampaignConfigResult.ProtoReflect.Descriptor instead.
func (*MutateConversionGoalCampaignConfigResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateConversionGoalCampaignConfigResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateConversionGoalCampaignConfigResult) GetConversionGoalCampaignConfig() *resources.ConversionGoalCampaignConfig {
	if x != nil {
		return x.ConversionGoalCampaignConfig
	}
	return nil
}

var File_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDesc = "" +
	"\n" +
	"Ogoogle/ads/googleads/v20/services/conversion_goal_campaign_config_service.proto\x12!google.ads.googleads.v20.services\x1a:google/ads/googleads/v20/enums/response_content_type.proto\x1aHgoogle/ads/googleads/v20/resources/conversion_goal_campaign_config.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\"\xe7\x02\n" +
	"*MutateConversionGoalCampaignConfigsRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12m\n" +
	"\n" +
	"operations\x18\x02 \x03(\v2H.google.ads.googleads.v20.services.ConversionGoalCampaignConfigOperationB\x03\xe0A\x02R\n" +
	"operations\x12#\n" +
	"\rvalidate_only\x18\x03 \x01(\bR\fvalidateOnly\x12\x7f\n" +
	"\x15response_content_type\x18\x04 \x01(\x0e2K.google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentTypeR\x13responseContentType\"\xcd\x01\n" +
	"%ConversionGoalCampaignConfigOperation\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12Z\n" +
	"\x06update\x18\x01 \x01(\v2@.google.ads.googleads.v20.resources.ConversionGoalCampaignConfigH\x00R\x06updateB\v\n" +
	"\toperation\"\x94\x01\n" +
	"+MutateConversionGoalCampaignConfigsResponse\x12e\n" +
	"\aresults\x18\x01 \x03(\v2K.google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigResultR\aresults\"\x95\x02\n" +
	"(MutateConversionGoalCampaignConfigResult\x12_\n" +
	"\rresource_name\x18\x01 \x01(\tB:\xfaA7\n" +
	"5googleads.googleapis.com/ConversionGoalCampaignConfigR\fresourceName\x12\x87\x01\n" +
	"\x1fconversion_goal_campaign_config\x18\x02 \x01(\v2@.google.ads.googleads.v20.resources.ConversionGoalCampaignConfigR\x1cconversionGoalCampaignConfig2\x9c\x03\n" +
	"#ConversionGoalCampaignConfigService\x12\xad\x02\n" +
	"#MutateConversionGoalCampaignConfigs\x12M.google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsRequest\x1aN.google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsResponse\"g\xdaA\x16customer_id,operations\x82\xd3\xe4\x93\x02H:\x01*\"C/v20/customers/{customer_id=*}/conversionGoalCampaignConfigs:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x94\x02\n" +
	"%com.google.ads.googleads.v20.servicesB(ConversionGoalCampaignConfigServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_goTypes = []any{
	(*MutateConversionGoalCampaignConfigsRequest)(nil),     // 0: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsRequest
	(*ConversionGoalCampaignConfigOperation)(nil),          // 1: google.ads.googleads.v20.services.ConversionGoalCampaignConfigOperation
	(*MutateConversionGoalCampaignConfigsResponse)(nil),    // 2: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsResponse
	(*MutateConversionGoalCampaignConfigResult)(nil),       // 3: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.ConversionGoalCampaignConfig)(nil),         // 6: google.ads.googleads.v20.resources.ConversionGoalCampaignConfig
}
var file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsRequest.operations:type_name -> google.ads.googleads.v20.services.ConversionGoalCampaignConfigOperation
	4, // 1: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsRequest.response_content_type:type_name -> google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v20.services.ConversionGoalCampaignConfigOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v20.services.ConversionGoalCampaignConfigOperation.update:type_name -> google.ads.googleads.v20.resources.ConversionGoalCampaignConfig
	3, // 4: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsResponse.results:type_name -> google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigResult
	6, // 5: google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigResult.conversion_goal_campaign_config:type_name -> google.ads.googleads.v20.resources.ConversionGoalCampaignConfig
	0, // 6: google.ads.googleads.v20.services.ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs:input_type -> google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsRequest
	2, // 7: google.ads.googleads.v20.services.ConversionGoalCampaignConfigService.MutateConversionGoalCampaignConfigs:output_type -> google.ads.googleads.v20.services.MutateConversionGoalCampaignConfigsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_init()
}
func file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_init() {
	if File_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes[1].OneofWrappers = []any{
		(*ConversionGoalCampaignConfigOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto = out.File
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_conversion_goal_campaign_config_service_proto_depIdxs = nil
}
