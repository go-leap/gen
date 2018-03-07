package main

// Code generated by github.com/metaleap/gogen-dump — DO NOT EDIT.

import (
	"bytes"
	"io"
	"unsafe"
)

func (me *fixed) writeTo(buf *bytes.Buffer) (err error) {

	buf.Write((*[2036]byte)(unsafe.Pointer(me))[:])

	return
}

func (me *fixed) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	if err := me.writeTo(&buf); err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

func (me *fixed) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = me.writeTo(&buf); err == nil {
		data = buf.Bytes()
	}
	return
}

func (me *fixed) UnmarshalBinary(data []byte) (err error) {

	*me = *((*fixed)(unsafe.Pointer(&data[0])))

	return
}

func (me *testStruct) writeTo(buf *bytes.Buffer) (err error) {

	var data bytes.Buffer

	if err = me.embName.writeTo(&data); err != nil {
		return
	}
	l_embName := (data.Len())
	buf.Write((*[8]byte)(unsafe.Pointer(&l_embName))[:])
	data.WriteTo(buf)

	buf.Write(((*[16]byte)(unsafe.Pointer(&(me.DingDong.Complex))))[:])

	buf.Write(((*[504]byte)(unsafe.Pointer(&(me.DingDong.FixedSize[0]))))[:])

	if me.Hm.Balance == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		for i_HmꓸBalance := 0; i_HmꓸBalance < 3; i_HmꓸBalance++ {
			if (*me.Hm.Balance)[i_HmꓸBalance] == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				if *(*me.Hm.Balance)[i_HmꓸBalance] == nil {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(1)
					buf.Write(((*[2]byte)(unsafe.Pointer(*(*me.Hm.Balance)[i_HmꓸBalance])))[:])
				}
			}
		}
	}

	buf.Write(((*[8]byte)(unsafe.Pointer(&(me.Hm.Hm.AccountAge))))[:])

	l_HmꓸHmꓸLookie := (len(me.Hm.Hm.Lookie))
	buf.Write((*[8]byte)(unsafe.Pointer(&l_HmꓸHmꓸLookie))[:])
	for i_HmꓸHmꓸLookie := 0; i_HmꓸHmꓸLookie < (l_HmꓸHmꓸLookie); i_HmꓸHmꓸLookie++ {
		if me.Hm.Hm.Lookie[i_HmꓸHmꓸLookie] == nil {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
			buf.Write((*[2036]byte)(unsafe.Pointer(me.Hm.Hm.Lookie[i_HmꓸHmꓸLookie]))[:])
		}
	}

	l_HmꓸHmꓸAny := (len(me.Hm.Hm.Any))
	buf.Write((*[8]byte)(unsafe.Pointer(&l_HmꓸHmꓸAny))[:])
	for i_HmꓸHmꓸAny := 0; i_HmꓸHmꓸAny < (l_HmꓸHmꓸAny); i_HmꓸHmꓸAny++ {
		switch t_i_HmꓸHmꓸAny := me.Hm.Hm.Any[i_HmꓸHmꓸAny].(type) {
		case *embName:
			buf.WriteByte(1)
			if t_i_HmꓸHmꓸAny == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				if err = t_i_HmꓸHmꓸAny.writeTo(&data); err != nil {
					return
				}
				l_i_HmꓸHmꓸAny := (data.Len())
				buf.Write((*[8]byte)(unsafe.Pointer(&l_i_HmꓸHmꓸAny))[:])
				data.WriteTo(buf)
			}
		case []embName:
			buf.WriteByte(2)
			l_i_HmꓸHmꓸAny := (len(t_i_HmꓸHmꓸAny))
			buf.Write((*[8]byte)(unsafe.Pointer(&l_i_HmꓸHmꓸAny))[:])
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				if err = t_i_HmꓸHmꓸAny[ii_i_HmꓸHmꓸAny].writeTo(&data); err != nil {
					return
				}
				l_ii_i_HmꓸHmꓸAny := (data.Len())
				buf.Write((*[8]byte)(unsafe.Pointer(&l_ii_i_HmꓸHmꓸAny))[:])
				data.WriteTo(buf)
			}
		case []*embName:
			buf.WriteByte(3)
			l_i_HmꓸHmꓸAny := (len(t_i_HmꓸHmꓸAny))
			buf.Write((*[8]byte)(unsafe.Pointer(&l_i_HmꓸHmꓸAny))[:])
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				if t_i_HmꓸHmꓸAny[ii_i_HmꓸHmꓸAny] == nil {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(1)
					if err = t_i_HmꓸHmꓸAny[ii_i_HmꓸHmꓸAny].writeTo(&data); err != nil {
						return
					}
					l_ii_i_HmꓸHmꓸAny := (data.Len())
					buf.Write((*[8]byte)(unsafe.Pointer(&l_ii_i_HmꓸHmꓸAny))[:])
					data.WriteTo(buf)
				}
			}
		case []*float32:
			buf.WriteByte(4)
			l_i_HmꓸHmꓸAny := (len(t_i_HmꓸHmꓸAny))
			buf.Write((*[8]byte)(unsafe.Pointer(&l_i_HmꓸHmꓸAny))[:])
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				if t_i_HmꓸHmꓸAny[ii_i_HmꓸHmꓸAny] == nil {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(1)
					buf.Write(((*[4]byte)(unsafe.Pointer(t_i_HmꓸHmꓸAny[ii_i_HmꓸHmꓸAny])))[:])
				}
			}
		default:
			buf.WriteByte(0)
		}
	}

	l_HmꓸFoo := (len(me.Hm.Foo))
	buf.Write((*[8]byte)(unsafe.Pointer(&l_HmꓸFoo))[:])
	for i_HmꓸFoo := 0; i_HmꓸFoo < (l_HmꓸFoo); i_HmꓸFoo++ {
		for ii_i_HmꓸFoo := 0; ii_i_HmꓸFoo < 2; ii_i_HmꓸFoo++ {
			l_ii_i_HmꓸFoo := (len(me.Hm.Foo[i_HmꓸFoo][ii_i_HmꓸFoo]))
			buf.Write((*[8]byte)(unsafe.Pointer(&l_ii_i_HmꓸFoo))[:])
			for mk_mk_mk__ii_i_HmꓸFoo, mv_mv_mv_ii_i_HmꓸFoo := range me.Hm.Foo[i_HmꓸFoo][ii_i_HmꓸFoo] {
				buf.Write(((*[4]byte)(unsafe.Pointer(&(mk_mk_mk__ii_i_HmꓸFoo))))[:])
				if mv_mv_mv_ii_i_HmꓸFoo == nil {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(1)
					if *mv_mv_mv_ii_i_HmꓸFoo == nil {
						buf.WriteByte(0)
					} else {
						buf.WriteByte(1)
						if **mv_mv_mv_ii_i_HmꓸFoo == nil {
							buf.WriteByte(0)
						} else {
							buf.WriteByte(1)
							l_mv_mv_mv_ii_i_HmꓸFoo := (len((***mv_mv_mv_ii_i_HmꓸFoo)))
							buf.Write((*[8]byte)(unsafe.Pointer(&l_mv_mv_mv_ii_i_HmꓸFoo))[:])
							for iiii_mv_mv_mv_ii_i_HmꓸFoo := 0; iiii_mv_mv_mv_ii_i_HmꓸFoo < (l_mv_mv_mv_ii_i_HmꓸFoo); iiii_mv_mv_mv_ii_i_HmꓸFoo++ {
								if (***mv_mv_mv_ii_i_HmꓸFoo)[iiii_mv_mv_mv_ii_i_HmꓸFoo] == nil {
									buf.WriteByte(0)
								} else {
									buf.WriteByte(1)
									buf.Write(((*[2]byte)(unsafe.Pointer((***mv_mv_mv_ii_i_HmꓸFoo)[iiii_mv_mv_mv_ii_i_HmꓸFoo])))[:])
								}
							}
						}
					}
				}
			}
		}
	}

	if me.Age == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		if *me.Age == nil {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
			if **me.Age == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				buf.Write(((*[8]byte)(unsafe.Pointer(**me.Age)))[:])
			}
		}
	}

	return
}

