package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestRootCmd(t *testing.T) {

	testRootPath := "/tmp"

	cmd := newRootCmd()

	cmd.SetArgs([]string{
		fmt.Sprintf("--root=%s", testRootPath),
		"https://github.com/KDE/dummy.git",
	})

	err := cmd.Execute()
	if err != nil {
		t.Error(err)
	}

	expectedRepoPath := "/tmp/github.com/KDE/dummy"
	if _, err := os.Stat(expectedRepoPath); os.IsNotExist(err) {
		t.Errorf("expected repo to be cloned to %s. But path does not exist", expectedRepoPath)
	}

	err = os.RemoveAll(expectedRepoPath)
	if err != nil {
		t.Errorf("error while cleanup: %v", err)
	}
}
