package parser

import (
	"fmt"
	"os"
	"os/exec"
)

func HelmRenderer(dir string) (string, error) {
	cmd := exec.Command("helm", "template", dir)

	file, err := os.Create("output/template.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	cmd.Stdout = file

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}

	return "template.yaml", nil
}
