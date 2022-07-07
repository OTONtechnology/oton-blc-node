package tx

import (
 fmt "fmt"
 _ "github.com/gogo/protobuf/gogoproto"
 proto "github.com/gogo/protobuf/proto"
 io "io"
 math "math"
 math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3

type Coin struct {
 Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
 Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Coin) Reset() { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage() {}
func (*Coin) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *Coin) XXX_Merge(src proto.Message) {
 xxx_messageInfo_Coin.Merge(m, src)
}
func (m *Coin) XXX_Size() int {
 return m.Size()
}
func (m *Coin) XXX_DiscardUnknown() {
 xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func (m *Coin) GetName() string {
 if m != nil {
  return m.Name
 }
 return ""
}

func (m *Coin) GetAmount() int64 {
 if m != nil {
  return m.Amount
 }
 return 0
}

type Input struct {
 Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
 Coins []Coin `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins"`
 Sequence int64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
 Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
 PubKey []byte `protobuf:"bytes,5,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *Input) Reset() { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage() {}
func (*Input) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_Input.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *Input) XXX_Merge(src proto.Message) {
 xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
 return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
 xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *Input) GetCoins() []Coin {
 if m != nil {
  return m.Coins
 }
 return nil
}

func (m *Input) GetSequence() int64 {
 if m != nil {
  return m.Sequence
 }
 return 0
}

func (m *Input) GetSignature() []byte {
 if m != nil {
  return m.Signature
 }
 return nil
}

func (m *Input) GetPubKey() []byte {
 if m != nil {
  return m.PubKey
 }
 return nil
}





type Output struct {
 Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
 Coins []Coin `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins"`
}

func (m *Output) Reset() { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage() {}
func (*Output) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_Output.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *Output) XXX_Merge(src proto.Message) {
 xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
 return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
 xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *Output) GetCoins() []Coin {
 if m != nil {
  return m.Coins
 }
 return nil
}




type Raw struct {
 Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
 Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (m *Raw) Reset() { *m = Raw{} }
func (m *Raw) String() string { return proto.CompactTextString(m) }
func (*Raw) ProtoMessage() {}
func (*Raw) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *Raw) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_Raw.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *Raw) XXX_Merge(src proto.Message) {
 xxx_messageInfo_Raw.Merge(m, src)
}
func (m *Raw) XXX_Size() int {
 return m.Size()
}
func (m *Raw) XXX_DiscardUnknown() {
 xxx_messageInfo_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_Raw proto.InternalMessageInfo

func (m *Raw) GetType() string {
 if m != nil {
  return m.Type
 }
 return ""
}

func (m *Raw) GetRaw() []byte {
 if m != nil {
  return m.Raw
 }
 return nil
}


type SendCoins struct {
 Gas int64 `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
 Fee Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee"`
 Inputs []Input `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs"`
 Outputs []Output `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs"`
}

func (m *SendCoins) Reset() { *m = SendCoins{} }
func (m *SendCoins) String() string { return proto.CompactTextString(m) }
func (*SendCoins) ProtoMessage() {}
func (*SendCoins) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{4}
}
func (m *SendCoins) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SendCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SendCoins.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SendCoins) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SendCoins.Merge(m, src)
}
func (m *SendCoins) XXX_Size() int {
 return m.Size()
}
func (m *SendCoins) XXX_DiscardUnknown() {
 xxx_messageInfo_SendCoins.DiscardUnknown(m)
}

var xxx_messageInfo_SendCoins proto.InternalMessageInfo

func (m *SendCoins) GetGas() int64 {
 if m != nil {
  return m.Gas
 }
 return 0
}

func (m *SendCoins) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SendCoins) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SendCoins) GetOutputs() []Output {
 if m != nil {
  return m.Outputs
 }
 return nil
}


type MintCoins struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`
 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`
 Outputs []Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs"`
 Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
 Amount int64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
 DecimalPoint uint32 `protobuf:"varint,6,opt,name=decimal_point,json=decimalPoint,proto3" json:"decimal_point,omitempty"`
 Movable bool `protobuf:"varint,7,opt,name=movable,proto3" json:"movable,omitempty"`
 Delta int64 `protobuf:"varint,8,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (m *MintCoins) Reset() { *m = MintCoins{} }
func (m *MintCoins) String() string { return proto.CompactTextString(m) }
func (*MintCoins) ProtoMessage() {}
func (*MintCoins) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{5}
}
func (m *MintCoins) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *MintCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_MintCoins.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *MintCoins) XXX_Merge(src proto.Message) {
 xxx_messageInfo_MintCoins.Merge(m, src)
}
func (m *MintCoins) XXX_Size() int {
 return m.Size()
}
func (m *MintCoins) XXX_DiscardUnknown() {
 xxx_messageInfo_MintCoins.DiscardUnknown(m)
}

var xxx_messageInfo_MintCoins proto.InternalMessageInfo

func (m *MintCoins) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *MintCoins) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *MintCoins) GetOutputs() []Output {
 if m != nil {
  return m.Outputs
 }
 return nil
}

func (m *MintCoins) GetName() string {
 if m != nil {
  return m.Name
 }
 return ""
}

func (m *MintCoins) GetAmount() int64 {
 if m != nil {
  return m.Amount
 }
 return 0
}

func (m *MintCoins) GetDecimalPoint() uint32 {
 if m != nil {
  return m.DecimalPoint
 }
 return 0
}

func (m *MintCoins) GetMovable() bool {
 if m != nil {
  return m.Movable
 }
 return false
}

func (m *MintCoins) GetDelta() int64 {
 if m != nil {
  return m.Delta
 }
 return 0
}


type SetSalePrice struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`


 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`


 Price Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
}

func (m *SetSalePrice) Reset() { *m = SetSalePrice{} }
func (m *SetSalePrice) String() string { return proto.CompactTextString(m) }
func (*SetSalePrice) ProtoMessage() {}
func (*SetSalePrice) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{6}
}
func (m *SetSalePrice) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SetSalePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SetSalePrice.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SetSalePrice) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SetSalePrice.Merge(m, src)
}
func (m *SetSalePrice) XXX_Size() int {
 return m.Size()
}
func (m *SetSalePrice) XXX_DiscardUnknown() {
 xxx_messageInfo_SetSalePrice.DiscardUnknown(m)
}

var xxx_messageInfo_SetSalePrice proto.InternalMessageInfo

func (m *SetSalePrice) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SetSalePrice) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SetSalePrice) GetName() string {
 if m != nil {
  return m.Name
 }
 return ""
}

func (m *SetSalePrice) GetPrice() Coin {
 if m != nil {
  return m.Price
 }
 return Coin{}
}


type CreateAMC struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`



 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`


 Address []byte `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`

 Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`

 Params []AMCParam `protobuf:"bytes,5,rep,name=params,proto3" json:"params"`
}

