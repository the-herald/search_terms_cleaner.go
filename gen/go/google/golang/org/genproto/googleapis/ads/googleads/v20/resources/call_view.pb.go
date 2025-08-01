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
// source: google/ads/googleads/v20/resources/call_view.proto

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

// A call view that includes data for call tracking of call-only ads or call
// extensions.
type CallView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the call view.
	// Call view resource names have the form:
	//
	// `customers/{customer_id}/callViews/{call_detail_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Country code of the caller.
	CallerCountryCode string `protobuf:"bytes,2,opt,name=caller_country_code,json=callerCountryCode,proto3" json:"caller_country_code,omitempty"`
	// Output only. Area code of the caller. Null if the call duration is shorter
	// than 15 seconds.
	CallerAreaCode string `protobuf:"bytes,3,opt,name=caller_area_code,json=callerAreaCode,proto3" json:"caller_area_code,omitempty"`
	// Output only. The advertiser-provided call duration in seconds.
	CallDurationSeconds int64 `protobuf:"varint,4,opt,name=call_duration_seconds,json=callDurationSeconds,proto3" json:"call_duration_seconds,omitempty"`
	// Output only. The advertiser-provided call start date time.
	StartCallDateTime string `protobuf:"bytes,5,opt,name=start_call_date_time,json=startCallDateTime,proto3" json:"start_call_date_time,omitempty"`
	// Output only. The advertiser-provided call end date time.
	EndCallDateTime string `protobuf:"bytes,6,opt,name=end_call_date_time,json=endCallDateTime,proto3" json:"end_call_date_time,omitempty"`
	// Output only. The call tracking display location.
	CallTrackingDisplayLocation enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation `protobuf:"varint,7,opt,name=call_tracking_display_location,json=callTrackingDisplayLocation,proto3,enum=google.ads.googleads.v20.enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation" json:"call_tracking_display_location,omitempty"`
	// Output only. The type of the call.
	Type enums.CallTypeEnum_CallType `protobuf:"varint,8,opt,name=type,proto3,enum=google.ads.googleads.v20.enums.CallTypeEnum_CallType" json:"type,omitempty"`
	// Output only. The status of the call.
	CallStatus    enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus `protobuf:"varint,9,opt,name=call_status,json=callStatus,proto3,enum=google.ads.googleads.v20.enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus" json:"call_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CallView) Reset() {
	*x = CallView{}
	mi := &file_google_ads_googleads_v20_resources_call_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallView) ProtoMessage() {}

func (x *CallView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_call_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallView.ProtoReflect.Descriptor instead.
func (*CallView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_call_view_proto_rawDescGZIP(), []int{0}
}

func (x *CallView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CallView) GetCallerCountryCode() string {
	if x != nil {
		return x.CallerCountryCode
	}
	return ""
}

func (x *CallView) GetCallerAreaCode() string {
	if x != nil {
		return x.CallerAreaCode
	}
	return ""
}

func (x *CallView) GetCallDurationSeconds() int64 {
	if x != nil {
		return x.CallDurationSeconds
	}
	return 0
}

func (x *CallView) GetStartCallDateTime() string {
	if x != nil {
		return x.StartCallDateTime
	}
	return ""
}

func (x *CallView) GetEndCallDateTime() string {
	if x != nil {
		return x.EndCallDateTime
	}
	return ""
}

func (x *CallView) GetCallTrackingDisplayLocation() enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation {
	if x != nil {
		return x.CallTrackingDisplayLocation
	}
	return enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation(0)
}

func (x *CallView) GetType() enums.CallTypeEnum_CallType {
	if x != nil {
		return x.Type
	}
	return enums.CallTypeEnum_CallType(0)
}

func (x *CallView) GetCallStatus() enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus {
	if x != nil {
		return x.CallStatus
	}
	return enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus(0)
}

var File_google_ads_googleads_v20_resources_call_view_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_call_view_proto_rawDesc = "" +
	"\n" +
	"2google/ads/googleads/v20/resources/call_view.proto\x12\"google.ads.googleads.v20.resources\x1aCgoogle/ads/googleads/v20/enums/call_tracking_display_location.proto\x1a.google/ads/googleads/v20/enums/call_type.proto\x1a=google/ads/googleads/v20/enums/google_voice_call_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xaa\x06\n" +
	"\bCallView\x12N\n" +
	"\rresource_name\x18\x01 \x01(\tB)\xe0A\x03\xfaA#\n" +
	"!googleads.googleapis.com/CallViewR\fresourceName\x123\n" +
	"\x13caller_country_code\x18\x02 \x01(\tB\x03\xe0A\x03R\x11callerCountryCode\x12-\n" +
	"\x10caller_area_code\x18\x03 \x01(\tB\x03\xe0A\x03R\x0ecallerAreaCode\x127\n" +
	"\x15call_duration_seconds\x18\x04 \x01(\x03B\x03\xe0A\x03R\x13callDurationSeconds\x124\n" +
	"\x14start_call_date_time\x18\x05 \x01(\tB\x03\xe0A\x03R\x11startCallDateTime\x120\n" +
	"\x12end_call_date_time\x18\x06 \x01(\tB\x03\xe0A\x03R\x0fendCallDateTime\x12\xa5\x01\n" +
	"\x1ecall_tracking_display_location\x18\a \x01(\x0e2[.google.ads.googleads.v20.enums.CallTrackingDisplayLocationEnum.CallTrackingDisplayLocationB\x03\xe0A\x03R\x1bcallTrackingDisplayLocation\x12N\n" +
	"\x04type\x18\b \x01(\x0e25.google.ads.googleads.v20.enums.CallTypeEnum.CallTypeB\x03\xe0A\x03R\x04type\x12u\n" +
	"\vcall_status\x18\t \x01(\x0e2O.google.ads.googleads.v20.enums.GoogleVoiceCallStatusEnum.GoogleVoiceCallStatusB\x03\xe0A\x03R\n" +
	"callStatus:Z\xeaAW\n" +
	"!googleads.googleapis.com/CallView\x122customers/{customer_id}/callViews/{call_detail_id}B\xff\x01\n" +
	"&com.google.ads.googleads.v20.resourcesB\rCallViewProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_call_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_call_view_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_call_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_call_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_call_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_call_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_call_view_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_call_view_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_call_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_call_view_proto_goTypes = []any{
	(*CallView)(nil), // 0: google.ads.googleads.v20.resources.CallView
	(enums.CallTrackingDisplayLocationEnum_CallTrackingDisplayLocation)(0), // 1: google.ads.googleads.v20.enums.CallTrackingDisplayLocationEnum.CallTrackingDisplayLocation
	(enums.CallTypeEnum_CallType)(0),                                       // 2: google.ads.googleads.v20.enums.CallTypeEnum.CallType
	(enums.GoogleVoiceCallStatusEnum_GoogleVoiceCallStatus)(0),             // 3: google.ads.googleads.v20.enums.GoogleVoiceCallStatusEnum.GoogleVoiceCallStatus
}
var file_google_ads_googleads_v20_resources_call_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.CallView.call_tracking_display_location:type_name -> google.ads.googleads.v20.enums.CallTrackingDisplayLocationEnum.CallTrackingDisplayLocation
	2, // 1: google.ads.googleads.v20.resources.CallView.type:type_name -> google.ads.googleads.v20.enums.CallTypeEnum.CallType
	3, // 2: google.ads.googleads.v20.resources.CallView.call_status:type_name -> google.ads.googleads.v20.enums.GoogleVoiceCallStatusEnum.GoogleVoiceCallStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_call_view_proto_init() }
func file_google_ads_googleads_v20_resources_call_view_proto_init() {
	if File_google_ads_googleads_v20_resources_call_view_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_call_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_call_view_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_call_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_call_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_call_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_call_view_proto = out.File
	file_google_ads_googleads_v20_resources_call_view_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_call_view_proto_depIdxs = nil
}
