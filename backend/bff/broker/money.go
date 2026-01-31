package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func MoneyScreen(c *gin.Context) {
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
						PaddingHorizontal: 20,
						PaddingTop:       10,
						MarginTop:        10,
						PaddingBottom:    20,
						BackgroundColor:  "#fff",
						BorderBottomWidth: 1,
						BorderColor:      "#f0f0f0",
					},
					Children: []bff.UISnippet{
						{
							Type: "Text",
							Data: bff.TextData{
								Text:      "Payments & Settlements",
								FontSize:  24,
								FontWeight: "bold",
								Color:     "#1a1a1a",
							},
						},
					},
				},
				// Main ScrollView
				{
					Type: "ScrollView",
					Data: bff.ViewData{
						Flex:               1,
						ShowsVerticalScrollIndicator: false,
					},
					Children: []bff.UISnippet{
						// Summary Cards Horizontal Scroll
						{
							Type: "ScrollView",
							Data: bff.ViewData{
								Horizontal: true,
								ShowsHorizontalScrollIndicator: false,
								MarginBottom: 20,
							},
							Children: []bff.UISnippet{
								{
									Type: "View",
									Data: bff.ViewData{
										FlexDirection: "row",
										PaddingHorizontal: 20,
										Gap: 12,
									},
									Children: []bff.UISnippet{
										// Card 1: Pending Payments
										createSummaryCard("1", "Pending Payments", "₹45,200", "12 Payments", "clock-outline", "#ff0000"),
										// Card 2: Completed Payments
										createSummaryCard("2", "Completed Payments", "₹1,85,500", "28 Payments", "check-circle-outline", "#28A745"),
										// Card 3: Total Paid This Month
										createSummaryCard("3", "Total Paid This Month", "₹2,30,700", "+15% from last month", "calendar-month", "#1a1a1a"),
									},
								},
							},
						},
						// Filter Bar
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								AlignItems:    "center",
								PaddingHorizontal: 16,
								MarginBottom:  24,
								Gap:           12,
							},
							Children: []bff.UISnippet{
								// Filter Buttons Scroll
								{
									Type: "ScrollView",
									Data: bff.ViewData{
										Horizontal: true,
										ShowsHorizontalScrollIndicator: false,
										Flex:       1,
									},
									Children: []bff.UISnippet{
										{
											Type: "View",
											Data: bff.ViewData{
												FlexDirection: "row",
											},
											Children: []bff.UISnippet{
												createFilterButton("All", true),
												createFilterButton("Pending", false),
												createFilterButton("Done", false),
											},
										},
									},
								},
								// Date Filter
								{
									Type: "TouchableOpacity",
									Data: bff.TouchableOpacityData{
										Style: bff.ViewData{
											FlexDirection: "row",
											AlignItems:    "center",
											PaddingHorizontal: 12,
											PaddingVertical:   8,
											BorderRadius:     8,
											BackgroundColor:  "#f8f9fa",
											BorderWidth:      1,
											BorderColor:      "#E5E5E5",
											Gap:              6,
										},
										OnPress: bff.ActionData{
											Type: "dateFilter",
										},
									},
									Children: []bff.UISnippet{
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "calendar-month",
												Size:  20,
												Color: "#666",
											},
										},
										{
											Type: "Text",
											Data: bff.TextData{
												Text:      "This Month",
												FontSize:  14,
												Color:     "#666",
												FontWeight: "500",
											},
										},
										{
											Type: "Icon",
											Data: bff.IconData{
												Name:  "chevron-down",
												Size:  16,
												Color: "#666",
											},
										},
									},
								},
							},
						},
						// Payments List Section
						{
							Type: "View",
							Data: bff.ViewData{
								PaddingHorizontal: 20,
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Recent Payments",
										FontSize:  18,
										FontWeight: "bold",
										Color:     "#1a1a1a",
										MarginBottom: 16,
									},
								},
								// Payment Cards List
								{
									Type: "View",
									Data: bff.ViewData{
										Gap: 12,
									},
									Children: []bff.UISnippet{
										// Payment Card 1
										createPaymentCard(
											"TRK-2458",
											"Rajesh Kumar",
											"+91 98765 43210",
											"MH12 AB 1234",
											"₹18,500",
											"Pending",
											"Waiting",
											map[string]string{
												"cargoType":     "Electronics",
												"from":          "Mumbai",
												"to":            "Delhi",
												"distance":      "1400 km",
												"commission":    "₹1,850",
												"payableAmount": "₹16,650",
											},
										),
										// Payment Card 2
										createPaymentCard(
											"TRK-2457",
											"Suresh Patel",
											"+91 98765 43211",
											"GJ01 CD 5678",
											"₹22,300",
											"Completed",
											"Approved",
											map[string]string{
												"cargoType":     "Textiles",
												"from":          "Ahmedabad",
												"to":            "Chennai",
												"distance":      "1600 km",
												"commission":    "₹2,230",
												"payableAmount": "₹20,070",
											},
										),
										// Payment Card 3
										createPaymentCard(
											"TRK-2456",
											"Vikram Singh",
											"+91 98765 43212",
											"DL09 EF 9012",
											"₹15,800",
											"Pending",
											"Approved",
											map[string]string{
												"cargoType":     "Automobile Parts",
												"from":          "Pune",
												"to":            "Bangalore",
												"distance":      "850 km",
												"commission":    "₹1,580",
												"payableAmount": "₹14,220",
											},
										),
										// Payment Card 4
										createPaymentCard(
											"TRK-2455",
											"Anil Sharma",
											"+91 98765 43213",
											"KA05 GH 3456",
											"₹28,900",
											"Completed",
											"Approved",
											map[string]string{
												"cargoType":     "Machinery",
												"from":          "Chennai",
												"to":            "Kolkata",
												"distance":      "1700 km",
												"commission":    "₹2,890",
												"payableAmount": "₹26,010",
											},
										),
									},
								},
							},
						},
						// Bottom Spacer for Floating Button
						{
							Type: "View",
							Data: bff.ViewData{
								Height: 100,
							},
						},
					},
				},
				// Footer with Floating Button
				{
					Type: "View",
					Data: bff.ViewData{
						Position:        "absolute",
						Bottom:          0,
						Left:            0,
						Right:           0,
						BackgroundColor: "#fff",
						PaddingHorizontal: 20,
						PaddingVertical: 16,
						BorderTopWidth:  1,
						BorderColor:     "#f0f0f0",
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
									ShadowColor:     "#ff0000",
									ShadowOffsetX:   0,
									ShadowOffsetY:   4,
									ShadowOpacity:   0.3,
									ShadowRadius:    8,
									Elevation:       4,
									Gap:             8,
								},
								OnPress: bff.ActionData{
									Type: "navigate",
									To:   "/new-payment",
								},
							},
							Children: []bff.UISnippet{
								{
									Type: "Icon",
									Data: bff.IconData{
										Name:  "plus",
										Size:  24,
										Color: "#fff",
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Initiate New Payment",
										Color:     "#fff",
										FontSize:  16,
										FontWeight: "600",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create response with dynamic data
	response := bff.ScreenResponse{
		Status: "success",
		Screen: "money",
		UI:     ui,
		Data: map[string]interface{}{
			"summaryCards": []map[string]interface{}{
				{
					"id":     "1",
					"title":  "Pending Payments",
					"amount": "₹45,200",
					"count":  "12 Payments",
					"icon":   "clock-outline",
					"color":  "#ff0000",
				},
				{
					"id":     "2",
					"title":  "Completed Payments",
					"amount": "₹1,85,500",
					"count":  "28 Payments",
					"icon":   "check-circle-outline",
					"color":  "#28A745",
				},
				{
					"id":     "3",
					"title":  "Total Paid This Month",
					"amount": "₹2,30,700",
					"count":  "+15% from last month",
					"icon":   "calendar-month",
					"color":  "#1a1a1a",
				},
			},
			"filters": []string{"All", "Pending", "Done"},
			"selectedFilter": "All",
			"payments": []map[string]interface{}{
				{
					"id":         "TRK-2458",
					"driverName": "Rajesh Kumar",
					"phone":      "+91 98765 43210",
					"truckNumber": "MH12 AB 1234",
					"amount":     "₹18,500",
					"status":     "Pending",
					"podStatus":  "Waiting",
					"tripDetails": map[string]string{
						"cargoType":     "Electronics",
						"from":          "Mumbai",
						"to":            "Delhi",
						"distance":      "1400 km",
						"commission":    "₹1,850",
						"payableAmount": "₹16,650",
					},
				},
				{
					"id":         "TRK-2457",
					"driverName": "Suresh Patel",
					"phone":      "+91 98765 43211",
					"truckNumber": "GJ01 CD 5678",
					"amount":     "₹22,300",
					"status":     "Completed",
					"podStatus":  "Approved",
					"tripDetails": map[string]string{
						"cargoType":     "Textiles",
						"from":          "Ahmedabad",
						"to":            "Chennai",
						"distance":      "1600 km",
						"commission":    "₹2,230",
						"payableAmount": "₹20,070",
					},
				},
				{
					"id":         "TRK-2456",
					"driverName": "Vikram Singh",
					"phone":      "+91 98765 43212",
					"truckNumber": "DL09 EF 9012",
					"amount":     "₹15,800",
					"status":     "Pending",
					"podStatus":  "Approved",
					"tripDetails": map[string]string{
						"cargoType":     "Automobile Parts",
						"from":          "Pune",
						"to":            "Bangalore",
						"distance":      "850 km",
						"commission":    "₹1,580",
						"payableAmount": "₹14,220",
					},
				},
				{
					"id":         "TRK-2455",
					"driverName": "Anil Sharma",
					"phone":      "+91 98765 43213",
					"truckNumber": "KA05 GH 3456",
					"amount":     "₹28,900",
					"status":     "Completed",
					"podStatus":  "Approved",
					"tripDetails": map[string]string{
						"cargoType":     "Machinery",
						"from":          "Chennai",
						"to":            "Kolkata",
						"distance":      "1700 km",
						"commission":    "₹2,890",
						"payableAmount": "₹26,010",
					},
				},
			},
			"statusStyles": map[string]map[string]string{
				"Pending": {
					"bgColor": "#FFE6E6",
					"color":   "#1a1a1a",
				},
				"Completed": {
					"bgColor": "#E8F5E8",
					"color":   "#1a1a1a",
				},
			},
			"podStatusStyles": map[string]map[string]string{
				"Approved": {
					"bgColor": "#E8F5E8",
					"color":   "#1a1a1a",
				},
				"Waiting": {
					"bgColor": "#FFF3CD",
					"color":   "#1a1a1a",
				},
			},
		},
	}

	c.JSON(200, response)
}

// Helper function to create summary card
func createSummaryCard(id, title, amount, count, icon, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "View",
		Data: bff.ViewData{
			BackgroundColor: "#fff",
			BorderRadius:    16,
			Padding:         20,
			Width:           280,
			ShadowColor:     "#000",
			ShadowOffsetX:   0,
			ShadowOffsetY:   2,
			ShadowOpacity:   0.05,
			ShadowRadius:    8,
			Elevation:       2,
			BorderWidth:     1,
			BorderColor:     "#f0f0f0",
		},
		Children: []bff.UISnippet{
			// Card Header
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					AlignItems:    "center",
					MarginBottom:  12,
				},
				Children: []bff.UISnippet{
					{
						Type: "Icon",
						Data: bff.IconData{
							Name:  icon,
							Size:  20,
							Color: color,
						},
					},
					{
						Type: "Text",
						Data: bff.TextData{
							Text:      title,
							FontSize:  14,
							Color:     "#666",
							FontWeight: "500",
							MarginLeft: 8,
						},
					},
				},
			},
			// Card Amount
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      amount,
					FontSize:  24,
					FontWeight: "bold",
					Color:     color,
					MarginBottom: 4,
				},
			},
			// Card Count
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      count,
					FontSize:  14,
					Color:     "#666",
				},
			},
		},
	}
}

