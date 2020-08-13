package gen

import (
	"fmt"
	"gitlab.com/osaki-lab/errcdgen"
	"go/ast"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	PkgName string
	Decls   []*Decl
}

type Decl struct {
	Name             string
	Code             string
	Format           string
	FormatVerbs      []string // set by fmt analyzer
	LogLevelEnable   bool
	LogLevel         errcdgen.Level
	StatusCodeEnable bool
	StatusCode       int
	DisableErr       bool
	Labels           []Label
	chainFuncName    string // inside fields
}

func (d Decl) LabelEnable() bool {
	return len(d.Labels) > 0
}

func (d Decl) Args() string {
	var resp = ""
	for _, v := range d.Labels {
		if resp != "" {
			resp += ","
		}
		resp += v.Name + " " + v.GoType
	}
	return resp
}

func (d Decl) ArgValues() string {
	var resp = ""
	for _, v := range d.Labels {
		if resp != "" {
			resp += ","
		}
		resp += v.Name
	}
	return resp
}

type Label struct {
	Index  int
	Name   string
	GoType string
}

func Traverse(n *ast.File) (*File, error) {

	var resp []*Decl
	for _, d := range n.Decls {

		//fs := token.NewFileSet()
		//ast.Print(fs, d)

		decls, err := traverseAst(d)
		if err != nil {
			return nil, err
		}
		if decls != nil {
			resp = append(resp, decls...)
		}
	}

	return &File{
		PkgName: n.Name.Name,
		Decls:   resp,
	}, nil
}

func traverseAst(n ast.Node) ([]*Decl, error) {
	var resp []*Decl

	switch n := n.(type) {
	case *ast.DeclStmt:
		decls, err := traverseAst(n.Decl)
		if err != nil {
			return nil, err
		}

		if decls != nil {
			resp = append(resp, decls...)
		}
	case *ast.GenDecl:
		// const is not target block
		if n.Tok.String() == "var" {
			for _, spec := range n.Specs {
				decls, err := traverseAst(spec)
				if err != nil {
					return nil, err
				}
				if decls != nil {
					resp = append(resp, decls...)
				}
			}
		}
	case *ast.ValueSpec:
		// カンマ区切りで左辺に複数宣言するのには対応しない
		declare := traverseSingle(n.Values[0])
		if declare != nil {
			declare.Name = n.Names[0].Name
			resp = append(resp, declare)
		}
	}

	return resp, nil
}

func traverseSingle(v ast.Expr) *Decl {
	switch v := v.(type) {
	case *ast.CallExpr:
		decl := traverseSingle(v.Fun)
		if decl == nil {
			return nil
		}

		switch decl.chainFuncName {
		case "NewCodeError":
			if len(v.Args) != 2 {
				fmt.Printf("invalid NewCodeErrorFun Args: %v\n", v.Args)
				return nil
			}

			arg0, ok := v.Args[0].(*ast.BasicLit)
			if !ok {
				return nil
			}
			arg1, ok := v.Args[1].(*ast.BasicLit)
			if !ok {
				return nil
			}

			format := strings.Trim(arg1.Value, `"`)
			verbs := Analyze(format)

			return &Decl{
				Code:        strings.Trim(arg0.Value, `"`),
				Format:      format,
				FormatVerbs: verbs.Verb,
			}

		case "Label":
			if len(v.Args) != 3 {
				fmt.Fprintf(os.Stderr, "Label is only allowed 3 args: %v", v.Args)
			}

			index, err := strconv.Atoi(v.Args[0].(*ast.BasicLit).Value)
			if err != nil {
				log.Println("label parse: ", v, err)
				return nil
			}
			decl.Labels = append(decl.Labels, Label{
				Index:  index,
				Name:   strings.Trim(v.Args[1].(*ast.BasicLit).Value, `"`),
				GoType: DetectGoType(v.Args[2]),
			})

		case "DisableError":
			decl.DisableErr = true
		case "WithStatusCode":
			decl.StatusCodeEnable = true
			arg0, ok := v.Args[0].(*ast.BasicLit)
			if !ok {
				return decl
			}
			decl.StatusCode, _ = strconv.Atoi(arg0.Value)
		case "TraceLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.TraceLevel
		case "DebugLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.DebugLevel
		case "InfoLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.InfoLevel
		case "WarnLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.WarnLevel
		case "ErrorLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.ErrorLevel
		case "FatalLevel":
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.FatalLevel
		}
		return decl

	case *ast.SelectorExpr:
		if x, ok := v.X.(*ast.Ident); ok && x.Name == "errcdgen" && v.Sel.Name == "NewCodeError" {
			return &Decl{
				chainFuncName: "NewCodeError",
			}
		}

		decl := traverseSingle(v.X)
		decl.chainFuncName = v.Sel.Name
		return decl
	}
	return nil
}

func DetectGoType(expr ast.Node) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind.String() == "INT" {
			return "int"
		} else if e.Kind.String() == "STRING" {
			return "string"
		} else {
			return "interface{}"
		}
	case *ast.CallExpr:
		return DetectGoType(e.Fun)
	case *ast.CompositeLit:
		switch et := e.Type.(type) {
		case *ast.ArrayType:
			goType := DetectGoType(et.Elt)
			return "[]" + goType
		case *ast.MapType:
			keyType := DetectGoType(et.Key)
			valueType := DetectGoType(et.Value)
			return "map[" + keyType + "]" + valueType
		}
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.Ident:
		return e.Name
	}

	return "" // unknown
}
