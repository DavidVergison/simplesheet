package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DavidVergison/simplesheet/bot/kankalib"
	"github.com/MakeNowJust/heredoc"
)

type Character struct {
	Attributes map[string]Statistic
	TestRolls map[string]TestRoll
}

type TestRoll struct {
	Label string
	AttributesNames []string
}

func InitCharacter(entityId string, kankaClient kankalib.KankaClient) (Character, error) {
	attributes, err := kankaClient.CharacterAttributes(entityId)
	if err != nil {
		return Character{}, err
	}

	character := Character{}
	character.ParseAttributes(attributes)

	character.TestRolls =  map[string]TestRoll{
		"Esprit": {
			Label: "Esprit",
			AttributesNames: []string{"ESPR"},
		},
		"Corps": {
			Label: "Corps",
			AttributesNames: []string{"CORP"},
		},
		"Coeur": {
			Label: "Coeur",
			AttributesNames: []string{"COEU"},
		},
		"Perception": {
			Label: "Perception",
			AttributesNames: []string{"PERC","ESPR"},
		},
		"Bluff": {
			Label: "Bluff",
			AttributesNames: []string{"BLUF","ESPR"},
		},
		"Farce": {
			Label: "Farce",
			AttributesNames: []string{"FARC","ESPR"},
		},
		"Tactique": {
			Label: "Tactique",
			AttributesNames: []string{"TACT","ESPR"},
		},
		"Précision": {
			Label: "Précision",
			AttributesNames: []string{"PREC","CORP"},
		},
		"Discrétion": {
			Label: "Discrétion",
			AttributesNames: []string{"DISC","CORP"},
		},
		"Endurance": {
			Label: "Endurance",
			AttributesNames: []string{"ENDU","CORP"},
		},
		"Bagarre": {
			Label: "Bagarre",
			AttributesNames: []string{"BAGA","CORP"},
		},
		"Rumeur": {
			Label: "Rumeur",
			AttributesNames: []string{"RUME","COEU"},
		},
		"Décorum": {
			Label: "Décorum",
			AttributesNames: []string{"DECO","COEU"},
		},
		"Persuasion": {
			Label: "Persuasion",
			AttributesNames: []string{"PERS","COEU"},
		},
		"Relationnel": {
			Label: "Relationnel",
			AttributesNames: []string{"RELA","COEU"},
		},	
		
	}

	return character, nil
}

type Statistic struct {
	label         string
	statisticName string
	bonus         []string
	value         int
}

type Parsing struct {
	label         string
	regexp        *regexp.Regexp
	statisticName string
	bonus         []string
}

func getRegexp(expr string) *regexp.Regexp {
	reg, _ := regexp.Compile(expr)
	return reg
}

func (char *Character) ParseAttributes(attributes kankalib.Attributes) {
	config := char.getParsingConfig()
	char.Attributes = map[string]Statistic{}

	for _, d := range attributes.Data {
		fmt.Println(d.Name)
	}

	for _, c := range config {
		attr, err := char.filterAttribute(attributes.Data, c)
		value := 0
		if err == nil {
			stringValue, ok := attr.Value.(string)
			if ok {
				value, _ = strconv.Atoi(stringValue)
			}
		}
		char.Attributes[c.statisticName] = Statistic{
			label:         c.label,
			statisticName: c.statisticName,
			bonus:         c.bonus,
			value:         value,
		}

	}
}

func (char Character) filterAttribute(data []kankalib.Data, pattern Parsing) (kankalib.Data, error) {
	for _, attr := range data {
		if pattern.regexp.MatchString(attr.Name) {
			return attr, nil
		}
	}

	return kankalib.Data{}, fmt.Errorf("Not Found : %s", pattern.label)
}

func (char Character) AttributeByLabel(label string) Statistic {
	for _, attr := range char.Attributes {
		if attr.label == label {
			return attr
		}
	}

	return Statistic{}
}

