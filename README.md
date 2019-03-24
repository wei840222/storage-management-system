# go-helm-rest
  
### Dep
helm, helm json

### API-Doc
`docker run -it --rm -p 3000:3000 -v $(PWD)/openapi.yaml:/usr/share/nginx/html/openapi.yaml -e SPEC_URL=openapi.yaml redocly/redoc`

### MongoDB
`docker run -it --rm -p 27017:27017 mongo`
  
### Minio
`helm install --set accessKey=aclab,secretKey=csie1226,persistence.size=2G,service.type=NodePort,resources.limits.memory=512Mi stable/minio`