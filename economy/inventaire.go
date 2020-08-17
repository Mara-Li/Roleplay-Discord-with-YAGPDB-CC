{{/* Databases */}}
{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$id := .User.ID }}
{{if $name}}
	{{$user = $name}}
	{{$idperso := (toRune (lower $name))}}
	{{$id = ""}}
	{{range $idperso}}
		{{- $id = (print $id .)}}
	{{- end}}
	{{$id = (toInt $id)}}
	{{dbSet $id "rerollName" $name}}
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}

{{$serverEco := sdict}}
{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{/* Inventory */}}
{{$inv := sdict}}
{{if ($userEco.Get "Inventory")}}
	{{$inv = sdict ($userEco.Get "Inventory")}}
{{end}}


{{$desc := "Ton inventaire est vide ! Si le shop est ouvert, tu peux aller acheter des trucs !"}}
{{$footer := print "Page: 1 / 1 | #" $id }}
{{$end :=""}}
{{$cslice := cslice}}
{{range $k,$v := $inv}}
	{{$cslice = $cslice.Append (printf " :white_small_square: ** %-10v **  : [%v]" $k $v)}}
{{end}}
{{if $cslice}}
{{/* hell starts */}}
	{{$page := "1"}}
	{{if .CmdArgs}}
		{{$index := 0}}
		{{if ge (len .CmdArgs) 2}}
			{{$index = 1}}
		{{end}}
		{{$page = or (toInt (index .CmdArgs $index)) 1}}
		{{$page = toString $page}}
	{{end}}
		{{$end = roundCeil (div (toFloat (len $cslice)) 10)}}
	{{$footer = print "Page: " $page "/" $end "| #" $id }}
	{{$start := (mult 10 (sub $page 1))}}
	{{$stop := (mult $page 10)}}
	{{$data := ""}}
	{{if ge $stop (len $cslice)}}
		{{$stop = (len $cslice)}}
	{{end}}
	{{if not (eq $page "0")}}
		{{if and (le $start $stop) (ge (len $cslice) $start) (le $stop (len $cslice))}}
			{{range (seq $start $stop)}}
				{{$data = (print $data "\n" (index $cslice .))}}
			{{else}}
				{{$data = "Il n'y a rien ici..."}}
{{$footer = print "Page: " $page "/" $end " | #" $id }}
			{{end}}
		{{else}}
			{{$data = "Il n'y a rien ici..."}}
{{$footer = print "Page: " $page "/" $end " | #" $id }}
		{{end}}
			{{$desc = print "" $data ""}}
	{{end}}
	
{{/* hell ends */}}
{{end}}
{{$author := (joinStr " " "Inventaire de :" (title $user))}}

{{$id := sendMessageRetID nil (cembed "author" (sdict "name" $author "icon_url" "https://i.imgur.com/iUmz9Gi.png") "color" 0x8CBAEF "description" $desc "footer" (sdict "text" $footer) )}}
{{addMessageReactions nil $id "◀️" "▶️" "🗑️"}}
{{deleteTrigger 1}}
