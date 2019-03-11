# go-helm-rest
  
### Dep
helm, helm json

### API-Doc
`docker run -it --rm -p 80:80 -v $(PWD)/openapi.yaml:/usr/share/nginx/html/openapi.yaml -e SPEC_URL=openapi.yaml redocly/redoc`
  
### Minio
`helm install --set accessKey=aclab,secretKey=csie1226,persistence.size=2G,service.type=NodePort,resources.limits.memory=512Mi stable/minio`