package goscholar

import (
	"testing"
	"fmt"
	"github.com/docopt/docopt-go"
)

// $ goscholar search --query "deep learning" --author "y bengio" --after 2015 --num 100 --start 20
func TestSearchQuery(t *testing.T) {
	args := []string{"goscholar", "search", "--query", "deep learning", "--author", "y bengio", "--after", "2015", "--num", "100", "--start", "20"}
	expected := "https://scholar.google.co.jp/scholar?hl=en&q=deep+learning+author:\"y+bengio\"&as_ylo=2015&as_yhi=&num=100&start=20"

	if err := CheckQuery(SearchQuery, args, expected); err != nil {
		t.Error(err)
	}
}

// $ goscholar find 8108748482885444188
func TestFindQuery(t *testing.T) {
	args := []string{"goscholar", "find","8108748482885444188"}
	expected := "https://scholar.google.co.jp/scholar?hl=en&cluster=8108748482885444188&num=1"

	if err := CheckQuery(FindQuery, args, expected); err != nil {
		t.Error(err)
	}
}

// $ goscholar cite 8108748482885444188 --after 2012 --num 40 --start 20
func TestCiteQuery(t *testing.T) {
	args := []string{"goscholar", "cite", "8108748482885444188", "--after", "2012", "--num", "40", "--start", "20"}
	expected := "https://scholar.google.co.jp/scholar?hl=en&cites=8108748482885444188&as_ylo=2012&as_yhi=&num=40&start=20"

	if err := CheckQuery(CiteQuery, args, expected); err != nil {
		t.Error(err)
	}
}

func TestCitePopQuery(t *testing.T) {
	// set params
	info := "XOJff8gPiHAJ"

	// exec NewQuery()
	url, _ := CitePopUpQuery(info)

	// check the results and expected results
	expected := "https://scholar.google.co.jp/scholar?q=info:XOJff8gPiHAJ:scholar.google.com/&output=cite&scirp=0&hl=en"
	if url != expected {
		t.Error(fmt.Sprintf("\nExpected: %v\n     URL: %v", expected, url))
	}
}

type FailQueryTestError struct {
	url string
	expected string
}

func (e FailQueryTestError) Error() string {
	return fmt.Sprintf("\nExpected: %v\n     URL: %v", e.expected, e.url)
}

func CheckQuery(query func(map[string]interface{}) (string, error),args []string, expected string) error {
	arguments, _ := docopt.Parse(USAGE, args[1:], true, VERSION, false)
	url, _ := query(arguments)

	if url != expected {
		return FailQueryTestError{url, expected}
	}

	return nil

}