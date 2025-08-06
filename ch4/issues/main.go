// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", len(result.Items))
	monthOld := []github.Issue{}
	yearOld := []github.Issue{}
	veryOld := []github.Issue{}
	now := time.Now()
	monthAgo := now.Add(-time.Hour * 24 * 30)
	yearAgo := now.Add(-time.Hour * 24 * 365)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.Compare(monthAgo) >= 0:
			monthOld = append(monthOld, *item)
		case item.CreatedAt.Compare(yearAgo) >= 0:
			yearOld = append(yearOld, *item)
		default:
			veryOld = append(veryOld, *item)
		}

	}
	fmt.Printf("%d monthOld issues:\n", len(monthOld))
	for _, item := range monthOld {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.RFC3339))
	}
	fmt.Printf("%d yearOld issues:\n", len(yearOld))
	for _, item := range yearOld {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.RFC3339))
	}
	fmt.Printf("%d veryOld issues:\n", len(veryOld))
	for _, item := range veryOld {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.Format(time.RFC3339))
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
