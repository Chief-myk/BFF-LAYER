package driver

import (
	"backend/bff"
	"strings"
	"github.com/gin-gonic/gin"
)

func PaymentScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		// Status Bar
		{
			Type: "STATUS_BAR",
			Data: bff.StatusBarData{
				BackgroundColor: "#ffffff",
				Style:           "dark",
			},
		},

		// Main Scroll View
		{
			Type: "SCROLL",
			Data: bff.ViewData{
				FlexGrow:        1,
				BackgroundColor: "#f5f7fa",
				PaddingTop:      16,
				PaddingBottom:   16,
			},
			Children: []bff.UISnippet{
				/* =======================
				   PAYMENT ALERT
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						FlexDirection:    "row",
						AlignItems:       "center",
						BackgroundColor:  "#E8F5E8",
						Padding:          20,
						BorderRadius:     16,
						MarginHorizontal: 16,
						MarginBottom:     16,
						BorderWidth:      1,
						BorderColor:      "#C8E6C9",
						ShadowColor:      "#000",
						ShadowOpacity:    0.05,
						ShadowRadius:     8,
						Elevation:        2,
					},
					Children: []bff.UISnippet{
						{
							Type: "ICON",
							Data: bff.IconData{
								Name:            "checkmark-circle",
								Size:            28,
								Color:           "#28A745",
								ContainerSize:   48,
								BorderRadius:    24,
								BackgroundColor: "#ffffff",
								Padding:         10,
							},
						},
						{
							Type: "VIEW",
							Data: bff.ViewData{
								Flex:        1,
								MarginLeft:  16,
								MarginRight: 16,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:         "Payment received for trip 4587",
										FontSize:     16,
										FontWeight:   "bold",
										Color:        "#1a1a1a",
										MarginBottom: 4,
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:         "₹5,200 credited to wallet",
										FontSize:     18,
										FontWeight:   "bold",
										Color:        "#28A745",
										MarginBottom: 4,
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:     "2 hours ago",
										FontSize: 14,
										Color:    "#666666",
										Opacity:  0.8,
									},
								},
							},
						},
						{
							Type: "ICON",
							Data: bff.IconData{
								Name:  "information-circle-outline",
								Size:  24,
								Color: "#007AFF",
							},
						},
					},
				},

				/* =======================
				   WALLET SUMMARY
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						PaddingHorizontal: 16,
						MarginBottom:      20,
					},
					Children: []bff.UISnippet{
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-between",
								AlignItems:     "center",
								MarginBottom:   16,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "Wallet Summary",
										FontSize:   22,
										FontWeight: "bold",
										Color:      "#1a237e",
									},
								},
								{
									Type: "VIEW",
									Data: bff.ViewData{
										FlexDirection:     "row",
										AlignItems:        "center",
										PaddingHorizontal: 12,
										PaddingVertical:   8,
										BackgroundColor:   "#f0f7ff",
										BorderRadius:      20,
										BorderWidth:       1,
										BorderColor:       "#007AFF20",
									},
									Children: []bff.UISnippet{
										{
											Type: "ICON",
											Data: bff.IconData{
												Name:  "refresh",
												Size:  16,
												Color: "#007AFF",
											},
										},
										{
											Type: "TEXT",
											Data: bff.TextData{
												Text:       "Refresh",
												FontSize:   14,
												FontWeight: "600",
												Color:      "#007AFF",
												MarginLeft: 6,
											},
										},
									},
								},
							},
						},
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-between",
								Gap:            12,
							},
							Children: []bff.UISnippet{
								walletCardEnhanced("wallet", "#28A745", "Current Balance", "₹4,520", "Available for use"),
								walletCardEnhanced("cash-multiple", "#007AFF", "Total Earnings", "₹8,500", "This month"),
								walletCardEnhanced("clock-outline", "#FF9500", "Pending", "₹2,300", "Awaiting clearance"),
							},
						},
					},
				},

				/* =======================
				   RECENT TRANSACTIONS
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						PaddingHorizontal: 16,
						MarginBottom:      24,
					},
					Children: []bff.UISnippet{
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-between",
								AlignItems:     "center",
								MarginBottom:   16,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "Recent Transactions",
										FontSize:   22,
										FontWeight: "bold",
										Color:      "#1a237e",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "See All",
										FontSize:   16,
										FontWeight: "600",
										Color:      "#007AFF",
									},
								},
							},
						},
						transactionCardEnhanced(
							"truck-check",
							true,
							"Commission Earned",
							"TRIP#4587 • Mumbai → Delhi",
							"₹2,500",
							"Success",
							"2024-09-15 • 10:30 AM",
						),
						transactionCardEnhanced(
							"bank-transfer",
							false,
							"Withdrawal",
							"Bank Transfer",
							"₹5,000",
							"Pending",
							"2024-09-13 • 09:45 AM",
						),
					},
				},

				/* =======================
				   TRIP WISE EARNINGS
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						PaddingHorizontal: 16,
						MarginBottom:      24,
					},
					Children: []bff.UISnippet{
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-between",
								AlignItems:     "center",
								MarginBottom:   16,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "Trip Wise Earnings",
										FontSize:   22,
										FontWeight: "bold",
										Color:      "#1a237e",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "See All",
										FontSize:   16,
										FontWeight: "600",
										Color:      "#007AFF",
									},
								},
							},
						},
						tripCardEnhanced(map[string]string{
							"id":       "TRIP#4587",
							"route":    "Mumbai → Delhi",
							"distance": "1,416 km",
							"amount":   "₹15,200",
							"status":   "Settled",
							"broker":   "ABC Logistics",
							"date":     "2024-09-15",
						}),
						tripCardEnhanced(map[string]string{
							"id":       "TRIP#4498",
							"route":    "Delhi → Kolkata",
							"distance": "1,533 km",
							"amount":   "₹16,800",
							"status":   "Pending",
							"broker":   "PQR Freight",
							"date":     "2024-09-13",
						}),
					},
				},

				/* =======================
				   SETTLEMENT DETAILS
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						PaddingHorizontal: 16,
						MarginBottom:      24,
					},
					Children: []bff.UISnippet{
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-between",
								AlignItems:     "center",
								MarginBottom:   16,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "Settlement Details",
										FontSize:   22,
										FontWeight: "bold",
										Color:      "#1a237e",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       "View All",
										FontSize:   16,
										FontWeight: "600",
										Color:      "#007AFF",
									},
								},
							},
						},
						settlementCardEnhanced(
							"SET#789456",
							"₹15,200",
							"ABC Logistics",
							"Bank Transfer",
							"2024-09-15",
							"Full payment received",
						),
						settlementCardEnhanced(
							"SET#789455",
							"₹8,500",
							"XYZ Transport",
							"UPI",
							"2024-09-14",
							"Commission deducted: ₹500",
						),
					},
				},

				/* =======================
				   EXTRA FEATURES
				======================= */
				{
					Type: "VIEW",
					Data: bff.ViewData{
						PaddingHorizontal: 16,
						MarginBottom:      32,
					},
					Children: []bff.UISnippet{
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:  "row",
								JustifyContent: "space-around",
								Gap:            12,
							},
							Children: []bff.UISnippet{
								featureButtonEnhanced("filter-variant", "Filter", "#666666"),
								featureButtonEnhanced("magnify", "Search", "#666666"),
								featureButtonEnhanced("file-export", "Export", "#666666"),
							},
						},
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "Payment",
		UI:     ui,
	}

	c.JSON(200, response)
}

