// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: serv/incentives/v1/incentives.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
	_ = time.Kitchen
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Incentive defines an instance that organizes distribution conditions for a
// given smart contract
type Incentive struct {
	// contract address of the smart contract to be incentivized
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// allocations is a slice of denoms and percentages of rewards to be allocated
	Allocations github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=allocations,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"allocations"`
	// epochs defines the number of remaining epochs for the incentive
	Epochs uint32 `protobuf:"varint,3,opt,name=epochs,proto3" json:"epochs,omitempty"`
	// start_time of the incentive distribution
	StartTime time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// total_gas is the cumulative gas spent by all gas meters of the incentive during the epoch
	TotalGas uint64 `protobuf:"varint,5,opt,name=total_gas,json=totalGas,proto3" json:"total_gas,omitempty"`
}

func (m *Incentive) Reset()         { *m = Incentive{} }
func (m *Incentive) String() string { return proto.CompactTextString(m) }
func (*Incentive) ProtoMessage()    {}
func (*Incentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b81e40854aec77, []int{0}
}

func (m *Incentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Incentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Incentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Incentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Incentive.Merge(m, src)
}

func (m *Incentive) XXX_Size() int {
	return m.Size()
}

func (m *Incentive) XXX_DiscardUnknown() {
	xxx_messageInfo_Incentive.DiscardUnknown(m)
}

var xxx_messageInfo_Incentive proto.InternalMessageInfo

func (m *Incentive) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *Incentive) GetAllocations() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func (m *Incentive) GetEpochs() uint32 {
	if m != nil {
		return m.Epochs
	}
	return 0
}

func (m *Incentive) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Incentive) GetTotalGas() uint64 {
	if m != nil {
		return m.TotalGas
	}
	return 0
}

// GasMeter tracks the cumulative gas spent per participant in one epoch
type GasMeter struct {
	// contract is the hex address of the incentivized smart contract
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// participant address that interacts with the incentive
	Participant string `protobuf:"bytes,2,opt,name=participant,proto3" json:"participant,omitempty"`
	// cumulative_gas spent during the epoch
	CumulativeGas uint64 `protobuf:"varint,3,opt,name=cumulative_gas,json=cumulativeGas,proto3" json:"cumulative_gas,omitempty"`
}

func (m *GasMeter) Reset()         { *m = GasMeter{} }
func (m *GasMeter) String() string { return proto.CompactTextString(m) }
func (*GasMeter) ProtoMessage()    {}
func (*GasMeter) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b81e40854aec77, []int{1}
}

func (m *GasMeter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GasMeter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasMeter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GasMeter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasMeter.Merge(m, src)
}

func (m *GasMeter) XXX_Size() int {
	return m.Size()
}

func (m *GasMeter) XXX_DiscardUnknown() {
	xxx_messageInfo_GasMeter.DiscardUnknown(m)
}

var xxx_messageInfo_GasMeter proto.InternalMessageInfo

func (m *GasMeter) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *GasMeter) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *GasMeter) GetCumulativeGas() uint64 {
	if m != nil {
		return m.CumulativeGas
	}
	return 0
}

// RegisterIncentiveProposal is a gov Content type to register an incentive
type RegisterIncentiveProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address to be registered
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	// allocations defines the denoms and percentage of rewards to be allocated
	Allocations github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,4,rep,name=allocations,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"allocations"`
	// epochs is the number of remaining epochs for the incentive
	Epochs uint32 `protobuf:"varint,5,opt,name=epochs,proto3" json:"epochs,omitempty"`
}

func (m *RegisterIncentiveProposal) Reset()         { *m = RegisterIncentiveProposal{} }
func (m *RegisterIncentiveProposal) String() string { return proto.CompactTextString(m) }
func (*RegisterIncentiveProposal) ProtoMessage()    {}
func (*RegisterIncentiveProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b81e40854aec77, []int{2}
}

func (m *RegisterIncentiveProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RegisterIncentiveProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterIncentiveProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RegisterIncentiveProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterIncentiveProposal.Merge(m, src)
}

func (m *RegisterIncentiveProposal) XXX_Size() int {
	return m.Size()
}

func (m *RegisterIncentiveProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterIncentiveProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterIncentiveProposal proto.InternalMessageInfo

func (m *RegisterIncentiveProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *RegisterIncentiveProposal) GetAllocations() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func (m *RegisterIncentiveProposal) GetEpochs() uint32 {
	if m != nil {
		return m.Epochs
	}
	return 0
}

// CancelIncentiveProposal is a gov Content type to cancel an incentive
type CancelIncentiveProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address of the incentivized smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *CancelIncentiveProposal) Reset()         { *m = CancelIncentiveProposal{} }
func (m *CancelIncentiveProposal) String() string { return proto.CompactTextString(m) }
func (*CancelIncentiveProposal) ProtoMessage()    {}
func (*CancelIncentiveProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_95b81e40854aec77, []int{3}
}

