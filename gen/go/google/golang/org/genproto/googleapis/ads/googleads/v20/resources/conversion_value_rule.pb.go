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
// source: google/ads/googleads/v20/resources/conversion_value_rule.proto

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

// A conversion value rule
type ConversionValueRule struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the conversion value rule.
	// Conversion value rule resource names have the form:
	//
	// `customers/{customer_id}/conversionValueRules/{conversion_value_rule_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the conversion value rule.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Action applied when the rule is triggered.
	Action *ConversionValueRule_ValueRuleAction `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	// Condition for Geo location that must be satisfied for the value rule to
	// apply.
	GeoLocationCondition *ConversionValueRule_ValueRuleGeoLocationCondition `protobuf:"bytes,4,opt,name=geo_location_condition,json=geoLocationCondition,proto3" json:"geo_location_condition,omitempty"`
	// Condition for device type that must be satisfied for the value rule to
	// apply.
	DeviceCondition *ConversionValueRule_ValueRuleDeviceCondition `protobuf:"bytes,5,opt,name=device_condition,json=deviceCondition,proto3" json:"device_condition,omitempty"`
	// Condition for audience that must be satisfied for the value rule to apply.
	AudienceCondition *ConversionValueRule_ValueRuleAudienceCondition `protobuf:"bytes,6,opt,name=audience_condition,json=audienceCondition,proto3" json:"audience_condition,omitempty"`
	// Condition for itinerary that must be satisfied for the value rule to apply.
	ItineraryCondition *ConversionValueRule_ValueRuleItineraryCondition `protobuf:"bytes,9,opt,name=itinerary_condition,json=itineraryCondition,proto3" json:"itinerary_condition,omitempty"`
	// Output only. The resource name of the conversion value rule's owner
	// customer. When the value rule is inherited from a manager customer,
	// owner_customer will be the resource name of the manager whereas the
	// customer in the resource_name will be of the requesting serving customer.
	// ** Read-only **
	OwnerCustomer string `protobuf:"bytes,7,opt,name=owner_customer,json=ownerCustomer,proto3" json:"owner_customer,omitempty"`
	// The status of the conversion value rule.
	Status        enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus `protobuf:"varint,8,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule) Reset() {
	*x = ConversionValueRule{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule) ProtoMessage() {}

func (x *ConversionValueRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule.ProtoReflect.Descriptor instead.
func (*ConversionValueRule) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0}
}

func (x *ConversionValueRule) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ConversionValueRule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConversionValueRule) GetAction() *ConversionValueRule_ValueRuleAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *ConversionValueRule) GetGeoLocationCondition() *ConversionValueRule_ValueRuleGeoLocationCondition {
	if x != nil {
		return x.GeoLocationCondition
	}
	return nil
}

func (x *ConversionValueRule) GetDeviceCondition() *ConversionValueRule_ValueRuleDeviceCondition {
	if x != nil {
		return x.DeviceCondition
	}
	return nil
}

func (x *ConversionValueRule) GetAudienceCondition() *ConversionValueRule_ValueRuleAudienceCondition {
	if x != nil {
		return x.AudienceCondition
	}
	return nil
}

func (x *ConversionValueRule) GetItineraryCondition() *ConversionValueRule_ValueRuleItineraryCondition {
	if x != nil {
		return x.ItineraryCondition
	}
	return nil
}

func (x *ConversionValueRule) GetOwnerCustomer() string {
	if x != nil {
		return x.OwnerCustomer
	}
	return ""
}

func (x *ConversionValueRule) GetStatus() enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus {
	if x != nil {
		return x.Status
	}
	return enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus(0)
}

// Action applied when rule is applied.
type ConversionValueRule_ValueRuleAction struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Specifies applied operation.
	Operation enums.ValueRuleOperationEnum_ValueRuleOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=google.ads.googleads.v20.enums.ValueRuleOperationEnum_ValueRuleOperation" json:"operation,omitempty"`
	// Specifies applied value.
	Value         float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleAction) Reset() {
	*x = ConversionValueRule_ValueRuleAction{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleAction) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleAction) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleAction.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleAction) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConversionValueRule_ValueRuleAction) GetOperation() enums.ValueRuleOperationEnum_ValueRuleOperation {
	if x != nil {
		return x.Operation
	}
	return enums.ValueRuleOperationEnum_ValueRuleOperation(0)
}

