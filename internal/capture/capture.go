package capture

import (
    "log"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/payne10/Void.git/internal/storage"
)

var (
    packets []gopacket.Packet
)

// StartCapture initializes packet capture on the specified network interface
func StartCapture(iface string) {
    handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        processPacket(packet)
    }
}

// processPacket processes each captured packet and stores it
func processPacket(packet gopacket.Packet) {
    packeData := packet.String()
    storage.InsertPacket(packetData)
    log.Println(packetData)
}

