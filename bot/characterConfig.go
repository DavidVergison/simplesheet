package main

/*
Maison
Année scolaire
-- Caractéristiques --
💭 ESPRIT 💭
💪 CORPS 💪
❤  COEUR  ❤
✨ MAGIE ✨
💭Bluff
💭Farce
💭Tactique
❤ Rumeur
💪Bagarre
💪Endurance
💭Perception
💪Précision
❤ Décorum
💪Discrétion
❤ Persuasion
❤ Relationnel
-- Bonus --
Etude des Moldus : Nul
Hist. de la Magie : Nul
Botanique : Nul
SAFARI : Nul
Potions : Nul
Astronomie : Bon
Enchantement : Bon
Divination : Bon
MAD : Bon
Action de Maison
-- Traits --
Points d'Eveil dispos =&gt;
Noble
Riche
3ème Oeil
Nul (Etude des Moldus)
Nul (Histoire de la Magie)
Nul (Botanique)
Nul (SAFARI)
Nul (Potions)
Bon (Astronomie)
Bon (Enchantement)
Bon (Divination)
Bon (MAD)
M.A.D.
-- 👨‍🎓 Cours 👩‍🎓 --
Enchantement
Maléfices
Potions
*/

func (char Character) getParsingConfig() map[string]Parsing {
	return map[string]Parsing{
		"ESPR":{
			label:         "ESPRIT",
			regexp:        getRegexp("ESPRIT"),
			statisticName: "ESPR",
		},
		"PERC":{
			label:         "Perception",
			regexp:        getRegexp("Perception$"),
			statisticName: "PERC",
		},
		"BLUF":{
			label:         "Bluff",
			regexp:        getRegexp("Bluff$"),
			statisticName: "BLUF",
		},
		"FARC": {
			label:         "Farce",
			regexp:        getRegexp("Farce$"),
			statisticName: "FARC",
		},
		"TACT": {
			label:         "Tactique",
			regexp:        getRegexp("Tactique$"),
			statisticName: "TACT",
		},
		"CORP": {
			label:         "CORPS",
			regexp:        getRegexp("CORPS"),
			statisticName: "CORP",
		},
		"PREC": {
			label:         "Précision",
			regexp:        getRegexp("Précision$"),
			statisticName: "PREC",
		},
		"DISC": {
			label:         "Discrétion",
			regexp:        getRegexp("Discrétion$"),
			statisticName: "DISC",
		},
		"ENDU": {
			label:         "Endurance",
			regexp:        getRegexp("Endurance$"),
			statisticName: "ENDU",
		},
		"BAGA": {
			label:         "Bagarre",
			regexp:        getRegexp("Bagarre$"),
			statisticName: "BAGA",
		},
		"COEU": {
			label:         "COEUR",
			regexp:        getRegexp("COEUR"),
			statisticName: "COEU",
		},
		"RUME":{
			label:         "Rumeur",
			regexp:        getRegexp("Rumeur$"),
			statisticName: "RUME",
		},
		"DECO":{
			label:         "Décorum",
			regexp:        getRegexp("Décorum$"),
			statisticName: "DECO",
		},
		"PERS":{
			label:         "Persuasion",
			regexp:        getRegexp("Persuasion$"),
			statisticName: "PERS",
		},
		"RELA":{
			label:         "Relationnel",
			regexp:        getRegexp("Relationnel$"),
			statisticName: "RELA",
		},
		"EDM":{
			label:         "Etude des Moldus",
			regexp:        getRegexp("^Etude des Moldus :"),
			statisticName: "EDM",
		},
		"HIS":{
			label:         "Hist. de la Magie",
			regexp:        getRegexp("^Hist. de la Magie :"),
			statisticName: "HIS",
		},
		"BOT":{
			label:         "Botanique",
			regexp:        getRegexp("^Botanique :"),
			statisticName: "BOT",
		},
		"SAF":{
			label:         "SAFARI",
			regexp:        getRegexp("^SAFARI :"),
			statisticName: "SAF",
		},
		"POT":{
			label:         "Potions",
			regexp:        getRegexp("^Potions :"),
			statisticName: "POT",
		},
		"AST":{
			label:         "Astronomie",
			regexp:        getRegexp("^Astronomie :"),
			statisticName: "AST",
		},
		"ENC":{
			label:         "Enchantement",
			regexp:        getRegexp("^Enchantement :"),
			statisticName: "ENC",
		},
		"DIV":{
			label:         "Divination",
			regexp:        getRegexp("^Divination :"),
			statisticName: "DIV",
		},
		"MAD":{
			label:         "MAD",
			regexp:        getRegexp("^MAD :"),
			statisticName: "MAD",
		},
		"MET":{
			label:         "Métamorphose",
			regexp:        getRegexp("^Métamorphose :"),
			statisticName: "MET",
		},
		"MAGI":{
			label:         "MAGIE",
			regexp:        getRegexp("MAGIE"),
			statisticName: "MAGI",
		},
		"VOL":{
			label:         "Vol",
			regexp:        getRegexp("Vol"),
			statisticName: "VOL",
		},
	}

}