func (me *testStruct) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	if err := me.writeTo(&buf); err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

func (me *testStruct) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = me.writeTo(&buf); err == nil {
		data = buf.Bytes()
	}
	return
}

func (me *testStruct) UnmarshalBinary(data []byte) (err error) {

	var pos int

	l_embName := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	if err = me.embName.UnmarshalBinary(data[pos : pos+l_embName]); err != nil {
		return
	}
	pos += l_embName

	me.DingDong.Complex = *((*complex128)(unsafe.Pointer(&data[pos])))
	pos += 16

	me.DingDong.FixedSize = *((*[9][7]float64)(unsafe.Pointer(&data[pos])))
	pos += 504

	if pos++; data[pos-1] == 0 {
		me.Hm.Balance = nil
	} else {
		v_HmꓸBalance := [3]**int16{}
		for i_HmꓸBalance := 0; i_HmꓸBalance < 3; i_HmꓸBalance++ {
			if pos++; data[pos-1] != 0 {
				if pos++; data[pos-1] != 0 {
					v_i_HmꓸBalance := *((*int16)(unsafe.Pointer(&data[pos])))
					pos += 2
					p0_i_HmꓸBalance := &v_i_HmꓸBalance
					v_HmꓸBalance[i_HmꓸBalance] = &p0_i_HmꓸBalance
				}
			}
		}
		me.Hm.Balance = &v_HmꓸBalance
	}

	me.Hm.Hm.AccountAge = *((*int)(unsafe.Pointer(&data[pos])))
	pos += 8

	l_HmꓸHmꓸLookie := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	me.Hm.Hm.Lookie = make([]*fixed, l_HmꓸHmꓸLookie)
	for i_HmꓸHmꓸLookie := 0; i_HmꓸHmꓸLookie < (l_HmꓸHmꓸLookie); i_HmꓸHmꓸLookie++ {
		if pos++; data[pos-1] != 0 {
			v_i_HmꓸHmꓸLookie := *((*fixed)(unsafe.Pointer(&data[pos])))
			pos += 2036
			me.Hm.Hm.Lookie[i_HmꓸHmꓸLookie] = &v_i_HmꓸHmꓸLookie
		}
	}

	l_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	me.Hm.Hm.Any = make([]interface{}, l_HmꓸHmꓸAny)
	for i_HmꓸHmꓸAny := 0; i_HmꓸHmꓸAny < (l_HmꓸHmꓸAny); i_HmꓸHmꓸAny++ {
		t_i_HmꓸHmꓸAny := data[pos]
		pos++
		switch t_i_HmꓸHmꓸAny {
		case 1:
			if pos++; data[pos-1] != 0 {
				v_i_HmꓸHmꓸAny := embName{}
				l_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
				pos += 8
				if err = v_i_HmꓸHmꓸAny.UnmarshalBinary(data[pos : pos+l_i_HmꓸHmꓸAny]); err != nil {
					return
				}
				pos += l_i_HmꓸHmꓸAny
				me.Hm.Hm.Any[i_HmꓸHmꓸAny] = &v_i_HmꓸHmꓸAny
			}
		case 2:
			l_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
			pos += 8
			me.Hm.Hm.Any[i_HmꓸHmꓸAny] = make([]embName, l_i_HmꓸHmꓸAny)
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				l_ii_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
				pos += 8
				if err = me.Hm.Hm.Any[i_HmꓸHmꓸAny].([]embName)[ii_i_HmꓸHmꓸAny].UnmarshalBinary(data[pos : pos+l_ii_i_HmꓸHmꓸAny]); err != nil {
					return
				}
				pos += l_ii_i_HmꓸHmꓸAny
			}
		case 3:
			l_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
			pos += 8
			me.Hm.Hm.Any[i_HmꓸHmꓸAny] = make([]*embName, l_i_HmꓸHmꓸAny)
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				if pos++; data[pos-1] != 0 {
					v_ii_i_HmꓸHmꓸAny := embName{}
					l_ii_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
					pos += 8
					if err = v_ii_i_HmꓸHmꓸAny.UnmarshalBinary(data[pos : pos+l_ii_i_HmꓸHmꓸAny]); err != nil {
						return
					}
					pos += l_ii_i_HmꓸHmꓸAny
					me.Hm.Hm.Any[i_HmꓸHmꓸAny].([]*embName)[ii_i_HmꓸHmꓸAny] = &v_ii_i_HmꓸHmꓸAny
				}
			}
		case 4:
			l_i_HmꓸHmꓸAny := (*((*int)(unsafe.Pointer(&data[pos]))))
			pos += 8
			me.Hm.Hm.Any[i_HmꓸHmꓸAny] = make([]*float32, l_i_HmꓸHmꓸAny)
			for ii_i_HmꓸHmꓸAny := 0; ii_i_HmꓸHmꓸAny < (l_i_HmꓸHmꓸAny); ii_i_HmꓸHmꓸAny++ {
				if pos++; data[pos-1] != 0 {
					v_ii_i_HmꓸHmꓸAny := *((*float32)(unsafe.Pointer(&data[pos])))
					pos += 4
					me.Hm.Hm.Any[i_HmꓸHmꓸAny].([]*float32)[ii_i_HmꓸHmꓸAny] = &v_ii_i_HmꓸHmꓸAny
				}
			}
		default:
			me.Hm.Hm.Any[i_HmꓸHmꓸAny] = nil
		}
	}

	l_HmꓸFoo := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	me.Hm.Foo = make([][2]map[rune]***[]*int16, l_HmꓸFoo)
	for i_HmꓸFoo := 0; i_HmꓸFoo < (l_HmꓸFoo); i_HmꓸFoo++ {
		for ii_i_HmꓸFoo := 0; ii_i_HmꓸFoo < 2; ii_i_HmꓸFoo++ {
			l_ii_i_HmꓸFoo := (*((*int)(unsafe.Pointer(&data[pos]))))
			pos += 8
			me.Hm.Foo[i_HmꓸFoo][ii_i_HmꓸFoo] = make(map[rune]***[]*int16, l_ii_i_HmꓸFoo)
			for iii_ii_i_HmꓸFoo := 0; iii_ii_i_HmꓸFoo < (l_ii_i_HmꓸFoo); iii_ii_i_HmꓸFoo++ {
				var mkv_mk_mk_mk__ii_i_HmꓸFoo rune
				var mkv_mv_mv_mv_ii_i_HmꓸFoo ***[]*int16
				mkv_mk_mk_mk__ii_i_HmꓸFoo = *((*rune)(unsafe.Pointer(&data[pos])))
				pos += 4
				if pos++; data[pos-1] != 0 {
					if pos++; data[pos-1] != 0 {
						if pos++; data[pos-1] != 0 {
							l_mv_mv_mv_ii_i_HmꓸFoo := (*((*int)(unsafe.Pointer(&data[pos]))))
							pos += 8
							v_mv_mv_mv_ii_i_HmꓸFoo := make([]*int16, l_mv_mv_mv_ii_i_HmꓸFoo)
							for iiii_mv_mv_mv_ii_i_HmꓸFoo := 0; iiii_mv_mv_mv_ii_i_HmꓸFoo < (l_mv_mv_mv_ii_i_HmꓸFoo); iiii_mv_mv_mv_ii_i_HmꓸFoo++ {
								if pos++; data[pos-1] != 0 {
									v_iiii_mv_mv_mv_ii_i_HmꓸFoo := *((*int16)(unsafe.Pointer(&data[pos])))
									pos += 2
									v_mv_mv_mv_ii_i_HmꓸFoo[iiii_mv_mv_mv_ii_i_HmꓸFoo] = &v_iiii_mv_mv_mv_ii_i_HmꓸFoo
								}
							}
							p0_mv_mv_mv_ii_i_HmꓸFoo := &v_mv_mv_mv_ii_i_HmꓸFoo
							p1_mv_mv_mv_ii_i_HmꓸFoo := &p0_mv_mv_mv_ii_i_HmꓸFoo
							mkv_mv_mv_mv_ii_i_HmꓸFoo = &p1_mv_mv_mv_ii_i_HmꓸFoo
						}
					}
				}
				me.Hm.Foo[i_HmꓸFoo][ii_i_HmꓸFoo][mkv_mk_mk_mk__ii_i_HmꓸFoo] = mkv_mv_mv_mv_ii_i_HmꓸFoo
			}
		}
	}

	if pos++; data[pos-1] == 0 {
		me.Age = nil
	} else {
		if pos++; data[pos-1] == 0 {
			me.Age = nil
		} else {
			if pos++; data[pos-1] == 0 {
				me.Age = nil
			} else {
				v_Age := *((*uint)(unsafe.Pointer(&data[pos]))) /*pos += 8 */
				p0_Age := &v_Age
				p1_Age := &p0_Age
				me.Age = &p1_Age
			}
		}
	}

	return
}

