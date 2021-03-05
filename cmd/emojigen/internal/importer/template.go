package importer

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

const keywordTemplateString = `// Code generated based on latest emoji dataset. DO NOT EDIT.

package {{ .Package }}
{{- define "image-data" -}}
{ {{ .Unified | quote }}, {{ .Character | quote }}, {{ .SheetX }}, {{ .SheetY }}, {{ .AddedIn | quote }}, map[Platform]bool{ PlatformApple: {{ .HasImgApple }}, PlatformGoogle: {{ .HasImgGoogle }}, PlatformTwitter: {{ .HasImgTwitter }}, PlatformFacebook: {{ .HasImgFacebook }} }, {{ .Obsoletes | quote }}, {{ .ObsoletedBy | quote }} }
{{- end -}}

// AllEmojis contains the list of all available keywords.
var All = []Info {
	{{- range .Emojis }}
	{ {{.ShortName | quote }}, {{.Category | quote }}, {{ .Text | quote }}, []string{ {{$s := separator ", "}}{{ range .ShortNames }}{{ call $s }}{{ . | quote }}{{ end }} }, ImageData{{ template "image-data" .EmojiImageData }}, {{ with .SkinVariations }}map[Modifier]ImageData{ {{$s := separator ", "}}{{- range $key, $value := . }}{{ call $s }}{{ $key | modifierConstant }}: {{ template "image-data" $value }}{{- end }}}{{else}}nil{{end}} },
	{{- end }}
}
`

var keywordTemplate = template.Must(
	template.
		New("keywords").
		Funcs(
			template.FuncMap{
				"separator":        separator,
				"quote":            quote,
				"modifierConstant": modifierConstant,
			},
		).
		Parse(keywordTemplateString),
)

// RenderTemplate renders the dataset template to the given io.Writer.
func RenderTemplate(w io.Writer, packageName string, emojis []EmojiInfo) error {
	return keywordTemplate.Execute(
		w,
		map[string]interface{}{
			"Package": packageName,
			"Emojis":  emojis,
		},
	)
}

// copied from "github.com/Masterminds/sprig"
func quote(str ...interface{}) string {
	out := make([]string, 0, len(str))
	for _, s := range str {
		if s != nil {
			out = append(out, fmt.Sprintf("%q", strval(s)))
		}
	}
	return strings.Join(out, " ")
}

// copied from "github.com/Masterminds/sprig"
func strval(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// separator emits the an empty string the first time
// and then the separator every subsequent time.
func separator(s string) func() string {
	i := -1
	return func() string {
		i++
		if i == 0 {
			return ""
		}
		return s
	}
}

func modifierConstant(src string) string {
	switch src {
	case "1F3FB":
		return "SkinToneLight"
	case "1F3FC":
		return "SkinToneMediumLight"
	case "1F3FD":
		return "SkinToneMedium"
	case "1F3FE":
		return "SkinToneMediumDark"
	case "1F3FF":
		return "SkinToneDark"
	default:
		panic(fmt.Errorf("unsupported modifier string %s", src))
	}
}
