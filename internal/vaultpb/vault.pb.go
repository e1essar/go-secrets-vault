// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/vault.proto

package vaultpb

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

type Secret struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nonce         []byte                 `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Ciphertext    []byte                 `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_proto_vault_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_proto_vault_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Secret) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Secret) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type SecretID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecretID) Reset() {
	*x = SecretID{}
	mi := &file_proto_vault_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecretID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretID) ProtoMessage() {}

func (x *SecretID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretID.ProtoReflect.Descriptor instead.
func (*SecretID) Descriptor() ([]byte, []int) {
	return file_proto_vault_proto_rawDescGZIP(), []int{1}
}

func (x *SecretID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_vault_proto protoreflect.FileDescriptor

const file_proto_vault_proto_rawDesc = "" +
	"\n" +
	"\x11proto/vault.proto\x12\x05vault\"N\n" +
	"\x06Secret\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05nonce\x18\x02 \x01(\fR\x05nonce\x12\x1e\n" +
	"\n" +
	"ciphertext\x18\x03 \x01(\fR\n" +
	"ciphertext\"\x1a\n" +
	"\bSecretID\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\x86\x01\n" +
	"\x05Vault\x12'\n" +
	"\x05Store\x12\r.vault.Secret\x1a\x0f.vault.SecretID\x12*\n" +
	"\bRetrieve\x12\x0f.vault.SecretID\x1a\r.vault.Secret\x12(\n" +
	"\x06Rotate\x12\x0f.vault.SecretID\x1a\r.vault.SecretB\x1aZ\x18internal/vaultpb;vaultpbb\x06proto3"

var (
	file_proto_vault_proto_rawDescOnce sync.Once
	file_proto_vault_proto_rawDescData []byte
)

func file_proto_vault_proto_rawDescGZIP() []byte {
	file_proto_vault_proto_rawDescOnce.Do(func() {
		file_proto_vault_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_vault_proto_rawDesc), len(file_proto_vault_proto_rawDesc)))
	})
	return file_proto_vault_proto_rawDescData
}

var file_proto_vault_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_vault_proto_goTypes = []any{
	(*Secret)(nil),   // 0: vault.Secret
	(*SecretID)(nil), // 1: vault.SecretID
}
var file_proto_vault_proto_depIdxs = []int32{
	0, // 0: vault.Vault.Store:input_type -> vault.Secret
	1, // 1: vault.Vault.Retrieve:input_type -> vault.SecretID
	1, // 2: vault.Vault.Rotate:input_type -> vault.SecretID
	1, // 3: vault.Vault.Store:output_type -> vault.SecretID
	0, // 4: vault.Vault.Retrieve:output_type -> vault.Secret
	0, // 5: vault.Vault.Rotate:output_type -> vault.Secret
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_vault_proto_init() }
func file_proto_vault_proto_init() {
	if File_proto_vault_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_vault_proto_rawDesc), len(file_proto_vault_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vault_proto_goTypes,
		DependencyIndexes: file_proto_vault_proto_depIdxs,
		MessageInfos:      file_proto_vault_proto_msgTypes,
	}.Build()
	File_proto_vault_proto = out.File
	file_proto_vault_proto_goTypes = nil
	file_proto_vault_proto_depIdxs = nil
}
