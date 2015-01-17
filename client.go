package worldbank

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://api.worldbank.org/" // countries/chn;bra/indicators/DPANUSIFS?date=2009Q1:2010Q3

type Series struct {
	Language   string
	Countries  []string
	Indicators []string
	MRV        string
	PerPage    string
	Start      Date
	End        Date
	Format     string
	Frequency  string
	Data       string
	Query      []string
}

type Date struct {
	Year, // 2014
	Subunit string // M08 or Q3
}

// Create a Query String
func (s *Series) Querify() {
	language := s.Language + "/"
	countries := "countries/" + strings.Join(s.Countries, ";") + "/"
	indicators := "indicators/" + strings.Join(s.Indicators, ";") + "?"
	format := "format=" + s.Format // + "&"
	date := "date=" +
		s.Start.Year + s.Start.Subunit + ":" +
		s.End.Year + s.End.Subunit // + "&"
	mrv := "MRV=" + s.MRV                   // + "&"
	perpage := "per_page=" + s.PerPage      // + "&"
	frequency := "frequency=" + s.Frequency // + "&"

	s.Query = []string{url, language, countries, indicators, strings.Join([]string{format, date, mrv, perpage, frequency}, "&")}
}

// Request Data Using Query String
func (s *Series) Request() {
	resp, err := http.Get(strings.Join(s.Query, ""))
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	s.Data = string(body)
}

// Write Data to File (XML or JSON)
func (s Series) Write(filepath string) {
	err := ioutil.WriteFile(filepath, []byte(s.Data), 0777)
	check(err)
}

func (s *Series) Reset() {
	*s = Series{}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
