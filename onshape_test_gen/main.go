package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/iancoleman/strcase"
	"github.com/onshape-public/go-client/onshape"
)

func main() {
	_, _ = onshape.NewAPIClientFromEnv()
	writeServiceTests()
}

type APIServiceTest struct {
	APIName  string
	Requests []string
}

type OpenAPISpecification struct {
	Paths MapSlice `json:"paths"`
}

func writeServiceTests() {
	testFiles := parseSpecificationRequests()

	for p, tf := range testFiles {
		path := "./../onshape_test/" + p

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("New tests created for", path)
			writeNewTestFile(path, *tf)
		} else {
			writeExistingTestFile(path, *tf)
		}
	}
}

func parseSpecificationRequests() map[string]*APIServiceTest {
	file, err := os.ReadFile("./../openapi.json")

	if err != nil {
		panic(err)
	}

	var spec OpenAPISpecification
	if err := json.Unmarshal(file, &spec); err != nil {
		panic(fmt.Sprint("Could not unmarshal", err))
	}

	testFiles := make(map[string]*APIServiceTest)

	for _, p := range spec.Paths {
		for _, q := range p.Value.(MapSlice) {
			q := q.Value.(MapSlice)

			apiName := Get(q, "tags").([]interface{})[0].(string)
			fileName := "api_" + strcase.ToSnake(apiName) + "_test.go"

			if testFiles[fileName] == nil {
				testFiles[fileName] = &APIServiceTest{
					APIName:  apiName,
					Requests: make([]string, 0),
				}
			}

			testFiles[fileName].Requests = append(testFiles[fileName].Requests, "Api"+strcase.ToCamel(Get(q, "operationId").(string))+"Request")
		}
	}

	return testFiles
}

func writeExistingTestFile(name string, st APIServiceTest) {
	st = filterForExistingTests(name, st)

	if len(st.Requests) > 0 {
		fmt.Println("Additional tests added to", name)
		writeMustacheTemplateFile(name, st, "api_template_test_extra.mustache")
	}
}

func writeNewTestFile(name string, st APIServiceTest) {
	writeMustacheTemplateFile(name, st, "api_template_test.mustache")
}

func writeMustacheTemplateFile(name string, st APIServiceTest, mus string) {
	fil, err := os.ReadFile("./" + mus)

	if err != nil {
		panic(err)
	}

	ab, err := mustache.Render(string(fil), st)

	if err != nil {
		panic(err)
	}

	c, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	c.WriteString(ab)
	c.Close()
}

func filterForExistingTests(name string, st APIServiceTest) APIServiceTest {
	data, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	sd := string(data)
	requests := make([]string, 0)

	for _, r := range st.Requests {
		if !strings.Contains(sd, "onshape."+r) && !strings.Contains(sd, st.APIName+"Api."+strings.TrimSuffix(strings.TrimPrefix(r, "Api"), "Request")) {
			requests = append(requests, r)
		}
	}

	return APIServiceTest{
		APIName:  st.APIName,
		Requests: requests,
	}
}