/*
Maléfices				1d20 + 3 (💭 + ✨ + Farce)
M.A.D.					1d20 + 5 (💭 + ✨ + MAD)
Enchantement			1d20 + 5 (💭 + ✨ + Enchantement)
Potions					1d20 + 2 (💭 + Précision + Potions)
Botanique				1d20 + x (💪+ [situationel] + Botanique)
Histoire de la Magie	1d20 + 0 (💭 + Endurance + Histoire)
Métamorphose			1d20 + x (💭 + ✨ + Métamorphose)
Vol						1d20 + x (💪 + Energie actuelle + Vol)
S.A.F.A.R.I.			1d20 + x (❤ + [situationel] + SAFARI)
Astronomie				1d20 + x (💭 + Perception + Astronomie)
Cartographier une constellation		1d20 + x (💪+ Précision + Astronomie)
*/
func (char Character) ToString() string {
	doc := heredoc.Doc(`
		ESPRIT     : @@ESPR@@    CORPS      : @@CORP@@    COEUR       : @@COEU@@
		Perception : @@PERC@@    Précision  : @@PREC@@    Rumeur      : @@RUME@@
		Bluff      : @@BLUF@@    Discrétion : @@DISC@@    Décorum     : @@DECO@@
		Farce      : @@FARC@@    Endurance  : @@ENDU@@    Persuasion  : @@PERS@@
		Tactique   : @@TACT@@    Bagarre    : @@BAGA@@    Relationnel : @@RELA@@

		Maléfices               1d20 + ESPRIT(@@ESPR@@) + MAGIE(@@MAGI@@)      + Farce(@@FARC@@)
		M.A.D.                  1d20 + ESPRIT(@@ESPR@@) + MAGIE(@@MAGI@@)      + MAD(@@MAD@@)
		Enchantement            1d20 + ESPRIT(@@ESPR@@) + MAGIE(@@MAGI@@)      + Enchantement(@@ENC@@)
		Potions                 1d20 + ESPRIT(@@ESPR@@) + Précision(@@PREC@@)  + Potions(@@POT@@)
		Botanique               1d20 + CORPS(@@CORP@@)  + Botanique(@@BOT@@)   + situationnel
		Histoire de la Magie    1d20 + ESPRIT(@@ESPR@@) + Endurance(@@ENDU@@)  + Histoire de la Magie(@@HIS@@)
		Métamorphose            1d20 + ESPRIT(@@ESPR@@) + MAGIE(@@MAGI@@)      + Métamorphose(@@MET@@)
		Vol                     1d20 + CORPS(@@CORP@@)  + Vol(@@VOL@@)         + Energie actuelle
		S.A.F.A.R.I.            1d20 + COEUR(@@COEU@@)  + SAFARI(@@SAF@@)      + situationel
		Astronomie              1d20 + ESPRIT(@@ESPR@@) + Perception(@@PERC@@) + Astronomie(@@AST@@)
		Cart. une constellation	1d20 + CORPS(@@CORP@@)  + Précision(@@PREC@@)  + Astronomie(@@AST@@)
	`)
	for k, v := range char.Attributes {
		doc = strings.Replace(doc, fmt.Sprintf("@@%s@@", k), strconv.Itoa(v.value), -1)
	}
	return fmt.Sprintf("```\n%s\n```", doc)
}

func (char Character) Roll(statistic string) string {

	roll := char.TestRolls[statistic]
	score := 0
	txt := ""
	for _, attributeName := range roll.AttributesNames {
		attribute:=char.Attributes[attributeName]
		score += attribute.value
		txt += fmt.Sprintf("%s(%d),", attribute.label, attribute.value)
	}
	txt = txt[:len(txt)-1]

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	dice := r.Intn(20) + 1
	total := dice + score

	return fmt.Sprintf("TOTAL : %d  => 1d20(%d) + %d [%s]", total, dice, score, txt)
}
