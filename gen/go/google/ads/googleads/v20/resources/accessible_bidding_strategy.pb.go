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
// source: google/ads/googleads/v20/resources/accessible_bidding_strategy.proto

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

// Represents a view of BiddingStrategies owned by and shared with the customer.
//
// In contrast to BiddingStrategy, this resource includes strategies owned by
// managers of the customer and shared with this customer - in addition to
// strategies owned by this customer. This resource does not provide metrics and
// only exposes a limited subset of the BiddingStrategy attributes.
type AccessibleBiddingStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the accessible bidding strategy.
	// AccessibleBiddingStrategy resource names have the form:
	//
	// `customers/{customer_id}/accessibleBiddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the bidding strategy.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The name of the bidding strategy.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The type of the bidding strategy.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v20.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// Output only. The ID of the Customer which owns the bidding strategy.
	OwnerCustomerId int64 `protobuf:"varint,5,opt,name=owner_customer_id,json=ownerCustomerId,proto3" json:"owner_customer_id,omitempty"`
	// Output only. descriptive_name of the Customer which owns the bidding
	// strategy.
	OwnerDescriptiveName string `protobuf:"bytes,6,opt,name=owner_descriptive_name,json=ownerDescriptiveName,proto3" json:"owner_descriptive_name,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//
	//	*AccessibleBiddingStrategy_MaximizeConversionValue_
	//	*AccessibleBiddingStrategy_MaximizeConversions_
	//	*AccessibleBiddingStrategy_TargetCpa_
	//	*AccessibleBiddingStrategy_TargetImpressionShare_
	//	*AccessibleBiddingStrategy_TargetRoas_
	//	*AccessibleBiddingStrategy_TargetSpend_
	Scheme        isAccessibleBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy) Reset() {
	*x = AccessibleBiddingStrategy{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy) ProtoMessage() {}

func (x *AccessibleBiddingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *AccessibleBiddingStrategy) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AccessibleBiddingStrategy) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccessibleBiddingStrategy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessibleBiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if x != nil {
		return x.Type
	}
	return enums.BiddingStrategyTypeEnum_BiddingStrategyType(0)
}

func (x *AccessibleBiddingStrategy) GetOwnerCustomerId() int64 {
	if x != nil {
		return x.OwnerCustomerId
	}
	return 0
}

func (x *AccessibleBiddingStrategy) GetOwnerDescriptiveName() string {
	if x != nil {
		return x.OwnerDescriptiveName
	}
	return ""
}

func (x *AccessibleBiddingStrategy) GetScheme() isAccessibleBiddingStrategy_Scheme {
	if x != nil {
		return x.Scheme
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetMaximizeConversionValue() *AccessibleBiddingStrategy_MaximizeConversionValue {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_MaximizeConversionValue_); ok {
			return x.MaximizeConversionValue
		}
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetMaximizeConversions() *AccessibleBiddingStrategy_MaximizeConversions {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_MaximizeConversions_); ok {
			return x.MaximizeConversions
		}
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetTargetCpa() *AccessibleBiddingStrategy_TargetCpa {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_TargetCpa_); ok {
			return x.TargetCpa
		}
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetTargetImpressionShare() *AccessibleBiddingStrategy_TargetImpressionShare {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_TargetImpressionShare_); ok {
			return x.TargetImpressionShare
		}
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetTargetRoas() *AccessibleBiddingStrategy_TargetRoas {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_TargetRoas_); ok {
			return x.TargetRoas
		}
	}
	return nil
}

func (x *AccessibleBiddingStrategy) GetTargetSpend() *AccessibleBiddingStrategy_TargetSpend {
	if x != nil {
		if x, ok := x.Scheme.(*AccessibleBiddingStrategy_TargetSpend_); ok {
			return x.TargetSpend
		}
	}
	return nil
}

type isAccessibleBiddingStrategy_Scheme interface {
	isAccessibleBiddingStrategy_Scheme()
}

type AccessibleBiddingStrategy_MaximizeConversionValue_ struct {
	// Output only. An automated bidding strategy to help get the most
	// conversion value for your campaigns while spending your budget.
	MaximizeConversionValue *AccessibleBiddingStrategy_MaximizeConversionValue `protobuf:"bytes,7,opt,name=maximize_conversion_value,json=maximizeConversionValue,proto3,oneof"`
}

type AccessibleBiddingStrategy_MaximizeConversions_ struct {
	// Output only. An automated bidding strategy to help get the most
	// conversions for your campaigns while spending your budget.
	MaximizeConversions *AccessibleBiddingStrategy_MaximizeConversions `protobuf:"bytes,8,opt,name=maximize_conversions,json=maximizeConversions,proto3,oneof"`
}

type AccessibleBiddingStrategy_TargetCpa_ struct {
	// Output only. A bidding strategy that sets bids to help get as many
	// conversions as possible at the target cost-per-acquisition (CPA) you set.
	TargetCpa *AccessibleBiddingStrategy_TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type AccessibleBiddingStrategy_TargetImpressionShare_ struct {
	// Output only. A bidding strategy that automatically optimizes towards a
	// chosen percentage of impressions.
	TargetImpressionShare *AccessibleBiddingStrategy_TargetImpressionShare `protobuf:"bytes,10,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type AccessibleBiddingStrategy_TargetRoas_ struct {
	// Output only. A bidding strategy that helps you maximize revenue while
	// averaging a specific target Return On Ad Spend (ROAS).
	TargetRoas *AccessibleBiddingStrategy_TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type AccessibleBiddingStrategy_TargetSpend_ struct {
	// Output only. A bid strategy that sets your bids to help get as many
	// clicks as possible within your budget.
	TargetSpend *AccessibleBiddingStrategy_TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*AccessibleBiddingStrategy_MaximizeConversionValue_) isAccessibleBiddingStrategy_Scheme() {}

