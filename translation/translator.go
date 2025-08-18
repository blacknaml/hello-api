package translation

func Translate(word string, language string) string {
	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	case "finnish":
		return "hei"
	default:
		return ""
	}
}
