package jobs

import (
	"database/sql"
	"os"

	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func JobsCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "jobs",
		Short: "Show open jobs",
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			db := util.Db()
			defer db.Close()

			selectJobs(db)
		},
	}

	return cmd
}

func selectJobs(db *sql.DB) {
	rows, err := db.Query("select id, title, description, city from jobs;")
	cobra.CheckErr(err)

	defer rows.Close()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("JOBS")
	t.AppendHeader(table.Row{"ID", "Title", "Description", "City"})

	for rows.Next() {
		var id int
		var title string
		var description string
		var city string

		err = rows.Scan(&id, &title, &description, &city)
		cobra.CheckErr(err)

		t.AppendRow([]interface{}{id, title, description, city})
	}

	err = rows.Err()
	cobra.CheckErr(err)

	t.SetStyle(table.StyleLight)
	t.Render()
}
