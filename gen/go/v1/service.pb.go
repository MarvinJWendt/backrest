// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/service.proto

package v1

import (
	types "github.com/garethgeorge/backrest/gen/go/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClearHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string  `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId     string  `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	OnlyFailed bool    `protobuf:"varint,3,opt,name=only_failed,json=onlyFailed,proto3" json:"only_failed,omitempty"`
	Ops        []int64 `protobuf:"varint,4,rep,packed,name=ops,proto3" json:"ops,omitempty"`
}

func (x *ClearHistoryRequest) Reset() {
	*x = ClearHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearHistoryRequest) ProtoMessage() {}

func (x *ClearHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearHistoryRequest.ProtoReflect.Descriptor instead.
func (*ClearHistoryRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ClearHistoryRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ClearHistoryRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *ClearHistoryRequest) GetOnlyFailed() bool {
	if x != nil {
		return x.OnlyFailed
	}
	return false
}

func (x *ClearHistoryRequest) GetOps() []int64 {
	if x != nil {
		return x.Ops
	}
	return nil
}

type ForgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId     string `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	SnapshotId string `protobuf:"bytes,3,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
}

func (x *ForgetRequest) Reset() {
	*x = ForgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetRequest) ProtoMessage() {}

func (x *ForgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetRequest.ProtoReflect.Descriptor instead.
func (*ForgetRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ForgetRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ForgetRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *ForgetRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

type ListSnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId string `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (x *ListSnapshotsRequest) Reset() {
	*x = ListSnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotsRequest) ProtoMessage() {}

func (x *ListSnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotsRequest.ProtoReflect.Descriptor instead.
func (*ListSnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSnapshotsRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ListSnapshotsRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

type GetOperationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string  `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	PlanId     string  `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	SnapshotId string  `protobuf:"bytes,4,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	FlowId     int64   `protobuf:"varint,6,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Ids        []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	LastN      int64   `protobuf:"varint,3,opt,name=last_n,json=lastN,proto3" json:"last_n,omitempty"` // limit to the last n operations
}

func (x *GetOperationsRequest) Reset() {
	*x = GetOperationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationsRequest) ProtoMessage() {}

func (x *GetOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationsRequest.ProtoReflect.Descriptor instead.
func (*GetOperationsRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetOperationsRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *GetOperationsRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *GetOperationsRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *GetOperationsRequest) GetFlowId() int64 {
	if x != nil {
		return x.FlowId
	}
	return 0
}

func (x *GetOperationsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetOperationsRequest) GetLastN() int64 {
	if x != nil {
		return x.LastN
	}
	return 0
}

type RestoreSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlanId     string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	RepoId     string `protobuf:"bytes,5,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	SnapshotId string `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Target     string `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *RestoreSnapshotRequest) Reset() {
	*x = RestoreSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreSnapshotRequest) ProtoMessage() {}

func (x *RestoreSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreSnapshotRequest.ProtoReflect.Descriptor instead.
func (*RestoreSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *RestoreSnapshotRequest) GetPlanId() string {
	if x != nil {
		return x.PlanId
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RestoreSnapshotRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type ListSnapshotFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoId     string `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	SnapshotId string `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	Path       string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListSnapshotFilesRequest) Reset() {
	*x = ListSnapshotFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotFilesRequest) ProtoMessage() {}

func (x *ListSnapshotFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotFilesRequest.ProtoReflect.Descriptor instead.
func (*ListSnapshotFilesRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListSnapshotFilesRequest) GetRepoId() string {
	if x != nil {
		return x.RepoId
	}
	return ""
}

func (x *ListSnapshotFilesRequest) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *ListSnapshotFilesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ListSnapshotFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string     `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Entries []*LsEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListSnapshotFilesResponse) Reset() {
	*x = ListSnapshotFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnapshotFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnapshotFilesResponse) ProtoMessage() {}

func (x *ListSnapshotFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnapshotFilesResponse.ProtoReflect.Descriptor instead.
func (*ListSnapshotFilesResponse) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListSnapshotFilesResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListSnapshotFilesResponse) GetEntries() []*LsEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type LogDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref string `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *LogDataRequest) Reset() {
	*x = LogDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDataRequest) ProtoMessage() {}

func (x *LogDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDataRequest.ProtoReflect.Descriptor instead.
func (*LogDataRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *LogDataRequest) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

type LsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Path  string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Uid   int64  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid   int64  `protobuf:"varint,5,opt,name=gid,proto3" json:"gid,omitempty"`
	Size  int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Mode  int64  `protobuf:"varint,7,opt,name=mode,proto3" json:"mode,omitempty"`
	Mtime string `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Atime string `protobuf:"bytes,9,opt,name=atime,proto3" json:"atime,omitempty"`
	Ctime string `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
}

func (x *LsEntry) Reset() {
	*x = LsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsEntry) ProtoMessage() {}

func (x *LsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsEntry.ProtoReflect.Descriptor instead.
func (*LsEntry) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *LsEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LsEntry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LsEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LsEntry) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LsEntry) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *LsEntry) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *LsEntry) GetMode() int64 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *LsEntry) GetMtime() string {
	if x != nil {
		return x.Mtime
	}
	return ""
}

func (x *LsEntry) GetAtime() string {
	if x != nil {
		return x.Atime
	}
	return ""
}

func (x *LsEntry) GetCtime() string {
	if x != nil {
		return x.Ctime
	}
	return ""
}

var File_v1_service_proto protoreflect.FileDescriptor

var file_v1_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x13, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x6f, 0x70, 0x73, 0x22, 0x62, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x15, 0x0a, 0x06,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x68, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x56, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x22, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x22, 0xd3, 0x01, 0x0a, 0x07, 0x4c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xe2, 0x08, 0x0a, 0x08, 0x42, 0x61,
	0x63, 0x6b, 0x72, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00,
	0x12, 0x21, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x08, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x69, 0x63,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x12, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x50, 0x72,
	0x75, 0x6e, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x12, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x68, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72,
	0x65, 0x74, 0x68, 0x67, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x72, 0x65,
	0x73, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_service_proto_rawDescOnce sync.Once
	file_v1_service_proto_rawDescData = file_v1_service_proto_rawDesc
)

func file_v1_service_proto_rawDescGZIP() []byte {
	file_v1_service_proto_rawDescOnce.Do(func() {
		file_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_service_proto_rawDescData)
	})
	return file_v1_service_proto_rawDescData
}

var file_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_service_proto_goTypes = []interface{}{
	(*ClearHistoryRequest)(nil),       // 0: v1.ClearHistoryRequest
	(*ForgetRequest)(nil),             // 1: v1.ForgetRequest
	(*ListSnapshotsRequest)(nil),      // 2: v1.ListSnapshotsRequest
	(*GetOperationsRequest)(nil),      // 3: v1.GetOperationsRequest
	(*RestoreSnapshotRequest)(nil),    // 4: v1.RestoreSnapshotRequest
	(*ListSnapshotFilesRequest)(nil),  // 5: v1.ListSnapshotFilesRequest
	(*ListSnapshotFilesResponse)(nil), // 6: v1.ListSnapshotFilesResponse
	(*LogDataRequest)(nil),            // 7: v1.LogDataRequest
	(*LsEntry)(nil),                   // 8: v1.LsEntry
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
	(*Config)(nil),                    // 10: v1.Config
	(*Repo)(nil),                      // 11: v1.Repo
	(*types.StringValue)(nil),         // 12: types.StringValue
	(*types.Int64Value)(nil),          // 13: types.Int64Value
	(*OperationEvent)(nil),            // 14: v1.OperationEvent
	(*OperationList)(nil),             // 15: v1.OperationList
	(*ResticSnapshotList)(nil),        // 16: v1.ResticSnapshotList
	(*types.BytesValue)(nil),          // 17: types.BytesValue
	(*types.StringList)(nil),          // 18: types.StringList
}
var file_v1_service_proto_depIdxs = []int32{
	8,  // 0: v1.ListSnapshotFilesResponse.entries:type_name -> v1.LsEntry
	9,  // 1: v1.Backrest.GetConfig:input_type -> google.protobuf.Empty
	10, // 2: v1.Backrest.SetConfig:input_type -> v1.Config
	11, // 3: v1.Backrest.AddRepo:input_type -> v1.Repo
	9,  // 4: v1.Backrest.GetOperationEvents:input_type -> google.protobuf.Empty
	3,  // 5: v1.Backrest.GetOperations:input_type -> v1.GetOperationsRequest
	2,  // 6: v1.Backrest.ListSnapshots:input_type -> v1.ListSnapshotsRequest
	5,  // 7: v1.Backrest.ListSnapshotFiles:input_type -> v1.ListSnapshotFilesRequest
	12, // 8: v1.Backrest.IndexSnapshots:input_type -> types.StringValue
	12, // 9: v1.Backrest.Backup:input_type -> types.StringValue
	12, // 10: v1.Backrest.Prune:input_type -> types.StringValue
	1,  // 11: v1.Backrest.Forget:input_type -> v1.ForgetRequest
	4,  // 12: v1.Backrest.Restore:input_type -> v1.RestoreSnapshotRequest
	12, // 13: v1.Backrest.Unlock:input_type -> types.StringValue
	12, // 14: v1.Backrest.Stats:input_type -> types.StringValue
	13, // 15: v1.Backrest.Cancel:input_type -> types.Int64Value
	7,  // 16: v1.Backrest.GetLogs:input_type -> v1.LogDataRequest
	13, // 17: v1.Backrest.GetDownloadURL:input_type -> types.Int64Value
	0,  // 18: v1.Backrest.ClearHistory:input_type -> v1.ClearHistoryRequest
	12, // 19: v1.Backrest.PathAutocomplete:input_type -> types.StringValue
	10, // 20: v1.Backrest.GetConfig:output_type -> v1.Config
	10, // 21: v1.Backrest.SetConfig:output_type -> v1.Config
	10, // 22: v1.Backrest.AddRepo:output_type -> v1.Config
	14, // 23: v1.Backrest.GetOperationEvents:output_type -> v1.OperationEvent
	15, // 24: v1.Backrest.GetOperations:output_type -> v1.OperationList
	16, // 25: v1.Backrest.ListSnapshots:output_type -> v1.ResticSnapshotList
	6,  // 26: v1.Backrest.ListSnapshotFiles:output_type -> v1.ListSnapshotFilesResponse
	9,  // 27: v1.Backrest.IndexSnapshots:output_type -> google.protobuf.Empty
	9,  // 28: v1.Backrest.Backup:output_type -> google.protobuf.Empty
	9,  // 29: v1.Backrest.Prune:output_type -> google.protobuf.Empty
	9,  // 30: v1.Backrest.Forget:output_type -> google.protobuf.Empty
	9,  // 31: v1.Backrest.Restore:output_type -> google.protobuf.Empty
	9,  // 32: v1.Backrest.Unlock:output_type -> google.protobuf.Empty
	9,  // 33: v1.Backrest.Stats:output_type -> google.protobuf.Empty
	9,  // 34: v1.Backrest.Cancel:output_type -> google.protobuf.Empty
	17, // 35: v1.Backrest.GetLogs:output_type -> types.BytesValue
	12, // 36: v1.Backrest.GetDownloadURL:output_type -> types.StringValue
	9,  // 37: v1.Backrest.ClearHistory:output_type -> google.protobuf.Empty
	18, // 38: v1.Backrest.PathAutocomplete:output_type -> types.StringList
	20, // [20:39] is the sub-list for method output_type
	1,  // [1:20] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v1_service_proto_init() }
func file_v1_service_proto_init() {
	if File_v1_service_proto != nil {
		return
	}
	file_v1_config_proto_init()
	file_v1_restic_proto_init()
	file_v1_operations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearHistoryRequest); i {
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
		file_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetRequest); i {
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
		file_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotsRequest); i {
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
		file_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationsRequest); i {
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
		file_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreSnapshotRequest); i {
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
		file_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotFilesRequest); i {
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
		file_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnapshotFilesResponse); i {
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
		file_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogDataRequest); i {
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
		file_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsEntry); i {
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
			RawDescriptor: file_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_service_proto_goTypes,
		DependencyIndexes: file_v1_service_proto_depIdxs,
		MessageInfos:      file_v1_service_proto_msgTypes,
	}.Build()
	File_v1_service_proto = out.File
	file_v1_service_proto_rawDesc = nil
	file_v1_service_proto_goTypes = nil
	file_v1_service_proto_depIdxs = nil
}
