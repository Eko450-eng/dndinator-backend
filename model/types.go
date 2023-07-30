package model

import (
	"fmt"
	"io"
	"strconv"
)

type Abillitie struct {
	ID   *int    `json:"id,omitempty" db:"id"`
	Name *string `json:"name,omitempty" db:"name"`
}

type User struct {
	ID       *int    `json:"id,omitempty" db:"id"`
	UserName *string `json:"username,omitempty" db:"username"`
	Email    *string `json:"email,omitempty" db:"email"`
	Password *string `json:"password,omitempty" db:"password"`
	// Characters *[]Character `json:"characterid,omitempty" db:"characterid"`
}

type Attack struct {
	ID          *int       `json:"id,omitempty" db:"id"`
	Name        *string    `json:"name,omitempty" db:"name"`
	Atk         *string    `json:"atk,omitempty" db:"atk"`
	Range       *string    `json:"range,omitempty" db:"range"`
	Type        *string    `json:"type,omitempty" db:"type"`
	Characterid *Character `json:"characterid,omitempty" db:"characterid"`
}

type Character struct {
	ID         *int       `json:"id,omitempty" db:"id"`
	Name       string     `json:"name" db:"name"`
	Lvl        *int       `json:"lvl,omitempty" db:"lvl"`
	Race       *string    `json:"race,omitempty" db:"race"`
	Alignment  *string    `json:"alignment,omitempty" db:"alignment"`
	Background *string    `json:"background,omitempty" db:"background"`
	Hp         *int       `json:"hp,omitempty" db:"hp"`
	Maxhp      *int       `json:"maxhp,omitempty" db:"maxhp"`
	Ac         *int       `json:"ac,omitempty" db:"ac"`
	Init       *int       `json:"init,omitempty" db:"init"`
	Party      *int       `json:"party,omitempty" db:"party"`
	Playerid   *string    `json:"playerid,omitempty" db:"playerid"`
	Abillitie  *Abillitie `json:"abillitie,omitempty" db:"abillitie"`
	Language   *Language  `json:"language,omitempty" db:"language"`
	Class      *Classes   `json:"class,omitempty" db:"class"`
}

type Class struct {
	ID    *int    `json:"id,omitempty" db:"id"`
	Index *string `json:"index,omitempty" db:"index"`
	Name  *string `json:"name,omitempty" db:"name"`
	URL   *string `json:"url,omitempty" db:"url"`
}

type Coins struct {
	ID          *int       `json:"id,omitempty" db:"id"`
	Pp          *int       `json:"pp,omitempty" db:"pp"`
	Gp          *int       `json:"gp,omitempty" db:"gp"`
	Ep          *int       `json:"ep,omitempty" db:"ep"`
	Sp          *int       `json:"sp,omitempty" db:"sp"`
	Cp          *int       `json:"cp,omitempty" db:"cp"`
	Characterid *Character `json:"characterid,omitempty" db:"characterid"`
}

type DetStat struct {
	ID          *int       `json:"id,omitempty" db:"id"`
	Name        *DetStats  `json:"name,omitempty" db:"name"`
	Value       *int       `json:"value,omitempty" db:"value"`
	Proficiency *bool      `json:"proficiency,omitempty" db:"proficiency"`
	Characterid *Character `json:"characterid,omitempty" db:"characterid"`
}

type Equipment struct {
	ID          *int       `json:"id,omitempty" db:"id"`
	Name        *string    `json:"name,omitempty" db:"name"`
	Amount      *int       `json:"amount,omitempty" db:"amount"`
	Characterid *Character `json:"characterid,omitempty" db:"characterid"`
}

type Language struct {
	ID   *int    `json:"id,omitempty" db:"id"`
	Name *string `json:"name,omitempty" db:"name"`
}

type Spell struct {
	ID          *int      `json:"id,omitempty" db:"id"`
	Name        string    `json:"name" db:"name"`
	CastingTime string    `json:"casting_time" db:"casting_time"`
	Range       string    `json:"range" db:"range"`
	Duration    string    `json:"duration" db:"duration"`
	Components  []*string `json:"components,omitempty" db:"components"`
	Desc        *string   `json:"desc,omitempty" db:"desc"`
	Higherdesc  *string   `json:"higherdesc,omitempty" db:"higherdesc"`
	Level       *int      `json:"level,omitempty" db:"level"`
	Classes     []*string `json:"classes,omitempty" db:"classes"`
}

type Spellbook struct {
	ID          *int `json:"id,omitempty" db:"id"`
	Tier        *int `json:"tier,omitempty" db:"tier"`
	Slots       *int `json:"slots,omitempty" db:"slots"`
	Characterid *int `json:"characterid,omitempty" db:"characterid"`
}

