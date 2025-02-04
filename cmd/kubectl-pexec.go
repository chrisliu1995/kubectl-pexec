package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/ringtail/kubectl-pexec/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-pexec", pflag.ExitOnError)
	flags.String("ignore-hostname", "false", "ignore hostname output")
	pflag.CommandLine = flags

	root := cmd.NewPExecCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
