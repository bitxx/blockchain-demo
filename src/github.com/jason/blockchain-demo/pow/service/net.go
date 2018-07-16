package service

import (
	"net/http"
	"os"
	"log"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/jason/blockchain-demo/pow/model"
	"io"
	"github.com/davecgh/go-spew/spew"
)

func Run() error{
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ",httpAddr)
	s := &http.Server{
		Addr:httpAddr,
		Handler:mux,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	if err:=s.ListenAndServe();err!=nil{
		return err
	}
	return nil
}


func makeMuxRouter() http.Handler{
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/",handleGetBlockchain)
	muxRouter.HandleFunc("/",handleWriteBlock)
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter,r *http.Request){
	bytes, err := json.MarshalIndent(model.Blockchain, "", "")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func handleWriteBlock(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var m model.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m);err!=nil{
		respondWithJSON(w,r,http.StatusBadRequest,r.Body)
	}
	defer r.Body.Close()

	model.Mutex.Lock()
	newBlock := GenerateBlock(model.Blockchain[len(model.Blockchain)-1],m.BPM)
	model.Mutex.Unlock()
	if isBlockValid(newBlock,model.Blockchain[len(model.Blockchain)-1]){
		model.Blockchain = append(model.Blockchain,newBlock)
		spew.Dump(model.Blockchain)
	}
	respondWithJSON(w,r,http.StatusCreated,newBlock)

}

func respondWithJSON(w http.ResponseWriter,r *http.Request,code int,payload interface{}){
	w.Header().Set("Content-Type","application/json")
	response, err := json.MarshalIndent(payload, "", " ")
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500:Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}