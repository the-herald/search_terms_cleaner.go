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
// source: google/ads/googleads/v20/errors/query_error.proto

package errors

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

// Enum describing possible query errors.
type QueryErrorEnum_QueryError int32

const (
	// Name unspecified.
	QueryErrorEnum_UNSPECIFIED QueryErrorEnum_QueryError = 0
	// The received error code is not known in this version.
	QueryErrorEnum_UNKNOWN QueryErrorEnum_QueryError = 1
	// Returned if all other query error reasons are not applicable.
	QueryErrorEnum_QUERY_ERROR QueryErrorEnum_QueryError = 50
	// A condition used in the query references an invalid enum constant.
	QueryErrorEnum_BAD_ENUM_CONSTANT QueryErrorEnum_QueryError = 18
	// Query contains an invalid escape sequence.
	QueryErrorEnum_BAD_ESCAPE_SEQUENCE QueryErrorEnum_QueryError = 7
	// Field name is invalid.
	QueryErrorEnum_BAD_FIELD_NAME QueryErrorEnum_QueryError = 12
	// Limit value is invalid (for example, not a number)
	QueryErrorEnum_BAD_LIMIT_VALUE QueryErrorEnum_QueryError = 15
	// Encountered number can not be parsed.
	QueryErrorEnum_BAD_NUMBER QueryErrorEnum_QueryError = 5
	// Invalid operator encountered.
	QueryErrorEnum_BAD_OPERATOR QueryErrorEnum_QueryError = 3
	// Parameter unknown or not supported.
	QueryErrorEnum_BAD_PARAMETER_NAME QueryErrorEnum_QueryError = 61
	// Parameter have invalid value.
	QueryErrorEnum_BAD_PARAMETER_VALUE QueryErrorEnum_QueryError = 62
	// Invalid resource type was specified in the FROM clause.
	QueryErrorEnum_BAD_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 45
	// Non-ASCII symbol encountered outside of strings.
	QueryErrorEnum_BAD_SYMBOL QueryErrorEnum_QueryError = 2
	// Value is invalid.
	QueryErrorEnum_BAD_VALUE QueryErrorEnum_QueryError = 4
	// Date filters fail to restrict date to a range smaller than 31 days.
	// Applicable if the query is segmented by date.
	QueryErrorEnum_DATE_RANGE_TOO_WIDE QueryErrorEnum_QueryError = 36
	// Filters on date/week/month/quarter have a start date after
	// end date.
	QueryErrorEnum_DATE_RANGE_TOO_NARROW QueryErrorEnum_QueryError = 60
	// Expected AND between values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_AND QueryErrorEnum_QueryError = 30
	// Expecting ORDER BY to have BY.
	QueryErrorEnum_EXPECTED_BY QueryErrorEnum_QueryError = 14
	// There was no dimension field selected.
	QueryErrorEnum_EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 37
	// Missing filters on date related fields.
	QueryErrorEnum_EXPECTED_FILTERS_ON_DATE_RANGE QueryErrorEnum_QueryError = 55
	// Missing FROM clause.
	QueryErrorEnum_EXPECTED_FROM QueryErrorEnum_QueryError = 44
	// The operator used in the conditions requires the value to be a list.
	QueryErrorEnum_EXPECTED_LIST QueryErrorEnum_QueryError = 41
	// Fields used in WHERE or ORDER BY clauses are missing from the SELECT
	// clause.
	QueryErrorEnum_EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 16
	// SELECT is missing at the beginning of query.
	QueryErrorEnum_EXPECTED_SELECT QueryErrorEnum_QueryError = 13
	// A list was passed as a value to a condition whose operator expects a
	// single value.
	QueryErrorEnum_EXPECTED_SINGLE_VALUE QueryErrorEnum_QueryError = 42
	// Missing one or both values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 29
	// Invalid date format. Expected 'YYYY-MM-DD'.
	QueryErrorEnum_INVALID_DATE_FORMAT QueryErrorEnum_QueryError = 38
	// Misaligned date value for the filter. The date should be the start of a
	// week/month/quarter if the filtered field is
	// segments.week/segments.month/segments.quarter.
	QueryErrorEnum_MISALIGNED_DATE_FOR_FILTER QueryErrorEnum_QueryError = 64
	// Value passed was not a string when it should have been. For example, it
	// was a number or unquoted literal.
	QueryErrorEnum_INVALID_STRING_VALUE QueryErrorEnum_QueryError = 57
	// A String value passed to the BETWEEN operator does not parse as a date.
	QueryErrorEnum_INVALID_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 26
	// The value passed to the DURING operator is not a Date range literal
	QueryErrorEnum_INVALID_VALUE_WITH_DURING_OPERATOR QueryErrorEnum_QueryError = 22
	// An invalid value was passed to the LIKE operator.
	QueryErrorEnum_INVALID_VALUE_WITH_LIKE_OPERATOR QueryErrorEnum_QueryError = 56
	// An operator was provided that is inapplicable to the field being
	// filtered.
	QueryErrorEnum_OPERATOR_FIELD_MISMATCH QueryErrorEnum_QueryError = 35
	// A Condition was found with an empty list.
	QueryErrorEnum_PROHIBITED_EMPTY_LIST_IN_CONDITION QueryErrorEnum_QueryError = 28
	// A condition used in the query references an unsupported enum constant.
	QueryErrorEnum_PROHIBITED_ENUM_CONSTANT QueryErrorEnum_QueryError = 54
	// Fields that are not allowed to be selected together were included in
	// the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 31
	// A field that is not orderable was included in the ORDER BY clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE QueryErrorEnum_QueryError = 40
	// A field that is not selectable was included in the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 23
	// A field that is not filterable was included in the WHERE clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 24
	// Resource type specified in the FROM clause is not supported by this
	// service.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 43
	// A field that comes from an incompatible resource was included in the
	// SELECT clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 48
	// A field that comes from an incompatible resource was included in the
	// WHERE clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 58
	// A metric incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 49
	// A segment incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 51
	// A segment in the SELECT clause is incompatible with a metric in the
	// SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 53
	// The value passed to the limit clause is too low.
	QueryErrorEnum_LIMIT_VALUE_TOO_LOW QueryErrorEnum_QueryError = 25
	// Query has a string containing a newline character.
	QueryErrorEnum_PROHIBITED_NEWLINE_IN_STRING QueryErrorEnum_QueryError = 8
	// List contains values of different types.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_IN_LIST QueryErrorEnum_QueryError = 10
	// The values passed to the BETWEEN operator are not of the same type.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 21
	// Query contains unterminated string.
	QueryErrorEnum_STRING_NOT_TERMINATED QueryErrorEnum_QueryError = 6
	// Too many segments are specified in SELECT clause.
	QueryErrorEnum_TOO_MANY_SEGMENTS QueryErrorEnum_QueryError = 34
	// Query is incomplete and cannot be parsed.
	QueryErrorEnum_UNEXPECTED_END_OF_QUERY QueryErrorEnum_QueryError = 9
	// FROM clause cannot be specified in this query.
	QueryErrorEnum_UNEXPECTED_FROM_CLAUSE QueryErrorEnum_QueryError = 47
	// Query contains one or more unrecognized fields.
	QueryErrorEnum_UNRECOGNIZED_FIELD QueryErrorEnum_QueryError = 32
	// Query has an unexpected extra part.
	QueryErrorEnum_UNEXPECTED_INPUT QueryErrorEnum_QueryError = 11
	// Metrics cannot be requested for a manager account. To retrieve metrics,
	// issue separate requests against each client account under the manager
	// account.
	QueryErrorEnum_REQUESTED_METRICS_FOR_MANAGER QueryErrorEnum_QueryError = 59
	// The number of values (right-hand-side operands) in a filter exceeds the
	// limit.
	QueryErrorEnum_FILTER_HAS_TOO_MANY_VALUES QueryErrorEnum_QueryError = 63
)

