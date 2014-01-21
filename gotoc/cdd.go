package gotoc

import (
	"bytes"
	"code.google.com/p/go.tools/go/types"
	"fmt"
	"io"
)

type DeclType int

const (
	FuncDecl DeclType = iota
	VarDecl
	ConstDecl
	TypeDecl
	ImportDecl
)

// CDD stores Go declaration translated to C declaration and definition.
type CDD struct {
	Origin types.Object              // object for this declaration
	Uses   map[types.Object]struct{} // used declarations

	Type   DeclType
	Export bool
	Inline bool

	Decl []byte
	Def  []byte
}

func newCDD(o types.Object, t DeclType) *CDD {
	cdd := &CDD{
		Origin: o,
		Type:   t,
		Uses:   make(map[types.Object]struct{}),
	}
	if t == FuncDecl && o.Name() == "main" && o.Pkg().Name() == "main" {
		cdd.Export = true
	} else {
		cdd.Export = o.IsExported()
	}
	return cdd
}

func (cdd *CDD) copyDecl(b *bytes.Buffer, suffix string) {
	n := b.Len()
	b.WriteString(suffix)
	cdd.Decl = append([]byte(nil), b.Bytes()...)
	b.Truncate(n)
}

func (cdd *CDD) copyDef(b *bytes.Buffer) {
	cdd.Def = append([]byte(nil), b.Bytes()...)
}

func (cdd *CDD) WriteDecl(wh, wc io.Writer) error {
	prefix := ""

	switch cdd.Type {
	case FuncDecl:
		if cdd.Inline {
			prefix = "static inline "
		} else if !cdd.Export {
			prefix = "static "
		}

	case VarDecl:
		if cdd.Export {
			prefix = "extern "
		} else {
			prefix = "static "
		}

	case ConstDecl:
		if !cdd.Export {
			return nil
		}

	case ImportDecl:
		if !cdd.Export {
			return nil
		}
	}

	w := wc
	if cdd.Export {
		w = wh
	}

	_, err := io.WriteString(w, prefix)
	if err != nil {
		return err
	}
	if len(cdd.Decl) == 0 {
		fmt.Fprintf(w, "<%s>", cdd.Origin.Type())
	} else {
		_, err = w.Write(cdd.Decl)
	}
	return err
}

func (cdd *CDD) WriteDef(wh, wc io.Writer) error {
	prefix := ""

	switch cdd.Type {
	case FuncDecl:
		if cdd.Inline {
			prefix = "static inline "
		} else if !cdd.Export {
			prefix = "static "
		}

	case VarDecl:
		if !cdd.Export {
			prefix = "static "
		}

	case ConstDecl:
		return nil

	case ImportDecl:
		return nil
	}

	w := wc
	if cdd.Inline {
		w = wh
	}

	_, err := io.WriteString(w, prefix)
	if err != nil {
		return err
	}
	if len(cdd.Decl) == 0 {
		fmt.Fprintf(w, "<%s>", cdd.Origin.Type())
	} else {
		_, err = w.Write(cdd.Def)
	}
	return err
}