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
// source: google/ads/googleads/v20/services/ad_group_label_service.proto

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

// Request message for
// [AdGroupLabelService.MutateAdGroupLabels][google.ads.googleads.v20.services.AdGroupLabelService.MutateAdGroupLabels].
type MutateAdGroupLabelsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the customer whose ad group labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on ad group labels.
	Operations []*AdGroupLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly  bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAdGroupLabelsRequest) Reset() {
	*x = MutateAdGroupLabelsRequest{}
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupLabelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupLabelsRequest) ProtoMessage() {}

func (x *MutateAdGroupLabelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupLabelsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupLabelsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAdGroupLabelsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupLabelsRequest) GetOperations() []*AdGroupLabelOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupLabelsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupLabelsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on an ad group label.
type AdGroupLabelOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*AdGroupLabelOperation_Create
	//	*AdGroupLabelOperation_Remove
	Operation     isAdGroupLabelOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupLabelOperation) Reset() {
	*x = AdGroupLabelOperation{}
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupLabelOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupLabelOperation) ProtoMessage() {}

func (x *AdGroupLabelOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupLabelOperation.ProtoReflect.Descriptor instead.
func (*AdGroupLabelOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescGZIP(), []int{1}
}

func (x *AdGroupLabelOperation) GetOperation() isAdGroupLabelOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *AdGroupLabelOperation) GetCreate() *resources.AdGroupLabel {
	if x != nil {
		if x, ok := x.Operation.(*AdGroupLabelOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *AdGroupLabelOperation) GetRemove() string {
	if x != nil {
		if x, ok := x.Operation.(*AdGroupLabelOperation_Remove); ok {
			return x.Remove
		}
	}
	return ""
}

type isAdGroupLabelOperation_Operation interface {
	isAdGroupLabelOperation_Operation()
}

type AdGroupLabelOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group
	// label.
	Create *resources.AdGroupLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupLabelOperation_Remove struct {
	// Remove operation: A resource name for the ad group label
	// being removed, in this format:
	//
	// `customers/{customer_id}/adGroupLabels/{ad_group_id}~{label_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupLabelOperation_Create) isAdGroupLabelOperation_Operation() {}

func (*AdGroupLabelOperation_Remove) isAdGroupLabelOperation_Operation() {}

// Response message for an ad group labels mutate.
type MutateAdGroupLabelsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results       []*MutateAdGroupLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAdGroupLabelsResponse) Reset() {
	*x = MutateAdGroupLabelsResponse{}
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupLabelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupLabelsResponse) ProtoMessage() {}

func (x *MutateAdGroupLabelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupLabelsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupLabelsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAdGroupLabelsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupLabelsResponse) GetResults() []*MutateAdGroupLabelResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for an ad group label mutate.
type MutateAdGroupLabelResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName  string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAdGroupLabelResult) Reset() {
	*x = MutateAdGroupLabelResult{}
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupLabelResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupLabelResult) ProtoMessage() {}

func (x *MutateAdGroupLabelResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupLabelResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupLabelResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupLabelResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v20_services_ad_group_label_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/services/ad_group_label_service.proto\x12!google.ads.googleads.v20.services\x1a7google/ads/googleads/v20/resources/ad_group_label.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x17google/rpc/status.proto\"\xef\x01\n" +
	"\x1aMutateAdGroupLabelsRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12]\n" +
	"\n" +
	"operations\x18\x02 \x03(\v28.google.ads.googleads.v20.services.AdGroupLabelOperationB\x03\xe0A\x02R\n" +
	"operations\x12'\n" +
	"\x0fpartial_failure\x18\x03 \x01(\bR\x0epartialFailure\x12#\n" +
	"\rvalidate_only\x18\x04 \x01(\bR\fvalidateOnly\"\xb6\x01\n" +
	"\x15AdGroupLabelOperation\x12J\n" +
	"\x06create\x18\x01 \x01(\v20.google.ads.googleads.v20.resources.AdGroupLabelH\x00R\x06create\x12D\n" +
	"\x06remove\x18\x02 \x01(\tB*\xfaA'\n" +
	"%googleads.googleapis.com/AdGroupLabelH\x00R\x06removeB\v\n" +
	"\toperation\"\xbc\x01\n" +
	"\x1bMutateAdGroupLabelsResponse\x12F\n" +
	"\x15partial_failure_error\x18\x03 \x01(\v2\x12.google.rpc.StatusR\x13partialFailureError\x12U\n" +
	"\aresults\x18\x02 \x03(\v2;.google.ads.googleads.v20.services.MutateAdGroupLabelResultR\aresults\"k\n" +
	"\x18MutateAdGroupLabelResult\x12O\n" +
	"\rresource_name\x18\x01 \x01(\tB*\xfaA'\n" +
	"%googleads.googleapis.com/AdGroupLabelR\fresourceName2\xcc\x02\n" +
	"\x13AdGroupLabelService\x12\xed\x01\n" +
	"\x13MutateAdGroupLabels\x12=.google.ads.googleads.v20.services.MutateAdGroupLabelsRequest\x1a>.google.ads.googleads.v20.services.MutateAdGroupLabelsResponse\"W\xdaA\x16customer_id,operations\x82\xd3\xe4\x93\x028:\x01*\"3/v20/customers/{customer_id=*}/adGroupLabels:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x84\x02\n" +
	"%com.google.ads.googleads.v20.servicesB\x18AdGroupLabelServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_ad_group_label_service_proto_goTypes = []any{
	(*MutateAdGroupLabelsRequest)(nil),  // 0: google.ads.googleads.v20.services.MutateAdGroupLabelsRequest
	(*AdGroupLabelOperation)(nil),       // 1: google.ads.googleads.v20.services.AdGroupLabelOperation
	(*MutateAdGroupLabelsResponse)(nil), // 2: google.ads.googleads.v20.services.MutateAdGroupLabelsResponse
	(*MutateAdGroupLabelResult)(nil),    // 3: google.ads.googleads.v20.services.MutateAdGroupLabelResult
	(*resources.AdGroupLabel)(nil),      // 4: google.ads.googleads.v20.resources.AdGroupLabel
	(*status.Status)(nil),               // 5: google.rpc.Status
}
var file_google_ads_googleads_v20_services_ad_group_label_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.MutateAdGroupLabelsRequest.operations:type_name -> google.ads.googleads.v20.services.AdGroupLabelOperation
	4, // 1: google.ads.googleads.v20.services.AdGroupLabelOperation.create:type_name -> google.ads.googleads.v20.resources.AdGroupLabel
	5, // 2: google.ads.googleads.v20.services.MutateAdGroupLabelsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 3: google.ads.googleads.v20.services.MutateAdGroupLabelsResponse.results:type_name -> google.ads.googleads.v20.services.MutateAdGroupLabelResult
	0, // 4: google.ads.googleads.v20.services.AdGroupLabelService.MutateAdGroupLabels:input_type -> google.ads.googleads.v20.services.MutateAdGroupLabelsRequest
	2, // 5: google.ads.googleads.v20.services.AdGroupLabelService.MutateAdGroupLabels:output_type -> google.ads.googleads.v20.services.MutateAdGroupLabelsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_ad_group_label_service_proto_init() }
func file_google_ads_googleads_v20_services_ad_group_label_service_proto_init() {
	if File_google_ads_googleads_v20_services_ad_group_label_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes[1].OneofWrappers = []any{
		(*AdGroupLabelOperation_Create)(nil),
		(*AdGroupLabelOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_ad_group_label_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_ad_group_label_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_ad_group_label_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_ad_group_label_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_ad_group_label_service_proto = out.File
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_ad_group_label_service_proto_depIdxs = nil
}
