package main

import (
	"fmt"
	"os"
)

type Entrada int

// Possiveis entradas do usuário
const (
	Nada = iota
	ParaCima
	ParaBaixo
	ParaEsquerda
	ParaDireita
	SairDoJogo // Tecla ESC
)

func leEntradaDoUsuario() Entrada {
	var m Entrada
	array := make([]byte, 10)

	lido, _ := os.Stdin.Read(array)

	if lido == 1 && array[0] == 0x1b {
		m = SairDoJogo
	} else if lido == 3 {
		if array[0] == 0x1b && array[1] == '[' {
			switch array[2] {
			case 'A':
				m = ParaCima
			case 'B':
				m = ParaBaixo
			case 'C':
				m = ParaDireita
			case 'D':
				m = ParaEsquerda
			}
		}
	}
	return m
}

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
	LimpaTela() // adicione esta linha
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
		m := leEntradaDoUsuario()
		if m == SairDoJogo {
			break
		}

		// Processa movimento dos fantasmas

		// Processa colisões

		// Dorme
	}
}
