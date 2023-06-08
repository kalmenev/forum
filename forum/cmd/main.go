package main

import (
	"log"
	"net/http"

	"forum/config"
	"forum/endpoint"
	"forum/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle(config.MAIN_PAGE, middleware.MesureRequestTime(http.HandlerFunc(endpoint.HomePageHandler)))
	mux.Handle(config.LOGIN_PAGE, middleware.MesureRequestTime(http.HandlerFunc(endpoint.SignInHandler)))
	mux.Handle(config.SIGN_UP_PAGE, middleware.MesureRequestTime(middleware.CheckSingUpForm(http.HandlerFunc(endpoint.SignUpHandler))))
	mux.Handle(config.ALL_POST_PAGE, middleware.MesureRequestTime(http.HandlerFunc(endpoint.PostPageHandler)))
	mux.Handle(config.CREATE_POST_PAGE, middleware.MesureRequestTime(http.HandlerFunc(endpoint.CreatePostHandler)))




	srv := http.Server{
		Addr:    config.ADDR,
		Handler: mux,
	}


	log.Println(
		srv.ListenAndServe(),
	)
}
