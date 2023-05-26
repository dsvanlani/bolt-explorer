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
	// get the first arge as filepath
	filepath := os.Args[1]
	if filepath == "" {
		fmt.Println("No filepath argument provided")
		os.Exit(1)
	}

	fmt.Println(filepath)

	// check that file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	db, err := bbolt.Open(filepath, 0666, nil)
	if err != nil {
		fmt.Println("Failed to open DB, %w", err)
		os.Exit(1)
	}

	keys, err := reader.ParseDB(db)
	if err != nil {
		fmt.Println("Failed to parse DB, %w", err)
		os.Exit(1)
	}

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
