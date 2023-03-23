package providers

import (
	"context"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/zclconf/go-cty/cty"
	"io"
	"log"
	"os"
	"path"
	"sirrend-terraform.com/terrap/internal/commons"
	terraform_utils "sirrend-terraform.com/terrap/internal/terraform-utils"
	"sirrend-terraform.com/terrap/internal/types"
	"sirrend-terraform.com/terrap/internal/utils"
	"sirrend-terraform.com/terrap/internal/workspace"
	"strings"
)

/*
@brief: CreateProviderBlockTemplate Creates the building blocks of a new hcl file
@
@returns: *hclwrite.Body - the body of the required_providers block
@		  *hclwrite.File - the main hclfile created
*/

func CreateProviderBlockTemplate() (*hclwrite.Body, *hclwrite.File) {
	hclFile := hclwrite.NewEmptyFile()

	// initialize the body of the new file object
	rootBody := hclFile.Body()

	tfBlock := rootBody.AppendNewBlock("terraform", nil)
	tfBlockBody := tfBlock.Body()
	reqProvsBlock := tfBlockBody.AppendNewBlock("required_providers", nil)
	reqProvsBlockBody := reqProvsBlock.Body()

	return reqProvsBlockBody, hclFile
}

/*
@brief: CreateTempProviderFile creates a temporary terraform providers file to download the newest version by
@
@params: u []types.UpgradeList - a list of all the packages to upgrade
@
@returns: *gabs.Container - the schema in gabs format
@		  error - if exist
*/

func CreateTempProviderFile(u []types.UpgradeList) (string, error) {
	// create new file on system
	folderName := ".temp/temp" + utils.GenerateString(3)
	err := utils.CreateDirIfNotExist(folderName, 0747)
	if err != nil {
		return "", err
	}

	tfFile, err := os.Create(folderName + "/init.tf")
	if err != nil {
		return "", err
	}

	body, hclFile := CreateProviderBlockTemplate()
	for _, values := range u {
		s := strings.Split(values.Source, "/")
		body.SetAttributeValue(s[len(s)-1],
			cty.ObjectVal(map[string]cty.Value{
				"source":  cty.StringVal(values.Source),
				"version": cty.StringVal(values.Version),
			}))
	}

	_, err = tfFile.Write(hclFile.Bytes())
	if err != nil {
		return "", err
	}

	err = tfFile.Close()
	if err != nil {
		return "", err
	}

	return folderName, nil
}

/*
@brief: GetSchemaByProvider returns the schema for the current given provider
@
@params: context workspace.Workspace - the workspace which holds all the data of the workspace
@        provider string - the provider to return the schema for
@returns: *gabs.Container - the schema in gabs format
@		  error - if exist
*/

func GetSchemaByProvider(context workspace.Workspace, provider string) (*gabs.Container, error) {
	ver, _ := GetVersionByProvider(context, provider)
	byteValue, err := utils.GetFileContentAsBytes(path.Join(commons.TERRAP_PROVIDERS_FOLDER, provider, ver))
	if err != nil {
		return gabs.New(), err
	} else {
		parsedJSON, err := gabs.ParseJSON(byteValue) // parse the json byte array as json
		if err != nil {
			return gabs.New(), err
		}

		return parsedJSON, nil
	}
}

/*
@brief: GetVersionByProvider returns the version in the terraform workspace for the current given provider
@
@params: context workspace.Workspace - the workspace which holds all the data of the workspace
@        provider string - the provider to return the version for
@returns: string - the version
@		  error - if exist
*/

func GetVersionByProvider(context workspace.Workspace, provider string) (string, error) {
	providersList := context.Providers

	for p, v := range providersList {
		if p == provider {
			return v.String(), nil
		}
	}

	return "", errors.New("provider not found")

}

/*
@brief: saveToConfigFolder saves the schemas of the providers to the config folder
@
@params: schemas - the schemas to print to the file
*/

func saveToConfigFolder(schemas interface{}) *os.File {
	// save to config folder
	r, err := utils.Marshal(schemas)
	if err != nil {
		fmt.Println(err)
	}

	u, err := os.UserHomeDir()
	configFolder := path.Join(u, ".terrap")

	// write to file
	f, _ := os.Create(path.Join(configFolder, "providers.json"))

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

/*
@brief: SaveSchemas saves the schemas of the providers of the folder to a file
@
@params: tf - terraform executer of type *tfexec.Terraform to execute the ProvidersSchema command
*/

func SaveSchemas(tf *tfexec.Terraform) (string, error) {
	// get terraform providers schema from folder
	schemas, err := tf.ProvidersSchema(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	f := saveToConfigFolder(schemas)
	return f.Name(), nil
}

/*
@brief: FindTfProviders finds the terraform providers in the given folder
@
@params: dir - the folder to find the terraform providers in
@
@returns: the terraform providers in the given folder
*/

func FindTfProviders(dir string, main *workspace.Workspace) map[string]*version.Version {
	tf := terraform_utils.NewTerraformExecuter(dir, (*main).ExecPath)

	_, providersList, err := tf.Version(context.Background(), true)

	if err != nil {
		log.Fatalf("error getting terraform providers: %s", err)
	}

	(*main).Providers = providersList

	return providersList
}
