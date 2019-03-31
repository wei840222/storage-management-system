package config

import "os"

type Config struct {
	DeployNamespace string
	RancherApiUrl   string
	RancherApiToken string
	LonghornApiUrl  string
	PrometheusUiUrl string
	MongoUrl        string
}

func (c Config) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func New() *Config {
	c := &Config{}
	c.DeployNamespace = c.getEnv("DEPLOY_NAMESPACE", "default")
	c.RancherApiUrl = c.getEnv("RANCHER_API_URL", "https://acl.csie.ntut.edu.tw:3000/v3/project/c-srzf4:p-62gsq")
	c.RancherApiToken = c.getEnv("RANCHER_API_TOKEN", "token-58xvs:cfhz9zqh4szcvjcfvbkjsn9s5gvklqr2n8pgx94spjfl29lrrfg9pz")
	c.LonghornApiUrl = c.getEnv("LONGHORN_API_URL", "https://acl.csie.ntut.edu.tw:3000/k8s/clusters/c-srzf4/api/v1/namespaces/longhorn-system/services/http:longhorn-frontend:80/proxy/v1/volumes")
	c.PrometheusUiUrl = c.getEnv("PROMETHEUS_UI_URL", "https://acl.csie.ntut.edu.tw:3000/k8s/clusters/c-srzf4/api/v1/namespaces/prometheus/services/http:prometheus-grafana:80/proxy")
	c.MongoUrl = c.getEnv("MONGO_URL", "mongodb://localhost:27017")
	return c
}
