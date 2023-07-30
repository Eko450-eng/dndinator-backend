package main

import (
	"dndinator/dndinator-api/database"
	"dndinator/dndinator-api/routes/character"
	"dndinator/dndinator-api/routes/spells"
	"dndinator/dndinator-api/routes/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/spells", spells.GetAllSpells)
	router.GET("/myspells", spells.GetAllSpells)
	router.GET("/spell", spells.GetSpell)

	router.POST("/addspellbook", spells.AddSpell)
	router.POST("/getmyspellbook", spells.GetMySpells)

	router.GET("/character", character.GetUserById)
	router.POST("/character", character.CreateCharacter)

	router.POST("/login", user.Login)
	router.POST("/loginJWTJankness", user.JANKLogin)
	router.POST("/signup", user.SignUp)

	gin.SetMode(gin.ReleaseMode)
	router.Run("192.168.0.38:4000")
	router.Run("localhost:4000")
}
