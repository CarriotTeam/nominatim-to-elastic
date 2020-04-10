package mqtt

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/app"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/services"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/utlis"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	services.Log.Printf("deviceId:%s , log: %s", strings.Split(message.Topic(), "/")[2], string(message.Payload()))
	data := utlis.DecodeToDeviceLog(message.Payload(), strings.Split(message.Topic(), "/")[2])
	if data != nil {
		app.OnNewLog(data)
	}
}

func ConnectToMQTT() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()
	server := flag.String("server", configs.Config.MQTTServer.Server, "The full url of the MQTT server to connect.")
	topic := flag.String("topic", configs.Config.MQTTServer.Topic, "Topic to subscribe to")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientId := flag.String("clientId", hostname+strconv.Itoa(time.Now().Second()), "A clientId for the connection")
	username := flag.String("username", configs.Config.MQTTServer.UserName, "A username to authenticate to the MQTT server")
	password := flag.String("password", configs.Config.MQTTServer.Password, "Password to match username")
	flag.Parse()

	connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientId).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		services.ConnectionStatus.Set(0)
		panic(token.Error())
	} else {
		services.ConnectionStatus.Set(1)
		fmt.Printf("Connected to %s\n", *server)
	}

	<-c
}
