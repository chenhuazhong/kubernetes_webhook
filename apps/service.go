package apps

import (
	"crypto/tls"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"kubernetes_webhook/settings"
	"net/http"
)

type APIServeMux struct {
	http.ServeMux
}

func (mux *APIServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		fmt.Println("url:", r.RequestURI)
	}()
	mux.ServeMux.ServeHTTP(w, r)
}

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
	mux.HandleFunc("/huazhongwebhook", ImagePull)
	//mux.HandleFunc("/validate", whsvr.serve)
	server.Handler = mux
	fmt.Println("start service------")
	starterr := server.ListenAndServeTLS("", "")

	if starterr != nil {
		panic(starterr)
	}
}


//
//var (
//	runtimeScheme = runtime.NewScheme()
//	codecs        = serializer.NewCodecFactory(runtimeScheme)
//	deserializer  = codecs.UniversalDeserializer()
//
//	// (https://github.com/kubernetes/kubernetes/issues/57982)
//	defaulter = runtime.ObjectDefaulter(runtimeScheme)
//)

func seri() runtime.Decoder{
	return serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
}

func ImagePull(writer http.ResponseWriter, request *http.Request) {
	adm_serializer := AdmissionReviewHeadler{
		request: request,
		//runtimeScheme: runtime.NewScheme(),
		//codecs: serializer.NewCodecFactory(runtime.NewScheme()),
		Deserializer: seri(),

	}
	adm_serializer.LoadAdmissionReview()
	deployment := v1.Deployment{}
	adm_serializer.Load(&deployment)
	fmt.Println(deployment)
}