func (m *CreateAMC) Reset() { *m = CreateAMC{} }
func (m *CreateAMC) String() string { return proto.CompactTextString(m) }
func (*CreateAMC) ProtoMessage() {}
func (*CreateAMC) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{7}
}
func (m *CreateAMC) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *CreateAMC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_CreateAMC.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *CreateAMC) XXX_Merge(src proto.Message) {
 xxx_messageInfo_CreateAMC.Merge(m, src)
}
func (m *CreateAMC) XXX_Size() int {
 return m.Size()
}
func (m *CreateAMC) XXX_DiscardUnknown() {
 xxx_messageInfo_CreateAMC.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAMC proto.InternalMessageInfo

func (m *CreateAMC) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *CreateAMC) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *CreateAMC) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *CreateAMC) GetName() string {
 if m != nil {
  return m.Name
 }
 return ""
}

func (m *CreateAMC) GetParams() []AMCParam {
 if m != nil {
  return m.Params
 }
 return nil
}


type AMCParam struct {
 Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
 Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AMCParam) Reset() { *m = AMCParam{} }
func (m *AMCParam) String() string { return proto.CompactTextString(m) }
func (*AMCParam) ProtoMessage() {}
func (*AMCParam) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{8}
}
func (m *AMCParam) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *AMCParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_AMCParam.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *AMCParam) XXX_Merge(src proto.Message) {
 xxx_messageInfo_AMCParam.Merge(m, src)
}
func (m *AMCParam) XXX_Size() int {
 return m.Size()
}
func (m *AMCParam) XXX_DiscardUnknown() {
 xxx_messageInfo_AMCParam.DiscardUnknown(m)
}

var xxx_messageInfo_AMCParam proto.InternalMessageInfo

func (m *AMCParam) GetKey() string {
 if m != nil {
  return m.Key
 }
 return ""
}

func (m *AMCParam) GetValue() int64 {
 if m != nil {
  return m.Value
 }
 return 0
}



type SetInAMC struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`


 Address []byte `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`

 Referal []byte `protobuf:"bytes,4,opt,name=referal,proto3" json:"referal,omitempty"`


 Sponsor []byte `protobuf:"bytes,5,opt,name=sponsor,proto3" json:"sponsor,omitempty"`
}

func (m *SetInAMC) Reset() { *m = SetInAMC{} }
func (m *SetInAMC) String() string { return proto.CompactTextString(m) }
func (*SetInAMC) ProtoMessage() {}
func (*SetInAMC) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{9}
}
func (m *SetInAMC) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SetInAMC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SetInAMC.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SetInAMC) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SetInAMC.Merge(m, src)
}
func (m *SetInAMC) XXX_Size() int {
 return m.Size()
}
func (m *SetInAMC) XXX_DiscardUnknown() {
 xxx_messageInfo_SetInAMC.DiscardUnknown(m)
}

var xxx_messageInfo_SetInAMC proto.InternalMessageInfo

func (m *SetInAMC) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SetInAMC) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SetInAMC) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *SetInAMC) GetReferal() []byte {
 if m != nil {
  return m.Referal
 }
 return nil
}

func (m *SetInAMC) GetSponsor() []byte {
 if m != nil {
  return m.Sponsor
 }
 return nil
}



type BuyInAMC struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`


 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Address []byte `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`

 Referal []byte `protobuf:"bytes,4,opt,name=referal,proto3" json:"referal,omitempty"`

 Value Coin `protobuf:"bytes,5,opt,name=value,proto3" json:"value"`
 Delta Coin `protobuf:"bytes,6,opt,name=delta,proto3" json:"delta"`
}

func (m *BuyInAMC) Reset() { *m = BuyInAMC{} }
func (m *BuyInAMC) String() string { return proto.CompactTextString(m) }
func (*BuyInAMC) ProtoMessage() {}
func (*BuyInAMC) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{10}
}
func (m *BuyInAMC) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *BuyInAMC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_BuyInAMC.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *BuyInAMC) XXX_Merge(src proto.Message) {
 xxx_messageInfo_BuyInAMC.Merge(m, src)
}
func (m *BuyInAMC) XXX_Size() int {
 return m.Size()
}
func (m *BuyInAMC) XXX_DiscardUnknown() {
 xxx_messageInfo_BuyInAMC.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInAMC proto.InternalMessageInfo

func (m *BuyInAMC) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *BuyInAMC) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *BuyInAMC) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *BuyInAMC) GetReferal() []byte {
 if m != nil {
  return m.Referal
 }
 return nil
}

func (m *BuyInAMC) GetValue() Coin {
 if m != nil {
  return m.Value
 }
 return Coin{}
}

func (m *BuyInAMC) GetDelta() Coin {
 if m != nil {
  return m.Delta
 }
 return Coin{}
}


type RefundInAMC struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`

 Buy []byte `protobuf:"bytes,4,opt,name=buy,proto3" json:"buy,omitempty"`

 Log string `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *RefundInAMC) Reset() { *m = RefundInAMC{} }
func (m *RefundInAMC) String() string { return proto.CompactTextString(m) }
func (*RefundInAMC) ProtoMessage() {}
func (*RefundInAMC) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{11}
}
func (m *RefundInAMC) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *RefundInAMC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_RefundInAMC.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *RefundInAMC) XXX_Merge(src proto.Message) {
 xxx_messageInfo_RefundInAMC.Merge(m, src)
}
func (m *RefundInAMC) XXX_Size() int {
 return m.Size()
}
func (m *RefundInAMC) XXX_DiscardUnknown() {
 xxx_messageInfo_RefundInAMC.DiscardUnknown(m)
}

var xxx_messageInfo_RefundInAMC proto.InternalMessageInfo

func (m *RefundInAMC) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *RefundInAMC) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *RefundInAMC) GetName() string {
 if m != nil {
  return m.Name
 }
 return ""
}

func (m *RefundInAMC) GetBuy() []byte {
 if m != nil {
  return m.Buy
 }
 return nil
}

func (m *RefundInAMC) GetLog() string {
 if m != nil {
  return m.Log
 }
 return ""
}




type ChangeAddressWant struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`



 Outputs []Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs"`
}

