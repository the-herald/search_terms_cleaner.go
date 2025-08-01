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
// source: google/ads/googleads/v20/resources/data_link.proto

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

// Represents the data sharing connection between  a Google
// Ads customer and another product's data.
type DataLink struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. Resource name of the product data link.
	// DataLink resource names have the form:
	//
	// `customers/{customer_id}/datalinks/{product_link_id}~{data_link_id}}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the link.
	// This field is read only.
	ProductLinkId *int64 `protobuf:"varint,2,opt,name=product_link_id,json=productLinkId,proto3,oneof" json:"product_link_id,omitempty"`
	// Output only. The ID of the data link.
	// This field is read only.
	DataLinkId *int64 `protobuf:"varint,3,opt,name=data_link_id,json=dataLinkId,proto3,oneof" json:"data_link_id,omitempty"`
	// Output only. The type of the data.
	Type enums.DataLinkTypeEnum_DataLinkType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v20.enums.DataLinkTypeEnum_DataLinkType" json:"type,omitempty"`
	// Output only. The status of the data link.
	Status enums.DataLinkStatusEnum_DataLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.DataLinkStatusEnum_DataLinkStatus" json:"status,omitempty"`
	// Data linked to this account.
	//
	// Types that are valid to be assigned to DataLinkEntity:
	//
	//	*DataLink_YoutubeVideo
	DataLinkEntity isDataLink_DataLinkEntity `protobuf_oneof:"data_link_entity"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DataLink) Reset() {
	*x = DataLink{}
	mi := &file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataLink) ProtoMessage() {}

func (x *DataLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataLink.ProtoReflect.Descriptor instead.
func (*DataLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_data_link_proto_rawDescGZIP(), []int{0}
}

func (x *DataLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DataLink) GetProductLinkId() int64 {
	if x != nil && x.ProductLinkId != nil {
		return *x.ProductLinkId
	}
	return 0
}

func (x *DataLink) GetDataLinkId() int64 {
	if x != nil && x.DataLinkId != nil {
		return *x.DataLinkId
	}
	return 0
}

func (x *DataLink) GetType() enums.DataLinkTypeEnum_DataLinkType {
	if x != nil {
		return x.Type
	}
	return enums.DataLinkTypeEnum_DataLinkType(0)
}

func (x *DataLink) GetStatus() enums.DataLinkStatusEnum_DataLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.DataLinkStatusEnum_DataLinkStatus(0)
}

func (x *DataLink) GetDataLinkEntity() isDataLink_DataLinkEntity {
	if x != nil {
		return x.DataLinkEntity
	}
	return nil
}

func (x *DataLink) GetYoutubeVideo() *YoutubeVideoIdentifier {
	if x != nil {
		if x, ok := x.DataLinkEntity.(*DataLink_YoutubeVideo); ok {
			return x.YoutubeVideo
		}
	}
	return nil
}

type isDataLink_DataLinkEntity interface {
	isDataLink_DataLinkEntity()
}

type DataLink_YoutubeVideo struct {
	// Immutable. A data link to YouTube video.
	YoutubeVideo *YoutubeVideoIdentifier `protobuf:"bytes,6,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

func (*DataLink_YoutubeVideo) isDataLink_DataLinkEntity() {}

