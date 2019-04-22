FROM node AS build-web
WORKDIR /app
ADD ./web /app
RUN yarn && yarn build

FROM redocly/redoc AS api-doc
ADD ./redoc/openapi.yaml /usr/share/nginx/html/openapi.yaml
RUN sed -i -e "s|%SPEC_URL%|api-doc/openapi.yaml|g" /usr/share/nginx/html/index.html

FROM golang:alpine
RUN apk add --update --no-cache git bash curl nodejs ca-certificates && \
    curl -L https://storage.googleapis.com/kubernetes-helm/helm-v2.13.1-linux-amd64.tar.gz |tar xvz && \
    mv linux-amd64/helm /usr/bin/helm && \
    chmod +x /usr/bin/helm
RUN mkdir -p /root/.helm/plugins && \
    helm plugin install https://github.com/Microsoft/helm-json-output --version master
ADD ./api /app
ADD ./api/kube.config /root/.kube/config
WORKDIR /app
RUN helm init && go mod download
COPY --from=build-web /app/dist /app/dist
COPY --from=api-doc /usr/share/nginx/html /app/dist/api-doc
CMD go run main.go