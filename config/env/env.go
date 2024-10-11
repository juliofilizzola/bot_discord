package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port          = ""
	TokenDiscord  = ""
	AvatarURL     = ""
	Username      = ""
	DatabaseURL   = ""
	DbType        = ""
	AutoMigrateDb = ""
	SetEnv        = ""
)

func Env() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	AutoMigrateDb = fmt.Sprint(os.Getenv("AUTO_MIGRATE"))
	DbType = fmt.Sprint(os.Getenv("DB_TYPE"))
	DatabaseURL = fmt.Sprint(os.Getenv("DATABASE_URL"))
	Port = fmt.Sprint(os.Getenv("PORT"))
	TokenDiscord = fmt.Sprint(os.Getenv("TOKEN_DISCORD"))
	Username = fmt.Sprint(os.Getenv("USER_NAME"))
	AvatarURL = fmt.Sprint(os.Getenv("AVATAR_URL"))
	SetEnv = fmt.Sprint(os.Getenv("SET_ENV"))
}
