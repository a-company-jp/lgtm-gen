package config

// ワンチャン使いそうなので置いておく（画像の削除など）
type (
	Service struct {
		Authentication Authentication `yaml:"authentication"`
	}
	Authentication struct {
		JwtSecret string `yaml:"jwt_secret"`
	}
)
