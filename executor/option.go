package executor

type Option func(opt *option)

type option struct {
	MaxTask int `toml:"max_task"`
}
