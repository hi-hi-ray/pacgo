package main

import "fmt"

type Labirinto struct {
	largura int
	altura  int
	mapa    []string
}

var labirinto Labirinto

func inicializarLabirinto() {
	labirinto = Labirinto{
		largura: 20,
		altura:  10,
		mapa: []string{
			"####################",
			"#                 F#",
			"#                  #",
			"#                  #",
			"#                  #",
			"#                  #",
			"#                  #",
			"#                  #",
			"#G                 #",
			"####################",
		},
	}
}

func desenhaTela() {
	for _, linha := range labirinto.mapa {
		fmt.Println(linha)
	}
}

func main() {
	// Inicializar terminal
	inicializa()
	defer finaliza() // executa apenas no fim do programa

	// Inicializar labirinto
	inicializarLabirinto()

	// Loop principal
	for {
		// Desenha tela
		desenhaTela()

		// Processa entrada do jogador

		// Processa movimento dos fantasmas

		// Processa colisões

		// Dorme

		break
	}
}
