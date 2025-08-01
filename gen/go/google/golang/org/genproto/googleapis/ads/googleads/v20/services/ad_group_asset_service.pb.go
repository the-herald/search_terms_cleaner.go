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
// source: google/ads/googleads/v20/services/ad_group_asset_service.proto

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
// [AdGroupAssetService.MutateAdGroupAssets][google.ads.googleads.v20.services.AdGroupAssetService.MutateAdGroupAssets].
type MutateAdGroupAssetsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer whose ad group assets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad group assets.
	Operations []*AdGroupAssetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (x *MutateAdGroupAssetsRequest) Reset() {
	*x = MutateAdGroupAssetsRequest{}
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetsRequest) ProtoMessage() {}

func (x *MutateAdGroupAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetsRequest.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAdGroupAssetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAdGroupAssetsRequest) GetOperations() []*AdGroupAssetOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAdGroupAssetsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateAdGroupAssetsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateAdGroupAssetsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, update, remove) on an ad group asset.
type AdGroupAssetOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*AdGroupAssetOperation_Create
	//	*AdGroupAssetOperation_Update
	//	*AdGroupAssetOperation_Remove
	Operation     isAdGroupAssetOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupAssetOperation) Reset() {
	*x = AdGroupAssetOperation{}
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupAssetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupAssetOperation) ProtoMessage() {}