func (m *ChangeAddressWant) Reset() { *m = ChangeAddressWant{} }
func (m *ChangeAddressWant) String() string { return proto.CompactTextString(m) }
func (*ChangeAddressWant) ProtoMessage() {}
func (*ChangeAddressWant) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{12}
}
func (m *ChangeAddressWant) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *ChangeAddressWant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_ChangeAddressWant.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *ChangeAddressWant) XXX_Merge(src proto.Message) {
 xxx_messageInfo_ChangeAddressWant.Merge(m, src)
}
func (m *ChangeAddressWant) XXX_Size() int {
 return m.Size()
}
func (m *ChangeAddressWant) XXX_DiscardUnknown() {
 xxx_messageInfo_ChangeAddressWant.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAddressWant proto.InternalMessageInfo

func (m *ChangeAddressWant) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *ChangeAddressWant) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *ChangeAddressWant) GetOutputs() []Output {
 if m != nil {
  return m.Outputs
 }
 return nil
}





type ChangeAddressApply struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`


 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Votes []Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes"`

 Address []byte `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ChangeAddressApply) Reset() { *m = ChangeAddressApply{} }
func (m *ChangeAddressApply) String() string { return proto.CompactTextString(m) }
func (*ChangeAddressApply) ProtoMessage() {}
func (*ChangeAddressApply) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{13}
}
func (m *ChangeAddressApply) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *ChangeAddressApply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_ChangeAddressApply.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *ChangeAddressApply) XXX_Merge(src proto.Message) {
 xxx_messageInfo_ChangeAddressApply.Merge(m, src)
}
func (m *ChangeAddressApply) XXX_Size() int {
 return m.Size()
}
func (m *ChangeAddressApply) XXX_DiscardUnknown() {
 xxx_messageInfo_ChangeAddressApply.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAddressApply proto.InternalMessageInfo

func (m *ChangeAddressApply) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *ChangeAddressApply) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *ChangeAddressApply) GetVotes() []Vote {
 if m != nil {
  return m.Votes
 }
 return nil
}

func (m *ChangeAddressApply) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}



type Vote struct {
 Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
 Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
 PubKey []byte `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *Vote) Reset() { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage() {}
func (*Vote) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{14}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *Vote) XXX_Merge(src proto.Message) {
 xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
 return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
 xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetAddress() []byte {
 if m != nil {
  return m.Address
 }
 return nil
}

func (m *Vote) GetSignature() []byte {
 if m != nil {
  return m.Signature
 }
 return nil
}

func (m *Vote) GetPubKey() []byte {
 if m != nil {
  return m.PubKey
 }
 return nil
}


type BuyNode struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`


 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`


 Outputs []Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs"`

 Holder []byte `protobuf:"bytes,4,opt,name=holder,proto3" json:"holder,omitempty"`

 Validator []byte `protobuf:"bytes,5,opt,name=validator,proto3" json:"validator,omitempty"`

 Power int64 `protobuf:"varint,6,opt,name=power,proto3" json:"power,omitempty"`

 Value Coin `protobuf:"bytes,7,opt,name=value,proto3" json:"value"`
 Delta Coin `protobuf:"bytes,8,opt,name=delta,proto3" json:"delta"`
}

func (m *BuyNode) Reset() { *m = BuyNode{} }
func (m *BuyNode) String() string { return proto.CompactTextString(m) }
func (*BuyNode) ProtoMessage() {}
func (*BuyNode) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{15}
}
func (m *BuyNode) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *BuyNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_BuyNode.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *BuyNode) XXX_Merge(src proto.Message) {
 xxx_messageInfo_BuyNode.Merge(m, src)
}
func (m *BuyNode) XXX_Size() int {
 return m.Size()
}
func (m *BuyNode) XXX_DiscardUnknown() {
 xxx_messageInfo_BuyNode.DiscardUnknown(m)
}

var xxx_messageInfo_BuyNode proto.InternalMessageInfo

func (m *BuyNode) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *BuyNode) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *BuyNode) GetOutputs() []Output {
 if m != nil {
  return m.Outputs
 }
 return nil
}

func (m *BuyNode) GetHolder() []byte {
 if m != nil {
  return m.Holder
 }
 return nil
}

func (m *BuyNode) GetValidator() []byte {
 if m != nil {
  return m.Validator
 }
 return nil
}

func (m *BuyNode) GetPower() int64 {
 if m != nil {
  return m.Power
 }
 return 0
}

func (m *BuyNode) GetValue() Coin {
 if m != nil {
  return m.Value
 }
 return Coin{}
}

func (m *BuyNode) GetDelta() Coin {
 if m != nil {
  return m.Delta
 }
 return Coin{}
}



type SetInteresNode struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Value Coin `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
}

func (m *SetInteresNode) Reset() { *m = SetInteresNode{} }
func (m *SetInteresNode) String() string { return proto.CompactTextString(m) }
func (*SetInteresNode) ProtoMessage() {}
func (*SetInteresNode) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{16}
}
func (m *SetInteresNode) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SetInteresNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SetInteresNode.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SetInteresNode) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SetInteresNode.Merge(m, src)
}
func (m *SetInteresNode) XXX_Size() int {
 return m.Size()
}
func (m *SetInteresNode) XXX_DiscardUnknown() {
 xxx_messageInfo_SetInteresNode.DiscardUnknown(m)
}

var xxx_messageInfo_SetInteresNode proto.InternalMessageInfo

func (m *SetInteresNode) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SetInteresNode) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SetInteresNode) GetValue() Coin {
 if m != nil {
  return m.Value
 }
 return Coin{}
}


type SetNodeReward struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Value Coin `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
}

func (m *SetNodeReward) Reset() { *m = SetNodeReward{} }
func (m *SetNodeReward) String() string { return proto.CompactTextString(m) }
func (*SetNodeReward) ProtoMessage() {}
func (*SetNodeReward) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{17}
}
func (m *SetNodeReward) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SetNodeReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SetNodeReward.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SetNodeReward) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SetNodeReward.Merge(m, src)
}
func (m *SetNodeReward) XXX_Size() int {
 return m.Size()
}
func (m *SetNodeReward) XXX_DiscardUnknown() {
 xxx_messageInfo_SetNodeReward.DiscardUnknown(m)
}

var xxx_messageInfo_SetNodeReward proto.InternalMessageInfo

func (m *SetNodeReward) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SetNodeReward) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SetNodeReward) GetValue() Coin {
 if m != nil {
  return m.Value
 }
 return Coin{}
}


type SetNewAge struct {
 Fee Coin `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`

 Inputs []Input `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs"`

 Params []AgeParam `protobuf:"bytes,3,rep,name=params,proto3" json:"params"`
}

func (m *SetNewAge) Reset() { *m = SetNewAge{} }
func (m *SetNewAge) String() string { return proto.CompactTextString(m) }
func (*SetNewAge) ProtoMessage() {}
func (*SetNewAge) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{18}
}
func (m *SetNewAge) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *SetNewAge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_SetNewAge.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *SetNewAge) XXX_Merge(src proto.Message) {
 xxx_messageInfo_SetNewAge.Merge(m, src)
}
func (m *SetNewAge) XXX_Size() int {
 return m.Size()
}
func (m *SetNewAge) XXX_DiscardUnknown() {
 xxx_messageInfo_SetNewAge.DiscardUnknown(m)
}

