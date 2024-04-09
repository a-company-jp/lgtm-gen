package main

import (
	"lgtm-gen/pkg/fs"
	infra "lgtm-gen/svc/pkg/infra/firestore"
	"log"
	"net/http"

	"lgtm-gen/svc/pkg/application/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},

		AllowMethods: []string{
			"POST",
			"PUT",
			"DELETE",
			"GET",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "LGTM GEN",
		})
	})

	f, err := fs.NewFireStore()
	if err != nil {
		log.Fatalf("failed to connect to firestore, err: %v", err)
	}

	apiV1 := r.Group("/api/v1")
	if err := Implement(apiV1, f); err != nil {
		log.Fatalf("Failed to start server...%v", err)
		return
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server... %v", err)
		return
	}
}

func Implement(rg *gin.RouterGroup, f *fs.FireStore) error {
	lgtmHandler := handler.NewLGTMHandler(infra.NewLGTMTable(f))

	rg.Handle("POST", "/lgtm", lgtmHandler.CreateLGTM())

	return nil
}
