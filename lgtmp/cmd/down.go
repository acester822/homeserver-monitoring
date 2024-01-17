// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/qclaogui/codelab-monitoring/internal/cmd/all"
	"github.com/qclaogui/codelab-monitoring/internal/cmd/logs"
	"github.com/qclaogui/codelab-monitoring/internal/cmd/metrics"
	"github.com/qclaogui/codelab-monitoring/internal/cmd/profiles"
	"github.com/qclaogui/codelab-monitoring/internal/cmd/traces"

	"github.com/spf13/cobra"
)

func NewCmdDown() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "down",
		Short: "Stop LGTMP stack.",
		Long:  "Provisioning LGTMP stack by Docker Compose.",
		Example: heredoc.Doc(`
			# lgtmp down <command>
			$ lgtmp down metrics
		`),
	}

	cmd.AddCommand(all.NewCmdAll())
	cmd.AddCommand(logs.NewCmdLogs())
	cmd.AddCommand(metrics.NewCmdMetrics())
	cmd.AddCommand(profiles.NewCmdProfiles())
	cmd.AddCommand(traces.NewCmdTraces())

	return cmd
}
