package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-leap/str"
)

func genDump() {
	for _, tdt := range tdot.Types {
		for _, tdf := range tdt.Fields {
			tdf.TmplR, tdf.TmplW = genForFieldOrVarOfNamedTypeRW(tdf.FName, "", tdt, tdf.typeIdent, 0, 0, tdf.taggedUnion)

			if tdf.isLast { // drop the very-last, thus ineffectual (and hence linter-triggering) assignment to pos
				lastp, lastpalt := ustr.Last(tdf.TmplR, "pos++"), ustr.Last(tdf.TmplR, "pos +=")
				if lastpalt > lastp {
					lastp = lastpalt
				}
				if lastp > 0 {
					off, offalt := ustr.Pos(tdf.TmplR[lastp:], ";"), ustr.Pos(tdf.TmplR[lastp:], "\n")
					if offalt > 0 && offalt < off {
						off = offalt
					}
					if off < 0 {
						off = len(tdf.TmplR) - lastp
					}
					tdf.TmplR = tdf.TmplR[:lastp] + "/*" + tdf.TmplR[lastp:lastp+off] + "*/" + tdf.TmplR[lastp+off:]
				}
			}
		}
	}

	filePathDst := filepath.Join(goPkgDirPath, genFileName)
	file, err := os.Create(filePathDst)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err = genViaTmpl(file); err != nil {
		panic(err)
	}
	os.Stdout.WriteString("generated: " + filePathDst)
}

