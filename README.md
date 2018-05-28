# Induscover Snap

The industrial discover or IndusCover Snap identifies and enumerates devices with the following standards:
BACnet, CoDeSys V2, EtherNet/IP [Rockwell Automation and others], Niagara Fox, Schneider Electric Modicon PLCs, Omron PLCs, PC Worx Protocol enabled PLCs, ProConOS enabled PLCs and Siemens SIMATIC S7 PLCs.

Build:
git clone ...
snapcraft

Run:

```
sudo snap install induscover
sudo snap connect induscover:network-control
```

This will publish on localhost:1883 to the MQTT topic `discover/plc/out` the data read that comes from the discovery process of sending the host to `discover/plc/in`.

If you haven't got an MQTT server running just do the below before installing induscover:

```
sudo snap install mosquitto
```

If you want to quickly test, you can use NodeRed (sudo snap install nodered) and import > clipboard the content of [nodered.flow](https://raw.githubusercontent.com/mectors/induscover/master/nodered.flow).

Alternatively, with the mosquitto client tools, in one terminal run:

```
mosquitto_sub -t discover/plc/out
```

And in another terminal:

```
mosquitto_pub -t discover/plc/in -m <host to interrogate>
```
