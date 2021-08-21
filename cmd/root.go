package cmd

import (
	"github.com/alipiry/jobbox-cli/cmd/apply"
	"github.com/alipiry/jobbox-cli/cmd/candidate"
	"github.com/alipiry/jobbox-cli/cmd/candidates"
	"github.com/alipiry/jobbox-cli/cmd/job"
	"github.com/alipiry/jobbox-cli/cmd/jobs"
	"github.com/alipiry/jobbox-cli/cmd/setup"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jobbox-cli",
	Short: "Make your own jobboard in terminal",
	Long: `Jobbox CLI is a tool to create your customized jobboard.
You can create or edit your job positions.
You can apply for your open job positions and save the candidates' info.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(setup.SetupCmd())
	rootCmd.AddCommand(job.JobCmd())
	rootCmd.AddCommand(jobs.JobsCmd())
	rootCmd.AddCommand(candidate.CandidateCmd())
	rootCmd.AddCommand(candidates.CandidatesCmd())
	rootCmd.AddCommand(apply.ApplyCmd())
}
