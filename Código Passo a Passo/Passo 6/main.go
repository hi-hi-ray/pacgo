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

type PacGo struct {
	posicao Posicao
	figura  string
}

var pacgo PacGo

func criarPacGo(posicao Posicao, figura string) {
	pacgo = PacGo{
		posicao: posicao,
		figura:  "G",
	}
}

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
			"#   #####  #####   #",
			"#       #  #       #",
			"#       #          #",
			"#       #  #       #",
			"#          #       #",
			"#   #####  #####   #",
			"#G                 #",
			"####################",
		},
	}

	for linha, linhaMapa := range labirinto.mapa {
		for coluna, caractere := range linhaMapa {
			switch caractere {
			case 'G':
				{
					criarPacGo(Posicao{linha, coluna}, "G")
				}
			}
		}
	}
}

func desenhaTela() {
	LimpaTela()

	// Imprime mapa
	for _, linha := range labirinto.mapa {
		for _, char := range linha {
			switch char {
			case '#':
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	// Imprime PacGo
	MoveCursor(pacgo.posicao)
	fmt.Printf("%s", pacgo.figura)

	// Move cursor para fora do labirinto
	MoveCursor(Posicao{labirinto.altura + 2, 0})
}

func moverPacGo(m Movimento) {
	var novaLinha = pacgo.posicao.linha
	var novaColuna = pacgo.posicao.coluna

	switch m {
	case ParaCima:
		novaLinha--
		if novaLinha < 0 {
			novaLinha = labirinto.altura - 1
		}
	case ParaBaixo:
		novaLinha++
		if novaLinha >= labirinto.altura {
			novaLinha = 0
		}
	case ParaDireita:
		novaColuna++
		if novaColuna >= labirinto.largura {
			novaColuna = 0
		}
	case ParaEsquerda:
		novaColuna--
		if novaColuna < 0 {
			novaColuna = labirinto.largura - 1
		}
	}

	conteudoDoMapa := labirinto.mapa[novaLinha][novaColuna]
	if conteudoDoMapa != '#' {
		pacgo.posicao.linha = novaLinha
		pacgo.posicao.coluna = novaColuna
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
		} else {
			moverPacGo(m)
		}

		// Processa movimento dos fantasmas

		// Processa colisões

		// Dorme
	}
}
