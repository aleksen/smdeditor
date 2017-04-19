package smdeditor

import (
	"bytes"

	"github.com/ponzu-cms/ponzu/management/editor"
)

var loaded bool

func Input(fieldName string, p interface{}, attrs map[string]string) []byte {
	html := &bytes.Buffer{}
	if !loaded {
		html.WriteString(`<link rel="stylesheet" href="https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.css">
<script src="https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.js"></script>
`)
		loaded = true
	}

	e := editor.NewElement("textarea", attrs["label"], fieldName, p, attrs)
	html.Write(editor.DOMElement(e))
	html.WriteString(`<script>
var simplemde = new SimpleMDE({ element: $("textarea[name=` + editor.TagNameFromStructField(fieldName, p) + `]")[0] });
</script>
`)
	return html.Bytes()
}
