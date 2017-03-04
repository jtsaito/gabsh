package main

/* Command line tool for retrieving particular values. Wraps the gabs lib.
 *
 * Required parameter: name of json file
 * Optional parameter: list of json path keys
 *
 * Output: newline separated list of values found at each json path put in.
 *
 * Example call:
 * go run main.go -file=~/.aws/cli/cache/sandbox--arn_aws_iam.json Credentials.SecretAccessKey Credentials.SessionToken Credentials.AccessKeyId
 */

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
)

func main() {
	fileName := flag.String("file", "", "name of file to parse")
	flag.Parse()
	jsonPaths := flag.Args()

	results, err := readJSONPaths(*fileName, jsonPaths)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

func readJSONPaths(fileName string, jsonPaths []string) (results []string, err error) {
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

	for _, p := range jsonPaths {
		value, ok := jsonParsed.Path(p).Data().(string)
		if !ok {
			err = fmt.Errorf("could parse JSON path:  %s", p)
			return
		}
		results = append(results, value)
	}

	return
}
