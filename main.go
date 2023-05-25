package main

import (
	"fmt"
	"os"

	// "github.com/observiq/observiq-otel-cli/fetcher"
	// "github.com/observiq/observiq-otel-cli/internal/otlp"

	tea "github.com/charmbracelet/bubbletea"
	reader "github.com/observiq/bolt-explorer/Reader"
	"github.com/observiq/bolt-explorer/model"
	"github.com/observiq/bolt-explorer/router"
	"github.com/observiq/bolt-explorer/style"
	"go.etcd.io/bbolt"
)

func main() {

	db, err := bbolt.Open("/Users/dsvanlani/.bindplane/storage", 0666, nil)
	if err != nil {
		fmt.Println("Failed to open DB, %w", err)
		os.Exit(1)
	}
	reader.ParseDB(db)
	keys, err := reader.ParseDB(db)

	router := router.NewRouter(keys, db)
	styles := style.DefaultStyles()
	app := model.NewApp(router, styles)
	program := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if err != nil {
		fmt.Println("Failed to parse DB, %w", err)
		os.Exit(1)
	}

	if err := program.Start(); err != nil {
		os.Exit(1)
	}
}