// Helper function to create filter button
func createFilterButton(label string, isActive bool) bff.UISnippet {
	bgColor := "#f8f9fa"
	textColor := "#666"
	borderColor := "#E5E5E5"
	
	if isActive {
		bgColor = "#ff0000"
		textColor = "#fff"
		borderColor = "#ff0000"
	}

	return bff.UISnippet{
		Type: "TouchableOpacity",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				PaddingHorizontal: 16,
				PaddingVertical:   8,
				BorderRadius:      20,
				BackgroundColor:   bgColor,
				MarginRight:       8,
				BorderWidth:       1,
				BorderColor:       borderColor,
			},
			OnPress: bff.ActionData{
				Type: "updateScreen",
				Data: map[string]interface{}{
					"selectedFilter": label,
				},
			},
		},
		Children: []bff.UISnippet{
			{
				Type: "Text",
				Data: bff.TextData{
					Text:      label,
					FontSize:  14,
					FontWeight: "500",
					Color:     textColor,
				},
			},
		},
	}
}

// Helper function to create payment card
func createPaymentCard(id, driverName, phone, truckNumber, amount, status, podStatus string, tripDetails map[string]string) bff.UISnippet {
	statusStyles := map[string]map[string]string{
		"Pending": {
			"bgColor": "#FFE6E6",
			"color":   "#1a1a1a",
		},
		"Completed": {
			"bgColor": "#E8F5E8",
			"color":   "#1a1a1a",
		},
	}

	podStatusStyles := map[string]map[string]string{
		"Approved": {
			"bgColor": "#E8F5E8",
			"color":   "#1a1a1a",
		},
		"Waiting": {
			"bgColor": "#FFF3CD",
			"color":   "#1a1a1a",
		},
	}

	return bff.UISnippet{
		Type: "TouchableOpacity",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				BackgroundColor: "#fff",
				BorderRadius:    12,
				Padding:         16,
				BorderWidth:     1,
				BorderColor:     "#f0f0f0",
				ShadowColor:     "#000",
				ShadowOffsetX:   0,
				ShadowOffsetY:   1,
				ShadowOpacity:   0.05,
				ShadowRadius:    4,
				Elevation:       1,
			},
			OnPress: bff.ActionData{
				Type: "showModal",
				Data: map[string]interface{}{
					"modal":   "paymentDetail",
					"paymentId": id,
				},
			},
		},
		Children: []bff.UISnippet{
			// Payment Header
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					AlignItems:    "flex-start",
					MarginBottom:  12,
				},
				Children: []bff.UISnippet{
					// Left: Payment Info
					{
						Type: "View",
						Data: bff.ViewData{
							Flex: 1,
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "Trip ID: " + id,
									FontSize:  14,
									FontWeight: "600",
									Color:     "#1a1a1a",
									MarginBottom: 4,
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      driverName,
									FontSize:  16,
									FontWeight: "bold",
									Color:     "#1a1a1a",
									MarginBottom: 2,
								},
							},
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      phone,
									FontSize:  14,
									Color:     "#666",
								},
							},
						},
					},
					// Right: Amount & Status
					{
						Type: "View",
						Data: bff.ViewData{
							AlignItems: "flex-end",
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      amount,
									FontSize:  18,
									FontWeight: "bold",
									Color:     "#1a1a1a",
									MarginBottom: 8,
								},
							},
							{
								Type: "View",
								Data: bff.ViewData{
									PaddingHorizontal: 8,
									PaddingVertical:   4,
									BorderRadius:      6,
									BackgroundColor:   statusStyles[status]["bgColor"],
								},
								Children: []bff.UISnippet{
									{
										Type: "Text",
										Data: bff.TextData{
											Text:      status,
											FontSize:  12,
											FontWeight: "600",
											Color:     statusStyles[status]["color"],
										},
									},
								},
							},
						},
					},
				},
			},
			// Payment Details
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					AlignItems:    "center",
					MarginBottom:  12,
				},
				Children: []bff.UISnippet{
					// Left: Truck Details
					{
						Type: "View",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           6,
						},
						Children: []bff.UISnippet{
							{
								Type: "Icon",
								Data: bff.IconData{
									Name:  "truck",
									Size:  16,
									Color: "#666",
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
					// Right: POD Status
					{
						Type: "View",
						Data: bff.ViewData{
							PaddingHorizontal: 8,
							PaddingVertical:   4,
							BorderRadius:      6,
							BackgroundColor:   podStatusStyles[podStatus]["bgColor"],
						},
						Children: []bff.UISnippet{
							{
								Type: "Text",
								Data: bff.TextData{
									Text:      "POD: " + podStatus,
									FontSize:  12,
									FontWeight: "600",
									Color:     podStatusStyles[podStatus]["color"],
								},
							},
						},
					},
				},
			},
			// Expand Section
			{
				Type: "View",
				Data: bff.ViewData{
					FlexDirection: "row",
					JustifyContent: "space-between",
					AlignItems:    "center",
					PaddingTop:    12,
					BorderTopWidth: 1,
					BorderColor:   "#f0f0f0",
				},
				Children: []bff.UISnippet{
					{
						Type: "Text",
						Data: bff.TextData{
							Text:      "Tap to view details",
							FontSize:  14,
							Color:     "#ff0000",
							FontWeight: "500",
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
	}
}

// Payment Detail Modal API
func PaymentDetailScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	paymentId := c.Query("paymentId")

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
												Text:      "Payment Details",
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
								// Modal Body
								{
									Type: "ScrollView",
									Data: bff.ViewData{
										Padding: 20,
									},
									Children: createPaymentDetailContent(paymentId),
								},
								// Modal Footer
								createPaymentDetailFooter(paymentId),
							},
						},
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "paymentDetail",
		UI:     ui,
		Data: map[string]interface{}{
			"paymentId": paymentId,
		},
	}

	c.JSON(200, response)
}

