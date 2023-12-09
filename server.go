// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stefchris/go-config"
)

func RunServer(callbacks map[string]Callback) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var request Request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			writeError(w, err.Error())
			return
		}

		response, err := request.handle(callbacks)
		if err != nil {
			writeError(w, err.Error())
			return
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			writeError(w, err.Error())
			return
		}

		w.Write(bytes)
	})

	addr := config.GetString("server.addr", ":8080")
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	go func() {
		log.Printf("Starting http server on %s", addr)
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done

	log.Println("Shutting down server...")
	err := srv.Shutdown(context.TODO())
	if err != nil {
		panic(err)
	}
	log.Println("Server stopped")
}
