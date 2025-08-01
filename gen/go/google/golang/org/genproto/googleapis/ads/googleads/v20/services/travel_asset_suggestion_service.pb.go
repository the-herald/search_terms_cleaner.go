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
// source: google/ads/googleads/v20/services/travel_asset_suggestion_service.proto

package services

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

// Request message for
// [TravelAssetSuggestionService.SuggestTravelAssets][google.ads.googleads.v20.services.TravelAssetSuggestionService.SuggestTravelAssets].
type SuggestTravelAssetsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The ID of the customer.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The language specifications in BCP 47 format (for example, en-US,
	// zh-CN, etc.) for the asset suggestions. Text will be in this language.
	// Usually matches one of the campaign target languages.
	LanguageOption string `protobuf:"bytes,2,opt,name=language_option,json=languageOption,proto3" json:"language_option,omitempty"`
	// The Google Maps Place IDs of hotels for which assets are requested. See
	// https://developers.google.com/places/web-service/place-id for more
	// information.
	PlaceIds      []string `protobuf:"bytes,4,rep,name=place_ids,json=placeIds,proto3" json:"place_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuggestTravelAssetsRequest) Reset() {
	*x = SuggestTravelAssetsRequest{}
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuggestTravelAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestTravelAssetsRequest) ProtoMessage() {}

func (x *SuggestTravelAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestTravelAssetsRequest.ProtoReflect.Descriptor instead.
func (*SuggestTravelAssetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestTravelAssetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *SuggestTravelAssetsRequest) GetLanguageOption() string {
	if x != nil {
		return x.LanguageOption
	}
	return ""
}

func (x *SuggestTravelAssetsRequest) GetPlaceIds() []string {
	if x != nil {
		return x.PlaceIds
	}
	return nil
}

// Response message for
// [TravelAssetSuggestionService.SuggestTravelAssets][google.ads.googleads.v20.services.TravelAssetSuggestionService.SuggestTravelAssets].
type SuggestTravelAssetsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Asset suggestions for each place ID submitted in the request.
	HotelAssetSuggestions []*HotelAssetSuggestion `protobuf:"bytes,1,rep,name=hotel_asset_suggestions,json=hotelAssetSuggestions,proto3" json:"hotel_asset_suggestions,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *SuggestTravelAssetsResponse) Reset() {
	*x = SuggestTravelAssetsResponse{}
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuggestTravelAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestTravelAssetsResponse) ProtoMessage() {}

func (x *SuggestTravelAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestTravelAssetsResponse.ProtoReflect.Descriptor instead.
func (*SuggestTravelAssetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestTravelAssetsResponse) GetHotelAssetSuggestions() []*HotelAssetSuggestion {
	if x != nil {
		return x.HotelAssetSuggestions
	}
	return nil
}

// Message containing the asset suggestions for a hotel.
type HotelAssetSuggestion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Google Places ID of the hotel.
	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
	// Suggested final URL for an AssetGroup.
	FinalUrl string `protobuf:"bytes,2,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	// Hotel name in requested language.
	HotelName string `protobuf:"bytes,3,opt,name=hotel_name,json=hotelName,proto3" json:"hotel_name,omitempty"`
	// Call to action type.
	CallToAction enums.CallToActionTypeEnum_CallToActionType `protobuf:"varint,4,opt,name=call_to_action,json=callToAction,proto3,enum=google.ads.googleads.v20.enums.CallToActionTypeEnum_CallToActionType" json:"call_to_action,omitempty"`
	// Text assets such as headline, description, etc.
	TextAssets []*HotelTextAsset `protobuf:"bytes,5,rep,name=text_assets,json=textAssets,proto3" json:"text_assets,omitempty"`
	// Image assets such as landscape/portrait/square, etc.
	ImageAssets []*HotelImageAsset `protobuf:"bytes,6,rep,name=image_assets,json=imageAssets,proto3" json:"image_assets,omitempty"`
	// The status of the hotel asset suggestion.
	Status        enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus `protobuf:"varint,7,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HotelAssetSuggestion) Reset() {
	*x = HotelAssetSuggestion{}
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelAssetSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelAssetSuggestion) ProtoMessage() {}

