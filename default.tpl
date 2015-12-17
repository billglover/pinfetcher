{{range $i, $e := .}}
### [{{.Description}}]({{.Href}})
{{.Extended}}

*Tags:* {{range $tagIndex, $tag := .TagArray}}{{if $tagIndex}}, {{end}}`{{$tag}}`{{end}}

{{end}}