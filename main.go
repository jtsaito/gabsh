package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
)

func main() {
	fileName := flag.String("name", "", "name of file to parse")
	flag.Parse()
	secr, token, id, err := readCredentials(*fileName)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(secr, token, id)
}

func readCredentials(fileName string) (secretAccessKey, SessionToken, AccessKeyID string, err error) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		return
	}

	secretAccessKey, ok := jsonParsed.Path("Credentials.SecretAccessKey").Data().(string)
	if !ok {
		err = errors.New("could find parse secretAccessKey")
		return
	}

	SessionToken, ok = jsonParsed.Path("Credentials.SessionToken").Data().(string)
	if !ok {
		err = errors.New("could find parse SessionToken")
		return
	}

	AccessKeyID, ok = jsonParsed.Path("Credentials.AccessKeyId").Data().(string)
	if !ok {
		err = errors.New("could find parse secretAccessKey")
		return
	}

	return
}

/*
for later ;)
func readCredential(jsonParsed *gabs.Container, key string) (string, error) {
	value, ok := jsonParsed.Path(key).Data().(string)
	if !ok {
		return "", errors.New("could find parse secretAccessKey")
	}
	return value, nil
}
*/
