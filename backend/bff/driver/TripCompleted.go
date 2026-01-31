package driver

import (
	"backend/bff"
	"time"
	"github.com/gin-gonic/gin"
)

func TripCompletedScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Trip data (this would normally come from a database)
	completedAt := time.Now()
	tripData := map[string]string{
		"tripId":        "TRK789456",
		"origin":        "Mumbai Warehouse",
		"destination":   "Delhi Distribution Center",
		"distance":      "1,450 km",
		"duration":      "2 days 4 hours",
		"payment":       "₹45,000",
		"rating":        "4.8",
		"completedAt":   completedAt.Format("Monday, 2 January 2006"),
		"completedTime": completedAt.Format("15:04"),
	}

	// Trip stats
	tripStats := []bff.RouteData{
		{ID: "distance", Name: "Distance Covered", Description: tripData["distance"]},
		{ID: "ontime", Name: "On-time Delivery", Description: "100%"},
		{ID: "fuel", Name: "Fuel Efficiency", Description: "6.2 km/l"},
		{ID: "safety", Name: "Safe Driving", Description: "98%"},
	}

	// Build UI snippets
	ui := []bff.UISnippet{
		{
			Type: "VIEW",
			Data: bff.ViewData{Flex: 1, BackgroundColor: "#F8FAFC"},
			Children: []bff.UISnippet{
				// Success Header
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:       "🎉 Trip Completed Successfully!",
						FontSize:   28,
						FontWeight: "800",
						Color:      "#1E293B",
						TextAlign:  "center",
					},
				},
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:      "You have successfully delivered the goods to the destination",
						FontSize:  16,
						Color:     "#6B7280",
						TextAlign: "center",
					},
				},
				// Trip Summary Card
				{
					Type: "VIEW",
					Data: bff.ViewData{Padding: 24, BorderRadius: 20, BackgroundColor: "#FFFFFF"},
					Children: []bff.UISnippet{
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:       "Trip Summary",
								FontSize:   20,
								FontWeight: "700",
								Color:      "#1E293B",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:       "Trip ID: " + tripData["tripId"],
								FontSize:   15,
								FontWeight: "700",
								Color:      "#FF0000",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:       "From: " + tripData["origin"],
								FontSize:   16,
								FontWeight: "600",
								Color:      "#1E293B",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:       "To: " + tripData["destination"],
								FontSize:   16,
								FontWeight: "600",
								Color:      "#1E293B",
							},
						},
						// Trip Details Grid
						{
							Type: "VIEW",
							Data: bff.ViewData{FlexDirection: "row", FlexGrow: 1, Gap: 12},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{Text: "Distance: " + tripData["distance"], FontSize: 18, FontWeight: "700"},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{Text: "Duration: " + tripData["duration"], FontSize: 18, FontWeight: "700"},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{Text: "Payment: " + tripData["payment"], FontSize: 18, FontWeight: "700"},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{Text: "Rating: " + tripData["rating"] + "/5", FontSize: 18, FontWeight: "700"},
								},
							},
						},
						// Completion time
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:       "Completed on " + tripData["completedAt"] + " at " + tripData["completedTime"],
								FontSize:   13,
								FontWeight: "600",
								Color:      "#065F46",
							},
						},
					},
				},
				// Trip Stats
				{
					Type: "VIEW",
					Data: bff.ViewData{FlexDirection: "row", Gap: 16},
					Children: func() []bff.UISnippet {
						var stats []bff.UISnippet
						for _, s := range tripStats {
							stats = append(stats, bff.UISnippet{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       s.Name + ": " + s.Description,
									FontSize:   16,
									FontWeight: "600",
									Color:      "#1E293B",
								},
							})
						}
						return stats
					}(),
				},
				// Action Buttons
				{
					Type: "BUTTON",
					Data: bff.ButtonData{
						Text:   "Go to Home",
						Action: bff.ActionData{Type: "NAVIGATE", Navigate: "/(footbar)/home"},
					},
				},
				{
					Type: "BUTTON",
					Data: bff.ButtonData{
						Text:   "Rate This Trip",
						Action: bff.ActionData{Type: "ACTION", Value: "rateTrip"},
					},
				},
				{
					Type: "BUTTON",
					Data: bff.ButtonData{
						Text:   "View Detailed Report",
						Action: bff.ActionData{Type: "ACTION", Value: "viewTripDetails"},
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "TripCompleted",
		UI:     ui,
	}

	c.JSON(200, response)
}
