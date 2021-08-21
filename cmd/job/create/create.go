package create

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alipiry/jobbox-cli/util"
	"github.com/spf13/cobra"
)

func CreateJobCmd() *cobra.Command {
	var title string
	var description string
	var city string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create job",
		Example: heredoc.Doc(`
			$ jobbox-cli job create -t "title" -d "desc" -c "city"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			job := util.Job{Title: title, Description: description, City: city}

			db := util.Db()
			defer db.Close()

			createJob(db, job)
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "Example: Front-end developer")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Example: JS, React, NextJS")
	cmd.Flags().StringVarP(&city, "city", "c", "", "Example: Tehran")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("city")

	return cmd
}

func createJob(db *sql.DB, job util.Job) {
	createJobQuery := `insert into jobs (title, description, city) values (?, ?, ?);`

	statement, err := db.Prepare(createJobQuery)
	cobra.CheckErr(err)

	fmt.Println("Creating new job...")
	statement.Exec(job.Title, job.Description, job.City)
	fmt.Println("New job created")
}
