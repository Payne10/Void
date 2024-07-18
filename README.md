# Void

`Void` is a comprehensive networking tool written in Go designed to sit between your router and the rest of your internal LAN. It captures and analyzes all network traffic, creating a visual map of the network and allowing for packet manipulation, firewall management, and machine learning-based traffic analysis.

## Features

- **Network Mapping**: Automatically discovers and maps all devices on the network.
- **Packet Capturing**: Captures and duplicates all network packets in real-time.
- **Packet Manipulation**: Allows modification and filtering of network packets.
- **Visualization**: Provides a graphical representation of the network.
- **Firewall Management**: Implements machine learning-based firewall rules.
- **Protocol Analysis**: Supports deep packet inspection and protocol analysis.
- **Scalability**: Optimized for high throughput and large networks.

## Installation

### Prerequisites

- Go 1.18 or higher
- `libpcap` library (for packet capturing)

### Clone the Repository

```bash
git clone https://github.com/Payne10/Void.git
cd void
```

### Install Dependencies

```bash
go mod download
```

### Build

```bash
go build -o Void cmd/main.go
```

## Usage

### Running the Application

To start the application, simply run:

```bash
sudo ./Void
```

Note: `sudo` is required to capture network packets.

### Network Mapping

Upon starting, `Void` will automatically create a visual map of the network using ARP requests and responses. The map can be accessed via a web-based interface provided by the application.

### Packet Capturing and Manipulation

`Void` will capture all packets on the specified network interface. Packet duplication and manipulation rules can be configured in the configuration file.

### Machine Learning Integration

`Void` integrates machine learning models for traffic analysis and firewall management. Models can be trained on custom datasets and loaded into the application.

## Configuration

The configuration file `config.yaml` allows customization of various aspects of the application:

```yaml
network:
  interface: eth0
  promiscuous: true

visualization:
  enabled: true
  output: ./network_map.png

firewall:
  ml_model: ./models/firewall_model.h5

logging:
  level: info
  output: ./logs/void.log
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add new feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Create a pull request

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or suggestions, please open an issue on GitHub or contact pakauffman@proton.me.
```
