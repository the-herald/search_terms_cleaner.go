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
// source: google/ads/googleads/v20/resources/user_list.proto

package resources

import (
	common "google.golang.org/genproto/googleapis/ads/googleads/v20/common"
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

// A user list. This is a list of users a customer may target.
type UserList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the user list.
	// User list resource names have the form:
	//
	// `customers/{customer_id}/userLists/{user_list_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the user list.
	Id *int64 `protobuf:"varint,25,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Output only. An option that indicates if a user may edit a list. Depends on
	// the list ownership and list type. For example, external remarketing user
	// lists are not editable.
	//
	// This field is read-only.
	ReadOnly *bool `protobuf:"varint,26,opt,name=read_only,json=readOnly,proto3,oneof" json:"read_only,omitempty"`
	// Name of this user list. Depending on its access_reason, the user list name
	// may not be unique (for example, if access_reason=SHARED)
	Name *string `protobuf:"bytes,27,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Description of this user list.
	Description *string `protobuf:"bytes,28,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Membership status of this user list. Indicates whether a user list is open
	// or active. Only open user lists can accumulate more users and can be
	// targeted to.
	MembershipStatus enums.UserListMembershipStatusEnum_UserListMembershipStatus `protobuf:"varint,6,opt,name=membership_status,json=membershipStatus,proto3,enum=google.ads.googleads.v20.enums.UserListMembershipStatusEnum_UserListMembershipStatus" json:"membership_status,omitempty"`
	// An ID from external system. It is used by user list sellers to correlate
	// IDs on their systems.
	IntegrationCode *string `protobuf:"bytes,29,opt,name=integration_code,json=integrationCode,proto3,oneof" json:"integration_code,omitempty"`
	// Number of days a user's cookie stays on your list since its most recent
	// addition to the list. This field must be between 0 and 540 inclusive.
	// However, for CRM based userlists, this field can be set to 10000 which
	// means no expiration. Beginning on April 7, 2025, using a value of 10000 to
	// indicate no expiration will no longer be supported.
	//
	// This field is ignored for logical_user_list and rule_based_user_list types.
	// Membership to lists of these types depends on the rules defined by the
	// lists.
	MembershipLifeSpan *int64 `protobuf:"varint,30,opt,name=membership_life_span,json=membershipLifeSpan,proto3,oneof" json:"membership_life_span,omitempty"`
	// Output only. Estimated number of users in this user list, on the Google
	// Display Network. This value is null if the number of users has not yet been
	// determined.
	//
	// This field is read-only.
	SizeForDisplay *int64 `protobuf:"varint,31,opt,name=size_for_display,json=sizeForDisplay,proto3,oneof" json:"size_for_display,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, on the
	// Google Display Network.
	//
	// This field is read-only.
	SizeRangeForDisplay enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,10,opt,name=size_range_for_display,json=sizeRangeForDisplay,proto3,enum=google.ads.googleads.v20.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_display,omitempty"`
	// Output only. Estimated number of users in this user list in the google.com
	// domain. These are the users available for targeting in Search campaigns.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForSearch *int64 `protobuf:"varint,32,opt,name=size_for_search,json=sizeForSearch,proto3,oneof" json:"size_for_search,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, for
	// Search ads.
	//
	// This field is read-only.
	SizeRangeForSearch enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,12,opt,name=size_range_for_search,json=sizeRangeForSearch,proto3,enum=google.ads.googleads.v20.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_search,omitempty"`
	// Output only. Type of this list.
	//
	// This field is read-only.
	Type enums.UserListTypeEnum_UserListType `protobuf:"varint,13,opt,name=type,proto3,enum=google.ads.googleads.v20.enums.UserListTypeEnum_UserListType" json:"type,omitempty"`
	// Indicating the reason why this user list membership status is closed. It is
	// only populated on lists that were automatically closed due to inactivity,
	// and will be cleared once the list membership status becomes open.
	ClosingReason enums.UserListClosingReasonEnum_UserListClosingReason `protobuf:"varint,14,opt,name=closing_reason,json=closingReason,proto3,enum=google.ads.googleads.v20.enums.UserListClosingReasonEnum_UserListClosingReason" json:"closing_reason,omitempty"`
	// Output only. Indicates the reason this account has been granted access to
	// the list. The reason can be SHARED, OWNED, LICENSED or SUBSCRIBED.
	//
	// This field is read-only.
	AccessReason enums.AccessReasonEnum_AccessReason `protobuf:"varint,15,opt,name=access_reason,json=accessReason,proto3,enum=google.ads.googleads.v20.enums.AccessReasonEnum_AccessReason" json:"access_reason,omitempty"`
	// Indicates if this share is still enabled. When a UserList is shared with
	// the user this field is set to ENABLED. Later the userList owner can decide
	// to revoke the share and make it DISABLED.
	// The default value of this field is set to ENABLED.
	AccountUserListStatus enums.UserListAccessStatusEnum_UserListAccessStatus `protobuf:"varint,16,opt,name=account_user_list_status,json=accountUserListStatus,proto3,enum=google.ads.googleads.v20.enums.UserListAccessStatusEnum_UserListAccessStatus" json:"account_user_list_status,omitempty"`
	// Indicates if this user list is eligible for Google Search Network.
	EligibleForSearch *bool `protobuf:"varint,33,opt,name=eligible_for_search,json=eligibleForSearch,proto3,oneof" json:"eligible_for_search,omitempty"`
	// Output only. Indicates this user list is eligible for Google Display
	// Network.
	//
	// This field is read-only.
	EligibleForDisplay *bool `protobuf:"varint,34,opt,name=eligible_for_display,json=eligibleForDisplay,proto3,oneof" json:"eligible_for_display,omitempty"`
	// Output only. Indicates match rate for Customer Match lists. The range of
	// this field is [0-100]. This will be null for other list types or when it's
	// not possible to calculate the match rate.
	//
	// This field is read-only.
	MatchRatePercentage *int32 `protobuf:"varint,24,opt,name=match_rate_percentage,json=matchRatePercentage,proto3,oneof" json:"match_rate_percentage,omitempty"`
	// The user list.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to UserList:
	//
	//	*UserList_CrmBasedUserList
	//	*UserList_SimilarUserList
	//	*UserList_RuleBasedUserList
	//	*UserList_LogicalUserList
	//	*UserList_BasicUserList
	//	*UserList_LookalikeUserList
	UserList      isUserList_UserList `protobuf_oneof:"user_list"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserList) Reset() {
	*x = UserList{}
	mi := &file_google_ads_googleads_v20_resources_user_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_user_list_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_user_list_proto_rawDescGZIP(), []int{0}
}

func (x *UserList) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *UserList) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *UserList) GetReadOnly() bool {
	if x != nil && x.ReadOnly != nil {
		return *x.ReadOnly
	}
	return false
}

func (x *UserList) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserList) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UserList) GetMembershipStatus() enums.UserListMembershipStatusEnum_UserListMembershipStatus {
	if x != nil {
		return x.MembershipStatus
	}
	return enums.UserListMembershipStatusEnum_UserListMembershipStatus(0)
}

func (x *UserList) GetIntegrationCode() string {
	if x != nil && x.IntegrationCode != nil {
		return *x.IntegrationCode
	}
	return ""
}

func (x *UserList) GetMembershipLifeSpan() int64 {
	if x != nil && x.MembershipLifeSpan != nil {
		return *x.MembershipLifeSpan
	}
	return 0
}

func (x *UserList) GetSizeForDisplay() int64 {
	if x != nil && x.SizeForDisplay != nil {
		return *x.SizeForDisplay
	}
	return 0
}

func (x *UserList) GetSizeRangeForDisplay() enums.UserListSizeRangeEnum_UserListSizeRange {
	if x != nil {
		return x.SizeRangeForDisplay
	}
	return enums.UserListSizeRangeEnum_UserListSizeRange(0)
}

func (x *UserList) GetSizeForSearch() int64 {
	if x != nil && x.SizeForSearch != nil {
		return *x.SizeForSearch
	}
	return 0
}

func (x *UserList) GetSizeRangeForSearch() enums.UserListSizeRangeEnum_UserListSizeRange {
	if x != nil {
		return x.SizeRangeForSearch
	}
	return enums.UserListSizeRangeEnum_UserListSizeRange(0)
}

func (x *UserList) GetType() enums.UserListTypeEnum_UserListType {
	if x != nil {
		return x.Type
	}
	return enums.UserListTypeEnum_UserListType(0)
}

func (x *UserList) GetClosingReason() enums.UserListClosingReasonEnum_UserListClosingReason {
	if x != nil {
		return x.ClosingReason
	}
	return enums.UserListClosingReasonEnum_UserListClosingReason(0)
}

func (x *UserList) GetAccessReason() enums.AccessReasonEnum_AccessReason {
	if x != nil {
		return x.AccessReason
	}
	return enums.AccessReasonEnum_AccessReason(0)
}

func (x *UserList) GetAccountUserListStatus() enums.UserListAccessStatusEnum_UserListAccessStatus {
	if x != nil {
		return x.AccountUserListStatus
	}
	return enums.UserListAccessStatusEnum_UserListAccessStatus(0)
}

func (x *UserList) GetEligibleForSearch() bool {
	if x != nil && x.EligibleForSearch != nil {
		return *x.EligibleForSearch
	}
	return false
}

func (x *UserList) GetEligibleForDisplay() bool {
	if x != nil && x.EligibleForDisplay != nil {
		return *x.EligibleForDisplay
	}
	return false
}

func (x *UserList) GetMatchRatePercentage() int32 {
	if x != nil && x.MatchRatePercentage != nil {
		return *x.MatchRatePercentage
	}
	return 0
}

func (x *UserList) GetUserList() isUserList_UserList {
	if x != nil {
		return x.UserList
	}
	return nil
}

func (x *UserList) GetCrmBasedUserList() *common.CrmBasedUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_CrmBasedUserList); ok {
			return x.CrmBasedUserList
		}
	}
	return nil
}

func (x *UserList) GetSimilarUserList() *common.SimilarUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_SimilarUserList); ok {
			return x.SimilarUserList
		}
	}
	return nil
}

func (x *UserList) GetRuleBasedUserList() *common.RuleBasedUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_RuleBasedUserList); ok {
			return x.RuleBasedUserList
		}
	}
	return nil
}

func (x *UserList) GetLogicalUserList() *common.LogicalUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_LogicalUserList); ok {
			return x.LogicalUserList
		}
	}
	return nil
}

func (x *UserList) GetBasicUserList() *common.BasicUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_BasicUserList); ok {
			return x.BasicUserList
		}
	}
	return nil
}

func (x *UserList) GetLookalikeUserList() *common.LookalikeUserListInfo {
	if x != nil {
		if x, ok := x.UserList.(*UserList_LookalikeUserList); ok {
			return x.LookalikeUserList
		}
	}
	return nil
}

type isUserList_UserList interface {
	isUserList_UserList()
}

type UserList_CrmBasedUserList struct {
	// User list of CRM users provided by the advertiser.
	CrmBasedUserList *common.CrmBasedUserListInfo `protobuf:"bytes,19,opt,name=crm_based_user_list,json=crmBasedUserList,proto3,oneof"`
}

type UserList_SimilarUserList struct {
	// Output only. User list which are similar to users from another UserList.
	// These lists are readonly and automatically created by google.
	SimilarUserList *common.SimilarUserListInfo `protobuf:"bytes,20,opt,name=similar_user_list,json=similarUserList,proto3,oneof"`
}

type UserList_RuleBasedUserList struct {
	// User list generated by a rule.
	RuleBasedUserList *common.RuleBasedUserListInfo `protobuf:"bytes,21,opt,name=rule_based_user_list,json=ruleBasedUserList,proto3,oneof"`
}

type UserList_LogicalUserList struct {
	// User list that is a custom combination of user lists and user interests.
	LogicalUserList *common.LogicalUserListInfo `protobuf:"bytes,22,opt,name=logical_user_list,json=logicalUserList,proto3,oneof"`
}

type UserList_BasicUserList struct {
	// User list targeting as a collection of conversion or remarketing actions.
	BasicUserList *common.BasicUserListInfo `protobuf:"bytes,23,opt,name=basic_user_list,json=basicUserList,proto3,oneof"`
}

type UserList_LookalikeUserList struct {
	// Immutable. Lookalike User List.
	LookalikeUserList *common.LookalikeUserListInfo `protobuf:"bytes,36,opt,name=lookalike_user_list,json=lookalikeUserList,proto3,oneof"`
}

func (*UserList_CrmBasedUserList) isUserList_UserList() {}

func (*UserList_SimilarUserList) isUserList_UserList() {}

func (*UserList_RuleBasedUserList) isUserList_UserList() {}

func (*UserList_LogicalUserList) isUserList_UserList() {}

func (*UserList_BasicUserList) isUserList_UserList() {}

func (*UserList_LookalikeUserList) isUserList_UserList() {}

var File_google_ads_googleads_v20_resources_user_list_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_user_list_proto_rawDesc = "" +
	"\n" +
	"2google/ads/googleads/v20/resources/user_list.proto\x12\"google.ads.googleads.v20.resources\x1a0google/ads/googleads/v20/common/user_lists.proto\x1a2google/ads/googleads/v20/enums/access_reason.proto\x1a<google/ads/googleads/v20/enums/user_list_access_status.proto\x1a=google/ads/googleads/v20/enums/user_list_closing_reason.proto\x1a@google/ads/googleads/v20/enums/user_list_membership_status.proto\x1a9google/ads/googleads/v20/enums/user_list_size_range.proto\x1a3google/ads/googleads/v20/enums/user_list_type.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xc7\x12\n" +
	"\bUserList\x12N\n" +
	"\rresource_name\x18\x01 \x01(\tB)\xe0A\x05\xfaA#\n" +
	"!googleads.googleapis.com/UserListR\fresourceName\x12\x18\n" +
	"\x02id\x18\x19 \x01(\x03B\x03\xe0A\x03H\x01R\x02id\x88\x01\x01\x12%\n" +
	"\tread_only\x18\x1a \x01(\bB\x03\xe0A\x03H\x02R\breadOnly\x88\x01\x01\x12\x17\n" +
	"\x04name\x18\x1b \x01(\tH\x03R\x04name\x88\x01\x01\x12%\n" +
	"\vdescription\x18\x1c \x01(\tH\x04R\vdescription\x88\x01\x01\x12\x82\x01\n" +
	"\x11membership_status\x18\x06 \x01(\x0e2U.google.ads.googleads.v20.enums.UserListMembershipStatusEnum.UserListMembershipStatusR\x10membershipStatus\x12.\n" +
	"\x10integration_code\x18\x1d \x01(\tH\x05R\x0fintegrationCode\x88\x01\x01\x125\n" +
	"\x14membership_life_span\x18\x1e \x01(\x03H\x06R\x12membershipLifeSpan\x88\x01\x01\x122\n" +
	"\x10size_for_display\x18\x1f \x01(\x03B\x03\xe0A\x03H\aR\x0esizeForDisplay\x88\x01\x01\x12\x81\x01\n" +
	"\x16size_range_for_display\x18\n" +
	" \x01(\x0e2G.google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRangeB\x03\xe0A\x03R\x13sizeRangeForDisplay\x120\n" +
	"\x0fsize_for_search\x18  \x01(\x03B\x03\xe0A\x03H\bR\rsizeForSearch\x88\x01\x01\x12\x7f\n" +
	"\x15size_range_for_search\x18\f \x01(\x0e2G.google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRangeB\x03\xe0A\x03R\x12sizeRangeForSearch\x12V\n" +
	"\x04type\x18\r \x01(\x0e2=.google.ads.googleads.v20.enums.UserListTypeEnum.UserListTypeB\x03\xe0A\x03R\x04type\x12v\n" +
	"\x0eclosing_reason\x18\x0e \x01(\x0e2O.google.ads.googleads.v20.enums.UserListClosingReasonEnum.UserListClosingReasonR\rclosingReason\x12g\n" +
	"\raccess_reason\x18\x0f \x01(\x0e2=.google.ads.googleads.v20.enums.AccessReasonEnum.AccessReasonB\x03\xe0A\x03R\faccessReason\x12\x86\x01\n" +
	"\x18account_user_list_status\x18\x10 \x01(\x0e2M.google.ads.googleads.v20.enums.UserListAccessStatusEnum.UserListAccessStatusR\x15accountUserListStatus\x123\n" +
	"\x13eligible_for_search\x18! \x01(\bH\tR\x11eligibleForSearch\x88\x01\x01\x12:\n" +
	"\x14eligible_for_display\x18\" \x01(\bB\x03\xe0A\x03H\n" +
	"R\x12eligibleForDisplay\x88\x01\x01\x12<\n" +
	"\x15match_rate_percentage\x18\x18 \x01(\x05B\x03\xe0A\x03H\vR\x13matchRatePercentage\x88\x01\x01\x12f\n" +
	"\x13crm_based_user_list\x18\x13 \x01(\v25.google.ads.googleads.v20.common.CrmBasedUserListInfoH\x00R\x10crmBasedUserList\x12g\n" +
	"\x11similar_user_list\x18\x14 \x01(\v24.google.ads.googleads.v20.common.SimilarUserListInfoB\x03\xe0A\x03H\x00R\x0fsimilarUserList\x12i\n" +
	"\x14rule_based_user_list\x18\x15 \x01(\v26.google.ads.googleads.v20.common.RuleBasedUserListInfoH\x00R\x11ruleBasedUserList\x12b\n" +
	"\x11logical_user_list\x18\x16 \x01(\v24.google.ads.googleads.v20.common.LogicalUserListInfoH\x00R\x0flogicalUserList\x12\\\n" +
	"\x0fbasic_user_list\x18\x17 \x01(\v22.google.ads.googleads.v20.common.BasicUserListInfoH\x00R\rbasicUserList\x12m\n" +
	"\x13lookalike_user_list\x18$ \x01(\v26.google.ads.googleads.v20.common.LookalikeUserListInfoB\x03\xe0A\x05H\x00R\x11lookalikeUserList:X\xeaAU\n" +
	"!googleads.googleapis.com/UserList\x120customers/{customer_id}/userLists/{user_list_id}B\v\n" +
	"\tuser_listB\x05\n" +
	"\x03_idB\f\n" +
	"\n" +
	"_read_onlyB\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_descriptionB\x13\n" +
	"\x11_integration_codeB\x17\n" +
	"\x15_membership_life_spanB\x13\n" +
	"\x11_size_for_displayB\x12\n" +
	"\x10_size_for_searchB\x16\n" +
	"\x14_eligible_for_searchB\x17\n" +
	"\x15_eligible_for_displayB\x18\n" +
	"\x16_match_rate_percentageB\xff\x01\n" +
	"&com.google.ads.googleads.v20.resourcesB\rUserListProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_user_list_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_user_list_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_user_list_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_user_list_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_user_list_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_user_list_proto_rawDesc), len(file_google_ads_googleads_v20_resources_user_list_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_user_list_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_user_list_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_user_list_proto_goTypes = []any{
	(*UserList)(nil), // 0: google.ads.googleads.v20.resources.UserList
	(enums.UserListMembershipStatusEnum_UserListMembershipStatus)(0), // 1: google.ads.googleads.v20.enums.UserListMembershipStatusEnum.UserListMembershipStatus
	(enums.UserListSizeRangeEnum_UserListSizeRange)(0),               // 2: google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRange
	(enums.UserListTypeEnum_UserListType)(0),                         // 3: google.ads.googleads.v20.enums.UserListTypeEnum.UserListType
	(enums.UserListClosingReasonEnum_UserListClosingReason)(0),       // 4: google.ads.googleads.v20.enums.UserListClosingReasonEnum.UserListClosingReason
	(enums.AccessReasonEnum_AccessReason)(0),                         // 5: google.ads.googleads.v20.enums.AccessReasonEnum.AccessReason
	(enums.UserListAccessStatusEnum_UserListAccessStatus)(0),         // 6: google.ads.googleads.v20.enums.UserListAccessStatusEnum.UserListAccessStatus
	(*common.CrmBasedUserListInfo)(nil),                              // 7: google.ads.googleads.v20.common.CrmBasedUserListInfo
	(*common.SimilarUserListInfo)(nil),                               // 8: google.ads.googleads.v20.common.SimilarUserListInfo
	(*common.RuleBasedUserListInfo)(nil),                             // 9: google.ads.googleads.v20.common.RuleBasedUserListInfo
	(*common.LogicalUserListInfo)(nil),                               // 10: google.ads.googleads.v20.common.LogicalUserListInfo
	(*common.BasicUserListInfo)(nil),                                 // 11: google.ads.googleads.v20.common.BasicUserListInfo
	(*common.LookalikeUserListInfo)(nil),                             // 12: google.ads.googleads.v20.common.LookalikeUserListInfo
}
var file_google_ads_googleads_v20_resources_user_list_proto_depIdxs = []int32{
	1,  // 0: google.ads.googleads.v20.resources.UserList.membership_status:type_name -> google.ads.googleads.v20.enums.UserListMembershipStatusEnum.UserListMembershipStatus
	2,  // 1: google.ads.googleads.v20.resources.UserList.size_range_for_display:type_name -> google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRange
	2,  // 2: google.ads.googleads.v20.resources.UserList.size_range_for_search:type_name -> google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRange
	3,  // 3: google.ads.googleads.v20.resources.UserList.type:type_name -> google.ads.googleads.v20.enums.UserListTypeEnum.UserListType
	4,  // 4: google.ads.googleads.v20.resources.UserList.closing_reason:type_name -> google.ads.googleads.v20.enums.UserListClosingReasonEnum.UserListClosingReason
	5,  // 5: google.ads.googleads.v20.resources.UserList.access_reason:type_name -> google.ads.googleads.v20.enums.AccessReasonEnum.AccessReason
	6,  // 6: google.ads.googleads.v20.resources.UserList.account_user_list_status:type_name -> google.ads.googleads.v20.enums.UserListAccessStatusEnum.UserListAccessStatus
	7,  // 7: google.ads.googleads.v20.resources.UserList.crm_based_user_list:type_name -> google.ads.googleads.v20.common.CrmBasedUserListInfo
	8,  // 8: google.ads.googleads.v20.resources.UserList.similar_user_list:type_name -> google.ads.googleads.v20.common.SimilarUserListInfo
	9,  // 9: google.ads.googleads.v20.resources.UserList.rule_based_user_list:type_name -> google.ads.googleads.v20.common.RuleBasedUserListInfo
	10, // 10: google.ads.googleads.v20.resources.UserList.logical_user_list:type_name -> google.ads.googleads.v20.common.LogicalUserListInfo
	11, // 11: google.ads.googleads.v20.resources.UserList.basic_user_list:type_name -> google.ads.googleads.v20.common.BasicUserListInfo
	12, // 12: google.ads.googleads.v20.resources.UserList.lookalike_user_list:type_name -> google.ads.googleads.v20.common.LookalikeUserListInfo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_user_list_proto_init() }
func file_google_ads_googleads_v20_resources_user_list_proto_init() {
	if File_google_ads_googleads_v20_resources_user_list_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_user_list_proto_msgTypes[0].OneofWrappers = []any{
		(*UserList_CrmBasedUserList)(nil),
		(*UserList_SimilarUserList)(nil),
		(*UserList_RuleBasedUserList)(nil),
		(*UserList_LogicalUserList)(nil),
		(*UserList_BasicUserList)(nil),
		(*UserList_LookalikeUserList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_user_list_proto_rawDesc), len(file_google_ads_googleads_v20_resources_user_list_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_user_list_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_user_list_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_user_list_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_user_list_proto = out.File
	file_google_ads_googleads_v20_resources_user_list_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_user_list_proto_depIdxs = nil
}
