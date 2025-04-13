package main

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test/content
var content embed.FS

//go:embed all:test/content/file.txt
var file string

func TestEmbedFile(t *testing.T) {
	require.Equal(t, "I am file", file)
}

func TestEmbedDir(t *testing.T) {
	// Read the directory content
	dirContent, err := content.ReadDir("test/content")
	require.NoError(t, err)

	// Check the number of files in the directory
	require.Equal(t, 2, len(dirContent))

	names := make([]string, len(dirContent))
	for i, entry := range dirContent {
		names[i] = entry.Name()
	}
	require.ElementsMatch(t, []string{"another-file.txt", "file.txt"}, names)
}
