package spells

import (
	"dndinator/dndinator-api/database"
	"dndinator/dndinator-api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMySpells(c *gin.Context) {
	type GetMySpellsResponse struct {
		Spell []*model.Spell
		Slots int
		Tier  int
	}

	var character model.Character
	var spellBook model.Spellbook
	var spellBookSpells []*model.Spellbook_spell
	var spells []*model.Spell
	var body *AddSpellType

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed ", err.Error())
		return
	}

	getCharacter := `SELECT * FROM character WHERE playerid=$1`
	getSpellbook := `SELECT * FROM spellbook WHERE characterid=$1`
	getSpellbookSpell := `SELECT * FROM spellbook_spell WHERE book=$1`
	getSpells := `SELECT * FROM spell WHERE id=$1`

	database.Db.QueryRowx(getCharacter, body.User).StructScan(&character)
	database.Db.QueryRowx(getSpellbook, character.ID).StructScan(&spellBook)
	rows, err := database.Db.Queryx(getSpellbookSpell, spellBook.ID)
	if err != nil {
		fmt.Println("This is what you have fucked up: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var spellBookSpell model.Spellbook_spell
		rows.StructScan(&spellBookSpell)
		spellBookSpells = append(spellBookSpells, &spellBookSpell)
		var spell model.Spell
		database.Db.QueryRowx(getSpells, spellBookSpell.Spell).StructScan(&spell)
		spells = append(spells, &spell)
	}

	res := GetMySpellsResponse{
		Spell: spells,
		Slots: *spellBook.Slots,
		Tier:  *spellBook.Tier,
	}
	c.IndentedJSON(http.StatusOK, res)
}

func GetAllSpells(c *gin.Context) {
	var data []*model.Spell

	query := `SELECT * FROM spell`
	rows, err := database.Db.Queryx(query)
	if err != nil {
		fmt.Println("This is what you have fucked up: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var spell model.Spell
		rows.StructScan(&spell)
		data = append(data, &spell)
	}

	c.IndentedJSON(http.StatusOK, data)
}

func GetSpell(c *gin.Context) {
	var data []*model.Spell

	query := `
  SELECT * FROM spell
  WHERE name=$1
  VALUES($1)
  `
	fmt.Println("Context: ", c.Request.Body)
	database.Db.QueryRowx(query, "Acid Rain").StructScan(data)

	c.IndentedJSON(http.StatusOK, data)
}

type AddSpellType struct {
	User  *string
	Spell *int
	Book  *int
}

// TODO CHECK IF SPELL IS ALREADY EQUIPPED
func AddSpell(c *gin.Context) {
	var data []*model.Spellbook_spell
	var body *AddSpellType

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed: ", err.Error())
		return
	}

	query := `
    INSERT INTO spellbook_spell
    (book, spell)
    VALUES($1, $2);`

	database.Db.QueryRowx(query, body.Book, body.Spell).StructScan(data)

	c.IndentedJSON(http.StatusOK, data)
}