func (m *CancelIncentiveProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *CancelIncentiveProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelIncentiveProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *CancelIncentiveProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelIncentiveProposal.Merge(m, src)
}

func (m *CancelIncentiveProposal) XXX_Size() int {
	return m.Size()
}

func (m *CancelIncentiveProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelIncentiveProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CancelIncentiveProposal proto.InternalMessageInfo

func (m *CancelIncentiveProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CancelIncentiveProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CancelIncentiveProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*Incentive)(nil), "serv.incentives.v1.Incentive")
	proto.RegisterType((*GasMeter)(nil), "serv.incentives.v1.GasMeter")
	proto.RegisterType((*RegisterIncentiveProposal)(nil), "serv.incentives.v1.RegisterIncentiveProposal")
	proto.RegisterType((*CancelIncentiveProposal)(nil), "serv.incentives.v1.CancelIncentiveProposal")
}

func init() {
	proto.RegisterFile("serv/incentives/v1/incentives.proto", fileDescriptor_95b81e40854aec77)
}

var fileDescriptor_95b81e40854aec77 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xbf, 0x6f, 0xd4, 0x30,
	0x18, 0x4d, 0xee, 0x47, 0x75, 0xe7, 0x53, 0x19, 0x42, 0x05, 0xe1, 0x40, 0x49, 0x74, 0x02, 0x29,
	0x12, 0xaa, 0x4d, 0xda, 0x8d, 0xf1, 0x0e, 0xe9, 0xc4, 0x80, 0x84, 0x22, 0x26, 0x96, 0xca, 0x71,
	0x4d, 0x6a, 0x91, 0xc4, 0x51, 0xfc, 0x5d, 0x04, 0x2b, 0x7f, 0x41, 0x27, 0x66, 0x66, 0xfe, 0x92,
	0x8e, 0x1d, 0x99, 0x28, 0xba, 0x5b, 0xf8, 0x33, 0x90, 0x9d, 0xa4, 0xa4, 0x4b, 0xc7, 0x5b, 0x12,
	0xbf, 0x67, 0x7f, 0x7a, 0xef, 0x7d, 0x9f, 0x8d, 0x9e, 0xf3, 0x3a, 0x97, 0x8a, 0x88, 0x82, 0xf1,
	0x02, 0x44, 0xcd, 0x15, 0xa9, 0xa3, 0x1e, 0xc2, 0x65, 0x25, 0x41, 0x3a, 0x0f, 0xcd, 0x29, 0xdc,
	0xe3, 0xeb, 0x68, 0xee, 0x31, 0xa9, 0x74, 0x6d, 0x42, 0x15, 0x27, 0x75, 0x94, 0x70, 0xa0, 0x11,
	0x61, 0x52, 0x14, 0x4d, 0xd1, 0xfc, 0x28, 0x95, 0xa9, 0x34, 0x4b, 0xa2, 0x57, 0x2d, 0xeb, 0xa7,
	0x52, 0xa6, 0x19, 0x27, 0x06, 0x25, 0x9b, 0x4f, 0x04, 0x44, 0xce, 0x15, 0xd0, 0xbc, 0x6c, 0x0e,
	0x2c, 0xbe, 0x0f, 0xd0, 0xf4, 0x6d, 0x27, 0xe4, 0xcc, 0xd1, 0x84, 0xc9, 0x02, 0x2a, 0xca, 0xc0,
	0xb5, 0x03, 0x3b, 0x9c, 0xc6, 0xb7, 0xd8, 0x51, 0x68, 0x46, 0xb3, 0x4c, 0x32, 0x0a, 0x42, 0x16,
	0xca, 0x1d, 0x04, 0xc3, 0x70, 0x76, 0xf2, 0x0c, 0x37, 0xb6, 0xb0, 0xb6, 0x85, 0x5b, 0x5b, 0xf8,
	0x0d, 0x67, 0x2b, 0x29, 0x8a, 0xe5, 0xe9, 0xd5, 0x6f, 0xdf, 0xfa, 0x79, 0xe3, 0xbf, 0x4c, 0x05,
	0x5c, 0x6c, 0x12, 0xcc, 0x64, 0x4e, 0xda, 0x18, 0xcd, 0xef, 0x58, 0x9d, 0x7f, 0x26, 0xf0, 0xb5,
	0xe4, 0xaa, 0xab, 0x51, 0x71, 0x5f, 0xc5, 0x79, 0x84, 0x0e, 0x78, 0x29, 0xd9, 0x85, 0x72, 0x87,
	0x81, 0x1d, 0x1e, 0xc6, 0x2d, 0x72, 0x56, 0x08, 0x29, 0xa0, 0x15, 0x9c, 0xe9, 0x3c, 0xee, 0x28,
	0xb0, 0xc3, 0xd9, 0xc9, 0x1c, 0x37, 0x61, 0x71, 0x17, 0x16, 0x7f, 0xe8, 0xc2, 0x2e, 0x27, 0xda,
	0xc9, 0xe5, 0x8d, 0x6f, 0xc7, 0x53, 0x53, 0xa7, 0x77, 0x9c, 0xa7, 0x68, 0x0a, 0x12, 0x68, 0x76,
	0x96, 0x52, 0xe5, 0x8e, 0x03, 0x3b, 0x1c, 0xc5, 0x13, 0x43, 0xac, 0xa9, 0x5a, 0x48, 0x34, 0x59,
	0x53, 0xf5, 0x8e, 0x03, 0xaf, 0xee, 0x6d, 0x4b, 0x80, 0x66, 0x25, 0xad, 0x40, 0x30, 0x51, 0xd2,
	0x02, 0xdc, 0x81, 0xd9, 0xee, 0x53, 0xce, 0x0b, 0xf4, 0x80, 0x6d, 0xf2, 0x4d, 0x46, 0x75, 0x8b,
	0x8d, 0xd6, 0xd0, 0x68, 0x1d, 0xfe, 0x67, 0xb5, 0xe0, 0xb7, 0x01, 0x7a, 0x12, 0xf3, 0x54, 0x28,
	0xe0, 0xd5, 0xed, 0x44, 0xde, 0x57, 0xb2, 0x94, 0x8a, 0x66, 0xce, 0x11, 0x1a, 0x83, 0x80, 0x8c,
	0xb7, 0xfa, 0x0d, 0xd0, 0xe2, 0xe7, 0x5c, 0xb1, 0x4a, 0x94, 0xba, 0x5d, 0x9d, 0x78, 0x8f, 0xba,
	0x63, 0x7d, 0x78, 0xff, 0x44, 0x47, 0x7b, 0x9e, 0xe8, 0xb8, 0x3f, 0xd1, 0xd7, 0xa3, 0xbf, 0x3f,
	0x7c, 0x6b, 0xa1, 0xd0, 0xe3, 0x15, 0x2d, 0x18, 0xcf, 0xf6, 0xd2, 0x81, 0x46, 0x74, 0xb9, 0xbe,
	0xda, 0x7a, 0xf6, 0xf5, 0xd6, 0xb3, 0xff, 0x6c, 0x3d, 0xfb, 0x72, 0xe7, 0x59, 0xd7, 0x3b, 0xcf,
	0xfa, 0xb5, 0xf3, 0xac, 0x8f, 0xc7, 0xbd, 0x98, 0xcd, 0xd3, 0x6d, 0xbe, 0x75, 0xf4, 0x8a, 0x7c,
	0xe9, 0x3f, 0x63, 0x93, 0x38, 0x39, 0x30, 0x37, 0xef, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0xe9, 0x28, 0xf9, 0xe7, 0x03, 0x00, 0x00,
}

