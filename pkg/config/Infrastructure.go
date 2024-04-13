package config

type Infrastructure struct {
	GoogleCloud GoogleCloud `yaml:"google_cloud"`
}

type GoogleCloud struct {
	ProjectID           string   `yaml:"project_id"`
	UseCredentialsFile  bool     `yaml:"use_credentials_file"`
	CredentialsFilePath string   `yaml:"credentials_file_path"`
	Firebase            Firebase `yaml:"firebase"`
}

type Firebase struct {
	DatabaseURL string `yaml:"database_url"`
}
