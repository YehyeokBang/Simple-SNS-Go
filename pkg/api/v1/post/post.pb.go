// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: pkg/api/v1/post/post.proto

package post

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName  string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type WritePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WritePostRequest) Reset() {
	*x = WritePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePostRequest) ProtoMessage() {}

func (x *WritePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePostRequest.ProtoReflect.Descriptor instead.
func (*WritePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *WritePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type WritePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *WritePostResponse) Reset() {
	*x = WritePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePostResponse) ProtoMessage() {}

func (x *WritePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePostResponse.ProtoReflect.Descriptor instead.
func (*WritePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *WritePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostsRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPostsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePostRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostResponse) Reset() {
	*x = UpdatePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResponse) ProtoMessage() {}

func (x *UpdatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePostRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_post_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_post_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_post_post_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePostResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_pkg_api_v1_post_post_proto protoreflect.FileDescriptor

var file_pkg_api_v1_post_post_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x31,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x10,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x11, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xa8, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x65, 0x68,
	0x79, 0x65, 0x6f, 0x6b, 0x42, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d,
	0x53, 0x4e, 0x53, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_v1_post_post_proto_rawDescOnce sync.Once
	file_pkg_api_v1_post_post_proto_rawDescData = file_pkg_api_v1_post_post_proto_rawDesc
)

func file_pkg_api_v1_post_post_proto_rawDescGZIP() []byte {
	file_pkg_api_v1_post_post_proto_rawDescOnce.Do(func() {
		file_pkg_api_v1_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_v1_post_post_proto_rawDescData)
	})
	return file_pkg_api_v1_post_post_proto_rawDescData
}

var file_pkg_api_v1_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_api_v1_post_post_proto_goTypes = []interface{}{
	(*Post)(nil),                  // 0: v1.post.Post
	(*WritePostRequest)(nil),      // 1: v1.post.WritePostRequest
	(*WritePostResponse)(nil),     // 2: v1.post.WritePostResponse
	(*GetPostsRequest)(nil),       // 3: v1.post.GetPostsRequest
	(*GetPostsResponse)(nil),      // 4: v1.post.GetPostsResponse
	(*UpdatePostRequest)(nil),     // 5: v1.post.UpdatePostRequest
	(*UpdatePostResponse)(nil),    // 6: v1.post.UpdatePostResponse
	(*DeletePostRequest)(nil),     // 7: v1.post.DeletePostRequest
	(*DeletePostResponse)(nil),    // 8: v1.post.DeletePostResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_pkg_api_v1_post_post_proto_depIdxs = []int32{
	9, // 0: v1.post.Post.created_at:type_name -> google.protobuf.Timestamp
	9, // 1: v1.post.Post.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: v1.post.WritePostResponse.post:type_name -> v1.post.Post
	0, // 3: v1.post.GetPostsResponse.posts:type_name -> v1.post.Post
	0, // 4: v1.post.UpdatePostResponse.post:type_name -> v1.post.Post
	1, // 5: v1.post.PostService.WritePost:input_type -> v1.post.WritePostRequest
	3, // 6: v1.post.PostService.GetPosts:input_type -> v1.post.GetPostsRequest
	5, // 7: v1.post.PostService.UpdatePost:input_type -> v1.post.UpdatePostRequest
	7, // 8: v1.post.PostService.DeletePost:input_type -> v1.post.DeletePostRequest
	2, // 9: v1.post.PostService.WritePost:output_type -> v1.post.WritePostResponse
	4, // 10: v1.post.PostService.GetPosts:output_type -> v1.post.GetPostsResponse
	6, // 11: v1.post.PostService.UpdatePost:output_type -> v1.post.UpdatePostResponse
	8, // 12: v1.post.PostService.DeletePost:output_type -> v1.post.DeletePostResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_api_v1_post_post_proto_init() }
func file_pkg_api_v1_post_post_proto_init() {
	if File_pkg_api_v1_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_v1_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePostRequest); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePostResponse); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsResponse); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostResponse); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_pkg_api_v1_post_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResponse); i {
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
			RawDescriptor: file_pkg_api_v1_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_v1_post_post_proto_goTypes,
		DependencyIndexes: file_pkg_api_v1_post_post_proto_depIdxs,
		MessageInfos:      file_pkg_api_v1_post_post_proto_msgTypes,
	}.Build()
	File_pkg_api_v1_post_post_proto = out.File
	file_pkg_api_v1_post_post_proto_rawDesc = nil
	file_pkg_api_v1_post_post_proto_goTypes = nil
	file_pkg_api_v1_post_post_proto_depIdxs = nil
}
