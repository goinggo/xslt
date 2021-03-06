// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
This program provides an sample to learn how a work pool can increase
performance and get more work done with less resources

Ardan Studios
12973 SW 112 ST, Suite 153
Miami, FL 33186
bill@ardanstudios.com

http://www.goinggo.net/2013/11/using-xslt-with-go.html
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// document defines a json document of key value pairs.
type document map[string]interface{}

func main() {
	// Process the xml against the stylesheet.
	jsonData, err := processXslt("stylesheet.xslt", "deals.xml")
	if err != nil {
		fmt.Printf("ProcessXslt: %s\n", err)
		os.Exit(1)
	}

	// An anonymous struct to unmarshal the json document
	// produced by the xslt processing.
	var documents struct {
		Deals []document `json:"deals"`
	}

	// Create a slice of the document
	if err := json.Unmarshal(jsonData, &documents); err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deals: %d\n\n", len(documents.Deals))

	// Display the documents
	for _, deal := range documents.Deals {
		fmt.Printf("DealId: %d\n", int(deal["dealid"].(float64)))
		fmt.Printf("Title: %s\n\n", deal["title"].(string))
	}
}

// processXslt runs the xml data through the stylesheet to produce the json document for insertion
func processXslt(xslFile string, xmlFile string) ([]byte, error) {
	cmd := exec.Cmd{
		Args: []string{"xsltproc", xslFile, xmlFile},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}

	// Process the xml against the stylsheet
	jsonString, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", jsonString)

	return []byte(jsonString), nil
}
