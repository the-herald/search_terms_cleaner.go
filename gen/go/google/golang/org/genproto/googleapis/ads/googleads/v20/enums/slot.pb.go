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
// source: google/ads/googleads/v20/enums/slot.proto

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

// Enumerates possible positions of the Ad.
type SlotEnum_Slot int32

const (
	// Not specified.
	SlotEnum_UNSPECIFIED SlotEnum_Slot = 0
	// The value is unknown in this version.
	SlotEnum_UNKNOWN SlotEnum_Slot = 1
	// Google search: Side.
	SlotEnum_SEARCH_SIDE SlotEnum_Slot = 2
	// Google search: Top.
	SlotEnum_SEARCH_TOP SlotEnum_Slot = 3
	// Google search: Other.
	SlotEnum_SEARCH_OTHER SlotEnum_Slot = 4
	// Google Display Network.
	SlotEnum_CONTENT SlotEnum_Slot = 5
	// Search partners: Top.
	SlotEnum_SEARCH_PARTNER_TOP SlotEnum_Slot = 6
	// Search partners: Other.
	SlotEnum_SEARCH_PARTNER_OTHER SlotEnum_Slot = 7
	// Cross-network.
	SlotEnum_MIXED SlotEnum_Slot = 8
)

// Enum value maps for SlotEnum_Slot.
var (
	SlotEnum_Slot_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "SEARCH_SIDE",
		3: "SEARCH_TOP",
		4: "SEARCH_OTHER",
		5: "CONTENT",
		6: "SEARCH_PARTNER_TOP",
		7: "SEARCH_PARTNER_OTHER",
		8: "MIXED",
	}
	SlotEnum_Slot_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"SEARCH_SIDE":          2,
		"SEARCH_TOP":           3,
		"SEARCH_OTHER":         4,
		"CONTENT":              5,
		"SEARCH_PARTNER_TOP":   6,
		"SEARCH_PARTNER_OTHER": 7,
		"MIXED":                8,
	}
)

func (x SlotEnum_Slot) Enum() *SlotEnum_Slot {
	p := new(SlotEnum_Slot)
	*p = x
	return p
}

func (x SlotEnum_Slot) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SlotEnum_Slot) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_slot_proto_enumTypes[0].Descriptor()
}

func (SlotEnum_Slot) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_slot_proto_enumTypes[0]
}

func (x SlotEnum_Slot) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SlotEnum_Slot.Descriptor instead.
func (SlotEnum_Slot) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_slot_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enumeration of possible positions of the Ad.
type SlotEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlotEnum) Reset() {
	*x = SlotEnum{}
	mi := &file_google_ads_googleads_v20_enums_slot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlotEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotEnum) ProtoMessage() {}

func (x *SlotEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_slot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotEnum.ProtoReflect.Descriptor instead.
func (*SlotEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_slot_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_slot_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_slot_proto_rawDesc = "" +
	"\n" +
	")google/ads/googleads/v20/enums/slot.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xae\x01\n" +
	"\bSlotEnum\"\xa1\x01\n" +
	"\x04Slot\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x0f\n" +
	"\vSEARCH_SIDE\x10\x02\x12\x0e\n" +
	"\n" +
	"SEARCH_TOP\x10\x03\x12\x10\n" +
	"\fSEARCH_OTHER\x10\x04\x12\v\n" +
	"\aCONTENT\x10\x05\x12\x16\n" +
	"\x12SEARCH_PARTNER_TOP\x10\x06\x12\x18\n" +
	"\x14SEARCH_PARTNER_OTHER\x10\a\x12\t\n" +
	"\x05MIXED\x10\bB\xe3\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\tSlotProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_slot_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_slot_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_slot_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_slot_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_slot_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_slot_proto_rawDesc), len(file_google_ads_googleads_v20_enums_slot_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_slot_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_slot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_slot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_slot_proto_goTypes = []any{
	(SlotEnum_Slot)(0), // 0: google.ads.googleads.v20.enums.SlotEnum.Slot
	(*SlotEnum)(nil),   // 1: google.ads.googleads.v20.enums.SlotEnum
}
var file_google_ads_googleads_v20_enums_slot_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_slot_proto_init() }
func file_google_ads_googleads_v20_enums_slot_proto_init() {
	if File_google_ads_googleads_v20_enums_slot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_slot_proto_rawDesc), len(file_google_ads_googleads_v20_enums_slot_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_slot_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_slot_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_slot_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_slot_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_slot_proto = out.File
	file_google_ads_googleads_v20_enums_slot_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_slot_proto_depIdxs = nil
}
