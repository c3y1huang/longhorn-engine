// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RemoveFileRequest struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFileRequest) Reset()         { *m = RemoveFileRequest{} }
func (m *RemoveFileRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveFileRequest) ProtoMessage()    {}
func (*RemoveFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *RemoveFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFileRequest.Unmarshal(m, b)
}
func (m *RemoveFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFileRequest.Marshal(b, m, deterministic)
}
func (m *RemoveFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFileRequest.Merge(m, src)
}
func (m *RemoveFileRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveFileRequest.Size(m)
}
func (m *RemoveFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFileRequest proto.InternalMessageInfo

func (m *RemoveFileRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

type RenameFileRequest struct {
	OldFileName          string   `protobuf:"bytes,1,opt,name=oldFileName,proto3" json:"oldFileName,omitempty"`
	NewFileName          string   `protobuf:"bytes,2,opt,name=newFileName,proto3" json:"newFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameFileRequest) Reset()         { *m = RenameFileRequest{} }
func (m *RenameFileRequest) String() string { return proto.CompactTextString(m) }
func (*RenameFileRequest) ProtoMessage()    {}
func (*RenameFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *RenameFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameFileRequest.Unmarshal(m, b)
}
func (m *RenameFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameFileRequest.Marshal(b, m, deterministic)
}
func (m *RenameFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameFileRequest.Merge(m, src)
}
func (m *RenameFileRequest) XXX_Size() int {
	return xxx_messageInfo_RenameFileRequest.Size(m)
}
func (m *RenameFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameFileRequest proto.InternalMessageInfo

func (m *RenameFileRequest) GetOldFileName() string {
	if m != nil {
		return m.OldFileName
	}
	return ""
}

func (m *RenameFileRequest) GetNewFileName() string {
	if m != nil {
		return m.NewFileName
	}
	return ""
}

type LaunchReceiverRequest struct {
	ToFileName           string   `protobuf:"bytes,1,opt,name=toFileName,proto3" json:"toFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchReceiverRequest) Reset()         { *m = LaunchReceiverRequest{} }
func (m *LaunchReceiverRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchReceiverRequest) ProtoMessage()    {}
func (*LaunchReceiverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *LaunchReceiverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchReceiverRequest.Unmarshal(m, b)
}
func (m *LaunchReceiverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchReceiverRequest.Marshal(b, m, deterministic)
}
func (m *LaunchReceiverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchReceiverRequest.Merge(m, src)
}
func (m *LaunchReceiverRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchReceiverRequest.Size(m)
}
func (m *LaunchReceiverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchReceiverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchReceiverRequest proto.InternalMessageInfo

func (m *LaunchReceiverRequest) GetToFileName() string {
	if m != nil {
		return m.ToFileName
	}
	return ""
}

type LaunchReceiverReply struct {
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchReceiverReply) Reset()         { *m = LaunchReceiverReply{} }
func (m *LaunchReceiverReply) String() string { return proto.CompactTextString(m) }
func (*LaunchReceiverReply) ProtoMessage()    {}
func (*LaunchReceiverReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *LaunchReceiverReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchReceiverReply.Unmarshal(m, b)
}
func (m *LaunchReceiverReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchReceiverReply.Marshal(b, m, deterministic)
}
func (m *LaunchReceiverReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchReceiverReply.Merge(m, src)
}
func (m *LaunchReceiverReply) XXX_Size() int {
	return xxx_messageInfo_LaunchReceiverReply.Size(m)
}
func (m *LaunchReceiverReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchReceiverReply.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchReceiverReply proto.InternalMessageInfo

func (m *LaunchReceiverReply) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type SendFileRequest struct {
	FromFileName         string   `protobuf:"bytes,1,opt,name=fromFileName,proto3" json:"fromFileName,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendFileRequest) Reset()         { *m = SendFileRequest{} }
func (m *SendFileRequest) String() string { return proto.CompactTextString(m) }
func (*SendFileRequest) ProtoMessage()    {}
func (*SendFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *SendFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendFileRequest.Unmarshal(m, b)
}
func (m *SendFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendFileRequest.Marshal(b, m, deterministic)
}
func (m *SendFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendFileRequest.Merge(m, src)
}
func (m *SendFileRequest) XXX_Size() int {
	return xxx_messageInfo_SendFileRequest.Size(m)
}
func (m *SendFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendFileRequest proto.InternalMessageInfo

func (m *SendFileRequest) GetFromFileName() string {
	if m != nil {
		return m.FromFileName
	}
	return ""
}

func (m *SendFileRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SendFileRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type CoalesceRequest struct {
	FromFileName         string   `protobuf:"bytes,1,opt,name=fromFileName,proto3" json:"fromFileName,omitempty"`
	ToFileName           string   `protobuf:"bytes,2,opt,name=toFileName,proto3" json:"toFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoalesceRequest) Reset()         { *m = CoalesceRequest{} }
func (m *CoalesceRequest) String() string { return proto.CompactTextString(m) }
func (*CoalesceRequest) ProtoMessage()    {}
func (*CoalesceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *CoalesceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoalesceRequest.Unmarshal(m, b)
}
func (m *CoalesceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoalesceRequest.Marshal(b, m, deterministic)
}
func (m *CoalesceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoalesceRequest.Merge(m, src)
}
func (m *CoalesceRequest) XXX_Size() int {
	return xxx_messageInfo_CoalesceRequest.Size(m)
}
func (m *CoalesceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CoalesceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CoalesceRequest proto.InternalMessageInfo

func (m *CoalesceRequest) GetFromFileName() string {
	if m != nil {
		return m.FromFileName
	}
	return ""
}

func (m *CoalesceRequest) GetToFileName() string {
	if m != nil {
		return m.ToFileName
	}
	return ""
}

type CreateBackupRequest struct {
	SnapshotFileName     string            `protobuf:"bytes,1,opt,name=snapshotFileName,proto3" json:"snapshotFileName,omitempty"`
	BackupTarget         string            `protobuf:"bytes,2,opt,name=backupTarget,proto3" json:"backupTarget,omitempty"`
	VolumeName           string            `protobuf:"bytes,3,opt,name=volumeName,proto3" json:"volumeName,omitempty"`
	Labels               []string          `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Credential           map[string]string `protobuf:"bytes,5,rep,name=credential,proto3" json:"credential,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateBackupRequest) Reset()         { *m = CreateBackupRequest{} }
func (m *CreateBackupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBackupRequest) ProtoMessage()    {}
func (*CreateBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *CreateBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBackupRequest.Unmarshal(m, b)
}
func (m *CreateBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBackupRequest.Marshal(b, m, deterministic)
}
func (m *CreateBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBackupRequest.Merge(m, src)
}
func (m *CreateBackupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBackupRequest.Size(m)
}
func (m *CreateBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBackupRequest proto.InternalMessageInfo

func (m *CreateBackupRequest) GetSnapshotFileName() string {
	if m != nil {
		return m.SnapshotFileName
	}
	return ""
}

func (m *CreateBackupRequest) GetBackupTarget() string {
	if m != nil {
		return m.BackupTarget
	}
	return ""
}

func (m *CreateBackupRequest) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

func (m *CreateBackupRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateBackupRequest) GetCredential() map[string]string {
	if m != nil {
		return m.Credential
	}
	return nil
}

type CreateBackupReply struct {
	Backup               string   `protobuf:"bytes,1,opt,name=backup,proto3" json:"backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBackupReply) Reset()         { *m = CreateBackupReply{} }
func (m *CreateBackupReply) String() string { return proto.CompactTextString(m) }
func (*CreateBackupReply) ProtoMessage()    {}
func (*CreateBackupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *CreateBackupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBackupReply.Unmarshal(m, b)
}
func (m *CreateBackupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBackupReply.Marshal(b, m, deterministic)
}
func (m *CreateBackupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBackupReply.Merge(m, src)
}
func (m *CreateBackupReply) XXX_Size() int {
	return xxx_messageInfo_CreateBackupReply.Size(m)
}
func (m *CreateBackupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBackupReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBackupReply proto.InternalMessageInfo

func (m *CreateBackupReply) GetBackup() string {
	if m != nil {
		return m.Backup
	}
	return ""
}

type RemoveBackupRequest struct {
	Backup               string   `protobuf:"bytes,1,opt,name=backup,proto3" json:"backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBackupRequest) Reset()         { *m = RemoveBackupRequest{} }
func (m *RemoveBackupRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveBackupRequest) ProtoMessage()    {}
func (*RemoveBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *RemoveBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBackupRequest.Unmarshal(m, b)
}
func (m *RemoveBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBackupRequest.Marshal(b, m, deterministic)
}
func (m *RemoveBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBackupRequest.Merge(m, src)
}
func (m *RemoveBackupRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveBackupRequest.Size(m)
}
func (m *RemoveBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBackupRequest proto.InternalMessageInfo

func (m *RemoveBackupRequest) GetBackup() string {
	if m != nil {
		return m.Backup
	}
	return ""
}

type RestoreBackupRequest struct {
	Backup               string   `protobuf:"bytes,1,opt,name=backup,proto3" json:"backup,omitempty"`
	SnapshotFileName     string   `protobuf:"bytes,2,opt,name=snapshotFileName,proto3" json:"snapshotFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreBackupRequest) Reset()         { *m = RestoreBackupRequest{} }
func (m *RestoreBackupRequest) String() string { return proto.CompactTextString(m) }
func (*RestoreBackupRequest) ProtoMessage()    {}
func (*RestoreBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *RestoreBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreBackupRequest.Unmarshal(m, b)
}
func (m *RestoreBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreBackupRequest.Marshal(b, m, deterministic)
}
func (m *RestoreBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreBackupRequest.Merge(m, src)
}
func (m *RestoreBackupRequest) XXX_Size() int {
	return xxx_messageInfo_RestoreBackupRequest.Size(m)
}
func (m *RestoreBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreBackupRequest proto.InternalMessageInfo

func (m *RestoreBackupRequest) GetBackup() string {
	if m != nil {
		return m.Backup
	}
	return ""
}

func (m *RestoreBackupRequest) GetSnapshotFileName() string {
	if m != nil {
		return m.SnapshotFileName
	}
	return ""
}

type RestoreBackupIncrementallyRequest struct {
	Backup                 string   `protobuf:"bytes,1,opt,name=backup,proto3" json:"backup,omitempty"`
	DeltaFileName          string   `protobuf:"bytes,2,opt,name=deltaFileName,proto3" json:"deltaFileName,omitempty"`
	LastRestoredBackupName string   `protobuf:"bytes,3,opt,name=lastRestoredBackupName,proto3" json:"lastRestoredBackupName,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *RestoreBackupIncrementallyRequest) Reset()         { *m = RestoreBackupIncrementallyRequest{} }
func (m *RestoreBackupIncrementallyRequest) String() string { return proto.CompactTextString(m) }
func (*RestoreBackupIncrementallyRequest) ProtoMessage()    {}
func (*RestoreBackupIncrementallyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *RestoreBackupIncrementallyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreBackupIncrementallyRequest.Unmarshal(m, b)
}
func (m *RestoreBackupIncrementallyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreBackupIncrementallyRequest.Marshal(b, m, deterministic)
}
func (m *RestoreBackupIncrementallyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreBackupIncrementallyRequest.Merge(m, src)
}
func (m *RestoreBackupIncrementallyRequest) XXX_Size() int {
	return xxx_messageInfo_RestoreBackupIncrementallyRequest.Size(m)
}
func (m *RestoreBackupIncrementallyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreBackupIncrementallyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreBackupIncrementallyRequest proto.InternalMessageInfo

func (m *RestoreBackupIncrementallyRequest) GetBackup() string {
	if m != nil {
		return m.Backup
	}
	return ""
}

func (m *RestoreBackupIncrementallyRequest) GetDeltaFileName() string {
	if m != nil {
		return m.DeltaFileName
	}
	return ""
}

func (m *RestoreBackupIncrementallyRequest) GetLastRestoredBackupName() string {
	if m != nil {
		return m.LastRestoredBackupName
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{11}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RemoveFileRequest)(nil), "RemoveFileRequest")
	proto.RegisterType((*RenameFileRequest)(nil), "RenameFileRequest")
	proto.RegisterType((*LaunchReceiverRequest)(nil), "LaunchReceiverRequest")
	proto.RegisterType((*LaunchReceiverReply)(nil), "LaunchReceiverReply")
	proto.RegisterType((*SendFileRequest)(nil), "SendFileRequest")
	proto.RegisterType((*CoalesceRequest)(nil), "CoalesceRequest")
	proto.RegisterType((*CreateBackupRequest)(nil), "CreateBackupRequest")
	proto.RegisterMapType((map[string]string)(nil), "CreateBackupRequest.CredentialEntry")
	proto.RegisterType((*CreateBackupReply)(nil), "CreateBackupReply")
	proto.RegisterType((*RemoveBackupRequest)(nil), "RemoveBackupRequest")
	proto.RegisterType((*RestoreBackupRequest)(nil), "RestoreBackupRequest")
	proto.RegisterType((*RestoreBackupIncrementallyRequest)(nil), "RestoreBackupIncrementallyRequest")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5f, 0x6b, 0x13, 0x4f,
	0x14, 0x6d, 0xb2, 0x4d, 0x7e, 0xed, 0x6d, 0xfb, 0x4b, 0x3a, 0x49, 0x43, 0xd8, 0x07, 0x89, 0x43,
	0x29, 0xb1, 0xe2, 0x28, 0x15, 0x54, 0x0a, 0x82, 0x5a, 0x5b, 0x50, 0xc4, 0x87, 0x8d, 0x22, 0x08,
	0x3e, 0x4c, 0x36, 0xb7, 0x4d, 0xe8, 0xec, 0xce, 0x3a, 0x3b, 0x89, 0xec, 0x47, 0xf1, 0xd5, 0x0f,
	0xe7, 0xe7, 0x90, 0xfd, 0x97, 0xcc, 0x6e, 0x36, 0x54, 0xdf, 0x66, 0x6e, 0xce, 0x39, 0xf7, 0xde,
	0x9d, 0x73, 0x6f, 0x60, 0x57, 0x05, 0x2e, 0x0b, 0x94, 0xd4, 0x92, 0x3e, 0x86, 0x43, 0x07, 0x3d,
	0xb9, 0xc0, 0xab, 0x99, 0x40, 0x07, 0xbf, 0xcf, 0x31, 0xd4, 0xc4, 0x86, 0x9d, 0xeb, 0x99, 0xc0,
	0x8f, 0xdc, 0xc3, 0x7e, 0x6d, 0x50, 0x1b, 0xee, 0x3a, 0xcb, 0x3b, 0xfd, 0x12, 0x13, 0x7c, 0xee,
	0x15, 0x08, 0x03, 0xd8, 0x93, 0x62, 0x72, 0x55, 0xe4, 0x98, 0xa1, 0x18, 0xe1, 0xe3, 0x8f, 0x25,
	0xa2, 0x9e, 0x22, 0x8c, 0x10, 0x7d, 0x0e, 0x47, 0x1f, 0xf8, 0xdc, 0x77, 0xa7, 0x0e, 0xba, 0x38,
	0x5b, 0xa0, 0xca, 0xc5, 0xef, 0x01, 0x68, 0x59, 0xd2, 0x36, 0x22, 0xf4, 0x01, 0x74, 0xca, 0xc4,
	0x40, 0x44, 0x84, 0xc0, 0x76, 0x20, 0x95, 0x4e, 0x08, 0x0d, 0x27, 0x39, 0xd3, 0x6f, 0xd0, 0x1a,
	0xa1, 0x3f, 0x31, 0x4b, 0xa7, 0xb0, 0x7f, 0xad, 0xa4, 0x57, 0xd2, 0x2f, 0xc4, 0x62, 0xa9, 0xa9,
	0x0c, 0x75, 0x56, 0x75, 0x72, 0x5e, 0xca, 0x5b, 0x86, 0xfc, 0x67, 0x68, 0x5d, 0x48, 0x2e, 0x30,
	0x74, 0xff, 0x49, 0xbe, 0xd8, 0x60, 0x7d, 0xad, 0xc1, 0x5f, 0x75, 0xe8, 0x5c, 0x28, 0xe4, 0x1a,
	0xdf, 0x70, 0xf7, 0x76, 0x1e, 0xe4, 0xda, 0xa7, 0xd0, 0x0e, 0x7d, 0x1e, 0x84, 0x53, 0xa9, 0x4b,
	0xfa, 0x6b, 0xf1, 0xb8, 0x8e, 0x71, 0x42, 0xfe, 0xc4, 0xd5, 0x0d, 0xe6, 0xad, 0x14, 0x62, 0x71,
	0x1d, 0x0b, 0x29, 0xe6, 0x5e, 0xaa, 0x64, 0xa5, 0x75, 0xac, 0x22, 0xa4, 0x07, 0x4d, 0xc1, 0xc7,
	0x28, 0xc2, 0xfe, 0xf6, 0xc0, 0x1a, 0xee, 0x3a, 0xd9, 0x8d, 0xbc, 0x05, 0x70, 0x15, 0x4e, 0xd0,
	0xd7, 0x33, 0x2e, 0xfa, 0x8d, 0x81, 0x35, 0xdc, 0x3b, 0x3b, 0x66, 0x15, 0x15, 0xc7, 0xb1, 0x0c,
	0x76, 0xe9, 0x6b, 0x15, 0x39, 0x06, 0xcf, 0x7e, 0x09, 0xad, 0xd2, 0xcf, 0xa4, 0x0d, 0xd6, 0x2d,
	0x46, 0x59, 0x4f, 0xf1, 0x91, 0x74, 0xa1, 0xb1, 0xe0, 0x62, 0x9e, 0x7f, 0xa5, 0xf4, 0x72, 0x5e,
	0x7f, 0x51, 0xa3, 0x0f, 0xe1, 0xb0, 0x98, 0x31, 0xf6, 0x40, 0x0f, 0x9a, 0x69, 0x87, 0x99, 0x46,
	0x76, 0xa3, 0x8f, 0xa0, 0x93, 0xba, 0xbe, 0xf8, 0x41, 0x37, 0xc1, 0xbf, 0x42, 0xd7, 0xc1, 0x50,
	0x4b, 0xf5, 0x77, 0xf8, 0xca, 0x87, 0xa9, 0x57, 0x3f, 0x0c, 0xfd, 0x59, 0x83, 0xfb, 0x05, 0xf1,
	0x77, 0xbe, 0xab, 0xd0, 0x43, 0x5f, 0x73, 0x21, 0xa2, 0xbb, 0x32, 0x1d, 0xc3, 0xc1, 0x04, 0x85,
	0xe6, 0xa5, 0x34, 0xc5, 0x20, 0x79, 0x06, 0x3d, 0xc1, 0x43, 0x9d, 0xa5, 0x99, 0xa4, 0x79, 0x8c,
	0x47, 0xde, 0xf0, 0x2b, 0xfd, 0x0f, 0x1a, 0x97, 0x5e, 0xa0, 0xa3, 0xb3, 0xdf, 0x16, 0xb4, 0x47,
	0x91, 0xef, 0xbe, 0xbe, 0x41, 0x5f, 0x8f, 0x50, 0x2d, 0x66, 0x2e, 0x92, 0x53, 0x80, 0xd5, 0xea,
	0x20, 0x84, 0xad, 0xed, 0x11, 0xbb, 0xc9, 0x12, 0x3a, 0xdd, 0x4a, 0xb1, 0xf9, 0xd6, 0x48, 0xb0,
	0xa5, 0x15, 0x62, 0x60, 0x5f, 0xc1, 0xff, 0xc5, 0x79, 0x26, 0x3d, 0x56, 0xb9, 0x19, 0xec, 0x2e,
	0xab, 0x18, 0x7c, 0xba, 0x45, 0x4e, 0x60, 0x27, 0x1f, 0x73, 0xd2, 0x66, 0xa5, 0x89, 0x37, 0x32,
	0x9d, 0xc0, 0x4e, 0x3e, 0xaf, 0xa4, 0xcd, 0x4a, 0xa3, 0x6b, 0xe0, 0xce, 0x61, 0xdf, 0xf4, 0x16,
	0xe9, 0x56, 0x99, 0xdb, 0x26, 0x6c, 0xcd, 0x80, 0x74, 0x8b, 0x30, 0xd8, 0x37, 0xad, 0x46, 0xba,
	0xac, 0xc2, 0x79, 0x46, 0xae, 0x27, 0x70, 0x50, 0xb0, 0x03, 0x39, 0x62, 0x55, 0xde, 0x33, 0x18,
	0xef, 0xc1, 0xde, 0x6c, 0x20, 0x42, 0xd9, 0x9d, 0xee, 0x5a, 0x69, 0x8d, 0x9b, 0xc9, 0xbf, 0xc2,
	0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xbd, 0x40, 0xd1, 0x22, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncAgentServiceClient is the client API for SyncAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncAgentServiceClient interface {
	RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*Empty, error)
	RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*Empty, error)
	LaunchReceiver(ctx context.Context, in *LaunchReceiverRequest, opts ...grpc.CallOption) (*LaunchReceiverReply, error)
	SendFile(ctx context.Context, in *SendFileRequest, opts ...grpc.CallOption) (*Empty, error)
	Coalesce(ctx context.Context, in *CoalesceRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupReply, error)
	RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...grpc.CallOption) (*Empty, error)
	RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*Empty, error)
	RestoreBackupIncrementally(ctx context.Context, in *RestoreBackupIncrementallyRequest, opts ...grpc.CallOption) (*Empty, error)
}

type syncAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncAgentServiceClient(cc *grpc.ClientConn) SyncAgentServiceClient {
	return &syncAgentServiceClient{cc}
}

func (c *syncAgentServiceClient) RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/RemoveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/RenameFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) LaunchReceiver(ctx context.Context, in *LaunchReceiverRequest, opts ...grpc.CallOption) (*LaunchReceiverReply, error) {
	out := new(LaunchReceiverReply)
	err := c.cc.Invoke(ctx, "/SyncAgentService/LaunchReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) SendFile(ctx context.Context, in *SendFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/SendFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) Coalesce(ctx context.Context, in *CoalesceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/Coalesce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) CreateBackup(ctx context.Context, in *CreateBackupRequest, opts ...grpc.CallOption) (*CreateBackupReply, error) {
	out := new(CreateBackupReply)
	err := c.cc.Invoke(ctx, "/SyncAgentService/CreateBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/RemoveBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/RestoreBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncAgentServiceClient) RestoreBackupIncrementally(ctx context.Context, in *RestoreBackupIncrementallyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/SyncAgentService/RestoreBackupIncrementally", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncAgentServiceServer is the server API for SyncAgentService service.
type SyncAgentServiceServer interface {
	RemoveFile(context.Context, *RemoveFileRequest) (*Empty, error)
	RenameFile(context.Context, *RenameFileRequest) (*Empty, error)
	LaunchReceiver(context.Context, *LaunchReceiverRequest) (*LaunchReceiverReply, error)
	SendFile(context.Context, *SendFileRequest) (*Empty, error)
	Coalesce(context.Context, *CoalesceRequest) (*Empty, error)
	CreateBackup(context.Context, *CreateBackupRequest) (*CreateBackupReply, error)
	RemoveBackup(context.Context, *RemoveBackupRequest) (*Empty, error)
	RestoreBackup(context.Context, *RestoreBackupRequest) (*Empty, error)
	RestoreBackupIncrementally(context.Context, *RestoreBackupIncrementallyRequest) (*Empty, error)
}

// UnimplementedSyncAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncAgentServiceServer struct {
}

func (*UnimplementedSyncAgentServiceServer) RemoveFile(ctx context.Context, req *RemoveFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}
func (*UnimplementedSyncAgentServiceServer) RenameFile(ctx context.Context, req *RenameFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFile not implemented")
}
func (*UnimplementedSyncAgentServiceServer) LaunchReceiver(ctx context.Context, req *LaunchReceiverRequest) (*LaunchReceiverReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LaunchReceiver not implemented")
}
func (*UnimplementedSyncAgentServiceServer) SendFile(ctx context.Context, req *SendFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
func (*UnimplementedSyncAgentServiceServer) Coalesce(ctx context.Context, req *CoalesceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coalesce not implemented")
}
func (*UnimplementedSyncAgentServiceServer) CreateBackup(ctx context.Context, req *CreateBackupRequest) (*CreateBackupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackup not implemented")
}
func (*UnimplementedSyncAgentServiceServer) RemoveBackup(ctx context.Context, req *RemoveBackupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBackup not implemented")
}
func (*UnimplementedSyncAgentServiceServer) RestoreBackup(ctx context.Context, req *RestoreBackupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreBackup not implemented")
}
func (*UnimplementedSyncAgentServiceServer) RestoreBackupIncrementally(ctx context.Context, req *RestoreBackupIncrementallyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreBackupIncrementally not implemented")
}

func RegisterSyncAgentServiceServer(s *grpc.Server, srv SyncAgentServiceServer) {
	s.RegisterService(&_SyncAgentService_serviceDesc, srv)
}

func _SyncAgentService_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/RemoveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).RemoveFile(ctx, req.(*RemoveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_RenameFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).RenameFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/RenameFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).RenameFile(ctx, req.(*RenameFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_LaunchReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchReceiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).LaunchReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/LaunchReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).LaunchReceiver(ctx, req.(*LaunchReceiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_SendFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).SendFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/SendFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).SendFile(ctx, req.(*SendFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_Coalesce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoalesceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).Coalesce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/Coalesce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).Coalesce(ctx, req.(*CoalesceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_CreateBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).CreateBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/CreateBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).CreateBackup(ctx, req.(*CreateBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_RemoveBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).RemoveBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/RemoveBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).RemoveBackup(ctx, req.(*RemoveBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_RestoreBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).RestoreBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/RestoreBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).RestoreBackup(ctx, req.(*RestoreBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncAgentService_RestoreBackupIncrementally_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreBackupIncrementallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncAgentServiceServer).RestoreBackupIncrementally(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SyncAgentService/RestoreBackupIncrementally",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncAgentServiceServer).RestoreBackupIncrementally(ctx, req.(*RestoreBackupIncrementallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SyncAgentService",
	HandlerType: (*SyncAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveFile",
			Handler:    _SyncAgentService_RemoveFile_Handler,
		},
		{
			MethodName: "RenameFile",
			Handler:    _SyncAgentService_RenameFile_Handler,
		},
		{
			MethodName: "LaunchReceiver",
			Handler:    _SyncAgentService_LaunchReceiver_Handler,
		},
		{
			MethodName: "SendFile",
			Handler:    _SyncAgentService_SendFile_Handler,
		},
		{
			MethodName: "Coalesce",
			Handler:    _SyncAgentService_Coalesce_Handler,
		},
		{
			MethodName: "CreateBackup",
			Handler:    _SyncAgentService_CreateBackup_Handler,
		},
		{
			MethodName: "RemoveBackup",
			Handler:    _SyncAgentService_RemoveBackup_Handler,
		},
		{
			MethodName: "RestoreBackup",
			Handler:    _SyncAgentService_RestoreBackup_Handler,
		},
		{
			MethodName: "RestoreBackupIncrementally",
			Handler:    _SyncAgentService_RestoreBackupIncrementally_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}