type Spellbook_spell struct {
	ID    *int `json:"id,omitempty" db:"id"`
	Spell *int `json:"spell,omitempty" db:"spell"`
	Book  *int `json:"book,omitempty" db:"book"`
}

type Spelllist struct {
	ID    *int    `json:"id,omitempty" db:"id"`
	Index *string `json:"index,omitempty" db:"index"`
	Name  *string `json:"name,omitempty" db:"name"`
	URL   *string `json:"url,omitempty" db:"url"`
}

type Stat struct {
	ID          *int       `json:"id,omitempty" db:"id"`
	Name        *Stats     `json:"name,omitempty" db:"name"`
	Value       *int       `json:"value,omitempty" db:"value"`
	Savingvalue *int       `json:"savingvalue,omitempty" db:"savingvalue"`
	Modifier    *int       `json:"modifier,omitempty" db:"modifier"`
	Shorthand   *string    `json:"shorthand,omitempty" db:"shorthand"`
	Characterid *Character `json:"characterid,omitempty" db:"characterid"`
}

type InputAddSpell struct {
	SpellID     *int `json:"spellId,omitempty" db:"spell_id"`
	Characterid *int `json:"characterid,omitempty" db:"characterid"`
}

type InputAttack struct {
	Name        *string `json:"name,omitempty" db:"name"`
	Atk         *string `json:"atk,omitempty" db:"atk"`
	Range       *string `json:"range,omitempty" db:"range"`
	Type        *string `json:"type,omitempty" db:"type"`
	Characterid *int    `json:"characterid,omitempty" db:"characterid"`
}

type InputCharacter struct {
	Name       string  `json:"name" db:"name"`
	Lvl        *int    `json:"lvl,omitempty" db:"lvl"`
	Race       *string `json:"race,omitempty" db:"race"`
	Alignment  *string `json:"alignment,omitempty" db:"alignment"`
	Background *string `json:"background,omitempty" db:"background"`
	Hp         *int    `json:"hp,omitempty" db:"hp"`
	Maxhp      *int    `json:"maxhp,omitempty" db:"maxhp"`
	Ac         *int    `json:"ac,omitempty" db:"ac"`
	Init       *int    `json:"init,omitempty" db:"init"`
	Party      *int    `json:"party,omitempty" db:"party"`
}

type InputDetStat struct {
	Name        *DetStats `json:"name,omitempty" db:"name"`
	Value       *int      `json:"value,omitempty" db:"value"`
	Proficiency *bool     `json:"proficiency,omitempty" db:"proficiency"`
	Characterid *int      `json:"characterid,omitempty" db:"characterid"`
}

type InputSpell struct {
	Name        string    `json:"name" db:"name"`
	CastingTime string    `json:"casting_time" db:"casting_time"`
	Range       string    `json:"range" db:"range"`
	Duration    string    `json:"duration" db:"duration"`
	Components  []*string `json:"components,omitempty" db:"components"`
	Desc        *string   `json:"desc,omitempty" db:"desc"`
	Higherdesc  *string   `json:"higherdesc,omitempty" db:"higherdesc"`
	Level       *int      `json:"level,omitempty" db:"level"`
	Classes     []*string `json:"classes,omitempty" db:"classes"`
}

type InputSpelllist struct {
	Index *string `json:"index,omitempty" db:"index"`
	Name  *string `json:"name,omitempty" db:"name"`
	URL   *string `json:"url,omitempty" db:"url"`
}

type InputStat struct {
	Name        *Stats  `json:"name,omitempty" db:"name"`
	Value       *int    `json:"value,omitempty" db:"value"`
	Savingvalue *int    `json:"savingvalue,omitempty" db:"savingvalue"`
	Modifier    *int    `json:"modifier,omitempty" db:"modifier"`
	Shorthand   *string `json:"shorthand,omitempty" db:"shorthand"`
	Characterid *int    `json:"characterid,omitempty" db:"characterid"`
}

type Classes string

const (
	ClassesArtificer Classes = "Artificer"
	ClassesBarbarian Classes = "Barbarian"
	ClassesBard      Classes = "Bard"
	ClassesBlood     Classes = "Blood"
	ClassesHunter    Classes = "Hunter"
	ClassesCleric    Classes = "Cleric"
	ClassesDruid     Classes = "Druid"
	ClassesFighter   Classes = "Fighter"
	ClassesMonk      Classes = "Monk"
	ClassesPaladin   Classes = "Paladin"
	ClassesRanger    Classes = "Ranger"
	ClassesRogue     Classes = "Rogue"
	ClassesSorcerer  Classes = "Sorcerer"
	ClassesWarlock   Classes = "Warlock"
	ClassesWizard    Classes = "Wizard"
)

