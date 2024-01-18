package movie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title string `json:"title"`
	Year int `json:"year"`
	Rating float64 `json:"rating"`
}

var movies []Movie

func MoivesHandler(w http.ResponseWriter, req *http.Request){
	method := req.Method
	if method=="POST"{
		body, err := io.ReadAll(req.Body)
		if err!= nil{
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		u:=Movie{}
		err = json.Unmarshal(body,&u)
		if err!= nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error: %s", err)
			return
		}
		movies = append(movies, u)


		r,err := json.Marshal(u)
		w.Write(r)
		return
	}
	if method =="GET"{
		r,err:= json.Marshal(movies)
		if err != nil{
			fmt.Fprintf(w,"err: %s",err)
			return
		}
		w.Header().Set("content-type","application/json; charset=utf-8")

		w.Write(r)
	}else{
	fmt.Fprintf(w, " %s something wrong", method)
	}

}