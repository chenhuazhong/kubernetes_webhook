package apps

import (
	"crypto/tls"
	"fmt"
	"kubernetes_webhook/settings"
	"net/http"
)


func Server() {
	pair, err := tls.LoadX509KeyPair(settings.CERT_FILE_PATH, settings.KEY_FILE_PATH)
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:      fmt.Sprintf(":%v", settings.HOST_PORT),
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{pair}},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/huazhongwebhook", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hahahahahh")
	})
	//mux.HandleFunc("/validate", whsvr.serve)
	server.Handler = mux
	starterr := server.ListenAndServeTLS("", "")

	if starterr != nil {
		panic(starterr)
	}
}