package pkgenv

import "github.com/joho/godotenv"

func New() {
	godotenv.Load()
}
