package main

import (
	"context"
	"github.com/Schariss/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)
	//Basic implementation
	//http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
	//	log.Println("goodbye")
	//})
	//Externalize logging and handlers
	//http.HandleFunc("/", hh.ServeHTTP)
	//Using our custom servermux
	sm := http.NewServeMux()
	//sm.Handle("/hello", hh)
	//sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)
	sm.Handle("/products/", hh)
	//http.ListenAndServe(":9090", nil)
	//Instead, we use our custom servermux
	//http.ListenAndServe(":9090", sm)
	//Create our own server to handle different problems such as timeout
	server := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	// start the server
	go func() {
		l.Println("Starting server on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			//l.Fatal(err)
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//listen and serve is a blocking method that wait for response until we call the shutdown method
	l.Println("end of main")

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}


