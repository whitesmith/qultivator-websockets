# Qultivator
Qultivator was a project built under 32h on the July's Whitesmith Hackathon.

## Qultivator Websockets Server
Webserver for the websockets connection on the Qultivator project

### Endpoints
- `/flower` to establish plant connection
- `/user` to establish user connection

### Payloads
- Users send info in the format: `{"id":"hq-plant", "action":"water", "value":1}`
- Users receive info in the format: `{"id":"hq-plant", "data":{"sT":26.37,"eH":54.80,"eT":26.40,"lT":4036,"lI":637,"sH":16}}`
- Plants send info in the format: `{"id":"hq-plant", "data":{"sT":26.37,"eH":54.80,"eT":26.40,"lT":4036,"lI":637,"sH":16}}`
- Plants receive info in the format: `waterON`

### Json data
- SoilTemperature float32 `sT`
- SoilHumidity float32 `sH`
- EnvironmentHumidity float32 `eH`
- EnvironmentTemperature float32 `eT`
- LightTemperature float32 `lT`
- LightIntensity float32 `lI`
