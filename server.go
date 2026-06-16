package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Criamos um roteador isolado (Mux) para evitar conflito com os espaços
	mux := http.NewServeMux()

	// 2. Rota para abrir o seu arquivo HTML principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Bem-vindos ao site oficial do Minecraft _ Minecraft.html")
	})

	// 3. Configura a pasta onde estão salvos os estilos, imagens e scripts
	pastaArquivos := http.Dir("Bem-vindos ao site oficial do Minecraft _ Minecraft_files")
	fileServer := http.FileServer(pastaArquivos)

	// 4. Criamos a rota exata da pasta de arquivos
	rota := "/Bem-vindos ao site oficial do Minecraft _ Minecraft_files/"
	
	// Registramos no nosso roteador seguro 'mux'
	mux.Handle(rota, http.StripPrefix(rota, fileServer))

	// 5. Iniciamos o servidor usando o nosso roteador modificado
	fmt.Println("Servidor rodando com sucesso em http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

