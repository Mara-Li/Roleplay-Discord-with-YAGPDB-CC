{{$groupe := sdict}}
{{with (dbGet .Server.ID "groupe")}}
	{{$groupe = sdict .Value}}
{{end}}

{{/* Seuil & dictionnaire */}}
{{$seuil := sdict}}
{{with (dbGet .Server.ID "seuil")}}
	{{$seuil = sdict .Value}}
{{end}}

{{/* Tour count :Incr each call unless CmdArgs = reset */}}
{{if .CmdArgs}}
	{{dbSet 0 "turn" 1}}
{{else}}
	{{$x := dbIncr 0 "turn" 1}}
{{end}}
{{$turn := (dbGet 0 "turn").Value}}
{{if not $turn}}
	{{$turn := 1}}
	{{dbSet 0 "turn" 1}}
{{end}}

{{range $i, $j := $groupe}}
	{{if le $j 0}}
		{{$j = 4}}
	{{else if le $j 2}}
		{{$j = 4}}
	{{else if ge $j 2}}
		{{$j = 6}}
	{{end}}
	{{$groupe.Set $i $j}}
{{end}}

{{$icon := (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png")}}
{{$embed := cembed
	"author" (sdict "name" "Vaisseau Nucleus" "icon_url" $icon)
	"title" (joinStr " " "TOUR :" (str (toInt $turn)))
	"color" 0x6B54BE
	"timestamp" currentTime}}
{{sendMessage nil $embed}}
{{dbSet .Server.ID "groupe" $groupe}}