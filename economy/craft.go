{{/* Databases */}}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{range $idperso}}
		{{- $id = add $id . }}
	{{- end}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}


{{$target := .User.ID}}
{{$mention := ""}}

{{$remove := reFind `^\$use` .Message.Content}}
{{$craft := reFind `^\$add` .Message.Content}}


{{if .CmdArgs}}
	{{if $craft}}
		{{if $name}}
			{{$target = $id}}
			{{$user = $name}}
			{{$mention = $name}}
		{{else}}
			{{$target = .User.ID}}
			{{$mention = joinStr "" "<@" $target ">"}}
			{{$user = (getMember $target).Nick}}
		{{end}}
		{{$userEco := sdict}}
		{{if $target}}
			{{with (dbGet $target "economy")}}
				{{$userEco = sdict .Value}}
			{{end}}
			{{$inv := sdict}}
			{{with ($userEco.Get "Inventory")}}
				{{$inv = sdict .}}
			{{end}}
			{{$item := title (index .CmdArgs 0)}}
			{{$amount := 1}}
			{{if gt (len .CmdArgs) 1 }}
				{{$amount = (toInt (index .CmdArgs 0))}}
				{{$item = title (index .CmdArgs 1)}}
				{{if $name}}
					{{$item = title (index .CmdArgs 0)}}
					{{$amount = 1}}
					{{if gt (len .CmdArgs) 2}}
						{{$amount = (toInt (index .CmdArgs 0))}}
						{{$item = title (index .CmdArgs 1)}}
					{{end}}
				{{end}}
			{{end}}
			{{if $inv.Get $item}}
				{{$inv.Set $item (add (toInt ($inv.Get $item)) $amount)}}
			{{else}}
				{{if ne $amount 0}}
					{{$inv.Set $item $amount}}
				{{end}}
			{{end}}
			{{if eq (toInt $amount) 1}}
			L'objet {{$item}} a été ajouté à l'inventaire de {{$mention}}
			{{else}}
				{{$amount}} {{$item}} ont été ajouté à l'inventaire de {{$mention}}
			{{end}}
			{{$userEco.Set "Inventory" $inv}}
			{{dbSet $target "economy" $userEco}}
		{{end}}
	{{end}}
{{end}}

{{if .CmdArgs}}
	{{if $remove}}
		{{if $name}}
			{{$target = $id}}
			{{$user = $name}}
			{{$mention = $name}}
		{{else}}
			{{$target = .User.ID}}
			{{$mention = joinStr "" "<@" $target ">"}}
			{{$user = (getMember $target).Nick}}
		{{end}}
		{{$userEco := sdict}}
		{{if $target}}
			{{with (dbGet $target "economy")}}
				{{$userEco = sdict .Value}}
			{{end}}
			{{$inv := sdict}}
			{{with ($userEco.Get "Inventory")}}
				{{$inv = sdict .}}
			{{end}}
			{{$item := (title (index .CmdArgs 0))}}
			{{$amount := 1}}
			{{if gt (len .CmdArgs) 1 }}
				{{if not $name}}
					{{$amount = (toInt (index .CmdArgs 0))}}
					{{if eq $amount 0}}
						{{$amount = str (index .CmdArgs 0)}}
					{{end}}
					{{$item = title (index .CmdArgs 1)}}
				{{else}}
					{{$item = title (index .CmdArgs 1)}}
					{{if gt (len .CmdArgs) 2}}
						{{$item = title (index .CmdArgs 2)}}
						{{$amount = (toInt (index .CmdArgs 1))}}
						{{if eq $amount 0}}
							{{$amount = str (index .CmdArgs 1)}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
			{{if $inv.Get $item}}
				{{$value := $inv.Get $item}}
				{{if and (ne (toInt $amount) (toInt $value)) (ne (toInt $amount) 0)}}
					{{if lt (toInt $amount) (toInt $value)}}
						{{$inv.Set $item (sub (toInt $value) (toInt $amount))}}
					{{else}}
						{{$inv.Set $item (toInt $amount)}}
					{{end}}
				{{else if eq (toInt $amount) (toInt $value)}}
					{{$inv.Del $item}}
				{{else if eq $amount "all"}}
					{{$inv.Del $item}}
				{{end}}
				{{$userEco.Set "Inventory" $inv}}
				{{if eq (str $amount) "all"}}
					{{$amount = "tous les"}}
				{{end}}
				{{$mention}} a utilisé {{$amount}} {{$item}} de son inventaire.
				{{dbSet $target "economy" $userEco}}
			{{else}}
				{{$mention}} : cet objet n'est pas présent dans votre inventaire : vous ne pouvez donc pas l'utiliser.
			{{end}}
		{{end}}
	{{end}}
{{end}}
