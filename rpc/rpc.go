package rpc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type RPCWithCache struct {
	// Cache  *Cache
	Server *http.Server
}

func New(listenAddr string /*cacheDBURL string,*/, rpcURL string) (*RPCWithCache, error) {
	server := &http.Server{
		Addr:    listenAddr,
		Handler: http.DefaultServeMux,
	}

	// ctx := context.Background()
	// db, err := pgxpool.New(ctx, cacheDBURL)

	// if err != nil {
	// 	return nil, err
	// }

	client := &http.Client{}

	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		var rpcRequest bytes.Buffer

		_, err := io.Copy(&rpcRequest, r.Body)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to copy request body", http.StatusInternalServerError)
			return
		}

		request, err := http.NewRequest(r.Method, rpcURL+r.URL.Path, &rpcRequest)
		request.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to create request", http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(request)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to make request", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		var rpcResponse bytes.Buffer
		_, err = io.Copy(&rpcResponse, resp.Body)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to copy response body", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Write(rpcResponse.Bytes())
	})

	// cache := NewCache(database.New(db))
	return (&RPCWithCache{ /*Cache: cache,*/ Server: server}), nil
}
