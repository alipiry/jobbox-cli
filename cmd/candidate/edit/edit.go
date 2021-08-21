package edit

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/spf13/cobra"
)

func UpdateCandidateCmd() *cobra.Command {
	var name string
	var email string
	var phoneNumber string
	var candidateId string

	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit candidate",
		Example: heredoc.Doc(`
			$ jobbox-cli candidate edit -i "candidateID" -n "name" -e "email" -p "phone_number"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			candidate := util.Candidate{Name: name, Email: email, PhoneNumber: phoneNumber}

			db := util.Db()
			defer db.Close()

			updateCandidate(db, candidateId, candidate)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Example: Ali Piry")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Example: ali@piry.dev")
	cmd.Flags().StringVarP(&phoneNumber, "phoneNumber", "p", "", "Example: 09121234567")
	cmd.Flags().StringVarP(&candidateId, "id", "i", "", "Example: 1")
	cmd.MarkFlagRequired("id")

	return cmd
}

func updateCandidate(db *sql.DB, candidateId string, candidate util.Candidate) {
	updateCandidateQuery := `update candidates set name = ?, email = ?, phone_number = ? where id = ?;`

	selectCandidateQuery := `select name, email, phone_number from candidates where id = ?;`
	var name string
	var email string
	var phoneNumber string

	row := db.QueryRow(selectCandidateQuery, candidateId)
	err := row.Scan(&name, &email, &phoneNumber)
	cobra.CheckErr(err)

	statement, err := db.Prepare(updateCandidateQuery)
	cobra.CheckErr(err)

	if len(candidate.Name) > 0 {
		name = candidate.Name
	}

	if len(candidate.Email) > 0 {
		email = candidate.Email
	}

	if len(candidate.PhoneNumber) > 0 {
		phoneNumber = candidate.PhoneNumber
	}

	fmt.Println("Updating candidate with id:", candidateId)
	statement.Exec(name, email, phoneNumber, candidateId)
	fmt.Println("Candidate with id:", candidateId, "updated")
}
