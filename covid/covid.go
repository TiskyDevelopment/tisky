package covid

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"tisky/errs"

	"github.com/olekukonko/tablewriter"
	"github.com/valyala/fastjson"
)

func Handle() {
	/* Creates a new flag */
	covidCmd := flag.NewFlagSet("covid", flag.ExitOnError)
	/* Command line country code variable. */
	countryCode := covidCmd.String("country", "", "country code")
	/* Last days command line variable. */
	lastDays := covidCmd.Int("days", 1, "last number of days")
	/* Args of the subcommand. */
	args := os.Args[2:]
	/* Check if subcommand's args exist. */
	if len(args) < 1 {
		covidCmd.Usage()
		os.Exit(0)
	}
	/* Parses the subcommand. */
	covidCmd.Parse(args)
	/* Check if days number is between 1 and 400. */
	if *lastDays <= 0 || *lastDays > 400 {
		errs.PrRed("[ ! ] Days variable can only be between 1 and 400.")
	}
	/* Check if country code was declared. */
	if *countryCode == "" {
		errs.NotEnoughArgs()
	}
	/* Store API response. */
	body := Request(*countryCode)
	/* Check if the country code is valid. */
	if strings.Contains(body, "provide a valid") {
		errs.BadUsage("tisky covid -country us -days 9")
	}
	/* Define JSON parser and parse the response */
	var p fastjson.Parser
	x, _ := p.Parse(body)
	dataObj := x.GetObject("data").String()
	today, _ := p.Parse(dataObj)
	/* Define array of elements. */
	var data [][]string
	for x := 0; x < *lastDays; x++ {
		/* Current day. */
		todayarr := today.GetArray("timeline")[x]
		/* Table elements. */
		confirmed := fmt.Sprint(todayarr.GetInt("new_confirmed"))
		recovered := fmt.Sprint(todayarr.GetInt("new_recovered"))
		deaths := fmt.Sprint(todayarr.GetInt("new_deaths"))
		active := fmt.Sprint(todayarr.GetInt("active"))
		data = append(data, []string{confirmed, recovered, deaths, active, fmt.Sprintf("%d days ago", x)})
	}
	/* Set the output to the console. */
	table := tablewriter.NewWriter(os.Stdout)
	/* Set table's headers. */
	table.SetHeader([]string{"Confirmed", "Recovered", "Deaths", "Active", "Time"})

	/* Load table cells. */
	for _, v := range data {
		table.Append(v)
	}
	/* Send output. */
	table.Render()
}

/* Get JSON data from url. */
func Request(code string) string {
	resp, err := http.Get(fmt.Sprintf("https://corona-api.com/countries/%v", code))
	if err != nil {
		errs.NoInternet()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
