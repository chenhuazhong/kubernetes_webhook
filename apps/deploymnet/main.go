package deploymnet

import (
	"fmt"
	"io/ioutil"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/apps/v1"
	_ "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"net/http"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()

	// (https://github.com/kubernetes/kubernetes/issues/57982)
	defaulter = runtime.ObjectDefaulter(runtimeScheme)
)



func ImagePull(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hahahahahh")
	data, err := ioutil.ReadAll(request.Body)

	ar := v1beta1.AdmissionReview{}

	if err != nil{
		fmt.Println(err)
	} else {
		obj, goup, err := deserializer.Decode(data, nil, &ar)
		if err != nil{
			panic(err)
		} else {
			fmt.Println(obj, goup)
			de := v1.Deployment{}

			de_de, group, err := deserializer.Decode(ar.Request.Object.Raw, nil, &de)
			if err != nil{
				panic(err)
			}else{
				fmt.Println(de_de, group)
			}
		}
	}
}