func (x *AdGroupAssetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupAssetOperation.ProtoReflect.Descriptor instead.
func (*AdGroupAssetOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescGZIP(), []int{1}
}

func (x *AdGroupAssetOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *AdGroupAssetOperation) GetOperation() isAdGroupAssetOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *AdGroupAssetOperation) GetCreate() *resources.AdGroupAsset {
	if x != nil {
		if x, ok := x.Operation.(*AdGroupAssetOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *AdGroupAssetOperation) GetUpdate() *resources.AdGroupAsset {
	if x != nil {
		if x, ok := x.Operation.(*AdGroupAssetOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

func (x *AdGroupAssetOperation) GetRemove() string {
	if x != nil {
		if x, ok := x.Operation.(*AdGroupAssetOperation_Remove); ok {
			return x.Remove
		}
	}
	return ""
}

type isAdGroupAssetOperation_Operation interface {
	isAdGroupAssetOperation_Operation()
}

type AdGroupAssetOperation_Create struct {
	// Create operation: No resource name is expected for the new ad group
	// asset.
	Create *resources.AdGroupAsset `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAssetOperation_Update struct {
	// Update operation: The ad group asset is expected to have a valid resource
	// name.
	Update *resources.AdGroupAsset `protobuf:"bytes,3,opt,name=update,proto3,oneof"`
}

type AdGroupAssetOperation_Remove struct {
	// Remove operation: A resource name for the removed ad group asset is
	// expected, in this format:
	//
	// `customers/{customer_id}/adGroupAssets/{ad_group_id}~{asset_id}~{field_type}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAssetOperation_Create) isAdGroupAssetOperation_Operation() {}

func (*AdGroupAssetOperation_Update) isAdGroupAssetOperation_Operation() {}

func (*AdGroupAssetOperation_Remove) isAdGroupAssetOperation_Operation() {}

// Response message for an ad group asset mutate.
type MutateAdGroupAssetsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results       []*MutateAdGroupAssetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAdGroupAssetsResponse) Reset() {
	*x = MutateAdGroupAssetsResponse{}
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetsResponse) ProtoMessage() {}

func (x *MutateAdGroupAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetsResponse.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAdGroupAssetsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateAdGroupAssetsResponse) GetResults() []*MutateAdGroupAssetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the ad group asset mutate.
type MutateAdGroupAssetResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated ad group asset with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	AdGroupAsset  *resources.AdGroupAsset `protobuf:"bytes,2,opt,name=ad_group_asset,json=adGroupAsset,proto3" json:"ad_group_asset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAdGroupAssetResult) Reset() {
	*x = MutateAdGroupAssetResult{}
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAdGroupAssetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAdGroupAssetResult) ProtoMessage() {}

func (x *MutateAdGroupAssetResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAdGroupAssetResult.ProtoReflect.Descriptor instead.
func (*MutateAdGroupAssetResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAdGroupAssetResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateAdGroupAssetResult) GetAdGroupAsset() *resources.AdGroupAsset {
	if x != nil {
		return x.AdGroupAsset
	}
	return nil
}

var File_google_ads_googleads_v20_services_ad_group_asset_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/services/ad_group_asset_service.proto\x12!google.ads.googleads.v20.services\x1a:google/ads/googleads/v20/enums/response_content_type.proto\x1a7google/ads/googleads/v20/resources/ad_group_asset.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a\x17google/rpc/status.proto\"\xf0\x02\n" +
	"\x1aMutateAdGroupAssetsRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12]\n" +
	"\n" +
	"operations\x18\x02 \x03(\v28.google.ads.googleads.v20.services.AdGroupAssetOperationB\x03\xe0A\x02R\n" +
	"operations\x12'\n" +
	"\x0fpartial_failure\x18\x03 \x01(\bR\x0epartialFailure\x12#\n" +
	"\rvalidate_only\x18\x04 \x01(\bR\fvalidateOnly\x12\x7f\n" +
	"\x15response_content_type\x18\x05 \x01(\x0e2K.google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentTypeR\x13responseContentType\"\xbf\x02\n" +
	"\x15AdGroupAssetOperation\x12;\n" +
	"\vupdate_mask\x18\x04 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12J\n" +
	"\x06create\x18\x01 \x01(\v20.google.ads.googleads.v20.resources.AdGroupAssetH\x00R\x06create\x12J\n" +
	"\x06update\x18\x03 \x01(\v20.google.ads.googleads.v20.resources.AdGroupAssetH\x00R\x06update\x12D\n" +
	"\x06remove\x18\x02 \x01(\tB*\xfaA'\n" +
	"%googleads.googleapis.com/AdGroupAssetH\x00R\x06removeB\v\n" +
	"\toperation\"\xbc\x01\n" +
	"\x1bMutateAdGroupAssetsResponse\x12F\n" +
	"\x15partial_failure_error\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x13partialFailureError\x12U\n" +
	"\aresults\x18\x02 \x03(\v2;.google.ads.googleads.v20.services.MutateAdGroupAssetResultR\aresults\"\xc3\x01\n" +
	"\x18MutateAdGroupAssetResult\x12O\n" +
	"\rresource_name\x18\x01 \x01(\tB*\xfaA'\n" +
	"%googleads.googleapis.com/AdGroupAssetR\fresourceName\x12V\n" +
	"\x0ead_group_asset\x18\x02 \x01(\v20.google.ads.googleads.v20.resources.AdGroupAssetR\fadGroupAsset2\xcc\x02\n" +
	"\x13AdGroupAssetService\x12\xed\x01\n" +
	"\x13MutateAdGroupAssets\x12=.google.ads.googleads.v20.services.MutateAdGroupAssetsRequest\x1a>.google.ads.googleads.v20.services.MutateAdGroupAssetsResponse\"W\xdaA\x16customer_id,operations\x82\xd3\xe4\x93\x028:\x01*\"3/v20/customers/{customer_id=*}/adGroupAssets:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x84\x02\n" +
	"%com.google.ads.googleads.v20.servicesB\x18AdGroupAssetServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_ad_group_asset_service_proto_goTypes = []any{
	(*MutateAdGroupAssetsRequest)(nil),                     // 0: google.ads.googleads.v20.services.MutateAdGroupAssetsRequest
	(*AdGroupAssetOperation)(nil),                          // 1: google.ads.googleads.v20.services.AdGroupAssetOperation
	(*MutateAdGroupAssetsResponse)(nil),                    // 2: google.ads.googleads.v20.services.MutateAdGroupAssetsResponse
	(*MutateAdGroupAssetResult)(nil),                       // 3: google.ads.googleads.v20.services.MutateAdGroupAssetResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.AdGroupAsset)(nil),                         // 6: google.ads.googleads.v20.resources.AdGroupAsset
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v20_services_ad_group_asset_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.MutateAdGroupAssetsRequest.operations:type_name -> google.ads.googleads.v20.services.AdGroupAssetOperation
	4, // 1: google.ads.googleads.v20.services.MutateAdGroupAssetsRequest.response_content_type:type_name -> google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v20.services.AdGroupAssetOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v20.services.AdGroupAssetOperation.create:type_name -> google.ads.googleads.v20.resources.AdGroupAsset
	6, // 4: google.ads.googleads.v20.services.AdGroupAssetOperation.update:type_name -> google.ads.googleads.v20.resources.AdGroupAsset
	7, // 5: google.ads.googleads.v20.services.MutateAdGroupAssetsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v20.services.MutateAdGroupAssetsResponse.results:type_name -> google.ads.googleads.v20.services.MutateAdGroupAssetResult
	6, // 7: google.ads.googleads.v20.services.MutateAdGroupAssetResult.ad_group_asset:type_name -> google.ads.googleads.v20.resources.AdGroupAsset
	0, // 8: google.ads.googleads.v20.services.AdGroupAssetService.MutateAdGroupAssets:input_type -> google.ads.googleads.v20.services.MutateAdGroupAssetsRequest
	2, // 9: google.ads.googleads.v20.services.AdGroupAssetService.MutateAdGroupAssets:output_type -> google.ads.googleads.v20.services.MutateAdGroupAssetsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_ad_group_asset_service_proto_init() }
func file_google_ads_googleads_v20_services_ad_group_asset_service_proto_init() {
	if File_google_ads_googleads_v20_services_ad_group_asset_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes[1].OneofWrappers = []any{
		(*AdGroupAssetOperation_Create)(nil),
		(*AdGroupAssetOperation_Update)(nil),
		(*AdGroupAssetOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_ad_group_asset_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_ad_group_asset_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_ad_group_asset_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_ad_group_asset_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_ad_group_asset_service_proto = out.File
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_ad_group_asset_service_proto_depIdxs = nil
}