func genForFieldOrVarOfNamedTypeRW(fieldName string, altNoMe string, tdt *tmplDotType, typeName string, numIndir int, iterDepth int, taggedUnion []string) (tmplR string, tmplW string) {
	nf, mfw, mfr, lf, nfr := fieldName, "me."+fieldName, "me."+fieldName, "l_"+ustr.Replace(fieldName, ".", "ꓸ"), ustr.Replace(fieldName, ".", "ꓸ")
	if altNoMe != "" {
		if mfw = altNoMe; iterDepth > 0 {
			if mfr = ustr.TrimR(altNoMe, ":"); !ustr.Suff(mfr, "["+nf+"]") {
				mfr += "[" + nf + "]"
			}
		}
	}
	if numIndir > 0 {
		mfr = "v_" + nfr + ":"
		mfw = "(" + ustr.Times("*", numIndir) + mfw + ")"
	} else if numIndir == 0 && iterDepth > 0 && altNoMe == "" && (ustr.Pref(nf, "mk_") || ustr.Pref(nf, "mv_")) {
		mfr = "mkv_" + nfr
	}
	var cast string
	if optSafeVarints {
		cast = "uint64"
	}
	switch typeName {
	case "bool":
		tmplW = "if " + mfw + " { buf.WriteByte(1) } else { buf.WriteByte(0) }"
		tmplR = mfr + "= (data[pos] == 1) ; pos++"
	case "uint8", "byte":
		tmplW = "buf.WriteByte(" + mfw + ")"
		tmplR = mfr + "= data[pos] ; pos++"
	case "string":
		tmplW = lf + " := " + cast + "(len(" + mfw + ")) ; " + genLenW(nfr) + " ; buf.WriteString(" + mfw + ")"
		tmplR = genLenR(nfr) + " ; " + mfr + "= string(data[pos : pos+" + lf + "]) ; pos += " + lf
	case "int16", "uint16":
		tmplW = genSizedW(nfr, mfw, "2")
		tmplR = genSizedR(mfr, typeName, "2")
	case "rune", "int32", "float32", "uint32":
		tmplW = genSizedW(nfr, mfw, "4")
		tmplR = genSizedR(mfr, typeName, "4")
	case "complex64", "float64", "uint64", "int64":
		tmplW = genSizedW(nfr, mfw, "8")
		tmplR = genSizedR(mfr, typeName, "8")
	case "complex128":
		tmplW = genSizedW(nfr, mfw, "16")
		tmplR = genSizedR(mfr, typeName, "16")
	case "uint", "uintptr":
		tmplW = genSizedW(nfr, mfw, "uint64")
		tmplR = genSizedR(mfr, typeName, "uint64")
	case "int":
		tmplW = genSizedW(nfr, mfw, "int64")
		tmplR = genSizedR(mfr, typeName, "int64")
	default:
		if typeName[0] == '*' {
			// POINTER

			var numindir int
			for _, r := range typeName {
				if r == '*' {
					numindir++
				} else {
					break
				}
			}
			tr, tw := genForFieldOrVarOfNamedTypeRW(nf, altNoMe, tdt, typeName[numindir:], numindir, iterDepth, taggedUnion)
			for i := 0; i < numindir; i++ {
				if ustr.Pref(mfr, "v_") || ustr.Pref(mfr, "mkv_") || ustr.Has(mfr, "[") {
					tmplR += "if pos++; data[pos-1] != 0 { "
				} else {
					tmplR += "if pos++; data[pos-1] == 0 { " + mfr + " = nil } else { "
				}
				tmplW += "if " + ustr.Times("*", i) + mfw + " == nil { buf.WriteByte(0) } else { buf.WriteByte(1) ; "
			}
			tmplR += "\n\t\t" + tr + " ; "
			for i := 0; i < numindir; i++ {
				if i == 0 {
					if i == numindir-1 {
						tmplR += mfr + "= &v_" + nfr
					} else {
						tmplR += "p0_" + nfr + " := &v_" + nfr + " ; "
					}
				} else if i == numindir-1 {
					tmplR += mfr + " = &p" + strconv.Itoa(i-1) + "_" + nfr
				} else {
					tmplR += "p" + strconv.Itoa(i) + "_" + nfr + " := &p" + strconv.Itoa(i-1) + "_" + nfr + " ; "
				}
			}
			tmplR += "\n\t" + ustr.Times("}", numindir)
			tmplW += "\n\t\t" + tw
			tmplW += "\n\t" + ustr.Times("}", numindir)
		} else if ismap, pclose := ustr.Pref(typeName, "map["), ustr.Pos(typeName[1:], "]")+1; pclose > 0 && (typeName[0] == '[' || ismap) {
			// ARRAY / SLICE / MAP

			slen := typeName[1:pclose]
			if slen == "" || ismap {
				if slen = "(" + lf + ")"; optSafeVarints {
					slen = "int" + slen
				}
				tmplR += genLenR(nfr) + " ; " + mfr + "= make(" + typeName + ", " + lf + ") ; "
				tmplW += lf + " := " + cast + "(len(" + mfw + ")) ; " + genLenW(nfr) + " ; "
			} else if numIndir > 0 {
				tmplR += mfr + "= " + typeName + "{} ; "
			}
			idx := ustr.Times("i", iterDepth+1) + "_" + nfr

			if ismap {
				mk, mv := ustr.Times("mk_", iterDepth+1)+"_"+nfr, ustr.Times("mv_", iterDepth+1)+nfr
				tmplR += "for " + idx + " := 0; " + idx + " < " + slen + "; " + idx + "++ {"
				tmplW += "for " + mk + ", " + mv + " := range " + mfr + " {"

				tmplR += "\n\t\tvar mkv_" + mk + " " + typeName[4:pclose]
				tmplR += "\n\t\tvar mkv_" + mv + " " + typeName[pclose+1:]
				tr, _ := genForFieldOrVarOfNamedTypeRW(mk, "", tdt, typeName[4:pclose], 0, iterDepth+1, taggedUnion)
				tmplR += "\n\t\t" + tr
				tr, _ = genForFieldOrVarOfNamedTypeRW(mv, "", tdt, typeName[pclose+1:], 0, iterDepth+1, taggedUnion)
				tmplR += "\n\t\t" + tr
				tmplR += "\n\t\t" + mfr + "[mkv_" + mk + "] = mkv_" + mv
				_, tw := genForFieldOrVarOfNamedTypeRW(mk, mk, tdt, typeName[4:pclose], 0, iterDepth+1, taggedUnion)
				tmplW += "\n\t\t" + tw
				_, tw = genForFieldOrVarOfNamedTypeRW(mv, mv, tdt, typeName[pclose+1:], 0, iterDepth+1, taggedUnion)
				tmplW += "\n\t\t" + tw

				tmplR += "\n\t}"
				tmplW += "\n\t}"
			} else {
				tmplR += "for " + idx + " := 0; " + idx + " < " + slen + "; " + idx + "++ {"
				tmplW += "for " + idx + " := 0; " + idx + " < " + slen + "; " + idx + "++ {"
				if ustr.Pref(mfr, "me.") && tdt.isIfaceSlice(mfr[3:]) {
					mfr = mfr + ".(" + typeName + ")"
				}
				tr, _ := genForFieldOrVarOfNamedTypeRW(idx, mfr, tdt, typeName[pclose+1:], 0, iterDepth+1, taggedUnion)
				tmplR += "\n\t\t" + tr
				_, tw := genForFieldOrVarOfNamedTypeRW(idx, mfw+"["+idx+"]", tdt, typeName[pclose+1:], 0, iterDepth+1, taggedUnion)
				tmplW += "\n\t\t" + tw
				tmplR += "\n\t}"
				tmplW += "\n\t}"
			}
		} else if len(taggedUnion) > 0 {
			// TAGGED INTERFACE

			tmplR += "t_" + nfr + " := data[pos] ; pos++ ; switch t_" + nfr + " {"
			tmplW += "switch t_" + nfr + " := " + mfw + ".(type) {"
			for ti, tu := range taggedUnion {
				tr, _ := genForFieldOrVarOfNamedTypeRW(nf, mfr, tdt, tu, 0, iterDepth, nil)
				tmplR += "\n\t\tcase " + strconv.Itoa(ti+1) + ":\n\t\t\t" + tr
				_, tw := genForFieldOrVarOfNamedTypeRW(nf, "t_"+nfr, tdt, tu, 0, iterDepth, nil)
				tmplW += "\n\t\tcase " + tu + ":\n\t\t\tbuf.WriteByte(" + strconv.Itoa(ti+1) + ") ; " + tw
			}
			tmplR += "\n\t\tdefault:\n\t\t\t" + mfr + " = nil"
			tmplW += "\n\t\tdefault:\n\t\t\tbuf.WriteByte(0)"
			tmplR += "\n\t}"
			tmplW += "\n\t}"
		} else {
			// OTHER

			if ustr.Pref(mfr, "v_") {
				tmplR = mfr + "= " + typeName + "{} ; "
			}
			tmplR += genLenR(nfr) + " ; if err = " + ustr.TrimR(mfr, ":") + ".UnmarshalBinary(data[pos : pos+" + lf + "]); err != nil { return } ; pos += " + lf
			tmplW += "d_" + nfr + ", e_" + nfr + " := " + mfw + ".MarshalBinary() ; if err = e_" + nfr + "; err != nil { return } ; " + lf + " := " + cast + "(len(d_" + nfr + ")) ; " + genLenW(nfr) + " ; buf.Write(d_" + nfr + ")"

			for tspec := range ts {
				if tspec.Name.Name == typeName {
					tdt.HasWData = true
					if ustr.Pref(mfw, "(") && ustr.Suff(mfw, ")") && mfw[1] == '*' {
						mfw = ustr.TrimL(mfw[1:len(mfw)-1], "*")
					}
					tmplW = "if err = " + mfw + ".writeTo(&data); err != nil { return } ; " + lf + " := " + cast + "(data.Len()) ; " + genLenW(nfr) + " ; data.WriteTo(buf)"
					break
				}
			}
		}
	}
	return
}

