package executor

type Option func(opt *option)

type option struct {
	CollectorCount int `toml:"collector_count"`
	ReceiverCount  int `toml:"receiver_count"`
	ExecuteCount   int `toml:"execute_count"`
}
