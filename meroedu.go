package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/alecthomas/kingpin.v2"

	_ "github.com/meroedu/meroedu/docs"
	_attachmentHttpDelivery "github.com/meroedu/meroedu/internal/attachment/delivery/http"
	_attachmentRepo "github.com/meroedu/meroedu/internal/attachment/repository/mysql"
	_attachmentStore "github.com/meroedu/meroedu/internal/attachment/storage/filesystem"
	_attachmentUcase "github.com/meroedu/meroedu/internal/attachment/usecase"
	_categoryHttpDelivery "github.com/meroedu/meroedu/internal/category/delivery/http"
	_categoryRepo "github.com/meroedu/meroedu/internal/category/repository/mysql"
	_categoryUcase "github.com/meroedu/meroedu/internal/category/usecase"
	_contentHttpDelivery "github.com/meroedu/meroedu/internal/content/delivery/http"
	_contentRepo "github.com/meroedu/meroedu/internal/content/repository/mysql"
	_contentStore "github.com/meroedu/meroedu/internal/content/storage/filesystem"
	_contentUcase "github.com/meroedu/meroedu/internal/content/usecase"
	_courseHttpDelivery "github.com/meroedu/meroedu/internal/course/delivery/http"
	_courseHttpDeliveryMiddleware "github.com/meroedu/meroedu/internal/course/delivery/http/middleware"
	_courseRepo "github.com/meroedu/meroedu/internal/course/repository/mysql"
	_courseUcase "github.com/meroedu/meroedu/internal/course/usecase"
	_healthHttpDelivery "github.com/meroedu/meroedu/internal/health/delivery/http"
	_lessonHttpDelivery "github.com/meroedu/meroedu/internal/lesson/delivery/http"
	_lessonRepo "github.com/meroedu/meroedu/internal/lesson/repository/mysql"
	_lessonUcase "github.com/meroedu/meroedu/internal/lesson/usecase"
	_tagHttpDelivery "github.com/meroedu/meroedu/internal/tag/delivery/http"
	_tagRepo "github.com/meroedu/meroedu/internal/tag/repository/mysql"
	_tagUcase "github.com/meroedu/meroedu/internal/tag/usecase"
	datastore "github.com/meroedu/meroedu/pkg/database"

	"github.com/meroedu/meroedu/internal/config"
	"github.com/meroedu/meroedu/pkg/log"
)

var (
	configPath = kingpin.Flag("config", "Location of config.yml").Default("./config.yml").String()
)

// @title Mero Edu API
// @version 0.1
// @description Mero Edu is a software application for the administration, documentation, tracking, reporting, automation and delivery of educational courses, training programs, or learning and development programs for school.

// @contact.name Mero Edu
// @contact.url https://meroedu.com

// @license.name MIT License
// @license.url https://github.com/meroedu/meroedu/blob/master/LICENSE
// @BasePath /
func main() {

	// Parse the CLI flags and load the config
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// Load the config
	config.ReadConfig(*configPath)
	db, err := datastore.NewDB()
	if err != nil {
		log.Error(err)
	}
	defer func() {
		log.Info("Closing database connection")
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}()

	e := echo.New()

	// Init Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	middL := _courseHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// healthcheck
	_healthHttpDelivery.NewHealthHandler(e)

	// contents
	contentRepository := _contentRepo.Init(db)
	contentStorage, err := _contentStore.Init()
	if err != nil {
		log.Fatalf("Error initializing content storage: %v", err)
	}
	contentUseCase := _contentUcase.NewContentUseCase(contentRepository, contentStorage, timeoutContext)
	_contentHttpDelivery.NewContentHandler(e, contentUseCase)

	// tags
	tagRepository := _tagRepo.Init(db)
	_tagHttpDelivery.NewTagHandler(e, _tagUcase.NewTagUseCase(tagRepository, timeoutContext))

	// Categories
	categoryRepository := _categoryRepo.Init(db)
	_categoryHttpDelivery.NewCategoryHandler(e, _categoryUcase.NewCategoryUseCase(categoryRepository, timeoutContext))

	// Attachment
	attachmentRepository := _attachmentRepo.Init(db)
	attachmentStorage, err := _attachmentStore.Init()
	if err != nil {
		log.Fatalf("Error initializing attachment storage: %v", err)
	}
	attachmentUseCase := _attachmentUcase.NewAttachmentUseCase(attachmentRepository, attachmentStorage, timeoutContext)
	_attachmentHttpDelivery.NewAttachmentHandler(e, attachmentUseCase)

	// Lessons
	lessonRepository := _lessonRepo.Init(db)
	lessonUseCase := _lessonUcase.NewLessonUseCase(lessonRepository, contentUseCase, timeoutContext)
	_lessonHttpDelivery.NewLessonHandler(e, lessonUseCase)

	// Courses
	courseRepository := _courseRepo.Init(db)
	_courseHttpDelivery.NewCourseHandler(e, _courseUcase.NewCourseUseCase(courseRepository, lessonUseCase, attachmentUseCase, timeoutContext))

	// Start HTTP Server
	go func() {
		if err := e.Start(viper.GetString("server.address")); err != nil {
			log.Info("Shutting down the server")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
