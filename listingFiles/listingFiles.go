package listingFiles

import (
	"io/ioutil"
	"log"
)

func ListFiles(directory string) ([]string, error){
	files, err  := ioutil.ReadDir(directory)
	if err != nil{
		log.Fatal(err)
		return nil, err
	}

	var fileNames []string
	for _, file := range files{
		fileNames = append(fileNames,file.Name())
	} 
	return fileNames, nil
}
