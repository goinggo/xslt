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

	http://www.goinggo.net/2013/10/using-xslt-with-go.html
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// document defines a json document of key value pairs
type document map[string]interface{}

// documents is used to unmarshal the json document produced by the xslt processing
type documents struct {
	Deals []document `json:"deals"`
}

func main() {
	// Process the xml against the stylesheet
	jsonData, err := processXslt(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("ProcessXslt: %s\n", err)
		os.Exit(1)
	}

	// Create a slice of the document
	dealDocs := documents{}
	err = json.Unmarshal(jsonData, &dealDocs)
	if err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deals: %d\n\n", len(dealDocs.Deals))

	// Display the documents
	for _, deal := range dealDocs.Deals {
		fmt.Printf("%v\n\n", deal)
		fmt.Printf("DealId: %d\n", int(deal["dealid"].(float64)))
		fmt.Printf("Title: %s\n\n\n", deal["title"].(string))
	}
}

// processXslt runs the xml data through the stylesheet to produce the json document for insertion
func processXslt(stylesheet string, xmldocument string) (jsonData []byte, err error) {
	cmd := new(exec.Cmd)
	cmd.Args = []string{"xsltproc", stylesheet, xmldocument}
	cmd.Env = os.Environ()
	if runtime.GOOS == "darwin" {
		cmd.Path = "./xsltproc_darwin"
	} else {
		cmd.Path = "./xsltproc_linux"
	}

	// Process the xml against the stylsheet
	jsonString, err := cmd.Output()
	if err != nil {
		return jsonData, err
	}

	// Convert to bytes
	jsonData = []byte(jsonString)

	return jsonData, err
}
