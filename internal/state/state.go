package state

import (
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"io"
	"log"
	"os"
	"sync"
)

var lock sync.Mutex

/*
@brief: Save saves a representation of v to the file at path.
@
@params: path string - the path to the file to save to
@		 v interface{} - the object to save
@
@returns: error - the error if any
*/

func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	r, err := utils.Marshal(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r) //write marshaled json to file

	return err
}

/*
@brief: Load loads the file at path into v.
@		Use os.IsNotExist() to see if the returned error is due
@		to the file being missing.Save saves a representation of v to the file at path.
@
@params: path string - the path to the file to load from
@		 v interface{} - the object to load into
@
@returns: error - the error if any
*/

func Load(path string, v *workspace.Workspace) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	return utils.Unmarshal(f, v)
}
