package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	traceId := ""
	if m := r.Context().Value("traceId"); m != nil {
		if value, ok := m.(string); ok {
			traceId = value
		}
	}
	w.Header().Add("traceId", traceId)
	w.Write([]byte("Welcome to golang"))
}

func traceId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := uuid.New().String()
		//通过contextWithValue 添加上traceId信息
		ctx := context.WithValue(r.Context(), "traceId", traceId)
		fmt.Printf("关联上 traceId=%v,\n", traceId)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func main() {
	welcomeHandler := http.HandlerFunc(welcome)
	http.Handle("/welcome", traceId(welcomeHandler))
	http.ListenAndServe(":9090", nil)
}

//启动程序后 执行命令 curl -I 127.0.0.1:9090/welcome
//
//运行结果
//程序自身输出
//关联上 traceId=37ecaf19-4e37-441a-b78e-85fd100a0b37,
//关联上 traceId=c208ee34-9d42-4446-a4c9-429b8a9bec67,

//curl -I 127.0.0.1:9090/welcome
//HTTP/1.1 200 OK
//Traceid: 37ecaf19-4e37-441a-b78e-85fd100a0b37
//Date: Wed, 16 Aug 2023 12:14:47 GMT
//Content-Length: 17
//Content-Type: text/plain; charset=utf-8
//
//ericyang@ericyangs-MacBook-Pro e3 % curl -I 127.0.0.1:9090/welcome
//HTTP/1.1 200 OK
//Traceid: c208ee34-9d42-4446-a4c9-429b8a9bec67
//Date: Wed, 16 Aug 2023 12:14:51 GMT
//Content-Length: 17
//Content-Type: text/plain; charset=utf-8
