package annotate

import (
	"bufio"
	"fmt"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"os"
)

func AddLineInPosition(resource handle_files.Resource, newLine string, pos int) {
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

	// Insert the new line at the given position
	lines = append(lines[:pos], append([]string{newLine}, lines[pos:]...)...)

	// Truncate the file to 0 bytes
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// Write the modified lines back to the file
	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	// write lines to file
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err = writer.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func AddAnnotationByRuleSet(resource handle_files.Resource, ruleSet rules_api.RuleSet) {
	for _, rule := range ruleSet.Rules {
		pos := FindAttributeInResourceDeclaration(resource, rule.Path)

		if pos != 0 {
			AddLineInPosition(resource, "\n# "+rule.Notification, pos)
		}
	}
}
