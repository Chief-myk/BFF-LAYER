package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func LoadScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		// Main Container
		{
			Type: "View",
			Data: bff.ViewData{
				Flex:            1,
				BackgroundColor: "#ffffff",
			},
			Children: []bff.UISnippet{
				// Header
				{
					Type: "View",
					Data: bff.ViewData{
						FlexDirection:   "row",
						JustifyContent:  "space-between",
						AlignItems:      "center",
						PaddingHorizontal: 20,
						PaddingTop:      10,
						PaddingBottom:   20,
						BackgroundColor: "#fff",
						BorderBottomWidth: 1,
						BorderColor:      "#f0f0f0",
					},
					Children: []bff.UISnippet{
						{
							Type: "View",
							Data: bff.ViewData{
								Flex: 1,
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Loads",
										FontSize:  28,
										FontWeight: "bold",
										Color:     "#1a1a1a",
										MarginBottom: 4,
									},
								},
								{
									Type: "View",
									Data: bff.ViewData{
										Width:           60,
										Height:          4,
										BackgroundColor: "#ff0000",
										BorderRadius:    2,
									},
								},
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								AlignItems:    "center",
							},
							Children: []bff.UISnippet{
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											MarginLeft: 16,
										},
										OnPress: bff.ActionData{
											Type: "search",
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "search",
												Size:  24,
												Color: "#1a1a1a",
											},
										},
									},
								},
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											MarginLeft: 16,
										},
										OnPress: bff.ActionData{
											Type: "filter",
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "filter",
												Size:  24,
												Color: "#1a1a1a",
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
					Type: "View",
					Data: bff.ViewData{
						FlexDirection: "row",
						PaddingHorizontal: 10,
						PaddingVertical:   7,
						MarginTop:        20,
						BorderBottomWidth: 1,
						BorderColor:      "#f0f0f0",
					},
					Children: []bff.UISnippet{
						createTab("Active Loads", true),
						createTab("Pending Loads", false),
						createTab("Completed Loads", false),
					},
				},
				// Loads List
				{
					Type: "FlatList",
					Data: bff.ViewData{
						Flex:               1,
						ShowsVerticalScrollIndicator: false,
					},
					Children: []bff.UISnippet{
						// Load Card 1
						createLoadCard("LD-7891", "Mumbai Warehouse", "Delhi Distribution Center", 
							"Electronics", "20ft Container Truck", "₹45,000", "Bidding Open", 5, 
							"1400 km", "15 Tons", "20x8x8 ft", "Fragile items - Handle with care",
							map[string]string{
								"biddingStart": "2024-01-15 10:00",
								"biddingEnd":   "2024-01-17 18:00",
							}),
						// Load Card 2
						createLoadCard("LD-7892", "Ahmedabad Factory", "Chennai Port", 
							"Textiles", "32ft Trailer", "₹38,000", "Bids Received", 3, 
							"1600 km", "22 Tons", "32x8x8 ft", "Waterproof packaging required",
							map[string]string{
								"biddingStart": "2024-01-14 09:00",
								"biddingEnd":   "2024-01-16 17:00",
							}),
					},
				},
			},
		},
	}

	// Create response with dynamic data
	response := bff.ScreenResponse{
		Status: "success",
		Screen: "load",
		UI:     ui,
		Data: map[string]interface{}{
			"tabs": []string{"Active Loads", "Pending Loads", "Completed Loads"},
			"activeTab": "Active Loads",
			"loads": map[string][]map[string]interface{}{
				"Active Loads": {
					{
						"id":          "LD-7891",
						"pickup":      "Mumbai Warehouse",
						"drop":        "Delhi Distribution Center",
						"cargoType":   "Electronics",
						"vehicleType": "20ft Container Truck",
						"budget":      "₹45,000",
						"status":      "Bidding Open",
						"bids":        5,
						"distance":    "1400 km",
						"weight":      "15 Tons",
						"dimensions":  "20x8x8 ft",
						"notes":       "Fragile items - Handle with care",
						"timeline": map[string]string{
							"biddingStart": "2024-01-15 10:00",
							"biddingEnd":   "2024-01-17 18:00",
						},
					},
					{
						"id":          "LD-7892",
						"pickup":      "Ahmedabad Factory",
						"drop":        "Chennai Port",
						"cargoType":   "Textiles",
						"vehicleType": "32ft Trailer",
						"budget":      "₹38,000",
						"status":      "Bids Received",
						"bids":        3,
						"distance":    "1600 km",
						"weight":      "22 Tons",
						"dimensions":  "32x8x8 ft",
						"notes":       "Waterproof packaging required",
						"timeline": map[string]string{
							"biddingStart": "2024-01-14 09:00",
							"biddingEnd":   "2024-01-16 17:00",
						},
					},
				},
				"Pending Loads": {
					{
						"id":          "LD-7893",
						"pickup":      "Pune Industrial Area",
						"drop":        "Bangalore Tech Park",
						"cargoType":   "Machinery Parts",
						"vehicleType": "10ft Truck",
						"budget":      "₹25,000",
						"status":      "Driver Assigned",
						"bids":        0,
						"distance":    "850 km",
						"weight":      "8 Tons",
						"dimensions":  "10x6x6 ft",
						"notes":       "Heavy machinery - secure properly",
						"timeline": map[string]string{
							"biddingStart": "2024-01-13 08:00",
							"biddingEnd":   "2024-01-15 16:00",
						},
					},
				},
				"Completed Loads": {
					{
						"id":          "LD-7890",
						"pickup":      "Chennai Warehouse",
						"drop":        "Kolkata Depot",
						"cargoType":   "Consumer Goods",
						"vehicleType": "24ft Truck",
						"budget":      "₹32,000",
						"status":      "Delivered",
						"bids":        0,
						"distance":    "1700 km",
						"weight":      "18 Tons",
						"dimensions":  "24x8x8 ft",
						"notes":       "Standard handling",
						"timeline": map[string]string{
							"biddingStart": "2024-01-10 10:00",
							"biddingEnd":   "2024-01-12 18:00",
						},
					},
				},
			},
			"bidsData": []map[string]interface{}{
				{
					"id":            "BID-001",
					"driverName":    "Rajesh Kumar",
					"truckType":     "20ft Container Truck",
					"rating":        4.8,
					"bidPrice":      "₹42,500",
					"phone":         "+91 98765 43210",
					"experience":    "5 years",
					"completedTrips": 234,
				},
				{
					"id":            "BID-002",
					"driverName":    "Suresh Patel",
					"truckType":     "22ft Truck",
					"rating":        4.6,
					"bidPrice":      "₹44,000",
					"phone":         "+91 98765 43211",
					"experience":    "3 years",
					"completedTrips": 156,
				},
				{
					"id":            "BID-003",
					"driverName":    "Vikram Singh",
					"truckType":     "20ft Container Truck",
					"rating":        4.9,
					"bidPrice":      "₹43,000",
					"phone":         "+91 98765 43212",
					"experience":    "7 years",
					"completedTrips": 312,
				},
			},
			"ownTrucksData": []map[string]interface{}{
				{
					"id":              "TRK-001",
					"truckType":       "20ft Container Truck",
					"capacity":        "15 Tons",
					"registration":    "MH 01 AB 1234",
					"driverName":      "Anil Sharma",
					"documentVerified": true,
				},
				{
					"id":              "TRK-002",
					"truckType":       "24ft Truck",
					"capacity":        "18 Tons",
					"registration":    "GJ 05 CD 5678",
					"driverName":      "Mahesh Reddy",
					"documentVerified": true,
				},
			},
			"statusStyles": map[string]map[string]string{
				"Bidding Open": {
					"bgColor": "#D4EDDA",
					"color":   "#1a1a1a",
				},
				"Bids Received": {
					"bgColor": "#CCE5FF",
					"color":   "#1a1a1a",
				},
				"Driver Assigned": {
					"bgColor": "#FFF3CD",
					"color":   "#1a1a1a",
				},
				"Delivered": {
					"bgColor": "#E2E3E5",
					"color":   "#1a1a1a",
				},
			},
		},
	}

	c.JSON(200, response)
}

