package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "github.com/yaoguangduan/protosync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"

type IntroDetailSync struct {
	address        string
	addressINDEX   int
	money          int32
	moneyINDEX     int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewIntroDetailSync() *IntroDetailSync {
	return &IntroDetailSync{
		addressINDEX:   0,
		moneyINDEX:     1,
		dirtyFieldMark: make([]uint8, 1),
	}
}
func (x *IntroDetailSync) Clear() *IntroDetailSync {
	x.SetAddress("")
	x.SetMoney(0)
	return x
}
func (x *IntroDetailSync) CopyToPb(r *IntroDetail) *IntroDetailSync {
	r.SetAddress(x.address)
	r.SetMoney(x.money)
	return x
}
func (x *IntroDetailSync) CopyFromPb(r *IntroDetail) *IntroDetailSync {
	if r.Address != nil {
		x.SetAddress(*r.Address)
	}
	if r.Money != nil {
		x.SetMoney(*r.Money)
	}
	return x
}
func (x *IntroDetailSync) MergeDirtyFromPb(r *IntroDetail) {
	if r.Address != nil {
		x.SetAddress(*r.Address)
	}
	if r.Money != nil {
		x.SetMoney(*r.Money)
	}
}
func (x *IntroDetailSync) MergeDirtyFromBytes(buf []byte) *IntroDetailSync {
	fds := syncdep.PreUnmarshal(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetAddress(syncdep.Bys2Str(rawF.Value.([]byte)))
		case 2:
			x.SetMoney(int32(rawF.Value.(uint64)))
		}
	}
	return x
}
func (x *IntroDetailSync) MergeDirtyToBytes() []byte {
	var buf []byte
	if x.isAddressDirty() {
		buf = protowire.AppendTag(buf, 1, 2)
		buf = syncdep.AppendFieldValue(buf, x.address)
	}
	if x.isMoneyDirty() {
		buf = protowire.AppendTag(buf, 2, 0)
		buf = syncdep.AppendFieldValue(buf, x.money)
	}
	return buf
}
func (x *IntroDetailSync) MergeDirtyToPb(r *IntroDetail) {
	if x.isAddressDirty() {
		r.SetAddress(x.address)
	}
	if x.isMoneyDirty() {
		r.SetMoney(x.money)
	}
}
func (x *IntroDetailSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *IntroDetailSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *IntroDetailSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *IntroDetailSync) FlushDirty(dirty bool) {
	if dirty || x.isAddressDirty() {
		x.setAddressDirty(dirty, true)
	}
	if dirty || x.isMoneyDirty() {
		x.setMoneyDirty(dirty, true)
	}
}
func (x *IntroDetailSync) setAddressDirty(dirty bool, recur bool) {
	x.SetDirty(x.addressINDEX, dirty, x)
}
func (x *IntroDetailSync) isAddressDirty() bool {
	idx := x.addressINDEX >> 3
	off := x.addressINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *IntroDetailSync) setMoneyDirty(dirty bool, recur bool) {
	x.SetDirty(x.moneyINDEX, dirty, x)
}
func (x *IntroDetailSync) isMoneyDirty() bool {
	idx := x.moneyINDEX >> 3
	off := x.moneyINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *IntroDetailSync) Key() interface{} {
	return nil
}
func (x *IntroDetailSync) SetKey(v interface{}) {
	return
}
func (x *IntroDetailSync) GetAddress() string {
	return x.address
}
func (x *IntroDetailSync) SetAddress(v string) *IntroDetailSync {
	if x.address == v {
		return x
	}
	x.address = v
	x.setAddressDirty(true, false)
	return x
}
func (x *IntroDetailSync) GetMoney() int32 {
	return x.money
}
func (x *IntroDetailSync) SetMoney(v int32) *IntroDetailSync {
	if x.money == v {
		return x
	}
	x.money = v
	x.setMoneyDirty(true, false)
	return x
}
func (xs *IntroDetail) SetAddress(v string) {
	xs.Address = &v
}
func (xs *IntroDetail) SetMoney(v int32) {
	xs.Money = &v
}
func (xs *IntroDetail) Unmarshal(buf []byte) error {
	for len(buf) > 0 {
		number, _, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return protowire.ParseError(n)
		}
		buf = buf[n:]
		switch number {
		case 1:
			v, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				return protowire.ParseError(n)
			}
			buf = buf[n:]
			xs.SetAddress(syncdep.Bys2Str(v))
			break
		case 2:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return protowire.ParseError(n)
			}
			buf = buf[n:]
			xs.SetMoney(int32(v))
			break
		}
	}
	return nil
}
func (xs *IntroDetail) Marshal() []byte {
	var buf []byte
	if xs.Address != nil {
		buf = protowire.AppendTag(buf, 1, protowire.BytesType)
		buf = protowire.AppendString(buf, *xs.Address)
	}
	if xs.Money != nil {
		buf = protowire.AppendTag(buf, 2, protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*xs.Money))
	}
	return buf
}
