// package main

// import (
// 	"log"
// 	"net/http"
// 	"backend/bff"
// 	"backend/bff/auth"
// 	"backend/bff/driver"
// 	"backend/bff/broker"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("👋 Hello World! 🚀🔥"))
// 	})
// 	http.HandleFunc("/bff/splash", bff.SplashScreenHandler)

// 	http.HandleFunc("/bff/auth/auth", auth.AuthScreenHandler)
// 	http.HandleFunc("/bff/auth/otp", auth.OtpScreenHandler)
// 	http.HandleFunc("/bff/auth/registration-role", auth.RegistrationRoleHandler)
// 	http.HandleFunc("/bff/auth/r1", auth.R1Screen)
// 	http.HandleFunc("/bff/auth/r2", auth.R2Screen)
// 	http.HandleFunc("/bff/auth/r3", auth.R3Screen)
// 	http.HandleFunc("/bff/auth/r4", auth.R4Screen)
// 	http.HandleFunc("/bff/auth/r5", auth.R5Screen)
// 	http.HandleFunc("/bff/auth/r6", auth.R6Screen)
// 	http.HandleFunc("/bff/auth/r7", auth.R7Screen)
// 	http.HandleFunc("/bff/auth/r8", auth.R8Screen)
// 	http.HandleFunc("/bff/auth/g1", auth.G1Screen)
// 	http.HandleFunc("/bff/auth/g2", auth.G2Screen)
// 	http.HandleFunc("/bff/auth/g3", auth.G3Screen)
// 	http.HandleFunc("/bff/auth/g4", auth.G4Screen)
// 	http.HandleFunc("/bff/auth/g5", auth.G5Screen)

// 	http.HandleFunc("/bff/driver/tripCompleted", driver.TripCompletedScreen)
// 	http.HandleFunc("/bff/driver/profile", driver.ProfileScreen)
// 	http.HandleFunc("/bff/driver/payment", driver.PaymentScreen)
// 	http.HandleFunc("/bff/driver/market", driver.MarketScreen)
// 	http.HandleFunc("/bff/driver/home", driver.HomeScreen)
// 	http.HandleFunc("/bff/driver/home/action", driver.HandleHomeAction)
// 	http.HandleFunc("/bff/driver/mytrip", driver.MyTripScreen)

// 	http.HandleFunc("/bff/broker/addload", broker.AddLoadScreen)
// 	http.HandleFunc("/bff/broker/addtruck", broker.AddTruckScreen)
// 	http.HandleFunc("/bff/broker/profile", broker.ProfileScreen)
// 	http.HandleFunc("/bff/broker/money", broker.MoneyScreen)
// 	http.HandleFunc("/bff/broker/load", broker.LoadScreen)
// 	http.HandleFunc("/bff/broker/home", broker.HomeScreen)
// 	http.HandleFunc("/bff/broker/livetrip", broker.LiveTripScreen)

// 	log.Println("BFF server running on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }


package main

import (
	"log"
	"backend/bff"
	"backend/bff/auth"
	"backend/bff/driver"
	"backend/bff/broker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	// Optional: allow CORS for all routes
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "👋 Hello World! 🚀🔥")
	})

	// BFF routes grouped by feature
	bffGroup := r.Group("/bff")
	{
		// Splash
		bffGroup.GET("/splash", bff.SplashScreenHandler)

		// Auth routes
		authGroup := bffGroup.Group("/auth")
		{
			authGroup.GET("/auth", auth.AuthScreenHandler)
			authGroup.GET("/otp", auth.OtpScreenHandler)
			authGroup.GET("/registration-role", auth.RegistrationRoleHandler)

			authGroup.GET("/r1", auth.R1Screen)
			authGroup.GET("/r2", auth.R2Screen)
			authGroup.GET("/r3", auth.R3Screen)
			authGroup.GET("/r4", auth.R4Screen)
			authGroup.GET("/r5", auth.R5Screen)
			authGroup.GET("/r6", auth.R6Screen)
			authGroup.GET("/r7", auth.R7Screen)
			authGroup.GET("/r8", auth.R8Screen)

			authGroup.GET("/g1", auth.G1Screen)
			authGroup.GET("/g2", auth.G2Screen)
			authGroup.GET("/g3", auth.G3Screen)
			authGroup.GET("/g4", auth.G4Screen)
			authGroup.GET("/g5", auth.G5Screen)
		}

		// Driver routes
		driverGroup := bffGroup.Group("/driver")
		{
			driverGroup.GET("/tripCompleted", driver.TripCompletedScreen)
			driverGroup.GET("/profile", driver.ProfileScreen)
			driverGroup.GET("/payment", driver.PaymentScreen)
			driverGroup.GET("/market", driver.MarketScreen)
			driverGroup.GET("/home", driver.HomeScreen)
			driverGroup.POST("/home/action", driver.HandleHomeAction) // POST if it's an action
			driverGroup.GET("/mytrip", driver.MyTripScreen)
		}

		// Broker routes
		brokerGroup := bffGroup.Group("/broker")
		{
			brokerGroup.GET("/addload", broker.AddLoadScreen)
			brokerGroup.GET("/addtruck", broker.AddTruckScreen)
			brokerGroup.GET("/profile", broker.ProfileScreen)
			brokerGroup.GET("/money", broker.MoneyScreen)
			brokerGroup.GET("/load", broker.LoadScreen)
			brokerGroup.GET("/home", broker.HomeScreen)
			brokerGroup.GET("/livetrip", broker.LiveTripScreen)
		}
	}

	// Start server
	log.Println("BFF server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