var xxx_messageInfo_SetNewAge proto.InternalMessageInfo

func (m *SetNewAge) GetFee() Coin {
 if m != nil {
  return m.Fee
 }
 return Coin{}
}

func (m *SetNewAge) GetInputs() []Input {
 if m != nil {
  return m.Inputs
 }
 return nil
}

func (m *SetNewAge) GetParams() []AgeParam {
 if m != nil {
  return m.Params
 }
 return nil
}




type AgeParam struct {
 Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
 Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
 Svalue string `protobuf:"bytes,3,opt,name=svalue,proto3" json:"svalue,omitempty"`
}

func (m *AgeParam) Reset() { *m = AgeParam{} }
func (m *AgeParam) String() string { return proto.CompactTextString(m) }
func (*AgeParam) ProtoMessage() {}
func (*AgeParam) Descriptor() ([]byte, []int) {
 return fileDescriptor_0fd2153dc07d3b5c, []int{19}
}
func (m *AgeParam) XXX_Unmarshal(b []byte) error {
 return m.Unmarshal(b)
}
func (m *AgeParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
 if deterministic {
  return xxx_messageInfo_AgeParam.Marshal(b, m, deterministic)
 } else {
  b = b[:cap(b)]
  n, err := m.MarshalToSizedBuffer(b)
  if err != nil {
   return nil, err
  }
  return b[:n], nil
 }
}
func (m *AgeParam) XXX_Merge(src proto.Message) {
 xxx_messageInfo_AgeParam.Merge(m, src)
}
func (m *AgeParam) XXX_Size() int {
 return m.Size()
}
func (m *AgeParam) XXX_DiscardUnknown() {
 xxx_messageInfo_AgeParam.DiscardUnknown(m)
}

var xxx_messageInfo_AgeParam proto.InternalMessageInfo

func (m *AgeParam) GetKey() string {
 if m != nil {
  return m.Key
 }
 return ""
}

func (m *AgeParam) GetValue() int64 {
 if m != nil {
  return m.Value
 }
 return 0
}

func (m *AgeParam) GetSvalue() string {
 if m != nil {
  return m.Svalue
 }
 return ""
}

