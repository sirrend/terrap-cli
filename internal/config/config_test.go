package config

import (
	"testing"
)

// test struct to hold all the cli tool configuration parameters
type conf struct {
	Username string
	Email    string
	Position string
	SendData bool
}

// Test if the folder creation is working
func TestFolderCreation(t *testing.T) {
	err := CreateConfigFolder()

	if err != nil {
		t.Errorf("Error creating folder: %s", err)
	} else {
		t.Logf("Created successfully")
	}
}

//Test if configuraation file saving is working
func TestSaveConfig(t *testing.T) {
	c := conf{Username: "test", Email: "test@gmail.com", Position: "test", SendData: true}

	err := SaveConfigurationFile(c)
	if err != nil {
		t.Errorf("Error saving configuration file: %s", err)
	} else {
		t.Logf("Saved successfully")
	}
}