func genSizedR(mfr string, typeName string, byteSize string) string {
	if byteSize == "int64" || byteSize == "uint64" {
		if optSafeVarints {
			return mfr + "= " + typeName + "(*((*" + byteSize + ")(unsafe.Pointer(&data[pos])))) ; pos += 8"
		}
		byteSize = "8"
	}
	return mfr + "= *((*" + typeName + ")(unsafe.Pointer(&data[pos]))) ; pos += " + byteSize
}

func genSizedW(fieldName string, mfw string, byteSize string) (s string) {
	if byteSize == "int64" || byteSize == "uint64" {
		if optSafeVarints {
			return byteSize + "_" + fieldName + " := " + byteSize + "(" + mfw + ") ; buf.Write(((*[8]byte)(unsafe.Pointer(&" + byteSize + "_" + fieldName + ")))[:])"
		}
		byteSize = "8"
	}
	if ustr.Pref(mfw, "(*") && ustr.Suff(mfw, ")") {
		return "buf.Write(((*[" + byteSize + "]byte)(unsafe.Pointer(" + mfw[2:len(mfw)-1] + ")))[:])"
	} else if mfw[0] == '*' {
		return "buf.Write(((*[" + byteSize + "]byte)(unsafe.Pointer(" + mfw[1:] + ")))[:])"
	}
	return "buf.Write(((*[" + byteSize + "]byte)(unsafe.Pointer(&(" + mfw + "))))[:])"
}

func genLenR(fieldName string) string {
	if optSafeVarints {
		return "l_" + fieldName + " := int(*((*uint64)(unsafe.Pointer(&data[pos])))) ; pos += 8"
	}
	return "l_" + fieldName + " := (*((*int)(unsafe.Pointer(&data[pos])))) ; pos += 8"
}

func genLenW(fieldName string) string {
	return "buf.Write((*[8]byte)(unsafe.Pointer(&l_" + fieldName + "))[:])"
}