func init() {
 proto.RegisterType((*Coin)(nil), "tx.Coin")
 proto.RegisterType((*Input)(nil), "tx.Input")
 proto.RegisterType((*Output)(nil), "tx.Output")
 proto.RegisterType((*Raw)(nil), "tx.Raw")
 proto.RegisterType((*SendCoins)(nil), "tx.SendCoins")
 proto.RegisterType((*MintCoins)(nil), "tx.MintCoins")
 proto.RegisterType((*SetSalePrice)(nil), "tx.SetSalePrice")
 proto.RegisterType((*CreateAMC)(nil), "tx.CreateAMC")
 proto.RegisterType((*AMCParam)(nil), "tx.AMCParam")
 proto.RegisterType((*SetInAMC)(nil), "tx.SetInAMC")
 proto.RegisterType((*BuyInAMC)(nil), "tx.BuyInAMC")
 proto.RegisterType((*RefundInAMC)(nil), "tx.RefundInAMC")
 proto.RegisterType((*ChangeAddressWant)(nil), "tx.ChangeAddressWant")
 proto.RegisterType((*ChangeAddressApply)(nil), "tx.ChangeAddressApply")
 proto.RegisterType((*Vote)(nil), "tx.Vote")
 proto.RegisterType((*BuyNode)(nil), "tx.BuyNode")
 proto.RegisterType((*SetInteresNode)(nil), "tx.SetInteresNode")
 proto.RegisterType((*SetNodeReward)(nil), "tx.SetNodeReward")
 proto.RegisterType((*SetNewAge)(nil), "tx.SetNewAge")
 proto.RegisterType((*AgeParam)(nil), "tx.AgeParam")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{

 0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6f, 0xeb, 0x44,
 0x10, 0xcf, 0x66, 0x6d, 0xc7, 0x9e, 0x97, 0xa2, 0xb2, 0x42, 0x0f, 0xeb, 0x09, 0x85, 0xc8, 0x54,
 0x22, 0x7a, 0x48, 0x3d, 0x94, 0x4f, 0x90, 0xe4, 0xc2, 0x03, 0xf5, 0x51, 0x39, 0x12, 0xef, 0x58,
 0x6d, 0xe2, 0xa9, 0x6b, 0xe1, 0x78, 0x8d, 0xbd, 0x6e, 0x9a, 0x1b, 0x7f, 0x6e, 0x9c, 0x2a, 0xc1,
 0x01, 0x71, 0xe1, 0xc0, 0x17, 0xe0, 0x63, 0xf4, 0xd8, 0x0b, 0x12, 0x27, 0x84, 0xda, 0x2f, 0x82,
 0x76, 0x6d, 0xe7, 0x8f, 0x88, 0x42, 0x0e, 0x51, 0xfb, 0x6e, 0x33, 0x3b, 0xe3, 0xdd, 0xdf, 0x6f,
 0xf6, 0xb7, 0x33, 0x06, 0x5b, 0x5e, 0x1f, 0xa7, 0x99, 0x90, 0x82, 0x35, 0xe5, 0xf5, 0x0b, 0x08,
 0x45, 0x28, 0x4a, 0xdf, 0x3b, 0x01, 0x63, 0x28, 0xa2, 0x84, 0x31, 0x30, 0x12, 0x3e, 0x45, 0x97,
 0x74, 0x49, 0xcf, 0xf1, 0xb5, 0xcd, 0x9e, 0x83, 0xc5, 0xa7, 0xa2, 0x48, 0xa4, 0xdb, 0xec, 0x92,
 0x1e, 0xf5, 0x2b, 0xcf, 0xfb, 0x95, 0x80, 0xf9, 0x2a, 0x49, 0x0b, 0xc9, 0x5c, 0x68, 0xf1, 0x20,
 0xc8, 0x30, 0xcf, 0xf5, 0x87, 0x6d, 0xbf, 0x76, 0xd9, 0x11, 0x98, 0x13, 0x11, 0x25, 0xb9, 0xdb,
 0xec, 0xd2, 0xde, 0xb3, 0x13, 0xfb, 0x58, 0x5e, 0x1f, 0xab, 0x83, 0x06, 0xc6, 0xed, 0xdf, 0x1f,
 0x36, 0xfc, 0x32, 0xc8, 0x5e, 0x80, 0x9d, 0xe3, 0x37, 0x05, 0x26, 0x13, 0x74, 0xa9, 0x3e, 0x63,
 0xe1, 0xb3, 0x0f, 0xc0, 0xc9, 0xa3, 0x30, 0xe1, 0xb2, 0xc8, 0xd0, 0x35, 0xf4, 0xee, 0xcb, 0x05,
 0xf6, 0x3e, 0xb4, 0xd2, 0x62, 0x7c, 0xfe, 0x35, 0xce, 0x5d, 0x53, 0xc7, 0xac, 0xb4, 0x18, 0x7f,
 0x81, 0x73, 0xef, 0x33, 0xb0, 0xbe, 0x2c, 0xe4, 0x1e, 0xc0, 0x79, 0x9f, 0x00, 0xf5, 0xf9, 0x4c,
 0x55, 0x46, 0xce, 0xd3, 0x45, 0x65, 0x94, 0xcd, 0x0e, 0x81, 0x66, 0x7c, 0xa6, 0xcb, 0xd2, 0xf6,
 0x95, 0xe9, 0xfd, 0x4c, 0xc0, 0x19, 0x61, 0x12, 0x0c, 0x35, 0xaf, 0x43, 0xa0, 0x21, 0x2f, 0x8f,
 0xa5, 0xbe, 0x32, 0x59, 0x17, 0xe8, 0x05, 0xa2, 0xfe, 0xe2, 0xbf, 0x07, 0xaa, 0x10, 0xfb, 0x18,
 0xac, 0x48, 0x15, 0x35, 0x77, 0xa9, 0x46, 0xe5, 0xa8, 0x24, 0x5d, 0xe6, 0x2a, 0xab, 0x0a, 0xb3,
 0x97, 0xd0, 0x12, 0x9a, 0x61, 0xee, 0x1a, 0x3a, 0x13, 0x54, 0x66, 0x49, 0xba, 0x4a, 0xad, 0x13,
 0xbc, 0x1f, 0x9a, 0xe0, 0x9c, 0x46, 0x89, 0x2c, 0x61, 0x55, 0x20, 0xc8, 0x2e, 0x20, 0x9a, 0x3b,
 0x83, 0xa0, 0xff, 0x03, 0x62, 0xa1, 0x2d, 0x63, 0xa3, 0xb6, 0xcc, 0x55, 0x6d, 0xb1, 0x8f, 0xe0,
 0x20, 0xc0, 0x49, 0x34, 0xe5, 0xf1, 0x79, 0x2a, 0xa2, 0x44, 0xba, 0x56, 0x97, 0xf4, 0x0e, 0xfc,
 0x76, 0xb5, 0x78, 0xa6, 0xd6, 0xd4, 0xcd, 0x4e, 0xc5, 0x15, 0x1f, 0xc7, 0xe8, 0xb6, 0xba, 0xa4,
 0x67, 0xfb, 0xb5, 0xcb, 0xde, 0x03, 0x33, 0xc0, 0x58, 0x72, 0xd7, 0xd6, 0xbb, 0x96, 0x8e, 0xf7,
 0x13, 0x81, 0xf6, 0x08, 0xe5, 0x88, 0xc7, 0x78, 0x96, 0x45, 0x13, 0xdc, 0x67, 0x21, 0x6a, 0x72,
 0x74, 0x85, 0xdc, 0x11, 0x98, 0xa9, 0x3a, 0x47, 0x33, 0xde, 0xa0, 0x2f, 0x1d, 0xf4, 0xfe, 0x20,
 0xe0, 0x0c, 0x33, 0xe4, 0x12, 0xfb, 0xa7, 0xc3, 0x7d, 0x42, 0x5a, 0x11, 0x3e, 0x5d, 0x17, 0xfe,
 0xa6, 0x9b, 0x78, 0x09, 0x56, 0xca, 0x33, 0x3e, 0xcd, 0x5d, 0x53, 0x6f, 0xdb, 0x56, 0xdb, 0xf6,
 0x4f, 0x87, 0x67, 0x6a, 0xb1, 0xde, 0xb9, 0xcc, 0xf0, 0x4e, 0xc0, 0xae, 0x23, 0x4a, 0xe3, 0xea,
 0xf5, 0x95, 0xcf, 0x42, 0x99, 0xaa, 0xf8, 0x57, 0x3c, 0x2e, 0xb0, 0x6a, 0x17, 0xa5, 0xe3, 0xfd,
 0x4e, 0xc0, 0x1e, 0xa1, 0x7c, 0x95, 0x3c, 0x1a, 0x4b, 0x17, 0x5a, 0x19, 0x5e, 0x60, 0xc6, 0xe3,
 0xaa, 0x6f, 0xd4, 0xae, 0x8a, 0xe4, 0xa9, 0x48, 0x72, 0x91, 0x55, 0x5d, 0xa3, 0x76, 0xbd, 0x3f,
 0x09, 0xd8, 0x83, 0x62, 0xfe, 0x56, 0xa0, 0x3c, 0xaa, 0xeb, 0x68, 0x6e, 0x96, 0x8f, 0x0e, 0xaa,
 0xac, 0x52, 0xea, 0xd6, 0xe6, 0xac, 0x52, 0xfa, 0x37, 0x04, 0x9e, 0xf9, 0x78, 0x51, 0x24, 0xc1,
 0xde, 0xa9, 0x6d, 0x52, 0xfe, 0x21, 0xd0, 0x71, 0x31, 0xaf, 0x08, 0x29, 0x53, 0xad, 0xc4, 0x22,
 0xd4, 0x54, 0x1c, 0x5f, 0x99, 0xde, 0x8f, 0x04, 0xde, 0x1d, 0x5e, 0xf2, 0x24, 0xc4, 0x7e, 0x59,
 0x8a, 0x37, 0x3c, 0x91, 0x4f, 0xd4, 0x9b, 0xbc, 0xdf, 0x08, 0xb0, 0x35, 0x30, 0xfd, 0x34, 0x8d,
 0xe7, 0xfb, 0x44, 0xa3, 0x6e, 0x53, 0x48, 0xac, 0xb1, 0xe8, 0xcd, 0xbe, 0x12, 0x12, 0x17, 0xb7,
 0xa9, 0x82, 0xab, 0x3a, 0x31, 0xd6, 0x74, 0xe2, 0xbd, 0x01, 0x43, 0xa5, 0x6f, 0x19, 0x67, 0x6b,
 0x93, 0xb2, 0xb9, 0x65, 0x52, 0xd2, 0xb5, 0x49, 0xf9, 0x4b, 0x13, 0x5a, 0x83, 0x62, 0xfe, 0x5a,
 0x04, 0xf8, 0x54, 0x93, 0xe1, 0x39, 0x58, 0x97, 0x22, 0x0e, 0x30, 0xab, 0x48, 0x57, 0x9e, 0x62,
 0x74, 0xc5, 0xe3, 0x28, 0xe0, 0x72, 0xf1, 0x52, 0x97, 0x0b, 0xaa, 0xcf, 0xa4, 0x62, 0x86, 0x99,
 0x56, 0x3e, 0xf5, 0x4b, 0x67, 0xf9, 0x6a, 0x5a, 0x3b, 0xbd, 0x1a, 0x7b, 0xdb, 0xab, 0xf9, 0x8e,
 0xc0, 0x3b, 0xba, 0x67, 0x49, 0xcc, 0x30, 0xdf, 0x77, 0x85, 0x16, 0x48, 0xe9, 0x16, 0xa4, 0xde,
 0xb7, 0x04, 0x0e, 0x46, 0x28, 0xd5, 0xe1, 0x3e, 0xce, 0x78, 0x16, 0x3c, 0x3e, 0x84, 0xef, 0xf5,
 0x4f, 0x8d, 0x7c, 0x8d, 0xb3, 0x7e, 0xb8, 0x67, 0x8d, 0xd4, 0x33, 0x87, 0xae, 0xcc, 0x9c, 0x10,
 0x37, 0xcd, 0x9c, 0xcf, 0xc1, 0xae, 0x23, 0xbb, 0xce, 0x1c, 0xa5, 0xab, 0x7c, 0xc9, 0xcf, 0xf1,
 0x2b, 0x6f, 0xe0, 0xde, 0xde, 0x77, 0xc8, 0xdd, 0x7d, 0x87, 0xfc, 0x73, 0xdf, 0x21, 0x37, 0x0f,
 0x9d, 0xc6, 0xdd, 0x43, 0xa7, 0xf1, 0xd7, 0x43, 0xa7, 0x31, 0xb6, 0xf4, 0xef, 0xf0, 0xa7, 0xff,
 0x06, 0x00, 0x00, 0xff, 0xff, 0x67, 0xec, 0x12, 0x62, 0x2a, 0x0b, 0x00, 0x00,
}

func (m *Coin) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if m.Amount != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Amount))
  i--
  dAtA[i] = 0x10
 }
 if len(m.Name) > 0 {
  i -= len(m.Name)
  copy(dAtA[i:], m.Name)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *Input) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.PubKey) > 0 {
  i -= len(m.PubKey)
  copy(dAtA[i:], m.PubKey)
  i = encodeVarintTx(dAtA, i, uint64(len(m.PubKey)))
  i--
  dAtA[i] = 0x2a
 }
 if len(m.Signature) > 0 {
  i -= len(m.Signature)
  copy(dAtA[i:], m.Signature)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
  i--
  dAtA[i] = 0x22
 }
 if m.Sequence != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
  i--
  dAtA[i] = 0x18
 }
 if len(m.Coins) > 0 {
  for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Coins) > 0 {
  for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *Raw) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *Raw) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Raw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Raw) > 0 {
  i -= len(m.Raw)
  copy(dAtA[i:], m.Raw)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Raw)))
  i--
  dAtA[i] = 0x12
 }
 if len(m.Type) > 0 {
  i -= len(m.Type)
  copy(dAtA[i:], m.Type)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Type)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *SendCoins) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SendCoins) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Outputs) > 0 {
  for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x22
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x12
 if m.Gas != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Gas))
  i--
  dAtA[i] = 0x8
 }
 return len(dAtA) - i, nil
}

