// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: m.proto

package protobuf

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_m_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_m_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UUIDUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *UUIDUser) Reset() {
	*x = UUIDUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDUser) ProtoMessage() {}

func (x *UUIDUser) ProtoReflect() protoreflect.Message {
	mi := &file_m_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDUser.ProtoReflect.Descriptor instead.
func (*UUIDUser) Descriptor() ([]byte, []int) {
	return file_m_proto_rawDescGZIP(), []int{1}
}

func (x *UUIDUser) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUIDUser string `protobuf:"bytes,1,opt,name=UUIDUser,proto3" json:"UUIDUser,omitempty"`
	UUIDPet  string `protobuf:"bytes,2,opt,name=UUIDPet,proto3" json:"UUIDPet,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Kind     string `protobuf:"bytes,4,opt,name=Kind,proto3" json:"Kind,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_m_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_m_proto_rawDescGZIP(), []int{2}
}

func (x *Pet) GetUUIDUser() string {
	if x != nil {
		return x.UUIDUser
	}
	return ""
}

func (x *Pet) GetUUIDPet() string {
	if x != nil {
		return x.UUIDPet
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type UUIDPet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *UUIDPet) Reset() {
	*x = UUIDPet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDPet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDPet) ProtoMessage() {}

func (x *UUIDPet) ProtoReflect() protoreflect.Message {
	mi := &file_m_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDPet.ProtoReflect.Descriptor instead.
func (*UUIDPet) Descriptor() ([]byte, []int) {
	return file_m_proto_rawDescGZIP(), []int{3}
}

func (x *UUIDPet) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_m_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_m_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_m_proto_rawDescGZIP(), []int{4}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

var File_m_proto protoreflect.FileDescriptor

var file_m_proto_rawDesc = []byte{
	0x0a, 0x07, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x55, 0x55, 0x49,
	0x44, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x03, 0x50, 0x65, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x55, 0x49, 0x44, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x55, 0x49, 0x44, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x55, 0x55, 0x49, 0x44, 0x50, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x55,
	0x55, 0x49, 0x44, 0x50, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x69,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x1d,
	0x0a, 0x07, 0x55, 0x55, 0x49, 0x44, 0x50, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x1f, 0x0a,
	0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0x94,
	0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x55, 0x55, 0x49, 0x44,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x09, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x08, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x04, 0x2e, 0x50, 0x65, 0x74, 0x1a, 0x04,
	0x2e, 0x50, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x12, 0x08,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x50, 0x65, 0x74, 0x1a, 0x04, 0x2e, 0x50, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x06, 0x50, 0x75, 0x74, 0x50, 0x65, 0x74, 0x12, 0x04, 0x2e, 0x50, 0x65, 0x74, 0x1a, 0x04,
	0x2e, 0x50, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75,
	0x74, 0x12, 0x08, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x50, 0x65, 0x74, 0x1a, 0x08, 0x2e, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_m_proto_rawDescOnce sync.Once
	file_m_proto_rawDescData = file_m_proto_rawDesc
)

func file_m_proto_rawDescGZIP() []byte {
	file_m_proto_rawDescOnce.Do(func() {
		file_m_proto_rawDescData = protoimpl.X.CompressGZIP(file_m_proto_rawDescData)
	})
	return file_m_proto_rawDescData
}

var file_m_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_m_proto_goTypes = []interface{}{
	(*User)(nil),     // 0: User
	(*UUIDUser)(nil), // 1: UUIDUser
	(*Pet)(nil),      // 2: Pet
	(*UUIDPet)(nil),  // 3: UUIDPet
	(*Nothing)(nil),  // 4: Nothing
}
var file_m_proto_depIdxs = []int32{
	1, // 0: AuthChecker.GetUser:input_type -> UUIDUser
	0, // 1: AuthChecker.CreateUser:input_type -> User
	0, // 2: AuthChecker.UpdateUser:input_type -> User
	1, // 3: AuthChecker.DeleteUser:input_type -> UUIDUser
	1, // 4: AuthChecker.CheckUser:input_type -> UUIDUser
	2, // 5: AuthChecker.CreatePet:input_type -> Pet
	3, // 6: AuthChecker.GetPet:input_type -> UUIDPet
	2, // 7: AuthChecker.PutPet:input_type -> Pet
	3, // 8: AuthChecker.DeletePut:input_type -> UUIDPet
	0, // 9: AuthChecker.GetUser:output_type -> User
	4, // 10: AuthChecker.CreateUser:output_type -> Nothing
	0, // 11: AuthChecker.UpdateUser:output_type -> User
	4, // 12: AuthChecker.DeleteUser:output_type -> Nothing
	4, // 13: AuthChecker.CheckUser:output_type -> Nothing
	2, // 14: AuthChecker.CreatePet:output_type -> Pet
	2, // 15: AuthChecker.GetPet:output_type -> Pet
	2, // 16: AuthChecker.PutPet:output_type -> Pet
	4, // 17: AuthChecker.DeletePut:output_type -> Nothing
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_m_proto_init() }
func file_m_proto_init() {
	if File_m_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_m_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_m_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDUser); i {
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
		file_m_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_m_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDPet); i {
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
		file_m_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
			RawDescriptor: file_m_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_m_proto_goTypes,
		DependencyIndexes: file_m_proto_depIdxs,
		MessageInfos:      file_m_proto_msgTypes,
	}.Build()
	File_m_proto = out.File
	file_m_proto_rawDesc = nil
	file_m_proto_goTypes = nil
	file_m_proto_depIdxs = nil
}