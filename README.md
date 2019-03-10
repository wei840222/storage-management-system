# go-helm-rest
  
### Minio
`helm install --set accessKey=aclab,secretKey=csie1226,persistence.size=2G,service.type=NodePort,resources.limits.memory=512Mi --debug stable/minio`