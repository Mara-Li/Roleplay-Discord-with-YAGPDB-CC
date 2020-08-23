{{/* User Variable */}}
{{$icon := $icon}}


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

{{/* Command Body */}}{{/* deep breaths */}}
{{if eq $store "open"}}
	{{$items := sdict}}
	{{with ($serverEco.Get "Items")}}
		{{$items = sdict .}}
	{{end}}
	{{$initem := ""}}
	{{$amount := 1}}
	{{with .CmdArgs}}
		{{$initem = ( title (index . 0))}}
		{{if ge (len .) 2}}
			{{$amount = or (toInt (index . 1)) 1}}
		{{end}}
	{{end}}

	{{$chargeur := reFind `(?i)chargeur` $initem}}
	{{$compo := reFind `(?i)(bc|lc|cb|sf|cu)` $initem}}
	{{if $compo}}
		{{if eq $compo "bc" "BC" "Bc"}}
			{{$initem = "[C] Biocomposant"}}
		{{else if eq $compo "lc" "LC" "Lc"}}
			{{$initem = "[C] Liquide Cytomorphe"}}
		{{else if eq $compo "cb" "CB" "Cb"}}
			{{$initem = "[C] Cellule Bionotropique"}}
		{{else if eq $compo "sf" "SF" "Sf"}}
			{{$initem = "[C] Substrat Ferreux"}}
		{{else if eq $compo "cu" "CU" "Cu"}}
			{{$initem = "[C] Composant Universel"}}
		{{end}}
	{{end}}

	{{if $chargeur}}
		{{$initem = reFind `(?i)(fusil|pistolet|canon)` $initem}}
		{{$initem = print "[CHARGEUR] " $initem}}
	{{end}}

	{{$item := sdict}}
	{{with $items.Get $initem}}
		{{$item = sdict .}}
	{{end}}
	{{$bal := (toInt ($userEco.Get "balance"))}}
	{{$inv := sdict}}
	{{with ($userEco.Get "Inventory")}}
		{{$inv = sdict .}}
	{{end}}
	{{if not $item}}
		ERREUR : **L'OBJET N'EXISTE PAS**
	{{else}}
	{{if eq (toInt $item.sellprice) 0}}
		{{$user = joinStr " " "Vente |" (title $user)}}
		{{sendMessage nil (cembed "author" (sdict "name" $user "icon_url" $icon) "color" 0x8CBAEF "description" (print "Le marchand n'achète pas ce genre d'objet...")) }}
	{{else if ge (toInt ($inv.Get $initem)) $amount}}
		{{$bal = add $bal (mult $item.sellprice $amount)}}
		{{$userEco.Set "balance" $bal}}
		{{$inv.Set $initem (sub ($inv.Get $initem) $amount)}}
		{{if eq ($inv.Get $initem) 0}}
			{{$inv.Del $initem}}
		{{end}}
		{{$userEco.Set "Inventory" $inv}}
		{{if not (eq (str $item.stock) "♾️")}}
			{{$item.Set "stock" (add $item.stock $amount)}}
		{{end}}
		{{$items.Set $initem $item}}
		{{$serverEco.Set "Items" $items}}
		{{$user = joinStr " " "Vente |" (title $user)}}
		{{sendMessage nil (cembed "author" (sdict "name" $user "icon_url" $icon) "color" 0x8CBAEF "description" (print "Vous avez vendu " $amount " " $initem " pour " (mult $item.sellprice $amount) " " $symbol " .")) }}
	{{else}}
		**ERREUR** : Quantité supérieure à celle possédée.
	{{end}}
	{{end}}
{{else}}
	{{$user = joinStr " " "Vente impossible |" (title $user)}}
		{{$embed := cembed
		"author" (sdict "name" $user "icon_url" $icon)
		"description" "Le magasin est actuellement indisponible, vous ne pouvez rien vendre. "
	"color" 0x8CBAEF}}
	{{sendMessage nil $embed}}
{{end}}
{{/* Database Updates */}}
{{dbSet $id "economy" $userEco}}
{{dbSet .Server.ID "economy" $serverEco}}
{{deleteTrigger 1}}
