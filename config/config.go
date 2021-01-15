package config

type Server struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
}
