package main

import (
	"fmt"
	articleHandler "gate-way/pkg/article/controllers"
	authHandler "gate-way/pkg/auth/controllers"
	"gate-way/pkg/pb"
	"gate-way/pkg/utils"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()

	// Auth Service
	clientConn, err := grpc.Dial(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect to Lotto service:", err)
	}
	authRoute := e.Group("/auth")
	aService := pb.NewAuthClient(clientConn)
	cAuth := authHandler.New(aService)
	utils.AuthService = aService
	authRoute.POST("/register", cAuth.Register)
	authRoute.POST("/login", cAuth.Login)

	// Article Service
	articleRoute := e.Group("/article")
	cArticle := articleHandler.New()
	articleRoute.GET("", cArticle.GetArticle)

	e.Logger.Fatal(e.Start(":1323"))
}
