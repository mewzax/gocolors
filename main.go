package gocolors

func Blue(text string) string {
	return "\033[94m" + text + "\033[0m"
}

func Green(text string) string {
	return "\033[92m" + text + "\033[0m"
}

func Red(text string) string {
	return "\033[91m" + text + "\033[0m"
}

func Yellow(text string) string {
	return "\033[93m" + text + "\033[0m"
}

func Cyan(text string) string {
	return "\033[96m" + text + "\033[0m"
}

func White(text string) string {
	return "\033[37m" + text + "\033[0m"
}

func Magenta(text string) string {
	return "\033[95m" + text + "\033[0m"
}

func Black(text string) string {
	return "\033[90m" + text + "\033[0m"
}