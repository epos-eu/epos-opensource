package display

import (
	"fmt"
	"net/url"

	"github.com/epos-eu/epos-opensource/db/sqlc"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func InfraList(rows [][]any, headers []string, title string) {
	if len(rows) == 0 {
		Info("No installed environments found")
		return
	}
	t := table.NewWriter()
	t.SetTitle(title)
	t.SetStyle(table.StyleRounded)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Colors = text.Colors{text.FgYellow, text.Bold}
	t.Style().Color.Border = text.Colors{text.FgGreen}
	t.Style().Color.Footer = text.Colors{text.FgGreen}
	t.Style().Color.Separator = text.Colors{text.FgGreen}
	t.Style().Color.Header = text.Colors{text.FgCyan}
	colConfigs := make([]table.ColumnConfig, len(headers))
	for i := range headers {
		colConfigs[i] = table.ColumnConfig{Number: i + 1, AlignHeader: text.AlignCenter}
	}
	t.SetColumnConfigs(colConfigs)
	headerAny := make([]any, len(headers))
	for i, h := range headers {
		headerAny[i] = h
	}
	t.AppendHeader(table.Row(headerAny))
	for _, row := range rows {
		t.AppendRow(table.Row(row))
	}
	rowMerge := table.RowConfig{
		AutoMerge:      true,
		AutoMergeAlign: text.AlignLeft,
	}
	footer := make([]any, len(headers))
	for i := range footer {
		footer[i] = copyright
	}
	t.AppendFooter(table.Row(footer), rowMerge)
	fmt.Println(t.Render())
}

func DockerList(dockers []sqlc.Docker, title string) {
	rows := make([][]any, len(dockers))
	for i, d := range dockers {
		gatewayURL, err := url.JoinPath(d.ApiUrl, "ui")
		if err != nil {
			panic("error building gateway url") // TODO
		}
		backofficeURL, err := url.JoinPath(d.BackofficeUrl, "home")
		if err != nil {
			panic("error building backoffice url") // TODO
		}
		rows[i] = []any{d.Name, d.Directory, d.GuiUrl, backofficeURL, gatewayURL}
	}
	headers := []string{"Name", "Directory", "GUI URL", "Backoffice URL", "API URL"}
	InfraList(rows, headers, title)
}

func KubernetesList(kubes []sqlc.Kubernetes, title string) {
	rows := make([][]any, len(kubes))
	for i, k := range kubes {
		gatewayURL, err := url.JoinPath(k.ApiUrl, "ui")
		if err != nil {
			panic("error building gateway url") // TODO
		}
		backofficeURL, err := url.JoinPath(k.BackofficeUrl, "home")
		if err != nil {
			panic("error building backoffice url") // TODO
		}
		rows[i] = []any{k.Name, k.Directory, k.Context, k.GuiUrl, backofficeURL, gatewayURL}
	}
	headers := []string{"Name", "Directory", "Context", "GUI URL", "Backoffice URL", "API URL"}
	InfraList(rows, headers, title)
}
