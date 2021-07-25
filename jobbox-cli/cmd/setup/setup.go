package setup

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alipiry/jobbox/jobbox-cli/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func SetupCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "setup",
		Short: "Setup DB for your jobboard",
		Run: func(cmd *cobra.Command, args []string) {
			dbNotExists, basePath, dbPath := util.DbNotExists()

			if dbNotExists {
				fmt.Println("Creating DB into", basePath)

				basePathError := os.MkdirAll(basePath, 0755)
				cobra.CheckErr(basePathError)

				dbFile, dbFileError := os.Create(dbPath)
				cobra.CheckErr(dbFileError)

				dbFile.Close()

				fmt.Println("DB created")

				db := util.Db()

				createJobsTable(db)

			} else {
				fmt.Println("DB already created!")
			}

		},
	}

	return cmd
}

func createJobsTable(db *sql.DB) {
	jobsTableQuery := `create table jobs (
		"id" integer primary key autoincrement,		
    "title" text not null,
    "description" text not null,
    "city" text not null,
    "created_at" text default (datetime('now'))
	  );`

	fmt.Println("Creating jobs table...")

	statement, err := db.Prepare(jobsTableQuery)

	cobra.CheckErr(err)

	statement.Exec()

	fmt.Println("Jobs table created")
}