func (m *MintCoins) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *MintCoins) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if m.Delta != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Delta))
  i--
  dAtA[i] = 0x40
 }
 if m.Movable {
  i--
  if m.Movable {
   dAtA[i] = 1
  } else {
   dAtA[i] = 0
  }
  i--
  dAtA[i] = 0x38
 }
 if m.DecimalPoint != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.DecimalPoint))
  i--
  dAtA[i] = 0x30
 }
 if m.Amount != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Amount))
  i--
  dAtA[i] = 0x28
 }
 if len(m.Name) > 0 {
  i -= len(m.Name)
  copy(dAtA[i:], m.Name)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Outputs) > 0 {
  for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *SetSalePrice) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SetSalePrice) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetSalePrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 {
  size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x22
 if len(m.Name) > 0 {
  i -= len(m.Name)
  copy(dAtA[i:], m.Name)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *CreateAMC) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *CreateAMC) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAMC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Params) > 0 {
  for iNdEx := len(m.Params) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Params[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x2a
  }
 }
 if len(m.Name) > 0 {
  i -= len(m.Name)
  copy(dAtA[i:], m.Name)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *AMCParam) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *AMCParam) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AMCParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if m.Value != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Value))
  i--
  dAtA[i] = 0x10
 }
 if len(m.Key) > 0 {
  i -= len(m.Key)
  copy(dAtA[i:], m.Key)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Key)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *SetInAMC) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SetInAMC) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetInAMC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Sponsor) > 0 {
  i -= len(m.Sponsor)
  copy(dAtA[i:], m.Sponsor)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Sponsor)))
  i--
  dAtA[i] = 0x2a
 }
 if len(m.Referal) > 0 {
  i -= len(m.Referal)
  copy(dAtA[i:], m.Referal)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Referal)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *BuyInAMC) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *BuyInAMC) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyInAMC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 {
  size, err := m.Delta.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x32
 {
  size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x2a
 if len(m.Referal) > 0 {
  i -= len(m.Referal)
  copy(dAtA[i:], m.Referal)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Referal)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *RefundInAMC) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *RefundInAMC) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefundInAMC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Log) > 0 {
  i -= len(m.Log)
  copy(dAtA[i:], m.Log)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Log)))
  i--
  dAtA[i] = 0x2a
 }
 if len(m.Buy) > 0 {
  i -= len(m.Buy)
  copy(dAtA[i:], m.Buy)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Buy)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Name) > 0 {
  i -= len(m.Name)
  copy(dAtA[i:], m.Name)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *ChangeAddressWant) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *ChangeAddressWant) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeAddressWant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Outputs) > 0 {
  for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *ChangeAddressApply) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *ChangeAddressApply) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChangeAddressApply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Votes) > 0 {
  for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.PubKey) > 0 {
  i -= len(m.PubKey)
  copy(dAtA[i:], m.PubKey)
  i = encodeVarintTx(dAtA, i, uint64(len(m.PubKey)))
  i--
  dAtA[i] = 0x1a
 }
 if len(m.Signature) > 0 {
  i -= len(m.Signature)
  copy(dAtA[i:], m.Signature)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Signature)))
  i--
  dAtA[i] = 0x12
 }
 if len(m.Address) > 0 {
  i -= len(m.Address)
  copy(dAtA[i:], m.Address)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func (m *BuyNode) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *BuyNode) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 {
  size, err := m.Delta.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x42
 {
  size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x3a
 if m.Power != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Power))
  i--
  dAtA[i] = 0x30
 }
 if len(m.Validator) > 0 {
  i -= len(m.Validator)
  copy(dAtA[i:], m.Validator)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
  i--
  dAtA[i] = 0x2a
 }
 if len(m.Holder) > 0 {
  i -= len(m.Holder)
  copy(dAtA[i:], m.Holder)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Holder)))
  i--
  dAtA[i] = 0x22
 }
 if len(m.Outputs) > 0 {
  for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *SetInteresNode) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SetInteresNode) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetInteresNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 {
  size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x1a
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *SetNodeReward) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SetNodeReward) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetNodeReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 {
  size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0x1a
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *SetNewAge) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *SetNewAge) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetNewAge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Params) > 0 {
  for iNdEx := len(m.Params) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Params[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x1a
  }
 }
 if len(m.Inputs) > 0 {
  for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
   {
    size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
    if err != nil {
     return 0, err
    }
    i -= size
    i = encodeVarintTx(dAtA, i, uint64(size))
   }
   i--
   dAtA[i] = 0x12
  }
 }
 {
  size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
  if err != nil {
   return 0, err
  }
  i -= size
  i = encodeVarintTx(dAtA, i, uint64(size))
 }
 i--
 dAtA[i] = 0xa
 return len(dAtA) - i, nil
}

