package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

/*
@brief: turn stream into byte slice
@
@params: stream - io.Reader
@
@Return: []byte slice
*/

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stream)
	return buf.Bytes()
}

/*
@brief: GetFileContentAsBytes returns file content as bytes
@
@params: path - string - path to dir
@
@returns: []byte - the encoded file content
@		  error - if exist
*/

func GetFileContentAsBytes(path string) ([]byte, error) {
	if DoesExist(path) {
		jsonFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(jsonFile)

		byteValue, _ := ioutil.ReadAll(jsonFile)

		return byteValue, nil
	}

	return []byte{}, errors.New("file does not exist")
}

/*
@brief: IsDir checks if a path is a dir
@
@params: p - the path to check
@
@returns: bool - true if path is a dir, false otherwise
*/

func IsDir(p string) bool {
	fileInfo, err := os.Stat(p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fileInfo.IsDir()
}

/*
@brief: IsInitialized checks if a state file is located in the given directory
@
@params: dir string - the folder to find the state file in
@
@returns: bool - true if exist, else false
*/

func IsInitialized(dir string) bool {
	if _, err := os.Stat(path.Join(dir, ".terrap.json")); err == nil {
		return true
	} else {
		return false
	}
}

/*
@brief: DoesExist checks if a given path exists
@
@params: path string - the path to the file / folder
@
@returns: bool - true if exists, else false
*/

func DoesExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

/*
@brief: Marshal is a function that marshals the object into an
@		io.Reader.
@		By default, it uses the JSON marshaller.
@
@params: v interface{} - the object to marshal
@
@returns: io.Reader - the marshalled object
@         error - the error if any
*/

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")

	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

/*
@brief: Unmarshal is a function that unmarshals the data from the
@		reader into the specified value.
@		By default, it uses the JSON unmarshaller.
@params: r io.Reader - the reader to unmarshal from
@		 v interface{} - the object to unmarshal into
@
@returns: error - the error if any
*/

var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

/*
@brief: PrettyPrintStruct will pretty print a given struct
@
@params: v interface{} - the struct to print
@
@returns: error - the error if any
*/

func PrettyPrintStruct(i interface{}) {
	r, err := Marshal(i)

	if err == nil {
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	} else {
		log.Fatal(err)
	}
}

// GetInput
/*
@brief: GetInput will gets input from user and returns it
@
@params: message string - the message to print to the user
@
@returns: string - the input from the user
*/
func GetInput(message string) string {
	in := bufio.NewReader(os.Stdin)   // input reader
	green := color.New(color.FgGreen) // cli color

	_, _ = green.Print(message)
	input, err := in.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

// MustUnquote
/*
@brief: MustUnquote get string input and returns it unquoted
@
@params: str string - the string to unquote
@
@returns: string - the new string
*/
func MustUnquote(str string) string {
	newStr, _ := strconv.Unquote(str)

	return newStr
}

// GetCodeUntilMatchingBrace
/*
@brief:
	GetCodeUntilMatchingBrace returns the code until the next matching bracket
@params:
	input string - the code as string to work on
@returns:
	string - string with matching sets of brackets
*/
func GetCodeUntilMatchingBrace(input string) string {
	var output string

	braceCount := 0
	for index, char := range input {
		if char == '{' { // if opening bracket
			braceCount++

		} else if char == '}' { // if closing bracket
			braceCount--

			if braceCount == 0 { // if brackets match
				output += string(input[index])

				return output
			}
		}
		output += string(char)
	}

	return output
}

// GetFirstKeyInMap
/*
@brief:
	GetFirstKeyInMap finds the first key in a map[string]interface{} object
@params:
	m map[string]interface{} - the map to go over
@returns:
	string - the first key, or "" if empty
*/
func GetFirstKeyInMap(m map[string]interface{}) string {
	for key, _ := range m {
		return key
	}

	return ""
}

// GetAbsPath
/*
@brief:
	GetAbsPath finds the absolute path of an object
@params:
	path - string - the path to fins abs of
@returns:
	string - the absolute path
*/
func GetAbsPath(path string) string {
	abs, _ := filepath.Abs(path)
	return abs
}

// IsItemInSlice
/*
@brief:
	IsItemInSlice checks if a given item is inside a given slice
@params:
	item - string - the item to look for
	items - []string - the slice to look in
@returns:
	bool - true if exists else false
*/
func IsItemInSlice(item string, items []string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}

	return false
}

// StripProviderPrefix
/*
@brief:
	StripProviderPrefix cleans the provider name from its registry prefix
@params:
	provider - string - the provider to strip
@returns:
	string - the stripped provide
*/
func StripProviderPrefix(provider string) string {
	return strings.ReplaceAll(provider, "registry.terraform.io/", "")
}

// IsHiddenFolder
/*
@brief:
	IsHiddenFolder checks if a given folder is hidden
@params:
	path - string - the dir to check
@returns:
	bool - true if hidden, otherwise false
*/
func IsHiddenFolder(path string) bool {
	pathLength := len(path)
	if pathLength > 1 {
		if path[0] == '.' && path[1] != '.' {
			return true
		}
	} else if pathLength == 1 && path[0] == '.' {
		return true
	}

	return false
}

// GetTable
/*
@brief:
	GetTable returns a new initialized table
@params:
	headers - []string - the list of headers to append to the new table
@returns:
	*tablewriter.Table - the new table
*/
func GetTable(headers []string) *tablewriter.Table {
	var headerColors []tablewriter.Colors
	columnAlignment := []int{0}
	columnColors := []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}

	table := tablewriter.NewWriter(os.Stdout) // init table
	table.SetHeader(headers)                  // add headers

	// set colors
	for i := 0; i < len(headers); i++ {
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.Bold, tablewriter.BgMagentaColor})
		columnColors = append(columnColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})
	}

	columnColors = columnColors[:len(columnColors)-1] // trim last element as the first one is set before the iteration

	table.SetHeaderColor(headerColors...)
	table.SetColumnColor(columnColors...)

	// set alignment
	for i := 0; i < len(headers)-1; i++ {
		columnAlignment = append(columnAlignment, 1)
	}
	table.SetColumnAlignment(columnAlignment)

	return table
}
