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
{{else if eq (len $user) 0}}
	{{$user = .User.Username}}
{{end}}

{{/* Databases */}}
{{$store := (dbGet .Server.ID "store").Value}}
{{$userEco := sdict}}
{{with (dbGet $id "economy")}}
	{{$userEco = sdict .Value}}
{{end}}
{{$serverEco := sdict}}

{{with (dbGet .Server.ID "economy")}}
	{{$serverEco = sdict .Value}}
{{end}}

{{$symbol := ""}}
{{if $serverEco.Get "symbol"}}
	{{$symbol = $serverEco.Get "symbol"}}
{{end}}

{{/* Command Body */}}
{{if eq $store "open"}}
	{{$items := sdict}}
	{{with ($serverEco.Get "Items")}}
		{{$items = sdict .}}
	{{end}}

	{{$initem := ""}}
	{{$amount := 1}}
	{{$desc := ""}}

	{{with .CmdArgs}}
	  {{$initem = ( (index . 0))}}
	  {{if ge (len .) 2}}
	    {{$amount = or (toInt (index . 1)) 1}}
	  {{end}}
	{{end}}

	{{$item := sdict}}

	{{with $items.Get $initem}}
		{{$item = sdict .}}
	{{end}}

	{{$bal := (toInt ($userEco.Get "balance"))}}
	{{if le $bal (toInt 0)}}
		{{$bal = 0}}
	{{end}}
	{{$inv := sdict}}
	{{with ($userEco.Get "Inventory")}}
		{{$inv = sdict .}}
	{{end}}

	{{$info := sdict}}
	{{with ($userEco.Get "desc")}}
		{{$info = sdict .}}
	{{end}}

	{{if not $item}}
		{{sendMessage nil "L'objet n'existe pas."}}
	{{else}}
	  {{if le $bal (toInt (mult .buyprice $amount))}}
			{{sendMessage nil "Tu n'as pas assez d'argent pour ça."}}
		{{else}}
		{{with $item}}
		    {{$infcheck := false}}
				{{if eq (str .stock) "♾️"}}
					{{$infcheck = true}}
				{{end}}
		    {{$stock := .stock}}
				{{if $infcheck}}
					{{$stock = 1000000000000000000}}
				{{end}}
		    {{if not (ge (toFloat $stock) (toFloat $amount) )}}
					{{sendMessage nil "Il n'y a pas de stock actuellement pour cet objet."}}
				{{else}}
		    	{{$bal = sub $bal (mult .buyprice $amount)}}
					{{if not $infcheck}}
					 	{{$item.Set "stock" (sub $stock $amount)}}
					{{end}}
					{{if .sii}}
						{{$i := ($items.Get ( $initem))}}
						{{with $i}}
							{{$desc = .desc}}
						{{end}}
						{{$inv.Set $initem (add (toInt ($inv.Get $initem)) $amount)}}
						{{$info.Set (joinStr " " $initem) $desc}}
					{{end}}
		      {{if not .reply}}
					{{$user = joinStr " " "Achat | " (title $user)}}
		        {{sendMessage nil (cembed
		          "author" (sdict "name" $user "icon_url" "https://i.imgur.com/3uiVkvv.png")
		          "description" (print "Vous avez acheté " $amount " " $initem " pour " (mult .buyprice $amount) " " $symbol ".")
		          "color" 0x8CBAEF)}}
		      {{end}}
		      {{$items.Set $initem $item}}
		      {{$serverEco.Set "Items" $items}}
		      {{$userEco.Set "balance" $bal}}
		      {{$userEco.Set "Inventory" $inv}}
				{{end}}
	    {{end}}
	  {{end}}
	{{end}}
	{{/* Database Updates */}}
	{{dbSet .Server.ID "economy" $serverEco}}
	{{dbSet $id "economy" $userEco}}
{{else}}
{{$user = joinStr " " "Achat impossible | " (title $user)}}
	{{sendMessage nil (cembed
		"author" (sdict "name" $user "icon_url" "https://i.imgur.com/3uiVkvv.png")
		"description" "La boutique est actuellement indisponible ! Vous ne pouvez rien acheter. "
		"color" 0x8CBAEF)}}
{{end}}
{{deleteTrigger 1}}