// Enum value maps for QueryErrorEnum_QueryError.
var (
	QueryErrorEnum_QueryError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		50: "QUERY_ERROR",
		18: "BAD_ENUM_CONSTANT",
		7:  "BAD_ESCAPE_SEQUENCE",
		12: "BAD_FIELD_NAME",
		15: "BAD_LIMIT_VALUE",
		5:  "BAD_NUMBER",
		3:  "BAD_OPERATOR",
		61: "BAD_PARAMETER_NAME",
		62: "BAD_PARAMETER_VALUE",
		45: "BAD_RESOURCE_TYPE_IN_FROM_CLAUSE",
		2:  "BAD_SYMBOL",
		4:  "BAD_VALUE",
		36: "DATE_RANGE_TOO_WIDE",
		60: "DATE_RANGE_TOO_NARROW",
		30: "EXPECTED_AND",
		14: "EXPECTED_BY",
		37: "EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE",
		55: "EXPECTED_FILTERS_ON_DATE_RANGE",
		44: "EXPECTED_FROM",
		41: "EXPECTED_LIST",
		16: "EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE",
		13: "EXPECTED_SELECT",
		42: "EXPECTED_SINGLE_VALUE",
		29: "EXPECTED_VALUE_WITH_BETWEEN_OPERATOR",
		38: "INVALID_DATE_FORMAT",
		64: "MISALIGNED_DATE_FOR_FILTER",
		57: "INVALID_STRING_VALUE",
		26: "INVALID_VALUE_WITH_BETWEEN_OPERATOR",
		22: "INVALID_VALUE_WITH_DURING_OPERATOR",
		56: "INVALID_VALUE_WITH_LIKE_OPERATOR",
		35: "OPERATOR_FIELD_MISMATCH",
		28: "PROHIBITED_EMPTY_LIST_IN_CONDITION",
		54: "PROHIBITED_ENUM_CONSTANT",
		31: "PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE",
		40: "PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE",
		23: "PROHIBITED_FIELD_IN_SELECT_CLAUSE",
		24: "PROHIBITED_FIELD_IN_WHERE_CLAUSE",
		43: "PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE",
		48: "PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE",
		58: "PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE",
		49: "PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
		51: "PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE",
		53: "PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
		25: "LIMIT_VALUE_TOO_LOW",
		8:  "PROHIBITED_NEWLINE_IN_STRING",
		10: "PROHIBITED_VALUE_COMBINATION_IN_LIST",
		21: "PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR",
		6:  "STRING_NOT_TERMINATED",
		34: "TOO_MANY_SEGMENTS",
		9:  "UNEXPECTED_END_OF_QUERY",
		47: "UNEXPECTED_FROM_CLAUSE",
		32: "UNRECOGNIZED_FIELD",
		11: "UNEXPECTED_INPUT",
		59: "REQUESTED_METRICS_FOR_MANAGER",
		63: "FILTER_HAS_TOO_MANY_VALUES",
	}
	QueryErrorEnum_QueryError_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"QUERY_ERROR":                      50,
		"BAD_ENUM_CONSTANT":                18,
		"BAD_ESCAPE_SEQUENCE":              7,
		"BAD_FIELD_NAME":                   12,
		"BAD_LIMIT_VALUE":                  15,
		"BAD_NUMBER":                       5,
		"BAD_OPERATOR":                     3,
		"BAD_PARAMETER_NAME":               61,
		"BAD_PARAMETER_VALUE":              62,
		"BAD_RESOURCE_TYPE_IN_FROM_CLAUSE": 45,
		"BAD_SYMBOL":                       2,
		"BAD_VALUE":                        4,
		"DATE_RANGE_TOO_WIDE":              36,
		"DATE_RANGE_TOO_NARROW":            60,
		"EXPECTED_AND":                     30,
		"EXPECTED_BY":                      14,
		"EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE":                37,
		"EXPECTED_FILTERS_ON_DATE_RANGE":                           55,
		"EXPECTED_FROM":                                            44,
		"EXPECTED_LIST":                                            41,
		"EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE":               16,
		"EXPECTED_SELECT":                                          13,
		"EXPECTED_SINGLE_VALUE":                                    42,
		"EXPECTED_VALUE_WITH_BETWEEN_OPERATOR":                     29,
		"INVALID_DATE_FORMAT":                                      38,
		"MISALIGNED_DATE_FOR_FILTER":                               64,
		"INVALID_STRING_VALUE":                                     57,
		"INVALID_VALUE_WITH_BETWEEN_OPERATOR":                      26,
		"INVALID_VALUE_WITH_DURING_OPERATOR":                       22,
		"INVALID_VALUE_WITH_LIKE_OPERATOR":                         56,
		"OPERATOR_FIELD_MISMATCH":                                  35,
		"PROHIBITED_EMPTY_LIST_IN_CONDITION":                       28,
		"PROHIBITED_ENUM_CONSTANT":                                 54,
		"PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE":            31,
		"PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE":                      40,
		"PROHIBITED_FIELD_IN_SELECT_CLAUSE":                        23,
		"PROHIBITED_FIELD_IN_WHERE_CLAUSE":                         24,
		"PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE":                  43,
		"PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE":                48,
		"PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE":                 58,
		"PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE":              49,
		"PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE":             51,
		"PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE": 53,
		"LIMIT_VALUE_TOO_LOW":                                      25,
		"PROHIBITED_NEWLINE_IN_STRING":                             8,
		"PROHIBITED_VALUE_COMBINATION_IN_LIST":                     10,
		"PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR":       21,
		"STRING_NOT_TERMINATED":                                    6,
		"TOO_MANY_SEGMENTS":                                        34,
		"UNEXPECTED_END_OF_QUERY":                                  9,
		"UNEXPECTED_FROM_CLAUSE":                                   47,
		"UNRECOGNIZED_FIELD":                                       32,
		"UNEXPECTED_INPUT":                                         11,
		"REQUESTED_METRICS_FOR_MANAGER":                            59,
		"FILTER_HAS_TOO_MANY_VALUES":                               63,
	}
)

