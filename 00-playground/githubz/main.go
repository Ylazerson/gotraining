// B''H

/*
go mod init sandbox/githubz
go run main.go
*/

// See page 113.

// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"sandbox/githubz/github"
)

//!+template
const templ = `{{.TotalCount}} issues sdklfjsdklfj:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//!-template

//!+daysAgo
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-daysAgo

//!+exec
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {

	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
