package cmd

import (
	"cli-taskmanager/pkgs/config"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type TaskCmd struct {
	config  *config.Config
	rootCmd *cobra.Command
	addCmd  *cobra.Command
	listCmd *cobra.Command
	doCmd   *cobra.Command
}

var rootCmd = &cobra.Command{
	Use:   "cli-taskmanager",
	Short: "A task management system in CLI, written with golang",
	Long:  "A barebone task management system CLI, written with golang and make use of Cobra, boltDB for its functionalities. This was meant to be an exercise for learning golang",
}

func NewCmdApp(config *config.Config) *TaskCmd {
	addCmd.Run = func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := config.Database.CreateTask(task)
		if err != nil {
			fmt.Printf("Error adding new task: %v\n", err)
		} else {
			fmt.Printf("New task added %d - %s\n", id, task)
		}
	}

	listCmd.Run = func(cmd *cobra.Command, args []string) {
		tasks, _ := config.Database.ListTasks()
		fmt.Println("---------------------")
		fmt.Printf("Total %d\n---------------------\n", len(tasks))
		for _, t := range tasks {
			fmt.Printf("   %d | %s\n", t.Key, t.Value)
		}
		fmt.Println("---------------------")
	}

	doCmd.Run = func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: %s is not an integer", arg)
			} else {
				ids = append(ids, id)
			}
		}
		errs := config.Database.DoTask(ids)
		if len(errs) > 0 {
			for _, err := range errs {
				fmt.Printf("Error deleting: %v\n", err)
			}
			return
		} else {
			fmt.Printf("Successfully deleted tasks: %v", ids)
		}

	}

	return &TaskCmd{
		config:  config,
		rootCmd: rootCmd,
		addCmd:  addCmd,
		listCmd: listCmd,
		doCmd:   doCmd,
	}
}

func (t *TaskCmd) Execute() {
	if err := t.rootCmd.Execute(); err != nil {
		log.Fatalf("Error building CMD %v", err)
	}
}
