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
// source: google/api/backend.proto

package serviceconfig

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

// Path Translation specifies how to combine the backend address with the
// request path in order to produce the appropriate forwarding URL for the
// request.
//
// Path Translation is applicable only to HTTP-based backends. Backends which
// do not accept requests over HTTP/HTTPS should leave `path_translation`
// unspecified.
type BackendRule_PathTranslation int32

const (
	BackendRule_PATH_TRANSLATION_UNSPECIFIED BackendRule_PathTranslation = 0
	// Use the backend address as-is, with no modification to the path. If the
	// URL pattern contains variables, the variable names and values will be
	// appended to the query string. If a query string parameter and a URL
	// pattern variable have the same name, this may result in duplicate keys in
	// the query string.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//     Method path:        /api/company/{cid}/user/{uid}
	//     Backend address:    https://example.cloudfunctions.net/getUser
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//     Request path: /api/company/widgetworks/user/johndoe
	//     Translated:
	//     https://example.cloudfunctions.net/getUser?cid=widgetworks&uid=johndoe
	//
	//     Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//     Translated:
	//     https://example.cloudfunctions.net/getUser?timezone=EST&cid=widgetworks&uid=johndoe
	BackendRule_CONSTANT_ADDRESS BackendRule_PathTranslation = 1
	// The request path will be appended to the backend address.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//     Method path:        /api/company/{cid}/user/{uid}
	//     Backend address:    https://example.appspot.com
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//     Request path: /api/company/widgetworks/user/johndoe
	//     Translated:
	//     https://example.appspot.com/api/company/widgetworks/user/johndoe
	//
	//     Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//     Translated:
	//     https://example.appspot.com/api/company/widgetworks/user/johndoe?timezone=EST
	BackendRule_APPEND_PATH_TO_ADDRESS BackendRule_PathTranslation = 2
)

// Enum value maps for BackendRule_PathTranslation.
var (
	BackendRule_PathTranslation_name = map[int32]string{
		0: "PATH_TRANSLATION_UNSPECIFIED",
		1: "CONSTANT_ADDRESS",
		2: "APPEND_PATH_TO_ADDRESS",
	}
	BackendRule_PathTranslation_value = map[string]int32{
		"PATH_TRANSLATION_UNSPECIFIED": 0,
		"CONSTANT_ADDRESS":             1,
		"APPEND_PATH_TO_ADDRESS":       2,
	}
)

func (x BackendRule_PathTranslation) Enum() *BackendRule_PathTranslation {
	p := new(BackendRule_PathTranslation)
	*p = x
	return p
}

func (x BackendRule_PathTranslation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackendRule_PathTranslation) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_backend_proto_enumTypes[0].Descriptor()
}

func (BackendRule_PathTranslation) Type() protoreflect.EnumType {
	return &file_google_api_backend_proto_enumTypes[0]
}

func (x BackendRule_PathTranslation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackendRule_PathTranslation.Descriptor instead.
func (BackendRule_PathTranslation) EnumDescriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{1, 0}
}

// `Backend` defines the backend configuration for a service.
type Backend struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A list of API backend rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules         []*BackendRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Backend) Reset() {
	*x = Backend{}
	mi := &file_google_api_backend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backend) ProtoMessage() {}

