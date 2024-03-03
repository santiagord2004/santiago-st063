package uploader

import(
	"fmt"
	"net/http"
	"os"
	"io"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseMultipartForm(10 << 20)
	if err!= nil{
		http.Error(w, "No se pudo parsear el formulario", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No se pudo obtener el archivo del formulario", http.StatusBadRequest)
		return
	}
	defer file.Close()

	newFile, err := os.Create(handler.Filename)
	if err !=nil{
		http.Error(w, "No se pudo crear el archivo", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(newFile, file)
	if err != nil{
		http.Error(w, "No se pudo copiar el archivo", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Archivo '%s' subido con exito", handler.Filename)
}

func Start() {
	http.HandleFunc("/upload", UploadHandler)
	fmt.Println("Uploader: Servidor iniciado en http://localhost:8080/upload")
	http.ListenAndServe(":8080", nil)
}