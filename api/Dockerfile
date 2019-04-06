FROM golang:alpine
RUN apk add --update --no-cache git bash curl nodejs ca-certificates && \
    curl -L https://storage.googleapis.com/kubernetes-helm/helm-v2.13.1-linux-amd64.tar.gz |tar xvz && \
    mv linux-amd64/helm /usr/bin/helm && \
    chmod +x /usr/bin/helm
RUN mkdir -p /root/.helm/plugins && \
    helm plugin install https://github.com/Microsoft/helm-json-output --version master
ADD . /app
ADD ./kube.config /root/.kube/config
WORKDIR /app
RUN helm init && go mod download
CMD go run main.go