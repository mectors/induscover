# Induscover Snap

The industrial discover or IndusCover Snap identifies and enumerates devices with the following standards:
BACnet, CoDeSys V2, EtherNet/IP [Rockwell Automation and others], Niagara Fox, Schneider Electric Modicon PLCs, Omron PLCs, PC Worx Protocol enabled PLCs, ProConOS enabled PLCs and Siemens SIMATIC S7 PLCs.

Build:
git clone ...
snapcraft

Run:

sudo snap install induscover
sudo snap connect induscover:network-control ubuntu-core:network-control

This will publish on localhost:1833 to the MQTT topic discover/plc/out the data read that comes from the discovery process of sending the host to discover/plc/in.

If you haven't got an MQTT server running just do
sudo snap install mosquitto

If you want to quickly test, you can use NodeRed (sudo snap install nodered) and import > clipboard the content of nodered.flow [https://raw.githubusercontent.com/mectors/induscover/master/nodered.flow]