var AllClasses = []Classes{
	ClassesArtificer,
	ClassesBarbarian,
	ClassesBard,
	ClassesBlood,
	ClassesHunter,
	ClassesCleric,
	ClassesDruid,
	ClassesFighter,
	ClassesMonk,
	ClassesPaladin,
	ClassesRanger,
	ClassesRogue,
	ClassesSorcerer,
	ClassesWarlock,
	ClassesWizard,
}

func (e Classes) IsValid() bool {
	switch e {
	case ClassesArtificer, ClassesBarbarian, ClassesBard, ClassesBlood, ClassesHunter, ClassesCleric, ClassesDruid, ClassesFighter, ClassesMonk, ClassesPaladin, ClassesRanger, ClassesRogue, ClassesSorcerer, ClassesWarlock, ClassesWizard:
		return true
	}
	return false
}

func (e Classes) String() string {
	return string(e)
}

func (e *Classes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Classes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Classes", str)
	}
	return nil
}

func (e Classes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DetStats string

const (
	DetStatsAcrobatics    DetStats = "Acrobatics"
	DetStatsAnimal        DetStats = "Animal"
	DetStatsHandling      DetStats = "Handling"
	DetStatsArcana        DetStats = "Arcana"
	DetStatsAthletics     DetStats = "Athletics"
	DetStatsDeception     DetStats = "Deception"
	DetStatsHistory       DetStats = "History"
	DetStatsInsight       DetStats = "Insight"
	DetStatsIntimidation  DetStats = "Intimidation"
	DetStatsInvestigation DetStats = "Investigation"
	DetStatsMedicine      DetStats = "Medicine"
	DetStatsNature        DetStats = "Nature"
	DetStatsPerception    DetStats = "Perception"
	DetStatsPerformance   DetStats = "Performance"
	DetStatsPersuasion    DetStats = "Persuasion"
	DetStatsReligion      DetStats = "Religion"
	DetStatsSleight       DetStats = "Sleight"
	DetStatsOf            DetStats = "of"
	DetStatsHand          DetStats = "Hand"
	DetStatsStealth       DetStats = "Stealth"
	DetStatsSurvival      DetStats = "Survival"
)

var AllDetStats = []DetStats{
	DetStatsAcrobatics,
	DetStatsAnimal,
	DetStatsHandling,
	DetStatsArcana,
	DetStatsAthletics,
	DetStatsDeception,
	DetStatsHistory,
	DetStatsInsight,
	DetStatsIntimidation,
	DetStatsInvestigation,
	DetStatsMedicine,
	DetStatsNature,
	DetStatsPerception,
	DetStatsPerformance,
	DetStatsPersuasion,
	DetStatsReligion,
	DetStatsSleight,
	DetStatsOf,
	DetStatsHand,
	DetStatsStealth,
	DetStatsSurvival,
}

func (e DetStats) IsValid() bool {
	switch e {
	case DetStatsAcrobatics, DetStatsAnimal, DetStatsHandling, DetStatsArcana, DetStatsAthletics, DetStatsDeception, DetStatsHistory, DetStatsInsight, DetStatsIntimidation, DetStatsInvestigation, DetStatsMedicine, DetStatsNature, DetStatsPerception, DetStatsPerformance, DetStatsPersuasion, DetStatsReligion, DetStatsSleight, DetStatsOf, DetStatsHand, DetStatsStealth, DetStatsSurvival:
		return true
	}
	return false
}

func (e DetStats) String() string {
	return string(e)
}

func (e *DetStats) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DetStats(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DetStats", str)
	}
	return nil
}

func (e DetStats) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Stats string

const (
	StatsStrength     Stats = "Strength"
	StatsDexterity    Stats = "Dexterity"
	StatsConstitution Stats = "Constitution"
	StatsIntelligence Stats = "Intelligence"
	StatsWisdom       Stats = "Wisdom"
	StatsCharisma     Stats = "Charisma"
)

var AllStats = []Stats{
	StatsStrength,
	StatsDexterity,
	StatsConstitution,
	StatsIntelligence,
	StatsWisdom,
	StatsCharisma,
}

func (e Stats) IsValid() bool {
	switch e {
	case StatsStrength, StatsDexterity, StatsConstitution, StatsIntelligence, StatsWisdom, StatsCharisma:
		return true
	}
	return false
}

func (e Stats) String() string {
	return string(e)
}

func (e *Stats) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Stats(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Stats", str)
	}
	return nil
}

func (e Stats) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
