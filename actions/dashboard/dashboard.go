package dashboard

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strconv"

	e "github.com/mt1976/crt/errors"
	t "github.com/mt1976/crt/language"
	support "github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
	probing "github.com/prometheus-community/pro-bing"
)

var C = config.Configuration

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	// crt.Clear()
	crt.InfoMessage(t.TxtDashboardChecking)
	p := page.New(t.TxtDashboardTitle)

	c := 0
	c++
	//p.Add("Testing Server/Service Dashboard", "", time.Now().Format("2006-01-02"))
	for i := 0; i < C.DashboardURINoEntries; i++ {
		//p.Add(C.DashboardURIName[i], "", "")
		crt.InfoMessage(fmt.Sprintf(t.TxtDashboardCheckingService, C.DashboardURIName[i]))
		result := CheckService(i)
		p.AddFieldValuePair(crt, C.DashboardURIName[i], result)
	}

	p.AddAction(t.SymActionQuit)
	p.AddAction(t.SymActionForward)
	p.AddAction(t.SymActionBack)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch nextAction {
		case t.SymActionForward:
			p.NextPage(crt)
		case t.SymActionBack:
			p.PreviousPage(crt)
		case t.SymActionQuit:
			ok = true
			return
		default:
			crt.InputError(e.ErrInvalidAction + t.SymSingleQuote + nextAction + t.SymSingleQuote)
		}
	}

}

func CheckService(i int) string {

	protocol := C.DashboardURIProtocol[i]
	host := C.DashboardURIHost[i]
	if host == "" {
		host = C.DashboardDefaultHost
	}
	if host == "" {
		panic(e.ErrDashboardNoHost)
	}
	port := C.DashboardURIPort[i]
	if port == "" {
		port = C.DashboardDefaultPort
	}
	query := C.DashboardURIQuery[i]
	operation := C.DashboardURIOperation[i]
	success := C.DashboardURISuccess[i]

	//Check Operation is a valid operation by checking if operation is in C.DashboardURIValidActions
	if !slices.Contains(C.DashboardURIValidActions, support.Upcase(operation)) {
		return e.ErrInvalidAction
	}

	if support.Upcase(operation) == "PING" {

		pinger, err := probing.NewPinger(host)
		if err != nil {
			return t.TxtStatusOffline + t.Space + support.PQuote(err.Error())
		}
		pinger.Count = 3
		err = pinger.Run() // Blocks until finished.
		if err != nil {
			return t.TxtStatusOffline + t.Space + support.PQuote(err.Error())
		}
		stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
		avgRtt := stats.AvgRtt

		return t.TxtStatusOnline + t.Space + support.PQuote(fmt.Sprintf("%v", avgRtt))
	}

	if support.Upcase(operation) == "HTTP" {
		//fmt.Printf("GET %v://%v:%v%v - %v %v\n", protocol, host, port, query, operation, success)
		var u url.URL

		u.Scheme = protocol
		u.Host = host + ":" + port
		u.Path = query

		//fmt.Println(u)

		return StatusCode(u.String(), "", success)
	}

	xx := fmt.Sprintf("%v://%v:%v%v - %v %v", protocol, host, port, query, operation, success)
	return xx
}

func StatusCode(PAGE string, AUTH string, SUCCESS string) (r string) {
	// Setup the request.
	req, err := http.NewRequest("GET", PAGE, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", AUTH)

	// Execute the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return t.TxtStatusOffline + t.Space + support.PQuote(t.TxtNoResponseFromServer)
	}

	// Close response body as required.
	defer resp.Body.Close()

	//fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode == 200 {
		return t.TxtStatusOnline + t.Space + support.PQuote(resp.Status)
	}
	//resp.StatusCode to string
	scString := strconv.Itoa(resp.StatusCode)
	if scString == SUCCESS {
		return t.TxtStatusOnline + t.Space + support.PQuote(resp.Status)
	}

	return t.TxtStatusOffline + t.Space + support.PQuote(resp.Status)
	// or fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}
