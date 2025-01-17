// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: proto/app.proto

package app

import (
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

// Vote to have the App api launched faster!
type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{0}
}

func (x *VoteRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{1}
}

func (x *VoteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the app
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// owner id
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// associated token
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// time of reservation
	Created string `protobuf:"bytes,4,opt,name=created,proto3" json:"created,omitempty"`
	// time reservation expires
	Expires string `protobuf:"bytes,5,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{2}
}

func (x *Reservation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reservation) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Reservation) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Reservation) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Reservation) GetExpires() string {
	if x != nil {
		return x.Expires
	}
	return ""
}

// Reserve your app name
type ReserveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of your app e.g helloworld
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReserveRequest) Reset() {
	*x = ReserveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveRequest) ProtoMessage() {}

func (x *ReserveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveRequest.ProtoReflect.Descriptor instead.
func (*ReserveRequest) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{3}
}

func (x *ReserveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReserveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservation *Reservation `protobuf:"bytes,1,opt,name=reservation,proto3" json:"reservation,omitempty"`
}

func (x *ReserveResponse) Reset() {
	*x = ReserveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveResponse) ProtoMessage() {}

func (x *ReserveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveResponse.ProtoReflect.Descriptor instead.
func (*ReserveResponse) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{4}
}

func (x *ReserveResponse) GetReservation() *Reservation {
	if x != nil {
		return x.Reservation
	}
	return nil
}

var File_proto_app_proto protoreflect.FileDescriptor

var file_proto_app_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x70, 0x22, 0x27, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x28, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x24, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x6c, 0x0a, 0x03, 0x41, 0x70,
	0x70, 0x12, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x56, 0x6f, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_app_proto_rawDescOnce sync.Once
	file_proto_app_proto_rawDescData = file_proto_app_proto_rawDesc
)

func file_proto_app_proto_rawDescGZIP() []byte {
	file_proto_app_proto_rawDescOnce.Do(func() {
		file_proto_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_app_proto_rawDescData)
	})
	return file_proto_app_proto_rawDescData
}

var file_proto_app_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_app_proto_goTypes = []interface{}{
	(*VoteRequest)(nil),     // 0: app.VoteRequest
	(*VoteResponse)(nil),    // 1: app.VoteResponse
	(*Reservation)(nil),     // 2: app.Reservation
	(*ReserveRequest)(nil),  // 3: app.ReserveRequest
	(*ReserveResponse)(nil), // 4: app.ReserveResponse
}
var file_proto_app_proto_depIdxs = []int32{
	2, // 0: app.ReserveResponse.reservation:type_name -> app.Reservation
	3, // 1: app.App.Reserve:input_type -> app.ReserveRequest
	0, // 2: app.App.Vote:input_type -> app.VoteRequest
	4, // 3: app.App.Reserve:output_type -> app.ReserveResponse
	1, // 4: app.App.Vote:output_type -> app.VoteResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_app_proto_init() }
func file_proto_app_proto_init() {
	if File_proto_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_proto_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
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
		file_proto_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_proto_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveRequest); i {
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
		file_proto_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReserveResponse); i {
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
			RawDescriptor: file_proto_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_app_proto_goTypes,
		DependencyIndexes: file_proto_app_proto_depIdxs,
		MessageInfos:      file_proto_app_proto_msgTypes,
	}.Build()
	File_proto_app_proto = out.File
	file_proto_app_proto_rawDesc = nil
	file_proto_app_proto_goTypes = nil
	file_proto_app_proto_depIdxs = nil
}
