// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: laptop_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopRequest) Reset() {
	*x = CreateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopRequest) ProtoMessage() {}

func (x *CreateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopRequest.ProtoReflect.Descriptor instead.
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLaptopRequest) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLaptopResponse) Reset() {
	*x = CreateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopResponse) ProtoMessage() {}

func (x *CreateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopResponse.ProtoReflect.Descriptor instead.
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLaptopResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchLaptopRequest) Reset() {
	*x = SearchLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopRequest) ProtoMessage() {}

func (x *SearchLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopRequest.ProtoReflect.Descriptor instead.
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchLaptopRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *SearchLaptopResponse) Reset() {
	*x = SearchLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchLaptopResponse) ProtoMessage() {}

func (x *SearchLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchLaptopResponse.ProtoReflect.Descriptor instead.
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{3}
}

func (x *SearchLaptopResponse) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data isUploadImageRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{4}
}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := x.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadImageRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId  string `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType string `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{5}
}

func (x *ImageInfo) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *ImageInfo) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{6}
}

func (x *UploadImageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UploadImageResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId string  `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	Score    float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RateLaptopRequest) Reset() {
	*x = RateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLaptopRequest) ProtoMessage() {}

func (x *RateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLaptopRequest.ProtoReflect.Descriptor instead.
func (*RateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{7}
}

func (x *RateLaptopRequest) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *RateLaptopRequest) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type RateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId     string  `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	RatedCount   uint32  `protobuf:"varint,2,opt,name=rated_count,json=ratedCount,proto3" json:"rated_count,omitempty"`
	AverageScore float64 `protobuf:"fixed64,3,opt,name=average_score,json=averageScore,proto3" json:"average_score,omitempty"`
}

func (x *RateLaptopResponse) Reset() {
	*x = RateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLaptopResponse) ProtoMessage() {}

func (x *RateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLaptopResponse.ProtoReflect.Descriptor instead.
func (*RateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_laptop_service_proto_rawDescGZIP(), []int{8}
}

func (x *RateLaptopResponse) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *RateLaptopResponse) GetRatedCount() uint32 {
	if x != nil {
		return x.RatedCount
	}
	return 0
}

func (x *RateLaptopResponse) GetAverageScore() float64 {
	if x != nil {
		return x.AverageScore
	}
	return 0
}

var File_laptop_service_proto protoreflect.FileDescriptor

var file_laptop_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66,
	0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4d, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x4e,
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x76,
	0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x39, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x46, 0x0a, 0x11, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x77, 0x0a, 0x12, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74,
	0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32, 0xb7, 0x04, 0x0a, 0x0d,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x2b,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72,
	0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x30, 0x01,
	0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61,
	0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70,
	0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x28, 0x01, 0x12,
	0x83, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x29,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72,
	0x2e, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x66, 0x72, 0x2e, 0x70, 0x63, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x28, 0x01, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_service_proto_rawDescOnce sync.Once
	file_laptop_service_proto_rawDescData = file_laptop_service_proto_rawDesc
)

func file_laptop_service_proto_rawDescGZIP() []byte {
	file_laptop_service_proto_rawDescOnce.Do(func() {
		file_laptop_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_service_proto_rawDescData)
	})
	return file_laptop_service_proto_rawDescData
}

var file_laptop_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_laptop_service_proto_goTypes = []interface{}{
	(*CreateLaptopRequest)(nil),  // 0: github.com.safr.pcbook.CreateLaptopRequest
	(*CreateLaptopResponse)(nil), // 1: github.com.safr.pcbook.CreateLaptopResponse
	(*SearchLaptopRequest)(nil),  // 2: github.com.safr.pcbook.SearchLaptopRequest
	(*SearchLaptopResponse)(nil), // 3: github.com.safr.pcbook.SearchLaptopResponse
	(*UploadImageRequest)(nil),   // 4: github.com.safr.pcbook.UploadImageRequest
	(*ImageInfo)(nil),            // 5: github.com.safr.pcbook.ImageInfo
	(*UploadImageResponse)(nil),  // 6: github.com.safr.pcbook.UploadImageResponse
	(*RateLaptopRequest)(nil),    // 7: github.com.safr.pcbook.RateLaptopRequest
	(*RateLaptopResponse)(nil),   // 8: github.com.safr.pcbook.RateLaptopResponse
	(*Laptop)(nil),               // 9: github.com.safr.pcbook.Laptop
	(*Filter)(nil),               // 10: github.com.safr.pcbook.Filter
}
var file_laptop_service_proto_depIdxs = []int32{
	9,  // 0: github.com.safr.pcbook.CreateLaptopRequest.laptop:type_name -> github.com.safr.pcbook.Laptop
	10, // 1: github.com.safr.pcbook.SearchLaptopRequest.filter:type_name -> github.com.safr.pcbook.Filter
	9,  // 2: github.com.safr.pcbook.SearchLaptopResponse.laptop:type_name -> github.com.safr.pcbook.Laptop
	5,  // 3: github.com.safr.pcbook.UploadImageRequest.info:type_name -> github.com.safr.pcbook.ImageInfo
	0,  // 4: github.com.safr.pcbook.LaptopService.CreateLaptop:input_type -> github.com.safr.pcbook.CreateLaptopRequest
	2,  // 5: github.com.safr.pcbook.LaptopService.SearchLaptop:input_type -> github.com.safr.pcbook.SearchLaptopRequest
	4,  // 6: github.com.safr.pcbook.LaptopService.UploadImage:input_type -> github.com.safr.pcbook.UploadImageRequest
	7,  // 7: github.com.safr.pcbook.LaptopService.RateLaptop:input_type -> github.com.safr.pcbook.RateLaptopRequest
	1,  // 8: github.com.safr.pcbook.LaptopService.CreateLaptop:output_type -> github.com.safr.pcbook.CreateLaptopResponse
	3,  // 9: github.com.safr.pcbook.LaptopService.SearchLaptop:output_type -> github.com.safr.pcbook.SearchLaptopResponse
	6,  // 10: github.com.safr.pcbook.LaptopService.UploadImage:output_type -> github.com.safr.pcbook.UploadImageResponse
	8,  // 11: github.com.safr.pcbook.LaptopService.RateLaptop:output_type -> github.com.safr.pcbook.RateLaptopResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_laptop_service_proto_init() }
func file_laptop_service_proto_init() {
	if File_laptop_service_proto != nil {
		return
	}
	file_filter_message_proto_init()
	file_laptop_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopResponse); i {
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
		file_laptop_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchLaptopResponse); i {
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
		file_laptop_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_laptop_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_laptop_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
		file_laptop_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLaptopRequest); i {
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
		file_laptop_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLaptopResponse); i {
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
	file_laptop_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laptop_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_laptop_service_proto_goTypes,
		DependencyIndexes: file_laptop_service_proto_depIdxs,
		MessageInfos:      file_laptop_service_proto_msgTypes,
	}.Build()
	File_laptop_service_proto = out.File
	file_laptop_service_proto_rawDesc = nil
	file_laptop_service_proto_goTypes = nil
	file_laptop_service_proto_depIdxs = nil
}
