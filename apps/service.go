package apps

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	v12 "k8s.io/api/admission/v1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"kubernetes_webhook/settings"
	"kubernetes_webhook/utils"
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
	mux.HandleFunc("/pod-mutating-sudecar", Sidecar)
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

func Serializer() runtime.Decoder{
	return serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
}


func Ident(writer http.ResponseWriter, request *http.Request, ) {
	err := recover()
	if  err != nil{
		return_redmission := v12.AdmissionReview{}
		return_redmission.Response = &v12.AdmissionResponse{
			Allowed: false,
			Result:  &metav1.Status{
				Code: 400,
				Message: fmt.Sprintf("%v", err),
			},

		}
		data_byte,_ := json.Marshal(return_redmission)
		writer.Write(data_byte)
	}
}


type SidecarContainer struct {
	Name string `json:"name"`
	Image string `json:"image"`
	ImagePullPolicy string `json:"image_pull_policy"`
	//Resources interface{}


}


func ImagePull(writer http.ResponseWriter, request *http.Request) {

	defer Ident(writer, request)

	adm_serializer := AdmissionReviewHeadler{
		Request: request,
		//runtimeScheme: runtime.NewScheme(),
		//codecs: serializer.NewCodecFactory(runtime.NewScheme()),
		Deserializer: Serializer(),
	}
	adm_serializer.LoadAdmissionReview()
	deployment := v1.Deployment{}
	adm_serializer.Load(&deployment)

	var path_list []utils.ApiResourceJsonPath
	print(1)
	s := deployment.Spec.Template.Spec.Containers
	for contarner_index, contaner := range s{
		fmt.Printf("处理第%s 容器: %s\n", contarner_index, contaner.Name)

		path_list = append(path_list, utils.ApiResourceJsonPath{
			Op: "replace",
			Path: fmt.Sprintf("/spec/template/spec/containers/%d/imagePullPolicy", contarner_index),
			Value: "IfNotPresent",
		})
	}
	var PatchTypeJSONPatch v12.PatchType = "JsonPath"
	path_byte_data,_ :=json.Marshal(path_list)
	return_redmission := v12.AdmissionReview{}
	return_redmission.Response = &v12.AdmissionResponse{
		Patch: path_byte_data,
		PatchType: &PatchTypeJSONPatch,
		UID: adm_serializer.Adm_obj.Request.UID,
		Allowed: true,
	}
	data, err := json.Marshal(return_redmission)
	if err != nil{
		panic(err)
	}
	writer.Write(data)

}

func Sidecar(writer http.ResponseWriter, request *http.Request)  {

	defer Ident(writer, request)
    fmt.Println("sidecar--------")
	adm_serializer := AdmissionReviewHeadler{
		Request: request,
		//runtimeScheme: runtime.NewScheme(),
		//codecs: serializer.NewCodecFactory(runtime.NewScheme()),
		Deserializer: Serializer(),
	}
	adm_serializer.LoadAdmissionReview()
	pod := corev1.Pod{}
	adm_serializer.Load(&pod)

	var path_list []utils.ApiResourceJsonPath
    //command_list := make([]string, 10)
    command_list := []string{"sleep", "33333"}
	sidecar_contarner_config := corev1.Container{
		Name: "busybox",
		Image: "busybox",
		ImagePullPolicy: "Always",
		Command: command_list,

	}
	contarner_list := []corev1.Container{}
	contarner_list = append(contarner_list, sidecar_contarner_config)
	//contarner_list = append(contarner_list, pod.Spec.Containers...)



	//contarner_list_byte, _ := json.Marshal(contarner_list)
	path_list = append(path_list, utils.ApiResourceJsonPath{
		Op: "add",
		Path: "/spec/containers/-",
		Value: sidecar_contarner_config,
	})
	var PatchTypeJSONPatch v12.PatchType = "JsonPath"
	path_byte_data,_ :=json.Marshal(path_list)
	return_redmission := v12.AdmissionReview{}
	return_redmission.Response = &v12.AdmissionResponse{
		Patch: path_byte_data,
		PatchType: &PatchTypeJSONPatch,
		UID: adm_serializer.Adm_obj.Request.UID,
		Allowed: true,
	}
	data, err := json.Marshal(return_redmission)
	if err != nil{
		panic(err)
	}
	writer.Write(data)

}
