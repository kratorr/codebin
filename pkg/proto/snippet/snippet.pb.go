// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: snippet.proto

package snippet

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Snippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Language int64  `protobuf:"varint,2,opt,name=language,proto3" json:"language,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Snippet) Reset() {
	*x = Snippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snippet) ProtoMessage() {}

func (x *Snippet) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snippet.ProtoReflect.Descriptor instead.
func (*Snippet) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{0}
}

func (x *Snippet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Snippet) GetLanguage() int64 {
	if x != nil {
		return x.Language
	}
	return 0
}

func (x *Snippet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snippet *Snippet `protobuf:"bytes,1,opt,name=snippet,proto3" json:"snippet,omitempty"`
}

func (x *CreateSnippetRequest) Reset() {
	*x = CreateSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetRequest) ProtoMessage() {}

func (x *CreateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetRequest.ProtoReflect.Descriptor instead.
func (*CreateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnippetRequest) GetSnippet() *Snippet {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type CreateSnippetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateSnippetResponse) Reset() {
	*x = CreateSnippetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetResponse) ProtoMessage() {}

func (x *CreateSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetResponse.ProtoReflect.Descriptor instead.
func (*CreateSnippetResponse) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSnippetResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type DeleteSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl *Snippet `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *DeleteSnippetRequest) Reset() {
	*x = DeleteSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnippetRequest) ProtoMessage() {}

func (x *DeleteSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnippetRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnippetRequest) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSnippetRequest) GetShortUrl() *Snippet {
	if x != nil {
		return x.ShortUrl
	}
	return nil
}

type ListSnippetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snippets []*Snippet `protobuf:"bytes,1,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *ListSnippetResponse) Reset() {
	*x = ListSnippetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnippetResponse) ProtoMessage() {}

func (x *ListSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnippetResponse.ProtoReflect.Descriptor instead.
func (*ListSnippetResponse) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{4}
}

func (x *ListSnippetResponse) GetSnippets() []*Snippet {
	if x != nil {
		return x.Snippets
	}
	return nil
}

type ReadLatestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadLatestRequest) Reset() {
	*x = ReadLatestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLatestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLatestRequest) ProtoMessage() {}

func (x *ReadLatestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLatestRequest.ProtoReflect.Descriptor instead.
func (*ReadLatestRequest) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{5}
}

type ReadLatestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snippets []*Snippet `protobuf:"bytes,1,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *ReadLatestResponse) Reset() {
	*x = ReadLatestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snippet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadLatestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadLatestResponse) ProtoMessage() {}

func (x *ReadLatestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snippet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadLatestResponse.ProtoReflect.Descriptor instead.
func (*ReadLatestResponse) Descriptor() ([]byte, []int) {
	return file_snippet_proto_rawDescGZIP(), []int{6}
}

func (x *ReadLatestResponse) GetSnippets() []*Snippet {
	if x != nil {
		return x.Snippets
	}
	return nil
}

var File_snippet_proto protoreflect.FileDescriptor

var file_snippet_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x07, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x07,
	0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x22, 0x34, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3d, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3b, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52,
	0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x61,
	0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a,
	0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x32, 0xfa, 0x02, 0x0a, 0x09, 0x53,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x56, 0x31, 0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x7a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x41, 0x92, 0x41, 0x2b, 0x62, 0x29, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69,
	0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x17, 0x0a, 0x06, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x32, 0x12, 0x0d, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x05, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69,
	0x70, 0x70, 0x65, 0x74, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x42, 0x93, 0x02, 0x92, 0x41, 0xd9, 0x01, 0x12, 0x36,
	0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x62, 0x69, 0x6e, 0x20, 0x41, 0x50, 0x49, 0x22, 0x1f, 0x0a,
	0x07, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x72, 0x1a, 0x14, 0x73, 0x6d, 0x69, 0x72, 0x6e, 0x6f,
	0x76, 0x76, 0x39, 0x31, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x06,
	0x30, 0x2e, 0x30, 0x2e, 0x31, 0x61, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x39, 0x30, 0x30, 0x31, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x59,
	0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f,
	0x76, 0x31, 0x3b, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snippet_proto_rawDescOnce sync.Once
	file_snippet_proto_rawDescData = file_snippet_proto_rawDesc
)

func file_snippet_proto_rawDescGZIP() []byte {
	file_snippet_proto_rawDescOnce.Do(func() {
		file_snippet_proto_rawDescData = protoimpl.X.CompressGZIP(file_snippet_proto_rawDescData)
	})
	return file_snippet_proto_rawDescData
}

var file_snippet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_snippet_proto_goTypes = []interface{}{
	(*Snippet)(nil),               // 0: Snippet
	(*CreateSnippetRequest)(nil),  // 1: CreateSnippetRequest
	(*CreateSnippetResponse)(nil), // 2: CreateSnippetResponse
	(*DeleteSnippetRequest)(nil),  // 3: DeleteSnippetRequest
	(*ListSnippetResponse)(nil),   // 4: ListSnippetResponse
	(*ReadLatestRequest)(nil),     // 5: ReadLatestRequest
	(*ReadLatestResponse)(nil),    // 6: ReadLatestResponse
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_snippet_proto_depIdxs = []int32{
	0, // 0: CreateSnippetRequest.snippet:type_name -> Snippet
	0, // 1: DeleteSnippetRequest.short_url:type_name -> Snippet
	0, // 2: ListSnippetResponse.snippets:type_name -> Snippet
	0, // 3: ReadLatestResponse.snippets:type_name -> Snippet
	1, // 4: SnippetV1.Create:input_type -> CreateSnippetRequest
	3, // 5: SnippetV1.Delete:input_type -> DeleteSnippetRequest
	5, // 6: SnippetV1.ReadLatest:input_type -> ReadLatestRequest
	7, // 7: SnippetV1.List:input_type -> google.protobuf.Empty
	2, // 8: SnippetV1.Create:output_type -> CreateSnippetResponse
	7, // 9: SnippetV1.Delete:output_type -> google.protobuf.Empty
	6, // 10: SnippetV1.ReadLatest:output_type -> ReadLatestResponse
	4, // 11: SnippetV1.List:output_type -> ListSnippetResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_snippet_proto_init() }
func file_snippet_proto_init() {
	if File_snippet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snippet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snippet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnippetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnippetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnippetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnippetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLatestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snippet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadLatestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_snippet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snippet_proto_goTypes,
		DependencyIndexes: file_snippet_proto_depIdxs,
		MessageInfos:      file_snippet_proto_msgTypes,
	}.Build()
	File_snippet_proto = out.File
	file_snippet_proto_rawDesc = nil
	file_snippet_proto_goTypes = nil
	file_snippet_proto_depIdxs = nil
}
