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
	scrapeCmd := flag.NewFlagSet("scrape", flag.ExitOnError)
	/* Command-line variables. */
	website := scrapeCmd.String("website", "pastebin.com", "only works with pastebin")
	content := scrapeCmd.String("content", "", "what to look for")
	output := scrapeCmd.String("output", "", "output the content to a file")
	colors.PrintYellow("[ ! ] Scrape doesn't support proxies yet. We highly suggest using a VPN.")
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
			towrite += fmt.Sprintf("\n======== SOURCE: %v ========\n%v\n\n", result, contentOfPaste)
		}
	}
	if *output != "" {
		file, err := os.Open(*output)
		if err != nil {
			file, err = os.Create(*output)
			if err != nil {
				errs.PrRed("[ ! ] There was an error creating that file.")
			}
		}
		file.Write([]byte(towrite))
		defer file.Close()

	}
	fmt.Println(towrite)

}

/* Fetches the arguments returned from the search. */
func Search(website string, content string) []string {
	/* Overwrites the website since it can only work with pastebin.com. */
	content = strings.Replace(content, " ", "+", -1)
	link := fmt.Sprintf("https://apis.explosionscratc.repl.co/google?q=site:%s+%s", website, content)
	resp := Request(link)
	var arr []string
	var p fastjson.Parser
	v, _ := p.Parse(resp)
	for _, m := range v.GetArray() {
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
