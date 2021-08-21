package candidate

import (
	"github.com/MakeNowJust/heredoc"
	createCmd "github.com/alipiry/jobbox-cli/cmd/candidate/create"
	updateCmd "github.com/alipiry/jobbox-cli/cmd/candidate/edit"
	"github.com/spf13/cobra"
)

func CandidateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "candidate <command>",
		Short: "Create or edit candidate",
		Example: heredoc.Doc(`
			$ jobbox-cli candidate create
			$ jobbox-cli candidate edit
		`),
	}

	cmd.AddCommand(createCmd.CreateCandidateCmd())
	cmd.AddCommand(updateCmd.UpdateCandidateCmd())

	return cmd
}
