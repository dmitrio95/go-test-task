package response

import (
	"bytes"
	"html/template"
)

const (
	sPageTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Entries}}
            <div><a href="{{.Link}}">{{.Name}}</a></div>
            <div>
                <table>
                    {{range $key, $value := .Data}}
                        <tr><td>{{$key}}</td><td>{{$value}}</td></tr>
                    {{else}}
                        <tr><td>No additional info</td></tr>
                    {{end}}
                </table>
            </div>
        {{else}}
            <div>No items</div>
        {{end}}
    </body>
</html>`
)

type htmlFormatter struct {
	pageTemplate *template.Template
}

func (f htmlFormatter) FormatResponse(data ResponseData) (string, error) {
	buf := bytes.NewBufferString("")
	err := f.pageTemplate.Execute(buf, data)
	return buf.String(), err
}

func newHTMLFormatter() *htmlFormatter {
	f := htmlFormatter{}
	f.pageTemplate = template.Must(template.New("page").Parse(sPageTemplate))

	return &f
}
