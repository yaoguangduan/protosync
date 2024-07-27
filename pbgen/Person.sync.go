package pbgen

//this file is generated by gsync, do not modify it manually !!!

import "gpsync/syncdep"
import "google.golang.org/protobuf/encoding/protowire"
import "slices"

type PersonSync struct {
	age            int32
	ageINDEX       int
	vipLevel       VipLevel
	vipLevelINDEX  int
	name           string
	nameINDEX      int
	actions        *syncdep.MapSync[string, *ActionInfoSync]
	actionsINDEX   int
	favor          *syncdep.ArraySync[string]
	favorINDEX     int
	loveSeq        *syncdep.ArraySync[ColorType]
	loveSeqINDEX   int
	isGirl         bool
	isGirlINDEX    int
	detail         *IntroDetailSync
	detailINDEX    int
	data           []byte
	dataINDEX      int
	dirtyFieldMark []uint8
	parent         syncdep.Sync
	indexInParent  int
}

func NewPersonSync() *PersonSync {
	return &PersonSync{
		ageINDEX:       0,
		vipLevelINDEX:  1,
		nameINDEX:      2,
		actionsINDEX:   3,
		favorINDEX:     4,
		loveSeqINDEX:   5,
		isGirlINDEX:    6,
		detailINDEX:    7,
		dataINDEX:      8,
		dirtyFieldMark: make([]uint8, 2),
	}
}
func (x *PersonSync) Clear() *PersonSync {
	x.SetAge(0)
	x.SetVipLevel(VipLevel_Level1)
	x.SetName("")
	if x.actions != nil {
		x.actions.Clear()
	}
	if x.favor != nil {
		x.favor.Clear()
	}
	if x.loveSeq != nil {
		x.loveSeq.Clear()
	}
	x.SetIsGirl(false)
	if x.detail != nil {
		x.detail.Clear()
	}
	x.SetData(make([]byte, 0))
	return x
}
func (x *PersonSync) CopyToPb(r *Person) *PersonSync {
	r.SetAge(x.age)
	r.SetVipLevel(x.vipLevel)
	r.SetName(x.name)
	if x.actions != nil && x.actions.Len() > 0 {
		tmp := make(map[string]*ActionInfo)
		x.actions.Each(func(k string, v *ActionInfoSync) bool {
			tmpV := &ActionInfo{}
			v.CopyToPb(tmpV)
			tmp[k] = tmpV
			return true
		})
		r.SetActions(tmp)
	}
	if x.favor != nil && x.favor.Len() > 0 {
		r.SetFavor(x.favor.ValueView())
	}
	if x.loveSeq != nil && x.loveSeq.Len() > 0 {
		r.SetLoveSeq(x.loveSeq.ValueView())
	}
	r.SetIsGirl(x.isGirl)
	if x.detail != nil {
		tmp := &IntroDetail{}
		x.detail.CopyToPb(tmp)
		r.SetDetail(tmp)
	}
	r.SetData(slices.Clone(x.data))
	return x
}
func (x *PersonSync) CopyFromPb(r *Person) *PersonSync {
	if r.Age != nil {
		x.SetAge(*r.Age)
	}
	if r.VipLevel != nil {
		x.SetVipLevel(*r.VipLevel)
	}
	if r.Name != nil {
		x.SetName(*r.Name)
	}
	for k, v := range r.Actions {
		if v != nil {
			vv := NewActionInfoSync()
			vv.CopyFromPb(v)
			x.GetActions().Put(k, vv)
		}
	}
	if r.Favor != nil {
		x.GetFavor().AddAll(r.Favor)
	}
	if r.LoveSeq != nil {
		x.GetLoveSeq().AddAll(r.LoveSeq)
	}
	if r.IsGirl != nil {
		x.SetIsGirl(*r.IsGirl)
	}
	if r.Detail != nil {
		x.GetDetail().CopyFromPb(r.Detail)
	}
	if len(r.Data) > 0 {
		x.SetData(slices.Clone(r.Data))
	}
	return x
}
func (x *PersonSync) MergeDirtyFromPb(r *Person) {
	if r.Age != nil {
		x.SetAge(*r.Age)
	}
	if r.VipLevel != nil {
		x.SetVipLevel(*r.VipLevel)
	}
	if r.Name != nil {
		x.SetName(*r.Name)
	}
	if x.actions != nil {
		x.GetActions().RemoveAll(r.ActionsDeleted)
	}
	for k, v := range r.Actions {
		var tmp = x.GetActions().Get(k)
		if tmp == nil {
			tmp = NewActionInfoSync()
			tmp.MergeDirtyFromPb(v)
			x.GetActions().Put(k, tmp)
		} else {
			tmp.MergeDirtyFromPb(v)
		}
	}
	if len(r.Favor) > 0 || r.FavorCleared {
		x.GetFavor().Clear()
		x.favor.AddAll(r.Favor)
	}
	if len(r.LoveSeq) > 0 || r.LoveSeqCleared {
		x.GetLoveSeq().Clear()
		x.loveSeq.AddAll(r.LoveSeq)
	}
	if r.IsGirl != nil {
		x.SetIsGirl(*r.IsGirl)
	}
	if r.Detail != nil {
		x.GetDetail().MergeDirtyFromPb(r.Detail)
	}
	if len(r.Data) > 0 {
		x.SetData(slices.Clone(r.Data))
	}
}
func (x *PersonSync) MergeDirtyFromBytes(buf []byte) *PersonSync {
	fds := syncdep.PreParseProtoBytes(buf)
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1004:
			if x.actions != nil {
				x.actions.Remove(syncdep.Bys2Str(rawF.Value.([]byte)))
			}
		case 2005:
			if rawF.Value.(uint64) > 0 {
				x.favor.Clear()
			}
		case 2006:
			if rawF.Value.(uint64) > 0 {
				x.loveSeq.Clear()
			}
		}
	}
	for _, rawF := range fds.Values {
		switch rawF.Number {
		case 1:
			x.SetAge(int32(rawF.Value.(uint64)))
		case 2:
			x.SetVipLevel(VipLevel(rawF.Value.(uint64)))
		case 3:
			x.SetName(syncdep.Bys2Str(rawF.Value.([]byte)))
		case 4:
			mapKV := syncdep.PreParseProtoBytes(rawF.Value.([]byte)).Values
			k := syncdep.GetMapKey[string](&mapKV[0])
			var tmp = x.GetActions().Get(k)
			if tmp == nil {
				tmp = NewActionInfoSync()
			}
			tmp.MergeDirtyFromBytes(mapKV[1].Value.([]byte))
			x.GetActions().Put(k, tmp)
		case 5:
			x.GetFavor().Add(syncdep.Bys2Str(rawF.Value.([]byte)))
		case 6:
			tmp := rawF.Value.([]byte)
			for len(tmp) > 0 {
				val, n := protowire.ConsumeVarint(tmp)
				if n < 0 {
					panic(n)
				}
				tmp = tmp[n:]
				x.GetLoveSeq().Add(ColorType(val))
			}
		case 7:
			x.SetIsGirl(rawF.Value.(uint64) > 0)
		case 8:
			x.GetDetail().MergeDirtyFromBytes(rawF.Value.([]byte))
		case 9:
			x.SetData(rawF.Value.([]byte))
		}
	}
	return x
}
func (x *PersonSync) MergeDirtyToBytes() []byte {
	var buf []byte
	if x.isAgeDirty() {
		buf = protowire.AppendTag(buf, 1, 0)
		buf = syncdep.AppendFieldValue(buf, x.age)
	}
	if x.isVipLevelDirty() {
		buf = protowire.AppendTag(buf, 2, 0)
		buf = syncdep.AppendFieldValue(buf, int32(x.vipLevel))
	}
	if x.isNameDirty() {
		buf = protowire.AppendTag(buf, 3, 2)
		buf = syncdep.AppendFieldValue(buf, x.name)
	}
	if x.isActionsDirty() {
		if len(x.actions.Deleted()) > 0 {
			deleted := x.actions.Deleted()
			for del := range deleted {
				buf = protowire.AppendTag(buf, 1004, protowire.BytesType)
				buf = protowire.AppendString(buf, del)
			}
		}
		if x.actions.Len() > 0 {
			x.actions.Each(func(k string, v *ActionInfoSync) bool {
				if !x.actions.ContainDirtied(k) {
					return true
				}
				buf = syncdep.AppendMapFieldKeyValue(buf, 4, k, v.MergeDirtyToBytes())
				return true
			})
		}
	}
	if x.isFavorDirty() {
		if x.favor != nil && x.favor.Len() > 0 {
			buf = protowire.AppendTag(buf, 2005, protowire.VarintType)
			buf = syncdep.AppendFieldValue(buf, false)
			x.favor.Each(func(i int, v string) bool {
				buf = protowire.AppendTag(buf, 5, 2)
				buf = protowire.AppendString(buf, v)
				return true
			})
		} else {
			buf = protowire.AppendTag(buf, 2005, protowire.VarintType)
			buf = syncdep.AppendFieldValue(buf, true)
		}
	}
	if x.isLoveSeqDirty() {
		if x.loveSeq != nil && x.loveSeq.Len() > 0 {
			buf = protowire.AppendTag(buf, 2006, protowire.VarintType)
			buf = syncdep.AppendFieldValue(buf, false)
			var packedBuf []byte
			x.loveSeq.Each(func(i int, v ColorType) bool {
				packedBuf = syncdep.AppendFieldValue(packedBuf, int32(v))
				return true
			})
			buf = protowire.AppendTag(buf, 6, 2)
			buf = protowire.AppendBytes(buf, packedBuf)
		} else {
			buf = protowire.AppendTag(buf, 2006, protowire.VarintType)
			buf = syncdep.AppendFieldValue(buf, true)
		}
	}
	if x.isIsGirlDirty() {
		buf = protowire.AppendTag(buf, 7, 0)
		buf = syncdep.AppendFieldValue(buf, x.isGirl)
	}
	if x.isDetailDirty() {
		if x.detail != nil {
			bytes := x.detail.MergeDirtyToBytes()
			buf = protowire.AppendTag(buf, 8, 2)
			buf = syncdep.AppendFieldValue(buf, bytes)
		}
	}
	if x.isDataDirty() {
		buf = protowire.AppendTag(buf, 9, 2)
		buf = syncdep.AppendFieldValue(buf, x.data)
	}
	return buf
}
func (x *PersonSync) MergeDirtyToPb(r *Person) {
	if x.isAgeDirty() {
		r.SetAge(x.age)
	}
	if x.isVipLevelDirty() {
		r.SetVipLevel(x.vipLevel)
	}
	if x.isNameDirty() {
		r.SetName(x.name)
	}
	if x.isActionsDirty() {
		updated := make([]string, 0)
		if r.Actions != nil {
			for k := range r.Actions {
				if x.actions.ContainDeleted(k) {
					delete(r.Actions, k)
				}
				if x.actions.ContainDirtied(k) {
					updated = append(updated, k)
					tmp := x.actions.Get(k)
					if r.Actions[k] == nil {
						r.Actions[k] = &ActionInfo{}
					}
					tmp.MergeDirtyToPb(r.Actions[k])
				}
			}
		} else {
			r.Actions = make(map[string]*ActionInfo)
		}
		for k := range x.actions.Dirtied() {
			if !slices.Contains(updated, k) {
				tmp := x.actions.Get(k)
				if r.Actions[k] == nil {
					r.Actions[k] = &ActionInfo{}
				}
				tmp.MergeDirtyToPb(r.Actions[k])
			}
		}
		if r.ActionsDeleted == nil && len(x.actions.Deleted()) > 0 {
			r.ActionsDeleted = make([]string, 0)
		}
		for k := range x.actions.Deleted() {
			if !slices.Contains(r.ActionsDeleted, k) {
				r.ActionsDeleted = append(r.ActionsDeleted, k)
			}
		}
	}
	if x.isFavorDirty() {
		count := x.favor.Len()
		r.Favor = make([]string, 0)
		if count > 0 {
			r.FavorCleared = false
			r.Favor = append(r.Favor, x.favor.ValueView()...)
		} else {
			r.FavorCleared = true
		}
	}
	if x.isLoveSeqDirty() {
		count := x.loveSeq.Len()
		r.LoveSeq = make([]ColorType, 0)
		if count > 0 {
			r.LoveSeqCleared = false
			r.LoveSeq = append(r.LoveSeq, x.loveSeq.ValueView()...)
		} else {
			r.LoveSeqCleared = true
		}
	}
	if x.isIsGirlDirty() {
		r.SetIsGirl(x.isGirl)
	}
	if x.isDetailDirty() {
		if r.Detail == nil {
			r.Detail = &IntroDetail{}
		}
		x.detail.MergeDirtyToPb(r.Detail)
	}
	if x.isDataDirty() {
		r.SetData(slices.Clone(x.data))
	}
}
func (x *PersonSync) SetDirty(index int, dirty bool, sync syncdep.Sync) {
	idx := index >> 3
	off := index & 7
	if dirty {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] | (1 << off)
		x.SetParentDirty()
	} else {
		x.dirtyFieldMark[idx] = x.dirtyFieldMark[idx] & ^(1 << off)
	}
}
func (x *PersonSync) SetParentDirty() {
	if x.parent != nil {
		x.parent.SetDirty(x.indexInParent, true, x)
	}
}
func (x *PersonSync) SetParent(sync syncdep.Sync, idx int) {
	x.parent = sync
	x.indexInParent = idx
}
func (x *PersonSync) FlushDirty(dirty bool) {
	if dirty || x.isAgeDirty() {
		x.setAgeDirty(dirty, true)
	}
	if dirty || x.isVipLevelDirty() {
		x.setVipLevelDirty(dirty, true)
	}
	if dirty || x.isNameDirty() {
		x.setNameDirty(dirty, true)
	}
	if dirty || x.isActionsDirty() {
		x.setActionsDirty(dirty, true)
	}
	if dirty || x.isFavorDirty() {
		x.setFavorDirty(dirty, true)
	}
	if dirty || x.isLoveSeqDirty() {
		x.setLoveSeqDirty(dirty, true)
	}
	if dirty || x.isIsGirlDirty() {
		x.setIsGirlDirty(dirty, true)
	}
	if dirty || x.isDetailDirty() {
		x.setDetailDirty(dirty, true)
	}
	if dirty || x.isDataDirty() {
		x.setDataDirty(dirty, true)
	}
}
func (x *PersonSync) setAgeDirty(dirty bool, recur bool) {
	x.SetDirty(x.ageINDEX, dirty, x)
}
func (x *PersonSync) isAgeDirty() bool {
	idx := x.ageINDEX >> 3
	off := x.ageINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setVipLevelDirty(dirty bool, recur bool) {
	x.SetDirty(x.vipLevelINDEX, dirty, x)
}
func (x *PersonSync) isVipLevelDirty() bool {
	idx := x.vipLevelINDEX >> 3
	off := x.vipLevelINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setNameDirty(dirty bool, recur bool) {
	x.SetDirty(x.nameINDEX, dirty, x)
}
func (x *PersonSync) isNameDirty() bool {
	idx := x.nameINDEX >> 3
	off := x.nameINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setActionsDirty(dirty bool, recur bool) {
	x.SetDirty(x.actionsINDEX, dirty, x)
	if recur && x.actions != nil {
		x.actions.FlushDirty(dirty)
	}
}
func (x *PersonSync) isActionsDirty() bool {
	idx := x.actionsINDEX >> 3
	off := x.actionsINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setFavorDirty(dirty bool, recur bool) {
	x.SetDirty(x.favorINDEX, dirty, x)
}
func (x *PersonSync) isFavorDirty() bool {
	idx := x.favorINDEX >> 3
	off := x.favorINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setLoveSeqDirty(dirty bool, recur bool) {
	x.SetDirty(x.loveSeqINDEX, dirty, x)
}
func (x *PersonSync) isLoveSeqDirty() bool {
	idx := x.loveSeqINDEX >> 3
	off := x.loveSeqINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setIsGirlDirty(dirty bool, recur bool) {
	x.SetDirty(x.isGirlINDEX, dirty, x)
}
func (x *PersonSync) isIsGirlDirty() bool {
	idx := x.isGirlINDEX >> 3
	off := x.isGirlINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setDetailDirty(dirty bool, recur bool) {
	x.SetDirty(x.detailINDEX, dirty, x)
	if recur && x.detail != nil {
		x.detail.FlushDirty(dirty)
	}
}
func (x *PersonSync) isDetailDirty() bool {
	idx := x.detailINDEX >> 3
	off := x.detailINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) setDataDirty(dirty bool, recur bool) {
	x.SetDirty(x.dataINDEX, dirty, x)
}
func (x *PersonSync) isDataDirty() bool {
	idx := x.dataINDEX >> 3
	off := x.dataINDEX & 7
	return (x.dirtyFieldMark[idx] & (1 << off)) != 0
}
func (x *PersonSync) Key() interface{} {
	return x.name
}
func (x *PersonSync) SetKey(v interface{}) {
	if x.parent != nil {
		if _, ok := x.parent.(*syncdep.MapSync[string, *PersonSync]); ok {
			panic("PersonSync key can not set")
		}
	}
	x.name = v.(string)
}
func (x *PersonSync) GetAge() int32 {
	return x.age
}
func (x *PersonSync) SetAge(v int32) *PersonSync {
	if x.age == v {
		return x
	}
	x.age = v
	x.setAgeDirty(true, false)
	return x
}
func (x *PersonSync) GetVipLevel() VipLevel {
	return x.vipLevel
}
func (x *PersonSync) SetVipLevel(v VipLevel) *PersonSync {
	if x.vipLevel == v {
		return x
	}
	x.vipLevel = v
	x.setVipLevelDirty(true, false)
	return x
}
func (x *PersonSync) GetName() string {
	return x.name
}
func (x *PersonSync) SetName(v string) *PersonSync {
	if x.name == v {
		return x
	}
	x.name = v
	x.setNameDirty(true, false)
	return x
}
func (x *PersonSync) GetActions() *syncdep.MapSync[string, *ActionInfoSync] {
	if x.actions == nil {
		x.actions = syncdep.NewMapSync[string, *ActionInfoSync]()
		x.actions.SetParent(x, x.actionsINDEX)
	}
	return x.actions
}
func (x *PersonSync) GetFavor() *syncdep.ArraySync[string] {
	if x.favor == nil {
		x.favor = syncdep.NewArraySync[string]()
		x.favor.SetParent(x, x.favorINDEX)
	}
	return x.favor
}
func (x *PersonSync) GetLoveSeq() *syncdep.ArraySync[ColorType] {
	if x.loveSeq == nil {
		x.loveSeq = syncdep.NewArraySync[ColorType]()
		x.loveSeq.SetParent(x, x.loveSeqINDEX)
	}
	return x.loveSeq
}
func (x *PersonSync) GetIsGirl() bool {
	return x.isGirl
}
func (x *PersonSync) SetIsGirl(v bool) *PersonSync {
	if x.isGirl == v {
		return x
	}
	x.isGirl = v
	x.setIsGirlDirty(true, false)
	return x
}
func (x *PersonSync) GetDetail() *IntroDetailSync {
	if x.detail == nil {
		x.detail = NewIntroDetailSync()
		x.detail.SetParent(x, x.detailINDEX)
	}
	return x.detail
}
func (x *PersonSync) SetDetail(v *IntroDetailSync) *PersonSync {
	if v != nil {
		v.SetParent(x, x.detailINDEX)
	}
	if x.detail != nil {
		x.detail.SetParent(nil, -1)
	}
	x.detail = v
	x.setDetailDirty(true, false)
	return x
}
func (x *PersonSync) GetData() []byte {
	return x.data
}
func (x *PersonSync) SetData(v []byte) *PersonSync {
	x.data = v
	x.setDataDirty(true, false)
	return x
}
func (xs *Person) SetAge(v int32) {
	xs.Age = &v
}
func (xs *Person) SetVipLevel(v VipLevel) {
	xs.VipLevel = &v
}
func (xs *Person) SetName(v string) {
	xs.Name = &v
}
func (xs *Person) SetActions(v map[string]*ActionInfo) {
	xs.Actions = v
}
func (xs *Person) SetFavor(v []string) {
	xs.Favor = v
}
func (xs *Person) SetLoveSeq(v []ColorType) {
	xs.LoveSeq = v
}
func (xs *Person) SetIsGirl(v bool) {
	xs.IsGirl = &v
}
func (xs *Person) SetDetail(v *IntroDetail) {
	xs.Detail = v
}
func (xs *Person) SetData(v []byte) {
	xs.Data = v
}
