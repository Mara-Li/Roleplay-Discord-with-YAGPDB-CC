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


{{/* Dict for weapon */}}
{{$arme := sdict}}
{{with (dbGet $id "arme")}}
	{{$arme = sdict .Value}}
{{end}}

{{$fusil := reFindAllSubmatches `fusil` .Message.Content}}
{{$fusil2 := reFindAllSubmatches `fusil2` .Message.Content}}
{{$pistolet := reFindAllSubmatches `pistolet` .Message.Content}}
{{$pistol2 := reFindAllSubmatches `pistolet2` .Message.Content }}
{{$canon := reFindAllSubmatches `canon` .Message.Content }}

{{$desc := ""}}
{{$msgc:=""}}
{{$msgp:=""}}
{{$msgp2 := ""}}
{{$msgf := ""}}
{{$msgf2 := ""}}
{{$img := "https://i.imgur.com/YeIsRmw.png"}}

{{if $fusil}}
	{{if ($arme.Get "fusil")}}
		{{$y := $arme.Get "fusil"}}
		{{$x := sub 12 $y}}
		{{if lt (toFloat $y) (toFloat 12)}}
	  	{{ $desc = (joinStr "" "Il reste " (toString (toInt $x)) "/12 balles de fusil.")}}
		{{else}}
			{{ $desc = "Fusil vide."}}
		{{end}}
	{{else}}
		{{ $desc = "Fusil chargé : Il y a actuellement 12 balles à disposition."}}
	{{end}}

{{else if $fusil2 }}
	{{if ($arme.Get "fusil2")}}
		{{$y := $arme.Get "fusil2"}}
		{{$x := sub 12 $y}}
		{{if lt (toFloat $y (toFloat) 12)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $x)) "/12 balles de fusil secondaire.")}}
		{{else}}
			{{ $desc = "Fusil secondaire vide. "}}
		{{end}}
	{{else}}
		{{ $desc = "Fusil secondaire chargé : Il y a actuellement 12 balles à disposition."}}
	{{end}}

{{else if $pistolet}}
	{{if ($arme.Get "pistol")}}
		{{$a := $arme.Get "pistol"}}
		{{$b := sub 8 $a}}
		{{if lt (toFloat $a) (toFloat 8)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $b)) "/8 balles de pistolet.")}}
		{{else}}
			{{ $desc = "Pistolet vide."}}
		{{end}}
	{{else}}
		{{$desc = "Pistolet chargé : Il y a actuellement 8 balles à disposition."}}
	{{end}}

{{else if $pistol2}}
	{{if ($arme.Get "pistol2")}}
		{{$a := $arme.Get "pistol2"}}
		{{$b := sub 8 $a}}
		{{if lt (toFloat $a) (toFloat 8)}}
				{{ $desc = (joinStr "" "Il reste " (toString (toInt $b)) "/8 balles de pistolet secondaire.")}}
			{{else}}
				{{ $desc = "Pistolet secondaire vide."}}
		{{end}}
	{{else}}
		{{$desc = "Pistolet secondaire chargé : Il y a actuellement 8 balles à disposition."}}
	{{end}}

{{else if $canon}}
	{{if ($arme.Get "canon")}}
		{{$c := $arme.Get "canon"}}
		{{$d := sub 20 $c}}
		{{if lt (toFloat $c) (toFloat 20)}}
			{{ $desc = (joinStr "" "Il reste " (toString (toInt $d)) "/20 balles de canon.")}}
		{{else}}
			{{ $desc = "Canon vide."}}
		{{end}}
	{{else}}
		{{ $desc = "Canon chargé : Il y a actuellement 20 balles à disposition."}}
	{{end}}

{{else}}
		{{if ($arme.Get "canon")}}
			{{$ca := toFloat ($arme.Get "canon")}}
			{{$canon := (sub 20 $ca)}}
			{{if lt (toFloat $ca) (toFloat 20)}}
				{{$msgc = (joinStr "" "Il reste " (toString (toInt $canon)) "/20 balles dans votre canon.")}}
			{{else}}
				{{$msgc = "Canon vide."}}
			{{end}}
		{{else}}
			{{$msgc = "Canon chargé : Il y a actuellement 20 balles à disposition."}}
		{{end}}

		{{if ($arme.Get "fusil")}}
			{{$fu := $arme.Get "fusil"}}
			{{$fusil := (toFloat (sub 12 $fu))}}
			{{if lt (toFloat $fu) (toFloat 12)}}
				{{$msgf = (joinStr "" "Il reste " (toString (toInt $fusil)) "/12 balles de fusil")}}
			{{else}}
				{{$msgf = "Fusil vide."}}
			{{end}}
		{{else}}
			{{$msgf = "Fusil chargé : Il y a actuellement 12 balles à disposition."}}
		{{end}}

		{{if ($arme.Get "fusil2")}}
			{{$fu2 := ($arme.Get "fusil2")}}
			{{$fusil2 := (toFloat (sub 12 $fu2 ))}}
			{{if lt (toFloat $fu2) (toFloat 12)}}
				{{$msgf2 = (joinStr "" "Il reste " (toString (toInt $fusil2)) "/12 balles de fusil secondaire.")}}
			{{else}}
				{{$msgf2 = "Fusil secondaire vide."}}
			{{end}}
		{{else}}
			{{$msgf2 = "Fusil secondaire chargé : Il y a actuellement 12 balles à disposition."}}
		{{end}}

		{{if ($arme.Get "pistol")}}
			{{$pi := ($arme.Get "pistol")}}
			{{$pistol := (toFloat (sub 8 $pi))}}
			{{if lt (toFloat $pi) (toFloat 8)}}
				{{$msgp = (joinStr "" "Il reste " (toString (toInt $pistolet )) "/8 balles de pistolet.")}}
			{{else}}
				{{$msgp = "Pistolet vide."}}
			{{end}}
		{{else}}
			{{$msgp = "Pistolet chargé : Il y a actuellement 8 balles à disposition."}}
		{{end}}
		{{if ($arme.Get "pistol2")}}
			{{$pi2 := ($arme.Get "pistol2")}}
			{{$pistol2 = (toFloat (sub 8 $pi2))}}
			{{if lt (toFloat $pi2) (toFloat 8)}}
				{{$msgp2 = (joinStr "" "Il reste " (toString (toInt $pistol2 )) "/8 balles de pistolet secondaire.")}}
			{{else}}
				{{$msgp2 = "Pistolet secondaire vide."}}
			{{end}}
		{{else}}
			{{$msgp2 = "Pistolet secondaire chargé : Il y a actuellement 8 balles à disposition."}}
		{{end}}

		{{$desc = (joinStr "" ":white_small_square: **Canon** : " $msgc "\n:white_small_square: **Fusil** : " $msgf "\n:white_small_square: **Fusil secondaire** : " $msgf2 "\n:white_small_square: **Pistolet** : " $msgp "\n:white_small_square: **Pistolet secondaire** : " $msgp2 "\n")}}
{{end}}

{{$embed := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"color" 0x6CAB8E
	"description" $desc}}
{{sendMessage nil $embed}}

{{deleteTrigger 1}}