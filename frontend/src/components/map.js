import { React, useState } from 'react';
import { MapContainer, TileLayer, Marker, Popup, useMapEvents } from 'react-leaflet';
import Card from './card';


function LocationMarker({ onLocationChange }) {
    const [position, setPosition] = useState(null)
    const map = useMapEvents({
        click(e) {
            setPosition(e.latlng);
            onLocationChange(e.latlng);
        },
    })

    return position === null ? null : (
        <Marker position={position}>
            <Popup>You are here</Popup>
        </Marker>
    )
}

const Map = ({ setAnswer }) => {

    const handleLocationChange = (location) => {
        setAnswer(location);
    };

    return (
        <MapContainer center={[51.505, -0.09]} zoom={2} style={{ height: '600px', width: '100%' }}>
            <TileLayer
                attribution='&copy; <a href="https://carto.com/">carto.com</a> contributors'
                url='https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}.png'
            />
            <LocationMarker onLocationChange={handleLocationChange} />
        </MapContainer>
    );
};

export default Map;
