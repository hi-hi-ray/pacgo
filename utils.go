package main

import (
	"os"
	"os/exec"
)

// Tela é a struct que define uma tela
type Tela struct{}

var tela Tela

func inicializa() {
	rawMode := exec.Command("/bin/stty", "cbreak", "-echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()

	tela = Tela{}
}

func finaliza() {
	rawMode := exec.Command("/bin/stty", "-cbreak", "echo")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()
}
