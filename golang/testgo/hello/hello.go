package hello

type language string

const (
	spanish language = "Spanish"
	french  language = "French"
)

type prefix string

const (
	englishHelloPrefix prefix = "Hello, "
	spanishHelloPrefix prefix = "Hola, "
	frenchHelloPrefix  prefix = "Bonjour, "
)

func greetingPrefix(lang language) (p prefix) {
	switch lang {
	case spanish:
		p = spanishHelloPrefix
	case french:
		p = frenchHelloPrefix
	default:
		p = englishHelloPrefix
	}
	return
}

func Hello(name string, lang language) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(lang)
	return string(prefix) + name
}
