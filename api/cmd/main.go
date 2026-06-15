package main

import (
	"flowmoney/api/internal/config"
	"flowmoney/api/internal/db"
	"flowmoney/api/internal/handlers"
	"flowmoney/api/internal/middlewares"
	"flowmoney/api/internal/repository"
	"flowmoney/api/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Config
	cfg := config.MustLoad()

	//DB
	database, err := db.ConnectDb(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	//Repositories
	uRepo := repository.NewUserRepo(database)
	cRepo := repository.NewCategoryRepo(database)
	tRepo := repository.NewTransactionRepo(database)
	bRepo := repository.NewBudgetRepo(database)

	//Services
	uService := service.NewUserService(*uRepo)
	cService := service.NewCategoryService(*cRepo)
	tService := service.NewTransactionService(*tRepo, *uRepo, *bRepo)
	bService := service.NewBudgetService(*bRepo, *tRepo)

	//Handlers
	uHandler := handlers.NewUserHandler(uService, cfg.Jwt)
	cHandler := handlers.NewCategoryHandler(cService)
	tHandler := handlers.NewTransactionHandlerImpl(tService)
	bHandler := handlers.NewBudgetHandlerImpl(bService)

	//Routes
	mux := http.NewServeMux()

	//Register
	mux.HandleFunc("POST /auth/register", uHandler.CreateUser)
	//Login
	mux.HandleFunc("POST /auth/login", uHandler.Login)

	//User routes
	mux.Handle("GET /user/{id}", middlewares.AuthMiddleware(http.HandlerFunc(uHandler.GetUserById), cfg.Jwt.JWTSecret))
	mux.Handle("PUT /update/{id}", middlewares.AuthMiddleware(http.HandlerFunc(uHandler.UpdateBalance), cfg.Jwt.JWTSecret))

	//Category routes
	mux.Handle("POST /category", middlewares.AuthMiddleware(http.HandlerFunc(cHandler.CreateCategory), cfg.Jwt.JWTSecret))
	mux.Handle("GET /category/{id}", middlewares.AuthMiddleware(http.HandlerFunc(cHandler.GetCategoryById), cfg.Jwt.JWTSecret))
	mux.Handle("GET /category/user/{id}", middlewares.AuthMiddleware(http.HandlerFunc(cHandler.GetByUserId), cfg.Jwt.JWTSecret))
	mux.Handle("PUT /category/update/{id}", middlewares.AuthMiddleware(http.HandlerFunc(cHandler.UpdateCategory), cfg.Jwt.JWTSecret))

	//Transaction routes
	mux.Handle("POST /transactions", middlewares.AuthMiddleware(http.HandlerFunc(tHandler.CreateTransaction), cfg.Jwt.JWTSecret))
	mux.Handle("GET /transactions/{id}", middlewares.AuthMiddleware(http.HandlerFunc(tHandler.GetTransactionById), cfg.Jwt.JWTSecret))
	mux.Handle("GET /transactions/user/{id}", middlewares.AuthMiddleware(http.HandlerFunc(tHandler.GetTransactionByUserId), cfg.Jwt.JWTSecret))

	//Budget routes
	mux.Handle("POST /budget", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.CreateBudget), cfg.Jwt.JWTSecret))
	mux.Handle("GET /budget/{id}", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.GetBudgetById), cfg.Jwt.JWTSecret))
	mux.Handle("GET /budget/category/{id}", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.GetBudgetByCategoryId), cfg.Jwt.JWTSecret))
	mux.Handle("GET /budget/month/{id}", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.GetByUserIdAndMonth), cfg.Jwt.JWTSecret))
	mux.Handle("PUT /budget/update/{id}", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.UpdateBudget), cfg.Jwt.JWTSecret))
	mux.Handle("DELETE /budget/{id}", middlewares.AuthMiddleware(http.HandlerFunc(bHandler.DeleteBudgetById), cfg.Jwt.JWTSecret))

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	fmt.Println("Сервер запущен!")
	http.ListenAndServe(addr, middlewares.LoggerMiddleware(mux))
}
