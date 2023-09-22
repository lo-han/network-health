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
	store "network-health/core/entity/device_list"
	"network-health/core/entity/logs"
	"network-health/infra/icmp"
	"network-health/infra/stdout"
	"network-health/infra/web"
	"network-health/infra/web/routes"

	kingpin "github.com/alecthomas/kingpin/v2"
	"github.com/goombaio/namegenerator"
	iris "github.com/kataras/iris/v12"
	"golang.org/x/sync/errgroup"
)

var (
	port      = kingpin.Flag("port", "Server's port").Short('p').Default("8080").Envar("SERVER_PORT").Int()
	devicesIP = kingpin.Flag("ip", "Devices IP").Short('a').Required().Envar("DEVICES_IP").Strings()
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

	for _, deviceIP := range *devicesIP {
		ipv4Address, err := icmp.NewIPv4Address(deviceIP)
		if err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error initializing app: invalid ip address '%s", deviceIP))
			return
		}

		name := nameGenerator.Generate()
		devices = append(devices, device.NewDevice(ipv4Address, name))
	}

	store := store.NewDeviceStore(len(devices), devices...)
	controller := controllers.NewController(store)

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
