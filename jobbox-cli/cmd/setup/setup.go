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
				defer db.Close()

				jobsQuery := `create table jobs (
					"id" integer primary key autoincrement,		
					"title" text not null,
					"description" text not null,
					"city" text not null,
					"created_at" text default (datetime('now'))
				);`

				candidatesQuery := `create table candidates (
					"id" integer primary key autoincrement,		
					"name" text not null,
					"email" text not null,
					"phone_number" text not null,
					"created_at" text default (datetime('now'))
				);`

				jobCandidatesQuery := `create table job_candidates (
					"id" integer primary key autoincrement,		
					"job_id" integer not null references jobs(id),
					"candidate_id" integer not null references candidates(id),
					"created_at" text default (datetime('now'))
				);`

				createTable(db, jobsQuery, "jobs")
				createTable(db, candidatesQuery, "candidates")
				createTable(db, jobCandidatesQuery, "job_candidates")

			} else {
				fmt.Println("DB already created!")
			}

		},
	}

	return cmd
}

func createTable(db *sql.DB, tableQuery string, tableName string) {
	statement, err := db.Prepare(tableQuery)
	cobra.CheckErr(err)

	fmt.Println("Creating", tableName, "table...")
	statement.Exec()
	fmt.Println(tableName, "table created")
}
