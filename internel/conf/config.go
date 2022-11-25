package conf

type Config struct {
	AppName string `mapstructure:"app_name"  validate:"required"`
	AppDesc string `mapstructure:"app_desc"  validate:"required"`
	Port    int    `mapstructure:"port"      validate:"required"`
	RunMode string `mapstructure:"run_mode"  validate:"required"`
}

func (s Config) Valid(m map[string]string) {
	return
}
