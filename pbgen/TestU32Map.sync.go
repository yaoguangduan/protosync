package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "gpsync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"

type TestU32MapSync struct {
	id             uint32
	idINDEX        int
	addition       string
	additionINDEX  int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewTestU32MapSync() *TestU32MapSync {
	return &TestU32MapSync{
		idINDEX:        0,
		additionINDEX:  1,
		dirtyFieldMark: make([]uint8, 1),
	}
}
func (x *TestU32MapSync) Clear() *TestU32MapSync {
	x.SetId(0)
	x.SetAddition("")
	return x
}
func (x *TestU32MapSync) CopyToPb(r *TestU32Map) *TestU32MapSync {
	r.SetId(x.id)
	r.SetAddition(x.addition)
	return x
}
func (x *TestU32MapSync) CopyFromPb(r *TestU32Map) *TestU32MapSync {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
	return x
}
func (x *TestU32MapSync) MergeDirtyFromPb(r *TestU32Map) {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
}
func (x *TestU32MapSync) MergeDirtyFromBytes(buf []byte) *TestU32MapSync {
	fds := syncdep.PreParseProtoBytes(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetId(uint32(rawF.Value.(uint64)))
		case 2:
			x.SetAddition(syncdep.Bys2Str(rawF.Value.([]byte)))
		}
	}
	return x
}
func (x *TestU32MapSync) MergeDirtyToBytes() []byte {
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
func (x *TestU32MapSync) MergeDirtyToPb(r *TestU32Map) {
	if x.isIdDirty() {
		r.SetId(x.id)
	}
	if x.isAdditionDirty() {
		r.SetAddition(x.addition)
	}
}
func (x *TestU32MapSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *TestU32MapSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *TestU32MapSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *TestU32MapSync) FlushDirty(dirty bool) {
	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}
	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}
}
func (x *TestU32MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(x.idINDEX, dirty, x)
}
func (x *TestU32MapSync) isIdDirty() bool {
	idx := x.idINDEX >> 3
	off := x.idINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestU32MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(x.additionINDEX, dirty, x)
}
func (x *TestU32MapSync) isAdditionDirty() bool {
	idx := x.additionINDEX >> 3
	off := x.additionINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestU32MapSync) Key() interface{} {
	return x.id
}
func (x *TestU32MapSync) SetKey(v interface{}) {
	if x.parent != nil {
		if _, ok := x.parent.(*syncdep.MapSync[uint32, *TestU32MapSync]); ok {
			panic("TestU32MapSync key can not set")
		}
	}
	x.id = v.(uint32)
}
func (x *TestU32MapSync) GetId() uint32 {
	return x.id
}
func (x *TestU32MapSync) SetId(v uint32) *TestU32MapSync {
	if x.id == v {
		return x
	}
	x.id = v
	x.setIdDirty(true, false)
	return x
}
func (x *TestU32MapSync) GetAddition() string {
	return x.addition
}
func (x *TestU32MapSync) SetAddition(v string) *TestU32MapSync {
	if x.addition == v {
		return x
	}
	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}
func (xs *TestU32Map) SetId(v uint32) {
	xs.Id = &v
}
func (xs *TestU32Map) SetAddition(v string) {
	xs.Addition = &v
}
