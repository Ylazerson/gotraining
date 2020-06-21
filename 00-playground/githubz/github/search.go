// B''H

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {

	// -- -------------------------------------------
	q := url.QueryEscape(strings.Join(terms, " "))

	// `arky`

	resp, err := http.Get(IssuesURL + "?q=" + q)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// -- -------------------------------------------
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// -- -------------------------------------------
	return &result, nil
}

//!-
