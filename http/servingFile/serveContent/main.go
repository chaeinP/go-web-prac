package main

import (
	"io"
	"os"
	"net/http"
)

func main(){
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpeg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request){
	f, err := os.Open("toby.jpeg")
	if (err != nil){
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	/*
	f.Name() 파일 이름
	fi.ModTime() 파일 최종 수정 시간
	f 실제 콘텐츠
	*/
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
