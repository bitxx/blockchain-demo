package service

import (
	"net/http"
	"os"
	"log"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/jason/blockchain-demo/blockchain/model"
	"io"
	"github.com/davecgh/go-spew/spew"
)

func Run() error{
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ",os.Getenv("ADDR"))
	s := &http.Server{
		Addr:":"+httpAddr,
		Handler:mux,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	if err := s.ListenAndServe();err !=nil{
		return err
	}
	return nil

}

func makeMuxRouter() http.Handler{
	router := mux.NewRouter()
	router.HandleFunc("/",handleGetBlockchain).Methods("GET")
	router.HandleFunc("/",handleWriteBlock).Methods("POST")
	return router
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request){
	bytes, err := json.MarshalIndent(model.Blockchain, "", " ")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request){
	var m Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err!=nil{
		respondWithJSON(w,r,http.StatusBadRequest,r.Body)
		return
	}

	defer r.Body.Close()
	newBlock, err := GenerateBlock(model.Blockchain[len(model.Blockchain)-1], m.BPM)
	if err!=nil{
		respondWithJSON(w,r,http.StatusInternalServerError,r.Body)
	}
	if isBlockValid(newBlock, model.Blockchain[len(model.Blockchain)-1]){
		newBlockchain := append(model.Blockchain,newBlock)
		spew.Dump(newBlockchain)
	}
}

func respondWithJSON(w http.ResponseWriter,r *http.Request,code int, payload interface{}){
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}