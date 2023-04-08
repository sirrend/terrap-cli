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
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type terraformVersion struct {
	Terraform_version string
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

func InstallTf() (execPath string, tv string) {
	if IsTerraformInstalled() {
		//log.Println("terraform is installed, using the installed executer.")

		var tVersion terraformVersion
		path, err := exec.LookPath("terraform")
		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, "Terrap failed with the following error: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		j, err := exec.Command("terraform", "version", "--json").Output()

		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, "Terrap failed while fetching Terraform version: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(j, &tVersion)

		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, "Terrap failed with the following error: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return path, tVersion.Terraform_version

	} else {
		_, _ = commons.YELLOW.Println("Installing terraform version:", ver.Version)

		// set installer details
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(ver.Version)),
		}

		// install terraform in context of the given directory
		execPath, err := installer.Install(context.Background())
		if err != nil {
			_, _ = commons.RED.Print(emoji.CrossMark, "Terrap failed while installing Terraform: ")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		return execPath, ver.Version
	}
}

/*
@brief: NewTerraformExecutor creates a new terraform executor
@
@params: dir - the directory to run terraform in
@        execPath - the path to the Terraform executable
@returns: *tfexec.Terraform - the terraform executor
*/

func NewTerraformExecuter(dir string, execPath string) *tfexec.Terraform {
	dir, _ = filepath.Abs(dir)
	tf, err := tfexec.NewTerraform(dir, execPath)
	if err != nil {
		_, _ = commons.RED.Print(emoji.CrossMark, "Terrap failed with the following error: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return tf
}

/*
@brief: terraformInit performs the `terraform init` command on the given folder
@
@params: dir - the folder to initialize
*/

func TerraformInit(dir string) (string, string, error) {
	execPath, terraformToolVersion := InstallTf()

	tf := NewTerraformExecuter(dir, execPath)

	// initialize terraform in the given folder
	err := tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		if strings.Contains(err.Error(), "net/http") {
			err = errors.New("Terrap experienced some networking issues, please make sure you have a stable internet connection")
			return "", "", err
		} else {
			return "", "", err
		}
	}

	return execPath, terraformToolVersion, nil
}
