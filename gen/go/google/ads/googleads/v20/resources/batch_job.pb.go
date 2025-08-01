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
// source: google/ads/googleads/v20/resources/batch_job.proto

package resources

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
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

// A list of mutates being processed asynchronously. The mutates are uploaded
// by the user. The mutates themselves aren't readable and the results of the
// job can only be read using BatchJobService.ListBatchJobResults.
type BatchJob struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the batch job.
	// Batch job resource names have the form:
	//
	// `customers/{customer_id}/batchJobs/{batch_job_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this batch job.
	Id *int64 `protobuf:"varint,7,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. The next sequence token to use when adding operations. Only
	// set when the batch job status is PENDING.
	NextAddSequenceToken *string `protobuf:"bytes,8,opt,name=next_add_sequence_token,json=nextAddSequenceToken,proto3,oneof" json:"next_add_sequence_token,omitempty"`
	// Output only. Contains additional information about this batch job.
	Metadata *BatchJob_BatchJobMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Status of this batch job.
	Status enums.BatchJobStatusEnum_BatchJobStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.BatchJobStatusEnum_BatchJobStatus" json:"status,omitempty"`
	// Output only. The resource name of the long-running operation that can be
	// used to poll for completion. Only set when the batch job status is RUNNING
	// or DONE.
	LongRunningOperation *string `protobuf:"bytes,9,opt,name=long_running_operation,json=longRunningOperation,proto3,oneof" json:"long_running_operation,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *BatchJob) Reset() {
	*x = BatchJob{}
	mi := &file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob) ProtoMessage() {}

func (x *BatchJob) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob.ProtoReflect.Descriptor instead.
func (*BatchJob) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_batch_job_proto_rawDescGZIP(), []int{0}
}

func (x *BatchJob) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BatchJob) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BatchJob) GetNextAddSequenceToken() string {
	if x != nil && x.NextAddSequenceToken != nil {
		return *x.NextAddSequenceToken
	}
	return ""
}

func (x *BatchJob) GetMetadata() *BatchJob_BatchJobMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BatchJob) GetStatus() enums.BatchJobStatusEnum_BatchJobStatus {
	if x != nil {
		return x.Status
	}
	return enums.BatchJobStatusEnum_BatchJobStatus(0)
}

func (x *BatchJob) GetLongRunningOperation() string {
	if x != nil && x.LongRunningOperation != nil {
		return *x.LongRunningOperation
	}
	return ""
}

// Additional information about the batch job. This message is also used as
// metadata returned in batch job Long Running Operations.
type BatchJob_BatchJobMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The time when this batch job was created.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:00"
	CreationDateTime *string `protobuf:"bytes,8,opt,name=creation_date_time,json=creationDateTime,proto3,oneof" json:"creation_date_time,omitempty"`
	// Output only. The time when this batch job started running.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:30"
	StartDateTime *string `protobuf:"bytes,7,opt,name=start_date_time,json=startDateTime,proto3,oneof" json:"start_date_time,omitempty"`
	// Output only. The time when this batch job was completed.
	// Formatted as yyyy-MM-dd HH:mm:ss. Example: "2018-03-05 09:16:00"
	CompletionDateTime *string `protobuf:"bytes,9,opt,name=completion_date_time,json=completionDateTime,proto3,oneof" json:"completion_date_time,omitempty"`
	// Output only. The fraction (between 0.0 and 1.0) of mutates that have been
	// processed. This is empty if the job hasn't started running yet.
	EstimatedCompletionRatio *float64 `protobuf:"fixed64,10,opt,name=estimated_completion_ratio,json=estimatedCompletionRatio,proto3,oneof" json:"estimated_completion_ratio,omitempty"`
	// Output only. The number of mutate operations in the batch job.
	OperationCount *int64 `protobuf:"varint,11,opt,name=operation_count,json=operationCount,proto3,oneof" json:"operation_count,omitempty"`
	// Output only. The number of mutate operations executed by the batch job.
	// Present only if the job has started running.
	ExecutedOperationCount *int64 `protobuf:"varint,12,opt,name=executed_operation_count,json=executedOperationCount,proto3,oneof" json:"executed_operation_count,omitempty"`
	// Immutable. The approximate upper bound for how long a batch job can be
	// executed, in seconds. If the job runs more than the given upper bound,
	// the job will be canceled.
	ExecutionLimitSeconds *int32 `protobuf:"varint,13,opt,name=execution_limit_seconds,json=executionLimitSeconds,proto3,oneof" json:"execution_limit_seconds,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *BatchJob_BatchJobMetadata) Reset() {
	*x = BatchJob_BatchJobMetadata{}
	mi := &file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchJob_BatchJobMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchJob_BatchJobMetadata) ProtoMessage() {}