func (m *Incentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Incentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Incentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalGas != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.TotalGas))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintIncentives(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.Epochs != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.Epochs))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GasMeter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasMeter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasMeter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CumulativeGas != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.CumulativeGas))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterIncentiveProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterIncentiveProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterIncentiveProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Epochs != 0 {
		i = encodeVarintIncentives(dAtA, i, uint64(m.Epochs))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentives(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CancelIncentiveProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelIncentiveProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelIncentiveProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintIncentives(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentives(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentives(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Incentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	if m.Epochs != 0 {
		n += 1 + sovIncentives(uint64(m.Epochs))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovIncentives(uint64(l))
	if m.TotalGas != 0 {
		n += 1 + sovIncentives(uint64(m.TotalGas))
	}
	return n
}

func (m *GasMeter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if m.CumulativeGas != 0 {
		n += 1 + sovIncentives(uint64(m.CumulativeGas))
	}
	return n
}

func (m *RegisterIncentiveProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovIncentives(uint64(l))
		}
	}
	if m.Epochs != 0 {
		n += 1 + sovIncentives(uint64(m.Epochs))
	}
	return n
}

func (m *CancelIncentiveProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovIncentives(uint64(l))
	}
	return n
}

func sovIncentives(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozIncentives(x uint64) (n int) {
	return sovIncentives(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Incentive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Incentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Incentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, types.DecCoin{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			m.Epochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epochs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalGas", wireType)
			}
			m.TotalGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *GasMeter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GasMeter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasMeter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeGas", wireType)
			}
			m.CumulativeGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *RegisterIncentiveProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterIncentiveProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterIncentiveProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, types.DecCoin{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			m.Epochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epochs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *CancelIncentiveProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CancelIncentiveProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelIncentiveProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIncentives
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentives
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentives(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentives
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipIncentives(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentives
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIncentives
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIncentives
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentives
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentives
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentives        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentives          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentives = fmt.Errorf("proto: unexpected end of group")
)
