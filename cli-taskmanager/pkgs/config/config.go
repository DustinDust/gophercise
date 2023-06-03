package config

import (
	"cli-taskmanager/pkgs/db"
)

type Config struct {
	Database *db.DB
}
