package main

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	reader "github.com/observiq/bolt-explorer/db_reader"
	"github.com/observiq/bolt-explorer/model"
	"github.com/observiq/bolt-explorer/router"
	"github.com/observiq/bolt-explorer/style"
	"go.etcd.io/bbolt"
)

func main() {
	// validate first arg is filepath
	filepath, err := getFilepath()
	if err != nil {
		fmt.Println(err)
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

// getFilepath return the filepath and validates that
// 1) the first argument exists and
// 2) the first argument is a file that exists.
func getFilepath() (string, error) {
	if len(os.Args) <= 1 {
		return "", errors.New("filepath argument is required")
	}

	// validate file path exists
	filepath := os.Args[1]
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return "", errors.New("database file does not exist")
	}

	return filepath, nil
}
