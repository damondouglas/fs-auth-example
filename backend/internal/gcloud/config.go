package gcloud

var (
	Config = &GCloudConfig{}
)

type GCloudConfig struct {
}

func (config *GCloudConfig) Project() (string, error) {
	return output("config", "get-value", "project")
}
