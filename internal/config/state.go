package config

import "bootfeeds/internal/database"

type State struct {
	DB  *database.Queries
	Cfg *Config
}
