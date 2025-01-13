package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFilepath(t *testing.T) {
	// mock os.Args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	os.Args = []string{"bolt-explorer", "testfiles/test.db"}

	filepath, err := getFilepath()
	require.NoError(t, err)
	require.Equal(t, "testfiles/test.db", filepath)
}

func TestGetFilepath_errors(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		errorMsg string
	}{
		{name: "no filepath arg", args: []string{}, errorMsg: "filepath argument is required"},
		{name: "file dne", args: []string{"bolt-explorer", "dne.db"}, errorMsg: "database file does not exist"},
	}

	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Args = test.args
			_, err := getFilepath()
			require.Error(t, err)
			require.Contains(t, err.Error(), test.errorMsg)
		})
	}
}
