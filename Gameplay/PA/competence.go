{{/* Compétence img*/}}
{{$imga:="https://i.imgur.com/zNofnyh.png"}}
{{$imgs :="https://i.imgur.com/9iRdtbM.png" }}
{{$imgm := "https://i.imgur.com/FCy00x2.png"}}

{{/* Get joueur */}}

{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id:= .User.ID}}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{$user = title $user}}
{{$idict := str $id}}

{{if .CmdArgs}}
	{{if eq (index .CmdArgs 0) "court"}}
		{{if not (index .CmdArgs 1)}}
			**ERREUR** : N'oubliez pas d'indiquer le nom de l'attaque
		{{else}}
			{{$name := (index .CmdArgs 1)}}
			{{$arg := dbGet $id $name}}
			{{if not $arg}}
				{{dbSet $id "cdatq" $name}}
				{{dbSet $id $name 1}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Début de la recharge de la competence " $name "...\n Please wait...")
					"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID 735938256038002818 $embed }}
			{{else}}
				{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "La compétence " $name " est toujours en cooldown... \n Please wait...")
				"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{end}}
	{{else if (index .CmdArgs 0) "long"}}
		{{if not (index .CmdArgs 1)}}
			**ERREUR** : N'oubliez pas d'indiquer le nom de l'attaque
		{{else}}
			{{$name := (index .CmdArgs 1)}}
			{{$arg := dbGet $id $name}}
			{{if not $arg}}
				{{dbSet $id $name 1}}
				{{dbSet $id "cdsupp" $name}}
				{{ $embed := cembed
					"author" (sdict "name" $user "icon_url" $imga)
					"description" (joinStr "" "Début de la recharge de la competence " $name "...\n Please wait...")
					"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID 735938256038002818 $embed }}
			{{else}}
				{{ $embed := cembed
				"author" (sdict "name" $user "icon_url" $imga)
				"description" (joinStr "" "La compétence " $name " est toujours en cooldown... \n Please wait...")
				"color" 0xDFAA58}}
				{{ $idM := sendMessageRetID nil $embed }}
				{{deleteMessage nil $idM 30}}
			{{end}}
		{{end}}
	{{else}}
		**USAGE** : `$comp (court|long) "nom"`
	{{end}}
{{else}}
	**USAGE** : `$comp (court|long) "nom"`
{{end}}