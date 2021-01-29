package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("charge watch", func() {
	Title("Charge Watcher")
	Server("chargewatch", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("X-Shared-Secret")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})
})

var _ = Service("chargewatch", func() {
	Method("healthcheck", func() {
		Result(String)
		HTTP(func() {
			GET("/healthcheck")
			Response(StatusOK)
		})
	})

	Method("listDevices", func() {
		Payload(func() {
			Attribute("userID", String, "userID")
			Required("userID")
		})
		Result(func() {
			Attribute("devices", ArrayOf(Device))
			Required("devices")
		})
		HTTP(func() {
			GET("/user/{userID}/devices")
			Response(StatusOK)
		})
	})

	Method("createDevice", func() {
		Payload(func() {
			Attribute("userID", String, "userID")
			Attribute("name", String, "device name")
			Attribute("description", String, "description", func() {
				Default("")
			})

			Required("userID", "name")
		})
		Result(func() {
			Attribute("device", Device)
			Required("device")
		})
		Error("StatusInternalServerError")
		HTTP(func() {
			POST("/user/{userID}/devices")
			Body(func() {
				Attribute("name")
				Attribute("description")
			})

			Response(StatusOK)
			Response("StatusInternalServerError", StatusInternalServerError)
		})
	})

	Method("updateCharge", func() {
		Payload(func() {
			Attribute("deviceID", String, "deviceID")
			Attribute("chargeValue", Int)
			Attribute("charging", Boolean)

			Required("deviceID", "chargeValue", "charging")
		})
		Result(func() {
			Attribute("device", Device)
			Required("device")
		})
		Error("StatusInternalServerError")
		Error("StatusBadRequest")
		HTTP(func() {
			POST("/devices/{deviceID}/charge")
			Body(func() {
				Attribute("chargeValue")
				Attribute("charging")
			})

			Response(StatusOK)
			Response("StatusBadRequest", StatusBadRequest)
			Response("StatusInternalServerError", StatusInternalServerError)
		})
	})

	Method("getChargeHistory", func() {
		Payload(func() {
			Attribute("deviceID", String, "deviceID")
			Required("deviceID")
		})
		Result(func() {
			Attribute("chargeHistory", ChargeHistory)
			Required("chargeHistory")
		})
		HTTP(func() {
			GET("/devices/{deviceID}/chargeHistory")
			Response(StatusOK)
		})
	})

	Method("updateDevice", func() {
		Payload(func() {
			Attribute("userID", String, "userID")
			Attribute("deviceID", String, "deviceID")
			Attribute("chargeValue", Int, "value [0,100]", func() {
				Minimum(0)
				Maximum(100)
			})
			Required("userID", "deviceID", "chargeValue")
		})
		Result(func() {
			Attribute("device", Device)
			Required("device")
		})
		HTTP(func() {
			POST("/user/{userID}/devices/{deviceID}")
			Body(func() {
				Attribute("chargeValue")
			})
			Response(StatusOK)
		})
	})

})

var Device = ResultType("Device", func() {
	Attribute("id", String)
	Attribute("userID", String)
	Attribute("name", String)
	Attribute("description", String)
	Attribute("state", String)
	Attribute("charge", Charge)

	Required("id", "userID", "name", "description", "state", "charge")
})

var Charge = Type("Charge", func() {
	Attribute("value", Int)
	Attribute("charging", Boolean)
	Attribute("timestamp", Int64)

	Required("value", "charging", "timestamp")
})

var ChargeHistory = Type("ChargeHistory", func() {
	Attribute("deviceID", String)
	Attribute("logs", ArrayOf(ChargeLog))

	Required("deviceID", "logs")

})

var ChargeLog = Type("ChargeLog", func() {
	Attribute("deviceID", String)
	Attribute("charge", Charge)

	Required("deviceID", "charge")

})
