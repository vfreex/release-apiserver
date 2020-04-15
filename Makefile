all: gen build
gen: proto
	apiserver-boot build generated
proto:
	apiserver-boot build generated --generator protobuf
build: apiserver controller-manager
apiserver:
	go build -o bin/apiserver cmd/apiserver/main.go
controller-manager:
	go build -o bin/controller-manager cmd/manager/main.go
python-client:
	openapi-generator generate \
      -i http://localhost:8080/openapi/v2 \
      -g python \
      -c manifests/openapi-codegen/python-config.yaml \
      -o python-openshift-release
run: apiserver controller-manager
	etcd || true
	bin/apiserver --etcd-servers=http://localhost:2379 \
		--secure-port=9443 \
		--insecure-port=8080 \
		--insecure-bind-address=127.0.0.1 \
		--tls-cert-file=apiserver.local.config/certificates/apiserver.crt \
		--tls-private-key-file=apiserver.local.config/certificates/apiserver.key \
		--storage-media-type=application/vnd.kubernetes.protobuf \
		--delegated-auth \
		--authentication-skip-lookup \
	    -v=10 \
		--client-ca-file config/certificates/ca.crt \
		#--kubeconfig= \
	    #--authentication-kubeconfig= \
	    #--authorization-kubeconfig= \
	    --kubeconfig=$$HOME/.kube/config \
		--authentication-kubeconfig=$$HOME/.kube/config \
		--authorization-kubeconfig=$$HOME/.kube/config \
test:
	go test ./...
.PHONY: all gen proto build apiserver controller-manager python-client run test
