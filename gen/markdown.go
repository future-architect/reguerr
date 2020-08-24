package gen

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"math"
	"strings"
)

func GenerateMarkdown(w io.Writer, decls []*Decl, opts ...Option) error {
	setting := NewSetting()
	for _, opt := range opts {
		opt(setting)
	}

	data := make([][]string, 0, len(decls))
	for _, v := range decls {
		if !v.LogLevelEnable {
			// when unsigned loglevel then setting from option
			v.LogLevel = setting.Level
		}
		if v.StatusCode == 0 {
			v.StatusCode = setting.StatusCode
		}

		data = append(data, []string{
			v.Code,
			v.Name,
			strings.Replace(fmt.Sprint(v.LogLevel), "Level", "", 1),
			fmt.Sprint(v.StatusCode),
			v.Format,
		})
	}

	w.Write([]byte("# Error Code List"))
	fmt.Fprintln(w)
	fmt.Fprintln(w)

	table := tablewriter.NewWriter(w)
	table.SetColWidth(math.MaxInt32) // tablewriter default column size is 30. this is too small.
	table.SetHeader([]string{"Code", "Name", "LogLevel", "StatusCode", "Format"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

	return nil
}
