package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "gpsync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"

type TestI64MapSync struct {
	id             int64
	idINDEX        int
	addition       string
	additionINDEX  int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewTestI64MapSync() *TestI64MapSync {
	return &TestI64MapSync{
		idINDEX:        0,
		additionINDEX:  1,
		dirtyFieldMark: make([]uint8, 1),
	}
}
func (x *TestI64MapSync) Clear() *TestI64MapSync {
	x.SetId(0)
	x.SetAddition("")
	return x
}
func (x *TestI64MapSync) CopyToPb(r *TestI64Map) *TestI64MapSync {
	r.SetId(x.id)
	r.SetAddition(x.addition)
	return x
}
func (x *TestI64MapSync) CopyFromPb(r *TestI64Map) *TestI64MapSync {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
	return x
}
func (x *TestI64MapSync) MergeDirtyFromPb(r *TestI64Map) {
	if r.Id != nil {
		x.SetId(*r.Id)
	}
	if r.Addition != nil {
		x.SetAddition(*r.Addition)
	}
}
func (x *TestI64MapSync) MergeDirtyFromBytes(buf []byte) *TestI64MapSync {
	fds := syncdep.PreParseProtoBytes(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetId(int64(rawF.Value.(uint64)))
		case 2:
			x.SetAddition(syncdep.Bys2Str(rawF.Value.([]byte)))
		}
	}
	return x
}
func (x *TestI64MapSync) MergeDirtyToBytes() []byte {
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
func (x *TestI64MapSync) MergeDirtyToPb(r *TestI64Map) {
	if x.isIdDirty() {
		r.SetId(x.id)
	}
	if x.isAdditionDirty() {
		r.SetAddition(x.addition)
	}
}
func (x *TestI64MapSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *TestI64MapSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *TestI64MapSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *TestI64MapSync) FlushDirty(dirty bool) {
	if dirty || x.isIdDirty() {
		x.setIdDirty(dirty, true)
	}
	if dirty || x.isAdditionDirty() {
		x.setAdditionDirty(dirty, true)
	}
}
func (x *TestI64MapSync) setIdDirty(dirty bool, recur bool) {
	x.SetDirty(x.idINDEX, dirty, x)
}
func (x *TestI64MapSync) isIdDirty() bool {
	idx := x.idINDEX >> 3
	off := x.idINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestI64MapSync) setAdditionDirty(dirty bool, recur bool) {
	x.SetDirty(x.additionINDEX, dirty, x)
}
func (x *TestI64MapSync) isAdditionDirty() bool {
	idx := x.additionINDEX >> 3
	off := x.additionINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *TestI64MapSync) Key() interface{} {
	return x.id
}
func (x *TestI64MapSync) SetKey(v interface{}) {
	if x.parent != nil {
		if _, ok := x.parent.(*syncdep.MapSync[int64, *TestI64MapSync]); ok {
			panic("TestI64MapSync key can not set")
		}
	}
	x.id = v.(int64)
}
func (x *TestI64MapSync) GetId() int64 {
	return x.id
}
func (x *TestI64MapSync) SetId(v int64) *TestI64MapSync {
	if x.id == v {
		return x
	}
	x.id = v
	x.setIdDirty(true, false)
	return x
}
func (x *TestI64MapSync) GetAddition() string {
	return x.addition
}
func (x *TestI64MapSync) SetAddition(v string) *TestI64MapSync {
	if x.addition == v {
		return x
	}
	x.addition = v
	x.setAdditionDirty(true, false)
	return x
}
func (xs *TestI64Map) SetId(v int64) {
	xs.Id = &v
}
func (xs *TestI64Map) SetAddition(v string) {
	xs.Addition = &v
}
