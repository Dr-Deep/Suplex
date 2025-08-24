package web

import (
	"net/http"

	"github.com/Dr-Deep/Suplex.git/internal"
)

/*
 * wahrscheinlich hinter reverse proxy ?
 */

type SuplexMux struct {
	socket string
	mux    *http.ServeMux

	bot *internal.SuplexBot
}

func NewSuplexMux(bot *internal.SuplexBot) (*http.ServeMux, error) {
	var mux = http.NewServeMux()

	mux.HandleFunc(
		"/",
		RouteJoin,
	)
	mux.HandleFunc(
		"/callback",
		RouteCallback,
	)

	//mux.HandleFunc()
	// 404?
	//50x?

	return mux, nil
}

func (mux *SuplexMux) Launch() error {
	return http.ListenAndServe(
		mux.socket,
		mux.mux,
	)
}
