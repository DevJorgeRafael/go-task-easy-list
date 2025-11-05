package main

import (
	"context"
	"fmt"
	"go-task-easy-list/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error cargando config:", err)
	}
	log.Println("Base de datos conectada")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Servidor corriendo en http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	<-quit
	log.Println("Apagando servidor...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error al apagar servidor:", err)
	}
	
	log.Println("Servidor detenido correctamente")
}