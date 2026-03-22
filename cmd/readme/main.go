package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func gitRoot(dir string) string {
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	root := gitRoot(cwd)
	if root == "" {
		fmt.Fprintln(os.Stderr, "not a git repository")
		os.Exit(1)
	}
	candidate := filepath.Join(root, "README.md")
	if _, err := os.Stat(candidate); err != nil {
		fmt.Fprintln(os.Stderr, "no README.md found")
		os.Exit(1)
	}
	c := exec.Command("bat", candidate)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
