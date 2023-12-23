package main

import (
	"context"
	"flag"
	"github.com/IlyaZayats/faculus/internal/db"
	"github.com/IlyaZayats/faculus/internal/handlers"
	"github.com/IlyaZayats/faculus/internal/repository"
	"github.com/IlyaZayats/faculus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		dbUrl    string
		listen   string
		logLevel string
	)
	//postgres://dynus:dynus@localhost:5555/dynus
	//postgres://dynus:dynus@postgres:5432/dynus
	flag.StringVar(&dbUrl, "db", "postgres://faculus:faculus@postgres:5432/faculus", "database connection url")
	flag.StringVar(&listen, "listen", ":8080", "server listen interface")
	flag.StringVar(&logLevel, "log-level", "error", "log level: panic, fatal, error, warning, info, debug, trace")

	flag.Parse()

	ctx := context.Background()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panicf("unable to get log level: %v", err)
	}
	logrus.SetLevel(level)

	dbc, err := db.NewPostgresPool(dbUrl)
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	facultyRepo, err := repository.NewPostgresFacultyRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build faculty repo: %v", err)
	}

	groupRepo, err := repository.NewPostgresGroupRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build group repo: %v", err)
	}

	studentRepo, err := repository.NewPostgresStudentRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build student repo: %v", err)
	}

	userRepo, err := repository.NewPostgresUserRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	facultyService, err := services.NewFacultyService(facultyRepo)
	if err != nil {
		logrus.Panicf("unable build faculty service: %v", err)
	}

	groupService, err := services.NewGroupService(groupRepo)
	if err != nil {
		logrus.Panicf("unable build group service: %v", err)
	}

	studentService, err := services.NewStudentService(studentRepo)
	if err != nil {
		logrus.Panicf("unable build student service: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	g := gin.New()

	_, err = handlers.NewFacultyHandlers(g, facultyService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	_, err = handlers.NewGroupHandlers(g, groupService)
	if err != nil {
		logrus.Panicf("unable build group handlers: %v", err)
	}

	_, err = handlers.NewStudentHandlers(g, studentService)
	if err != nil {
		logrus.Panicf("unable build student handlers: %v", err)
	}

	_, err = handlers.NewUserHandlers(g, userService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			}
		}
	}(childCtx)

	<-doneC

}
