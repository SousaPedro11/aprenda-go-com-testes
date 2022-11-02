package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaAlemao = "Hallo, "

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = prefixoOlaFrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "alemao":
		prefixo = prefixoOlaAlemao
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoDeSaudacao(idioma)

	return prefixo + nome + "!"
}

func main() {
}
