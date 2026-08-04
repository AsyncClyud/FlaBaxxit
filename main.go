package main

import (
	"context"
	"flabaxxit/internal/config"
	posthandler "flabaxxit/internal/handler/post"
	userhandler "flabaxxit/internal/handler/user"
	"flabaxxit/internal/middleware"
	"flabaxxit/internal/router"
	postservice "flabaxxit/internal/service/post"
	userservice "flabaxxit/internal/service/user"
	"flabaxxit/internal/storage"
	poststorage "flabaxxit/internal/storage/post"
	userstorage "flabaxxit/internal/storage/user"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Overload(); err != nil {
		log.Println(".Env not found")
	}

	cfg := config.Load()
	ctx := context.Background()
	rdb := storage.ConnectRedis(ctx, cfg.RedisURL())
	db := storage.ConnectDataBase(ctx, os.Getenv("DATABASE_URL"))
	defer db.Close()

	postDB := poststorage.NewPostRepo(db, rdb)
	userDB := userstorage.NewUserRepo(db)
	postService := postservice.NewPostService(*postDB)
	authService := userservice.NewAuthService(*userDB, []byte(cfg.JWTSecret))
	postHandler := posthandler.NewPostHandler(*postService, *authService, *cfg)
	userHandler := userhandler.NewUserHandler(*authService, *cfg)
	middleware := middleware.NewAuthMiddleware(*authService)

	router := router.Router(*postDB, *userDB, *postHandler, *userHandler, *authService, *middleware)

	router_err := router.Run()
	if router_err != nil {
		log.Fatalf("Router error: %v", router_err)
	}

}
