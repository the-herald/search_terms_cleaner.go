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
// source: google/ads/googleads/v20/services/offline_user_data_job_service.proto

package services

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	common "google.golang.org/genproto/googleapis/ads/googleads/v20/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v20/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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
// [OfflineUserDataJobService.CreateOfflineUserDataJob][google.ads.googleads.v20.services.OfflineUserDataJobService.CreateOfflineUserDataJob].
type CreateOfflineUserDataJobRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer for which to create an offline user data
	// job.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The offline user data job to be created.
	Job *resources.OfflineUserDataJob `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// If true, match rate range for the offline user data job is calculated and
	// made available in the resource.
	EnableMatchRateRangePreview bool `protobuf:"varint,5,opt,name=enable_match_rate_range_preview,json=enableMatchRateRangePreview,proto3" json:"enable_match_rate_range_preview,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *CreateOfflineUserDataJobRequest) Reset() {
	*x = CreateOfflineUserDataJobRequest{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOfflineUserDataJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfflineUserDataJobRequest) ProtoMessage() {}

func (x *CreateOfflineUserDataJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfflineUserDataJobRequest.ProtoReflect.Descriptor instead.
func (*CreateOfflineUserDataJobRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOfflineUserDataJobRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateOfflineUserDataJobRequest) GetJob() *resources.OfflineUserDataJob {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *CreateOfflineUserDataJobRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *CreateOfflineUserDataJobRequest) GetEnableMatchRateRangePreview() bool {
	if x != nil {
		return x.EnableMatchRateRangePreview
	}
	return false
}

// Response message for
// [OfflineUserDataJobService.CreateOfflineUserDataJob][google.ads.googleads.v20.services.OfflineUserDataJobService.CreateOfflineUserDataJob].
type CreateOfflineUserDataJobResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the OfflineUserDataJob.
	ResourceName  string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOfflineUserDataJobResponse) Reset() {
	*x = CreateOfflineUserDataJobResponse{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOfflineUserDataJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfflineUserDataJobResponse) ProtoMessage() {}

func (x *CreateOfflineUserDataJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfflineUserDataJobResponse.ProtoReflect.Descriptor instead.
func (*CreateOfflineUserDataJobResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOfflineUserDataJobResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [OfflineUserDataJobService.RunOfflineUserDataJob][google.ads.googleads.v20.services.OfflineUserDataJobService.RunOfflineUserDataJob].
type RunOfflineUserDataJobRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The resource name of the OfflineUserDataJob to run.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly  bool `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunOfflineUserDataJobRequest) Reset() {
	*x = RunOfflineUserDataJobRequest{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunOfflineUserDataJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunOfflineUserDataJobRequest) ProtoMessage() {}

func (x *RunOfflineUserDataJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunOfflineUserDataJobRequest.ProtoReflect.Descriptor instead.
func (*RunOfflineUserDataJobRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{2}
}

func (x *RunOfflineUserDataJobRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *RunOfflineUserDataJobRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Request message for
// [OfflineUserDataJobService.AddOfflineUserDataJobOperations][google.ads.googleads.v20.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations].
type AddOfflineUserDataJobOperationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The resource name of the OfflineUserDataJob.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// True to enable partial failure for the offline user data job.
	EnablePartialFailure *bool `protobuf:"varint,4,opt,name=enable_partial_failure,json=enablePartialFailure,proto3,oneof" json:"enable_partial_failure,omitempty"`
	// True to enable warnings for the offline user data job. When enabled, a
	// warning will not block the OfflineUserDataJobOperation, and will also
	// return warning messages about malformed field values.
	EnableWarnings *bool `protobuf:"varint,6,opt,name=enable_warnings,json=enableWarnings,proto3,oneof" json:"enable_warnings,omitempty"`
	// Required. The list of operations to be done.
	Operations []*OfflineUserDataJobOperation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly  bool `protobuf:"varint,5,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOfflineUserDataJobOperationsRequest) Reset() {
	*x = AddOfflineUserDataJobOperationsRequest{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOfflineUserDataJobOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOfflineUserDataJobOperationsRequest) ProtoMessage() {}

func (x *AddOfflineUserDataJobOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOfflineUserDataJobOperationsRequest.ProtoReflect.Descriptor instead.
func (*AddOfflineUserDataJobOperationsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddOfflineUserDataJobOperationsRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AddOfflineUserDataJobOperationsRequest) GetEnablePartialFailure() bool {
	if x != nil && x.EnablePartialFailure != nil {
		return *x.EnablePartialFailure
	}
	return false
}

func (x *AddOfflineUserDataJobOperationsRequest) GetEnableWarnings() bool {
	if x != nil && x.EnableWarnings != nil {
		return *x.EnableWarnings
	}
	return false
}

func (x *AddOfflineUserDataJobOperationsRequest) GetOperations() []*OfflineUserDataJobOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *AddOfflineUserDataJobOperationsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Operation to be made for the AddOfflineUserDataJobOperationsRequest.
type OfflineUserDataJobOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Operation to be made for the AddOfflineUserDataJobOperationsRequest.
	//
	// Types that are valid to be assigned to Operation:
	//
	//	*OfflineUserDataJobOperation_Create
	//	*OfflineUserDataJobOperation_Remove
	//	*OfflineUserDataJobOperation_RemoveAll
	Operation     isOfflineUserDataJobOperation_Operation `protobuf_oneof:"operation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineUserDataJobOperation) Reset() {
	*x = OfflineUserDataJobOperation{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineUserDataJobOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineUserDataJobOperation) ProtoMessage() {}

func (x *OfflineUserDataJobOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineUserDataJobOperation.ProtoReflect.Descriptor instead.
func (*OfflineUserDataJobOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{4}
}

func (x *OfflineUserDataJobOperation) GetOperation() isOfflineUserDataJobOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *OfflineUserDataJobOperation) GetCreate() *common.UserData {
	if x != nil {
		if x, ok := x.Operation.(*OfflineUserDataJobOperation_Create); ok {
			return x.Create
		}
	}
	return nil
}

func (x *OfflineUserDataJobOperation) GetRemove() *common.UserData {
	if x != nil {
		if x, ok := x.Operation.(*OfflineUserDataJobOperation_Remove); ok {
			return x.Remove
		}
	}
	return nil
}

func (x *OfflineUserDataJobOperation) GetRemoveAll() bool {
	if x != nil {
		if x, ok := x.Operation.(*OfflineUserDataJobOperation_RemoveAll); ok {
			return x.RemoveAll
		}
	}
	return false
}

type isOfflineUserDataJobOperation_Operation interface {
	isOfflineUserDataJobOperation_Operation()
}

type OfflineUserDataJobOperation_Create struct {
	// Add the provided data to the transaction. Data cannot be retrieved after
	// being uploaded.
	Create *common.UserData `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type OfflineUserDataJobOperation_Remove struct {
	// Remove the provided data from the transaction. Data cannot be retrieved
	// after being uploaded.
	Remove *common.UserData `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

type OfflineUserDataJobOperation_RemoveAll struct {
	// Remove all previously provided data. This is only supported for Customer
	// Match.
	RemoveAll bool `protobuf:"varint,3,opt,name=remove_all,json=removeAll,proto3,oneof"`
}

func (*OfflineUserDataJobOperation_Create) isOfflineUserDataJobOperation_Operation() {}

func (*OfflineUserDataJobOperation_Remove) isOfflineUserDataJobOperation_Operation() {}

func (*OfflineUserDataJobOperation_RemoveAll) isOfflineUserDataJobOperation_Operation() {}

// Response message for
// [OfflineUserDataJobService.AddOfflineUserDataJobOperations][google.ads.googleads.v20.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations].
type AddOfflineUserDataJobOperationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// Non blocking errors that pertain to operation failures in the warnings
	// mode. Returned only when enable_warnings = true.
	Warning       *status.Status `protobuf:"bytes,2,opt,name=warning,proto3" json:"warning,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddOfflineUserDataJobOperationsResponse) Reset() {
	*x = AddOfflineUserDataJobOperationsResponse{}
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddOfflineUserDataJobOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOfflineUserDataJobOperationsResponse) ProtoMessage() {}

func (x *AddOfflineUserDataJobOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOfflineUserDataJobOperationsResponse.ProtoReflect.Descriptor instead.
func (*AddOfflineUserDataJobOperationsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddOfflineUserDataJobOperationsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *AddOfflineUserDataJobOperationsResponse) GetWarning() *status.Status {
	if x != nil {
		return x.Warning
	}
	return nil
}

var File_google_ads_googleads_v20_services_offline_user_data_job_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDesc = "" +
	"\n" +
	"Egoogle/ads/googleads/v20/services/offline_user_data_job_service.proto\x12!google.ads.googleads.v20.services\x1a7google/ads/googleads/v20/common/offline_user_data.proto\x1a>google/ads/googleads/v20/resources/offline_user_data_job.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a#google/longrunning/operations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x17google/rpc/status.proto\"\x81\x02\n" +
	"\x1fCreateOfflineUserDataJobRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12M\n" +
	"\x03job\x18\x02 \x01(\v26.google.ads.googleads.v20.resources.OfflineUserDataJobB\x03\xe0A\x02R\x03job\x12#\n" +
	"\rvalidate_only\x18\x03 \x01(\bR\fvalidateOnly\x12D\n" +
	"\x1fenable_match_rate_range_preview\x18\x05 \x01(\bR\x1benableMatchRateRangePreview\"y\n" +
	" CreateOfflineUserDataJobResponse\x12U\n" +
	"\rresource_name\x18\x01 \x01(\tB0\xfaA-\n" +
	"+googleads.googleapis.com/OfflineUserDataJobR\fresourceName\"\x9d\x01\n" +
	"\x1cRunOfflineUserDataJobRequest\x12X\n" +
	"\rresource_name\x18\x01 \x01(\tB3\xe0A\x02\xfaA-\n" +
	"+googleads.googleapis.com/OfflineUserDataJobR\fresourceName\x12#\n" +
	"\rvalidate_only\x18\x02 \x01(\bR\fvalidateOnly\"\xa4\x03\n" +
	"&AddOfflineUserDataJobOperationsRequest\x12X\n" +
	"\rresource_name\x18\x01 \x01(\tB3\xe0A\x02\xfaA-\n" +
	"+googleads.googleapis.com/OfflineUserDataJobR\fresourceName\x129\n" +
	"\x16enable_partial_failure\x18\x04 \x01(\bH\x00R\x14enablePartialFailure\x88\x01\x01\x12,\n" +
	"\x0fenable_warnings\x18\x06 \x01(\bH\x01R\x0eenableWarnings\x88\x01\x01\x12c\n" +
	"\n" +
	"operations\x18\x03 \x03(\v2>.google.ads.googleads.v20.services.OfflineUserDataJobOperationB\x03\xe0A\x02R\n" +
	"operations\x12#\n" +
	"\rvalidate_only\x18\x05 \x01(\bR\fvalidateOnlyB\x19\n" +
	"\x17_enable_partial_failureB\x12\n" +
	"\x10_enable_warnings\"\xd5\x01\n" +
	"\x1bOfflineUserDataJobOperation\x12C\n" +
	"\x06create\x18\x01 \x01(\v2).google.ads.googleads.v20.common.UserDataH\x00R\x06create\x12C\n" +
	"\x06remove\x18\x02 \x01(\v2).google.ads.googleads.v20.common.UserDataH\x00R\x06remove\x12\x1f\n" +
	"\n" +
	"remove_all\x18\x03 \x01(\bH\x00R\tremoveAllB\v\n" +
	"\toperation\"\x9f\x01\n" +
	"'AddOfflineUserDataJobOperationsResponse\x12F\n" +
	"\x15partial_failure_error\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x13partialFailureError\x12,\n" +
	"\awarning\x18\x02 \x01(\v2\x12.google.rpc.StatusR\awarning2\xb2\a\n" +
	"\x19OfflineUserDataJobService\x12\xfb\x01\n" +
	"\x18CreateOfflineUserDataJob\x12B.google.ads.googleads.v20.services.CreateOfflineUserDataJobRequest\x1aC.google.ads.googleads.v20.services.CreateOfflineUserDataJobResponse\"V\xdaA\x0fcustomer_id,job\x82\xd3\xe4\x93\x02>:\x01*\"9/v20/customers/{customer_id=*}/offlineUserDataJobs:create\x12\xa4\x02\n" +
	"\x1fAddOfflineUserDataJobOperations\x12I.google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsRequest\x1aJ.google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsResponse\"j\xdaA\x18resource_name,operations\x82\xd3\xe4\x93\x02I:\x01*\"D/v20/{resource_name=customers/*/offlineUserDataJobs/*}:addOperations\x12\xa8\x02\n" +
	"\x15RunOfflineUserDataJob\x12?.google.ads.googleads.v20.services.RunOfflineUserDataJobRequest\x1a\x1d.google.longrunning.Operation\"\xae\x01\xcaAV\n" +
	"\x15google.protobuf.Empty\x12=google.ads.googleads.v20.resources.OfflineUserDataJobMetadata\xdaA\rresource_name\x82\xd3\xe4\x93\x02?:\x01*\":/v20/{resource_name=customers/*/offlineUserDataJobs/*}:run\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x8a\x02\n" +
	"%com.google.ads.googleads.v20.servicesB\x1eOfflineUserDataJobServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_goTypes = []any{
	(*CreateOfflineUserDataJobRequest)(nil),         // 0: google.ads.googleads.v20.services.CreateOfflineUserDataJobRequest
	(*CreateOfflineUserDataJobResponse)(nil),        // 1: google.ads.googleads.v20.services.CreateOfflineUserDataJobResponse
	(*RunOfflineUserDataJobRequest)(nil),            // 2: google.ads.googleads.v20.services.RunOfflineUserDataJobRequest
	(*AddOfflineUserDataJobOperationsRequest)(nil),  // 3: google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsRequest
	(*OfflineUserDataJobOperation)(nil),             // 4: google.ads.googleads.v20.services.OfflineUserDataJobOperation
	(*AddOfflineUserDataJobOperationsResponse)(nil), // 5: google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsResponse
	(*resources.OfflineUserDataJob)(nil),            // 6: google.ads.googleads.v20.resources.OfflineUserDataJob
	(*common.UserData)(nil),                         // 7: google.ads.googleads.v20.common.UserData
	(*status.Status)(nil),                           // 8: google.rpc.Status
	(*longrunningpb.Operation)(nil),                 // 9: google.longrunning.Operation
}
var file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_depIdxs = []int32{
	6, // 0: google.ads.googleads.v20.services.CreateOfflineUserDataJobRequest.job:type_name -> google.ads.googleads.v20.resources.OfflineUserDataJob
	4, // 1: google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsRequest.operations:type_name -> google.ads.googleads.v20.services.OfflineUserDataJobOperation
	7, // 2: google.ads.googleads.v20.services.OfflineUserDataJobOperation.create:type_name -> google.ads.googleads.v20.common.UserData
	7, // 3: google.ads.googleads.v20.services.OfflineUserDataJobOperation.remove:type_name -> google.ads.googleads.v20.common.UserData
	8, // 4: google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsResponse.partial_failure_error:type_name -> google.rpc.Status
	8, // 5: google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsResponse.warning:type_name -> google.rpc.Status
	0, // 6: google.ads.googleads.v20.services.OfflineUserDataJobService.CreateOfflineUserDataJob:input_type -> google.ads.googleads.v20.services.CreateOfflineUserDataJobRequest
	3, // 7: google.ads.googleads.v20.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations:input_type -> google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsRequest
	2, // 8: google.ads.googleads.v20.services.OfflineUserDataJobService.RunOfflineUserDataJob:input_type -> google.ads.googleads.v20.services.RunOfflineUserDataJobRequest
	1, // 9: google.ads.googleads.v20.services.OfflineUserDataJobService.CreateOfflineUserDataJob:output_type -> google.ads.googleads.v20.services.CreateOfflineUserDataJobResponse
	5, // 10: google.ads.googleads.v20.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations:output_type -> google.ads.googleads.v20.services.AddOfflineUserDataJobOperationsResponse
	9, // 11: google.ads.googleads.v20.services.OfflineUserDataJobService.RunOfflineUserDataJob:output_type -> google.longrunning.Operation
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_init() }
func file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_init() {
	if File_google_ads_googleads_v20_services_offline_user_data_job_service_proto != nil {
		return
	}
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[3].OneofWrappers = []any{}
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes[4].OneofWrappers = []any{
		(*OfflineUserDataJobOperation_Create)(nil),
		(*OfflineUserDataJobOperation_Remove)(nil),
		(*OfflineUserDataJobOperation_RemoveAll)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_offline_user_data_job_service_proto = out.File
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_offline_user_data_job_service_proto_depIdxs = nil
}
