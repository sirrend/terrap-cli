package terraform_utils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	ver "github.com/hashicorp/terraform/version"
	"log"
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
@returns: execPath - the path to the terraform executable
@         tv - the terraform version
*/

func InstallTf() (execPath string, tv string) {
	if IsTerraformInstalled() {
		//log.Println("terraform is installed, using the installed executer.")

		var tVersion terraformVersion
		path, err := exec.LookPath("terraform")
		if err != nil {
			log.Fatal(err)
		}
		j, err := exec.Command("terraform", "version", "--json").Output()

		if err != nil {
			log.Fatalf("error getting terraform version: %s", err)
		}

		err = json.Unmarshal(j, &tVersion)

		if err != nil {
			log.Fatalf("error unmarshalling terraform version: %s", err)
		}

		return path, tVersion.Terraform_version

	} else {
		log.Println("installing terraform version ", ver.Version)

		// set installer details
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion(ver.Version)),
		}

		// install terraform in context of the given directory
		execPath, err := installer.Install(context.Background())
		if err != nil {
			log.Fatalf("error installing Terraform: %s", err)
		}

		return execPath, ver.Version
	}
}

/*
@brief: NewTerraformExecuter creates a new terraform executer
@
@params: dir - the directory to run terraform in
@        execPath - the path to the Terraform executable
@returns: *tfexec.Terraform - the terraform executer
*/

func NewTerraformExecuter(dir string, execPath string) *tfexec.Terraform {
	dir, _ = filepath.Abs(dir)
	tf, err := tfexec.NewTerraform(dir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	return tf
}

/*
@brief: terraformInit performs the `terraform init` command on the given folder
@
@params: dir - the folder to initialize
*/

func TerraformInit(dir string) (string, string, error) {
	execPath, terraformVersion := InstallTf()

	tf := NewTerraformExecuter(dir, execPath)

	// initialize terraform in the given folder
	err := tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		if strings.Contains(err.Error(), "net/http") {
			//log.Panicf("Terraform init error: Terrap experienced networking issues, please make sure you have a stable internet connection.")
			err = errors.New("terrap experienced networking issues, please make sure you have a stable internet connection")
			return "", "", err
		} else {
			//log.Panicf("Terraform init error: %s", err.Error())
			return "", "", err
		}
	}

	return execPath, terraformVersion, nil
}
