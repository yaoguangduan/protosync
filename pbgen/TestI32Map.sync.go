package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "github.com/yaoguangduan/protosync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"

type TestI32MapSync struct {
	id             int32
	idINDEX        int
	addition       string
	additionINDEX  int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewTestI32MapSync() *TestI32MapSync {
	return &TestI32MapSync{
		idINDEX:        0,
		additionINDEX:  1,
		dirtyFieldMark: make([]uint8, 1),
	}
}
func (x *TestI32MapSync) Clear() *TestI32MapSync {
	x.SetId(0)
	x.SetAddition("")
	return x
}
func (x *TestI32MapSync) CopyToPb(r *TestI32Map) *TestI32MapSync {
	r.SetId(x.id)
	r.SetAddition(x.addition)
	return x
}
func (x *TestI32MapSync) CopyFromPb(r *TestI32Map) *TestI32MapSync {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
	return x
}
func (x *TestI32MapSync) MergeDirtyFromPb(r *TestI32Map) {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
}
func (x *TestI32MapSync) MergeDirtyFromBytes(buf []byte) *TestI32MapSync {
	fds := syncdep.PreParseProtoBytes(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetId(int32(rawF.Value.(uint64)))
		case 2:
			x.SetAddition(syncdep.Bys2Str(rawF.Value.([]byte)))
		}
	}
	return x
}
func (x *TestI32MapSync) MergeDirtyToBytes() []byte {
	var buf []byte
	if x.isIdDirty() {
		buf = protowire.AppendTag(buf, 1, 0)
		buf = syncdep.AppendFieldValue(buf, x.id)
	}
	if x.isAdditionDirty() {
		buf = protowire.AppendTag(buf, 2, 2)
		buf = syncdep.AppendFieldValue(buf, x.addition)
	}
	return buf
}
func (x *TestI32MapSync) MergeDirtyToPb(r *TestI32Map) {
	if x.isIdDirty() {
		r.SetId(x.id)
	}
	if x.isAdditionDirty() {
		r.SetAddition(x.addition)
	}
}
func (x *TestI32MapSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *TestI32MapSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *TestI32MapSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *TestI32MapSync) FlushDirty(dirty bool) {
	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}
	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}
}
func (x *TestI32MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(x.idINDEX, dirty, x)
}
func (x *TestI32MapSync) isIdDirty() bool {
	idx := x.idINDEX >> 3
	off := x.idINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestI32MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(x.additionINDEX, dirty, x)
}
func (x *TestI32MapSync) isAdditionDirty() bool {
	idx := x.additionINDEX >> 3
	off := x.additionINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestI32MapSync) Key() interface{} {
	return x.id
}
func (x *TestI32MapSync) SetKey(v interface{}) {
	if x.parent != nil {
		if _, ok := x.parent.(*syncdep.MapSync[int32, *TestI32MapSync]); ok {
			panic("TestI32MapSync key can not set")
		}
	}
	x.id = v.(int32)
}
func (x *TestI32MapSync) GetId() int32 {
	return x.id
}
func (x *TestI32MapSync) SetId(v int32) *TestI32MapSync {
	if x.id == v {
		return x
	}
	x.id = v
	x.setIdDirty(true, false)
	return x
}
func (x *TestI32MapSync) GetAddition() string {
	return x.addition
}
func (x *TestI32MapSync) SetAddition(v string) *TestI32MapSync {
	if x.addition == v {
		return x
	}
	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}
func (xs *TestI32Map) SetId(v int32) {
	xs.Id = &v
}
func (xs *TestI32Map) SetAddition(v string) {
	xs.Addition = &v
}
func (xs *TestI32Map) Unmarshal(buf []byte) error {
	for len(buf) > 0 {
		number, _, n := protowire.ConsumeTag(buf)
		if n < 0 {
			return protowire.ParseError(n)
		}
		buf = buf[n:]
		switch number {
		case 1:
			v, n := protowire.ConsumeVarint(buf)
			if n < 0 {
				return protowire.ParseError(n)
			}
			buf = buf[n:]
			xs.SetId(int32(v))
			break
		case 2:
			v, n := protowire.ConsumeBytes(buf)
			if n < 0 {
				return protowire.ParseError(n)
			}
			buf = buf[n:]
			xs.SetAddition(syncdep.Bys2Str(v))
			break
		}
	}
	return nil
}
func (xs *TestI32Map) Marshal() []byte {
	var buf []byte
	if xs.Id != nil {
		buf = protowire.AppendTag(buf, 1, protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*xs.Id))
	}
	if xs.Addition != nil {
		buf = protowire.AppendTag(buf, 2, protowire.BytesType)
		buf = protowire.AppendString(buf, *xs.Addition)
	}
	return buf
}
