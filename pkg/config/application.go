package config

type (
	Application struct {
		Server Server `yaml:"server"`
		GCS    GCS    `yaml:"gcs"`
	}
	Server struct {
		Backend Backend `yaml:"backend"`
	}
	Backend struct {
		Protocol string `yaml:"protocol"`
		Domain   string `yaml:"domain"`
		Port     string `yaml:"port"`
	}
	//	Concatenating protocol + domain + port should form a valid URL

	GCS struct {
		BucketName string `yaml:"bucket_name"`
	}
)
