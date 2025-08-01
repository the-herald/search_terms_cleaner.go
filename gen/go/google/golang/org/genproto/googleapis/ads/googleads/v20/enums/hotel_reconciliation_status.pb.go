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
// source: google/ads/googleads/v20/enums/hotel_reconciliation_status.proto

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

// Status of the hotel booking reconciliation.
type HotelReconciliationStatusEnum_HotelReconciliationStatus int32

const (
	// Not specified.
	HotelReconciliationStatusEnum_UNSPECIFIED HotelReconciliationStatusEnum_HotelReconciliationStatus = 0
	// The value is unknown in this version.
	HotelReconciliationStatusEnum_UNKNOWN HotelReconciliationStatusEnum_HotelReconciliationStatus = 1
	// Bookings are for a future date, or a stay is underway but the check-out
	// date hasn't passed. An active reservation can't be reconciled.
	HotelReconciliationStatusEnum_RESERVATION_ENABLED HotelReconciliationStatusEnum_HotelReconciliationStatus = 2
	// Check-out has already taken place, or the booked dates have passed
	// without cancellation. Bookings that are not reconciled within 45 days of
	// the check-out date are billed based on the original booking price.
	HotelReconciliationStatusEnum_RECONCILIATION_NEEDED HotelReconciliationStatusEnum_HotelReconciliationStatus = 3
	// These bookings have been reconciled. Reconciled bookings are billed 45
	// days after the check-out date.
	HotelReconciliationStatusEnum_RECONCILED HotelReconciliationStatusEnum_HotelReconciliationStatus = 4
	// This booking was marked as canceled. Canceled stays with a value greater
	// than zero (due to minimum stay rules or cancellation fees) are billed 45
	// days after the check-out date.
	HotelReconciliationStatusEnum_CANCELED HotelReconciliationStatusEnum_HotelReconciliationStatus = 5
)

// Enum value maps for HotelReconciliationStatusEnum_HotelReconciliationStatus.
var (
	HotelReconciliationStatusEnum_HotelReconciliationStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "RESERVATION_ENABLED",
		3: "RECONCILIATION_NEEDED",
		4: "RECONCILED",
		5: "CANCELED",
	}
	HotelReconciliationStatusEnum_HotelReconciliationStatus_value = map[string]int32{
		"UNSPECIFIED":           0,
		"UNKNOWN":               1,
		"RESERVATION_ENABLED":   2,
		"RECONCILIATION_NEEDED": 3,
		"RECONCILED":            4,
		"CANCELED":              5,
	}
)

func (x HotelReconciliationStatusEnum_HotelReconciliationStatus) Enum() *HotelReconciliationStatusEnum_HotelReconciliationStatus {
	p := new(HotelReconciliationStatusEnum_HotelReconciliationStatus)
	*p = x
	return p
}

func (x HotelReconciliationStatusEnum_HotelReconciliationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelReconciliationStatusEnum_HotelReconciliationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_enumTypes[0].Descriptor()
}

func (HotelReconciliationStatusEnum_HotelReconciliationStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_enumTypes[0]
}

func (x HotelReconciliationStatusEnum_HotelReconciliationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelReconciliationStatusEnum_HotelReconciliationStatus.Descriptor instead.
func (HotelReconciliationStatusEnum_HotelReconciliationStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for HotelReconciliationStatus.
type HotelReconciliationStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HotelReconciliationStatusEnum) Reset() {
	*x = HotelReconciliationStatusEnum{}
	mi := &file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelReconciliationStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelReconciliationStatusEnum) ProtoMessage() {}

func (x *HotelReconciliationStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelReconciliationStatusEnum.ProtoReflect.Descriptor instead.
func (*HotelReconciliationStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDesc = "" +
	"\n" +
	"@google/ads/googleads/v20/enums/hotel_reconciliation_status.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xad\x01\n" +
	"\x1dHotelReconciliationStatusEnum\"\x8b\x01\n" +
	"\x19HotelReconciliationStatus\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x17\n" +
	"\x13RESERVATION_ENABLED\x10\x02\x12\x19\n" +
	"\x15RECONCILIATION_NEEDED\x10\x03\x12\x0e\n" +
	"\n" +
	"RECONCILED\x10\x04\x12\f\n" +
	"\bCANCELED\x10\x05B\xf8\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x1eHotelReconciliationStatusProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_goTypes = []any{
	(HotelReconciliationStatusEnum_HotelReconciliationStatus)(0), // 0: google.ads.googleads.v20.enums.HotelReconciliationStatusEnum.HotelReconciliationStatus
	(*HotelReconciliationStatusEnum)(nil),                        // 1: google.ads.googleads.v20.enums.HotelReconciliationStatusEnum
}
var file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_init() }
func file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_init() {
	if File_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto = out.File
	file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_hotel_reconciliation_status_proto_depIdxs = nil
}
