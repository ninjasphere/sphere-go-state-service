package main

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
	"github.com/juju/loggo"
	"github.com/streadway/amqp"
	"github.com/wolfeidau/punter"
)

var (
	debug       = kingpin.Flag("debug", "Enable debug mode.").Bool()
	redisURL    = kingpin.Flag("redis", "REDIS url.").Default("redis://localhost:6379").OverrideDefaultFromEnvar("REDIS_URL").String()
	rabbitmqURL = kingpin.Flag("rabbitmq", "rabbitmq url.").Default("amqp://guest:guest@localhost:5672").OverrideDefaultFromEnvar("RABBIT_URL").String()

	log = loggo.GetLogger("state-service")

	routingKey = "*.$cloud.*.channel.*.event.state"
	userRegex  = regexp.MustCompile(`^(?P<user_id>\w+).\$cloud.(?P<device_id>\w_]+).channel.(?P<channel_id>[a-zA-Z0-9-_]+).event.state$`)
)

func main() {

	kingpin.Version(Version)
	kingpin.Parse()

	// apply flags
	if *debug {
		loggo.GetLogger("").SetLogLevel(loggo.DEBUG)
	} else {
		loggo.GetLogger("").SetLogLevel(loggo.INFO)
	}

	rurl, err := url.Parse(*redisURL)
	if err != nil {
		panic(err)
	}

	ss := &stateStore{
		pool: newPool(rurl.Host),
	}

	consumer, err := punter.NewConsumer(*rabbitmqURL, "amq.topic", "topic", "stateservice", routingKey, "stateservice-consumer", ss.stateHandler)
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-sc
	log.Warningf("Got signal: %v", s)

	log.Warningf("shutting down consumer")
	if err := consumer.Shutdown(); err != nil {
		log.Errorf("error during shutdown: %s", err)
	}

}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

type stateStore struct {
	pool *redis.Pool
}

func (ss *stateStore) stateHandler(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {

		start := time.Now()

		log.Debugf(
			"got %dB delivery: [%v]",
			len(d.Body),
			d.DeliveryTag,
		)

		err := ss.savePayload(d.Body, d.RoutingKey)

		if err != nil {
			log.Errorf("failed to process payload: %s", err)
		}

		log.Debugf("Time Taken: %v", time.Since(start))

		d.Ack(false)
	}
	log.Debugf("handle: deliveries channel closed")
	done <- nil
}

// cache the state in redis using a key based on state:{user_id}:{device_id}:{channel_id}
func (ss *stateStore) savePayload(body []byte, routingKey string) error {

	params := getParams(routingKey)

	if params == nil {
		return fmt.Errorf("bad routing key - " + routingKey)
	}

	// state:123:b6b984190f:on-off
	key := fmt.Sprintf("state:%s:%s:%s", params["user_id"], params["device_id"], params["channel_id"])

	c := ss.pool.Get()

	n, err := c.Do("SET", key, string(body))

	if err != nil {
		return err
	}

	log.Debugf(spew.Sprintf("n = %v", n))

	return nil
}

func getParams(routingKey string) map[string]string {

	matches := userRegex.FindAllStringSubmatch(routingKey, -1)

	if matches == nil {
		return nil
	}

	params := make(map[string]string)

	for i, attr := range userRegex.SubexpNames() {
		if attr == "" {
			continue
		}
		params[attr] = matches[0][i]
	}

	return params
}