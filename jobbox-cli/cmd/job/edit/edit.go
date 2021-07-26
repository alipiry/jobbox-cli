package edit

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/spf13/cobra"
)

func UpdateJobCmd() *cobra.Command {
	var title string
	var description string
	var city string
	var jobId string

	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit job",
		Example: heredoc.Doc(`
			$ jobbox-cli job edit -i "jobID" -t "title" -d "desc" -c "city"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			job := util.Job{Title: title, Description: description, City: city}

			db := util.Db()
			defer db.Close()

			updateJob(db, jobId, job)
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "Example: Front-end developer")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Example: JS, React, NextJS")
	cmd.Flags().StringVarP(&city, "city", "c", "", "Example: Tehran")
	cmd.Flags().StringVarP(&jobId, "id", "i", "", "Example: 1")
	cmd.MarkFlagRequired("id")

	return cmd
}

func updateJob(db *sql.DB, jobId string, job util.Job) {
	updateJobQuery := `update jobs set title = ?, description = ?, city = ? where id = ?;`

	selectJobQuery := `select title, description, city from jobs where id = ?;`
	var title string
	var description string
	var city string

	row := db.QueryRow(selectJobQuery, jobId)
	err := row.Scan(&title, &description, &city)
	cobra.CheckErr(err)

	statement, err := db.Prepare(updateJobQuery)
	cobra.CheckErr(err)

	if len(job.Title) > 0 {
		title = job.Title
	}

	if len(job.Description) > 0 {
		description = job.Description
	}

	if len(job.City) > 0 {
		city = job.City
	}

	fmt.Println("Updating job with id:", jobId)
	statement.Exec(title, description, city, jobId)
	fmt.Println("Job with id:", jobId, "updated")
}
