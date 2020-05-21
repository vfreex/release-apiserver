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
      -c ./openapi-config-python.yaml \
      -o python-openshift-release
run: apiserver controller-manager
	etcd || true
	bin/apiserver --etcd-servers=http://localhost:2379 \
		--secure-port=9443 \
		--insecure-port=8080 \
		--insecure-bind-address=127.0.0.1 \
		--storage-media-type=application/vnd.kubernetes.protobuf \
	    -v=7 \
		--client-ca-file ca.crt \
		--disable-admission-plugins=MutatingAdmissionWebhook \
	    --disable-admission-plugins=ValidatingAdmissionWebhook
test:
	go test ./...
.PHONY: all gen proto build apiserver controller-manager python-client run test
