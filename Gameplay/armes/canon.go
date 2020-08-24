
{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}

{{$img := "https://i.imgur.com/YeIsRmw.png"}}

{{/* Groupe dictionnaire */}}
{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Get player */}}
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


{{/* get PA */}}
{{$pa := $groupe.Get (str $id)}}
{{if not $pa}}
	{{$groupe.Set (str $id) 4}}
	{{$pa = $groupe.Get (str $id)}}
{{end}}
{{dbSet .Server.ID "groupe" $groupe}}

{{/* Dict for weapon */}}
{{$arme := sdict}}
{{with (dbGet $id "arme")}}
	{{$arme = sdict .Value}}
{{end}}

{{$desc := ""}}

{{/* Function */}}
{{if gt $pa 0}}
	{{if not ($arme.Get "canon")}}
	  {{$arme.Set "canon" 1}}
		{{$desc = (joinStr "" "Il reste 19/20 balles de canon.")}}
	{{else}}
		{{$arme.Set "canon" (add ($arme.Get "canon") 1)}}
	  {{$y := ($arme.Get "canon")}}
	  {{$x := sub 20 $y}}
	  {{if lt (toFloat $y) (toFloat 20)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $x)) "/20 balles de canon.")}}
		{{else if eq (toFloat $y) (toFloat 12)}}
			{{$desc = "Dernière balle utilisée."}}
	  {{else}}
			{{ $desc = "Canon vide."}}
	  {{end}}
	{{end}}
{{else}}
	{{$desc = "PA insuffisants pour réaliser l'action."}}
{{end}}


{{$embed := cembed
"author" (sdict "name" $user "icon_url" $img)
"color"  0x6CAB8E
"description" $desc}}
{{ $idM := sendMessageRetID nil $embed }}
{{deleteMessage nil $idM 30}}

{{dbSet $id "arme" $arme}}