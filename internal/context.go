package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Context struct {
	Env map[string]string
}

func NewContext() *Context {
	if os.Getenv("IS_PROD") != "" {
		godotenv.Load()
	}

	return &Context{
		Env: map[string]string{
			"": "",
		},
	}
}
