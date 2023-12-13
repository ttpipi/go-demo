package main

import (
	"context"
	"flag"
	"math/rand"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/metrics"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	metricsAddr = flag.String("metrics", "8081", "The address that will expose /metrics for Prometheus")

	logger = watermill.NewStdLogger(true, true)
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	flag.Parse()

	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	router.AddHandler(
		"metrics-example",
		"sub-topic", pubsub,
		"pub_topic", pubsub,
		handler,
	)	

	prometheusRegistry, closeMetricsServer := metrics.CreateRegistryAndServeHTTP(*metricsAddr)
	defer closeMetricsServer()

	metricsBuilder := metrics.NewPrometheusMetricsBuilder(prometheusRegistry, "", "")
	metricsBuilder.AddPrometheusRouterMetrics(router)

	pubWithMetrics, err := metricsBuilder.DecoratePublisher(pubsub)
	if err != nil {
		panic(err)
	}

	subWithMetrics, err := metricsBuilder.DecorateSubscriber(pubsub)
	if err != nil {
		panic(err)
	}

	routerClosed := make(chan struct{})
	go produceMessages(routerClosed, pubWithMetrics)
	go consumeMessages(subWithMetrics)

	_ = router.Run(context.Background())
	close(routerClosed)

}

func handler(msg *message.Message) ([]*message.Message, error) {
	numOutgoing := random.Intn(4)
	outgoing := make([]*message.Message, numOutgoing)
	
	for i := 0; i < numOutgoing; i++ {
		outgoing[i] = msg.Copy()
	}
	return outgoing, nil
}