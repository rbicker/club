package translations

import (
_ "github.com/rbicker/club/internal/server"
_ "golang.org/x/text/language"
_ "golang.org/x/text/message"
)

//go:generate go run golang.org/x/text/cmd/gotext -srclang=en update -out=catalog.go -lang=en,de

