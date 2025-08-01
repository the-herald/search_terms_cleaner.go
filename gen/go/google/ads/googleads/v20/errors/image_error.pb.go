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
// source: google/ads/googleads/v20/errors/image_error.proto

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

// Enum describing possible image errors.
type ImageErrorEnum_ImageError int32

const (
	// Enum unspecified.
	ImageErrorEnum_UNSPECIFIED ImageErrorEnum_ImageError = 0
	// The received error code is not known in this version.
	ImageErrorEnum_UNKNOWN ImageErrorEnum_ImageError = 1
	// The image is not valid.
	ImageErrorEnum_INVALID_IMAGE ImageErrorEnum_ImageError = 2
	// The image could not be stored.
	ImageErrorEnum_STORAGE_ERROR ImageErrorEnum_ImageError = 3
	// There was a problem with the request.
	ImageErrorEnum_BAD_REQUEST ImageErrorEnum_ImageError = 4
	// The image is not of legal dimensions.
	ImageErrorEnum_UNEXPECTED_SIZE ImageErrorEnum_ImageError = 5
	// Animated image are not permitted.
	ImageErrorEnum_ANIMATED_NOT_ALLOWED ImageErrorEnum_ImageError = 6
	// Animation is too long.
	ImageErrorEnum_ANIMATION_TOO_LONG ImageErrorEnum_ImageError = 7
	// There was an error on the server.
	ImageErrorEnum_SERVER_ERROR ImageErrorEnum_ImageError = 8
	// Image cannot be in CMYK color format.
	ImageErrorEnum_CMYK_JPEG_NOT_ALLOWED ImageErrorEnum_ImageError = 9
	// Flash images are not permitted.
	ImageErrorEnum_FLASH_NOT_ALLOWED ImageErrorEnum_ImageError = 10
	// Flash images must support clickTag.
	ImageErrorEnum_FLASH_WITHOUT_CLICKTAG ImageErrorEnum_ImageError = 11
	// A flash error has occurred after fixing the click tag.
	ImageErrorEnum_FLASH_ERROR_AFTER_FIXING_CLICK_TAG ImageErrorEnum_ImageError = 12
	// Unacceptable visual effects.
	ImageErrorEnum_ANIMATED_VISUAL_EFFECT ImageErrorEnum_ImageError = 13
	// There was a problem with the flash image.
	ImageErrorEnum_FLASH_ERROR ImageErrorEnum_ImageError = 14
	// Incorrect image layout.
	ImageErrorEnum_LAYOUT_PROBLEM ImageErrorEnum_ImageError = 15
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_IMAGE_FILE ImageErrorEnum_ImageError = 16
	// There was an error storing the image.
	ImageErrorEnum_ERROR_STORING_IMAGE ImageErrorEnum_ImageError = 17
	// The aspect ratio of the image is not allowed.
	ImageErrorEnum_ASPECT_RATIO_NOT_ALLOWED ImageErrorEnum_ImageError = 18
	// Flash cannot have network objects.
	ImageErrorEnum_FLASH_HAS_NETWORK_OBJECTS ImageErrorEnum_ImageError = 19
	// Flash cannot have network methods.
	ImageErrorEnum_FLASH_HAS_NETWORK_METHODS ImageErrorEnum_ImageError = 20
	// Flash cannot have a Url.
	ImageErrorEnum_FLASH_HAS_URL ImageErrorEnum_ImageError = 21
	// Flash cannot use mouse tracking.
	ImageErrorEnum_FLASH_HAS_MOUSE_TRACKING ImageErrorEnum_ImageError = 22
	// Flash cannot have a random number.
	ImageErrorEnum_FLASH_HAS_RANDOM_NUM ImageErrorEnum_ImageError = 23
	// Ad click target cannot be '_self'.
	ImageErrorEnum_FLASH_SELF_TARGETS ImageErrorEnum_ImageError = 24
	// GetUrl method should only use '_blank'.
	ImageErrorEnum_FLASH_BAD_GETURL_TARGET ImageErrorEnum_ImageError = 25
	// Flash version is not supported.
	ImageErrorEnum_FLASH_VERSION_NOT_SUPPORTED ImageErrorEnum_ImageError = 26
	// Flash movies need to have hard coded click URL or clickTAG
	ImageErrorEnum_FLASH_WITHOUT_HARD_CODED_CLICK_URL ImageErrorEnum_ImageError = 27
	// Uploaded flash file is corrupted.
	ImageErrorEnum_INVALID_FLASH_FILE ImageErrorEnum_ImageError = 28
	// Uploaded flash file can be parsed, but the click tag can not be fixed
	// properly.
	ImageErrorEnum_FAILED_TO_FIX_CLICK_TAG_IN_FLASH ImageErrorEnum_ImageError = 29
	// Flash movie accesses network resources
	ImageErrorEnum_FLASH_ACCESSES_NETWORK_RESOURCES ImageErrorEnum_ImageError = 30
	// Flash movie attempts to call external javascript code
	ImageErrorEnum_FLASH_EXTERNAL_JS_CALL ImageErrorEnum_ImageError = 31
	// Flash movie attempts to call flash system commands
	ImageErrorEnum_FLASH_EXTERNAL_FS_CALL ImageErrorEnum_ImageError = 32
	// Image file is too large.
	ImageErrorEnum_FILE_TOO_LARGE ImageErrorEnum_ImageError = 33
	// Image data is too large.
	ImageErrorEnum_IMAGE_DATA_TOO_LARGE ImageErrorEnum_ImageError = 34
	// Error while processing the image.
	ImageErrorEnum_IMAGE_PROCESSING_ERROR ImageErrorEnum_ImageError = 35
	// Image is too small.
	ImageErrorEnum_IMAGE_TOO_SMALL ImageErrorEnum_ImageError = 36
	// Input was invalid.
	ImageErrorEnum_INVALID_INPUT ImageErrorEnum_ImageError = 37
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_FILE ImageErrorEnum_ImageError = 38
	// Image constraints are violated, but details like ASPECT_RATIO_NOT_ALLOWED
	// can't be provided. This happens when asset spec contains more than one
	// constraint and different criteria of different constraints are violated.
	ImageErrorEnum_IMAGE_CONSTRAINTS_VIOLATED ImageErrorEnum_ImageError = 39
	// Image format is not allowed.
	ImageErrorEnum_FORMAT_NOT_ALLOWED ImageErrorEnum_ImageError = 40
)

