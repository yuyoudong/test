package settings

type Config struct {
	LogPath string `yaml:"logPath"`
}

type ConfigContains struct {
	Config Config `yaml:"config"`
}

var ConfigInstance ConfigContains

type SwagInfo struct {
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
}

var SwagConfig SwagConf

type SwagConf struct {
	Doc SwagInfo `yaml:"doc"`
}

var MQConf MQConfig

type MQConfig struct {
	NSQ NSQConfig `yaml:"nsq"`
}

//NSQConfig  configuration info
type NSQConfig struct {
	NSQLookupdHost string `yaml:"nsqlookupdHost"`
	NSQLookupdPort string `yaml:"nsqlookupdPort"`
	NSQDHost       string `yaml:"nsqdHost"`
	NSQDPort       string `yaml:"nsqdPort"`
	NSQDHTTPPort   string `yaml:"nsqdHttpPort"`
}
