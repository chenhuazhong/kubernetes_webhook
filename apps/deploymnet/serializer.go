package deploymnet
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	v1 "k8s.io/api/admission/v1"
//	"k8s.io/apimachinery/pkg/apis/meta/v1beta1"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/klog/v2"
//	"net/http"
//)
//
//func serve(w http.ResponseWriter, r *http.Request) {
//	var body []byte
//	if r.Body != nil {
//		if data, err := ioutil.ReadAll(r.Body); err == nil {
//			body = data
//		}
//	}
//
//	// verify the content type is accurate
//	contentType := r.Header.Get("Content-Type")
//	if contentType != "application/json" {
//		fmt.Errorf("contentType=%s, expect application/json", contentType)
//		return
//	}
//
//	klog.V(2).Info(fmt.Sprintf("handling request: %s", body))
//
//	deserializer := codecs.UniversalDeserializer()
//	obj, gvk, err := deserializer.Decode(body, nil, nil)
//	if err != nil {
//		msg := fmt.Sprintf("Request could not be decoded: %v", err)
//		klog.Error(msg)
//		http.Error(w, msg, http.StatusBadRequest)
//		return
//	}
//
//	var responseObj runtime.Object
//	switch *gvk {
//	case v1beta1.SchemeGroupVersion.WithKind("AdmissionReview"):
//		requestedAdmissionReview, ok := obj.(*v1beta1.AdmissionReview)
//		if !ok {
//			klog.Errorf("Expected v1beta1.AdmissionReview but got: %T", obj)
//			return
//		}
//		responseAdmissionReview := &v1beta1.AdmissionReview{}
//		responseAdmissionReview.SetGroupVersionKind(*gvk)
//		responseAdmissionReview.Response = admit.v1beta1(*requestedAdmissionReview)
//		responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID
//		responseObj = responseAdmissionReview
//	case v1.SchemeGroupVersion.WithKind("AdmissionReview"):
//		requestedAdmissionReview, ok := obj.(*v1.AdmissionReview)
//		if !ok {
//			klog.Errorf("Expected v1.AdmissionReview but got: %T", obj)
//			return
//		}
//		responseAdmissionReview := &v1.AdmissionReview{}
//		responseAdmissionReview.SetGroupVersionKind(*gvk)
//		responseAdmissionReview.Response = admit.v1(*requestedAdmissionReview)
//		responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID
//		responseObj = responseAdmissionReview
//	default:
//		msg := fmt.Sprintf("Unsupported group version kind: %v", gvk)
//		klog.Error(msg)
//		http.Error(w, msg, http.StatusBadRequest)
//		return
//	}
//
//	klog.V(2).Info(fmt.Sprintf("sending response: %v", responseObj))
//	respBytes, err := json.Marshal(responseObj)
//	if err != nil {
//		klog.Error(err)
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	if _, err := w.Write(respBytes); err != nil {
//		klog.Error(err)
//	}
//}