package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneRepository(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "gondalf_")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	url := "https://github.com/developia-io/gandalf.git"

	err = CloneRepository(url, tempDir)

	assert.NoError(t, err)

	_, err = os.Stat(tempDir + "/.git")
	assert.False(t, os.IsNotExist(err))
}
