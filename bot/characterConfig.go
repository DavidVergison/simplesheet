package main

/*
💭ESPRIT
💪CORPS
❤ COEUR
💭ESPRIT - COEUR ❤
Bluff
Farce
Tactique
Rumeur
💭ESPRIT - CORPS 💪
Bagarre
Endurance
Perception
Précision
❤ COEUR - CORPS💪
Décorum
Discrétion
Persuasion
Relationnel
-- Traits --
Noble
Riche
3ème Oeil
Nul (Etude des Moldus)
-- Bonus --
Etude des Moldus : Nul
✨MAGIE
Points d'Eveil dispos =&gt;
Nul (Histoire de la Magie)
Nul (Botanique)
Nul (SAFARI)
Nul (Potions)
Bon (Astronomie)
Bon (Enchantement)
Hist. de la Magie : Nul
Botanique : Nul
SAFARI : Nul
Potions : Nul
Astronomie : Bon
Enchantement : Bon
Divination : Bon
MAD : Bon
Bon (Divination)
Bon (MAD)
Maison
Année scolaire
Action de Maison
*/

func (char Character) getParsingConfig() []Parsing {
	return []Parsing{
		{
			label:         "ESPRIT",
			regexp:        getRegexp("ESPRIT$"),
			statisticName: "ESPR",
			bonus:         []string{},
		},
		{
			label:         "Perception",
			regexp:        getRegexp("^Perception$"),
			statisticName: "PERC",
			bonus:         []string{"ESPR"},
		},
		{
			label:         "Bluff",
			regexp:        getRegexp("^Bluff$"),
			statisticName: "BLUF",
			bonus:         []string{"ESPR"},
		},
		{
			label:         "Farce",
			regexp:        getRegexp("^Farce$"),
			statisticName: "FARC",
			bonus:         []string{"ESPR"},
		},
		{
			label:         "Tactique",
			regexp:        getRegexp("^Tactique$"),
			statisticName: "TACT",
			bonus:         []string{"ESPR"},
		},
		{
			label:         "CORPS",
			regexp:        getRegexp("CORPS$"),
			statisticName: "CORP",
			bonus:         []string{},
		},
		{
			label:         "Précision",
			regexp:        getRegexp("^Précision$"),
			statisticName: "PREC",
			bonus:         []string{"CORP"},
		},
		{
			label:         "Discrétion",
			regexp:        getRegexp("^Discrétion$"),
			statisticName: "DISC",
			bonus:         []string{"CORP"},
		},
		{
			label:         "Endurance",
			regexp:        getRegexp("^Endurance$"),
			statisticName: "ENDU",
			bonus:         []string{"CORP"},
		},
		{
			label:         "Bagarre",
			regexp:        getRegexp("^Bagarre$"),
			statisticName: "BAGA",
			bonus:         []string{"CORP"},
		},
		{
			label:         "COEUR",
			regexp:        getRegexp("COEUR$"),
			statisticName: "COEU",
			bonus:         []string{},
		},
		{
			label:         "Rumeur",
			regexp:        getRegexp("Rumeur$"),
			statisticName: "RUME",
			bonus:         []string{"COEU"},
		},
		{
			label:         "Décorum",
			regexp:        getRegexp("Décorum$"),
			statisticName: "DECO",
			bonus:         []string{"COEU"},
		},
		{
			label:         "Persuasion",
			regexp:        getRegexp("Persuasion$"),
			statisticName: "PERS",
			bonus:         []string{"COEU"},
		},
		{
			label:         "Relationnel",
			regexp:        getRegexp("Relationnel$"),
			statisticName: "RELA",
			bonus:         []string{"COEU"},
		},
	}

}
