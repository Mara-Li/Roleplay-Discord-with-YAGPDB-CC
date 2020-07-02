{{$user:= .Member.Nick}}
  {{if eq (len $user) 0}}
    {{$user = .User.Username}}
  {{end}}

{{$d:= (randInt 1 10)}}
{{$real:= $d}}
{{$arg1:=""}}
{{$arg2:= ""}}
{{$res := ""}}
{{$m := $d}}
{{$i := $d}}
{{$x := $i}}
{{$y := $m}}
{{$v := (joinStr "" "(" (toString $real) ")")}}
{{$comm:=""}}

{{$seuil := (toInt 8)}}
{{$idb := (toInt 0)}}

{{$arg := 0}}
{{$argm := 0}}
{{$c := ""}}

{{$manuel := reFind `s[1-9]` .Message.Content}}
{{$force:=reFindAllSubmatches `(force|Force)` .Message.Content}}
{{$agi := reFindAllSubmatches `(agilité|Agilité|agi|Agi)` .Message.Content}}
{{$endu := reFindAllSubmatches `(Endurance|endurance|endu|Endu)` .Message.Content}}
{{$preci := reFindAllSubmatches `(Précision|précision|préci|Préci)` .Message.Content}}
{{$intell := reFindAllSubmatches `(Intelligence|intelligence|intell|Intell|intel|Intel)` .Message.Content}}
{{$karma := reFindAllSubmatches `(Karma|karma)` .Message.Content}}

