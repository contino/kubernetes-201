FROM scratch
MAINTAINER Hibri Marzook <hibri.marzook@contino.io>
ADD sample-k8s-controller /sample-k8s-controller
ENTRYPOINT ["/sample-k8s-controller", "-logtostderr=true", "-v=2", "-stderrthreshold=INFO"]