func (m *AgeParam) Marshal() (dAtA []byte, err error) {
 size := m.Size()
 dAtA = make([]byte, size)
 n, err := m.MarshalToSizedBuffer(dAtA[:size])
 if err != nil {
  return nil, err
 }
 return dAtA[:n], nil
}

func (m *AgeParam) MarshalTo(dAtA []byte) (int, error) {
 size := m.Size()
 return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AgeParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
 i := len(dAtA)
 _ = i
 var l int
 _ = l
 if len(m.Svalue) > 0 {
  i -= len(m.Svalue)
  copy(dAtA[i:], m.Svalue)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Svalue)))
  i--
  dAtA[i] = 0x1a
 }
 if m.Value != 0 {
  i = encodeVarintTx(dAtA, i, uint64(m.Value))
  i--
  dAtA[i] = 0x10
 }
 if len(m.Key) > 0 {
  i -= len(m.Key)
  copy(dAtA[i:], m.Key)
  i = encodeVarintTx(dAtA, i, uint64(len(m.Key)))
  i--
  dAtA[i] = 0xa
 }
 return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
 offset -= sovTx(v)
 base := offset
 for v >= 1<<7 {
  dAtA[offset] = uint8(v&0x7f | 0x80)
  v >>= 7
  offset++
 }
 dAtA[offset] = uint8(v)
 return base
}
func (m *Coin) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Name)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if m.Amount != 0 {
  n += 1 + sovTx(uint64(m.Amount))
 }
 return n
}

func (m *Input) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if len(m.Coins) > 0 {
  for _, e := range m.Coins {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if m.Sequence != 0 {
  n += 1 + sovTx(uint64(m.Sequence))
 }
 l = len(m.Signature)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.PubKey)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *Output) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if len(m.Coins) > 0 {
  for _, e := range m.Coins {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 return n
}

func (m *Raw) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Type)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Raw)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *SendCoins) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 if m.Gas != 0 {
  n += 1 + sovTx(uint64(m.Gas))
 }
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Outputs) > 0 {
  for _, e := range m.Outputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 return n
}

func (m *MintCoins) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Outputs) > 0 {
  for _, e := range m.Outputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Name)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if m.Amount != 0 {
  n += 1 + sovTx(uint64(m.Amount))
 }
 if m.DecimalPoint != 0 {
  n += 1 + sovTx(uint64(m.DecimalPoint))
 }
 if m.Movable {
  n += 2
 }
 if m.Delta != 0 {
  n += 1 + sovTx(uint64(m.Delta))
 }
 return n
}

func (m *SetSalePrice) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Name)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = m.Price.Size()
 n += 1 + l + sovTx(uint64(l))
 return n
}

func (m *CreateAMC) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Name)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if len(m.Params) > 0 {
  for _, e := range m.Params {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 return n
}

func (m *AMCParam) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Key)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if m.Value != 0 {
  n += 1 + sovTx(uint64(m.Value))
 }
 return n
}

func (m *SetInAMC) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Referal)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Sponsor)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *BuyInAMC) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Referal)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = m.Value.Size()
 n += 1 + l + sovTx(uint64(l))
 l = m.Delta.Size()
 n += 1 + l + sovTx(uint64(l))
 return n
}

func (m *RefundInAMC) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Name)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Buy)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Log)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *ChangeAddressWant) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Outputs) > 0 {
  for _, e := range m.Outputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 return n
}

func (m *ChangeAddressApply) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Votes) > 0 {
  for _, e := range m.Votes {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *Vote) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Address)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Signature)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.PubKey)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func (m *BuyNode) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Outputs) > 0 {
  for _, e := range m.Outputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = len(m.Holder)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 l = len(m.Validator)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if m.Power != 0 {
  n += 1 + sovTx(uint64(m.Power))
 }
 l = m.Value.Size()
 n += 1 + l + sovTx(uint64(l))
 l = m.Delta.Size()
 n += 1 + l + sovTx(uint64(l))
 return n
}

func (m *SetInteresNode) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = m.Value.Size()
 n += 1 + l + sovTx(uint64(l))
 return n
}

func (m *SetNodeReward) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 l = m.Value.Size()
 n += 1 + l + sovTx(uint64(l))
 return n
}

