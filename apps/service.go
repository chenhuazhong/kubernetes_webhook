package apps

import (
	"crypto/tls"
	"fmt"
	"kubernetes_webhook/settings"
	"net/http"
)


func Server() {
	pair, err := tls.LoadX509KeyPair("/home/huazhong/k8s/key/server-cert.pem", "/home/huazhong/k8s/key/server-key.pem")
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