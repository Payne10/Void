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

