package apps

import (
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"kubernetes_webhook/utils"
	"net/http"
)
//codecs := serializer.NewCodecFactory(runtimeScheme)
type  AdmissionReviewHeadler struct {
	codecs serializer.CodecFactory
	runtimeScheme *runtime.Scheme
	Deserializer runtime.Decoder
	request *http.Request
	Resource utils.APIResource
	Adm_obj v1.AdmissionReview

}

func (adm *AdmissionReviewHeadler) LoadAdmissionReview() {
	adm_json_data, err := ioutil.ReadAll(adm.request.Body)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	adm.Adm_obj = v1.AdmissionReview{}
	_, _, decode_err := adm.Deserializer.Decode(adm_json_data, nil, &adm.Adm_obj)
	if decode_err != nil{
		panic(err)
	}
}

func (adm *AdmissionReviewHeadler) Load(resou runtime.Object) {
	_, _, err := adm.Deserializer.Decode(adm.Adm_obj.Request.Object.Raw, nil, resou)
	if err != nil{
		panic(err)
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