// Enhanced Helper Functions
func walletCardEnhanced(icon string, color string, title string, value string, subtitle string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Flex:            1,
			BackgroundColor: "#ffffff",
			Padding:         20,
			BorderRadius:    16,
			AlignItems:      "center",
			ShadowColor:     "#000",
			ShadowOpacity:   0.08,
			ShadowRadius:    12,
			Elevation:       3,
			BorderWidth:     1,
			BorderColor:     color + "20",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Width:           60,
					Height:          60,
					BorderRadius:    30,
					BackgroundColor: color + "15",
					AlignItems:      "center",
					JustifyContent:  "center",
					MarginBottom:    16,
				},
				Children: []bff.UISnippet{
					{
						Type: "ICON",
						Data: bff.IconData{
							Name:  icon,
							Size:  28,
							Color: color,
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         title,
					FontSize:     14,
					FontWeight:   "500",
					Color:        "#666666",
					MarginBottom: 8,
					TextAlign:    "center",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         value,
					FontSize:     20,
					FontWeight:   "bold",
					Color:        "#1a237e",
					MarginBottom: 4,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:      subtitle,
					FontSize:  12,
					Color:     "#999999",
					TextAlign: "center",
				},
			},
		},
	}
}

func transactionCardEnhanced(
	icon string,
	isCredit bool,
	title string,
	subtitle string,
	amount string,
	status string,
	time string,
) bff.UISnippet {
	color := "#DC3545"
	statusBg := "#FFE5E5"
	statusColor := "#DC3545"
	if isCredit {
		color = "#28A745"
		statusBg = "#E8F5E8"
		statusColor = "#28A745"
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#ffffff",
			BorderRadius:    16,
			Padding:         20,
			MarginBottom:    12,
			ShadowColor:     "#000",
			ShadowOpacity:   0.05,
			ShadowRadius:    8,
			Elevation:       2,
			BorderWidth:     1,
			BorderColor:     "#f0f0f0",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					AlignItems:    "center",
					MarginBottom:  12,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							Width:           48,
							Height:          48,
							BorderRadius:    24,
							BackgroundColor: color + "15",
							AlignItems:      "center",
							JustifyContent:  "center",
							MarginRight:     16,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  icon,
									Size:  24,
									Color: color,
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							Flex: 1,
						},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:         title,
									FontSize:     16,
									FontWeight:   "600",
									Color:        "#1a237e",
									MarginBottom: 4,
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:     subtitle,
									FontSize: 14,
									Color:    "#666666",
								},
							},
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       amount,
							FontSize:   18,
							FontWeight: "bold",
							Color:      color,
						},
					},
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "center",
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							PaddingHorizontal: 12,
							PaddingVertical:   6,
							BackgroundColor:   statusBg,
							BorderRadius:      20,
							BorderWidth:       1,
							BorderColor:       statusColor + "30",
						},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       status,
									FontSize:   12,
									FontWeight: "600",
									Color:      statusColor,
								},
							},
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:     time,
							FontSize: 12,
							Color:    "#999999",
						},
					},
				},
			},
		},
	}
}

