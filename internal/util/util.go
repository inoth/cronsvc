package util

func First[T any](defaultArg T, args []T) T {
	if len(args) > 0 {
		defaultArg = args[0]
	}
	return defaultArg
}
