const createMarkers = (locations, map) => {
    locations.forEach(function (loc) {
        L.marker([loc.Lat, loc.Lng]).addTo(map).bindPopup(loc.Name);
    });
}

const addNewMarkerListener = (map) => {
    map.on('click', function (e) {
        var lat = e.latlng.lat;
        var lng = e.latlng.lng;
        var name = prompt("Enter name for this location:");
        var newCenter = map.getCenter();
        var zoom = map.getZoom();

        if (name) {
            htmx.ajax('POST', '/add-location', {
                values: {
                    name: name,
                    lat: lat,
                    lng: lng,
                    newCenterLat: newCenter.lat,
                    newCenterLng: newCenter.lng,
                    zoom: zoom
                }
            });
        }
    });
}

const initializeMap = (center, locations, zoom) => {
    var map = L.map('map').setView([center.Lat, center.Lng], zoom);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; OpenStreetMap contributors'
    }).addTo(map);

    createMarkers(locations, map);
    addNewMarkerListener(map);
}
