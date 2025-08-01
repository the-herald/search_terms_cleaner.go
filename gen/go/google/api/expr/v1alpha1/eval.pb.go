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
// source: google/api/expr/v1alpha1/eval.proto

package expr

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// The state of an evaluation.
//
// Can represent an inital, partial, or completed state of evaluation.
type EvalState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique values referenced in this message.
	Values []*ExprValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// An ordered list of results.
	//
	// Tracks the flow of evaluation through the expression.
	// May be sparse.
	Results       []*EvalState_Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvalState) Reset() {
	*x = EvalState{}
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvalState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalState) ProtoMessage() {}

func (x *EvalState) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalState.ProtoReflect.Descriptor instead.
func (*EvalState) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP(), []int{0}
}

func (x *EvalState) GetValues() []*ExprValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *EvalState) GetResults() []*EvalState_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// The value of an evaluated expression.
type ExprValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An expression can resolve to a value, error or unknown.
	//
	// Types that are valid to be assigned to Kind:
	//
	//	*ExprValue_Value
	//	*ExprValue_Error
	//	*ExprValue_Unknown
	Kind          isExprValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExprValue) Reset() {
	*x = ExprValue{}
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExprValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExprValue) ProtoMessage() {}

func (x *ExprValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExprValue.ProtoReflect.Descriptor instead.
func (*ExprValue) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP(), []int{1}
}

func (x *ExprValue) GetKind() isExprValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *ExprValue) GetValue() *Value {
	if x != nil {
		if x, ok := x.Kind.(*ExprValue_Value); ok {
			return x.Value
		}
	}
	return nil
}

func (x *ExprValue) GetError() *ErrorSet {
	if x != nil {
		if x, ok := x.Kind.(*ExprValue_Error); ok {
			return x.Error
		}
	}
	return nil
}

func (x *ExprValue) GetUnknown() *UnknownSet {
	if x != nil {
		if x, ok := x.Kind.(*ExprValue_Unknown); ok {
			return x.Unknown
		}
	}
	return nil
}

type isExprValue_Kind interface {
	isExprValue_Kind()
}

type ExprValue_Value struct {
	// A concrete value.
	Value *Value `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}

type ExprValue_Error struct {
	// The set of errors in the critical path of evalution.
	//
	// Only errors in the critical path are included. For example,
	// `(<error1> || true) && <error2>` will only result in `<error2>`,
	// while `<error1> || <error2>` will result in both `<error1>` and
	// `<error2>`.
	//
	// Errors cause by the presence of other errors are not included in the
	// set. For example `<error1>.foo`, `foo(<error1>)`, and `<error1> + 1` will
	// only result in `<error1>`.
	//
	// Multiple errors *might* be included when evaluation could result
	// in different errors. For example `<error1> + <error2>` and
	// `foo(<error1>, <error2>)` may result in `<error1>`, `<error2>` or both.
	// The exact subset of errors included for this case is unspecified and
	// depends on the implementation details of the evaluator.
	Error *ErrorSet `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type ExprValue_Unknown struct {
	// The set of unknowns in the critical path of evaluation.
	//
	// Unknown behaves identically to Error with regards to propagation.
	// Specifically, only unknowns in the critical path are included, unknowns
	// caused by the presence of other unknowns are not included, and multiple
	// unknowns *might* be included included when evaluation could result in
	// different unknowns. For example:
	//
	//     (<unknown[1]> || true) && <unknown[2]> -> <unknown[2]>
	//     <unknown[1]> || <unknown[2]> -> <unknown[1,2]>
	//     <unknown[1]>.foo -> <unknown[1]>
	//     foo(<unknown[1]>) -> <unknown[1]>
	//     <unknown[1]> + <unknown[2]> -> <unknown[1]> or <unknown[2[>
	//
	// Unknown takes precidence over Error in cases where a `Value` can short
	// circuit the result:
	//
	//     <error> || <unknown> -> <unknown>
	//     <error> && <unknown> -> <unknown>
	//
	// Errors take precidence in all other cases:
	//
	//     <unknown> + <error> -> <error>
	//     foo(<unknown>, <error>) -> <error>
	Unknown *UnknownSet `protobuf:"bytes,3,opt,name=unknown,proto3,oneof"`
}

func (*ExprValue_Value) isExprValue_Kind() {}

func (*ExprValue_Error) isExprValue_Kind() {}

func (*ExprValue_Unknown) isExprValue_Kind() {}

// A set of errors.
//
// The errors included depend on the context. See `ExprValue.error`.
type ErrorSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The errors in the set.
	Errors        []*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorSet) Reset() {
	*x = ErrorSet{}
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorSet) ProtoMessage() {}

