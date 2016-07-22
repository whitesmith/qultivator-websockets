package models

type Payload struct {
	Id string `json:"id" form:"id"`
	Flower struct {
	       SoilTemperature float32 `json:"sT,omitempty" form:"sT"`
	       SoilHumidity float32 `json:"sH,omitempty" form:"sH"`
	       EnvironmentHumidity float32 `json:"eH,omitempty" form:"eH"`
	       EnvironmentTemperature float32 `json:"eT,omitempty" form:"eT"`
	       LightTemperature float32 `json:"lT,omitempty" form:"lT"`
	       LightIntensity float32 `json:"lI,omitempty" form:"lI"`
        } `json:"state" form:"state"`
}
