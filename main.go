package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/devopsluka/doc-wizzard/docs"
	"github.com/devopsluka/doc-wizzard/parser"
	"github.com/fatih/color"
)

func main() {
	color.Cyan("Welcome to the DevOps Toolkit!\nThis is the Doc Wizzard.\n")

	var chartPath string
	if len(os.Args) == 2 {
		chartPath = os.Args[1]
	} else {
		fmt.Print("Enter the path to the Helm chart directory: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chartPath = scanner.Text()
	}

	docs.CleanSheet("output/documentation.txt")

	if err := parser.ParseHelmChart(chartPath); err != nil {
		if chartPath == "helm-chart" {
			filePath, err := parser.HelmRenderer(chartPath)
			if err != nil {
				log.Fatalf("Error in trying to template Helm Chart - You might have put wrong directory name: %v", err)
			}

			fmt.Println("Helm chart rendered to", filePath)
		} else {
			parser.ParseHelmChart(chartPath)
		}
	}
}
