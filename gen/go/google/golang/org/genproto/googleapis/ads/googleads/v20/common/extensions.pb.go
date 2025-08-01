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
// source: google/ads/googleads/v20/common/extensions.proto

package common

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

// Represents a Call extension.
type CallFeedItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The advertiser's phone number to append to the ad.
	// This string must not be empty.
	PhoneNumber *string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	// Uppercase two-letter country code of the advertiser's phone number.
	// This string must not be empty.
	CountryCode *string `protobuf:"bytes,8,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
	// Indicates whether call tracking is enabled. By default, call tracking is
	// not enabled.
	CallTrackingEnabled *bool `protobuf:"varint,9,opt,name=call_tracking_enabled,json=callTrackingEnabled,proto3,oneof" json:"call_tracking_enabled,omitempty"`
	// The conversion action to attribute a call conversion to. If not set a
	// default conversion action is used. This field only has effect if
	// call_tracking_enabled is set to true. Otherwise this field is ignored.
	CallConversionAction *string `protobuf:"bytes,10,opt,name=call_conversion_action,json=callConversionAction,proto3,oneof" json:"call_conversion_action,omitempty"`
	// If true, disable call conversion tracking. call_conversion_action should
	// not be set if this is true. Optional.
	CallConversionTrackingDisabled *bool `protobuf:"varint,11,opt,name=call_conversion_tracking_disabled,json=callConversionTrackingDisabled,proto3,oneof" json:"call_conversion_tracking_disabled,omitempty"`
	// Enum value that indicates whether this call extension uses its own call
	// conversion setting (or just have call conversion disabled), or following
	// the account level setting.
	CallConversionReportingState enums.CallConversionReportingStateEnum_CallConversionReportingState `protobuf:"varint,6,opt,name=call_conversion_reporting_state,json=callConversionReportingState,proto3,enum=google.ads.googleads.v20.enums.CallConversionReportingStateEnum_CallConversionReportingState" json:"call_conversion_reporting_state,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *CallFeedItem) Reset() {
	*x = CallFeedItem{}
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallFeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallFeedItem) ProtoMessage() {}

func (x *CallFeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallFeedItem.ProtoReflect.Descriptor instead.
func (*CallFeedItem) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *CallFeedItem) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *CallFeedItem) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

func (x *CallFeedItem) GetCallTrackingEnabled() bool {
	if x != nil && x.CallTrackingEnabled != nil {
		return *x.CallTrackingEnabled
	}
	return false
}

func (x *CallFeedItem) GetCallConversionAction() string {
	if x != nil && x.CallConversionAction != nil {
		return *x.CallConversionAction
	}
	return ""
}

func (x *CallFeedItem) GetCallConversionTrackingDisabled() bool {
	if x != nil && x.CallConversionTrackingDisabled != nil {
		return *x.CallConversionTrackingDisabled
	}
	return false
}

func (x *CallFeedItem) GetCallConversionReportingState() enums.CallConversionReportingStateEnum_CallConversionReportingState {
	if x != nil {
		return x.CallConversionReportingState
	}
	return enums.CallConversionReportingStateEnum_CallConversionReportingState(0)
}

