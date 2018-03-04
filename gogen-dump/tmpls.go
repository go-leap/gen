package main

import (
	"strings"
)

type tmplDotFile struct {
	ProgHint string
	PName    string
	Types    []tmplDotType
	Imps     map[string]string
}

type tmplDotType struct {
	TName    string
	Fields   []tmplDotField
	HasWData bool
}

func (me *tmplDotType) isIfaceSlice(name string) bool {
	if i := strings.Index(name, "["); i > 0 {
		name = name[:i]
		for _, f := range me.Fields {
			if f.FName == name {
				return f.isIfaceSlice
			}
		}
	}
	return false
}

type tmplDotField struct {
	FName string
	TmplW string
	TmplR string

	isIfaceSlice bool
}

const tmplPkg = `package {{.PName}}

// Code generated by {{.ProgHint}} — DO NOT EDIT.

import (
	"bytes"
	"io"
	"unsafe"
	{{range $pkgname, $pkgimppath := .Imps}}
	{{ $pkgname }} "{{$pkgimppath}}"
	{{- end}}
)

{{range .Types}}
func (me *{{.TName}}) writeTo(buf *bytes.Buffer) (err error) {
	{{if .HasWData}}var data bytes.Buffer{{end}}
	{{range .Fields}}
	{{.TmplW}}
	{{end}}
	return
}

func (me *{{.TName}}) WriteTo(w io.Writer) (int64, error) {
	if data, err := me.MarshalBinary(); err == nil {
		var n int
		n, err = w.Write(data)
		return int64(n), err
	} else {
		return 0, err
	}
}

func (me *{{.TName}}) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = me.writeTo(&buf); err == nil {
		data = buf.Bytes()
	}
	return
}

func (me *{{.TName}}) UnmarshalBinary(data []byte) (err error) {
	var pos int
	{{range .Fields}}
	{{.TmplR}}
	{{end}}
	return
}
{{end}}
`
