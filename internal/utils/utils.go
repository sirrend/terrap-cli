package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/commons"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strings"
)

/*
@brief: GetProgressBar returns an initialized progress bar
@
@params: barText - string - the progress bar initial text
@
@return: progress bar
*/

func GetProgressBar(barText string) *progressbar.ProgressBar {
	return progressbar.NewOptions(100,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionShowDescriptionAtLineEnd(),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionSetDescription("[cyan]"+barText+"[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}

/*
@brief: GetVersionFromString grep version with semantic format from text
@
@params: strWithVersion - string - the string to match the version from
@
@return: the matched version
*/

func GetVersionFromString(strWithVersion string) string {
	r, _ := regexp.Compile("(\\d{1}|\\d{2}|\\d{3}|\\d{4})[.](\\d{1}|\\d{2}|\\d{3}|\\d{4})[.](\\d{1}|\\d{2}|\\d{3}|\\d{4})")
	match := r.FindString(strWithVersion)

	return match
}

/*
@brief: write interface into file
@
@params: object - interface{} - can be any object, fileName - string - the file name to be created
@
@Return: error
*/

func WriteInterfaceToFile(object interface{}, fileName string) error {
	marshal, _ := Marshal(object)
	err := os.WriteFile(fileName, StreamToByte(marshal), 0644)
	if err != nil {
		return err
	}

	return nil
}

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
@brief: PrintSlice print out the file content
@
@params: path - string - path to file
*/

func PrintSlice(path string) {
	lines, err := ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range lines {
		fmt.Printf("%v\n", value)
	}
}

/*
@brief: ReadLines read file's content
@
@params: path - string - path to file
@
@returns: []string - the lines from the file received
@		  error - if exist
*/

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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
@brief: CreateDirIfNotExist creates a dir if not present
@
@params: dir - string - path to dir
@        perm - os.FileMode - permission of the folder
@
@returns: error - if exist
*/

func CreateDirIfNotExist(dir string, perm os.FileMode) error {
	if !DoesExist(dir) {
		err := os.MkdirAll(dir, perm)
		if err != nil {
			return err
		}
	}

	return nil
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
@brief: FindFileByExtension finds the files with the given extension in the given folder
@
@params: dir string - the folder to find the wanted files in
@
@returns: []string - list of files with {extension} in the given folder
*/

func FindFilesByExtension(dir string, ext string) []string {
	extFiles := make([]string, 0)
	fmt.Println("scanning for terraform files in", dir)
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)

	} else {
		for _, file := range files {
			if file.IsDir() && file.Name()[0:1] != "." { // check if it's a directory and not a hidden directory
				FindFilesByExtension(file.Name(), ext)
			} else {
				if strings.HasSuffix(file.Name(), ext) {
					fmt.Println(file.Name())
					extFiles = append(extFiles, dir+file.Name())
				}
			}
		}
	}

	return extFiles
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
	} else {
		log.Fatal(err)
	}
}

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

/*
@brief: RemoveIndex will remove index from slice
@
@params: s - []string - a slice of strings to remove from
@		 index - int - the index to remove
@
@returns: []string - the new slice
*/

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

/*
@brief: GenerateString will generate a random string with the <n> size
@
@params: n - int - numb of chars
@
@returns: string - the new generated string
*/

func GenerateString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

/*
@brief: ProgressBarPanic prettify the progress bar panic operation
@
@params: pb - pointer to progressbar.ProgressBar type
@ 		 errMsg - error - the error message to output
*/

func ProgressBarPanic(pb *progressbar.ProgressBar, text string, errMsg error) {
	pb.Describe(text + "\n\n")
	fmt.Println(commons.RED_COLOR + "| " + strings.ReplaceAll(errMsg.Error(), "\n", "\n| ") + commons.RESET_COLOR)
	os.Exit(1)
}
