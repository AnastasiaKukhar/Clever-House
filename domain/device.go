package domain

type Device struct {
	Id      uint64
	HouseId uint64
	UserId  uint64
	Name    string
	Type    DeviceType
	Units   *string //поле може бути порожнім//
}

type DeviceType string

const (
	LightSensor       DeviceType = "LIGHT_SENSOR"
	TemperatureSensor DeviceType = "TEMPERATURE_SENSOR"
	HumiditySensor    DeviceType = "HUMIDITY_SENSOR"
)