// The identifier for YouTube video
type YoutubeVideoIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The ID of the hosting channel of the video. This is a string
	// value with “UC” prefix. For example, "UCK8sQmJBp8GCxrOtXWBpyEA".
	ChannelId *string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3,oneof" json:"channel_id,omitempty"`
	// Immutable. The ID of the video associated with the video link. This is the
	// 11 character string value used in the YouTube video URL. For example, video
	// ID is jV1vkHv4zq8 from the YouTube video URL
	// "https://www.youtube.com/watch?v=jV1vkHv4zq8&t=2s".
	VideoId       *string `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3,oneof" json:"video_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YoutubeVideoIdentifier) Reset() {
	*x = YoutubeVideoIdentifier{}
	mi := &file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YoutubeVideoIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YoutubeVideoIdentifier) ProtoMessage() {}

func (x *YoutubeVideoIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YoutubeVideoIdentifier.ProtoReflect.Descriptor instead.
func (*YoutubeVideoIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_data_link_proto_rawDescGZIP(), []int{1}
}

func (x *YoutubeVideoIdentifier) GetChannelId() string {
	if x != nil && x.ChannelId != nil {
		return *x.ChannelId
	}
	return ""
}

func (x *YoutubeVideoIdentifier) GetVideoId() string {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return ""
}

var File_google_ads_googleads_v20_resources_data_link_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_data_link_proto_rawDesc = "" +
	"\n" +
	"2google/ads/googleads/v20/resources/data_link.proto\x12\"google.ads.googleads.v20.resources\x1a5google/ads/googleads/v20/enums/data_link_status.proto\x1a3google/ads/googleads/v20/enums/data_link_type.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xfd\x04\n" +
	"\bDataLink\x12N\n" +
	"\rresource_name\x18\x01 \x01(\tB)\xe0A\x05\xfaA#\n" +
	"!googleads.googleapis.com/DataLinkR\fresourceName\x120\n" +
	"\x0fproduct_link_id\x18\x02 \x01(\x03B\x03\xe0A\x03H\x01R\rproductLinkId\x88\x01\x01\x12*\n" +
	"\fdata_link_id\x18\x03 \x01(\x03B\x03\xe0A\x03H\x02R\n" +
	"dataLinkId\x88\x01\x01\x12V\n" +
	"\x04type\x18\x04 \x01(\x0e2=.google.ads.googleads.v20.enums.DataLinkTypeEnum.DataLinkTypeB\x03\xe0A\x03R\x04type\x12^\n" +
	"\x06status\x18\x05 \x01(\x0e2A.google.ads.googleads.v20.enums.DataLinkStatusEnum.DataLinkStatusB\x03\xe0A\x03R\x06status\x12f\n" +
	"\ryoutube_video\x18\x06 \x01(\v2:.google.ads.googleads.v20.resources.YoutubeVideoIdentifierB\x03\xe0A\x05H\x00R\fyoutubeVideo:j\xeaAg\n" +
	"!googleads.googleapis.com/DataLink\x12Bcustomers/{customer_id}/dataLinks/{product_link_id}~{data_link_id}B\x12\n" +
	"\x10data_link_entityB\x12\n" +
	"\x10_product_link_idB\x0f\n" +
	"\r_data_link_id\"\x82\x01\n" +
	"\x16YoutubeVideoIdentifier\x12'\n" +
	"\n" +
	"channel_id\x18\x01 \x01(\tB\x03\xe0A\x05H\x00R\tchannelId\x88\x01\x01\x12#\n" +
	"\bvideo_id\x18\x02 \x01(\tB\x03\xe0A\x05H\x01R\avideoId\x88\x01\x01B\r\n" +
	"\v_channel_idB\v\n" +
	"\t_video_idB\xff\x01\n" +
	"&com.google.ads.googleads.v20.resourcesB\rDataLinkProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_data_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_data_link_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_data_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_data_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_data_link_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_data_link_proto_rawDesc), len(file_google_ads_googleads_v20_resources_data_link_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_data_link_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_data_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v20_resources_data_link_proto_goTypes = []any{
	(*DataLink)(nil),                             // 0: google.ads.googleads.v20.resources.DataLink
	(*YoutubeVideoIdentifier)(nil),               // 1: google.ads.googleads.v20.resources.YoutubeVideoIdentifier
	(enums.DataLinkTypeEnum_DataLinkType)(0),     // 2: google.ads.googleads.v20.enums.DataLinkTypeEnum.DataLinkType
	(enums.DataLinkStatusEnum_DataLinkStatus)(0), // 3: google.ads.googleads.v20.enums.DataLinkStatusEnum.DataLinkStatus
}
var file_google_ads_googleads_v20_resources_data_link_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v20.resources.DataLink.type:type_name -> google.ads.googleads.v20.enums.DataLinkTypeEnum.DataLinkType
	3, // 1: google.ads.googleads.v20.resources.DataLink.status:type_name -> google.ads.googleads.v20.enums.DataLinkStatusEnum.DataLinkStatus
	1, // 2: google.ads.googleads.v20.resources.DataLink.youtube_video:type_name -> google.ads.googleads.v20.resources.YoutubeVideoIdentifier
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_data_link_proto_init() }
func file_google_ads_googleads_v20_resources_data_link_proto_init() {
	if File_google_ads_googleads_v20_resources_data_link_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[0].OneofWrappers = []any{
		(*DataLink_YoutubeVideo)(nil),
	}
	file_google_ads_googleads_v20_resources_data_link_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_data_link_proto_rawDesc), len(file_google_ads_googleads_v20_resources_data_link_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_data_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_data_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_data_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_data_link_proto = out.File
	file_google_ads_googleads_v20_resources_data_link_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_data_link_proto_depIdxs = nil
}
