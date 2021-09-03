package gcloud

var (
	// Config exposes GCloudConfig
	Config = &GCloudConfig{}
)

// GCloudConfig wraps the gcloud config command
type GCloudConfig struct {
}

// Project provides the output of gcloud config get-value project
func (config *GCloudConfig) Project() (string, error) {
	return output("config", "get-value", "project")
}
