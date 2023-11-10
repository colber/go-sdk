
package main

import (
	"log"
	"os"
	"bufio"
	"io"
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"time"
)

type (
	File struct {
		Id       			string  	
		Type				string 		
		Tags				[]string 	
		CreatedAt			time.Time	
		UpdatedAt			time.Time	
		Status  			string		
	}

	Chunk struct {
		Id       			string  	`reindex:"id,,pk"`
		FileId       		string  	`reindex:"file_id"`
		Seq					int64
		Content				[]byte
	}
)

func main() {

	

	chunkSize := 4*1024

	// open input file
	fileName := "./data/input.webp"
	fi, err := os.Open(fileName)
	check(err)
	defer func() {
        if err := fi.Close(); err != nil {
            check(err)
        }
    }()

	// open output file
	fo, err := os.Create("./data/output.webp")
	check(err)
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			check(err)
		}
	}()
	// make a write buffer
	// w := bufio.NewWriter(fo)
	
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(fi)
	buf := make([]byte, 0, chunkSize)
	
	// create a new file on server
	file := openFile("")

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		
		// write a chunk
        // if _, err := w.Write(buf); err != nil {
        //     panic(err)
        // }

		// write a chunk to the server file
		chunk:= &Chunk{
			FileId:file.Id,
			Seq:nChunks,
			Content:buf,
		}
		writeToFile(chunk)

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			check(err)
		}
		
		nChunks++
		nBytes += int64(len(buf))

		// process buf
		if err != nil && err != io.EOF {
			check(err)
		}
	}

	// err = w.Flush()
	check(err)

	log.Println("Bytes:", nBytes, "Chunks:", nChunks)

	

	log.Println("FileId:", file.Id)

}



func openFile(id string) (*File){
	

	url := "http://localhost:8080/media/api/v1/files/open"

	body := []byte(`{
		"Id": "",
		"Type": "",
		"Tags": null
	}`)
	// Create a HTTP post request
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	check(err)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	file := &File{}
	err = json.NewDecoder(res.Body).Decode(file)
	check(err)
	
	return file

}

func writeToFile(chunk *Chunk) (*Chunk){
	

	b, err := json.Marshal(chunk)
    check(err)

	url := "http://localhost:8080/media/api/v1/files/write"
	
	
	// Create a HTTP post request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	check(err)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	out := &Chunk{}
	err = json.NewDecoder(res.Body).Decode(out)
	check(err)
	
	return out
	
}

func getFile(id string){
	url := fmt.Sprintf("http://localhost:8080/media/api/v1/files?id=%s",id)
	response, err := http.Get(url)
	check(err)

	responseData, err := ioutil.ReadAll(response.Body)
	check(err)

	fmt.Println(string(responseData))
}

func check(err error) {
    if err != nil {
		log.Println(err)
        panic(err)
    }
}