func (*AccessibleBiddingStrategy_MaximizeConversions_) isAccessibleBiddingStrategy_Scheme() {}

func (*AccessibleBiddingStrategy_TargetCpa_) isAccessibleBiddingStrategy_Scheme() {}

func (*AccessibleBiddingStrategy_TargetImpressionShare_) isAccessibleBiddingStrategy_Scheme() {}

func (*AccessibleBiddingStrategy_TargetRoas_) isAccessibleBiddingStrategy_Scheme() {}

func (*AccessibleBiddingStrategy_TargetSpend_) isAccessibleBiddingStrategy_Scheme() {}

// An automated bidding strategy to help get the most conversion value for
// your campaigns while spending your budget.
type AccessibleBiddingStrategy_MaximizeConversionValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The target return on ad spend (ROAS) option. If set, the bid
	// strategy will maximize revenue while averaging the target return on ad
	// spend. If the target ROAS is high, the bid strategy may not be able to
	// spend the full budget. If the target ROAS is not set, the bid strategy
	// will aim to achieve the highest possible ROAS for the budget.
	TargetRoas    float64 `protobuf:"fixed64,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_MaximizeConversionValue) Reset() {
	*x = AccessibleBiddingStrategy_MaximizeConversionValue{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_MaximizeConversionValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_MaximizeConversionValue) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_MaximizeConversionValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_MaximizeConversionValue.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_MaximizeConversionValue) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AccessibleBiddingStrategy_MaximizeConversionValue) GetTargetRoas() float64 {
	if x != nil {
		return x.TargetRoas
	}
	return 0
}

// An automated bidding strategy to help get the most conversions for your
// campaigns while spending your budget.
type AccessibleBiddingStrategy_MaximizeConversions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The target cost per acquisition (CPA) option. This is the
	// average amount that you would like to spend per acquisition.
	TargetCpaMicros int64 `protobuf:"varint,2,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_MaximizeConversions) Reset() {
	*x = AccessibleBiddingStrategy_MaximizeConversions{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_MaximizeConversions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_MaximizeConversions) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_MaximizeConversions) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_MaximizeConversions.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_MaximizeConversions) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AccessibleBiddingStrategy_MaximizeConversions) GetTargetCpaMicros() int64 {
	if x != nil {
		return x.TargetCpaMicros
	}
	return 0
}

