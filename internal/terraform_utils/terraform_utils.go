package terraform_utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	ver "github.com/hashicorp/terraform/version"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type terraformVersion struct {
	Terraform_version string
}

// RemoveTempTerraformExecutor
/*
@brief:
	RemoveTempTerraformExecutor deletes the Terraform executor received
@params:
	execPath - string - the path to remove
@returns:
	error - if exists, else nil
*/
func RemoveTempTerraformExecutor(execPath string) error {
	if utils.FileExists(execPath) {
		splintedPath := strings.Split(execPath, "/")
		splintedPath = splintedPath[:len(splintedPath)-1]

		execPath = strings.Join(splintedPath, "/")
		err := os.RemoveAll(execPath)
		if err != nil {
			err = errors.New("error: failed to remove temporary terraform executor")
			return err
		}
	}

	return nil
}

/*
@brief: IsTerraformInstalled checks if terraform is installed
@
@returns: true if terraform is installed, false otherwise
*/

func IsTerraformInstalled() bool {
	_, err := exec.LookPath("terraform") // checks if terraform is in path
	return err == nil
}

/*
@brief: InstallTf installs terraform if not present
@
@returns: execPath - the path to the Terraform executable
@         tv - the Terraform version
*/

func InstallTf() (execPath string, isTempVersion bool, tv string) {
	terraformUserVersion := os.Getenv("TERRAP_TERRAFORM_VERSION") // user decided he wants a specific version

	if IsTerraformInstalled() && terraformUserVersion == "" {
		var tVersion terraformVersion
		path, err := exec.LookPath("terraform")
		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, " Terrap failed with the following error: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		j, err := exec.Command("terraform", "version", "--json").Output()

		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, " Terrap failed while fetching Terraform version: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(j, &tVersion)

		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, " Terrap failed with the following error: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return path, false, tVersion.Terraform_version

	} else if terraformUserVersion != "" {
		_, _ = commons.YELLOW.Println(emoji.DesktopComputer, " Using Terraform version:", terraformUserVersion)

		// set installer details
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(terraformUserVersion)),
		}

		// install terraform in context of the given directory
		execPath, err := installer.Install(context.Background())
		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, " Terrap failed while installing Terraform: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return execPath, true, ver.Version
	} else {
		_, _ = commons.YELLOW.Println(emoji.DesktopComputer, " Using Terraform version:", ver.Version)

		// set installer details
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(ver.Version)),
		}

		// install terraform in context of the given directory
		execPath, err := installer.Install(context.Background())
		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, " Terrap failed while installing Terraform: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return execPath, true, ver.Version
	}
}

// FindTfProviders
/*
@brief: FindTfProviders
	finds the Terraform providers in the given folder
@params:
	dir - the folder to find the Terraform providers in
@returns:
	the Terraform providers in the given folder
*/
func FindTfProviders(dir string, main *workspace.Workspace) map[string]*version.Version {
	tf := NewTerraformExecutor(dir, (*main).ExecPath)

	_, providersList, err := tf.Version(context.Background(), true)

	if err != nil {
		log.Fatalf("error getting terraform providers: %s", err)
	}

	(*main).Providers = providersList

	return providersList
}

// NewTerraformExecutor
/*
@brief: NewTerraformExecutor creates a new terraform executor

@params: dir - the directory to run terraform in
        execPath - the path to the Terraform executable
@returns:
	*tfexec.Terraform - the terraform executor
*/
func NewTerraformExecutor(dir string, execPath string) *tfexec.Terraform {
	dir, _ = filepath.Abs(dir)
	tf, err := tfexec.NewTerraform(dir, execPath)
	if err != nil {
		_, _ = commons.RED.Print("\n", emoji.CrossMark, " Terrap failed with the following error: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return tf
}

/*
@brief:
	terraformInit performs the `terraform init` command on the given folder
@params:
	dir - the folder to initialize
@returns:
	execPath - string - the Terraform executor
	isTempVersion - bool - is a temporary terraform executor
	terraformToolVersion - string - the tool version
	err - error - if exists, else nil
*/

func TerraformInit(dir string) (execPath string, isTempVersion bool, terraformToolVersion string, err error) {
	execPath, isTempVersion, terraformToolVersion = InstallTf()

	tf := NewTerraformExecutor(dir, execPath)

	// initialize terraform in the given folder
	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		if strings.Contains(err.Error(), "net/http") {
			err = errors.New("error: Terrap experienced some networking issues, please make sure you have a stable internet connection")
			return "", false, "", err
		} else {
			return "", false, "", err
		}
	}

	return execPath, isTempVersion, terraformToolVersion, nil
}

// terraformErrorParser
/*
@brief:
	terraformErrorParser gets an error and parses it to `reason: explanation`
@params:
	err - error - the error to parse
@returns:
	map[string]string - the parsed error as a map of `reason: explanation`
*/
func terraformErrorParser(err error) []map[string]string {
	var parsedError []map[string]string
	errorLines := strings.Split(err.Error(), "\n")

	for lineIndex, line := range errorLines {
		if strings.Contains(line, "Error:") || strings.Contains(line, "Error ") {
			var description string

			for _, descriptionLine := range errorLines[lineIndex+1:] {
				if !strings.Contains(descriptionLine, "Error: ") || !strings.Contains(descriptionLine, "Error ") {
					description += descriptionLine + "\n"
				}

				if strings.Contains(descriptionLine, "Error: ") || strings.Contains(descriptionLine, "Error ") {
					break
				}
			}

			if strings.Contains(line, "Error: ") {
				parsedError = append(parsedError, map[string]string{strings.Split(line, "Error: ")[1]: description})
			} else if strings.Contains(line, "Error ") {
				parsedError = append(parsedError, map[string]string{line: description})
			}
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
func removeIdenticalErrorParagraphs(tfErrors []map[string]string) []map[string]string {
	var (
		paragraphs       []string
		paraToRemove     []string
		modifiedTfErrors = tfErrors
	)

	// find all common paragraphs in errors
	if len(tfErrors) > 1 {
		for _, err := range tfErrors {
			for _, description := range err {
				for _, para := range strings.Split(description, "\n\n") { // split on empty line
					if !utils.IsItemInSlice(para, paragraphs) { // find non unique paragraphs
						paragraphs = append(paragraphs, para)

					} else if !utils.IsItemInSlice(para, paraToRemove) && para != "\n" { // add paragraphs to remove if unique
						paraToRemove = append(paraToRemove, para)
					}
				}
			}
		}

		for _, err := range modifiedTfErrors {
			for k, description := range err {
				for _, remove := range paraToRemove {
					err[k] = strings.ReplaceAll(description, remove, "")
				}
			}
		}
	}

	return modifiedTfErrors
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

	for _, err := range removeIdenticalErrorParagraphs(tfErrors) {
		for k, description := range err {
			commons.RED.Print("  " + k + ":\n")

			for _, line := range strings.Split(description, "\n") {
				if line != "" {
					fmt.Println("    " + line)
				}
			}

			fmt.Println() // separate next paragraph
		}
	}
}
