package config

const (
	ServerKey         = "server-config"
	PipelineNameKey   = "meshsync-pipeline"
	ResourcesKey      = "resources"
	GlobalResourceKey = "global"
	LocalResourceKey  = "local"
	ListenersKey      = "listeners"
	LogStreamsKey     = "log-streams"

	BrokerURL     = "broker-url"
	RequestStream = "request-stream"
	LogStream     = "log-stream"
	ExecShell     = "exec-shell"
)

type PipelineConfigs []PipelineConfig

type PipelineConfig struct {
	Name      string `json:"name" yaml:"name"`
	PublishTo string `json:"publish-to" yaml:"publish-to"`
}

type ListenerConfigs []ListenerConfig

type ListenerConfig struct {
	Name           string `json:"name" yaml:"name"`
	ConnectionName string `json:"connection-name" yaml:"connection-name"`
	PublishTo      string `json:"publish-to" yaml:"publish-to"`
	SubscribeTo    string `json:"subscribe-to" yaml:"subscribe-to"`
}