func (me *embName) writeTo(buf *bytes.Buffer) (err error) {

	l_FirstName := (len(me.FirstName))
	buf.Write((*[8]byte)(unsafe.Pointer(&l_FirstName))[:])
	buf.WriteString(me.FirstName)

	l_MiddleNames := (len(me.MiddleNames))
	buf.Write((*[8]byte)(unsafe.Pointer(&l_MiddleNames))[:])
	for i_MiddleNames := 0; i_MiddleNames < (l_MiddleNames); i_MiddleNames++ {
		if me.MiddleNames[i_MiddleNames] == nil {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
			if *me.MiddleNames[i_MiddleNames] == nil {
				buf.WriteByte(0)
			} else {
				buf.WriteByte(1)
				if **me.MiddleNames[i_MiddleNames] == nil {
					buf.WriteByte(0)
				} else {
					buf.WriteByte(1)
					for ii_i_MiddleNames := 0; ii_i_MiddleNames < 5; ii_i_MiddleNames++ {
						if (***me.MiddleNames[i_MiddleNames])[ii_i_MiddleNames] == nil {
							buf.WriteByte(0)
						} else {
							buf.WriteByte(1)
							l_ii_i_MiddleNames := (len((*(***me.MiddleNames[i_MiddleNames])[ii_i_MiddleNames])))
							buf.Write((*[8]byte)(unsafe.Pointer(&l_ii_i_MiddleNames))[:])
							buf.WriteString((*(***me.MiddleNames[i_MiddleNames])[ii_i_MiddleNames]))
						}
					}
				}
			}
		}
	}

	if me.LastName == nil {
		buf.WriteByte(0)
	} else {
		buf.WriteByte(1)
		if *me.LastName == nil {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
			l_LastName := (len((**me.LastName)))
			buf.Write((*[8]byte)(unsafe.Pointer(&l_LastName))[:])
			buf.WriteString((**me.LastName))
		}
	}

	return
}

