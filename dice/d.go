{{$name := reFind `(\#\S*)` .Message.Content}}
{{$name = joinStr "" (split $name "#")}}
{{$user := .Member.Nick}}
{{$img := "https://i.imgur.com/Khq39Um.png"}}
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

{{$user = title $user}}

{{$stats := sdict}}
{{with (dbGet $id "stats")}}
	{{$stats = sdict .Value}}
{{end}}

{{$d:= (randInt 1 10)}}

{{$arg1:=""}}
{{$arg2:= ""}}
{{$v:=$d}}
{{$arg := 0}}
{{$argm := 0}}
{{$c:=""}}


{{$msg := ""}}

{{$seuil := (toInt 8)}}
{{$idb := (toInt 0)}}

{{$manuel := reFind `s[1-9]` .Message.Content}}
{{$force:=reFindAllSubmatches `(force|Force)` .Message.Content}}
{{$agi := reFindAllSubmatches `(agilit(é|e)|Agilit(é|e)|\dagi|\dAgi)` .Message.Content}}
{{$endu := reFindAllSubmatches `(Endurance|endurance|\dendu|\dEndu)` .Message.Content}}
{{$preci := reFindAllSubmatches `(Pr(é|e)cision|pr(é|e)cision|pr(é|e)ci|Pr(é|e)ci)` .Message.Content}}
{{$intell := reFindAllSubmatches `(Intelligence|intelligence|intell|Intell|intel|Intel)` .Message.Content}}
{{$karma := reFindAllSubmatches `(Karma|karma)` .Message.Content}}

{{if $manuel}}
	{{$seuil = (toInt (joinStr "" (split $manuel "s")))}}
	{{$idb = (toInt 0)}}
{{else if $force}}
	{{$seuil = (toInt ($stats.Get "force"))}}
	{{$idb = (toInt ($stats.Get "i_force"))}}
{{else if $agi}}
	{{$seuil = (toInt ($stats.Get "agi"))}}
	{{$idb = (toInt ($stats.Get "i_agi"))}}
{{else if $endu}}
	{{$seuil = (toInt ($stats.Get "endurance"))}}
	{{$idb = (toInt ($stats.Get "i_endu"))}}
{{else if $preci}}
	{{$seuil = (toInt ($stats.Get "preci"))}}
	{{$idb = (toInt ($stats.Get "i_preci"))}}
{{else if $intell}}
	{{$seuil = (toInt ($stats.Get "intelligence"))}}
	{{$idb = (toInt ($stats.Get "i_intel"))}}
{{else if $karma}}
	{{$seuil = (toInt ($stats.Get "karma"))}}
	{{$idb = (toInt 0)}}
{{else}}
	{{$seuil = (toInt 8)}}
	{{$idb = (toInt 0)}}
{{end}}

{{$imp := ""}}

{{if le $idb 1}}
	{{$imp = "Implant"}}
{{else}}
	{{$imp = "Implants"}}
{{end}}
{{$mimp := $idb}}