func (m *SetNewAge) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = m.Fee.Size()
 n += 1 + l + sovTx(uint64(l))
 if len(m.Inputs) > 0 {
  for _, e := range m.Inputs {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 if len(m.Params) > 0 {
  for _, e := range m.Params {
   l = e.Size()
   n += 1 + l + sovTx(uint64(l))
  }
 }
 return n
}

func (m *AgeParam) Size() (n int) {
 if m == nil {
  return 0
 }
 var l int
 _ = l
 l = len(m.Key)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 if m.Value != 0 {
  n += 1 + sovTx(uint64(m.Value))
 }
 l = len(m.Svalue)
 if l > 0 {
  n += 1 + l + sovTx(uint64(l))
 }
 return n
}

func sovTx(x uint64) (n int) {
 return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
 return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Coin) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: Coin: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Name = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 2:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
   }
   m.Amount = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Amount |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *Input) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: Input: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Coins = append(m.Coins, Coin{})
   if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
   }
   m.Sequence = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Sequence |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
   if m.Signature == nil {
    m.Signature = []byte{}
   }
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
   if m.PubKey == nil {
    m.PubKey = []byte{}
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *Output) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: Output: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Coins = append(m.Coins, Coin{})
   if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *Raw) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: Raw: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: Raw: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Type = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Raw = append(m.Raw[:0], dAtA[iNdEx:postIndex]...)
   if m.Raw == nil {
    m.Raw = []byte{}
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SendCoins) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SendCoins: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SendCoins: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
   }
   m.Gas = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Gas |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Outputs = append(m.Outputs, Output{})
   if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *MintCoins) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: MintCoins: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: MintCoins: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Outputs = append(m.Outputs, Output{})
   if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Name = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 5:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
   }
   m.Amount = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Amount |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 6:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field DecimalPoint", wireType)
   }
   m.DecimalPoint = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.DecimalPoint |= uint32(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 7:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Movable", wireType)
   }
   var v int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    v |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   m.Movable = bool(v != 0)
  case 8:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
   }
   m.Delta = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Delta |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SetSalePrice) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SetSalePrice: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SetSalePrice: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Name = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *CreateAMC) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: CreateAMC: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: CreateAMC: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Name = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Params = append(m.Params, AMCParam{})
   if err := m.Params[len(m.Params)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *AMCParam) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: AMCParam: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: AMCParam: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Key = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 2:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   m.Value = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Value |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SetInAMC) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SetInAMC: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SetInAMC: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Referal", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Referal = append(m.Referal[:0], dAtA[iNdEx:postIndex]...)
   if m.Referal == nil {
    m.Referal = []byte{}
   }
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Sponsor", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Sponsor = append(m.Sponsor[:0], dAtA[iNdEx:postIndex]...)
   if m.Sponsor == nil {
    m.Sponsor = []byte{}
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *BuyInAMC) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: BuyInAMC: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: BuyInAMC: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Referal", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Referal = append(m.Referal[:0], dAtA[iNdEx:postIndex]...)
   if m.Referal == nil {
    m.Referal = []byte{}
   }
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 6:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Delta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *RefundInAMC) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: RefundInAMC: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: RefundInAMC: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Name = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Buy", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Buy = append(m.Buy[:0], dAtA[iNdEx:postIndex]...)
   if m.Buy == nil {
    m.Buy = []byte{}
   }
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Log = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *ChangeAddressWant) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: ChangeAddressWant: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: ChangeAddressWant: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Outputs = append(m.Outputs, Output{})
   if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *ChangeAddressApply) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: ChangeAddressApply: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: ChangeAddressApply: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Votes = append(m.Votes, Vote{})
   if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *Vote) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: Vote: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
   if m.Address == nil {
    m.Address = []byte{}
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
   if m.Signature == nil {
    m.Signature = []byte{}
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
   if m.PubKey == nil {
    m.PubKey = []byte{}
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *BuyNode) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: BuyNode: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: BuyNode: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Outputs = append(m.Outputs, Output{})
   if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 4:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Holder", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Holder = append(m.Holder[:0], dAtA[iNdEx:postIndex]...)
   if m.Holder == nil {
    m.Holder = []byte{}
   }
   iNdEx = postIndex
  case 5:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
   }
   var byteLen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    byteLen |= int(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
   if byteLen < 0 {
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + byteLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
   if m.Validator == nil {
    m.Validator = []byte{}
   }
   iNdEx = postIndex
  case 6:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
   }
   m.Power = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Power |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 7:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 8:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Delta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SetInteresNode) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SetInteresNode: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SetInteresNode: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SetNodeReward) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SetNodeReward: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SetNodeReward: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *SetNewAge) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: SetNewAge: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: SetNewAge: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 2:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Inputs = append(m.Inputs, Input{})
   if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
   }
   var msglen int
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + msglen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Params = append(m.Params, AgeParam{})
   if err := m.Params[len(m.Params)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
    return err
   }
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func (m *AgeParam) Unmarshal(dAtA []byte) error {
 l := len(dAtA)
 iNdEx := 0
 for iNdEx < l {
  preIndex := iNdEx
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return ErrIntOverflowTx
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
   return fmt.Errorf("proto: AgeParam: wiretype end group for non-group")
  }
  if fieldNum <= 0 {
   return fmt.Errorf("proto: AgeParam: illegal tag %d (wire type %d)", fieldNum, wire)
  }
  switch fieldNum {
  case 1:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Key = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  case 2:
   if wireType != 0 {
    return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
   }
   m.Value = 0
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
    }
    if iNdEx >= l {
     return io.ErrUnexpectedEOF
    }
    b := dAtA[iNdEx]
    iNdEx++
    m.Value |= int64(b&0x7F) << shift
    if b < 0x80 {
     break
    }
   }
  case 3:
   if wireType != 2 {
    return fmt.Errorf("proto: wrong wireType = %d for field Svalue", wireType)
   }
   var stringLen uint64
   for shift := uint(0); ; shift += 7 {
    if shift >= 64 {
     return ErrIntOverflowTx
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
    return ErrInvalidLengthTx
   }
   postIndex := iNdEx + intStringLen
   if postIndex < 0 {
    return ErrInvalidLengthTx
   }
   if postIndex > l {
    return io.ErrUnexpectedEOF
   }
   m.Svalue = string(dAtA[iNdEx:postIndex])
   iNdEx = postIndex
  default:
   iNdEx = preIndex
   skippy, err := skipTx(dAtA[iNdEx:])
   if err != nil {
    return err
   }
   if skippy < 0 {
    return ErrInvalidLengthTx
   }
   if (iNdEx + skippy) < 0 {
    return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
 l := len(dAtA)
 iNdEx := 0
 depth := 0
 for iNdEx < l {
  var wire uint64
  for shift := uint(0); ; shift += 7 {
   if shift >= 64 {
    return 0, ErrIntOverflowTx
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
     return 0, ErrIntOverflowTx
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
     return 0, ErrIntOverflowTx
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
    return 0, ErrInvalidLengthTx
   }
   iNdEx += length
  case 3:
   depth++
  case 4:
   if depth == 0 {
    return 0, ErrUnexpectedEndOfGroupTx
   }
   depth--
  case 5:
   iNdEx += 4
  default:
   return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
  }
  if iNdEx < 0 {
   return 0, ErrInvalidLengthTx
  }
  if depth == 0 {
   return iNdEx, nil
  }
 }
 return 0, io.ErrUnexpectedEOF
}

var (
 ErrInvalidLengthTx = fmt.Errorf("proto: negative length found during unmarshaling")
 ErrIntOverflowTx = fmt.Errorf("proto: integer overflow")
 ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
