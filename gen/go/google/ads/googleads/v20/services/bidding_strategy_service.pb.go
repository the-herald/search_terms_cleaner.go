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
// source: google/ads/googleads/v20/services/bidding_strategy_service.proto

package services

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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
// [BiddingStrategyService.MutateBiddingStrategies][google.ads.googleads.v20.services.BiddingStrategyService.MutateBiddingStrategies].
type MutateBiddingStrategiesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer whose bidding strategies are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual bidding
	// strategies.
	Operations []*BiddingStrategyOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v20.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MutateBiddingStrategiesRequest) Reset() {
	*x = MutateBiddingStrategiesRequest{}
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingStrategiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingStrategiesRequest) ProtoMessage() {}

func (x *MutateBiddingStrategiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingStrategiesRequest.ProtoReflect.Descriptor instead.
func (*MutateBiddingStrategiesRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateBiddingStrategiesRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateBiddingStrategiesRequest) GetOperations() []*BiddingStrategyOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateBiddingStrategiesRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateBiddingStrategiesRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateBiddingStrategiesRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, update, remove) on a bidding strategy.
type BiddingStrategyOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*BiddingStrategyOperation_Create
	//	*BiddingStrategyOperation_Update
	//	*BiddingStrategyOperation_Remove
	Operation     isBiddingStrategyOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BiddingStrategyOperation) Reset() {
	*x = BiddingStrategyOperation{}
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BiddingStrategyOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategyOperation) ProtoMessage() {}

