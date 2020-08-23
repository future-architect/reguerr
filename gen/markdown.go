package gen

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gitlab.com/osaki-lab/reguerr"
	"io"
	"strings"
)

func GenerateMarkdown(w io.Writer, decls []*Decl) error {
	data := make([][]string, 0, len(decls))
	for _, v := range decls {

		if v.LogLevel == 0 {
			v.LogLevel = reguerr.DefaultErrorLevel
		}
		if v.StatusCode == 0 {
			v.StatusCode = reguerr.DefaultStatusCode
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
	table.SetHeader([]string{"Code", "Name", "LogLevel", "StatusCode", "Format"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

	return nil
}