func (x *BatchJob_BatchJobMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchJob_BatchJobMetadata.ProtoReflect.Descriptor instead.
func (*BatchJob_BatchJobMetadata) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_batch_job_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BatchJob_BatchJobMetadata) GetCreationDateTime() string {
	if x != nil && x.CreationDateTime != nil {
		return *x.CreationDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetStartDateTime() string {
	if x != nil && x.StartDateTime != nil {
		return *x.StartDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetCompletionDateTime() string {
	if x != nil && x.CompletionDateTime != nil {
		return *x.CompletionDateTime
	}
	return ""
}

func (x *BatchJob_BatchJobMetadata) GetEstimatedCompletionRatio() float64 {
	if x != nil && x.EstimatedCompletionRatio != nil {
		return *x.EstimatedCompletionRatio
	}
	return 0
}

func (x *BatchJob_BatchJobMetadata) GetOperationCount() int64 {
	if x != nil && x.OperationCount != nil {
		return *x.OperationCount
	}
	return 0
}

func (x *BatchJob_BatchJobMetadata) GetExecutedOperationCount() int64 {
	if x != nil && x.ExecutedOperationCount != nil {
		return *x.ExecutedOperationCount
	}
	return 0
}

func (x *BatchJob_BatchJobMetadata) GetExecutionLimitSeconds() int32 {
	if x != nil && x.ExecutionLimitSeconds != nil {
		return *x.ExecutionLimitSeconds
	}
	return 0
}

var File_google_ads_googleads_v20_resources_batch_job_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_batch_job_proto_rawDesc = "" +
	"\n" +
	"2google/ads/googleads/v20/resources/batch_job.proto\x12\"google.ads.googleads.v20.resources\x1a5google/ads/googleads/v20/enums/batch_job_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xb9\t\n" +
	"\bBatchJob\x12N\n" +
	"\rresource_name\x18\x01 \x01(\tB)\xe0A\x05\xfaA#\n" +
	"!googleads.googleapis.com/BatchJobR\fresourceName\x12\x18\n" +
	"\x02id\x18\a \x01(\x03B\x03\xe0A\x03H\x00R\x02id\x88\x01\x01\x12?\n" +
	"\x17next_add_sequence_token\x18\b \x01(\tB\x03\xe0A\x03H\x01R\x14nextAddSequenceToken\x88\x01\x01\x12^\n" +
	"\bmetadata\x18\x04 \x01(\v2=.google.ads.googleads.v20.resources.BatchJob.BatchJobMetadataB\x03\xe0A\x03R\bmetadata\x12^\n" +
	"\x06status\x18\x05 \x01(\x0e2A.google.ads.googleads.v20.enums.BatchJobStatusEnum.BatchJobStatusB\x03\xe0A\x03R\x06status\x12>\n" +
	"\x16long_running_operation\x18\t \x01(\tB\x03\xe0A\x03H\x02R\x14longRunningOperation\x88\x01\x01\x1a\xe9\x04\n" +
	"\x10BatchJobMetadata\x126\n" +
	"\x12creation_date_time\x18\b \x01(\tB\x03\xe0A\x03H\x00R\x10creationDateTime\x88\x01\x01\x120\n" +
	"\x0fstart_date_time\x18\a \x01(\tB\x03\xe0A\x03H\x01R\rstartDateTime\x88\x01\x01\x12:\n" +
	"\x14completion_date_time\x18\t \x01(\tB\x03\xe0A\x03H\x02R\x12completionDateTime\x88\x01\x01\x12F\n" +
	"\x1aestimated_completion_ratio\x18\n" +
	" \x01(\x01B\x03\xe0A\x03H\x03R\x18estimatedCompletionRatio\x88\x01\x01\x121\n" +
	"\x0foperation_count\x18\v \x01(\x03B\x03\xe0A\x03H\x04R\x0eoperationCount\x88\x01\x01\x12B\n" +
	"\x18executed_operation_count\x18\f \x01(\x03B\x03\xe0A\x03H\x05R\x16executedOperationCount\x88\x01\x01\x12@\n" +
	"\x17execution_limit_seconds\x18\r \x01(\x05B\x03\xe0A\x05H\x06R\x15executionLimitSeconds\x88\x01\x01B\x15\n" +
	"\x13_creation_date_timeB\x12\n" +
	"\x10_start_date_timeB\x17\n" +
	"\x15_completion_date_timeB\x1d\n" +
	"\x1b_estimated_completion_ratioB\x12\n" +
	"\x10_operation_countB\x1b\n" +
	"\x19_executed_operation_countB\x1a\n" +
	"\x18_execution_limit_seconds:X\xeaAU\n" +
	"!googleads.googleapis.com/BatchJob\x120customers/{customer_id}/batchJobs/{batch_job_id}B\x05\n" +
	"\x03_idB\x1a\n" +
	"\x18_next_add_sequence_tokenB\x19\n" +
	"\x17_long_running_operationB\xff\x01\n" +
	"&com.google.ads.googleads.v20.resourcesB\rBatchJobProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_batch_job_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_batch_job_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_batch_job_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_batch_job_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_batch_job_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_batch_job_proto_rawDesc), len(file_google_ads_googleads_v20_resources_batch_job_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_batch_job_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v20_resources_batch_job_proto_goTypes = []any{
	(*BatchJob)(nil),                             // 0: google.ads.googleads.v20.resources.BatchJob
	(*BatchJob_BatchJobMetadata)(nil),            // 1: google.ads.googleads.v20.resources.BatchJob.BatchJobMetadata
	(enums.BatchJobStatusEnum_BatchJobStatus)(0), // 2: google.ads.googleads.v20.enums.BatchJobStatusEnum.BatchJobStatus
}
var file_google_ads_googleads_v20_resources_batch_job_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.BatchJob.metadata:type_name -> google.ads.googleads.v20.resources.BatchJob.BatchJobMetadata
	2, // 1: google.ads.googleads.v20.resources.BatchJob.status:type_name -> google.ads.googleads.v20.enums.BatchJobStatusEnum.BatchJobStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_batch_job_proto_init() }
func file_google_ads_googleads_v20_resources_batch_job_proto_init() {
	if File_google_ads_googleads_v20_resources_batch_job_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[0].OneofWrappers = []any{}
	file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_batch_job_proto_rawDesc), len(file_google_ads_googleads_v20_resources_batch_job_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_batch_job_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_batch_job_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_batch_job_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_batch_job_proto = out.File
	file_google_ads_googleads_v20_resources_batch_job_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_batch_job_proto_depIdxs = nil
}
