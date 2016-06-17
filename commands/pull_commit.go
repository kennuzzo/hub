package commands

import "log"

var cmdPRCommit = &Command{
	Run:          PrCommit,
	GitExtension: false,
	Usage: `
	pr-commit [-fo] [-b <BASE>] [-h <HEAD>] [-a <USER>] [-M <MILESTONE>] [-l <LABELS>]
	pr-commit -m <MESSAGE>
	pr-commit -F <FILE>
	pr-commit -i <ISSUE> `,
	Long: `Create a GitHub pull request with id.
  It accepts the same values as the pull-request.
`,
}

func init() {
	CmdRunner.Use(cmdPRCommit)
}

func PrCommit(cmd *Command, args *Args) {
	// utils.Check(github.FormatError("creating Merge request", err))
	log.Println("createing pull-request")
	log.Println("amending current commit")
	log.Println("pushing changes.")
	pullRequest(cmd, args)
	return
}
