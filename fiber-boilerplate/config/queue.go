package config

type QueueConfiguration struct {
	Queue_User  string
	Queue_Pass  string
	Queue_Host  string
	Queue_Port  string
	Queue_VHost string
}

var QueueConfig *QueueConfiguration //nolint:gochecknoglobals

func LoadQueueConfig() {
	loadDefaultQueueConfig()
	ViperConfig.Unmarshal(&QueueConfig)
}

func loadDefaultQueueConfig() {
	ViperConfig.SetDefault("QUEUE_USER", "guest")
	ViperConfig.SetDefault("QUEUE_PASS", "guest")
	ViperConfig.SetDefault("QUEUE_HOST", "localhost")
	ViperConfig.SetDefault("QUEUE_PORT", "5672")
	ViperConfig.SetDefault("QUEUE_VHOST", "/")
}
