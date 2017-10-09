package provisioner

// RunnerData - sample data from configuration
type RunnerData struct {
	BucketName   string `mapstructure:"bucket-name" json:"bucket-name"`
	RemoteFolder string `mapstructure:"remote-folder" json:"remote-folder"`
	TempFolder   string `mapstructure:"temp-folder"`
}