{{if $manuel}}
	{{$seuil = (toInt (joinStr "" (split $manuel "s")))}}
	{{$idb = (toInt 0)}}
{{else if $force}}
	{{$seuil = (toInt (dbGet .User.ID "force").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_force").Value)}}
{{else if $agi}}
	{{$seuil = (toInt (dbGet .User.ID "agi").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_agi").Value)}}
{{else if $endu}}
	{{$seuil = (toInt (dbGet .User.ID "endurance").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_endu").Value)}}
{{else if $preci}}
	{{$seuil = (toInt (dbGet .User.ID "preci").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_preci").Value)}}
{{else if $intell}}
	{{$seuil = (toInt (dbGet .User.ID "intelligence").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_intel").Value)}}
{{else if $karma}}
	{{$seuil = (toInt (dbGet .User.ID "karma").Value)}}
	{{$idb = (toInt (dbGet .User.ID "i_karma").Value)}}
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


{{if .CmdArgs}}
	{{if $manuel}}
		{{$c = split (joinStr " " .CmdArgs) $manuel}}
	{{else}}
		{{$c = (joinStr " " .CmdArgs)}}
	{{end}}

	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
	  {{if ne (toInt (index .CmdArgs 0)) (toInt 0)}}
	    {{$i := (toInt  (index .CmdArgs 0))}}
			{{$x := $i}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{$arg = 1}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1) }}
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
			{{else}}
				{{$d = add $i $d}}
				{{$d = sub $d $idb}}
			{{end}}

			{{if gt $d (toInt 10)}}
				{{$d = toInt 10}}
			{{else if lt $d (toInt 0)}}
				{{$d = (toInt 0)}}
			{{end}}

			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " (toString $x) " : " $arg1 " | " $imp " : " $idb " | Seuil : " $seuil )}}
  	{{else}}
			{{$d = sub $d $idb}}
			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
		{{end}}
	{{else}}
		{{if ne (toInt (index .CmdArgs 0)) (toInt 0) }}
			{{$i = (toInt  (index .CmdArgs 0))}}
			{{$x := $i}}
			{{if gt (toInt (index .CmdArgs 0)) (toInt 0)}}
				{{$arg = 1}}
				{{if eq (toInt (index .CmdArgs 0)) (toInt 1)}}
					{{$arg1 = "pénalité"}}
				{{else}}
					{{$arg1 = "pénalités"}}
				{{end}}
			{{else}}
				{{$arg = 0}}
				{{$i = mult $i (toInt -1)}}
				{{$x = mult $x (toInt -1)}}
				{{$arg1 = "Bonus"}}
			{{end}}

			{{if ne (toInt (index .CmdArgs 1)) (toInt 0)}}
				{{$m := (toInt (index .CmdArgs 1))}}
				{{$y := $m}}
				{{if gt $m (toInt 0)}}
					{{$argm = 1}}
					{{if eq $m (toInt 1) }}
						{{$arg2 = "Pénalité"}}
					{{else}}
						{{$arg2 = "Pénalités"}}
					{{end}}
				{{else}}
					{{$argm = 0}}
					{{$m = mult $m (toInt -1)}}
					{{$y = mult $y (toInt -1)}}
					{{$arg2 = "Bonus"}}
				{{end}}

				{{if eq $arg 0}}
					{{$i = add $i $idb}}
					{{$d = sub $d $i}}
					{{$mimp = $i}}
					{{if eq $argm 0}}
						{{$d = sub $d $m}}
					{{else}}
						{{$d = add $d $m}}
					{{end}}
				{{else}}
					{{$d = add $d $i}}
					{{if eq $argm 0}}
						{{$m = add $m $idb}}
						{{$d = sub $d $m }}
						{{$mimp = $m}}
					{{else}}
						{{$d = add $d $m}}
						{{$d = sub $d $idb}}
						{{$mimp = $d}}
					{{end}}
				{{end}}


				{{if gt $d (toInt 10)}}
					{{$d = toInt 10}}
				{{else if lt $d (toInt 0)}}
					{{$d = (toInt 0)}}
				{{end}}


				{{$res = (joinStr "" "Dé : " $d " " $v " | " $arg1 " : " $x " | " $arg2 " : " $y " | " $imp " : " $idb " | Seuil : " $seuil )}}
			{{else}}

						{{if eq $arg 0}}
							{{$i = add $i $idb}}
							{{$d = sub $d $i}}
							{{$mimp = $i}}
						{{else}}
							{{$d = add $i $d}}
							{{$d = sub $d $idb}}
							{{$mimp = $idb}}
						{{end}}

						{{if gt $d (toInt 10)}}
							{{$d = toInt 10}}
						{{else if lt $d (toInt 0)}}
							{{$d = (toInt 0)}}
						{{end}}

				{{$res = (joinStr "" "Dé : " (toString $d) " " $v " |  " $arg1 " : " (toString $x) " | " $imp " : " $idb " | Seuil : " $seuil)}}
			{{end}}
		{{else}}
			{{$d = sub $d $idb}}
			{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
		{{end}}
	{{end}}

	{{if lt (toFloat (len .CmdArgs)) (toFloat 2)}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{$comm = ""}}
		{{else}}
			{{$comm = (joinStr " " $c ": ")}}
		{{end}}
	{{else}}
		{{if ne (toInt 0) (toInt (index .CmdArgs 0)) }}
			{{if ne (toFloat (index .CmdArgs 1)) (toFloat 0)}}
				{{if eq (toFloat 2) (toFloat (len .CmdArgs))}}
					{{$comm = ""}}
				{{else}}
					{{$comm = (joinStr " " (slice $c 2) ": ") }}
				{{end}}
			{{else}}
				{{$comm = (joinStr " " (slice $c 1) ": ") }}
			{{end}}
		{{else}}
			{{$comm = (joinStr " " $c ": ") }}
		{{end}}
	{{end}}
{{else}}
	{{$res = (joinStr "" "Dé : " (toString $d) " " $v " | " $imp " : " $idb " | Seuil : " $seuil )}}
{{end}}


{{$urc := cembed
	"description" (joinStr "" "**" $user "** ▬ " $comm "**Ultra critique** : Votre cible a un empoisonnement de 10 PV sur 3 tours (-30 PV en tout).\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x4D399E }}

{{$rc := cembed
	"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite critique** : Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x4D399E }}

{{$r := ""}}

{{if eq $seuil (toInt 2)}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "**Réussite** : Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x4D399E }}

{{else if or (eq $seuil (toInt 3)) (eq $seuil (toInt 4))}}
	{{$r = cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible a un empoisonnement de 8 PV sur 3 tours (-24 PV en tout).\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x4D399E }}

{{else if or (eq $seuil (toInt 5)) (eq $seuil (toInt 6))}}
	{{$r:= cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible a un empoisonnement de 4 PV sur 3 tours (-12 PV en tout).\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x4D399E }}

{{else if or (eq $seuil (toInt 7)) (eq $seuil (toInt 8))}}
	{{$embed := cembed
		"description" (joinStr "" "**" $user "** ▬ " $comm "*Réussite* : Votre cible a un empoisonnement de 2 PV sur 3 tours (-6 PV en tout).\n"
		"<:next:723131844643651655> *[" $res "]*" )
		"color" 0x4D399E }}
{{end}}

{{$echec := cembed
	"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec de l'empoisonnement**...\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x4D399E }}

{{$ec:= cembed
	"description" (joinStr "" "**" $user "** ▬ " $comm "**Echec critique de l'empoisonnement :** Votre cible gagne une régénération de 3 PV sur 3 tours (+9 PV en tout).\n"
	"<:next:723131844643651655> *[" $res "]*" )
	"color" 0x4D399E }}

{{if eq $d (toInt 0)}}
	{{sendMessage nil $urc}}
{{else if eq $d (toInt 1)}}
	{{sendMessage nil $rc}}
{{else if le $d $seuil}}
	{{if ge $mimp (toInt 1)}}
		{{if eq $d $seuil}}
			{{sendMessage nil $echec}}
		{{else}}
			{{sendMessage nil $r}}
		{{end}}
	{{else}}
		{{sendMessage nil $r}}
	{{end}}
{{else if or (gt $d $seuil) (lt $d (toInt 9)) }}
	{{sendMessage nil $echec}}
{{else if eq $d (toInt 10)}}
	{{sendMessage nil $ec}}
{{else if eq $d (toInt 9)}}
	{{sendMessage nil $echec}}
{{else}}
{{end}}

{{deleteTrigger 1}}
