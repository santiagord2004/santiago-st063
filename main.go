package main

import (
	"fmt"
	"main/listingFiles"
	"main/uploader"
	"os"
)

func main() {
	directory :="./"

	if len(os.Args) > 1 && os.Args[1] == "upload" {
        // Iniciar el servicio de carga de archivos
        uploader.Start()
        return
    }

	files, err := listingFiles.ListFiles(directory)
	if err != nil {
		fmt.Println("No se encontraron archivos: ", err)
		return
	}

	fmt.Println("archivos encontrados: ")
	for _, fileNames := range files {
		fmt.Println(fileNames)
	}
}