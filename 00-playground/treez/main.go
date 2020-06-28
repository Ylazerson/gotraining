// B''H

/*
go mod init sandbox/treez
go install
treez https://golang.org https://google.com
*/

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// -- --------------------------------------
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

// -- --------------------------------------
func outline(url string) error {

	// -- ----------------------------------
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// -- ----------------------------------
	doc, err := html.Parse(resp.Body)

	if err != nil {
		return err
	}

	// -- ----------------------------------
	forEachNode(doc, startElement, endElement)

	return nil
}

// -- --------------------------------------
/*
forEachNode
- Calls functions pre(x) and post(x) for each node x in the tree rooted at n.
- Both functions are optional.
- pre is called before the children are visited
- post is called after the children are visited
*/
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

	// -- ----------------------------------
	if pre != nil {
		pre(n)
	}

	// -- ----------------------------------
	// Note the recursion!
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	// -- ----------------------------------
	if post != nil {
		post(n)
	}
}

// -- --------------------------------------
/*
Note how the functions indent:
- The * adverb in %*s prints a string padded with a variable number of spaces.
- The width and the string are provided by the arguments depth*4 and "".
*/
var depth int

// -- --------------------------------------
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*4, "", n.Data)
		depth++
	}
}

// -- --------------------------------------
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*4, "", n.Data)
	}
}

// -- --------------------------------------
// html notes:
/*
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
*/
