package create

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/spf13/cobra"
)

func CreateCandidateCmd() *cobra.Command {
	var name string
	var email string
	var phoneNumber string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create candidate",
		Example: heredoc.Doc(`
			$ jobbox-cli candidate create -n "name" -e "email" -p "phone_number"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			candidate := util.Candidate{Name: name, Email: email, PhoneNumber: phoneNumber}

			db := util.Db()
			defer db.Close()

			createCandidate(db, candidate)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Example: Ali Piry")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Example: ali@piry.dev")
	cmd.Flags().StringVarP(&phoneNumber, "phoneNumber", "p", "", "Example: 09121234567")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("phoneNumber")

	return cmd
}

func createCandidate(db *sql.DB, candidate util.Candidate) {
	createCandidateQuery := `insert into candidates (name, email, phone_number) values (?, ?, ?);`

	statement, err := db.Prepare(createCandidateQuery)
	cobra.CheckErr(err)

	fmt.Println("Creating new candidate...")
	statement.Exec(candidate.Name, candidate.Email, candidate.PhoneNumber)
	fmt.Println("New candidate created")
}
