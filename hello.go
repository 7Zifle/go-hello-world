package main

func Hello(language, name string) string {
	greetingPrefix := ""
	switch language {
	case "english":
		greetingPrefix = "Hello, "
	case "french":
		greetingPrefix = "Bonjour, "
	case "spanish":
		greetingPrefix = "Hola, "
	}
	if name == "" {
		name = "buddy"
	}
	return greetingPrefix + name + "!"
}

func main() {

}