func tripCardEnhanced(data map[string]string) bff.UISnippet {
	statusColor := "#28A745"
	statusBg := "#E8F5E8"
	if data["status"] == "Pending" {
		statusColor = "#FF9500"
		statusBg = "#FFF3E0"
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#ffffff",
			BorderRadius:    20,
			Padding:         24,
			MarginBottom:    16,
			ShadowColor:     "#000",
			ShadowOpacity:   0.08,
			ShadowRadius:    12,
			Elevation:       3,
			BorderWidth:     1,
			BorderColor:     "#f0f0f0",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "center",
					MarginBottom:   16,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:         data["id"],
									FontSize:     12,
									FontWeight:   "600",
									Color:        "#666666",
									MarginBottom: 4,
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       data["route"],
									FontSize:   18,
									FontWeight: "bold",
									Color:      "#1a237e",
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							PaddingHorizontal: 16,
							PaddingVertical:   8,
							BackgroundColor:   statusBg,
							BorderRadius:      20,
							BorderWidth:       1,
							BorderColor:       statusColor + "30",
						},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       data["status"],
									FontSize:   14,
									FontWeight: "600",
									Color:      statusColor,
								},
							},
						},
					},
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					MarginBottom:   16,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           8,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  "map-outline",
									Size:  18,
									Color: "#666666",
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:     data["distance"],
									FontSize: 14,
									Color:    "#666666",
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           8,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  "business-outline",
									Size:  18,
									Color: "#666666",
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:     data["broker"],
									FontSize: 14,
									Color:    "#666666",
								},
							},
						},
					},
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					AlignItems:     "center",
					JustifyContent: "space-between",
					PaddingTop:     16,
					BorderTopWidth: 1,
					BorderTopColor: "#f0f0f0",
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           8,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  "calendar-outline",
									Size:  18,
									Color: "#666666",
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:     data["date"],
									FontSize: 14,
									Color:    "#666666",
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:         "Earnings",
									FontSize:     12,
									FontWeight:   "500",
									Color:        "#666666",
									MarginBottom: 4,
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       data["amount"],
									FontSize:   22,
									FontWeight: "bold",
									Color:      "#28A745",
								},
							},
						},
					},
				},
			},
		},
	}
}

