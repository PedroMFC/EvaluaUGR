package tests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"testing"
)


func BenchmarkMux(b *testing.B) {
	postBody, _ := json.Marshal(map[string]string{

	})

   for i := 0; i < b.N; i++ {
	   http.Get("http://localhost:8081/valoraciones/AAA")
	   http.Post("http://localhost:8081/valoraciones/AAA/5", "",  bytes.NewBuffer(postBody))
   }
}

func BenchmarkEcho(b *testing.B) {
	postBody, _ := json.Marshal(map[string]string{

	})

   for i := 0; i < b.N; i++ {
	   http.Get("http://localhost:8082/valoraciones/AAA")
	   http.Post("http://localhost:8082/valoraciones/AAA/5", "",  bytes.NewBuffer(postBody))
   }
}

func BenchmarkGin(b *testing.B) {
	postBody, _ := json.Marshal(map[string]string{

	 })

	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8080/valoraciones/AAA")
		http.Post("http://localhost:8080/valoraciones/AAA/5", "",  bytes.NewBuffer(postBody))
	}
}