func (x *BiddingStrategyOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategyOperation.ProtoReflect.Descriptor instead.
func (*BiddingStrategyOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescGZIP(), []int{1}
}

func (x *BiddingStrategyOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *BiddingStrategyOperation) GetOperation() isBiddingStrategyOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *BiddingStrategyOperation) GetCreate() *resources.BiddingStrategy {
	if x != nil {
		if x, ok := x.Operation.(*BiddingStrategyOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *BiddingStrategyOperation) GetUpdate() *resources.BiddingStrategy {
	if x != nil {
		if x, ok := x.Operation.(*BiddingStrategyOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

func (x *BiddingStrategyOperation) GetRemove() string {
	if x != nil {
		if x, ok := x.Operation.(*BiddingStrategyOperation_Remove); ok {
			return x.Remove
		}
	}
	return ""
}

type isBiddingStrategyOperation_Operation interface {
	isBiddingStrategyOperation_Operation()
}

type BiddingStrategyOperation_Create struct {
	// Create operation: No resource name is expected for the new bidding
	// strategy.
	Create *resources.BiddingStrategy `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type BiddingStrategyOperation_Update struct {
	// Update operation: The bidding strategy is expected to have a valid
	// resource name.
	Update *resources.BiddingStrategy `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type BiddingStrategyOperation_Remove struct {
	// Remove operation: A resource name for the removed bidding strategy is
	// expected, in this format:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*BiddingStrategyOperation_Create) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Update) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Remove) isBiddingStrategyOperation_Operation() {}

// Response message for bidding strategy mutate.
type MutateBiddingStrategiesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results       []*MutateBiddingStrategyResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateBiddingStrategiesResponse) Reset() {
	*x = MutateBiddingStrategiesResponse{}
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingStrategiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingStrategiesResponse) ProtoMessage() {}

func (x *MutateBiddingStrategiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingStrategiesResponse.ProtoReflect.Descriptor instead.
func (*MutateBiddingStrategiesResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateBiddingStrategiesResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateBiddingStrategiesResponse) GetResults() []*MutateBiddingStrategyResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the bidding strategy mutate.
type MutateBiddingStrategyResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated bidding strategy with only mutable fields after mutate. The
	// field will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	BiddingStrategy *resources.BiddingStrategy `protobuf:"bytes,2,opt,name=bidding_strategy,json=biddingStrategy,proto3" json:"bidding_strategy,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MutateBiddingStrategyResult) Reset() {
	*x = MutateBiddingStrategyResult{}
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateBiddingStrategyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingStrategyResult) ProtoMessage() {}

func (x *MutateBiddingStrategyResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingStrategyResult.ProtoReflect.Descriptor instead.
func (*MutateBiddingStrategyResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateBiddingStrategyResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateBiddingStrategyResult) GetBiddingStrategy() *resources.BiddingStrategy {
	if x != nil {
		return x.BiddingStrategy
	}
	return nil
}

var File_google_ads_googleads_v20_services_bidding_strategy_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDesc = "" +
	"\n" +
	"@google/ads/googleads/v20/services/bidding_strategy_service.proto\x12!google.ads.googleads.v20.services\x1a:google/ads/googleads/v20/enums/response_content_type.proto\x1a9google/ads/googleads/v20/resources/bidding_strategy.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a\x17google/rpc/status.proto\"\xf7\x02\n" +
	"\x1eMutateBiddingStrategiesRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12`\n" +
	"\n" +
	"operations\x18\x02 \x03(\v2;.google.ads.googleads.v20.services.BiddingStrategyOperationB\x03\xe0A\x02R\n" +
	"operations\x12'\n" +
	"\x0fpartial_failure\x18\x03 \x01(\bR\x0epartialFailure\x12#\n" +
	"\rvalidate_only\x18\x04 \x01(\bR\fvalidateOnly\x12\x7f\n" +
	"\x15response_content_type\x18\x05 \x01(\x0e2K.google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentTypeR\x13responseContentType\"\xcb\x02\n" +
	"\x18BiddingStrategyOperation\x12;\n" +
	"\vupdate_mask\x18\x04 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12M\n" +
	"\x06create\x18\x01 \x01(\v23.google.ads.googleads.v20.resources.BiddingStrategyH\x00R\x06create\x12M\n" +
	"\x06update\x18\x02 \x01(\v23.google.ads.googleads.v20.resources.BiddingStrategyH\x00R\x06update\x12G\n" +
	"\x06remove\x18\x03 \x01(\tB-\xfaA*\n" +
	"(googleads.googleapis.com/BiddingStrategyH\x00R\x06removeB\v\n" +
	"\toperation\"\xc3\x01\n" +
	"\x1fMutateBiddingStrategiesResponse\x12F\n" +
	"\x15partial_failure_error\x18\x03 \x01(\v2\x12.google.rpc.StatusR\x13partialFailureError\x12X\n" +
	"\aresults\x18\x02 \x03(\v2>.google.ads.googleads.v20.services.MutateBiddingStrategyResultR\aresults\"\xd1\x01\n" +
	"\x1bMutateBiddingStrategyResult\x12R\n" +
	"\rresource_name\x18\x01 \x01(\tB-\xfaA*\n" +
	"(googleads.googleapis.com/BiddingStrategyR\fresourceName\x12^\n" +
	"\x10bidding_strategy\x18\x02 \x01(\v23.google.ads.googleads.v20.resources.BiddingStrategyR\x0fbiddingStrategy2\xdf\x02\n" +
	"\x16BiddingStrategyService\x12\xfd\x01\n" +
	"\x17MutateBiddingStrategies\x12A.google.ads.googleads.v20.services.MutateBiddingStrategiesRequest\x1aB.google.ads.googleads.v20.services.MutateBiddingStrategiesResponse\"[\xdaA\x16customer_id,operations\x82\xd3\xe4\x93\x02<:\x01*\"7/v20/customers/{customer_id=*}/biddingStrategies:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x87\x02\n" +
	"%com.google.ads.googleads.v20.servicesB\x1bBiddingStrategyServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_bidding_strategy_service_proto_goTypes = []any{
	(*MutateBiddingStrategiesRequest)(nil),                 // 0: google.ads.googleads.v20.services.MutateBiddingStrategiesRequest
	(*BiddingStrategyOperation)(nil),                       // 1: google.ads.googleads.v20.services.BiddingStrategyOperation
	(*MutateBiddingStrategiesResponse)(nil),                // 2: google.ads.googleads.v20.services.MutateBiddingStrategiesResponse
	(*MutateBiddingStrategyResult)(nil),                    // 3: google.ads.googleads.v20.services.MutateBiddingStrategyResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.BiddingStrategy)(nil),                      // 6: google.ads.googleads.v20.resources.BiddingStrategy
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v20_services_bidding_strategy_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.MutateBiddingStrategiesRequest.operations:type_name -> google.ads.googleads.v20.services.BiddingStrategyOperation
	4, // 1: google.ads.googleads.v20.services.MutateBiddingStrategiesRequest.response_content_type:type_name -> google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v20.services.BiddingStrategyOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v20.services.BiddingStrategyOperation.create:type_name -> google.ads.googleads.v20.resources.BiddingStrategy
	6, // 4: google.ads.googleads.v20.services.BiddingStrategyOperation.update:type_name -> google.ads.googleads.v20.resources.BiddingStrategy
	7, // 5: google.ads.googleads.v20.services.MutateBiddingStrategiesResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v20.services.MutateBiddingStrategiesResponse.results:type_name -> google.ads.googleads.v20.services.MutateBiddingStrategyResult
	6, // 7: google.ads.googleads.v20.services.MutateBiddingStrategyResult.bidding_strategy:type_name -> google.ads.googleads.v20.resources.BiddingStrategy
	0, // 8: google.ads.googleads.v20.services.BiddingStrategyService.MutateBiddingStrategies:input_type -> google.ads.googleads.v20.services.MutateBiddingStrategiesRequest
	2, // 9: google.ads.googleads.v20.services.BiddingStrategyService.MutateBiddingStrategies:output_type -> google.ads.googleads.v20.services.MutateBiddingStrategiesResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_bidding_strategy_service_proto_init() }
func file_google_ads_googleads_v20_services_bidding_strategy_service_proto_init() {
	if File_google_ads_googleads_v20_services_bidding_strategy_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes[1].OneofWrappers = []any{
		(*BiddingStrategyOperation_Create)(nil),
		(*BiddingStrategyOperation_Update)(nil),
		(*BiddingStrategyOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_bidding_strategy_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_bidding_strategy_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_bidding_strategy_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_bidding_strategy_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_bidding_strategy_service_proto = out.File
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_bidding_strategy_service_proto_depIdxs = nil
}