func (me *embName) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	if err := me.writeTo(&buf); err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

func (me *embName) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = me.writeTo(&buf); err == nil {
		data = buf.Bytes()
	}
	return
}

func (me *embName) UnmarshalBinary(data []byte) (err error) {

	var pos int

	l_FirstName := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	me.FirstName = string(data[pos : pos+l_FirstName])
	pos += l_FirstName

	l_MiddleNames := (*((*int)(unsafe.Pointer(&data[pos]))))
	pos += 8
	me.MiddleNames = make([]***[5]*string, l_MiddleNames)
	for i_MiddleNames := 0; i_MiddleNames < (l_MiddleNames); i_MiddleNames++ {
		if pos++; data[pos-1] != 0 {
			if pos++; data[pos-1] != 0 {
				if pos++; data[pos-1] != 0 {
					v_i_MiddleNames := [5]*string{}
					for ii_i_MiddleNames := 0; ii_i_MiddleNames < 5; ii_i_MiddleNames++ {
						if pos++; data[pos-1] != 0 {
							l_ii_i_MiddleNames := (*((*int)(unsafe.Pointer(&data[pos]))))
							pos += 8
							v_ii_i_MiddleNames := string(data[pos : pos+l_ii_i_MiddleNames])
							pos += l_ii_i_MiddleNames
							v_i_MiddleNames[ii_i_MiddleNames] = &v_ii_i_MiddleNames
						}
					}
					p0_i_MiddleNames := &v_i_MiddleNames
					p1_i_MiddleNames := &p0_i_MiddleNames
					me.MiddleNames[i_MiddleNames] = &p1_i_MiddleNames
				}
			}
		}
	}

	if pos++; data[pos-1] == 0 {
		me.LastName = nil
	} else {
		if pos++; data[pos-1] == 0 {
			me.LastName = nil
		} else {
			l_LastName := (*((*int)(unsafe.Pointer(&data[pos]))))
			pos += 8
			v_LastName := string(data[pos : pos+l_LastName]) /*pos += l_LastName */
			p0_LastName := &v_LastName
			me.LastName = &p0_LastName
		}
	}

	return
}
