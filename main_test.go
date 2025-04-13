package spawnforge

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestCopyFs(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()

	// Copy the embedded filesystem to the temporary directory
	err := copyFs(content, tmpDir)
	require.NoError(t, err)

	// Check if the files exist in the temporary directory
	require.FileExists(t, tmpDir+"/test/content/file.txt")
	require.FileExists(t, tmpDir+"/test/content/another-file.txt")

	// Read the content of the copied file
	data, err := os.ReadFile(tmpDir + "/test/content/file.txt")
	require.NoError(t, err)
	require.Equal(t, "I am file", string(data))
}