func (x *Backend) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backend.ProtoReflect.Descriptor instead.
func (*Backend) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Backend) GetRules() []*BackendRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax
	// details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The address of the API backend.
	//
	// The scheme is used to determine the backend protocol and security.
	// The following schemes are accepted:
	//
	//    SCHEME        PROTOCOL    SECURITY
	//    http://       HTTP        None
	//    https://      HTTP        TLS
	//    grpc://       gRPC        None
	//    grpcs://      gRPC        TLS
	//
	// It is recommended to explicitly include a scheme. Leaving out the scheme
	// may cause constrasting behaviors across platforms.
	//
	// If the port is unspecified, the default is:
	// - 80 for schemes without TLS
	// - 443 for schemes with TLS
	//
	// For HTTP backends, use [protocol][google.api.BackendRule.protocol]
	// to specify the protocol version.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request. The default
	// varies based on the request protocol and deployment environment.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Deprecated, do not use.
	//
	// Deprecated: Marked as deprecated in google/api/backend.proto.
	MinDeadline float64 `protobuf:"fixed64,4,opt,name=min_deadline,json=minDeadline,proto3" json:"min_deadline,omitempty"`
	// The number of seconds to wait for the completion of a long running
	// operation. The default is no deadline.
	OperationDeadline float64                     `protobuf:"fixed64,5,opt,name=operation_deadline,json=operationDeadline,proto3" json:"operation_deadline,omitempty"`
	PathTranslation   BackendRule_PathTranslation `protobuf:"varint,6,opt,name=path_translation,json=pathTranslation,proto3,enum=google.api.BackendRule_PathTranslation" json:"path_translation,omitempty"`
	// Authentication settings used by the backend.
	//
	// These are typically used to provide service management functionality to
	// a backend served on a publicly-routable URL. The `authentication`
	// details should match the authentication behavior used by the backend.
	//
	// For example, specifying `jwt_audience` implies that the backend expects
	// authentication via a JWT.
	//
	// When authentication is unspecified, the resulting behavior is the same
	// as `disable_auth` set to `true`.
	//
	// Refer to https://developers.google.com/identity/protocols/OpenIDConnect for
	// JWT ID token.
	//
	// Types that are valid to be assigned to Authentication:
	//
	//	*BackendRule_JwtAudience
	//	*BackendRule_DisableAuth
	Authentication isBackendRule_Authentication `protobuf_oneof:"authentication"`
	// The protocol used for sending a request to the backend.
	// The supported values are "http/1.1" and "h2".
	//
	// The default value is inferred from the scheme in the
	// [address][google.api.BackendRule.address] field:
	//
	//    SCHEME        PROTOCOL
	//    http://       http/1.1
	//    https://      http/1.1
	//    grpc://       h2
	//    grpcs://      h2
	//
	// For secure HTTP backends (https://) that support HTTP/2, set this field
	// to "h2" for improved performance.
	//
	// Configuring this field to non-default values is only supported for secure
	// HTTP backends. This field will be ignored for all other backends.
	//
	// See
	// https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids
	// for more details on the supported values.
	Protocol string `protobuf:"bytes,9,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The map between request protocol and the backend address.
	OverridesByRequestProtocol map[string]*BackendRule `protobuf:"bytes,10,rep,name=overrides_by_request_protocol,json=overridesByRequestProtocol,proto3" json:"overrides_by_request_protocol,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *BackendRule) Reset() {
	*x = BackendRule{}
	mi := &file_google_api_backend_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackendRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRule) ProtoMessage() {}

func (x *BackendRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRule.ProtoReflect.Descriptor instead.
func (*BackendRule) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{1}
}

func (x *BackendRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *BackendRule) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BackendRule) GetDeadline() float64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

// Deprecated: Marked as deprecated in google/api/backend.proto.
func (x *BackendRule) GetMinDeadline() float64 {
	if x != nil {
		return x.MinDeadline
	}
	return 0
}

func (x *BackendRule) GetOperationDeadline() float64 {
	if x != nil {
		return x.OperationDeadline
	}
	return 0
}

func (x *BackendRule) GetPathTranslation() BackendRule_PathTranslation {
	if x != nil {
		return x.PathTranslation
	}
	return BackendRule_PATH_TRANSLATION_UNSPECIFIED
}

func (x *BackendRule) GetAuthentication() isBackendRule_Authentication {
	if x != nil {
		return x.Authentication
	}
	return nil
}

func (x *BackendRule) GetJwtAudience() string {
	if x != nil {
		if x, ok := x.Authentication.(*BackendRule_JwtAudience); ok {
			return x.JwtAudience
		}
	}
	return ""
}

func (x *BackendRule) GetDisableAuth() bool {
	if x != nil {
		if x, ok := x.Authentication.(*BackendRule_DisableAuth); ok {
			return x.DisableAuth
		}
	}
	return false
}

func (x *BackendRule) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *BackendRule) GetOverridesByRequestProtocol() map[string]*BackendRule {
	if x != nil {
		return x.OverridesByRequestProtocol
	}
	return nil
}

type isBackendRule_Authentication interface {
	isBackendRule_Authentication()
}

type BackendRule_JwtAudience struct {
	// The JWT audience is used when generating a JWT ID token for the backend.
	// This ID token will be added in the HTTP "authorization" header, and sent
	// to the backend.
	JwtAudience string `protobuf:"bytes,7,opt,name=jwt_audience,json=jwtAudience,proto3,oneof"`
}

type BackendRule_DisableAuth struct {
	// When disable_auth is true, a JWT ID token won't be generated and the
	// original "Authorization" HTTP header will be preserved. If the header is
	// used to carry the original token and is expected by the backend, this
	// field must be set to true to preserve the header.
	DisableAuth bool `protobuf:"varint,8,opt,name=disable_auth,json=disableAuth,proto3,oneof"`
}

func (*BackendRule_JwtAudience) isBackendRule_Authentication() {}