// Helper function to create tab
func createTab(label string, isActive bool) bff.UISnippet {
	bgColor := "transparent"
	textColor := "#666"

	if isActive {
		bgColor = "#ff0000"
		textColor = "#fff"
	}

	return bff.UISnippet{
		Type: "TouchableOpacity",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				Flex:             1,
				PaddingVertical:  12,
				AlignItems:       "center",
				BorderRadius:     8,
				MarginHorizontal: 4,
				BackgroundColor:  bgColor,
			},
			OnPress: bff.ActionData{
				Type: "updateScreen",
				Data: map[string]interface{}{
					"activeTab": label,
				},
			},
		},
		Children: []bff.UISnippet{
			{
				Type: "Text",
				Data: bff.TextData{
					Text:       label,
					FontSize:   12,
					FontWeight: "900",
					Color:      textColor,
				},
			},
		},
	}
}


// Helper function to create load card
func createLoadCard(id, pickup, drop, cargoType, vehicleType, budget, status string, 
	bids int, distance, weight, dimensions, notes string, timeline map[string]string) bff.UISnippet {
	
	statusStyles := map[string]map[string]string{
		"Bidding Open": {
			"bgColor": "#D4EDDA",
			"color":   "#1a1a1a",
		},
		"Bids Received": {
			"bgColor": "#CCE5FF",
			"color":   "#1a1a1a",
		},
		"Driver Assigned": {
			"bgColor": "#FFF3CD",
			"color":   "#1a1a1a",
		},
		"Delivered": {
			"bgColor": "#E2E3E5",
			"color":   "#1a1a1a",
		},
	}

	statusStyle := statusStyles[status]
	if statusStyle == nil {
		statusStyle = statusStyles["Bidding Open"]
	}

	return bff.UISnippet{
		Type: "TouchableOpacity",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				BackgroundColor: "#fff",
				BorderRadius:    16,
				Padding:         20,
				Margin:          20,
				ShadowColor:     "#000",
				ShadowOffsetX:   0,
				ShadowOffsetY:   4,
				ShadowOpacity:   0.1,
				ShadowRadius:    8,
				Elevation:       4,
				BorderWidth:     1,
				BorderColor:     "#f0f0f0",
			},
			OnPress: bff.ActionData{
				Type: "showModal",
				Data: map[string]interface{}{
					"modal": "loadDetail",
					"loadId": id,
				},
			},
		},
		Children: []bff.UISnippet{
			// Card Header
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					AlignItems:    "flex-start",
					MarginBottom:  16,
				},
				Children: []bff.UISnippet{
					// Route Info
					{
						Type: "View",
						Data: bff.ViewData{
							Flex: 1,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      pickup + " → " + drop,
									FontSize:  16,
									Color:     "#1a1a1a",
									MarginBottom: 4,
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      distance,
									FontSize:  14,
									Color:     "#666",
								},
							},
						},
					},
					// Bid Badge
					{
						Type: "View",
						Data: bff.ViewData{
							BackgroundColor: "#ff0000",
							PaddingHorizontal: 12,
							PaddingVertical:   6,
							BorderRadius:      12,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      string(rune(bids)) + " Bids",
									Color:     "#fff",
									FontSize:  12,
									FontWeight: "600",
								},
							},
						},
					},
				},
			},
			// Card Details
			{
				Type: "View",
				Data: bff.ViewData{
					MarginBottom: 16,
				},
				Children: []bff.UISnippet{
					// Cargo Type
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							JustifyContent: "space-between",
							AlignItems:    "center",
							MarginBottom:  8,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Cargo:",
									FontSize:  14,
									Color:     "#666",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      cargoType,
									FontSize:  14,
									Color:     "#1a1a1a",
									FontWeight: "500",
								},
							},
						},
					},
					// Vehicle Type
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							JustifyContent: "space-between",
							AlignItems:    "center",
							MarginBottom:  8,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Vehicle:",
									FontSize:  14,
									Color:     "#666",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      vehicleType,
									FontSize:  14,
									Color:     "#1a1a1a",
									FontWeight: "500",
								},
							},
						},
					},
					// Budget
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							JustifyContent: "space-between",
							AlignItems:    "center",
							MarginBottom:  8,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Budget:",
									FontSize:  14,
									Color:     "#666",
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      budget,
									FontSize:  16,
									Color:     "#ff0000",
									FontWeight: "bold",
								},
							},
						},
					},
				},
			},
			// Card Footer
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					AlignItems:    "center",
				},
				Children: []bff.UISnippet{
					// Status Badge
					{
						Type: "View",
						Data: bff.ViewData{
							PaddingHorizontal: 12,
							PaddingVertical:   6,
							BorderRadius:      12,
							BackgroundColor:   statusStyle["bgColor"],
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      status,
									FontSize:  12,
									FontWeight: "600",
									Color:     statusStyle["color"],
								},
							},
						},
					},
					// View Details
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "View Details",
									Color:     "#ff0000",
									FontSize:  14,
									FontWeight: "600",
									MarginRight: 4,
								},
							},
							{
								Type: "Icon",
								Data: bff.IconData{
									Name:  "chevron-right",
									Size:  20,
									Color: "#ff0000",
								},
							},
						},
					},
				},
			},
		},
	}
}

