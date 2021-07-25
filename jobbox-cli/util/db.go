package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Job struct {
	Title       string
	Description string
	City        string
}

func DbNotExists() (bool, string, string) {
	home, homeError := os.UserHomeDir()
	cobra.CheckErr(homeError)

	basePath := home + "/.jobbox"
	dbPath := basePath + "/app.db"

	_, pathError := os.Stat(dbPath)

	return os.IsNotExist(pathError), basePath, dbPath
}

func Db() *sql.DB {
	_, _, dbPath := DbNotExists()

	db, _ := sql.Open("sqlite3", dbPath)

	return db
}

func ShowErrorMessageIfDbNotExists() {
	dbNotExists, _, _ := DbNotExists()

	if dbNotExists {
		fmt.Println(`DB does not exist, please run setup command:
jobbox-cli setup`)
		return
	}
}
