package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: repograbber <repositório1:branch1> <repositório2:branch2> ...")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	repositories := os.Args[1:]

	for _, repoBranch := range repositories {
		p := strings.ReplaceAll(repoBranch, "https://", "")
		parts := strings.Split(p, ":")
		if len(parts) != 2 {
			fmt.Printf("Erro: Formato inválido para repositório:branch: %s\n", repoBranch)
			continue
		}

		repository := "https://" + parts[0]
		branch := parts[1]

		wg.Add(1)
		go func(repository, branch string) {
			defer wg.Done()
			cmd := exec.Command("git", "clone", "-b", branch, repository)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				fmt.Printf("Erro ao clonar o repositório %s: %v\n", repository, err)
				return
			}
			fmt.Printf("Repositório %s (branch: %s) clonado com sucesso!\n", repository, branch)

		}(repository, branch)
		wg.Wait()
	}
}
