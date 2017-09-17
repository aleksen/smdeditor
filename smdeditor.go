package smdeditor

import (
	"bytes"

	"github.com/ponzu-cms/ponzu/management/editor"
)

const (
	smdCSS = "https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.css"
	smdJS  = "https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.js"
)

func startScript(fieldName string, p interface{}) string {
	return `
<script type="text/javascript">
$(function() {
	var $ = $ || window.jQuery,
		head = $('head'),
		existing_scripts = head.find('script'),
		script_loaded = false,
		existing_css = head.find('link[rel=stylesheet]'),
		css_loaded = false;

	function initSMD() {
		var md = new SimpleMDE({ 
			element: $("textarea[name=` + editor.TagNameFromStructField(fieldName, p) + `]")[0] 
		});
	}

	// set marker if resources are already in DOM
	for (var i = 0; i < existing_css.length; i++) {
		if ($(existing_css[i]).attr('href') === '` + smdCSS + `') {
			css_loaded = true;
		}
	}

	for (var i = 0; i < existing_scripts.length; i++) {
		if ($(existing_scripts[i]).attr('src') === '` + smdJS + `') {
			script_loaded = true;
		}
	}

	// if not present, append to head
	if (!css_loaded) {
		var css = $('<link>');
		css.attr('rel', 'stylesheet');
		css.attr('href', '` + smdCSS + `');
		head.append(css);
	}

	if (!script_loaded) {
		var script = $('<script>');
		head.append(script);
		script.on('load', initSMD);
		script.attr('src', '` + smdJS + `');
	} else {
		if (SimpleMDE !== undefined) {
			initSMD();
		}
	}
});
</script>`
}

// Input returns an HTML textarea element pre-configured for Markdown support
func Input(fieldName string, p interface{}, attrs map[string]string) []byte {
	html := &bytes.Buffer{}
	_, err := html.WriteString(startScript(fieldName, p))
	if err != nil {
		return []byte("Markdown Editor failed.")
	}

	e := editor.NewElement("textarea", attrs["label"], fieldName, p, attrs)
	_, err = html.Write(editor.DOMElement(e))
	if err != nil {
		return []byte("DOM Element failed to be written in Markdown Editor")
	}

	return html.Bytes()
}
