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
}

func InitCharacter(entityId string, kankaClient kankalib.KankaClient) (Character, error) {
	attributes, err := kankaClient.CharacterAttributes(entityId)
	if err != nil {
		return Character{}, err
	}

	character := Character{}
	character.ParseAttributes(attributes)
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

	for _, c := range config {
		stringValue, ok := char.filterAttribute(attributes.Data, c).Value.(string)
		if ok {
			value, _ := strconv.Atoi(stringValue)			
			char.Attributes[c.statisticName] = Statistic{
				label:         c.label,
				statisticName: c.statisticName,
				bonus:         c.bonus,
				value:         value,
			}
		}
	}
}

func (char Character) filterAttribute(data []kankalib.Data, pattern Parsing) kankalib.Data {
	for _, attr := range data {
		if pattern.regexp.MatchString(attr.Name) {
			return attr
		}
	}

	return kankalib.Data{}
}

func (char Character) AttributeByLabel(label string) Statistic{
	for _, attr := range char.Attributes {
		if attr.label == label {
			return attr
		}
	}

	return Statistic{}
}

func (char Character) ToString() string {
	doc := heredoc.Doc(`
		ESPRIT     : @@ESPR@@    CORPS      : @@CORP@@    COEUR       : @@COEU@@
		Perception : @@PERC@@    Précision  : @@PREC@@    Rumeur      : @@RUME@@
		Bluff      : @@BLUF@@    Discrétion : @@DISC@@    Décorum     : @@DECO@@
		Farce      : @@FARC@@    Endurance  : @@ENDU@@    Persuasion  : @@PERS@@
		Tactique   : @@TACT@@    Bagarre    : @@BAGA@@    Relationnel : @@RELA@@
	`)
	for k, v := range char.Attributes {
		doc = strings.Replace(doc, fmt.Sprintf("@@%s@@", k), strconv.Itoa(v.value), -1)
	}
	return fmt.Sprintf("```\n%s\n```", doc)
}

func (char Character) Roll(statistic string) string {

	stat := char.AttributeByLabel(statistic)
	txt := fmt.Sprintf("%s(%d)", stat.label, stat.value)
	bonus := 0
	for _, b := range stat.bonus {
		bonusStat := char.Attributes[b]
		bonus = bonus + bonusStat.value
		txt += fmt.Sprintf(" + %s(%d)", bonusStat.label, bonusStat.value)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	dice := r.Intn(20) + 1
	total := dice + bonus + stat.value

	return fmt.Sprintf("TOTAL : %d  => 1d20(%d) + %s", total, dice, txt)
}
