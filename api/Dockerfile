FROM golang:1.12-stretch
RUN apt-get update && apt-get install -y git bash curl nodejs ca-certificates && \
    curl -L https://storage.googleapis.com/kubernetes-helm/helm-v2.13.1-linux-amd64.tar.gz |tar xvz && \
    mv linux-amd64/helm /usr/bin/helm && \
    chmod +x /usr/bin/helm
RUN mkdir -p /root/.helm/plugins && helm plugin install https://github.com/Microsoft/helm-json-output --version master
ADD . /app
ADD ./kube.config /root/.kube/config
WORKDIR /app
RUN helm init && go mod download
CMD alias node=nodejs && go run main.go