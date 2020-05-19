package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// Generates uuid v4
func Uuid() string {
	u := uuid.Must(uuid.NewV4(), nil).String()
	return strings.ReplaceAll(u, "-", "")
}