// Enum value maps for ImageErrorEnum_ImageError.
var (
	ImageErrorEnum_ImageError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_IMAGE",
		3:  "STORAGE_ERROR",
		4:  "BAD_REQUEST",
		5:  "UNEXPECTED_SIZE",
		6:  "ANIMATED_NOT_ALLOWED",
		7:  "ANIMATION_TOO_LONG",
		8:  "SERVER_ERROR",
		9:  "CMYK_JPEG_NOT_ALLOWED",
		10: "FLASH_NOT_ALLOWED",
		11: "FLASH_WITHOUT_CLICKTAG",
		12: "FLASH_ERROR_AFTER_FIXING_CLICK_TAG",
		13: "ANIMATED_VISUAL_EFFECT",
		14: "FLASH_ERROR",
		15: "LAYOUT_PROBLEM",
		16: "PROBLEM_READING_IMAGE_FILE",
		17: "ERROR_STORING_IMAGE",
		18: "ASPECT_RATIO_NOT_ALLOWED",
		19: "FLASH_HAS_NETWORK_OBJECTS",
		20: "FLASH_HAS_NETWORK_METHODS",
		21: "FLASH_HAS_URL",
		22: "FLASH_HAS_MOUSE_TRACKING",
		23: "FLASH_HAS_RANDOM_NUM",
		24: "FLASH_SELF_TARGETS",
		25: "FLASH_BAD_GETURL_TARGET",
		26: "FLASH_VERSION_NOT_SUPPORTED",
		27: "FLASH_WITHOUT_HARD_CODED_CLICK_URL",
		28: "INVALID_FLASH_FILE",
		29: "FAILED_TO_FIX_CLICK_TAG_IN_FLASH",
		30: "FLASH_ACCESSES_NETWORK_RESOURCES",
		31: "FLASH_EXTERNAL_JS_CALL",
		32: "FLASH_EXTERNAL_FS_CALL",
		33: "FILE_TOO_LARGE",
		34: "IMAGE_DATA_TOO_LARGE",
		35: "IMAGE_PROCESSING_ERROR",
		36: "IMAGE_TOO_SMALL",
		37: "INVALID_INPUT",
		38: "PROBLEM_READING_FILE",
		39: "IMAGE_CONSTRAINTS_VIOLATED",
		40: "FORMAT_NOT_ALLOWED",
	}
	ImageErrorEnum_ImageError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"INVALID_IMAGE":                      2,
		"STORAGE_ERROR":                      3,
		"BAD_REQUEST":                        4,
		"UNEXPECTED_SIZE":                    5,
		"ANIMATED_NOT_ALLOWED":               6,
		"ANIMATION_TOO_LONG":                 7,
		"SERVER_ERROR":                       8,
		"CMYK_JPEG_NOT_ALLOWED":              9,
		"FLASH_NOT_ALLOWED":                  10,
		"FLASH_WITHOUT_CLICKTAG":             11,
		"FLASH_ERROR_AFTER_FIXING_CLICK_TAG": 12,
		"ANIMATED_VISUAL_EFFECT":             13,
		"FLASH_ERROR":                        14,
		"LAYOUT_PROBLEM":                     15,
		"PROBLEM_READING_IMAGE_FILE":         16,
		"ERROR_STORING_IMAGE":                17,
		"ASPECT_RATIO_NOT_ALLOWED":           18,
		"FLASH_HAS_NETWORK_OBJECTS":          19,
		"FLASH_HAS_NETWORK_METHODS":          20,
		"FLASH_HAS_URL":                      21,
		"FLASH_HAS_MOUSE_TRACKING":           22,
		"FLASH_HAS_RANDOM_NUM":               23,
		"FLASH_SELF_TARGETS":                 24,
		"FLASH_BAD_GETURL_TARGET":            25,
		"FLASH_VERSION_NOT_SUPPORTED":        26,
		"FLASH_WITHOUT_HARD_CODED_CLICK_URL": 27,
		"INVALID_FLASH_FILE":                 28,
		"FAILED_TO_FIX_CLICK_TAG_IN_FLASH":   29,
		"FLASH_ACCESSES_NETWORK_RESOURCES":   30,
		"FLASH_EXTERNAL_JS_CALL":             31,
		"FLASH_EXTERNAL_FS_CALL":             32,
		"FILE_TOO_LARGE":                     33,
		"IMAGE_DATA_TOO_LARGE":               34,
		"IMAGE_PROCESSING_ERROR":             35,
		"IMAGE_TOO_SMALL":                    36,
		"INVALID_INPUT":                      37,
		"PROBLEM_READING_FILE":               38,
		"IMAGE_CONSTRAINTS_VIOLATED":         39,
		"FORMAT_NOT_ALLOWED":                 40,
	}
)

