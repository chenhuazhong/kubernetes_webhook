package apps

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func ResouceValidate(w http.ResponseWriter, r *http.Request)  {
	defer Ident(w, r)
	fmt.Println("validate--------")
	adm_serializer := NewAdmiSsionReviewHeadler(r)
	adm_serializer.LoadAdmissionReview()
	pod := corev1.Pod{}
	adm_serializer.Load(&pod)
	for _, contarner := range pod.Spec.Containers{
		if contarner.Resources.Limits == nil || contarner.Resources.Requests == nil{
			adm_return := v1.AdmissionReview{}
			adm_return.Response = &v1.AdmissionResponse{
				Allowed: false,
				Result: &metav1.Status{
					Code: 400,
					Message: "contarner.Resources.Limits || contarner.Resources.Requests 资源限制不能为空",
				},
			}
			data_byte, err := json.Marshal(adm_return)
			if err != nil{
				panic(err)
			} else{
				_, err := w.Write(data_byte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				break

			}
		}
	}
}