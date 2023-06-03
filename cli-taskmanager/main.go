package main

import (
	"cli-taskmanager/cmd"
	"cli-taskmanager/pkgs/config"
	"cli-taskmanager/pkgs/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	db := db.Initialize(filepath.Join(home, "task.db"))
	defer db.Connection.Close()
	c := config.Config{
		Database: db,
	}
	app := cmd.NewCmdApp(&c)
	app.Execute()
}