func settlementCardEnhanced(
	id string,
	amount string,
	broker string,
	mode string,
	date string,
	remarks string,
) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#ffffff",
			BorderRadius:    16,
			Padding:         20,
			MarginBottom:    12,
			ShadowColor:     "#000",
			ShadowOpacity:   0.05,
			ShadowRadius:    8,
			Elevation:       2,
			BorderWidth:     1,
			BorderColor:     "#f0f0f0",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "center",
					MarginBottom:   16,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           8,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  "receipt-outline",
									Size:  20,
									Color: "#666666",
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       id,
									FontSize:   16,
									FontWeight: "600",
									Color:      "#1a237e",
								},
							},
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       amount,
							FontSize:   20,
							FontWeight: "bold",
							Color:      "#28A745",
						},
					},
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Gap: 12,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           12,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:           40,
									Height:          40,
									BorderRadius:    20,
									BackgroundColor: "#f0f7ff",
									AlignItems:      "center",
									JustifyContent:  "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON",
										Data: bff.IconData{
											Name:  "business-outline",
											Size:  20,
											Color: "#007AFF",
										},
									},
								},
							},
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex: 1,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Broker",
											FontSize:     12,
											FontWeight:   "500",
											Color:        "#999999",
											MarginBottom: 2,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       broker,
											FontSize:   16,
											FontWeight: "600",
											Color:      "#1a237e",
										},
									},
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           12,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:           40,
									Height:          40,
									BorderRadius:    20,
									BackgroundColor: "#f0f7ff",
									AlignItems:      "center",
									JustifyContent:  "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON",
										Data: bff.IconData{
											Name:  "card-outline",
											Size:  20,
											Color: "#007AFF",
										},
									},
								},
							},
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex: 1,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Payment Mode",
											FontSize:     12,
											FontWeight:   "500",
											Color:        "#999999",
											MarginBottom: 2,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       mode,
											FontSize:   16,
											FontWeight: "600",
											Color:      "#1a237e",
										},
									},
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           12,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:           40,
									Height:          40,
									BorderRadius:    20,
									BackgroundColor: "#f0f7ff",
									AlignItems:      "center",
									JustifyContent:  "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON",
										Data: bff.IconData{
											Name:  "calendar-outline",
											Size:  20,
											Color: "#007AFF",
										},
									},
								},
							},
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex: 1,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Date",
											FontSize:     12,
											FontWeight:   "500",
											Color:        "#999999",
											MarginBottom: 2,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       date,
											FontSize:   16,
											FontWeight: "600",
											Color:      "#1a237e",
										},
									},
								},
							},
						},
					},
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							AlignItems:    "center",
							Gap:           12,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:           40,
									Height:          40,
									BorderRadius:    20,
									BackgroundColor: "#f0f7ff",
									AlignItems:      "center",
									JustifyContent:  "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON",
										Data: bff.IconData{
											Name:  "document-text-outline",
											Size:  20,
											Color: "#007AFF",
										},
									},
								},
							},
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex: 1,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Remarks",
											FontSize:     12,
											FontWeight:   "500",
											Color:        "#999999",
											MarginBottom: 2,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       remarks,
											FontSize:   16,
											FontWeight: "600",
											Color:      "#666666",
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
}

