package candidates

import (
	"database/sql"
	"os"

	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func CandidatesCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "candidates",
		Short: "Show existing candidates",
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			db := util.Db()
			defer db.Close()

			selectCandidates(db)
		},
	}

	return cmd
}

func selectCandidates(db *sql.DB) {
	rows, err := db.Query("select id, name, email, phone_number from candidates;")
	cobra.CheckErr(err)

	defer rows.Close()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("CANDIDATES")
	t.AppendHeader(table.Row{"ID", "Title", "Email", "PhoneNumber"})

	for rows.Next() {
		var id int
		var name string
		var email string
		var phoneNumber string

		err = rows.Scan(&id, &name, &email, &phoneNumber)
		cobra.CheckErr(err)

		t.AppendRow([]interface{}{id, name, email, phoneNumber})
	}

	err = rows.Err()
	cobra.CheckErr(err)

	t.SetStyle(table.StyleLight)
	t.Render()
}
