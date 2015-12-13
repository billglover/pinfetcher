# Blog Post

{{range $i, $e := .}}
## [{{.Description}}]({{.Href}})
{{.Extended}}

{{range $tagIndex, $tag := .TagArray}}{{if $tagIndex}}, {{end}}`{{$tag}}`{{end}}

{{end}}