// Represents a callout extension.
type CalloutFeedItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The callout text.
	// The length of this string should be between 1 and 25, inclusive.
	CalloutText   *string `protobuf:"bytes,2,opt,name=callout_text,json=calloutText,proto3,oneof" json:"callout_text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalloutFeedItem) Reset() {
	*x = CalloutFeedItem{}
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalloutFeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalloutFeedItem) ProtoMessage() {}

func (x *CalloutFeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalloutFeedItem.ProtoReflect.Descriptor instead.
func (*CalloutFeedItem) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_extensions_proto_rawDescGZIP(), []int{1}
}

func (x *CalloutFeedItem) GetCalloutText() string {
	if x != nil && x.CalloutText != nil {
		return *x.CalloutText
	}
	return ""
}

// Represents a sitelink.
type SitelinkFeedItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// URL display text for the sitelink.
	// The length of this string should be between 1 and 25, inclusive.
	LinkText *string `protobuf:"bytes,9,opt,name=link_text,json=linkText,proto3,oneof" json:"link_text,omitempty"`
	// First line of the description for the sitelink.
	// If this value is set, line2 must also be set.
	// The length of this string should be between 0 and 35, inclusive.
	Line1 *string `protobuf:"bytes,10,opt,name=line1,proto3,oneof" json:"line1,omitempty"`
	// Second line of the description for the sitelink.
	// If this value is set, line1 must also be set.
	// The length of this string should be between 0 and 35, inclusive.
	Line2 *string `protobuf:"bytes,11,opt,name=line2,proto3,oneof" json:"line2,omitempty"`
	// A list of possible final URLs after all cross domain redirects.
	FinalUrls []string `protobuf:"bytes,12,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of possible final mobile URLs after all cross domain redirects.
	FinalMobileUrls []string `protobuf:"bytes,13,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// URL template for constructing a tracking URL.
	TrackingUrlTemplate *string `protobuf:"bytes,14,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3,oneof" json:"tracking_url_template,omitempty"`
	// A list of mappings to be used for substituting URL custom parameter tags in
	// the tracking_url_template, final_urls, and/or final_mobile_urls.
	UrlCustomParameters []*CustomParameter `protobuf:"bytes,7,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// Final URL suffix to be appended to landing page URLs served with
	// parallel tracking.
	FinalUrlSuffix *string `protobuf:"bytes,15,opt,name=final_url_suffix,json=finalUrlSuffix,proto3,oneof" json:"final_url_suffix,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SitelinkFeedItem) Reset() {
	*x = SitelinkFeedItem{}
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SitelinkFeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SitelinkFeedItem) ProtoMessage() {}

func (x *SitelinkFeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_extensions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SitelinkFeedItem.ProtoReflect.Descriptor instead.
func (*SitelinkFeedItem) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_extensions_proto_rawDescGZIP(), []int{2}
}

func (x *SitelinkFeedItem) GetLinkText() string {
	if x != nil && x.LinkText != nil {
		return *x.LinkText
	}
	return ""
}

func (x *SitelinkFeedItem) GetLine1() string {
	if x != nil && x.Line1 != nil {
		return *x.Line1
	}
	return ""
}

func (x *SitelinkFeedItem) GetLine2() string {
	if x != nil && x.Line2 != nil {
		return *x.Line2
	}
	return ""
}

func (x *SitelinkFeedItem) GetFinalUrls() []string {
	if x != nil {
		return x.FinalUrls
	}
	return nil
}

func (x *SitelinkFeedItem) GetFinalMobileUrls() []string {
	if x != nil {
		return x.FinalMobileUrls
	}
	return nil
}

func (x *SitelinkFeedItem) GetTrackingUrlTemplate() string {
	if x != nil && x.TrackingUrlTemplate != nil {
		return *x.TrackingUrlTemplate
	}
	return ""
}

func (x *SitelinkFeedItem) GetUrlCustomParameters() []*CustomParameter {
	if x != nil {
		return x.UrlCustomParameters
	}
	return nil
}

func (x *SitelinkFeedItem) GetFinalUrlSuffix() string {
	if x != nil && x.FinalUrlSuffix != nil {
		return *x.FinalUrlSuffix
	}
	return ""
}

var File_google_ads_googleads_v20_common_extensions_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_common_extensions_proto_rawDesc = "" +
	"\n" +
	"0google/ads/googleads/v20/common/extensions.proto\x12\x1fgoogle.ads.googleads.v20.common\x1a6google/ads/googleads/v20/common/custom_parameter.proto\x1aDgoogle/ads/googleads/v20/enums/call_conversion_reporting_state.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xc6\x04\n" +
	"\fCallFeedItem\x12&\n" +
	"\fphone_number\x18\a \x01(\tH\x00R\vphoneNumber\x88\x01\x01\x12&\n" +
	"\fcountry_code\x18\b \x01(\tH\x01R\vcountryCode\x88\x01\x01\x127\n" +
	"\x15call_tracking_enabled\x18\t \x01(\bH\x02R\x13callTrackingEnabled\x88\x01\x01\x129\n" +
	"\x16call_conversion_action\x18\n" +
	" \x01(\tH\x03R\x14callConversionAction\x88\x01\x01\x12N\n" +
	"!call_conversion_tracking_disabled\x18\v \x01(\bH\x04R\x1ecallConversionTrackingDisabled\x88\x01\x01\x12\xa4\x01\n" +
	"\x1fcall_conversion_reporting_state\x18\x06 \x01(\x0e2].google.ads.googleads.v20.enums.CallConversionReportingStateEnum.CallConversionReportingStateR\x1ccallConversionReportingStateB\x0f\n" +
	"\r_phone_numberB\x0f\n" +
	"\r_country_codeB\x18\n" +
	"\x16_call_tracking_enabledB\x19\n" +
	"\x17_call_conversion_actionB$\n" +
	"\"_call_conversion_tracking_disabled\"J\n" +
	"\x0fCalloutFeedItem\x12&\n" +
	"\fcallout_text\x18\x02 \x01(\tH\x00R\vcalloutText\x88\x01\x01B\x0f\n" +
	"\r_callout_text\"\xd4\x03\n" +
	"\x10SitelinkFeedItem\x12 \n" +
	"\tlink_text\x18\t \x01(\tH\x00R\blinkText\x88\x01\x01\x12\x19\n" +
	"\x05line1\x18\n" +
	" \x01(\tH\x01R\x05line1\x88\x01\x01\x12\x19\n" +
	"\x05line2\x18\v \x01(\tH\x02R\x05line2\x88\x01\x01\x12\x1d\n" +
	"\n" +
	"final_urls\x18\f \x03(\tR\tfinalUrls\x12*\n" +
	"\x11final_mobile_urls\x18\r \x03(\tR\x0ffinalMobileUrls\x127\n" +
	"\x15tracking_url_template\x18\x0e \x01(\tH\x03R\x13trackingUrlTemplate\x88\x01\x01\x12d\n" +
	"\x15url_custom_parameters\x18\a \x03(\v20.google.ads.googleads.v20.common.CustomParameterR\x13urlCustomParameters\x12-\n" +
	"\x10final_url_suffix\x18\x0f \x01(\tH\x04R\x0efinalUrlSuffix\x88\x01\x01B\f\n" +
	"\n" +
	"_link_textB\b\n" +
	"\x06_line1B\b\n" +
	"\x06_line2B\x18\n" +
	"\x16_tracking_url_templateB\x13\n" +
	"\x11_final_url_suffixB\xef\x01\n" +
	"#com.google.ads.googleads.v20.commonB\x0fExtensionsProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/common;common\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Common\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Common\xea\x02#Google::Ads::GoogleAds::V20::Commonb\x06proto3"

var (
	file_google_ads_googleads_v20_common_extensions_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_common_extensions_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_common_extensions_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_common_extensions_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_common_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_extensions_proto_rawDesc), len(file_google_ads_googleads_v20_common_extensions_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_common_extensions_proto_rawDescData
}

var file_google_ads_googleads_v20_common_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v20_common_extensions_proto_goTypes = []any{
	(*CallFeedItem)(nil),     // 0: google.ads.googleads.v20.common.CallFeedItem
	(*CalloutFeedItem)(nil),  // 1: google.ads.googleads.v20.common.CalloutFeedItem
	(*SitelinkFeedItem)(nil), // 2: google.ads.googleads.v20.common.SitelinkFeedItem
	(enums.CallConversionReportingStateEnum_CallConversionReportingState)(0), // 3: google.ads.googleads.v20.enums.CallConversionReportingStateEnum.CallConversionReportingState
	(*CustomParameter)(nil), // 4: google.ads.googleads.v20.common.CustomParameter
}
var file_google_ads_googleads_v20_common_extensions_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v20.common.CallFeedItem.call_conversion_reporting_state:type_name -> google.ads.googleads.v20.enums.CallConversionReportingStateEnum.CallConversionReportingState
	4, // 1: google.ads.googleads.v20.common.SitelinkFeedItem.url_custom_parameters:type_name -> google.ads.googleads.v20.common.CustomParameter
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_common_extensions_proto_init() }
func file_google_ads_googleads_v20_common_extensions_proto_init() {
	if File_google_ads_googleads_v20_common_extensions_proto != nil {
		return
	}
	file_google_ads_googleads_v20_common_custom_parameter_proto_init()
	file_google_ads_googleads_v20_common_extensions_proto_msgTypes[0].OneofWrappers = []any{}
	file_google_ads_googleads_v20_common_extensions_proto_msgTypes[1].OneofWrappers = []any{}
	file_google_ads_googleads_v20_common_extensions_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_extensions_proto_rawDesc), len(file_google_ads_googleads_v20_common_extensions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_common_extensions_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_common_extensions_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_common_extensions_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_common_extensions_proto = out.File
	file_google_ads_googleads_v20_common_extensions_proto_goTypes = nil
	file_google_ads_googleads_v20_common_extensions_proto_depIdxs = nil
}
