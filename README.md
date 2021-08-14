## 动态准入控制器 webhook

### 要求
* go 1.15
* kubernetes 1.20.*

### 简介
k8s提供两个动态准入控制器入口，分别是 mutating webhook 和validating webhook。  
该项目将围绕这两个控制器实现。

### 注入配置 MutatingWebhookConfiguration
### 校验配置 ValidatingWebhookConfiguration





### 参考
https://github.com/kubernetes/kubernetes/blob/release-1.21/test/images/agnhost/webhook/main.go
https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#webhook-configuration
https://github.com/kubernetes/kubernetes/blob/release-1.21/test/images/agnhost/webhook/main.go
https://www.qikqiak.com/post/k8s-admission-webhook/
https://github.com/morvencao/kube-mutating-webhook-tutorial
https://github.com/cnych/admission-webhook-example