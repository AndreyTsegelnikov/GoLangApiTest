package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

)

var (
	appname     = "backend-api-test"
	version     = "1.0.0"
	publicAddr  = "0.0.0.0:8080"
	privateAddr = "0.0.0.0:8081"
)

func init() {
	fmt.Println("App version:", version)
	fmt.Println("App name:", appname)
	fmt.Println(`PUBLIC Started at:`, `http://`+publicAddr)
	fmt.Println(`PRIVATE Started at:`, `http://`+privateAddr)
	fmt.Println("Started:", time.Now())
}

func main() {
	env.EnvInit()
	env.InitConfig()
	services.InitService()

	application := app.NewApp(publicAddr, privateAddr, appname, version)

	application.GetRouter().Use(handler.Count)
	api := application.GetRouter().Group(appname + "/api/:ver")

	api.GET(`/launch`, handler.LaunchTest)
	api.POST(`/launch`, handler.LaunchTest)
	api.POST(`/webhook`, handler.LaunchTestPost)
	api.GET(`/another`, handler.LaunchTestAnother)

	go application.ServeHTTP()
	go test.StartTest(time.Duration(env.Config.Tests.LaunchPeriodInMinutes))
	go test.StartSheduleTest(env.Config.Tests.WorkingHoursStart, env.Config.Tests.WorkingHoursStop, time.Duration(env.Config.Tests.LaunchPeriodInMinutes))

	ossig := make(chan os.Signal, 1)

	signal.Notify(ossig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-ossig

	log.Println("signal", sig.String(), "stop signal from os")
}
