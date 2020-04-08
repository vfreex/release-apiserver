# [Experimental] OpenShift Release APIServer

## Build

```bash
go build -o bin/apiserver cmd/apiserver/main.go
```

## Run the apiserver locally

### Directly execute

```bash
# assume etc is running at http://localhost:2379
bin/apiserver \
    --etcd-servers=http://localhost:2379 \
    --secure-port=9443 \
    --insecure-port=8080 \
    --insecure-bind-address=127.0.0.1 
    --delegated-auth=false
```

### Use apiserver-boot
```bash
apiserver-boot run local --generate=false
```