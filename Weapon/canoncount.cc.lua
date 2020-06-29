
{{/* Each time the bot sees the trigger, it will count until it reaches the value set in the "if lt".
It will also count the number of balls used, and will return this message to tell the user that it has no more balls.

If you change the value of the if, you must change the value in the "$x := sub".  */}}


{{if not (dbGet .User.ID "canon")}}
  {{dbSet .User.ID "canon" 0}}
  {{$incr := dbIncr .User.ID "canon" 1}}
  {{$y := (dbGet .User.ID "canon").Value}}
  {{$x := sub 20 $y}}
  {{if lt $y (toFloat 20)}}
    {{ $embed := cembed
      "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges de canon !")}}
      {{ $id := sendMessageRetID nil $embed }}
  		{{deleteMessage nil $id 30}}
  {{else}}
    {{ $embed := cembed
      "description" (joinStr "" .User.Mention ", votre canon est vide.")}}
      {{ $id := sendMessageRetID nil $embed }}
  		{{deleteMessage nil $id 30}}
  {{end}}
{{else}}
  {{$incr := dbIncr .User.ID "canon" 1}}
  {{$y := (dbGet .User.ID "canon").Value}}
  {{$x := sub 20 $y}}
  {{if lt $y (toFloat 20)}}
    {{ $embed := cembed
    "description" (joinStr "" .User.Mention ", il vous reste " (toString (toInt $x)) " charges de canon !")}}
    {{ $id := sendMessageRetID nil $embed }}
    {{deleteMessage nil $id 30}}
  {{else}}
    {{ $embed := cembed
    "description" (joinStr "" .User.Mention ", votre canon est vide.")}}
    {{ $id := sendMessageRetID nil $embed }}
		{{deleteMessage nil $id 30}}
  {{end}}
{{end}}