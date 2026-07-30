package main

import (
	"blog/internal/config"
	posthandler "blog/internal/handler/post"
	userhandler "blog/internal/handler/user"
	"blog/internal/middleware"
	"blog/internal/router"
	postservice "blog/internal/service/post"
	userservice "blog/internal/service/user"
	"blog/internal/storage"
	poststorage "blog/internal/storage/post"
	userstorage "blog/internal/storage/user"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	if err := godotenv.Overload(); err != nil {
		log.Println(".Env not found")
	}

	cfg := config.Load()
	ctx := context.Background()
	redis_url, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := storage.ConnectRedis(ctx, redis_url)
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
