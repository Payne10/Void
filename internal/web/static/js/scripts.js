document.addEventListener('DOMContentLoaded', function() {
    // Fetch ARP scan results
    fetch('/scan')
        .then(response => response.json())
        .then(data => {
            // Create Cytoscape elements from the ARP scan data
            const elements = data.map((device, index) => ({
                data: {
                    id: device.ip,
                    label: `${device.ip}\n${device.mac}\n${device.vendor}`
                }
            }));

            // Create edges (assuming a simple star topology with a central router)
            const edges = data.map((device, index) => ({
                data: {
                    id: `edge-${index}`,
                    source: 'router', // assuming the central router has ID 'router'
                    target: device.ip
                }
            }));

            // Initialize Cytoscape with elements and edges
            var cy = cytoscape({
                container: document.getElementById('cy'),
                elements: [
                    { data: { id: 'router', label: 'Router' } }, // central router
                    ...elements,
                    ...edges
                ],
                style: [
                    {
                        selector: 'node',
                        style: {
                            'background-color': '#666',
                            'label': 'data(label)',
                            'text-valign': 'center',
                            'color': '#fff',
                            'font-size': '10px'
                        }
                    },
                    {
                        selector: 'edge',
                        style: {
                            'width': 2,
                            'line-color': '#ccc',
                            'target-arrow-color': '#ccc',
                            'target-arrow-shape': 'triangle'
                        }
                    }
                ],
                layout: {
                    name: 'circle'
                },
                userZoomingEnabled: true,
                userPanningEnabled: true,
                boxSelectionEnabled: false,
                autoungrabify: true
            });

            cy.on('tap', 'node', function(evt){
                var node = evt.target;
                console.log('Tapped node ' + node.id());
            });
        })
        .catch(error => console.error('Error fetching ARP scan data:', error));

    document.getElementById('loadPackets').addEventListener('click', function() {
        fetch('/packets')
            .then(response => response.json())
            .then(data => {
                let packetData = '';
                data.forEach(packet => {
                    packetData += packet + '\n\n';
                });
                document.getElementById('packetData').textContent = packetData;
            })
            .catch(error => console.error('Error fetching packets:', error));
    });
});

