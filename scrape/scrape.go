package scrape

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"tisky/colors"
	"tisky/errs"

	"github.com/valyala/fastjson"
)

func Handle() {
	/*
	  Warn about using a VPN.
	  TODO: Add proxies.
	*/
	colors.PrintYellow("[ ! ] Scrape doesn't support proxies yet. We highly suggest using a VPN.")
	/* Creates a new flag set called scrape. */
	scrapeCmd := flag.NewFlagSet("scrape", flag.ExitOnError)
	/* Command-line variables. */
	website := scrapeCmd.String("website", "pastebin.com", "only works with pastebin")
	content := scrapeCmd.String("content", "", "what to look for")
	output := scrapeCmd.String("output", "", "output the content to a file")
	/* Get subcommand's args. */
	args := os.Args[2:]
	/* Exit if not enough arguments. */
	if len(args) < 1 {
		scrapeCmd.Usage()
		os.Exit(0)
	}
	/* Parses the subcommand. */
	scrapeCmd.Parse(args)
	/* Search for the data's. */
	results := Search(*website, *content)
	/* To-write text declaration. */
	towrite := ""
	/* Print the content and output in a file if the output variable isn't empty. */
	for _, result := range results {
		contentOfPaste := ReadContent(result)
		fmt.Println(contentOfPaste)
		if *output != "" {
			towrite += fmt.Sprintf("\n======== SOURCE: %v ========\n\n%v\n", result, contentOfPaste)
		}
	}
	/* If user wants to output in a file what has scraped.*/
	if *output != "" {
		os.Remove(*output)
		file, err := os.Open(*output)
		if err != nil {
			file, err = os.Create(*output)
			if err != nil {
				errs.PrRed("[ ! ] There was an error creating that file.")
			}
		}
		/* Writes what has scraped. */
		file.Write([]byte(towrite))
		defer file.Close()

	}
	/* Outputs what has been scraped. */
	fmt.Println(towrite)

}

/* Fetches the arguments returned from the search. */
func Search(website string, content string) []string {
	/* Makes the content compatible with the google search API. */
	content = strings.Replace(content, " ", "+", -1)
	/* The API below is not mine, it is a random one from repl.it.*/
	link := fmt.Sprintf("https://apis.explosionscratc.repl.co/google?q=site:%s+%s", website, content)
	/* Read the API and outputs an array with all the links. */
	resp := Request(link)
	if strings.Contains(resp, "only") {
		errs.PrRed("[ ! ] You got ratelimited from the API for 15 minutes.")
	}

	var arr []string
	var p fastjson.Parser
	v, _ := p.Parse(resp)
	for _, m := range v.GetArray() {
		/* Gets the link from json. */
		link := string(m.GetStringBytes("link"))
		arr = append(arr, link)
		fmt.Printf("%v[ + ] Found %v %v\n", colors.Green, link, colors.White)
	}
	return arr
}

/* Get JSON data from url. */
func Request(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		errs.NoInternet()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

/* Reads paste's content. */
func ReadContent(link string) string {
	link = strings.Replace(link, "pastebin.com/", "pastebin.com/raw/", -1)
	content := Request(link)
	return content
}