// Helper function to create payment detail content
func createPaymentDetailContent(paymentId string) []bff.UISnippet {
	return []bff.UISnippet{
		// Trip Summary Section
		{
			Type: "View",
			Data: bff.ViewData{
				MarginBottom: 24,
			},
			Children: []bff.UISnippet{
				{
					Type: "Text",
					Data: bff.TextData{
						Text:      "Trip Summary",
						FontSize:  18,
						FontWeight: "bold",
						Color:     "#1a1a1a",
						MarginBottom: 12,
					},
				},
				{
					Type: "View",
					Data: bff.ViewData{
						BackgroundColor: "#f8f9fa",
						BorderRadius:    12,
						Padding:         16,
						Gap:             8,
					},
					Children: []bff.UISnippet{
						{
							Type: "Text",
							Data: bff.TextData{
								Text:      "Trip ID: " + paymentId,
								FontSize:  16,
								FontWeight: "600",
								Color:     "#1a1a1a",
							},
						},
						{
							Type: "Text",
							Data: bff.TextData{
								Text:      "Cargo: Electronics",
								FontSize:  16,
								FontWeight: "600",
								Color:     "#1a1a1a",
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								AlignItems:    "center",
								Gap:           8,
							},
							Children: []bff.UISnippet{
								{
									Type: "Icon",
									Data: bff.IconData{
										Name:  "map-marker",
										Size:  16,
										Color: "#ff0000",
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Mumbai → Delhi",
										FontSize:  14,
										Color:     "#666",
										FontWeight: "500",
									},
								},
							},
						},
						{
							Type: "Text",
							Data: bff.TextData{
								Text:      "1400 km",
								FontSize:  14,
								Color:     "#666",
							},
						},
					},
				},
			},
		},
		// Payment Breakdown Section
		{
			Type: "View",
			Data: bff.ViewData{
				MarginBottom: 24,
			},
			Children: []bff.UISnippet{
				{
					Type: "Text",
					Data: bff.TextData{
						Text:      "Payment Breakdown",
						FontSize:  18,
						FontWeight: "bold",
						Color:     "#1a1a1a",
						MarginBottom: 12,
					},
				},
				{
					Type: "View",
					Data: bff.ViewData{
						BackgroundColor: "#f8f9fa",
						BorderRadius:    12,
						Padding:         16,
						Gap:             12,
					},
					Children: []bff.UISnippet{
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								JustifyContent: "space-between",
								AlignItems:    "center",
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Trip Amount",
										FontSize:  14,
										Color:     "#666",
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "₹18,500",
										FontSize:  14,
										FontWeight: "600",
										Color:     "#1a1a1a",
									},
								},
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								JustifyContent: "space-between",
								AlignItems:    "center",
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Commission",
										FontSize:  14,
										Color:     "#666",
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "₹1,850",
										FontSize:  14,
										FontWeight: "600",
										Color:     "#1a1a1a",
									},
								},
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								Height:           1,
								BackgroundColor: "#E5E5E5",
							},
						},
						{
							Type: "View",
							Data: bff.ViewData{
								FlexDirection: "row",
								JustifyContent: "space-between",
								AlignItems:    "center",
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Final Payable Amount",
										FontSize:  16,
										FontWeight: "bold",
										Color:     "#1a1a1a",
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "₹16,650",
										FontSize:  18,
										FontWeight: "bold",
										Color:     "#ff0000",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Helper function to create payment detail footer
func createPaymentDetailFooter(paymentId string) bff.UISnippet {
	return bff.UISnippet{
		Type: "View",
		Data: bff.ViewData{
			FlexDirection: "row",
			Padding:       20,
			Gap:           12,
			BorderTopWidth: 1,
			BorderColor:   "#f0f0f0",
		},
		Children: []bff.UISnippet{
			// Download Invoice Button
			{
				Type: "TouchableOpacity",
				Data: bff.TouchableOpacityData{
					Style: bff.ViewData{
						Flex:             1,
						FlexDirection:   "row",
						BackgroundColor: "#fff",
						PaddingVertical: 16,
						PaddingHorizontal: 12,
						BorderRadius:    12,
						AlignItems:      "center",
						JustifyContent:  "center",
						BorderWidth:     1,
						BorderColor:     "#ff0000",
						Gap:             2,
					},
					OnPress: bff.ActionData{
						Type: "downloadInvoice",
						Data: map[string]interface{}{
							"paymentId": paymentId,
						},
					},
				},
				Children: []bff.UISnippet{
					{
						Type: "Icon",
						Data: bff.IconData{
							Name:  "file-download",
							Size:  20,
							Color: "#ff0000",
						},
					},
					{
						Type: "Text",
						Data: bff.TextData{
							Text:      "Download Invoice",
							Color:     "#ff0000",
							FontSize:  16,
							FontWeight: "600",
						},
					},
				},
			},
			// Make Payment Button
			{
				Type: "TouchableOpacity",
				Data: bff.TouchableOpacityData{
					Style: bff.ViewData{
						Flex:             1,
						FlexDirection:   "row",
						BackgroundColor: "#ff0000",
						PaddingVertical: 16,
						BorderRadius:    12,
						AlignItems:      "center",
						JustifyContent:  "center",
						Gap:             8,
					},
					OnPress: bff.ActionData{
						Type: "makePayment",
						Data: map[string]interface{}{
							"paymentId": paymentId,
						},
					},
				},
				Children: []bff.UISnippet{
					{
						Type: "Icon",
						Data: bff.IconData{
							Name:  "credit-card",
							Size:  20,
							Color: "#fff",
						},
					},
					{
						Type: "Text",
						Data: bff.TextData{
							Text:      "Make Payment",
							Color:     "#fff",
							FontSize:  16,
							FontWeight: "600",
						},
					},
				},
			},
		},
	}
}