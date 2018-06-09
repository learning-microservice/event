package logger

var Config = newConfig()

type config struct {
	Level      string
	FilePath   string
	MaxSize    int // megabytes
	MaxBackups int
	MaxAge     int // days
	Compress   bool
}

func newConfig() *config {
	return &config{
		Level:    "info",
		FilePath: "", // console log
	}
}
