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
// source: google/ads/googleads/v20/common/dates.proto

package common

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
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

// A date range.
type DateRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The start date, in yyyy-mm-dd format. This date is inclusive.
	StartDate *string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format. This date is inclusive.
	EndDate       *string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_dates_proto_rawDescGZIP(), []int{0}
}

func (x *DateRange) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *DateRange) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

// The year month range inclusive of the start and end months.
// Eg: A year month range to represent Jan 2020 would be: (Jan 2020, Jan 2020).
type YearMonthRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The inclusive start year month.
	Start *YearMonth `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The inclusive end year month.
	End           *YearMonth `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YearMonthRange) Reset() {
	*x = YearMonthRange{}
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YearMonthRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YearMonthRange) ProtoMessage() {}

func (x *YearMonthRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YearMonthRange.ProtoReflect.Descriptor instead.
func (*YearMonthRange) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_dates_proto_rawDescGZIP(), []int{1}
}

func (x *YearMonthRange) GetStart() *YearMonth {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *YearMonthRange) GetEnd() *YearMonth {
	if x != nil {
		return x.End
	}
	return nil
}

// Year month.
type YearMonth struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The year (for example, 2020).
	Year int64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// The month of the year. (for example, FEBRUARY).
	Month         enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,2,opt,name=month,proto3,enum=google.ads.googleads.v20.enums.MonthOfYearEnum_MonthOfYear" json:"month,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YearMonth) Reset() {
	*x = YearMonth{}
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YearMonth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YearMonth) ProtoMessage() {}

func (x *YearMonth) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_common_dates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YearMonth.ProtoReflect.Descriptor instead.
func (*YearMonth) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_common_dates_proto_rawDescGZIP(), []int{2}
}

func (x *YearMonth) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *YearMonth) GetMonth() enums.MonthOfYearEnum_MonthOfYear {
	if x != nil {
		return x.Month
	}
	return enums.MonthOfYearEnum_MonthOfYear(0)
}

var File_google_ads_googleads_v20_common_dates_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_common_dates_proto_rawDesc = "" +
	"\n" +
	"+google/ads/googleads/v20/common/dates.proto\x12\x1fgoogle.ads.googleads.v20.common\x1a2google/ads/googleads/v20/enums/month_of_year.proto\"k\n" +
	"\tDateRange\x12\"\n" +
	"\n" +
	"start_date\x18\x03 \x01(\tH\x00R\tstartDate\x88\x01\x01\x12\x1e\n" +
	"\bend_date\x18\x04 \x01(\tH\x01R\aendDate\x88\x01\x01B\r\n" +
	"\v_start_dateB\v\n" +
	"\t_end_date\"\x90\x01\n" +
	"\x0eYearMonthRange\x12@\n" +
	"\x05start\x18\x01 \x01(\v2*.google.ads.googleads.v20.common.YearMonthR\x05start\x12<\n" +
	"\x03end\x18\x02 \x01(\v2*.google.ads.googleads.v20.common.YearMonthR\x03end\"r\n" +
	"\tYearMonth\x12\x12\n" +
	"\x04year\x18\x01 \x01(\x03R\x04year\x12Q\n" +
	"\x05month\x18\x02 \x01(\x0e2;.google.ads.googleads.v20.enums.MonthOfYearEnum.MonthOfYearR\x05monthB\xea\x01\n" +
	"#com.google.ads.googleads.v20.commonB\n" +
	"DatesProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/common;common\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Common\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Common\xea\x02#Google::Ads::GoogleAds::V20::Commonb\x06proto3"

var (
	file_google_ads_googleads_v20_common_dates_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_common_dates_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_common_dates_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_common_dates_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_common_dates_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_dates_proto_rawDesc), len(file_google_ads_googleads_v20_common_dates_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_common_dates_proto_rawDescData
}

var file_google_ads_googleads_v20_common_dates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v20_common_dates_proto_goTypes = []any{
	(*DateRange)(nil),                      // 0: google.ads.googleads.v20.common.DateRange
	(*YearMonthRange)(nil),                 // 1: google.ads.googleads.v20.common.YearMonthRange
	(*YearMonth)(nil),                      // 2: google.ads.googleads.v20.common.YearMonth
	(enums.MonthOfYearEnum_MonthOfYear)(0), // 3: google.ads.googleads.v20.enums.MonthOfYearEnum.MonthOfYear
}
var file_google_ads_googleads_v20_common_dates_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v20.common.YearMonthRange.start:type_name -> google.ads.googleads.v20.common.YearMonth
	2, // 1: google.ads.googleads.v20.common.YearMonthRange.end:type_name -> google.ads.googleads.v20.common.YearMonth
	3, // 2: google.ads.googleads.v20.common.YearMonth.month:type_name -> google.ads.googleads.v20.enums.MonthOfYearEnum.MonthOfYear
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_common_dates_proto_init() }
func file_google_ads_googleads_v20_common_dates_proto_init() {
	if File_google_ads_googleads_v20_common_dates_proto != nil {
		return
	}
	file_google_ads_googleads_v20_common_dates_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_common_dates_proto_rawDesc), len(file_google_ads_googleads_v20_common_dates_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_common_dates_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_common_dates_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_common_dates_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_common_dates_proto = out.File
	file_google_ads_googleads_v20_common_dates_proto_goTypes = nil
	file_google_ads_googleads_v20_common_dates_proto_depIdxs = nil
}