{{ if .CmdArgs}}
	{{$c = joinStr " " .CmdArgs}}
	{{if $name}}
		{{$c = joinStr " " (split $c $name) }}
		{{$c = joinStr " " (split $c "#")}}
	{{end}}

	{{if $manuel}}
		{{$c = joinStr " " (split $c $manuel) }}
	{{end}}

	{{if ne (toFloat (index .CmdArgs 0)) (toFloat 0)}}
		{{$i:=(toInt (index .CmdArgs 0)) }}
		{{$x := $i}}
		{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
			{{$arg = 1}}
			{{if eq (toInt (index .CmdArgs 0)) (toInt 1)}}
				{{$arg1 = "Pénalité"}}
			{{else}}
				{{$arg1 = "Pénalités"}}
			{{end}}
		{{else}}
			{{$i = mult $i (toInt -1)}}
			{{$x = mult $x (toInt -1)}}
			{{$arg = 0}}
			{{$arg1 = "Bonus"}}

		{{end}}

		{{if eq $arg 0}}
			{{$i = add $i $idb}}
			{{$d = sub $d $i}}
			{{$mimp = $i}}
		{{else}}
			{{$d = add $i $d}}
			{{if ge (toFloat (len .CmdArgs)) (toFloat 2 )}}
				{{if eq (toInt (index .CmdArgs 1)) 0}}
					{{$d = sub $d $idb}}
					{{$mimp = $idb}}
				{{end}}
			{{end}}
		{{end}}

		{{if gt $d (toInt 10)}}
			{{$d = toInt 10}}
		{{else if lt $d (toInt 0)}}
			{{$d = (toInt 0)}}
		{{end}}

		{{if eq $d (toInt 0)}}
			{{$msg ="**Ultra critique** "}}
		{{else if eq $d (toInt 1)}}
			{{$msg = "**Réussite critique** "}}
		{{else if le $d $seuil}}
				{{if ge $mimp (toInt 1)}}
					{{if eq $d $seuil}}
						{{$msg = "**Echec** "}}
					{{else}}
						{{$msg = "**Réussite** "}}
					{{end}}
				{{else}}
					{{$msg = "**Réussite** "}}
				{{end}}
		{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
			{{$msg = "**Echec** "}}
		{{else if eq $d (toInt 10)}}
			{{$msg = "**Echec critique** "}}
		{{else if eq $d (toInt 9)}}
			{{$msg = "**Echec** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}

			{{if eq (toFloat 1)  (toFloat (len .CmdArgs))}}
				{{$embed := cembed
				"author" (sdict "name" $user "icon_url" $img)
				"description" (joinStr "" $msg "\n"
				"<:next:723131844643651655> [*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
				"color" 0x63AFE1}}
				{{sendMessage nil $embed}}
			{{else}}
				{{$c =  joinStr " " (slice .CmdArgs 1)}}

				{{if $manuel}}
					{{$c = joinStr " " (split $c $manuel)}}
				{{end}}

				{{if $name}}
					{{$c = joinStr " " (split $c $name)}}
					{{$c = joinStr " " (split $c "#")}}
				{{end}}

				{{$c = joinStr " " $c }}

				{{if eq (toFloat (index .CmdArgs 1)) (toFloat 0)}}
					{{$embed := cembed
					"author" (sdict "name" $user "icon_url" $img)
					"description" (joinStr "" $msg " ▬ " $c " : " "\n"
					"<:next:723131844643651655> [*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
					"color" 0x63AFE1}}
					{{sendMessage nil $embed}}

				{{else if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
					{{$m := (toInt (index .CmdArgs 1)) }}
					{{$y := $m}}
					{{if gt $m (toInt 0)}}
						{{$argm = 1}}
						{{if eq $m (toInt 1)}}
							{{$arg2 = "Pénalité "}}
						{{else}}
							{{$arg2 = "Pénalités "}}
						{{end}}
					{{else}}
						{{$argm = 0}}
						{{$m = mult $m (toInt -1)}}
						{{$y = mult $y (toInt -1)}}
						{{$arg2 = "Bonus "}}
					{{end}}


					{{if eq $arg 0}}
						{{if eq $argm 0}}
							{{$d = sub $d $m}}
						{{else}}
							{{$d = add $m $d}}
						{{end}}
					{{else}}
						{{if eq $argm 0}}
							{{$m = add $m $idb}}
							{{$mimp = $m}}
							{{$d = sub $d $m}}
						{{else}}
							{{$d = add $m $d}}
							{{$mimp = $idb}}
						{{end}}
					{{end}}


					{{if gt $d (toInt 10)}}
						{{$d = toInt 10}}
					{{else if lt $d (toInt 0)}}
						{{$d = (toInt 0)}}
					{{end}}

					{{if eq $d (toInt 0)}}
						{{$msg ="**Ultra critique** "}}
					{{else if eq $d (toInt 1)}}
						{{$msg = "**Réussite critique** "}}
					{{else if le $d $seuil}}
							{{if ge $mimp (toInt 1)}}
								{{if eq $d $seuil}}
									{{$msg = "**Echec** "}}
								{{else}}
									{{$msg = "**Réussite** "}}
								{{end}}
							{{else}}
								{{$msg = "**Réussite** "}}
							{{end}}
					{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
						{{$msg = "**Echec** "}}
					{{else if eq $d (toInt 10)}}
						{{$msg = "**Echec critique** "}}
					{{else if eq $d (toInt 9)}}
						{{$msg = "**Echec** "}}
					{{else}}
						{{$msg = " "}}
					{{end}}

				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$embed := cembed
						"author" (sdict "name" $user "icon_url" $img)
						"description" (joinStr "" $msg "\n"
						"<:next:723131844643651655>[*Dé : " $d " (" $v ") | " $x " : " $arg1 " | " $y " " $arg2 " | " $imp " : " $idb " | Seuil : " $seuil "*]")
						"color" 0x63AFE1}}
					{{sendMessage nil $embed}}

					{{else}}
					{{$c =  joinStr " " (slice .CmdArgs 2)}}
					{{if $manuel}}
						{{$c = joinStr " " (split $c $manuel)}}
					{{end}}

					{{if $name}}
						{{$c = joinStr " " (split $c $name)}}
						{{$c = joinStr " " (split $c "#")}}
					{{end}}

					{{$c = joinStr " " $c }}

						{{$embed := cembed
							"author" (sdict "name" $user "icon_url" $img)
						"description" (joinStr "" $msg " ▬ " $c " : " "\n"
						"<:next:723131844643651655>[*Dé : " $d " (" $v ") | " $arg1 " : " $x " | " $arg2 " : " $y " | " $imp " : " $idb " | Seuil : " $seuil "*]")
						"color" 0x63AFE1}}
						{{sendMessage nil $embed}}


				{{end}}
			{{end}}
		{{end}}

	{{else}}
		{{$d = sub $d $idb}}

		{{if gt $d (toInt 10)}}
			{{$d = toInt 10}}
		{{else if lt $d (toInt 0)}}
			{{$d = (toInt 0)}}
		{{end}}

		{{if eq $d (toInt 0)}}
			{{$msg ="**Ultra critique** "}}
		{{else if eq $d (toInt 1)}}
			{{$msg = "**Réussite critique** "}}
		{{else if le $d $seuil}}
				{{if ge $idb (toInt 1)}}
					{{if eq $d $seuil}}
						{{$msg = "**Echec** "}}
					{{else}}
						{{$msg = "**Réussite** "}}
					{{end}}
				{{else}}
					{{$msg = "**Réussite** "}}
				{{end}}
		{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
			{{$msg = "**Echec** "}}
		{{else if eq $d (toInt 10)}}
			{{$msg = "**Echec critique** "}}
		{{else if eq $d (toInt 9)}}
			{{$msg = "**Echec** "}}
		{{else}}
			{{$msg = " "}}
		{{end}}

		{{$embed := cembed
			"author" (sdict "name" $user "icon_url" $img)
		"description" (joinStr "" $msg " ▬ " (joinStr " " $c) " : " "\n"
	"<:next:723131844643651655>[*Dé : " $d  " (" $v ") "  " | " $imp " : " $idb " | Seuil : " $seuil "* ]")
		"color" 0x63AFE1}}
		{{sendMessage nil $embed}}
	{{end}}

{{else}}
{{if gt $d (toInt 10)}}
	{{$d = toInt 10}}
{{else if lt $d (toInt 0)}}
	{{$d = (toInt 0)}}
{{end}}

{{if eq $d (toInt 0)}}
	{{$msg ="**Ultra critique** "}}
{{else if eq $d (toInt 1)}}
	{{$msg = "**Réussite critique** "}}
{{else if le $d $seuil}}
		{{if ge $idb (toInt 1)}}
			{{if eq $d $seuil}}
				{{$msg = "**Echec** "}}
			{{else}}
				{{$msg = "**Réussite** "}}
			{{end}}
		{{else}}
			{{$msg = "**Réussite** "}}
		{{end}}
{{else if or (gt $d $seuil) (lt $d (toInt 9))}}
	{{$msg = "**Echec** "}}
{{else if eq $d (toInt 10)}}
	{{$msg = "**Echec critique** "}}
{{else if eq $d (toInt 9)}}
	{{$msg = "**Echec** "}}
{{else}}
	{{$msg = " "}}
{{end}}
	{{$embed := cembed
	"author" (sdict "name" $user "icon_url" $img)
	"description" (joinStr "" $msg "\n"
	"<:next:723131844643651655>[*Dé : " $d "* ]")
	"color" 0x63AFE1}}
	{{sendMessage nil $embed}}
{{end}}

{{deleteTrigger 1}}
