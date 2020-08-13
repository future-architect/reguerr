package gen

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
)

func GenerateMarkdown(w io.Writer, decls []*Decl) error {
	data := make([][]string, 0, len(decls))
	for _, v := range decls {
		logLevel := fmt.Sprint(v.LogLevel)
		if v.LogLevel == 0 {
			logLevel = "default"
		}

		statusCode := fmt.Sprint(v.StatusCode)
		if v.StatusCode == 0 {
			statusCode = "default"
		}

		data = append(data, []string{
			v.Code,
			v.Name,
			logLevel,
			statusCode,
			v.Format,
		})
	}

	w.Write([]byte("# Error Code List"))
	fmt.Fprintln(w)
	fmt.Fprintln(w)

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Code", "Name", "LogLevel", "StatusCode", "Format"})
	table.AppendBulk(data)
	table.Render()

	return nil
}