// An automated bid strategy that sets bids to help get as many conversions as
// possible at the target cost-per-acquisition (CPA) you set.
type AccessibleBiddingStrategy_TargetCpa struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. Average CPA target.
	// This target should be greater than or equal to minimum billable unit
	// based on the currency for the account.
	TargetCpaMicros *int64 `protobuf:"varint,1,opt,name=target_cpa_micros,json=targetCpaMicros,proto3,oneof" json:"target_cpa_micros,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_TargetCpa) Reset() {
	*x = AccessibleBiddingStrategy_TargetCpa{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_TargetCpa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_TargetCpa) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_TargetCpa) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_TargetCpa.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_TargetCpa) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AccessibleBiddingStrategy_TargetCpa) GetTargetCpaMicros() int64 {
	if x != nil && x.TargetCpaMicros != nil {
		return *x.TargetCpaMicros
	}
	return 0
}

// An automated bidding strategy that sets bids so that a certain percentage
// of search ads are shown at the top of the first page (or other targeted
// location).
type AccessibleBiddingStrategy_TargetImpressionShare struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The targeted location on the search results page.
	Location enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation `protobuf:"varint,1,opt,name=location,proto3,enum=google.ads.googleads.v20.enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation" json:"location,omitempty"`
	// The chosen fraction of ads to be shown in the targeted location in
	// micros. For example, 1% equals 10,000.
	LocationFractionMicros *int64 `protobuf:"varint,2,opt,name=location_fraction_micros,json=locationFractionMicros,proto3,oneof" json:"location_fraction_micros,omitempty"`
	// Output only. The highest CPC bid the automated bidding system is
	// permitted to specify. This is a required field entered by the advertiser
	// that sets the ceiling and specified in local micros.
	CpcBidCeilingMicros *int64 `protobuf:"varint,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3,oneof" json:"cpc_bid_ceiling_micros,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) Reset() {
	*x = AccessibleBiddingStrategy_TargetImpressionShare{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_TargetImpressionShare) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_TargetImpressionShare.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_TargetImpressionShare) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 3}
}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) GetLocation() enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation {
	if x != nil {
		return x.Location
	}
	return enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation(0)
}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) GetLocationFractionMicros() int64 {
	if x != nil && x.LocationFractionMicros != nil {
		return *x.LocationFractionMicros
	}
	return 0
}

func (x *AccessibleBiddingStrategy_TargetImpressionShare) GetCpcBidCeilingMicros() int64 {
	if x != nil && x.CpcBidCeilingMicros != nil {
		return *x.CpcBidCeilingMicros
	}
	return 0
}

// An automated bidding strategy that helps you maximize revenue while
// averaging a specific target return on ad spend (ROAS).
type AccessibleBiddingStrategy_TargetRoas struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The chosen revenue (based on conversion data) per unit of
	// spend.
	TargetRoas    *float64 `protobuf:"fixed64,1,opt,name=target_roas,json=targetRoas,proto3,oneof" json:"target_roas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_TargetRoas) Reset() {
	*x = AccessibleBiddingStrategy_TargetRoas{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_TargetRoas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_TargetRoas) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_TargetRoas) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_TargetRoas.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_TargetRoas) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 4}
}

func (x *AccessibleBiddingStrategy_TargetRoas) GetTargetRoas() float64 {
	if x != nil && x.TargetRoas != nil {
		return *x.TargetRoas
	}
	return 0
}

// An automated bid strategy that sets your bids to help get as many clicks
// as possible within your budget.
type AccessibleBiddingStrategy_TargetSpend struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The spend target under which to maximize clicks.
	// A TargetSpend bidder will attempt to spend the smaller of this value
	// or the natural throttling spend amount.
	// If not specified, the budget is used as the spend target.
	// This field is deprecated and should no longer be used. See
	// https://ads-developers.googleblog.com/2020/05/reminder-about-sunset-creation-of.html
	// for details.
	//
	// Deprecated: Marked as deprecated in google/ads/googleads/v20/resources/accessible_bidding_strategy.proto.
	TargetSpendMicros *int64 `protobuf:"varint,1,opt,name=target_spend_micros,json=targetSpendMicros,proto3,oneof" json:"target_spend_micros,omitempty"`
	// Output only. Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *int64 `protobuf:"varint,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3,oneof" json:"cpc_bid_ceiling_micros,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AccessibleBiddingStrategy_TargetSpend) Reset() {
	*x = AccessibleBiddingStrategy_TargetSpend{}
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessibleBiddingStrategy_TargetSpend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessibleBiddingStrategy_TargetSpend) ProtoMessage() {}

