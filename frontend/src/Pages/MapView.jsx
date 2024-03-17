import React, { useEffect, useState } from 'react'
import { MapContainer, Polyline, TileLayer, Marker, Popup } from 'react-leaflet';
import MarkerClusterGroup from 'react-leaflet-cluster';
import 'leaflet/dist/leaflet.css';

const MapView = () => {
  const [initialCoordinates, setInitialCoordinates] = useState([]);
  const [mapPoints, setMapPoints] = useState([]);
  const [positions, setPositions] = useState([]);

  useEffect(() => {
    const coordinates =[
      [
        [
          78.8045938515902,
          21.384452367356985
        ],
        [
          78.8045938515902,
          21.014236297073978
        ],
        [
          79.37816644715724,
          21.014236297073978
        ],
        [
          79.37816644715724,
          21.384452367356985
        ],
        [
          78.8045938515902,
          21.384452367356985
        ]
      ]
    ]

    const co = coordinates[0]
    const first = reversArray(co[0]);
    const map = co.map((e,i)=>{
        return {location:reversArray(e)}
    })
    console.log(first,map);

    setInitialCoordinates(first)
    setMapPoints(map)
    const linePositions = [];
        for (let i = 0; i < mapPoints.length - 1; i++) {
            linePositions.push([mapPoints[i].location, mapPoints[i + 1].location]);
        }
        setPositions(linePositions);
  }, [])

  const reversArray = (val) =>{
    const f = [val[1],val[0]]
    return f;
  }
  

  return (
    <div>
    
  {initialCoordinates.length >0 && <MapContainer zoom={8} style={{margin:"auto",marginTop:"100px", height: '500px', width: '90%', borderRadius: '10px', }} center={initialCoordinates}>
    <TileLayer url="http://{s}.google.com/vt/lyrs=s,h&x={x}&y={y}&z={z}" maxZoom={20} subdomains={['mt0', 'mt1', 'mt2', 'mt3']} />
    {positions.length > 0 && positions.map((position, index) => <Polyline key={index} positions={position} color={"white"} />)}
    <MarkerClusterGroup chunkedLoading>
      {mapPoints.map((e, i) => (
        <Marker key={i} position={e.location}/>
      ))}
    </MarkerClusterGroup>
  </MapContainer>}
  </div>
  )
}

export default MapView

