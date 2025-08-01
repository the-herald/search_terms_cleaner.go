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
// source: google/ads/googleads/v20/services/asset_group_listing_group_filter_service.proto

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
// [AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters][google.ads.googleads.v20.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters].
// partial_failure is not supported because the tree needs to be validated
// together.
type MutateAssetGroupListingGroupFiltersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer whose asset group listing group filters
	// are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual asset group
	// listing group filters.
	Operations []*AssetGroupListingGroupFilterOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,4,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v20.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *MutateAssetGroupListingGroupFiltersRequest) Reset() {
	*x = MutateAssetGroupListingGroupFiltersRequest{}
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAssetGroupListingGroupFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFiltersRequest) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFiltersRequest.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFiltersRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetOperations() []*AssetGroupListingGroupFilterOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateAssetGroupListingGroupFiltersRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove) on an asset group listing group filter.
type AssetGroupListingGroupFilterOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*AssetGroupListingGroupFilterOperation_Create
	//	*AssetGroupListingGroupFilterOperation_Update
	//	*AssetGroupListingGroupFilterOperation_Remove
	Operation     isAssetGroupListingGroupFilterOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetGroupListingGroupFilterOperation) Reset() {
	*x = AssetGroupListingGroupFilterOperation{}
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetGroupListingGroupFilterOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupListingGroupFilterOperation) ProtoMessage() {}

func (x *AssetGroupListingGroupFilterOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupListingGroupFilterOperation.ProtoReflect.Descriptor instead.
func (*AssetGroupListingGroupFilterOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{1}
}

func (x *AssetGroupListingGroupFilterOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetOperation() isAssetGroupListingGroupFilterOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetCreate() *resources.AssetGroupListingGroupFilter {
	if x != nil {
		if x, ok := x.Operation.(*AssetGroupListingGroupFilterOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetUpdate() *resources.AssetGroupListingGroupFilter {
	if x != nil {
		if x, ok := x.Operation.(*AssetGroupListingGroupFilterOperation_Update); ok {
			return x.Update
		}
	}
	return nil
}

func (x *AssetGroupListingGroupFilterOperation) GetRemove() string {
	if x != nil {
		if x, ok := x.Operation.(*AssetGroupListingGroupFilterOperation_Remove); ok {
			return x.Remove
		}
	}
	return ""
}

type isAssetGroupListingGroupFilterOperation_Operation interface {
	isAssetGroupListingGroupFilterOperation_Operation()
}

type AssetGroupListingGroupFilterOperation_Create struct {
	// Create operation: No resource name is expected for the new asset group
	// listing group filter.
	Create *resources.AssetGroupListingGroupFilter `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AssetGroupListingGroupFilterOperation_Update struct {
	// Update operation: The asset group listing group filter is expected to
	// have a valid resource name.
	Update *resources.AssetGroupListingGroupFilter `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AssetGroupListingGroupFilterOperation_Remove struct {
	// Remove operation: A resource name for the removed asset group listing
	// group filter is expected, in this format:
	// `customers/{customer_id}/assetGroupListingGroupFilters/{asset_group_id}~{listing_group_filter_id}`
	// An entity can be removed only if it's not referenced by other
	// parent_listing_group_id. If multiple entities are being deleted, the
	// mutates must be in the correct order.
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AssetGroupListingGroupFilterOperation_Create) isAssetGroupListingGroupFilterOperation_Operation() {
}

func (*AssetGroupListingGroupFilterOperation_Update) isAssetGroupListingGroupFilterOperation_Operation() {
}

func (*AssetGroupListingGroupFilterOperation_Remove) isAssetGroupListingGroupFilterOperation_Operation() {
}

// Response message for an asset group listing group filter mutate.
type MutateAssetGroupListingGroupFiltersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// All results for the mutate.
	Results       []*MutateAssetGroupListingGroupFilterResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MutateAssetGroupListingGroupFiltersResponse) Reset() {
	*x = MutateAssetGroupListingGroupFiltersResponse{}
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAssetGroupListingGroupFiltersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFiltersResponse) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFiltersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFiltersResponse.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFiltersResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateAssetGroupListingGroupFiltersResponse) GetResults() []*MutateAssetGroupListingGroupFilterResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the asset group listing group filter mutate.
type MutateAssetGroupListingGroupFilterResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated AssetGroupListingGroupFilter with only mutable fields after
	// mutate. The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	AssetGroupListingGroupFilter *resources.AssetGroupListingGroupFilter `protobuf:"bytes,2,opt,name=asset_group_listing_group_filter,json=assetGroupListingGroupFilter,proto3" json:"asset_group_listing_group_filter,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *MutateAssetGroupListingGroupFilterResult) Reset() {
	*x = MutateAssetGroupListingGroupFilterResult{}
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MutateAssetGroupListingGroupFilterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateAssetGroupListingGroupFilterResult) ProtoMessage() {}

func (x *MutateAssetGroupListingGroupFilterResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateAssetGroupListingGroupFilterResult.ProtoReflect.Descriptor instead.
func (*MutateAssetGroupListingGroupFilterResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateAssetGroupListingGroupFilterResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateAssetGroupListingGroupFilterResult) GetAssetGroupListingGroupFilter() *resources.AssetGroupListingGroupFilter {
	if x != nil {
		return x.AssetGroupListingGroupFilter
	}
	return nil
}

var File_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDesc = "" +
	"\n" +
	"Pgoogle/ads/googleads/v20/services/asset_group_listing_group_filter_service.proto\x12!google.ads.googleads.v20.services\x1a:google/ads/googleads/v20/enums/response_content_type.proto\x1aIgoogle/ads/googleads/v20/resources/asset_group_listing_group_filter.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\"\xe7\x02\n" +
	"*MutateAssetGroupListingGroupFiltersRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12m\n" +
	"\n" +
	"operations\x18\x02 \x03(\v2H.google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperationB\x03\xe0A\x02R\n" +
	"operations\x12#\n" +
	"\rvalidate_only\x18\x03 \x01(\bR\fvalidateOnly\x12\x7f\n" +
	"\x15response_content_type\x18\x04 \x01(\x0e2K.google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentTypeR\x13responseContentType\"\xff\x02\n" +
	"%AssetGroupListingGroupFilterOperation\x12;\n" +
	"\vupdate_mask\x18\x04 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12Z\n" +
	"\x06create\x18\x01 \x01(\v2@.google.ads.googleads.v20.resources.AssetGroupListingGroupFilterH\x00R\x06create\x12Z\n" +
	"\x06update\x18\x02 \x01(\v2@.google.ads.googleads.v20.resources.AssetGroupListingGroupFilterH\x00R\x06update\x12T\n" +
	"\x06remove\x18\x03 \x01(\tB:\xfaA7\n" +
	"5googleads.googleapis.com/AssetGroupListingGroupFilterH\x00R\x06removeB\v\n" +
	"\toperation\"\x94\x01\n" +
	"+MutateAssetGroupListingGroupFiltersResponse\x12e\n" +
	"\aresults\x18\x01 \x03(\v2K.google.ads.googleads.v20.services.MutateAssetGroupListingGroupFilterResultR\aresults\"\x96\x02\n" +
	"(MutateAssetGroupListingGroupFilterResult\x12_\n" +
	"\rresource_name\x18\x01 \x01(\tB:\xfaA7\n" +
	"5googleads.googleapis.com/AssetGroupListingGroupFilterR\fresourceName\x12\x88\x01\n" +
	" asset_group_listing_group_filter\x18\x02 \x01(\v2@.google.ads.googleads.v20.resources.AssetGroupListingGroupFilterR\x1cassetGroupListingGroupFilter2\x9c\x03\n" +
	"#AssetGroupListingGroupFilterService\x12\xad\x02\n" +
	"#MutateAssetGroupListingGroupFilters\x12M.google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersRequest\x1aN.google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersResponse\"g\xdaA\x16customer_id,operations\x82\xd3\xe4\x93\x02H:\x01*\"C/v20/customers/{customer_id=*}/assetGroupListingGroupFilters:mutate\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x94\x02\n" +
	"%com.google.ads.googleads.v20.servicesB(AssetGroupListingGroupFilterServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_goTypes = []any{
	(*MutateAssetGroupListingGroupFiltersRequest)(nil),     // 0: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersRequest
	(*AssetGroupListingGroupFilterOperation)(nil),          // 1: google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperation
	(*MutateAssetGroupListingGroupFiltersResponse)(nil),    // 2: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersResponse
	(*MutateAssetGroupListingGroupFilterResult)(nil),       // 3: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFilterResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.AssetGroupListingGroupFilter)(nil),         // 6: google.ads.googleads.v20.resources.AssetGroupListingGroupFilter
}
var file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersRequest.operations:type_name -> google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperation
	4, // 1: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersRequest.response_content_type:type_name -> google.ads.googleads.v20.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperation.create:type_name -> google.ads.googleads.v20.resources.AssetGroupListingGroupFilter
	6, // 4: google.ads.googleads.v20.services.AssetGroupListingGroupFilterOperation.update:type_name -> google.ads.googleads.v20.resources.AssetGroupListingGroupFilter
	3, // 5: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersResponse.results:type_name -> google.ads.googleads.v20.services.MutateAssetGroupListingGroupFilterResult
	6, // 6: google.ads.googleads.v20.services.MutateAssetGroupListingGroupFilterResult.asset_group_listing_group_filter:type_name -> google.ads.googleads.v20.resources.AssetGroupListingGroupFilter
	0, // 7: google.ads.googleads.v20.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters:input_type -> google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersRequest
	2, // 8: google.ads.googleads.v20.services.AssetGroupListingGroupFilterService.MutateAssetGroupListingGroupFilters:output_type -> google.ads.googleads.v20.services.MutateAssetGroupListingGroupFiltersResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_init()
}
func file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_init() {
	if File_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes[1].OneofWrappers = []any{
		(*AssetGroupListingGroupFilterOperation_Create)(nil),
		(*AssetGroupListingGroupFilterOperation_Update)(nil),
		(*AssetGroupListingGroupFilterOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto = out.File
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_asset_group_listing_group_filter_service_proto_depIdxs = nil
}
