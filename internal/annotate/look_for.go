package annotate

import (
	"bufio"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"strings"
)

func FindAttributeInResourceDeclaration(resource handle_files.Resource, attribute string) int {
	var lines []string

	// Open the file for reading
	file, err := os.OpenFile(resource.Pos.Filename, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the file into a slice of strings
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var attributeLine string
	for pos, line := range lines {
		if strings.Contains(line, resource.Name) {
			codeBlock := utils.GetCodeUntilMatchingBrace(strings.Join(lines[pos:], "\n"))

			for _, functionLine := range strings.Split(codeBlock, "\n") {
				if strings.Contains(functionLine, attribute) {
					attributeLine = functionLine
					break
				}
			}
		}

		if strings.Contains(line, attributeLine) {
			return pos
		}
	}

	return 0
}
