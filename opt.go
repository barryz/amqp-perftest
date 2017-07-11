package main

import (
	"github.com/urfave/cli"
)

var (
	uri              string  // amqp uri
	exchange         string  // exchange name
	flag             string  // message flag
	interval         int     // sampling interval in seconds
	autoAck          bool    // auto ack
	multiAckEvery    bool    // multi ack every
	randomRoutingKey bool    // use random routing key per message
	routingKey       string  // routing key
	frameMax         int     // frame max
	ptxSize          int     // producer tx size
	ctxSize          int     // consumer tx size
	preDeclare       bool    // allow use of predeclared objects
	globalQos        int     // channel prefetch count
	qos              int     // consumer prefetch count
	consumerRate     float64 // consumer rate limit
	producerRate     float64 // producer rate limit
	messageSize      int     // message size in bytes
	exchangeType     string  // exchange type
	queueName        string  // queue name
	producerCount    int     // producer count
	consumerCount    int     // consumer count
	duration         int     // run duration in seconds(unlimited by default)
	heartbeat        int     // heartbeat interval
	pMessages        int     // producer message count
	cMessages        int     // consumer messages count
	confirm          int     // max unconfirmed publishers
)

func parseArgs() *cli.App {
	app := cli.NewApp()
	app.Name = "Q_PerfTest"
	app.Usage = "RabbitMQ PerfTest Tool"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "l, uri", Value: "", Usage: "amqp uri", Destination: &uri},
		cli.StringFlag{Name: "e, exchange", Value: "", Usage: "exchange name", Destination: &exchange},
		cli.StringFlag{Name: "f, flag", Value: "", Usage: "message flag", Destination: &flag},
		cli.IntFlag{Name: "i, interval", Value: 0, Usage: "sampling interval in seconds", Destination: &interval},
		cli.BoolFlag{Name: "a, autoack", Usage: "auto ack", Destination: &autoAck},
		cli.BoolFlag{Name: "A, multiAckEvery", Usage: "multi ack every", Destination: &multiAckEvery},
		cli.BoolFlag{Name: "K, randomRoutingKey", Usage: "use random routing key per message", Destination: &randomRoutingKey},
		cli.StringFlag{Name: "k, routingKey", Value: "", Usage: "routing key", Destination: &routingKey},
		cli.IntFlag{Name: "M, framemax", Value: 0, Usage: "frame max", Destination: &frameMax},
		cli.IntFlag{Name: "m, ptxsize", Value: 0, Usage: "producer tx size", Destination: &ptxSize},
		cli.IntFlag{Name: "n, ctxsize", Value: 0, Usage: "consumer tx size", Destination: &ctxSize},
		cli.BoolFlag{Name: "p, predeclare", Usage: "allow use of predeclared objects", Destination: &preDeclare},
		cli.IntFlag{Name: "Q, globalQos", Value: 0, Usage: "channel prefetch count", Destination: &globalQos},
		cli.IntFlag{Name: "q, qos", Value: 0, Usage: "consumer prefetch count", Destination: &qos},
		cli.Float64Flag{Name: "R, consumerRate", Value: 0.0, Usage: "consumer rate limit", Destination: &consumerRate},
		cli.Float64Flag{Name: "r, producerRate", Value: 0.0, Usage: "producer rate limit", Destination: &producerRate},
		cli.IntFlag{Name: "s, size", Value: 0, Usage: "message size in bytes", Destination: &messageSize},
		cli.StringFlag{Name: "t, type", Value: "", Usage: "exchange type", Destination: &exchangeType},
		cli.StringFlag{Name: "u, queue", Value: "", Usage: "queue name", Destination: &queueName},
		cli.IntFlag{Name: "x, producers", Value: 0, Usage: "producer count", Destination: &producerCount},
		cli.IntFlag{Name: "y, consumers", Value: 0, Usage: "consumer count", Destination: &consumerCount},
		cli.IntFlag{Name: "z, time", Value: 0, Usage: "run duration in seconds", Destination: &duration},
		cli.IntFlag{Name: "b, heartbeat", Value: 0, Usage: "heartbeat interval", Destination: &heartbeat},
		cli.IntFlag{Name: "C, pmessages", Value: 0, Usage: "producer message count", Destination: &pMessages},
		cli.IntFlag{Name: "D, cmessages", Value: 0, Usage: "consumer message count", Destination: &cMessages},
		cli.IntFlag{Name: "c, confirm", Value: 0, Usage: "max unconfirmed publishers", Destination: &confirm},
	}

	return app
}