func (x *HotelAssetSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelAssetSuggestion.ProtoReflect.Descriptor instead.
func (*HotelAssetSuggestion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{2}
}

func (x *HotelAssetSuggestion) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

func (x *HotelAssetSuggestion) GetFinalUrl() string {
	if x != nil {
		return x.FinalUrl
	}
	return ""
}

func (x *HotelAssetSuggestion) GetHotelName() string {
	if x != nil {
		return x.HotelName
	}
	return ""
}

func (x *HotelAssetSuggestion) GetCallToAction() enums.CallToActionTypeEnum_CallToActionType {
	if x != nil {
		return x.CallToAction
	}
	return enums.CallToActionTypeEnum_CallToActionType(0)
}

func (x *HotelAssetSuggestion) GetTextAssets() []*HotelTextAsset {
	if x != nil {
		return x.TextAssets
	}
	return nil
}

func (x *HotelAssetSuggestion) GetImageAssets() []*HotelImageAsset {
	if x != nil {
		return x.ImageAssets
	}
	return nil
}

func (x *HotelAssetSuggestion) GetStatus() enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus {
	if x != nil {
		return x.Status
	}
	return enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus(0)
}

// A single text asset suggestion for a hotel.
type HotelTextAsset struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Asset text in requested language.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The text asset type. For example, HEADLINE, DESCRIPTION, etc.
	AssetFieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,2,opt,name=asset_field_type,json=assetFieldType,proto3,enum=google.ads.googleads.v20.enums.AssetFieldTypeEnum_AssetFieldType" json:"asset_field_type,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *HotelTextAsset) Reset() {
	*x = HotelTextAsset{}
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelTextAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelTextAsset) ProtoMessage() {}

func (x *HotelTextAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelTextAsset.ProtoReflect.Descriptor instead.
func (*HotelTextAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{3}
}

func (x *HotelTextAsset) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *HotelTextAsset) GetAssetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.AssetFieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

