package annotate

import (
	"bufio"
	"fmt"
	"github.com/sirrend/terrap-cli/internal/files_handler"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"strings"
)

func FindAttributeInResourceDeclaration(resource files_handler.Resource, path string) int {
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

	// declare variables for method
	posSum := 0
	splintedPath := strings.Split(path, ".")

	for pos, line := range lines {
		if strings.Contains(line, fmt.Sprintf("%s \"%s\" \"%s\"", resource.Type, resource.Name, resource.Alias)) {
			codeBlock := utils.GetCodeUntilMatchingBrace(strings.Join(lines[pos:], "\n"))

		OuterLoop:
			for _, component := range splintedPath {
				splintedCodeBlock := strings.Split(codeBlock, "\n")
				for functionPos, functionLine := range splintedCodeBlock {
					if strings.Contains(functionLine, component) && !strings.Contains(functionLine, "#") { // validate not a comment
						if component == splintedPath[len(splintedPath)-1] {
							posSum += functionPos + pos // add to line in file sum
							return posSum               // return position in file

						} else {
							codeBlock = strings.Join(strings.Split(codeBlock, "\n")[functionPos:], "\n") // continue from next codeBlock line after break
							posSum += functionPos                                                        // add to line in file sum

							break
						}
					}

					if splintedCodeBlock[len(splintedCodeBlock)-1] == functionLine {
						break OuterLoop
					}
				}
			}
		}
	}

	return 0
}