func featureButtonEnhanced(icon string, label string, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TOUCHABLE_OPACITY",
		Data: bff.TouchableOpacityData{
			Style: bff.ViewData{
				Flex:              1,
				FlexDirection:     "row",
				AlignItems:        "center",
				JustifyContent:    "center",
				BackgroundColor:   "#ffffff",
				PaddingVertical:   16,
				PaddingHorizontal: 8,
				BorderRadius:      12,
				ShadowColor:       "#000",
				ShadowOpacity:     0.05,
				ShadowRadius:      8,
				Elevation:         2,
				BorderWidth:       1,
				BorderColor:       "#f0f0f0",
			},
			OnPress: bff.ActionData{
				Type:  "ACTION",
				Value: "button_" + strings.ToLower(label),
			},
		},
		Children: []bff.UISnippet{
			{
				Type: "ICON",
				Data: bff.IconData{
					Name:  icon,
					Size:  20,
					Color: color,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       label,
					FontSize:   14,
					FontWeight: "600",
					Color:      color,
					MarginLeft: 8,
				},
			},
		},
	}
}

// Keep original functions for compatibility (not removing any words)
func card(bg string, layout string, children ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type: layout,
		Data: bff.ViewData{
			BackgroundColor: bg,
			Padding:         16,
			BorderRadius:    12,
		},
		Children: children,
	}
}

func sectionHeader(title string, action string) bff.UISnippet {
	return row(
		text(title, 20, true, "#1a1a1a"),
		text(action, 14, false, "#007AFF"),
	)
}

func walletCard(icon string, color string, title string, value string, subtitle string) bff.UISnippet {
	return card(
		"#ffffff",
		"COLUMN",
		iconAtom(icon, color),
		text(title, 14, false, "#666"),
		text(value, 16, true, "#1a1a1a"),
		text(subtitle, 12, false, "#666"),
	)
}

func transactionCard(
	icon string,
	isCredit bool,
	title string,
	trip string,
	route string,
	amount string,
	status string,
	time string,
) bff.UISnippet {
	color := "#DC3545"
	if isCredit {
		color = "#28A745"
	}

	return card(
		"#ffffff",
		"COLUMN",
		row(
			iconAtom(icon, color),
			column(
				text(title, 16, true, "#1a1a1a"),
				text(trip, 14, false, "#666"),
				text(time, 12, false, "#999"),
			),
			text(amount, 16, true, color),
		),
	)
}

func settlementCard(
	id string,
	amount string,
	broker string,
	mode string,
	date string,
	remarks string,
) bff.UISnippet {
	return card(
		"#ffffff",
		"COLUMN",
		row(
			text(id, 16, true, "#1a1a1a"),
			text(amount, 16, true, "#28A745"),
		),
		column(
			text("Broker: "+broker, 14, false, "#666"),
			text("Payment Mode: "+mode, 14, false, "#666"),
			text("Date: "+date, 14, false, "#666"),
			text("Remarks: "+remarks, 14, false, "#666"),
		),
	)
}

func featureButton(icon string, label string) bff.UISnippet {
	return bff.UISnippet{
		Type: "ROW",
		Children: []bff.UISnippet{
			iconAtom(icon, "#666"),
			textAtom(label, "#666", false),
		},
	}
}

func row(children ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type:     "ROW",
		Children: children,
	}
}

func column(children ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type:     "COLUMN",
		Children: children,
	}
}

func textAtom(value string, color string, bold bool) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:  value,
			Color: color,
			Bold:  bold,
		},
	}
}

func text(value string, size int, bold bool, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:     value,
			Color:    color,
			Bold:     bold,
			FontSize: size,
		},
	}
}
