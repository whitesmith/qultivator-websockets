# Qultivator
Qultivator is a project built under 32h for the July's Whitesmith Hackathon.

It is a IoT project that allows plants to be individually monitored and taken cared of in real-time.
Several environmental data can be gathered using sensors, and both a light and irrigation system can be controlled in real-time using the Qultivator website.

- https://qultivator.whitesmith.co
- https://github.com/whitesmith/qultivator-firmware
- https://github.com/whitesmith/qultivator-website

## Qultivator Websockets Server
Webserver for the websockets connection on the Qultivator project.

This repository includes example client applications for communication testing and the server implementation, which connects plant websockets and user websockets, enabling bi-directional real-time comunication between users and plants.

### Endpoints
- `/flower` to establish plant connection
- `/user` to establish user connection

### Payloads
- Users send info in the format: `{"id":"hq-plant", "action":"water", "value":1}`
- Users receive info in the format: `{"id":"hq-plant", "data":{"sT":26.37,"eH":54.80,"eT":26.40,"lT":4036,"lI":637,"sH":16}}`
- Plants send info in the format: `{"id":"hq-plant", "data":{"sT":26.37,"eH":54.80,"eT":26.40,"lT":4036,"lI":637,"sH":16}}`
- Plants receive info in the format: `waterON`

### JSON data
- SoilTemperature float32 `sT`
- SoilHumidity float32 `sH`
- EnvironmentHumidity float32 `eH`
- EnvironmentTemperature float32 `eT`
- LightTemperature float32 `lT`
- LightIntensity float32 `lI`