func (x QueryErrorEnum_QueryError) Enum() *QueryErrorEnum_QueryError {
	p := new(QueryErrorEnum_QueryError)
	*p = x
	return p
}

func (x QueryErrorEnum_QueryError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryErrorEnum_QueryError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_query_error_proto_enumTypes[0].Descriptor()
}

func (QueryErrorEnum_QueryError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_query_error_proto_enumTypes[0]
}

func (x QueryErrorEnum_QueryError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryErrorEnum_QueryError.Descriptor instead.
func (QueryErrorEnum_QueryError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_query_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible query errors.
type QueryErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryErrorEnum) Reset() {
	*x = QueryErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_query_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryErrorEnum) ProtoMessage() {}

func (x *QueryErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_query_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryErrorEnum.ProtoReflect.Descriptor instead.
func (*QueryErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_query_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_query_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_query_error_proto_rawDesc = "" +
	"\n" +
	"1google/ads/googleads/v20/errors/query_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x93\x0e\n" +
	"\x0eQueryErrorEnum\"\x80\x0e\n" +
	"\n" +
	"QueryError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x0f\n" +
	"\vQUERY_ERROR\x102\x12\x15\n" +
	"\x11BAD_ENUM_CONSTANT\x10\x12\x12\x17\n" +
	"\x13BAD_ESCAPE_SEQUENCE\x10\a\x12\x12\n" +
	"\x0eBAD_FIELD_NAME\x10\f\x12\x13\n" +
	"\x0fBAD_LIMIT_VALUE\x10\x0f\x12\x0e\n" +
	"\n" +
	"BAD_NUMBER\x10\x05\x12\x10\n" +
	"\fBAD_OPERATOR\x10\x03\x12\x16\n" +
	"\x12BAD_PARAMETER_NAME\x10=\x12\x17\n" +
	"\x13BAD_PARAMETER_VALUE\x10>\x12$\n" +
	" BAD_RESOURCE_TYPE_IN_FROM_CLAUSE\x10-\x12\x0e\n" +
	"\n" +
	"BAD_SYMBOL\x10\x02\x12\r\n" +
	"\tBAD_VALUE\x10\x04\x12\x17\n" +
	"\x13DATE_RANGE_TOO_WIDE\x10$\x12\x19\n" +
	"\x15DATE_RANGE_TOO_NARROW\x10<\x12\x10\n" +
	"\fEXPECTED_AND\x10\x1e\x12\x0f\n" +
	"\vEXPECTED_BY\x10\x0e\x12-\n" +
	")EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE\x10%\x12\"\n" +
	"\x1eEXPECTED_FILTERS_ON_DATE_RANGE\x107\x12\x11\n" +
	"\rEXPECTED_FROM\x10,\x12\x11\n" +
	"\rEXPECTED_LIST\x10)\x12.\n" +
	"*EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE\x10\x10\x12\x13\n" +
	"\x0fEXPECTED_SELECT\x10\r\x12\x19\n" +
	"\x15EXPECTED_SINGLE_VALUE\x10*\x12(\n" +
	"$EXPECTED_VALUE_WITH_BETWEEN_OPERATOR\x10\x1d\x12\x17\n" +
	"\x13INVALID_DATE_FORMAT\x10&\x12\x1e\n" +
	"\x1aMISALIGNED_DATE_FOR_FILTER\x10@\x12\x18\n" +
	"\x14INVALID_STRING_VALUE\x109\x12'\n" +
	"#INVALID_VALUE_WITH_BETWEEN_OPERATOR\x10\x1a\x12&\n" +
	"\"INVALID_VALUE_WITH_DURING_OPERATOR\x10\x16\x12$\n" +
	" INVALID_VALUE_WITH_LIKE_OPERATOR\x108\x12\x1b\n" +
	"\x17OPERATOR_FIELD_MISMATCH\x10#\x12&\n" +
	"\"PROHIBITED_EMPTY_LIST_IN_CONDITION\x10\x1c\x12\x1c\n" +
	"\x18PROHIBITED_ENUM_CONSTANT\x106\x121\n" +
	"-PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE\x10\x1f\x12'\n" +
	"#PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE\x10(\x12%\n" +
	"!PROHIBITED_FIELD_IN_SELECT_CLAUSE\x10\x17\x12$\n" +
	" PROHIBITED_FIELD_IN_WHERE_CLAUSE\x10\x18\x12+\n" +
	"'PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE\x10+\x12-\n" +
	")PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE\x100\x12,\n" +
	"(PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE\x10:\x12/\n" +
	"+PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE\x101\x120\n" +
	",PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE\x103\x12<\n" +
	"8PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE\x105\x12\x17\n" +
	"\x13LIMIT_VALUE_TOO_LOW\x10\x19\x12 \n" +
	"\x1cPROHIBITED_NEWLINE_IN_STRING\x10\b\x12(\n" +
	"$PROHIBITED_VALUE_COMBINATION_IN_LIST\x10\n" +
	"\x126\n" +
	"2PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR\x10\x15\x12\x19\n" +
	"\x15STRING_NOT_TERMINATED\x10\x06\x12\x15\n" +
	"\x11TOO_MANY_SEGMENTS\x10\"\x12\x1b\n" +
	"\x17UNEXPECTED_END_OF_QUERY\x10\t\x12\x1a\n" +
	"\x16UNEXPECTED_FROM_CLAUSE\x10/\x12\x16\n" +
	"\x12UNRECOGNIZED_FIELD\x10 \x12\x14\n" +
	"\x10UNEXPECTED_INPUT\x10\v\x12!\n" +
	"\x1dREQUESTED_METRICS_FOR_MANAGER\x10;\x12\x1e\n" +
	"\x1aFILTER_HAS_TOO_MANY_VALUES\x10?B\xef\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x0fQueryErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_query_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_query_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_query_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_query_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_query_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_query_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_query_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_query_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_query_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_query_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_query_error_proto_goTypes = []any{
	(QueryErrorEnum_QueryError)(0), // 0: google.ads.googleads.v20.errors.QueryErrorEnum.QueryError
	(*QueryErrorEnum)(nil),         // 1: google.ads.googleads.v20.errors.QueryErrorEnum
}
var file_google_ads_googleads_v20_errors_query_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_query_error_proto_init() }
func file_google_ads_googleads_v20_errors_query_error_proto_init() {
	if File_google_ads_googleads_v20_errors_query_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_query_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_query_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_query_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_query_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_query_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_query_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_query_error_proto = out.File
	file_google_ads_googleads_v20_errors_query_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_query_error_proto_depIdxs = nil
}
