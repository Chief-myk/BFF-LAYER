package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func LiveTripScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		// Main Container View
		{
			Type: "View",
			Data: bff.ViewData{
				Flex:            1,
				BackgroundColor: "#FFFFFF",
			},
			Children: []bff.UISnippet{
				// StatusBar
				{
					Type: "StatusBar",
					Data: bff.StatusBarData{
						BackgroundColor: "#FFFFFF",
						Style:           "dark",
					},
				},
				// Header
				{
					Type: "View",
					Data: bff.ViewData{
						FlexDirection:   "row",
						JustifyContent:  "space-between",
						AlignItems:      "center",
						PaddingHorizontal: 20,
						PaddingTop:      10,
						PaddingBottom:   15,
						BackgroundColor: "#FFFFFF",
						BorderBottomWidth: 1,
						BorderColor:      "#F0F0F0",
					},
					Children: []bff.UISnippet{
						{
							Type: "Text",
							Data: bff.TextData{
								Text:      "Live Trips",
								FontSize:  24,
								FontWeight: "700",
								Color:     "#333",
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								Gap:           15,
							},
							Children: []bff.UISnippet{
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											Padding: 5,
										},
										OnPress: bff.ActionData{
											Type: "navigate",
											To:   "/refresh",
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "refresh",
												Size:  24,
												Color: "#333",
											},
										},
									},
								},
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											Padding: 5,
										},
										OnPress: bff.ActionData{
											Type: "navigate",
											To:   "/filter",
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "filter",
												Size:  24,
												Color: "#333",
											},
										},
									},
								},
							},
						},
					},
				},
				// Tabs Container
				{
					Type: "ScrollView",
					Data: bff.ViewData{
						Horizontal:     true,
						ShowsHorizontalScrollIndicator: false,
						BackgroundColor: "#FFFFFF",
						BorderBottomWidth: 1,
						BorderColor:      "#F0F0F0",
						MaxHeight:       100,
					},
					Children: []bff.UISnippet{
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								PaddingHorizontal: 20,
								PaddingVertical:   30,
								Gap:              10,
							},
							Children: []bff.UISnippet{
								// All Tab
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											PaddingHorizontal: 20,
											PaddingVertical:   8,
											BorderRadius:      20,
											BackgroundColor:   "#ff0000", // Active by default
										},
										OnPress: bff.ActionData{
											Type: "updateScreen",
											Data: map[string]interface{}{
												"selectedTab": "all",
											},
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "All",
												FontSize:  14,
												FontWeight: "600",
												Color:     "#FFFFFF",
											},
										},
									},
								},
								// In Transit Tab
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											PaddingHorizontal: 20,
											PaddingVertical:   8,
											BorderRadius:      20,
											BackgroundColor:   "#F8F9FA",
										},
										OnPress: bff.ActionData{
											Type: "updateScreen",
											Data: map[string]interface{}{
												"selectedTab": "transit",
											},
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "In Transit",
												FontSize:  14,
												FontWeight: "600",
												Color:     "#666",
											},
										},
									},
								},
								// Delayed Tab
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											PaddingHorizontal: 20,
											PaddingVertical:   8,
											BorderRadius:      20,
											BackgroundColor:   "#F8F9FA",
										},
										OnPress: bff.ActionData{
											Type: "updateScreen",
											Data: map[string]interface{}{
												"selectedTab": "delayed",
											},
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "Delayed",
												FontSize:  14,
												FontWeight: "600",
												Color:     "#666",
											},
										},
									},
								},
								// Completed Tab
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											PaddingHorizontal: 20,
											PaddingVertical:   8,
											BorderRadius:      20,
											BackgroundColor:   "#F8F9FA",
										},
										OnPress: bff.ActionData{
											Type: "updateScreen",
											Data: map[string]interface{}{
												"selectedTab": "completed",
											},
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "Completed",
												FontSize:  14,
												FontWeight: "600",
												Color:     "#666",
											},
										},
									},
								},
							},
						},
					},
				},
				// Trips List
				{
					Type: "ScrollView",
					Data: bff.ViewData{
						Flex:               1,
						ShowsVerticalScrollIndicator: false,
						Padding:            16,
					},
					Children: []bff.UISnippet{
						// Trip Card 1
						createTripCard("TRIP-001", "Rajesh Kumar", "HR 38 XX 1234", "Mumbai, Maharashtra", "Delhi, NCR", "1,420 km", "in-transit", "12 hrs 30 min", "2 min ago", true),
						// Trip Card 2
						createTripCard("TRIP-002", "Amit Sharma", "DL 01 AB 5678", "Chennai, Tamil Nadu", "Bangalore, Karnataka", "350 km", "loading", "8 hrs 15 min", "5 min ago", false),
						// Trip Card 3
						createTripCard("TRIP-003", "Vikram Singh", "UP 32 CD 9012", "Kolkata, West Bengal", "Pune, Maharashtra", "1,850 km", "near-destination", "1 hr 45 min", "1 min ago", false),
					},
				},
			},
		},
	}

	// Create response with dynamic data
	response := bff.ScreenResponse{
		Status: "success",
		Screen: "liveTrip",
		UI:     ui,
		Data: map[string]interface{}{
			"tabs": []map[string]interface{}{
				{"id": "all", "label": "All"},
				{"id": "transit", "label": "In Transit"},
				{"id": "delayed", "label": "Delayed"},
				{"id": "completed", "label": "Completed"},
			},
			"statusConfig": map[string]interface{}{
				"in-transit": map[string]string{
					"label":    "In Transit",
					"color":    "#2E86AB",
					"bgColor":  "#E8F4FD",
				},
				"pickup": map[string]string{
					"label":    "Reached Pickup",
					"color":    "#F39C12",
					"bgColor":  "#FEF9E7",
				},
				"loading": map[string]string{
					"label":    "Loading",
					"color":    "#8E44AD",
					"bgColor":  "#F4ECF7",
				},
				"unloading": map[string]string{
					"label":    "Unloading",
					"color":    "#16A085",
					"bgColor":  "#E8F6F3",
				},
				"near-destination": map[string]string{
					"label":    "Near Destination",
					"color":    "#27AE60",
					"bgColor":  "#EAFAF1",
				},
			},
			"trips": []map[string]interface{}{
				{
					"id": "TRIP-001",
					"driver": map[string]interface{}{
						"name":  "Rajesh Kumar",
						"phone": "+91 98765 43210",
						"photo": nil,
					},
					"truck": map[string]interface{}{
						"number": "HR 38 XX 1234",
						"type":   "Open Half Body",
					},
					"route": map[string]interface{}{
						"from":     "Mumbai, Maharashtra",
						"to":       "Delhi, NCR",
						"distance": "1,420 km",
					},
					"status":       "in-transit",
					"eta":          "12 hrs 30 min",
					"lastUpdated":  "2 min ago",
					"currentLocation": "Near Jaipur, RJ",
					"speed":        "65 km/h",
					"distanceLeft": "320 km",
					"timeline": map[string]bool{
						"started":             true,
						"reachedPickup":       true,
						"loading":             true,
						"inTransit":           true,
						"unloading":           false,
						"reachedDestination":  false,
						"podUploaded":         false,
					},
				},
				{
					"id": "TRIP-002",
					"driver": map[string]interface{}{
						"name":  "Amit Sharma",
						"phone": "+91 98765 43211",
						"photo": nil,
					},
					"truck": map[string]interface{}{
						"number": "DL 01 AB 5678",
						"type":   "Container",
					},
					"route": map[string]interface{}{
						"from":     "Chennai, Tamil Nadu",
						"to":       "Bangalore, Karnataka",
						"distance": "350 km",
					},
					"status":       "loading",
					"eta":          "8 hrs 15 min",
					"lastUpdated":  "5 min ago",
					"currentLocation": "Chennai Warehouse",
					"speed":        "0 km/h",
					"distanceLeft": "350 km",
					"timeline": map[string]bool{
						"started":             true,
						"reachedPickup":       true,
						"loading":             true,
						"inTransit":           false,
						"unloading":           false,
						"reachedDestination":  false,
						"podUploaded":         false,
					},
				},
				{
					"id": "TRIP-003",
					"driver": map[string]interface{}{
						"name":  "Vikram Singh",
						"phone": "+91 98765 43212",
						"photo": nil,
					},
					"truck": map[string]interface{}{
						"number": "UP 32 CD 9012",
						"type":   "Trailer",
					},
					"route": map[string]interface{}{
						"from":     "Kolkata, West Bengal",
						"to":       "Pune, Maharashtra",
						"distance": "1,850 km",
					},
					"status":       "near-destination",
					"eta":          "1 hr 45 min",
					"lastUpdated":  "1 min ago",
					"currentLocation": "Pune Outer Ring",
					"speed":        "45 km/h",
					"distanceLeft": "35 km",
					"timeline": map[string]bool{
						"started":             true,
						"reachedPickup":       true,
						"loading":             true,
						"inTransit":           true,
						"unloading":           false,
						"reachedDestination":  false,
						"podUploaded":         false,
					},
				},
			},
		},
	}

	c.JSON(200, response)
}

