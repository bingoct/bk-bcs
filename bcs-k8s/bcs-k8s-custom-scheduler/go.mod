module github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-custom-scheduler

go 1.14

replace (
	github.com/Tencent/bk-bcs/bcs-common => ../../../bk-bcs/bcs-common
	github.com/Tencent/bk-bcs/bcs-k8s/bcs-k8s-custom-scheduler => ./
	github.com/Tencent/bk-bcs/bcs-k8s/kubernetes => ../../bcs-k8s/kubernetes
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	github.com/go-logr/logr => github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr => github.com/go-logr/zapr v0.1.1
	go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4
	k8s.io/klog => k8s.io/klog v1.0.0
)

require (
	github.com/Tencent/bk-bcs/bcs-common v0.0.0-00010101000000-000000000000
	github.com/Tencent/bk-bcs/bcs-k8s/kubernetes v0.0.0-00010101000000-000000000000
	github.com/emicklei/go-restful v2.15.0+incompatible
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kubernetes v1.14.10
)
