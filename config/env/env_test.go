package env

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestEnvLoadsAllVariables(t *testing.T) {
	os.Clearenv()
	err := godotenv.Write(map[string]string{
		"AUTO_MIGRATE":  "true",
		"DB_TYPE":       "postgres",
		"DATABASE_URL":  "postgres://user:pass@localhost/db",
		"PORT":          "8080",
		"TOKEN_DISCORD": "token",
		"USER_NAME":     "username",
		"AVATAR_URL":    "http://avatar.url",
		"SET_ENV":       "development",
	}, ".env")
	if err != nil {
		return
	}

	Env()

	assert.Equal(t, "true", AutoMigrateDb)
	assert.Equal(t, "postgres", DbType)
	assert.Equal(t, "postgres://user:pass@localhost/db", DatabaseURL)
	assert.Equal(t, "8080", Port)
	assert.Equal(t, "token", TokenDiscord)
	assert.Equal(t, "username", Username)
	assert.Equal(t, "http://avatar.url", AvatarURL)
	assert.Equal(t, "development", SetEnv)
}

func TestEnvHandlesEmptyEnvVariables(t *testing.T) {
	os.Clearenv()
	err := godotenv.Write(map[string]string{
		"AUTO_MIGRATE":  "",
		"DB_TYPE":       "",
		"DATABASE_URL":  "",
		"PORT":          "",
		"TOKEN_DISCORD": "",
		"USER_NAME":     "",
		"AVATAR_URL":    "",
		"SET_ENV":       "",
	}, ".env")

	if err != nil {
		return
	}

	Env()

	assert.Equal(t, "", AutoMigrateDb)
	assert.Equal(t, "", DbType)
	assert.Equal(t, "", DatabaseURL)
	assert.Equal(t, "", Port)
	assert.Equal(t, "", TokenDiscord)
	assert.Equal(t, "", Username)
	assert.Equal(t, "", AvatarURL)
	assert.Equal(t, "", SetEnv)
}