// A single image asset suggestion for a hotel.
type HotelImageAsset struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// URI for the image.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The Image asset type. For example, MARKETING_IMAGE,
	// PORTRAIT_MARKETING_IMAGE, etc.
	AssetFieldType enums.AssetFieldTypeEnum_AssetFieldType `protobuf:"varint,2,opt,name=asset_field_type,json=assetFieldType,proto3,enum=google.ads.googleads.v20.enums.AssetFieldTypeEnum_AssetFieldType" json:"asset_field_type,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *HotelImageAsset) Reset() {
	*x = HotelImageAsset{}
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelImageAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelImageAsset) ProtoMessage() {}

func (x *HotelImageAsset) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelImageAsset.ProtoReflect.Descriptor instead.
func (*HotelImageAsset) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP(), []int{4}
}

func (x *HotelImageAsset) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HotelImageAsset) GetAssetFieldType() enums.AssetFieldTypeEnum_AssetFieldType {
	if x != nil {
		return x.AssetFieldType
	}
	return enums.AssetFieldTypeEnum_AssetFieldType(0)
}

var File_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDesc = "" +
	"\n" +
	"Ggoogle/ads/googleads/v20/services/travel_asset_suggestion_service.proto\x12!google.ads.googleads.v20.services\x1a5google/ads/googleads/v20/enums/asset_field_type.proto\x1a8google/ads/googleads/v20/enums/call_to_action_type.proto\x1aBgoogle/ads/googleads/v20/enums/hotel_asset_suggestion_status.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\"\x8d\x01\n" +
	"\x1aSuggestTravelAssetsRequest\x12$\n" +
	"\vcustomer_id\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"customerId\x12,\n" +
	"\x0flanguage_option\x18\x02 \x01(\tB\x03\xe0A\x02R\x0elanguageOption\x12\x1b\n" +
	"\tplace_ids\x18\x04 \x03(\tR\bplaceIds\"\x8e\x01\n" +
	"\x1bSuggestTravelAssetsResponse\x12o\n" +
	"\x17hotel_asset_suggestions\x18\x01 \x03(\v27.google.ads.googleads.v20.services.HotelAssetSuggestionR\x15hotelAssetSuggestions\"\xf8\x03\n" +
	"\x14HotelAssetSuggestion\x12\x19\n" +
	"\bplace_id\x18\x01 \x01(\tR\aplaceId\x12\x1b\n" +
	"\tfinal_url\x18\x02 \x01(\tR\bfinalUrl\x12\x1d\n" +
	"\n" +
	"hotel_name\x18\x03 \x01(\tR\thotelName\x12k\n" +
	"\x0ecall_to_action\x18\x04 \x01(\x0e2E.google.ads.googleads.v20.enums.CallToActionTypeEnum.CallToActionTypeR\fcallToAction\x12R\n" +
	"\vtext_assets\x18\x05 \x03(\v21.google.ads.googleads.v20.services.HotelTextAssetR\n" +
	"textAssets\x12U\n" +
	"\fimage_assets\x18\x06 \x03(\v22.google.ads.googleads.v20.services.HotelImageAssetR\vimageAssets\x12q\n" +
	"\x06status\x18\a \x01(\x0e2Y.google.ads.googleads.v20.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatusR\x06status\"\x91\x01\n" +
	"\x0eHotelTextAsset\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\x12k\n" +
	"\x10asset_field_type\x18\x02 \x01(\x0e2A.google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldTypeR\x0eassetFieldType\"\x90\x01\n" +
	"\x0fHotelImageAsset\x12\x10\n" +
	"\x03uri\x18\x01 \x01(\tR\x03uri\x12k\n" +
	"\x10asset_field_type\x18\x02 \x01(\x0e2A.google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldTypeR\x0eassetFieldType2\xd9\x02\n" +
	"\x1cTravelAssetSuggestionService\x12\xf1\x01\n" +
	"\x13SuggestTravelAssets\x12=.google.ads.googleads.v20.services.SuggestTravelAssetsRequest\x1a>.google.ads.googleads.v20.services.SuggestTravelAssetsResponse\"[\xdaA\x1bcustomer_id,language_option\x82\xd3\xe4\x93\x027:\x01*\"2/v20/customers/{customer_id=*}:suggestTravelAssets\x1aE\xcaA\x18googleads.googleapis.com\xd2A'https://www.googleapis.com/auth/adwordsB\x8d\x02\n" +
	"%com.google.ads.googleads.v20.servicesB!TravelAssetSuggestionServiceProtoP\x01ZIgoogle.golang.org/genproto/googleapis/ads/googleads/v20/services;services\xa2\x02\x03GAA\xaa\x02!Google.Ads.GoogleAds.V20.Services\xca\x02!Google\\Ads\\GoogleAds\\V20\\Services\xea\x02%Google::Ads::GoogleAds::V20::Servicesb\x06proto3"

var (
	file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDescData
}

var file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_goTypes = []any{
	(*SuggestTravelAssetsRequest)(nil),                                   // 0: google.ads.googleads.v20.services.SuggestTravelAssetsRequest
	(*SuggestTravelAssetsResponse)(nil),                                  // 1: google.ads.googleads.v20.services.SuggestTravelAssetsResponse
	(*HotelAssetSuggestion)(nil),                                         // 2: google.ads.googleads.v20.services.HotelAssetSuggestion
	(*HotelTextAsset)(nil),                                               // 3: google.ads.googleads.v20.services.HotelTextAsset
	(*HotelImageAsset)(nil),                                              // 4: google.ads.googleads.v20.services.HotelImageAsset
	(enums.CallToActionTypeEnum_CallToActionType)(0),                     // 5: google.ads.googleads.v20.enums.CallToActionTypeEnum.CallToActionType
	(enums.HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus)(0), // 6: google.ads.googleads.v20.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatus
	(enums.AssetFieldTypeEnum_AssetFieldType)(0),                         // 7: google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
}
var file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v20.services.SuggestTravelAssetsResponse.hotel_asset_suggestions:type_name -> google.ads.googleads.v20.services.HotelAssetSuggestion
	5, // 1: google.ads.googleads.v20.services.HotelAssetSuggestion.call_to_action:type_name -> google.ads.googleads.v20.enums.CallToActionTypeEnum.CallToActionType
	3, // 2: google.ads.googleads.v20.services.HotelAssetSuggestion.text_assets:type_name -> google.ads.googleads.v20.services.HotelTextAsset
	4, // 3: google.ads.googleads.v20.services.HotelAssetSuggestion.image_assets:type_name -> google.ads.googleads.v20.services.HotelImageAsset
	6, // 4: google.ads.googleads.v20.services.HotelAssetSuggestion.status:type_name -> google.ads.googleads.v20.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatus
	7, // 5: google.ads.googleads.v20.services.HotelTextAsset.asset_field_type:type_name -> google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
	7, // 6: google.ads.googleads.v20.services.HotelImageAsset.asset_field_type:type_name -> google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
	0, // 7: google.ads.googleads.v20.services.TravelAssetSuggestionService.SuggestTravelAssets:input_type -> google.ads.googleads.v20.services.SuggestTravelAssetsRequest
	1, // 8: google.ads.googleads.v20.services.TravelAssetSuggestionService.SuggestTravelAssets:output_type -> google.ads.googleads.v20.services.SuggestTravelAssetsResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_init() }
func file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_init() {
	if File_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDesc), len(file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto = out.File
	file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_goTypes = nil
	file_google_ads_googleads_v20_services_travel_asset_suggestion_service_proto_depIdxs = nil
}
