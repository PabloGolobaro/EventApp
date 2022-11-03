package httputil

import (
	"log"
	"net"
	"net/http"
)

func RedirectToHTTPS(domain, tlsPort string) {
	httpSrv := http.Server{
		Addr: domain + ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := r.URL
			u.Host = net.JoinHostPort(domain, tlsPort)
			u.Scheme = "https"
			log.Println(u.String())
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}),
	}
	log.Println(httpSrv.ListenAndServe())
}
