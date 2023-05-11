package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ahmadfathan/todolist/cmd/helper"
	"github.com/ahmadfathan/todolist/internal/config"

	webhandler "github.com/ahmadfathan/todolist/internal/handler/todolist/http"
	logHelper "github.com/ahmadfathan/todolist/pkg/log"
	"github.com/ahmadfathan/todolist/pkg/middleware"
	"github.com/go-chi/chi"

	activityUsecase "github.com/ahmadfathan/todolist/internal/usecase/activity"
	todoUsecase "github.com/ahmadfathan/todolist/internal/usecase/todo"

	activityRepo "github.com/ahmadfathan/todolist/internal/repo/activity"
	todoRepo "github.com/ahmadfathan/todolist/internal/repo/todo"
)

func main() {

	appConfig := config.GetAppConfig()

	err := startApp(appConfig)

	if err != nil {
		logHelper.ErrLog.Fatalln("[main app] failed starting the app..")
		fmt.Println(err.Error())
	}
}

type Server struct {
	HTTP helper.HttpConfig
}

func startApp(appConfig config.AppConfig) error {

	// connect to DB
	dbDriver := "mysql"

	connectionString := config.GetDBConnectionString()

	db, err := sqlx.Connect(dbDriver, connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	// create activity repo
	activityRepo, err := activityRepo.New(db)

	if err != nil {
		logHelper.ErrLog.Errorln("[main app] error while initiating activity repo:", err)
		return err
	}

	// create todo repo
	todoRepo, err := todoRepo.New(db)

	if err != nil {
		logHelper.ErrLog.Errorln("[main app] error while initiating todo repo:", err)
		return err
	}

	// create activity usecase
	activityUc := activityUsecase.New(activityRepo)

	// create todo usecase
	todoUc := todoUsecase.New(todoRepo)

	// create new handler
	webHandler := webhandler.New(activityUc, todoUc)

	router := newRoutes(webHandler)

	logHelper.ErrLog.Infoln("[main app] starting the router..")

	return helper.StartServer(router, appConfig.Server.HTTP)
}

func newRoutes(webhandler *webhandler.Handler) *chi.Mux {
	router := chi.NewRouter()
	mw := middleware.NewSet()

	// register the endpoint here

	// activity
	router.Method(http.MethodGet, "/activity-groups", mw.HandlerFunc(webhandler.GetAllActivity))
	router.Method(http.MethodGet, "/activity-groups/{activity-id}", mw.HandlerFunc(webhandler.GetActivityByID))
	router.Method(http.MethodPost, "/activity-groups", mw.HandlerFunc(webhandler.CreateActivity))
	router.Method(http.MethodPatch, "/activity-groups/{activity-id}", mw.HandlerFunc(webhandler.UpdateActivity))
	router.Method(http.MethodDelete, "/activity-groups/{activity-id}", mw.HandlerFunc(webhandler.DeleteActivity))

	// todo
	router.Method(http.MethodGet, "/todo-items", mw.HandlerFunc(webhandler.GetAllTodo))
	router.Method(http.MethodGet, "/todo-items/{todo-id}", mw.HandlerFunc(webhandler.GetTodoByID))
	router.Method(http.MethodPost, "/todo-items", mw.HandlerFunc(webhandler.CreateTodo))
	router.Method(http.MethodPatch, "/todo-items/{todo-id}", mw.HandlerFunc(webhandler.UpdateTodo))
	router.Method(http.MethodDelete, "/todo-items/{todo-id}", mw.HandlerFunc(webhandler.DeleteTodo))

	return router
}
