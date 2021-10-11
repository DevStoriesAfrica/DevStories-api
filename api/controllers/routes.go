package controllers

import (
	"DevStories/api/middlewares"
	"net/http"
)

func (server *Server) InitializeRoutes() {
	//ToDo: Add home endpoint to direct users to homepage-GitHub README or API documentation
	server.Router.HandleFunc("/user/login", middlewares.JsonMiddleware(server.LoginUser)).Methods(http.MethodPost)
	server.Router.HandleFunc("/user/register", middlewares.JsonMiddleware(server.CreateUser)).Methods(http.MethodPost)

	server.Router.HandleFunc("/user/{id}", middlewares.JsonMiddleware(server.GetUser)).Methods(http.MethodGet)
	server.Router.HandleFunc("/user/{id}", middlewares.AuthorizationMiddleware(middlewares.JsonMiddleware(server.UpdateUser))).Methods(http.MethodPut)
	server.Router.HandleFunc("/user/{id}", middlewares.AuthorizationMiddleware(middlewares.JsonMiddleware(server.DeleteUser))).Methods(http.MethodDelete)
}