// Additional API endpoint for load detail modal
func LoadDetailScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Parse load ID from request
	loadId := c.Query("loadId")

	ui := []bff.UISnippet{
		{
			Type: "Modal",
			Data: map[string]interface{}{
				"animationType": "slide",
				"transparent":   true,
				"visible":       true,
			},
			Children: []bff.UISnippet{
				{
					Type: "View",
					Data: bff.ViewData{
						Flex:             1,
						BackgroundColor: "rgba(0, 0, 0, 0.5)",
						JustifyContent:   "flex-end",
					},
					Children: []bff.UISnippet{
						{
							Type: "View",
							Data: bff.ViewData{
								BackgroundColor: "#fff",
								BorderTopLeftRadius:  24,
								BorderTopRightRadius: 24,
								MaxHeight:            "90%",
							},
							Children: []bff.UISnippet{
								// Modal Header
								{
									Type: "View",
									Data: bff.ViewData{
										FlexDirection: "row",
										JustifyContent: "space-between",
										AlignItems:    "center",
										Padding:       20,
										BorderBottomWidth: 1,
										BorderColor:   "#f0f0f0",
									},
									Children: []bff.UISnippet{
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "Load Details",
												FontSize:  20,
												FontWeight: "bold",
												Color:     "#1a1a1a",
											},
										},
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Padding: 4,
												},
												OnPress: bff.ActionData{
													Type: "closeModal",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "Icon",
													Data: bff.IconData{
														Name:  "close",
														Size:  24,
														Color: "#666",
													},
												},
											},
										},
									},
								},
								// Modal Body (ScrollView)
								{
									Type: "ScrollView",
									Data: bff.ViewData{
										Padding: 20,
									},
									Children: []bff.UISnippet{
										// Load Information Section
										createDetailSection("Load Information", []bff.UISnippet{
											createDetailItem("Pickup Location", "Mumbai Warehouse"),
											createDetailItem("Drop Location", "Delhi Distribution Center"),
											createDetailItem("Cargo Type", "Electronics"),
											createDetailItem("Weight", "15 Tons"),
											createDetailItem("Dimensions", "20x8x8 ft"),
											createDetailItem("Budget", "₹45,000", true),
											createDetailItem("Notes", "Fragile items - Handle with care"),
										}),
										// Bidding Timeline Section
										createDetailSection("Bidding Timeline", []bff.UISnippet{
											createTimelineItem("Bidding Started", "2024-01-15 10:00"),
											createTimelineItem("Bidding Ends", "2024-01-17 18:00"),
										}),
										// Driver Bids Section
										{
											Type: "View",
											Data: bff.ViewData{
												MarginBottom: 24,
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														FlexDirection: "row",
														JustifyContent: "space-between",
														AlignItems:    "center",
														MarginBottom:  12,
													},
													Children: []bff.UISnippet{
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Driver Bids",
																FontSize:  18,
																FontWeight: "bold",
																Color:     "#1a1a1a",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "3 Bids Received",
																FontSize:  14,
																Color:     "#666",
																FontWeight: "500",
															},
														},
													},
												},
												// Bid Cards would go here
											},
										},
										// Place Your Bid Button
										{
											Type: "View",
											Data: bff.ViewData{
												MarginTop:    20,
												MarginBottom: 30,
											},
											Children: []bff.UISnippet{
												{
													Type: "TouchableOpacity",
													Data: bff.TouchableOpacityData{
														Style: bff.ViewData{
															BackgroundColor: "#ff0000",
															FlexDirection:   "row",
															AlignItems:      "center",
															JustifyContent:  "center",
															PaddingVertical: 16,
															BorderRadius:    12,
															ShadowColor:     "#000",
															ShadowOffsetX:   0,
															ShadowOffsetY:   4,
															ShadowOpacity:   0.2,
															ShadowRadius:    8,
															Elevation:       4,
														},
														OnPress: bff.ActionData{
															Type: "placeBid",
															Data: map[string]interface{}{
																"loadId": loadId,
															},
														},
													},
													Children: []bff.UISnippet{
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Place Your Bid",
																Color:     "#fff",
																FontSize:  18,
																FontWeight: "bold",
																MarginLeft: 8,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "loadDetail",
		UI:     ui,
		Data: map[string]interface{}{
			"loadId": loadId,
		},
	}

	c.JSON(200, response)
}

// Helper function to create detail section
func createDetailSection(title string, children []bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type: "View",
		Data: bff.ViewData{
			MarginBottom: 24,
		},
		Children: append([]bff.UISnippet{
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      title,
					FontSize:  18,
					FontWeight: "bold",
					Color:     "#1a1a1a",
					MarginBottom: 12,
				},
			},
		}, children...),
	}
}

// Helper function to create detail item
func createDetailItem(label, value string, isBudget ...bool) bff.UISnippet {
	isBudgetValue := len(isBudget) > 0 && isBudget[0]
	color := "#1a1a1a"
	if isBudgetValue {
		color = "#ff0000"
	}

	return bff.UISnippet{
		Type: "View",
		Data: bff.ViewData{
			MarginBottom: 12,
		},
		Children: []bff.UISnippet{
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      label,
					FontSize:  14,
					Color:     "#666",
				},
			},
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      value,
					FontSize:  14,
					Color:     color,
					FontWeight: "500",
				},
			},
		},
	}
}



// Helper function to create timeline item
func createTimelineItem(label, value string) bff.UISnippet {
	return bff.UISnippet{
		Type: "View",
		Data: bff.ViewData{
			FlexDirection: "row",
			JustifyContent: "space-between",
			AlignItems:    "center",
			MarginBottom:  8,
		},
		Children: []bff.UISnippet{
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      label,
					FontSize:  14,
					Color:     "#666",
				},
			},
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      value,
					FontSize:  14,
					Color:     "#1a1a1a",
					FontWeight: "500",
				},
			},
		},
	}
}