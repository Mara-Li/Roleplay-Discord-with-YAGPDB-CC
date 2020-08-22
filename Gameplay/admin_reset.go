{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := ""}}
{{if .CmdArgs}}
	{{if not $name}}
	{{$id = (userArg (index .CmdArgs 0)).ID}}
	{{$user = (getMember $id).Nick}}
	{{else if $name}}
		{{$user = $name}}
		{{$idperso := (toRune (lower $name))}}
		{{range $idperso}}
			{{- $id = add $id . }}
		{{- end}}
	{{end}}
{{end}}
{{$user = title $user}}
{{$idict := str $id}}

{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{$bool := "false"}}
{{range $i, $j := $groupe}}
	{{- if eq $idict $i}}
		{{- $bool = "true"}}
	{{- end -}}
{{end}}


{{dbDel $id "recharge"}}
{{dbDel $id "arme"}}
{{$atq := (dbGet $id "cdatq").Value}}
{{$supp := (dbGet $id "cdsupp").Value}}
{{dbDel $id $atq}}
{{dbDel $id $supp}}
{{dbDel $id "cdatq"}}
{{dbDel $id "cdsupp"}}

{{if eq $bool "true"}}
	{{$groupe.Del $idict}}
{{end}}
{{dbSet .Server.ID "groupe" $groupe}}

Toutes les variables d'armes et de PA de {{$user}} ont été supprimé de la DB !
