package internal

import (
    "fmt"
    "log"
    "net"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/layers"
)

func StartLLDPScan() {
    // Get a list of all network interfaces
    interfaces, err := net.Interfaces()
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over each interface and start capturing LLDP packets
    for _, iface := range interfaces {
        // Skip loopback interfaces and interfaces that are down
        if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
            continue
        }

        go captureLLDPOnInterface(iface.Name)
    }

    // Prevent the main function from exiting
    select {}
}

func captureLLDPOnInterface(interfaceName string) {
    // Open device for packet capture
    handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
    if err != nil {
        log.Printf("Error opening interface %s: %v", interfaceName, err)
        return
    }
    defer handle.Close()

    // Set filter to capture only LLDP packets
    err = handle.SetBPFFilter("ether proto 0x88cc")
    if err != nil {
        log.Printf("Error setting BPF filter on interface %s: %v", interfaceName, err)
        return
    }

    log.Printf("Started capturing on interface %s", interfaceName)

    // Start packet capture loop
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        processLLDPPacket(packet)
    }
}

func processLLDPPacket(packet gopacket.Packet) {
    lldpLayer := packet.Layer(layers.LayerTypeLinkLayerDiscovery)
    if lldpLayer != nil {
        lldp, _ := lldpLayer.(*layers.LinkLayerDiscovery)
        fmt.Printf("LLDP packet received from: %s\n", lldp.ChassisID.String())
        for _, tlv := range lldp.Values {
            switch tlv.Type {
            case layers.LLDPTLVChassisID:
                fmt.Printf("Chassis ID: %s\n", tlv.Value)
            case layers.LLDPTLVPortID:
                fmt.Printf("Port ID: %s\n", tlv.Value)
            case layers.LLDPTLVSystemName:
                fmt.Printf("System Name: %s\n", tlv.Value)
            case layers.LLDPTLVPortDescription:
                fmt.Printf("Port Description: %s\n", tlv.Value)
            case layers.LLDPTLVSystemDescription:
                fmt.Printf("System Description: %s\n", tlv.Value)
            case layers.LLDPTLVSystemCapabilities:
                fmt.Printf("System Capabilities: %s\n", tlv.Value)
            }
        }
    }
}