func (x *AccessibleBiddingStrategy_TargetSpend) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessibleBiddingStrategy_TargetSpend.ProtoReflect.Descriptor instead.
func (*AccessibleBiddingStrategy_TargetSpend) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP(), []int{0, 5}
}

// Deprecated: Marked as deprecated in google/ads/googleads/v20/resources/accessible_bidding_strategy.proto.
func (x *AccessibleBiddingStrategy_TargetSpend) GetTargetSpendMicros() int64 {
	if x != nil && x.TargetSpendMicros != nil {
		return *x.TargetSpendMicros
	}
	return 0
}

func (x *AccessibleBiddingStrategy_TargetSpend) GetCpcBidCeilingMicros() int64 {
	if x != nil && x.CpcBidCeilingMicros != nil {
		return *x.CpcBidCeilingMicros
	}
	return 0
}

var File_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDesc = "" +
	"\n" +
	"Dgoogle/ads/googleads/v20/resources/accessible_bidding_strategy.proto\x12\"google.ads.googleads.v20.resources\x1a:google/ads/googleads/v20/enums/bidding_strategy_type.proto\x1aEgoogle/ads/googleads/v20/enums/target_impression_share_location.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xdb\x10\n" +
	"\x19AccessibleBiddingStrategy\x12_\n" +
	"\rresource_name\x18\x01 \x01(\tB:\xe0A\x03\xfaA4\n" +
	"2googleads.googleapis.com/AccessibleBiddingStrategyR\fresourceName\x12\x13\n" +
	"\x02id\x18\x02 \x01(\x03B\x03\xe0A\x03R\x02id\x12\x17\n" +
	"\x04name\x18\x03 \x01(\tB\x03\xe0A\x03R\x04name\x12d\n" +
	"\x04type\x18\x04 \x01(\x0e2K.google.ads.googleads.v20.enums.BiddingStrategyTypeEnum.BiddingStrategyTypeB\x03\xe0A\x03R\x04type\x12/\n" +
	"\x11owner_customer_id\x18\x05 \x01(\x03B\x03\xe0A\x03R\x0fownerCustomerId\x129\n" +
	"\x16owner_descriptive_name\x18\x06 \x01(\tB\x03\xe0A\x03R\x14ownerDescriptiveName\x12\x98\x01\n" +
	"\x19maximize_conversion_value\x18\a \x01(\v2U.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversionValueB\x03\xe0A\x03H\x00R\x17maximizeConversionValue\x12\x8b\x01\n" +
	"\x14maximize_conversions\x18\b \x01(\v2Q.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversionsB\x03\xe0A\x03H\x00R\x13maximizeConversions\x12m\n" +
	"\n" +
	"target_cpa\x18\t \x01(\v2G.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetCpaB\x03\xe0A\x03H\x00R\ttargetCpa\x12\x92\x01\n" +
	"\x17target_impression_share\x18\n" +
	" \x01(\v2S.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetImpressionShareB\x03\xe0A\x03H\x00R\x15targetImpressionShare\x12p\n" +
	"\vtarget_roas\x18\v \x01(\v2H.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetRoasB\x03\xe0A\x03H\x00R\n" +
	"targetRoas\x12s\n" +
	"\ftarget_spend\x18\f \x01(\v2I.google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetSpendB\x03\xe0A\x03H\x00R\vtargetSpend\x1a?\n" +
	"\x17MaximizeConversionValue\x12$\n" +
	"\vtarget_roas\x18\x01 \x01(\x01B\x03\xe0A\x03R\n" +
	"targetRoas\x1aF\n" +
	"\x13MaximizeConversions\x12/\n" +
	"\x11target_cpa_micros\x18\x02 \x01(\x03B\x03\xe0A\x03R\x0ftargetCpaMicros\x1aW\n" +
	"\tTargetCpa\x124\n" +
	"\x11target_cpa_micros\x18\x01 \x01(\x03B\x03\xe0A\x03H\x00R\x0ftargetCpaMicros\x88\x01\x01B\x14\n" +
	"\x12_target_cpa_micros\x1a\xd0\x02\n" +
	"\x15TargetImpressionShare\x12\x80\x01\n" +
	"\blocation\x18\x01 \x01(\x0e2_.google.ads.googleads.v20.enums.TargetImpressionShareLocationEnum.TargetImpressionShareLocationB\x03\xe0A\x03R\blocation\x12=\n" +
	"\x18location_fraction_micros\x18\x02 \x01(\x03H\x00R\x16locationFractionMicros\x88\x01\x01\x12=\n" +
	"\x16cpc_bid_ceiling_micros\x18\x03 \x01(\x03B\x03\xe0A\x03H\x01R\x13cpcBidCeilingMicros\x88\x01\x01B\x1b\n" +
	"\x19_location_fraction_microsB\x19\n" +
	"\x17_cpc_bid_ceiling_micros\x1aG\n" +
	"\n" +
	"TargetRoas\x12)\n" +
	"\vtarget_roas\x18\x01 \x01(\x01B\x03\xe0A\x03H\x00R\n" +
	"targetRoas\x88\x01\x01B\x0e\n" +
	"\f_target_roas\x1a\xbb\x01\n" +
	"\vTargetSpend\x12:\n" +
	"\x13target_spend_micros\x18\x01 \x01(\x03B\x05\xe0A\x03\x18\x01H\x00R\x11targetSpendMicros\x88\x01\x01\x12=\n" +
	"\x16cpc_bid_ceiling_micros\x18\x02 \x01(\x03B\x03\xe0A\x03H\x01R\x13cpcBidCeilingMicros\x88\x01\x01B\x16\n" +
	"\x14_target_spend_microsB\x19\n" +
	"\x17_cpc_bid_ceiling_micros:\x82\x01\xeaA\x7f\n" +
	"2googleads.googleapis.com/AccessibleBiddingStrategy\x12Icustomers/{customer_id}/accessibleBiddingStrategies/{bidding_strategy_id}B\b\n" +
	"\x06schemeB\x90\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x1eAccessibleBiddingStrategyProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDesc), len(file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_goTypes = []any{
	(*AccessibleBiddingStrategy)(nil),                                          // 0: google.ads.googleads.v20.resources.AccessibleBiddingStrategy
	(*AccessibleBiddingStrategy_MaximizeConversionValue)(nil),                  // 1: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversionValue
	(*AccessibleBiddingStrategy_MaximizeConversions)(nil),                      // 2: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversions
	(*AccessibleBiddingStrategy_TargetCpa)(nil),                                // 3: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetCpa
	(*AccessibleBiddingStrategy_TargetImpressionShare)(nil),                    // 4: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetImpressionShare
	(*AccessibleBiddingStrategy_TargetRoas)(nil),                               // 5: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetRoas
	(*AccessibleBiddingStrategy_TargetSpend)(nil),                              // 6: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetSpend
	(enums.BiddingStrategyTypeEnum_BiddingStrategyType)(0),                     // 7: google.ads.googleads.v20.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	(enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation)(0), // 8: google.ads.googleads.v20.enums.TargetImpressionShareLocationEnum.TargetImpressionShareLocation
}
var file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_depIdxs = []int32{
	7, // 0: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.type:type_name -> google.ads.googleads.v20.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	1, // 1: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.maximize_conversion_value:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversionValue
	2, // 2: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.maximize_conversions:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.MaximizeConversions
	3, // 3: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.target_cpa:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetCpa
	4, // 4: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.target_impression_share:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetImpressionShare
	5, // 5: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.target_roas:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetRoas
	6, // 6: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.target_spend:type_name -> google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetSpend
	8, // 7: google.ads.googleads.v20.resources.AccessibleBiddingStrategy.TargetImpressionShare.location:type_name -> google.ads.googleads.v20.enums.TargetImpressionShareLocationEnum.TargetImpressionShareLocation
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_init() }
func file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_init() {
	if File_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[0].OneofWrappers = []any{
		(*AccessibleBiddingStrategy_MaximizeConversionValue_)(nil),
		(*AccessibleBiddingStrategy_MaximizeConversions_)(nil),
		(*AccessibleBiddingStrategy_TargetCpa_)(nil),
		(*AccessibleBiddingStrategy_TargetImpressionShare_)(nil),
		(*AccessibleBiddingStrategy_TargetRoas_)(nil),
		(*AccessibleBiddingStrategy_TargetSpend_)(nil),
	}
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[3].OneofWrappers = []any{}
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[4].OneofWrappers = []any{}
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[5].OneofWrappers = []any{}
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDesc), len(file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto = out.File
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_accessible_bidding_strategy_proto_depIdxs = nil
}
