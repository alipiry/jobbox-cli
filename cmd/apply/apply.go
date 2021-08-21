package apply

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alipiry/jobbox/jobbox-cli/util"
	"github.com/spf13/cobra"
)

func ApplyCmd() *cobra.Command {
	var name string
	var email string
	var phoneNumber string
	var jobId string

	var cmd = &cobra.Command{
		Use:   "apply",
		Short: "Apply for open jobs",
		Example: heredoc.Doc(`
			$ jobbox-cli apply -i "jobID" -n "name" -e "email" -p "phone_number"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			util.ShowErrorMessageIfDbNotExists()

			candidate := util.Candidate{Name: name, Email: email, PhoneNumber: phoneNumber}

			db := util.Db()
			defer db.Close()

			createJobCandidate(db, jobId, candidate)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Example: Ali Piry")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Example: ali@piry.dev")
	cmd.Flags().StringVarP(&phoneNumber, "phoneNumber", "p", "", "Example: 09121234567")
	cmd.Flags().StringVarP(&jobId, "id", "i", "", "Example: 1")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("phoneNumber")
	cmd.MarkFlagRequired("id")

	return cmd
}

func createJobCandidate(db *sql.DB, jobId string, candidate util.Candidate) {
	jobSelect := `select id from jobs where id = ?;`

	err := db.QueryRow(jobSelect, jobId).Scan(&jobId)
	if err != nil {
		fmt.Println("job does not exist!")
		return
	}

	createCandidateQuery := `insert into candidates (name, email, phone_number) values (?, ?, ?);`

	statement, err := db.Prepare(createCandidateQuery)
	cobra.CheckErr(err)

	fmt.Println("Creating new candidate...")
	res, err := statement.Exec(candidate.Name, candidate.Email, candidate.PhoneNumber)
	cobra.CheckErr(err)

	candidateId, err := res.LastInsertId()
	cobra.CheckErr(err)

	fmt.Println("New candidate created")

	createJobCandidateQuery := `insert into job_candidates (job_id, candidate_id) values (?, ?);`

	_statement, err := db.Prepare(createJobCandidateQuery)
	cobra.CheckErr(err)

	fmt.Println("Creating new application...")
	_statement.Exec(jobId, candidateId)

	fmt.Println("New application created")

}
