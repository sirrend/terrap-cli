package terraform_utils

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
	"strings"
)

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

type TerraformError struct {
	reason      string
	description string
}

// terraformErrorParser
/*
@brief:
	terraformErrorParser gets an error and parses it to `reason: explanation`
@params:
	err - error - the error to parse
@returns:
	[]TerraformError - a list of all errors as TerraformError objects
*/
func terraformErrorParser(err error) []TerraformError {
	var parsedError []TerraformError
	errorLines := strings.Split(err.Error(), "\n")

	for lineIndex, line := range errorLines {
		if strings.Contains(line, "Error:") || strings.Contains(line, "Error ") {
			var (
				description  string
				lastWasEmpty bool
			)

			for i, descriptionLine := range errorLines[lineIndex+1:] {
				if !strings.Contains(descriptionLine, "Error: ") && !strings.Contains(descriptionLine, "Error ") {
					if i != 0 && (descriptionLine == "" && lastWasEmpty) {
						lastWasEmpty = true
						continue
					} else {
						lastWasEmpty = false
						description += descriptionLine + "\n"
					}
				}

				if strings.Contains(descriptionLine, "Error: ") || strings.Contains(descriptionLine, "Error ") {
					break
				}
			}

			var tempError TerraformError
			if strings.Contains(line, "Error: ") {
				tempError.reason = utils.RemoveLastDot(strings.Split(line, "Error: ")[1])
			} else {
				tempError.reason = utils.RemoveLastDot(line)
			}

			tempError.description = description
			parsedError = append(parsedError, tempError)
		}
	}

	return parsedError
}

// removeIdenticalErrorParagraphs
/*
@brief:
	removeIdenticalErrorParagraphs finds identical error paragraphs in different error messages and trim them from each one
@params:
	err - error - the error to parse
*/
func removeIdenticalErrorParagraphs(tfErrors []TerraformError) []TerraformError {
	var (
		paragraphs       []string
		paraToRemove     []string
		modifiedTfErrors = tfErrors
	)

	// find all common paragraphs in errors
	if len(tfErrors) > 1 {
		for _, err := range tfErrors {
			for _, para := range strings.Split(err.description, "\n\n") { // split on empty line
				if !utils.IsItemInSlice(para, paragraphs) { // find non unique paragraphs
					paragraphs = append(paragraphs, para)

				} else if !utils.IsItemInSlice(para, paraToRemove) && para != "\n" { // add paragraphs to remove if unique
					paraToRemove = append(paraToRemove, para)
				}
			}
		}

		for _, err := range modifiedTfErrors {
			for _, remove := range paraToRemove {
				err.description = strings.ReplaceAll(err.description, remove, "")
			}
		}
	}

	return modifiedTfErrors
}

// CheckForSpecificErrors
/*
@brief:
	CheckForSpecificErrors prints Terrap oriented error messages for pre-known errors
@params:
	description - string - the new description message
*/
func CheckForSpecificErrors(err TerraformError) (description string) {
	if strings.Contains(strings.ToLower(err.reason), strings.ToLower("refreshing state: AccessDenied: Access Denied")) {
		return "Couldn't initialize a new Terrap workspace as the credentials in context weren't sufficient.\n" +
			"Please specify all needed configurations, such as environment variables, like you normally would when working with your local Terraform workspace."

	} else if strings.Contains(strings.ToLower(err.reason), strings.ToLower("error configuring S3 Backend")) {
		return "Couldn't initialize a new Terrap workspace as no credentials were found in context.\n" +
			"Please make sure your `~/.aws/credentials` file exists and properly configured."

	} else if strings.Contains(strings.ToLower(err.description), "version") &&
		strings.Contains(strings.ToLower(err.description), "terraform") &&
		!strings.Contains(strings.ToLower(err.description), "current platform") {

		return err.description +
			emoji.HammerAndWrench.String() +
			commons.YELLOW.Sprintf("\033[1m%s\033[0m",
				"  Possible Solution:\n"+
					"     | The error might be caused by using a Terraform version that is not compatible with the current infrastructure.\n"+
					"     | Please try using the `TERRAP_TERRAFORM_VERSION` environment variable to configure a specific Terraform version to use.")
	}

	return ""
}

// TerraformErrorPrettyPrint
/*
@brief:
	TerraformErrorPrettyPrint prints the terraform error given
@params:
	err - error - the error to parse
*/
func TerraformErrorPrettyPrint(err error) {
	tfErrors := terraformErrorParser(err)

	if len(tfErrors) > 1 {
		commons.RED.Println("Terrap failed with the following errors: ")
	} else if len(tfErrors) == 1 {
		commons.RED.Println("Terrap failed with the following error: ")
	}

	for _, errObj := range removeIdenticalErrorParagraphs(tfErrors) { // show only unique messages
		if terrapDesc := CheckForSpecificErrors(errObj); terrapDesc != "" {
			errObj.description = terrapDesc
		}

		// print error
		commons.RED.Print("  " + errObj.reason + ":\n")
		for i, line := range strings.Split(errObj.description, "\n") {
			if i == 0 && line == "" { // skip first line if empty
				continue
			}

			fmt.Println("    " + line)
		}

		fmt.Println() // separate next paragraph
	}
}
