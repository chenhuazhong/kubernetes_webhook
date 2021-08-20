## 动态准入控制器 webhook

### 要求
* go 1.15
* kubernetes 1.20.*
* 开发环境和k8s集群 网络互通（开发调试时apiserver webhook模块调本地服务）
* 开发环境有kubectl 和kubeconfig （本地开发环境通过kuberconfig链接集群）
* openssl 生成秘钥


### 简介
    如下图所示k8s提供两个动态准入控制器入口，分别是 mutating webhook 和validating webhook。  
    该项目将围绕这两个控制器实现。
    这里写几个webhook示例
    1. 将pod的镜像拉取规则改为 IfNotPresent
    2. 给特定标签或者注解的pod植入sidecar
    
    
![image.png](https://i.loli.net/2021/08/21/l4q5hnMTaU2XyFg.png)

### 注入配置 MutatingWebhookConfiguration
    mutating webhook是在k8s创建资源的过程中动态的修改支援的配置信息的钩子功能
    (如把镜像拉取规则改为IfNotPresent， 植入sidecar容器配置)
### 校验配置 ValidatingWebhookConfiguration
    validating webhook 在支援创建的过程中作用是检查资源的配置是否符合我们的业务规则，
    如我们不允许 pod中没有做资源限制的配置，所有容器都要求有cpu和内存的使用限制。
    没有配置资源限制的pod不能通过检验，这个功能可以用validating webhook 实现。


### 开发环境调试
1. 创建秘钥部署证书
    ```shell script
    # 生成证书 创建 webhook配置
    # 传参数为webhook调用的ip，也是本地服务绑定的ip
    bash ./install_dev.sh 192.168.179.128
    
    ```
2. 运行本地服务
    
    ```shell script
    go run ./main.go
    ```
3. 创建一个pod 测试下
    ```shell script
    kubectl run nginx --image=nginx
    ```
### 部署
1. 直接执行install.sh即可
    ```shell script
    bash install.sh
    ```


### 参考
https://github.com/kubernetes/kubernetes/blob/release-1.21/test/images/agnhost/webhook/main.go  
https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#webhook-configuration   
https://www.qikqiak.com/post/k8s-admission-webhook/  
https://github.com/morvencao/kube-mutating-webhook-tutorial  
https://github.com/cnych/admission-webhook-example  
https://www.jianshu.com/p/1de38a3f50f3  
https://jsonpatch.com/  

### 安装k8s集群
https://github.com/chenhuazhong/localkube
