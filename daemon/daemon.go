package daemon

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"toronto-phreak/config"
	"toronto-phreak/ui"
	"toronto-phreak/model"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	ui.RenderTemplate(w, "layout", model.Articles)
}

// Any non https traffic is redirected to the https equivilent.
func redirect(w http.ResponseWriter, req *http.Request) {
	h	:= strings.Split(req.Host, ":")[0]
    t	:= "https://" + h + ":" + strconv.FormatInt(config.Conf.SslPort, 10) + req.URL.Path
    if len(req.URL.RawQuery) > 0 {
        t += "?" + req.URL.RawQuery
    }
	fmt.Printf("Redirecting http to...           %s", t)
    http.Redirect(w, req, t, http.StatusTemporaryRedirect)
}

func keyHandler(s string) int {
	// TODO: Make this less messy
	helpstr	:= `Command       Function
-------       --------
help, h       This help screen.
quit, q       Quit application.`

	fmt.Printf("\033[1A")
	if s == "help" || s == "h" {
		fmt.Println(helpstr)
		return 0
	} else if s == "quit" || s == "q" {
		fmt.Println("Quitting...")
		fmt.Println("")
		return 1
	} else {
		fmt.Println("Unknown command...")
		fmt.Println("")
		return 0
	}
	return 0
}

func httpStart() error {
	var err error
	fmt.Printf("Starting HTTP on port...    %d\n", config.Conf.HttpPort)

	if config.Conf.ServeSsl && config.Conf.ForceSsl {
		err	= http.ListenAndServe(":" + strconv.FormatInt(config.Conf.HttpPort, 10), http.HandlerFunc(redirect))
	} else {
		err	= http.ListenAndServe(":" + strconv.FormatInt(config.Conf.HttpPort, 10), nil)
	}

	if err != nil {
		fmt.Printf("                         FAILED:\n%s\n", err)
		return err
	}
	return nil
}

func httpsStart() error {
	fmt.Printf("Starting HTTPS on port...   %d\n", config.Conf.SslPort)

	if config.Conf.ForceSsl {
		fmt.Println("Force SSL...                TRUE")
	}

	err	:= http.ListenAndServeTLS(":" + strconv.FormatInt(config.Conf.SslPort, 10), config.Conf.ServerCert, config.Conf.ServerKey, nil)
	if err != nil {
		fmt.Printf("                         FAILED:\n%s\n", err)
		return err
	}
	return nil
}

func Start() error {
	// Build the index page.
	http.HandleFunc("/", indexHandler)
	// Serve static stuff.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Printf("\033[2J\033[0;0H")
	fmt.Println("Starting goweb, type h for help...   ")
	fmt.Println("")

	if config.Conf.ServeHttp {
		go httpStart()	// TODO: Errorcheck this.
	}
	if config.Conf.ServeSsl {
		go httpsStart()	// TODO: Errorcheck this.
	}

	for i := 0; i != 1; {
		var key string
		fmt.Scanln(&key)
		i = keyHandler(key)
	}

	return nil
}