// Helper function to create trip card
func createTripCard(id, driverName, truckNumber, from, to, distance, status, eta, lastUpdated string, isExpanded bool) bff.UISnippet {
	statusConfig := map[string]map[string]string{
		"in-transit": {
			"label":    "In Transit",
			"color":    "#2E86AB",
			"bgColor":  "#E8F4FD",
		},
		"loading": {
			"label":    "Loading",
			"color":    "#8E44AD",
			"bgColor":  "#F4ECF7",
		},
		"near-destination": {
			"label":    "Near Destination",
			"color":    "#27AE60",
			"bgColor":  "#EAFAF1",
		},
	}

	statusData := statusConfig[status]
	if statusData == nil {
		statusData = statusConfig["in-transit"] // Default
	}

	return bff.UISnippet{
		Type: "TouchableOpacity",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				BackgroundColor: "#FFFFFF",
				BorderRadius:    16,
				Padding:         18,
				MarginBottom:    16,
				ShadowColor:     "#000",
				ShadowOffsetX:   0,
				ShadowOffsetY:   2,
				ShadowOpacity:   0.1,
				ShadowRadius:    6,
				Elevation:       3,
				BorderWidth:     1,
				BorderColor:     "#F0F0F0",
			},
			OnPress: bff.ActionData{
				Type: "updateScreen",
				Data: map[string]interface{}{
					"expandedTrip": id,
				},
			},
		},
		Children: []bff.UISnippet{
			// Card Header
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "flex-start",
					MarginBottom:   16,
				},
				Children: []bff.UISnippet{
					// Driver Info
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Flex:          1,
						},
						Children: []bff.UISnippet{
							{
								Type: "View",
								Data: bff.ViewData{
									Width:           40,
									Height:          40,
									BorderRadius:    20,
									BackgroundColor: "#FFE6E8",
									JustifyContent:  "center",
									AlignItems:      "center",
									MarginRight:     12,
								},
								Children: []bff.UISnippet{
									{
										Type: "Icon",
										Data: bff.IconData{
											Name:  "person",
											Size:  20,
											Color: "#ff0000",
										},
									},
								},
							},
							{
								Type: "View",
								Data: bff.ViewData{},
								Children: []bff.UISnippet{
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      driverName,
											FontSize:  16,
											FontWeight: "600",
											Color:     "#333",
											MarginBottom: 2,
										},
									},
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      truckNumber,
											FontSize:  14,
											Color:     "#666",
											FontWeight: "500",
										},
									},
								},
							},
						},
					},
					// Trip Meta
					{
						Type: "View",
						Data: bff.ViewData{
							AlignItems: "flex-end",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      id,
									FontSize:  12,
									Color:     "#999",
									MarginBottom: 6,
									FontWeight: "500",
								},
							},
							{
								Type: "View",
								Data: bff.ViewData{
									PaddingHorizontal: 12,
									PaddingVertical:   4,
									BorderRadius:      12,
									BackgroundColor:   statusData["bgColor"],
								},
								Children: []bff.UISnippet{
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      statusData["label"],
											FontSize:  12,
											FontWeight: "600",
											Color:     statusData["color"],
										},
									},
								},
							},
						},
					},
				},
			},
			// Route Section
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					AlignItems:    "center",
					MarginBottom:  16,
				},
				Children: []bff.UISnippet{
					{
						Type: "View",
						Data: bff.ViewData{
							Flex: 1,
						},
						Children: []bff.UISnippet{
							// From Location
							{
								Type: "View",
								Data: bff.ViewData{
									FlexDirection: "row",
									AlignItems:    "center",
									MarginBottom:  8,
								},
								Children: []bff.UISnippet{
									{
										Type: "View",
										Data: bff.ViewData{
											Width:           12,
											Height:          12,
											BorderRadius:    6,
											BackgroundColor: "#27AE60",
											MarginRight:     12,
										},
									},
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      from,
											FontSize:  14,
											Color:     "#333",
											FontWeight: "500",
											Flex:      1,
										},
									},
								},
							},
							// Route Line
							{
								Type: "View",
								Data: bff.ViewData{
									Width:           2,
									Height:          12,
									BackgroundColor: "#E0E0E0",
									MarginLeft:      5,
									MarginBottom:    8,
								},
							},
							// To Location
							{
								Type: "View",
								Data: bff.ViewData{
									FlexDirection: "row",
									AlignItems:    "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "View",
										Data: bff.ViewData{
											Width:           12,
											Height:          12,
											BorderRadius:    6,
											BackgroundColor: "#ff0000",
											MarginRight:     12,
										},
									},
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      to,
											FontSize:  14,
											Color:     "#333",
											FontWeight: "500",
											Flex:      1,
										},
									},
								},
							},
						},
					},
					// Map Icon
					{
						Type: "TouchableOpacity",
						Data: bff.TouchableOpacityData{
							Style: bff.ViewData{
								Padding:         8,
								BackgroundColor: "#FFE6E8",
								BorderRadius:    10,
								MarginLeft:      12,
							},
							OnPress: bff.ActionData{
								Type: "navigate",
								To:   "/map/" + id,
							},
						},
						Children: []bff.UISnippet{
							{
								Type: "Icon",
								Data: bff.IconData{
									Name:  "location-on",
									Size:  20,
									Color: "#ff0000",
								},
							},
						},
					},
				},
			},
			// Info Section
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					MarginBottom:   16,
				},
				Children: []bff.UISnippet{
					{
						Type: "View",
						Data: bff.ViewData{
							AlignItems: "center",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "ETA",
									FontSize:  12,
									Color:     "#999",
									MarginBottom: 4,
									FontWeight: "500",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      eta,
									FontSize:  14,
									FontWeight: "600",
									Color:     "#333",
								},
							},
						},
					},
					{
						Type: "View",
						Data: bff.ViewData{
							AlignItems: "center",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Distance",
									FontSize:  12,
									Color:     "#999",
									MarginBottom: 4,
									FontWeight: "500",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      distance,
									FontSize:  14,
									FontWeight: "600",
									Color:     "#333",
								},
							},
						},
					},
					{
						Type: "View",
						Data: bff.ViewData{
							AlignItems: "center",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Updated",
									FontSize:  12,
									Color:     "#999",
									MarginBottom: 4,
									FontWeight: "500",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      lastUpdated,
									FontSize:  14,
									FontWeight: "600",
									Color:     "#333",
								},
							},
						},
					},
				},
			},
			// Action Buttons
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					Gap:           12,
				},
				Children: []bff.UISnippet{
					{
						Type: "TouchableOpacity",
						Data: bff.TouchableOpacityData{
							Style: bff.ViewData{
								Flex:             1,
								FlexDirection:    "row",
								AlignItems:       "center",
								JustifyContent:   "center",
								PaddingVertical:  12,
								BorderRadius:     10,
								BackgroundColor:  "#ff0000",
								Gap:              8,
							},
							OnPress: bff.ActionData{
								Type: "call",
								Data: map[string]interface{}{
									"phone": "+91 98765 43210",
								},
							},
						},
						Children: []bff.UISnippet{
							{
								Type: "Icon",
								Data: bff.IconData{
									Name:  "call",
									Size:  16,
									Color: "#FFFFFF",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Call",
									Color:     "#FFFFFF",
									FontSize:  16,
									FontWeight: "600",
								},
							},
						},
					},
					{
						Type: "TouchableOpacity",
						Data: bff.TouchableOpacityData{
							Style: bff.ViewData{
								Flex:             1,
								FlexDirection:    "row",
								AlignItems:       "center",
								JustifyContent:   "center",
								PaddingVertical:  12,
								BorderRadius:     10,
								BackgroundColor:  "#FFFFFF",
								BorderWidth:      1,
								BorderColor:      "#ff0000",
								Gap:              8,
							},
							OnPress: bff.ActionData{
								Type: "message",
								Data: map[string]interface{}{
									"driver": driverName,
								},
							},
						},
						Children: []bff.UISnippet{
							{
								Type: "Icon",
								Data: bff.IconData{
									Name:  "chatbubble",
									Size:  16,
									Color: "#ff0000",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Message",
									Color:     "#ff0000",
									FontSize:  16,
									FontWeight: "600",
								},
							},
						},
					},
				},
			},
		},
	}
}