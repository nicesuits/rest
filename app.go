package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initalizes DB
func (a *App) Initialize(user, password, dbName string) {}

// Run runs DB
func (a *App) Run(addr string) {}
