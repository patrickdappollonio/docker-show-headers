package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

var port = ":" + envdefault("PORT", "8080")

func envdefault(key, defval string) string {
	if s, found := os.LookupEnv(key); found {
		return s
	}

	return defval
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wr := tabwriter.NewWriter(w, 6, 1, 4, ' ', 0)

		fmt.Fprintln(wr, "REQUEST INFORMATION:")
		fmt.Fprintln(wr, "--------------------")
		fmt.Fprintln(wr, "URL:\t", r.URL.String())
		fmt.Fprintln(wr, "Method:\t", r.Method)
		fmt.Fprintln(wr, "Protocol:\t", r.Proto)
		fmt.Fprintln(wr, "Host:\t", r.Host)
		fmt.Fprintln(wr, "Client Address:\t", r.RemoteAddr)

		if len(r.Header) > 0 {
			fmt.Fprintln(wr, "")
			fmt.Fprintln(wr, "REQUEST HEADERS:")
			fmt.Fprintln(wr, "----------------")

			for k, v := range r.Header {
				for _, m := range v {
					fmt.Fprintln(wr, k+":\t"+m)
				}
			}
		}

		if err := wr.Flush(); err != nil {
			http.Error(w, "Error flushing output headers: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
