package gen

import (
	"fmt"
	"strconv"

	"gitlab.com/osaki-lab/reguerr"
)

type File struct {
	PkgName string
	Decls   []*Decl
}

type Decl struct {
	Name             string
	Code             string
	Format           string
	LogLevelEnable   bool
	LogLevel         reguerr.Level
	StatusCodeEnable bool
	StatusCode       int
	DisableErr       bool
	Labels           []Label
	chainFuncName    string // inside fields
	CallBuild        bool
}

type Label struct {
	Index  int
	Name   string
	GoType string
}

func (d Decl) ExistArgs() bool {
	return len(Analyze(d.Format)) > 0
}

func (d Decl) Args() string {
	var resp = ""

	verbs := Analyze(d.Format)

	argNo := 1
	labelMap := d.labelMap()
	for i, _ := range verbs {
		if resp != "" {
			resp += ","
		}

		label, ok := labelMap[i]
		if ok {
			resp += label.Name + " " + label.GoType
			continue
		}

		resp += "arg" + strconv.Itoa(argNo) + " interface{}"
		argNo++
	}

	return resp
}

func (d Decl) ArgValues() string {
	var resp = ""

	verbs := Analyze(d.Format)

	argNo := 1
	labelMap := d.labelMap()
	for i, _ := range verbs {
		if resp != "" {
			resp += ","
		}

		label, ok := labelMap[i]
		if ok {
			resp += label.Name
			continue
		}

		resp += "arg" + strconv.Itoa(argNo)
		argNo++
	}
	return resp
}

func (d Decl) labelMap() map[int]Label {
	m := map[int]Label{}
	for _, l := range d.Labels {
		m[l.Index] = l
	}
	return m
}

func (d Decl) MessageTemplate() string {
	if d.DisableErr {
		return fmt.Sprintf("New%s is the error indicating [%s] %s", d.Name, d.Code, d.Format)
	}
	return fmt.Sprintf("New%s is the error indicating [%s] %s: $err", d.Name, d.Code, d.Format)

}
