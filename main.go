package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"network-health/controllers"
	device "network-health/core/entity/device"
	store "network-health/core/entity/device_store"
	"network-health/core/entity/logs"
	time_usecase "network-health/core/usecases/time"
	"network-health/infra/ipv4"
	"network-health/infra/stdout"
	"network-health/infra/web"
	"network-health/infra/web/routes"

	kingpin "github.com/alecthomas/kingpin/v2"
	"github.com/goombaio/namegenerator"
	iris "github.com/kataras/iris/v12"
	"golang.org/x/sync/errgroup"
)

var (
	port             = kingpin.Flag("port", "Server's port").Short('p').Default("8080").Envar("SERVER_PORT").Int()
	devicesAddresses = kingpin.Flag("address", "Devices addresses").Short('a').Required().Envar("DEVICES_ADDRS").Strings()
)

func main() {
	kingpin.Parse()

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logs.SetLogger(stdout.NewSTDOutLogger())

	router := routes.Router{}
	app := iris.New()
	router.Route(app)

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	var devices []*device.Device

	for _, deviceIP := range *devicesAddresses {
		ipv4Address, err := ipv4.NewIPv4Address(deviceIP)
		if err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error initializing app: invalid ip address '%s", deviceIP))
			return
		}

		name := nameGenerator.Generate()
		devices = append(devices, device.NewDevice(ipv4Address, name))
	}

	store, _ := store.NewDeviceStore(devices...)
	controller := controllers.NewController(store, &time_usecase.GoTime{})

	web.SetController(controller)

	go func() {
		if err := app.Run(iris.Addr(fmt.Sprintf(":%d", *port))); err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error on starting http listener: %s", err.Error()))
		}
	}()

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		<-gCtx.Done()

		app.Shutdown(context.Background())

		return nil
	})

	if err := g.Wait(); err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("exit reason: %s \n", err))
	}
}
