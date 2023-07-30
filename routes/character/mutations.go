package character

import (
	"dndinator/dndinator-api/database"
	"dndinator/dndinator-api/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	var data model.Character
	userId := c.Query("id")

	query := `
  SELECT * FROM character
  WHERE playerid=$1`

	database.Db.QueryRowx(query, userId).StructScan(&data)
	fmt.Println(data)

	if *data.ID <= 0 {
		c.IndentedJSON(500, "Please create a character first")
	} else {
		c.IndentedJSON(200, data)
	}
}

func CreateCharacter(c *gin.Context) {
	var data model.Character
	var body *model.Character

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed")
		return
	}

	query := `

        INSERT INTO "character"
        ("name", lvl, race, alignment, background, hp, maxhp, ac, init, playerid, party)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	err := database.Db.QueryRowx(query, body.Name, body.Lvl, body.Race, body.Alignment, body.Background, body.Hp, body.Maxhp, body.Ac, body.Init, body.Playerid, body.Party).StructScan(&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.IndentedJSON(200, &data)
}
