module knative.dev/eventing-autoscaler-keda

go 1.16

require (
	github.com/kedacore/keda/v2 v2.3.0
	go.uber.org/zap v1.18.1
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.24.1-0.20210712225150-5add4858c67d
	knative.dev/eventing-awssqs v0.24.1-0.20210712140124-93961c1cf002
	knative.dev/eventing-kafka v0.24.1-0.20210712161722-380c3a136f73
	knative.dev/eventing-redis v0.24.1-0.20210712215750-1737aceaa488
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210712150822-e8973c6acbf7
	knative.dev/reconciler-test v0.0.0-20210712152422-2b7893cc2eb4
)

replace (
	k8s.io/api => k8s.io/api v0.19.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.7
	k8s.io/apiserver => k8s.io/apiserver v0.19.7
	k8s.io/client-go => k8s.io/client-go v0.19.7
	k8s.io/code-generator => k8s.io/code-generator v0.19.7
)
