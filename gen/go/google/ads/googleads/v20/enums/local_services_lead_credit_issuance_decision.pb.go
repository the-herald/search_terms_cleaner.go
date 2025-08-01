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
// source: google/ads/googleads/v20/enums/local_services_lead_credit_issuance_decision.proto

package enums

import (
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

// Decision of bonus credit issued or rejected.
type LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision int32

const (
	// Not specified.
	LocalServicesLeadCreditIssuanceDecisionEnum_UNSPECIFIED LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 0
	// Used for return value only. Represents value unknown in this version.
	LocalServicesLeadCreditIssuanceDecisionEnum_UNKNOWN LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 1
	// Bonus credit is issued successfully and bonus credit cap has not reached
	// the threshold after issuing this bonus credit.
	LocalServicesLeadCreditIssuanceDecisionEnum_SUCCESS_NOT_REACHED_THRESHOLD LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 2
	// Bonus credit is issued successfully and bonus credit cap has reached the
	// threshold after issuing this bonus credit.
	LocalServicesLeadCreditIssuanceDecisionEnum_SUCCESS_REACHED_THRESHOLD LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 3
	// Bonus credit is not issued because the provider has reached the bonus
	// credit cap.
	LocalServicesLeadCreditIssuanceDecisionEnum_FAIL_OVER_THRESHOLD LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 4
	// Bonus credit is not issued because this lead is not eligible for bonus
	// credit.
	LocalServicesLeadCreditIssuanceDecisionEnum_FAIL_NOT_ELIGIBLE LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision = 5
)

// Enum value maps for LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision.
var (
	LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "SUCCESS_NOT_REACHED_THRESHOLD",
		3: "SUCCESS_REACHED_THRESHOLD",
		4: "FAIL_OVER_THRESHOLD",
		5: "FAIL_NOT_ELIGIBLE",
	}
	LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"SUCCESS_NOT_REACHED_THRESHOLD": 2,
		"SUCCESS_REACHED_THRESHOLD":     3,
		"FAIL_OVER_THRESHOLD":           4,
		"FAIL_NOT_ELIGIBLE":             5,
	}
)

func (x LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) Enum() *LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision {
	p := new(LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision)
	*p = x
	return p
}

func (x LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_enumTypes[0].Descriptor()
}

func (LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_enumTypes[0]
}

func (x LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision.Descriptor instead.
func (LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible credit issuance decisions for a lead.
type LocalServicesLeadCreditIssuanceDecisionEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocalServicesLeadCreditIssuanceDecisionEnum) Reset() {
	*x = LocalServicesLeadCreditIssuanceDecisionEnum{}
	mi := &file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalServicesLeadCreditIssuanceDecisionEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalServicesLeadCreditIssuanceDecisionEnum) ProtoMessage() {}

func (x *LocalServicesLeadCreditIssuanceDecisionEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalServicesLeadCreditIssuanceDecisionEnum.ProtoReflect.Descriptor instead.
func (*LocalServicesLeadCreditIssuanceDecisionEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDesc = "" +
	"\n" +
	"Qgoogle/ads/googleads/v20/enums/local_services_lead_credit_issuance_decision.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xd8\x01\n" +
	"+LocalServicesLeadCreditIssuanceDecisionEnum\"\xa8\x01\n" +
	"\x16CreditIssuanceDecision\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12!\n" +
	"\x1dSUCCESS_NOT_REACHED_THRESHOLD\x10\x02\x12\x1d\n" +
	"\x19SUCCESS_REACHED_THRESHOLD\x10\x03\x12\x17\n" +
	"\x13FAIL_OVER_THRESHOLD\x10\x04\x12\x15\n" +
	"\x11FAIL_NOT_ELIGIBLE\x10\x05B\x86\x02\n" +
	"\"com.google.ads.googleads.v20.enumsB,LocalServicesLeadCreditIssuanceDecisionProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDesc), len(file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_goTypes = []any{
	(LocalServicesLeadCreditIssuanceDecisionEnum_CreditIssuanceDecision)(0), // 0: google.ads.googleads.v20.enums.LocalServicesLeadCreditIssuanceDecisionEnum.CreditIssuanceDecision
	(*LocalServicesLeadCreditIssuanceDecisionEnum)(nil),                     // 1: google.ads.googleads.v20.enums.LocalServicesLeadCreditIssuanceDecisionEnum
}
var file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_init()
}
func file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_init() {
	if File_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDesc), len(file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto = out.File
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_local_services_lead_credit_issuance_decision_proto_depIdxs = nil
}
