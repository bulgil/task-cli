package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/bulgil/task-cli/internal/routes"
	"github.com/bulgil/task-cli/internal/storage"
)

func main() {
	strg := storage.NewStorage("./tasks.json")
	router := routes.NewRouter(strg)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		router.Route(input)
	}
}
