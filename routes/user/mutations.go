package user

import (
	"dndinator/dndinator-api/database"
	"dndinator/dndinator-api/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserExists(username, email string) bool {
	var data model.User

	query := `
  SELECT * FROM users
  WHERE email=$1 OR username=$2`

	database.Db.QueryRowx(query, email, username).StructScan(&data)
	if len(*data.UserName) >= 1 {
		return true
	}

	return false
}

func Login(c *gin.Context) {
	var data model.User
	var body *model.User

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed")
		return
	}

	query := `
  SELECT * FROM users
  WHERE username=$1 AND password=$2`

	database.Db.QueryRowx(query, body.UserName, body.Password).StructScan(&data)

	if *data.ID > 0 {
		// token, err := auth.CreateJWT()
		// if err != nil {
		// 	return
		// }

		c.IndentedJSON(200, &data)
	} else {
		c.IndentedJSON(500, &data)
	}
}

func SignUp(c *gin.Context) {
	var body *model.User

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed")
		return
	}

	// TODO Check if user exists
	existing := UserExists(*body.UserName, *body.Password)
	if existing {
		c.IndentedJSON(500, "This username or password already exists")
	} else {
		query := `
        INSERT INTO users
        (username, email, "password")
        VALUES($1, $2, $3)`

		database.Db.QueryRowx(query, body.UserName, body.Email, body.Password)
		c.IndentedJSON(200, "Success")
	}
}

func JANKLogin(c *gin.Context) {
	var data model.User
	var body *model.User

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind failed")
		return
	}

	query := `
  SELECT * FROM users
  WHERE username=$1`

	database.Db.QueryRowx(query, body.UserName).StructScan(&data)

	if *data.ID > 0 {
		c.IndentedJSON(200, &data)
	} else {
		c.IndentedJSON(500, &data)
	}
}