func (x ImageErrorEnum_ImageError) Enum() *ImageErrorEnum_ImageError {
	p := new(ImageErrorEnum_ImageError)
	*p = x
	return p
}

func (x ImageErrorEnum_ImageError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageErrorEnum_ImageError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_image_error_proto_enumTypes[0].Descriptor()
}

func (ImageErrorEnum_ImageError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_image_error_proto_enumTypes[0]
}

func (x ImageErrorEnum_ImageError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageErrorEnum_ImageError.Descriptor instead.
func (ImageErrorEnum_ImageError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_image_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible image errors.
type ImageErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageErrorEnum) Reset() {
	*x = ImageErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_image_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageErrorEnum) ProtoMessage() {}

func (x *ImageErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_image_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageErrorEnum.ProtoReflect.Descriptor instead.
func (*ImageErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_image_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_image_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_image_error_proto_rawDesc = "" +
	"\n" +
	"1google/ads/googleads/v20/errors/image_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xc2\b\n" +
	"\x0eImageErrorEnum\"\xaf\b\n" +
	"\n" +
	"ImageError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x11\n" +
	"\rINVALID_IMAGE\x10\x02\x12\x11\n" +
	"\rSTORAGE_ERROR\x10\x03\x12\x0f\n" +
	"\vBAD_REQUEST\x10\x04\x12\x13\n" +
	"\x0fUNEXPECTED_SIZE\x10\x05\x12\x18\n" +
	"\x14ANIMATED_NOT_ALLOWED\x10\x06\x12\x16\n" +
	"\x12ANIMATION_TOO_LONG\x10\a\x12\x10\n" +
	"\fSERVER_ERROR\x10\b\x12\x19\n" +
	"\x15CMYK_JPEG_NOT_ALLOWED\x10\t\x12\x15\n" +
	"\x11FLASH_NOT_ALLOWED\x10\n" +
	"\x12\x1a\n" +
	"\x16FLASH_WITHOUT_CLICKTAG\x10\v\x12&\n" +
	"\"FLASH_ERROR_AFTER_FIXING_CLICK_TAG\x10\f\x12\x1a\n" +
	"\x16ANIMATED_VISUAL_EFFECT\x10\r\x12\x0f\n" +
	"\vFLASH_ERROR\x10\x0e\x12\x12\n" +
	"\x0eLAYOUT_PROBLEM\x10\x0f\x12\x1e\n" +
	"\x1aPROBLEM_READING_IMAGE_FILE\x10\x10\x12\x17\n" +
	"\x13ERROR_STORING_IMAGE\x10\x11\x12\x1c\n" +
	"\x18ASPECT_RATIO_NOT_ALLOWED\x10\x12\x12\x1d\n" +
	"\x19FLASH_HAS_NETWORK_OBJECTS\x10\x13\x12\x1d\n" +
	"\x19FLASH_HAS_NETWORK_METHODS\x10\x14\x12\x11\n" +
	"\rFLASH_HAS_URL\x10\x15\x12\x1c\n" +
	"\x18FLASH_HAS_MOUSE_TRACKING\x10\x16\x12\x18\n" +
	"\x14FLASH_HAS_RANDOM_NUM\x10\x17\x12\x16\n" +
	"\x12FLASH_SELF_TARGETS\x10\x18\x12\x1b\n" +
	"\x17FLASH_BAD_GETURL_TARGET\x10\x19\x12\x1f\n" +
	"\x1bFLASH_VERSION_NOT_SUPPORTED\x10\x1a\x12&\n" +
	"\"FLASH_WITHOUT_HARD_CODED_CLICK_URL\x10\x1b\x12\x16\n" +
	"\x12INVALID_FLASH_FILE\x10\x1c\x12$\n" +
	" FAILED_TO_FIX_CLICK_TAG_IN_FLASH\x10\x1d\x12$\n" +
	" FLASH_ACCESSES_NETWORK_RESOURCES\x10\x1e\x12\x1a\n" +
	"\x16FLASH_EXTERNAL_JS_CALL\x10\x1f\x12\x1a\n" +
	"\x16FLASH_EXTERNAL_FS_CALL\x10 \x12\x12\n" +
	"\x0eFILE_TOO_LARGE\x10!\x12\x18\n" +
	"\x14IMAGE_DATA_TOO_LARGE\x10\"\x12\x1a\n" +
	"\x16IMAGE_PROCESSING_ERROR\x10#\x12\x13\n" +
	"\x0fIMAGE_TOO_SMALL\x10$\x12\x11\n" +
	"\rINVALID_INPUT\x10%\x12\x18\n" +
	"\x14PROBLEM_READING_FILE\x10&\x12\x1e\n" +
	"\x1aIMAGE_CONSTRAINTS_VIOLATED\x10'\x12\x16\n" +
	"\x12FORMAT_NOT_ALLOWED\x10(B\xef\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x0fImageErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_image_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_image_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_image_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_image_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_image_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_image_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_image_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_image_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_image_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_image_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_image_error_proto_goTypes = []any{
	(ImageErrorEnum_ImageError)(0), // 0: google.ads.googleads.v20.errors.ImageErrorEnum.ImageError
	(*ImageErrorEnum)(nil),         // 1: google.ads.googleads.v20.errors.ImageErrorEnum
}
var file_google_ads_googleads_v20_errors_image_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_image_error_proto_init() }
func file_google_ads_googleads_v20_errors_image_error_proto_init() {
	if File_google_ads_googleads_v20_errors_image_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_image_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_image_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_image_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_image_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_image_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_image_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_image_error_proto = out.File
	file_google_ads_googleads_v20_errors_image_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_image_error_proto_depIdxs = nil
}