func (x *ConversionValueRule_ValueRuleAction) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Condition on Geo dimension.
type ConversionValueRule_ValueRuleGeoLocationCondition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Geo locations that advertisers want to exclude.
	ExcludedGeoTargetConstants []string `protobuf:"bytes,1,rep,name=excluded_geo_target_constants,json=excludedGeoTargetConstants,proto3" json:"excluded_geo_target_constants,omitempty"`
	// Excluded Geo location match type.
	ExcludedGeoMatchType enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType `protobuf:"varint,2,opt,name=excluded_geo_match_type,json=excludedGeoMatchType,proto3,enum=google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType" json:"excluded_geo_match_type,omitempty"`
	// Geo locations that advertisers want to include.
	GeoTargetConstants []string `protobuf:"bytes,3,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Included Geo location match type.
	GeoMatchType  enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType `protobuf:"varint,4,opt,name=geo_match_type,json=geoMatchType,proto3,enum=google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType" json:"geo_match_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) Reset() {
	*x = ConversionValueRule_ValueRuleGeoLocationCondition{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleGeoLocationCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleGeoLocationCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleGeoLocationCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetExcludedGeoTargetConstants() []string {
	if x != nil {
		return x.ExcludedGeoTargetConstants
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetExcludedGeoMatchType() enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType {
	if x != nil {
		return x.ExcludedGeoMatchType
	}
	return enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType(0)
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetGeoTargetConstants() []string {
	if x != nil {
		return x.GeoTargetConstants
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleGeoLocationCondition) GetGeoMatchType() enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType {
	if x != nil {
		return x.GeoMatchType
	}
	return enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType(0)
}

// Condition on Device dimension.
type ConversionValueRule_ValueRuleDeviceCondition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Value for device type condition.
	DeviceTypes   []enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType `protobuf:"varint,1,rep,packed,name=device_types,json=deviceTypes,proto3,enum=google.ads.googleads.v20.enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType" json:"device_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) Reset() {
	*x = ConversionValueRule_ValueRuleDeviceCondition{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleDeviceCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleDeviceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleDeviceCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleDeviceCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ConversionValueRule_ValueRuleDeviceCondition) GetDeviceTypes() []enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType {
	if x != nil {
		return x.DeviceTypes
	}
	return nil
}

// Condition on Audience dimension.
type ConversionValueRule_ValueRuleAudienceCondition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// User Lists.
	UserLists []string `protobuf:"bytes,1,rep,name=user_lists,json=userLists,proto3" json:"user_lists,omitempty"`
	// User Interests.
	UserInterests []string `protobuf:"bytes,2,rep,name=user_interests,json=userInterests,proto3" json:"user_interests,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) Reset() {
	*x = ConversionValueRule_ValueRuleAudienceCondition{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleAudienceCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleAudienceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleAudienceCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleAudienceCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) GetUserLists() []string {
	if x != nil {
		return x.UserLists
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleAudienceCondition) GetUserInterests() []string {
	if x != nil {
		return x.UserInterests
	}
	return nil
}

// Condition on Itinerary dimension.
type ConversionValueRule_ValueRuleItineraryCondition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Range for the number of days between the date of the booking and the
	// start of the itinerary.
	AdvanceBookingWindow *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow `protobuf:"bytes,1,opt,name=advance_booking_window,json=advanceBookingWindow,proto3" json:"advance_booking_window,omitempty"`
	// Range for the itinerary length in number of nights.
	TravelLength *ConversionValueRule_ValueRuleItineraryTravelLength `protobuf:"bytes,2,opt,name=travel_length,json=travelLength,proto3" json:"travel_length,omitempty"`
	// The days of the week on which this itinerary's travel can start.
	TravelStartDay *ConversionValueRule_ValueRuleItineraryTravelStartDay `protobuf:"bytes,3,opt,name=travel_start_day,json=travelStartDay,proto3" json:"travel_start_day,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleItineraryCondition) Reset() {
	*x = ConversionValueRule_ValueRuleItineraryCondition{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleItineraryCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleItineraryCondition) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleItineraryCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleItineraryCondition.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleItineraryCondition) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 4}
}

func (x *ConversionValueRule_ValueRuleItineraryCondition) GetAdvanceBookingWindow() *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow {
	if x != nil {
		return x.AdvanceBookingWindow
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleItineraryCondition) GetTravelLength() *ConversionValueRule_ValueRuleItineraryTravelLength {
	if x != nil {
		return x.TravelLength
	}
	return nil
}

func (x *ConversionValueRule_ValueRuleItineraryCondition) GetTravelStartDay() *ConversionValueRule_ValueRuleItineraryTravelStartDay {
	if x != nil {
		return x.TravelStartDay
	}
	return nil
}

// Range for the number of days between the date of the booking and the
// start of the itinerary.
type ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Minimum number of days between the date of the booking the start date.
	MinDays *int32 `protobuf:"varint,3,opt,name=min_days,json=minDays,proto3,oneof" json:"min_days,omitempty"`
	// Maximum number of days between the date of the booking the start date.
	MaxDays       *int32 `protobuf:"varint,4,opt,name=max_days,json=maxDays,proto3,oneof" json:"max_days,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) Reset() {
	*x = ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 5}
}

func (x *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) GetMinDays() int32 {
	if x != nil && x.MinDays != nil {
		return *x.MinDays
	}
	return 0
}

func (x *ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow) GetMaxDays() int32 {
	if x != nil && x.MaxDays != nil {
		return *x.MaxDays
	}
	return 0
}

// Range for the itinerary length in number of nights.
type ConversionValueRule_ValueRuleItineraryTravelLength struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Minimum number of nights between the start date and the end date.
	MinNights int32 `protobuf:"varint,1,opt,name=min_nights,json=minNights,proto3" json:"min_nights,omitempty"`
	// Maximum number of days between the start date and the end date.
	MaxNights     int32 `protobuf:"varint,2,opt,name=max_nights,json=maxNights,proto3" json:"max_nights,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleItineraryTravelLength) Reset() {
	*x = ConversionValueRule_ValueRuleItineraryTravelLength{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleItineraryTravelLength) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleItineraryTravelLength) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleItineraryTravelLength) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleItineraryTravelLength.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleItineraryTravelLength) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 6}
}

func (x *ConversionValueRule_ValueRuleItineraryTravelLength) GetMinNights() int32 {
	if x != nil {
		return x.MinNights
	}
	return 0
}

func (x *ConversionValueRule_ValueRuleItineraryTravelLength) GetMaxNights() int32 {
	if x != nil {
		return x.MaxNights
	}
	return 0
}

// The days of the week on which an itinerary's travel can start.
type ConversionValueRule_ValueRuleItineraryTravelStartDay struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The travel can start on Monday.
	Monday bool `protobuf:"varint,1,opt,name=monday,proto3" json:"monday,omitempty"`
	// The travel can start on Tuesday.
	Tuesday bool `protobuf:"varint,2,opt,name=tuesday,proto3" json:"tuesday,omitempty"`
	// The travel can start on Wednesday.
	Wednesday bool `protobuf:"varint,3,opt,name=wednesday,proto3" json:"wednesday,omitempty"`
	// The travel can start on Thursday.
	Thursday bool `protobuf:"varint,4,opt,name=thursday,proto3" json:"thursday,omitempty"`
	// The travel can start on Friday.
	Friday bool `protobuf:"varint,5,opt,name=friday,proto3" json:"friday,omitempty"`
	// The travel can start on Saturday.
	Saturday bool `protobuf:"varint,6,opt,name=saturday,proto3" json:"saturday,omitempty"`
	// The travel can start on Sunday.
	Sunday        bool `protobuf:"varint,7,opt,name=sunday,proto3" json:"sunday,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) Reset() {
	*x = ConversionValueRule_ValueRuleItineraryTravelStartDay{}
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRule_ValueRuleItineraryTravelStartDay) ProtoMessage() {}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRule_ValueRuleItineraryTravelStartDay.ProtoReflect.Descriptor instead.
func (*ConversionValueRule_ValueRuleItineraryTravelStartDay) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP(), []int{0, 7}
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetMonday() bool {
	if x != nil {
		return x.Monday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetTuesday() bool {
	if x != nil {
		return x.Tuesday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetWednesday() bool {
	if x != nil {
		return x.Wednesday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetThursday() bool {
	if x != nil {
		return x.Thursday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetFriday() bool {
	if x != nil {
		return x.Friday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetSaturday() bool {
	if x != nil {
		return x.Saturday
	}
	return false
}

func (x *ConversionValueRule_ValueRuleItineraryTravelStartDay) GetSunday() bool {
	if x != nil {
		return x.Sunday
	}
	return false
}

var File_google_ads_googleads_v20_resources_conversion_value_rule_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/resources/conversion_value_rule.proto\x12\"google.ads.googleads.v20.resources\x1aAgoogle/ads/googleads/v20/enums/conversion_value_rule_status.proto\x1a;google/ads/googleads/v20/enums/value_rule_device_type.proto\x1aGgoogle/ads/googleads/v20/enums/value_rule_geo_location_match_type.proto\x1a9google/ads/googleads/v20/enums/value_rule_operation.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xa9\x17\n" +
	"\x13ConversionValueRule\x12Y\n" +
	"\rresource_name\x18\x01 \x01(\tB4\xe0A\x05\xfaA.\n" +
	",googleads.googleapis.com/ConversionValueRuleR\fresourceName\x12\x13\n" +
	"\x02id\x18\x02 \x01(\x03B\x03\xe0A\x03R\x02id\x12_\n" +
	"\x06action\x18\x03 \x01(\v2G.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleActionR\x06action\x12\x8b\x01\n" +
	"\x16geo_location_condition\x18\x04 \x01(\v2U.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleGeoLocationConditionR\x14geoLocationCondition\x12{\n" +
	"\x10device_condition\x18\x05 \x01(\v2P.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleDeviceConditionR\x0fdeviceCondition\x12\x81\x01\n" +
	"\x12audience_condition\x18\x06 \x01(\v2R.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAudienceConditionR\x11audienceCondition\x12\x84\x01\n" +
	"\x13itinerary_condition\x18\t \x01(\v2S.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryConditionR\x12itineraryCondition\x12P\n" +
	"\x0eowner_customer\x18\a \x01(\tB)\xe0A\x03\xfaA#\n" +
	"!googleads.googleapis.com/CustomerR\rownerCustomer\x12o\n" +
	"\x06status\x18\b \x01(\x0e2W.google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatusR\x06status\x1a\x90\x01\n" +
	"\x0fValueRuleAction\x12g\n" +
	"\toperation\x18\x01 \x01(\x0e2I.google.ads.googleads.v20.enums.ValueRuleOperationEnum.ValueRuleOperationR\toperation\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value\x1a\x97\x04\n" +
	"\x1dValueRuleGeoLocationCondition\x12r\n" +
	"\x1dexcluded_geo_target_constants\x18\x01 \x03(\tB/\xfaA,\n" +
	"*googleads.googleapis.com/GeoTargetConstantR\x1aexcludedGeoTargetConstants\x12\x96\x01\n" +
	"\x17excluded_geo_match_type\x18\x02 \x01(\x0e2_.google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchTypeR\x14excludedGeoMatchType\x12a\n" +
	"\x14geo_target_constants\x18\x03 \x03(\tB/\xfaA,\n" +
	"*googleads.googleapis.com/GeoTargetConstantR\x12geoTargetConstants\x12\x85\x01\n" +
	"\x0egeo_match_type\x18\x04 \x01(\x0e2_.google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchTypeR\fgeoMatchType\x1a\x8a\x01\n" +
	"\x18ValueRuleDeviceCondition\x12n\n" +
	"\fdevice_types\x18\x01 \x03(\x0e2K.google.ads.googleads.v20.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceTypeR\vdeviceTypes\x1a\xb6\x01\n" +
	"\x1aValueRuleAudienceCondition\x12E\n" +
	"\n" +
	"user_lists\x18\x01 \x03(\tB&\xfaA#\n" +
	"!googleads.googleapis.com/UserListR\tuserLists\x12Q\n" +
	"\x0euser_interests\x18\x02 \x03(\tB*\xfaA'\n" +
	"%googleads.googleapis.com/UserInterestR\ruserInterests\x1a\xb6\x03\n" +
	"\x1bValueRuleItineraryCondition\x12\x94\x01\n" +
	"\x16advance_booking_window\x18\x01 \x01(\v2^.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryAdvanceBookingWindowR\x14advanceBookingWindow\x12{\n" +
	"\rtravel_length\x18\x02 \x01(\v2V.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelLengthR\ftravelLength\x12\x82\x01\n" +
	"\x10travel_start_day\x18\x03 \x01(\v2X.google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelStartDayR\x0etravelStartDay\x1a\x82\x01\n" +
	"&ValueRuleItineraryAdvanceBookingWindow\x12\x1e\n" +
	"\bmin_days\x18\x03 \x01(\x05H\x00R\aminDays\x88\x01\x01\x12\x1e\n" +
	"\bmax_days\x18\x04 \x01(\x05H\x01R\amaxDays\x88\x01\x01B\v\n" +
	"\t_min_daysB\v\n" +
	"\t_max_days\x1a^\n" +
	"\x1eValueRuleItineraryTravelLength\x12\x1d\n" +
	"\n" +
	"min_nights\x18\x01 \x01(\x05R\tminNights\x12\x1d\n" +
	"\n" +
	"max_nights\x18\x02 \x01(\x05R\tmaxNights\x1a\xda\x01\n" +
	" ValueRuleItineraryTravelStartDay\x12\x16\n" +
	"\x06monday\x18\x01 \x01(\bR\x06monday\x12\x18\n" +
	"\atuesday\x18\x02 \x01(\bR\atuesday\x12\x1c\n" +
	"\twednesday\x18\x03 \x01(\bR\twednesday\x12\x1a\n" +
	"\bthursday\x18\x04 \x01(\bR\bthursday\x12\x16\n" +
	"\x06friday\x18\x05 \x01(\bR\x06friday\x12\x1a\n" +
	"\bsaturday\x18\x06 \x01(\bR\bsaturday\x12\x16\n" +
	"\x06sunday\x18\a \x01(\bR\x06sunday:z\xeaAw\n" +
	",googleads.googleapis.com/ConversionValueRule\x12Gcustomers/{customer_id}/conversionValueRules/{conversion_value_rule_id}B\x8a\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x18ConversionValueRuleProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDesc), len(file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_ads_googleads_v20_resources_conversion_value_rule_proto_goTypes = []any{
	(*ConversionValueRule)(nil),                                                // 0: google.ads.googleads.v20.resources.ConversionValueRule
	(*ConversionValueRule_ValueRuleAction)(nil),                                // 1: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAction
	(*ConversionValueRule_ValueRuleGeoLocationCondition)(nil),                  // 2: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleGeoLocationCondition
	(*ConversionValueRule_ValueRuleDeviceCondition)(nil),                       // 3: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleDeviceCondition
	(*ConversionValueRule_ValueRuleAudienceCondition)(nil),                     // 4: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAudienceCondition
	(*ConversionValueRule_ValueRuleItineraryCondition)(nil),                    // 5: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryCondition
	(*ConversionValueRule_ValueRuleItineraryAdvanceBookingWindow)(nil),         // 6: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryAdvanceBookingWindow
	(*ConversionValueRule_ValueRuleItineraryTravelLength)(nil),                 // 7: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelLength
	(*ConversionValueRule_ValueRuleItineraryTravelStartDay)(nil),               // 8: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelStartDay
	(enums.ConversionValueRuleStatusEnum_ConversionValueRuleStatus)(0),         // 9: google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatus
	(enums.ValueRuleOperationEnum_ValueRuleOperation)(0),                       // 10: google.ads.googleads.v20.enums.ValueRuleOperationEnum.ValueRuleOperation
	(enums.ValueRuleGeoLocationMatchTypeEnum_ValueRuleGeoLocationMatchType)(0), // 11: google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	(enums.ValueRuleDeviceTypeEnum_ValueRuleDeviceType)(0),                     // 12: google.ads.googleads.v20.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceType
}
var file_google_ads_googleads_v20_resources_conversion_value_rule_proto_depIdxs = []int32{
	1,  // 0: google.ads.googleads.v20.resources.ConversionValueRule.action:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAction
	2,  // 1: google.ads.googleads.v20.resources.ConversionValueRule.geo_location_condition:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleGeoLocationCondition
	3,  // 2: google.ads.googleads.v20.resources.ConversionValueRule.device_condition:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleDeviceCondition
	4,  // 3: google.ads.googleads.v20.resources.ConversionValueRule.audience_condition:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAudienceCondition
	5,  // 4: google.ads.googleads.v20.resources.ConversionValueRule.itinerary_condition:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryCondition
	9,  // 5: google.ads.googleads.v20.resources.ConversionValueRule.status:type_name -> google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatus
	10, // 6: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleAction.operation:type_name -> google.ads.googleads.v20.enums.ValueRuleOperationEnum.ValueRuleOperation
	11, // 7: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleGeoLocationCondition.excluded_geo_match_type:type_name -> google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	11, // 8: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleGeoLocationCondition.geo_match_type:type_name -> google.ads.googleads.v20.enums.ValueRuleGeoLocationMatchTypeEnum.ValueRuleGeoLocationMatchType
	12, // 9: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleDeviceCondition.device_types:type_name -> google.ads.googleads.v20.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceType
	6,  // 10: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryCondition.advance_booking_window:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryAdvanceBookingWindow
	7,  // 11: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryCondition.travel_length:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelLength
	8,  // 12: google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryCondition.travel_start_day:type_name -> google.ads.googleads.v20.resources.ConversionValueRule.ValueRuleItineraryTravelStartDay
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_conversion_value_rule_proto_init() }
func file_google_ads_googleads_v20_resources_conversion_value_rule_proto_init() {
	if File_google_ads_googleads_v20_resources_conversion_value_rule_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDesc), len(file_google_ads_googleads_v20_resources_conversion_value_rule_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_conversion_value_rule_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_conversion_value_rule_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_conversion_value_rule_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_conversion_value_rule_proto = out.File
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_conversion_value_rule_proto_depIdxs = nil
}
