package main

import (
  "fmt"
  "log"
  "flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
  xj "github.com/basgys/goxml2json"
  "time"
  "os"
  "os/exec"
  "strings"
)
var conns = flag.Int("conns", 10, "how many conns (0 means infinite)")
var host = flag.String("host", "localhost:1883", "hostname of broker")
var clientID = flag.String("clientid", "rfid", "the mqtt clientid")
var user = flag.String("user", "", "username")
var pass = flag.String("pass", "", "password")

var topic = "discovery"
var intopic = "discover/plc/in"
var outtopic = "discover/plc/out"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func nmap(protocol string, host string) string {
  snap := os.Getenv("SNAP")
  xml, err := exec.Command(snap+"/usr/bin/nmap", "-oX", "-", "-sU", "-p", "47808", "--script", protocol, host).Output()
  if  err != nil {
    log.Fatal(err)
  }

  return string(xml)
}

var local MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
  host :=	string(msg.Payload())
  log.Println("Host:"+host)
  xmlout := nmap("BACnet-discover-enumerate", host)
  log.Println("BACnet:"+xmlout)
  xmlout = xmlout + "\n" + nmap("codesys-v2-discover", host)
  xmlout = xmlout + "\n" + nmap("enip-enumerate", host)
  xmlout = xmlout + "\n" + nmap("fox-info", host)
  xmlout = xmlout + "\n" + nmap("modicon-info", host)
  xmlout = xmlout + "\n" + nmap("omron-info", host)
  xmlout = xmlout + "\n" + nmap("pcworx-info", host)
  xmlout = xmlout + "\n" + nmap("proconos-info", host)
  xmlout = xmlout + "\n" + nmap("s7-enumerate", host)
  log.Println("All:"+xmlout)

  json, err := xj.Convert(strings.NewReader(xmlout))
  if err != nil {
      panic("That's embarrassing...")
  }
  if !Clocal.IsConnected() {
    if token := Clocal.Connect(); token.Wait() && token.Error() != nil {
      panic(token.Error())
    }
  }

  if ptoken := Clocal.Publish(outtopic, 0, false,	json.String()); ptoken.Wait() && ptoken.Error() != nil {
    panic(ptoken.Error())
  }
}

var Clocal MQTT.Client

func main() {
	flag.Parse()

  // Prepare the local MQTT connection
	opts2 := MQTT.NewClientOptions().AddBroker("tcp://"+*host)
  opts2.SetClientID(*clientID)
  opts2.SetDefaultPublishHandler(local)

  //create and start a client using the above ClientOptions
  Clocal = MQTT.NewClient(opts2)
  if token := Clocal.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }
	fmt.Println("Connected locally")
  defer Clocal.Disconnect(250)

	// Say we are ready for action
	Clocal.Publish(topic, 0, false,	intopic)

  go func() {
	  // start listening for discover/plc/in
	  if token := Clocal.Subscribe(intopic, 0, nil); token.Wait() && token.Error() != nil {
	    fmt.Println(token.Error())
	    os.Exit(1)
	  }
		fmt.Println("Waiting for messages on: "+intopic)
	}()

  // loop while waiting for commands to come in
	for {
		if (!Clocal.IsConnected()) {
			if token := Clocal.Connect(); token.Wait() && token.Error() != nil {
		    panic(token.Error())
		  }
		}
    time.Sleep(1000 * time.Millisecond)
  }
}
