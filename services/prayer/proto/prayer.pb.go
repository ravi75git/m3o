// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/prayer.proto

package prayer

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

type PrayerTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// date for prayer times in YYYY-MM-DD format
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// fajr time
	Fajr string `protobuf:"bytes,2,opt,name=fajr,proto3" json:"fajr,omitempty"`
	// time of sunrise
	Sunrise string `protobuf:"bytes,3,opt,name=sunrise,proto3" json:"sunrise,omitempty"`
	// zuhr time
	Zuhr string `protobuf:"bytes,4,opt,name=zuhr,proto3" json:"zuhr,omitempty"`
	// asr time
	Asr string `protobuf:"bytes,5,opt,name=asr,proto3" json:"asr,omitempty"`
	// maghrib time
	Maghrib string `protobuf:"bytes,6,opt,name=maghrib,proto3" json:"maghrib,omitempty"`
	// isha time
	Isha string `protobuf:"bytes,7,opt,name=isha,proto3" json:"isha,omitempty"`
}

func (x *PrayerTime) Reset() {
	*x = PrayerTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prayer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrayerTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrayerTime) ProtoMessage() {}

func (x *PrayerTime) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prayer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrayerTime.ProtoReflect.Descriptor instead.
func (*PrayerTime) Descriptor() ([]byte, []int) {
	return file_proto_prayer_proto_rawDescGZIP(), []int{0}
}

func (x *PrayerTime) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PrayerTime) GetFajr() string {
	if x != nil {
		return x.Fajr
	}
	return ""
}

func (x *PrayerTime) GetSunrise() string {
	if x != nil {
		return x.Sunrise
	}
	return ""
}

func (x *PrayerTime) GetZuhr() string {
	if x != nil {
		return x.Zuhr
	}
	return ""
}

func (x *PrayerTime) GetAsr() string {
	if x != nil {
		return x.Asr
	}
	return ""
}

func (x *PrayerTime) GetMaghrib() string {
	if x != nil {
		return x.Maghrib
	}
	return ""
}

func (x *PrayerTime) GetIsha() string {
	if x != nil {
		return x.Isha
	}
	return ""
}

type TimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional date in YYYY-MM-DD format, otherwise uses today
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// number of days to request times for
	Days int32 `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	// location to retrieve prayer times for.
	// this can be a specific address, city, etc
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// optional latitude used in place of location
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// optional longitude used in place of location
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *TimesRequest) Reset() {
	*x = TimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prayer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimesRequest) ProtoMessage() {}

func (x *TimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prayer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimesRequest.ProtoReflect.Descriptor instead.
func (*TimesRequest) Descriptor() ([]byte, []int) {
	return file_proto_prayer_proto_rawDescGZIP(), []int{1}
}

func (x *TimesRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *TimesRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *TimesRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *TimesRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *TimesRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type TimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// date of request
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	// number of days
	Days int32 `protobuf:"varint,2,opt,name=days,proto3" json:"days,omitempty"`
	// location for the request
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// latitude of location
	Latitude float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// longitude of location
	Longitude float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// prayer times for the given location
	Times []*PrayerTime `protobuf:"bytes,6,rep,name=times,proto3" json:"times,omitempty"`
}

func (x *TimesResponse) Reset() {
	*x = TimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prayer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimesResponse) ProtoMessage() {}

func (x *TimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prayer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimesResponse.ProtoReflect.Descriptor instead.
func (*TimesResponse) Descriptor() ([]byte, []int) {
	return file_proto_prayer_proto_rawDescGZIP(), []int{2}
}

func (x *TimesResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *TimesResponse) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *TimesResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *TimesResponse) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *TimesResponse) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *TimesResponse) GetTimes() []*PrayerTime {
	if x != nil {
		return x.Times
	}
	return nil
}

var File_proto_prayer_proto protoreflect.FileDescriptor

var file_proto_prayer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x22, 0xa2, 0x01, 0x0a,
	0x0a, 0x50, 0x72, 0x61, 0x79, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x61, 0x6a, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x61, 0x6a, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x7a, 0x75, 0x68, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x75, 0x68,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x73, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x67, 0x68, 0x72, 0x69, 0x62, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x67, 0x68, 0x72, 0x69, 0x62, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x68, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x68,
	0x61, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x61, 0x79, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x32, 0x40, 0x0a, 0x06, 0x50, 0x72,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_prayer_proto_rawDescOnce sync.Once
	file_proto_prayer_proto_rawDescData = file_proto_prayer_proto_rawDesc
)

func file_proto_prayer_proto_rawDescGZIP() []byte {
	file_proto_prayer_proto_rawDescOnce.Do(func() {
		file_proto_prayer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prayer_proto_rawDescData)
	})
	return file_proto_prayer_proto_rawDescData
}

var file_proto_prayer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_prayer_proto_goTypes = []interface{}{
	(*PrayerTime)(nil),    // 0: prayer.PrayerTime
	(*TimesRequest)(nil),  // 1: prayer.TimesRequest
	(*TimesResponse)(nil), // 2: prayer.TimesResponse
}
var file_proto_prayer_proto_depIdxs = []int32{
	0, // 0: prayer.TimesResponse.times:type_name -> prayer.PrayerTime
	1, // 1: prayer.Prayer.Times:input_type -> prayer.TimesRequest
	2, // 2: prayer.Prayer.Times:output_type -> prayer.TimesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_prayer_proto_init() }
func file_proto_prayer_proto_init() {
	if File_proto_prayer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prayer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrayerTime); i {
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
		file_proto_prayer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimesRequest); i {
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
		file_proto_prayer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimesResponse); i {
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
			RawDescriptor: file_proto_prayer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_prayer_proto_goTypes,
		DependencyIndexes: file_proto_prayer_proto_depIdxs,
		MessageInfos:      file_proto_prayer_proto_msgTypes,
	}.Build()
	File_proto_prayer_proto = out.File
	file_proto_prayer_proto_rawDesc = nil
	file_proto_prayer_proto_goTypes = nil
	file_proto_prayer_proto_depIdxs = nil
}