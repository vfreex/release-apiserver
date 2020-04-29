FROM golang:1.14 as builder

WORKDIR /go/src/app
COPY . .
RUN make build

FROM fedora:32
# We can actually use FROM scratch, but let's use Fedora for testing
LABEL name="release-apiserver" \
    maintainer="OpenShift ART <aos-team-art@redhat.com>"

COPY --from=builder /go/src/app/bin /usr/local/bin/
CMD ["/usr/local/bin/apiserver"]
