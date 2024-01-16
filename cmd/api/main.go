package main

import (
	"api-agendamentos/adapter/postgres"
	"api-agendamentos/configs"
	"api-agendamentos/internal/delivery/http/routes"
	"html/template"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func loadTemplates(pattern string) *template.Template {
	tmpl, err := template.New("").Funcs(template.FuncMap{}).ParseGlob(pattern)
	if err != nil {
		panic(err)
	}
	return tmpl
}

func main() {
	globalConfig := configs.NewParsedConfig()
	postgresDatabase := postgres.NewPostgresDatabase(globalConfig.PostgresURL)
	connectionPG, err := postgresDatabase.OpenConnection()
	if err != nil {
		panic("Try to connect to database but return error")
	}
	defer connectionPG.Close()

	router := gin.Default()
	// // Define your backend server
	// backendURL, _ := url.Parse("http://backend-server:8080")

	// // Create a reverse proxy
	// proxy := httputil.NewSingleHostReverseProxy(backendURL)
	// // Use the ReverseProxy middleware
	// router.Use(func(c *gin.Context) {
	// 	proxy.ServeHTTP(c.Writer, c.Request)
	// 	c.Abort()
	// })
	router.Use(cors.Default())
	routes.InitRoutes(connectionPG, router)
	router.Run(":8080")
}
