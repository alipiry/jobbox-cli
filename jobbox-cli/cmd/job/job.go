package job

import (
	"github.com/MakeNowJust/heredoc"
	createCmd "github.com/alipiry/jobbox/jobbox-cli/cmd/job/create"
	updateCmd "github.com/alipiry/jobbox/jobbox-cli/cmd/job/edit"
	"github.com/spf13/cobra"
)

func JobCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "job <command>",
		Short: "Create or edit job",
		Example: heredoc.Doc(`
			$ jobbox-cli job create 
			$ jobbox-cli job edit 
		`),
	}

	cmd.AddCommand(createCmd.CreateJobCmd())
	cmd.AddCommand(updateCmd.UpdateJobCmd())

	return cmd
}