func (*BackendRule_DisableAuth) isBackendRule_Authentication() {}

var File_google_api_backend_proto protoreflect.FileDescriptor

const file_google_api_backend_proto_rawDesc = "" +
	"\n" +
	"\x18google/api/backend.proto\x12\n" +
	"google.api\"8\n" +
	"\aBackend\x12-\n" +
	"\x05rules\x18\x01 \x03(\v2\x17.google.api.BackendRuleR\x05rules\"\xcc\x05\n" +
	"\vBackendRule\x12\x1a\n" +
	"\bselector\x18\x01 \x01(\tR\bselector\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x1a\n" +
	"\bdeadline\x18\x03 \x01(\x01R\bdeadline\x12%\n" +
	"\fmin_deadline\x18\x04 \x01(\x01B\x02\x18\x01R\vminDeadline\x12-\n" +
	"\x12operation_deadline\x18\x05 \x01(\x01R\x11operationDeadline\x12R\n" +
	"\x10path_translation\x18\x06 \x01(\x0e2'.google.api.BackendRule.PathTranslationR\x0fpathTranslation\x12#\n" +
	"\fjwt_audience\x18\a \x01(\tH\x00R\vjwtAudience\x12#\n" +
	"\fdisable_auth\x18\b \x01(\bH\x00R\vdisableAuth\x12\x1a\n" +
	"\bprotocol\x18\t \x01(\tR\bprotocol\x12z\n" +
	"\x1doverrides_by_request_protocol\x18\n" +
	" \x03(\v27.google.api.BackendRule.OverridesByRequestProtocolEntryR\x1aoverridesByRequestProtocol\x1af\n" +
	"\x1fOverridesByRequestProtocolEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12-\n" +
	"\x05value\x18\x02 \x01(\v2\x17.google.api.BackendRuleR\x05value:\x028\x01\"e\n" +
	"\x0fPathTranslation\x12 \n" +
	"\x1cPATH_TRANSLATION_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10CONSTANT_ADDRESS\x10\x01\x12\x1a\n" +
	"\x16APPEND_PATH_TO_ADDRESS\x10\x02B\x10\n" +
	"\x0eauthenticationBn\n" +
	"\x0ecom.google.apiB\fBackendProtoP\x01ZEgoogle.golang.org/genproto/googleapis/api/serviceconfig;serviceconfig\xa2\x02\x04GAPIb\x06proto3"

var (
	file_google_api_backend_proto_rawDescOnce sync.Once
	file_google_api_backend_proto_rawDescData []byte
)

func file_google_api_backend_proto_rawDescGZIP() []byte {
	file_google_api_backend_proto_rawDescOnce.Do(func() {
		file_google_api_backend_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_api_backend_proto_rawDesc), len(file_google_api_backend_proto_rawDesc)))
	})
	return file_google_api_backend_proto_rawDescData
}

var file_google_api_backend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_api_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_backend_proto_goTypes = []any{
	(BackendRule_PathTranslation)(0), // 0: google.api.BackendRule.PathTranslation
	(*Backend)(nil),                  // 1: google.api.Backend
	(*BackendRule)(nil),              // 2: google.api.BackendRule
	nil,                              // 3: google.api.BackendRule.OverridesByRequestProtocolEntry
}
var file_google_api_backend_proto_depIdxs = []int32{
	2, // 0: google.api.Backend.rules:type_name -> google.api.BackendRule
	0, // 1: google.api.BackendRule.path_translation:type_name -> google.api.BackendRule.PathTranslation
	3, // 2: google.api.BackendRule.overrides_by_request_protocol:type_name -> google.api.BackendRule.OverridesByRequestProtocolEntry
	2, // 3: google.api.BackendRule.OverridesByRequestProtocolEntry.value:type_name -> google.api.BackendRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_api_backend_proto_init() }
func file_google_api_backend_proto_init() {
	if File_google_api_backend_proto != nil {
		return
	}
	file_google_api_backend_proto_msgTypes[1].OneofWrappers = []any{
		(*BackendRule_JwtAudience)(nil),
		(*BackendRule_DisableAuth)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_api_backend_proto_rawDesc), len(file_google_api_backend_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_backend_proto_goTypes,
		DependencyIndexes: file_google_api_backend_proto_depIdxs,
		EnumInfos:         file_google_api_backend_proto_enumTypes,
		MessageInfos:      file_google_api_backend_proto_msgTypes,
	}.Build()
	File_google_api_backend_proto = out.File
	file_google_api_backend_proto_goTypes = nil
	file_google_api_backend_proto_depIdxs = nil
}
