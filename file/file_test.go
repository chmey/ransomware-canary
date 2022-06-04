package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFileHash(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	testFileContent := []byte("test0123")
	testFileName := "testFileHash"

	t.Run("consistent hash", func(t *testing.T) {
		err := os.WriteFile(testFileName, testFileContent, 0644)
		require.NoError(err)
		h, err := GetFileHash(testFileName)
		require.NoError(err)
		h2, err := GetFileHash(testFileName)
		require.NoError(err)

		err = os.Remove(testFileName)
		require.NoError(err)

		assert.Equal(h, h2)

	})
	t.Run("detects change", func(t *testing.T) {
		err := os.WriteFile(testFileName, testFileContent, 0644)
		require.NoError(err)
		h, err := GetFileHash(testFileName)
		require.NoError(err)

		err = os.WriteFile(testFileName, []byte("changed"), 0644)
		require.NoError(err)

		h2, err := GetFileHash(testFileName)
		require.NoError(err)

		err = os.Remove(testFileName)
		require.NoError(err)

		assert.NotEqual(h, h2)
	})

	t.Run("missing file", func(t *testing.T) {
		require.NoFileExists(testFileName)
		_, err := GetFileHash(testFileName)
		assert.Error(err)
	})
}

func TestMakeFilePath(t *testing.T) {
	fileName := "testFile"
	_, err := MakeFilePath(fileName)
	require.NoError(t, err)
}
