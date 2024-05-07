package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/dvordrova/gop/codegen"
	"github.com/dvordrova/gop/tui/service"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	var rootCmd = &cobra.Command{Use: "gop"}

	var serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Create/edit service",
		Long:  `This is a longer description of service.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(service.NewModel(), tea.WithAltScreen())
			m, err := p.Run()
			_ = m
			return err
		},
	}
	var projectCmd = &cobra.Command{
		Use:   "project",
		Short: "[not implemented] create/edit project",
		Long:  `This is a longer description of project.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
	}

	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate service",
		Long:  `This is a longer description of gen.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return codegen.GenServiceStruct()
		},
	}

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "[not implemented] start service/project",
		Long:  `This is a longer description of start.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
	}

	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "[not implemented] stop running service/project",
		Long:  `This is a longer description of stop.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("not implemented yet")
		},
	}

	rootCmd.AddCommand(serviceCmd, projectCmd, genCmd, startCmd, stopCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