func (x *ErrorSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorSet.ProtoReflect.Descriptor instead.
func (*ErrorSet) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorSet) GetErrors() []*status.Status {
	if x != nil {
		return x.Errors
	}
	return nil
}

// A set of expressions for which the value is unknown.
//
// The unknowns included depend on the context. See `ExprValue.unknown`.
type UnknownSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ids of the expressions with unknown values.
	Exprs         []int64 `protobuf:"varint,1,rep,packed,name=exprs,proto3" json:"exprs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnknownSet) Reset() {
	*x = UnknownSet{}
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnknownSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownSet) ProtoMessage() {}

func (x *UnknownSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownSet.ProtoReflect.Descriptor instead.
func (*UnknownSet) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP(), []int{3}
}

func (x *UnknownSet) GetExprs() []int64 {
	if x != nil {
		return x.Exprs
	}
	return nil
}

// A single evalution result.
type EvalState_Result struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the expression this result if for.
	Expr int64 `protobuf:"varint,1,opt,name=expr,proto3" json:"expr,omitempty"`
	// The index in `values` of the resulting value.
	Value         int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvalState_Result) Reset() {
	*x = EvalState_Result{}
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvalState_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvalState_Result) ProtoMessage() {}

func (x *EvalState_Result) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_expr_v1alpha1_eval_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvalState_Result.ProtoReflect.Descriptor instead.
func (*EvalState_Result) Descriptor() ([]byte, []int) {
	return file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EvalState_Result) GetExpr() int64 {
	if x != nil {
		return x.Expr
	}
	return 0
}

func (x *EvalState_Result) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_google_api_expr_v1alpha1_eval_proto protoreflect.FileDescriptor

const file_google_api_expr_v1alpha1_eval_proto_rawDesc = "" +
	"\n" +
	"#google/api/expr/v1alpha1/eval.proto\x12\x18google.api.expr.v1alpha1\x1a$google/api/expr/v1alpha1/value.proto\x1a\x17google/rpc/status.proto\"\xc2\x01\n" +
	"\tEvalState\x12;\n" +
	"\x06values\x18\x01 \x03(\v2#.google.api.expr.v1alpha1.ExprValueR\x06values\x12D\n" +
	"\aresults\x18\x03 \x03(\v2*.google.api.expr.v1alpha1.EvalState.ResultR\aresults\x1a2\n" +
	"\x06Result\x12\x12\n" +
	"\x04expr\x18\x01 \x01(\x03R\x04expr\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value\"\xca\x01\n" +
	"\tExprValue\x127\n" +
	"\x05value\x18\x01 \x01(\v2\x1f.google.api.expr.v1alpha1.ValueH\x00R\x05value\x12:\n" +
	"\x05error\x18\x02 \x01(\v2\".google.api.expr.v1alpha1.ErrorSetH\x00R\x05error\x12@\n" +
	"\aunknown\x18\x03 \x01(\v2$.google.api.expr.v1alpha1.UnknownSetH\x00R\aunknownB\x06\n" +
	"\x04kind\"6\n" +
	"\bErrorSet\x12*\n" +
	"\x06errors\x18\x01 \x03(\v2\x12.google.rpc.StatusR\x06errors\"\"\n" +
	"\n" +
	"UnknownSet\x12\x14\n" +
	"\x05exprs\x18\x01 \x03(\x03R\x05exprsBl\n" +
	"\x1ccom.google.api.expr.v1alpha1B\tEvalProtoP\x01Z<google.golang.org/genproto/googleapis/api/expr/v1alpha1;expr\xf8\x01\x01b\x06proto3"

var (
	file_google_api_expr_v1alpha1_eval_proto_rawDescOnce sync.Once
	file_google_api_expr_v1alpha1_eval_proto_rawDescData []byte
)

func file_google_api_expr_v1alpha1_eval_proto_rawDescGZIP() []byte {
	file_google_api_expr_v1alpha1_eval_proto_rawDescOnce.Do(func() {
		file_google_api_expr_v1alpha1_eval_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_api_expr_v1alpha1_eval_proto_rawDesc), len(file_google_api_expr_v1alpha1_eval_proto_rawDesc)))
	})
	return file_google_api_expr_v1alpha1_eval_proto_rawDescData
}

var file_google_api_expr_v1alpha1_eval_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_api_expr_v1alpha1_eval_proto_goTypes = []any{
	(*EvalState)(nil),        // 0: google.api.expr.v1alpha1.EvalState
	(*ExprValue)(nil),        // 1: google.api.expr.v1alpha1.ExprValue
	(*ErrorSet)(nil),         // 2: google.api.expr.v1alpha1.ErrorSet
	(*UnknownSet)(nil),       // 3: google.api.expr.v1alpha1.UnknownSet
	(*EvalState_Result)(nil), // 4: google.api.expr.v1alpha1.EvalState.Result
	(*Value)(nil),            // 5: google.api.expr.v1alpha1.Value
	(*status.Status)(nil),    // 6: google.rpc.Status
}
var file_google_api_expr_v1alpha1_eval_proto_depIdxs = []int32{
	1, // 0: google.api.expr.v1alpha1.EvalState.values:type_name -> google.api.expr.v1alpha1.ExprValue
	4, // 1: google.api.expr.v1alpha1.EvalState.results:type_name -> google.api.expr.v1alpha1.EvalState.Result
	5, // 2: google.api.expr.v1alpha1.ExprValue.value:type_name -> google.api.expr.v1alpha1.Value
	2, // 3: google.api.expr.v1alpha1.ExprValue.error:type_name -> google.api.expr.v1alpha1.ErrorSet
	3, // 4: google.api.expr.v1alpha1.ExprValue.unknown:type_name -> google.api.expr.v1alpha1.UnknownSet
	6, // 5: google.api.expr.v1alpha1.ErrorSet.errors:type_name -> google.rpc.Status
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_api_expr_v1alpha1_eval_proto_init() }
func file_google_api_expr_v1alpha1_eval_proto_init() {
	if File_google_api_expr_v1alpha1_eval_proto != nil {
		return
	}
	file_google_api_expr_v1alpha1_value_proto_init()
	file_google_api_expr_v1alpha1_eval_proto_msgTypes[1].OneofWrappers = []any{
		(*ExprValue_Value)(nil),
		(*ExprValue_Error)(nil),
		(*ExprValue_Unknown)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_api_expr_v1alpha1_eval_proto_rawDesc), len(file_google_api_expr_v1alpha1_eval_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_expr_v1alpha1_eval_proto_goTypes,
		DependencyIndexes: file_google_api_expr_v1alpha1_eval_proto_depIdxs,
		MessageInfos:      file_google_api_expr_v1alpha1_eval_proto_msgTypes,
	}.Build()
	File_google_api_expr_v1alpha1_eval_proto = out.File
	file_google_api_expr_v1alpha1_eval_proto_goTypes = nil
	file_google_api_expr_v1alpha1_eval_proto_depIdxs = nil
}
