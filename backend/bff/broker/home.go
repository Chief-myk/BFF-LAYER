package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func HomeScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		// SafeAreaView
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
						PaddingVertical: 16,
						BackgroundColor: "#FFFFFF",
						BorderBottomWidth: 1,
						BorderColor:      "#F0F0F0",
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
										Text:      "LogiBroker",
										FontSize:  24,
										FontWeight: "bold",
										Color:     "#ff0000",
										MarginBottom: 4,
									},
								},
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Welcome back! 👋",
										FontSize:  14,
										Color:     "#666",
									},
								},
							},
						},
					},
				},
				// ScrollView
				{
					Type: "ScrollView",
					Data: bff.ViewData{
						Flex: 1,
					},
					Children: []bff.UISnippet{
						// Stats Section
						{
							Type: "View",
							Data: bff.ViewData{
								PaddingHorizontal: 20,
								PaddingVertical: 16,
							},
							Children: []bff.UISnippet{
								// Horizontal ScrollView for stats
								{
									Type: "ScrollView",
									Data: bff.ViewData{
										Horizontal: true,
									},
									Children: []bff.UISnippet{
										// Stat Card 1 - Active Loads
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width:  160,
													Height: 120,
													BorderRadius: 16,
													MarginRight: 12,
													Overflow: "hidden",
													ShadowColor: "#000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 2,
													ShadowOpacity: 0.1,
													ShadowRadius: 8,
													Elevation: 4,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/loads",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Flex: 1,
														BackgroundColor: "#ff0000",
														Padding: 16,
														AlignItems: "center",
														JustifyContent: "center",
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "package-variant",
																Size:  24,
																Color: "#fff",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "12",
																FontSize:  28,
																FontWeight: "bold",
																Color:     "#FFFFFF",
																MarginTop: 8,
																MarginBottom: 4,
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Active Loads",
																FontSize:  12,
																Color:     "#FFFFFF",
																TextAlign: "center",
																FontWeight: "500",
															},
														},
													},
												},
											},
										},
										// Stat Card 2 - Pending Bids
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width:  160,
													Height: 120,
													BorderRadius: 16,
													MarginRight: 12,
													BackgroundColor: "#FFFFFF",
													BorderWidth: 1,
													BorderColor: "#F0F0F0",
													ShadowColor: "#000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 2,
													ShadowOpacity: 0.1,
													ShadowRadius: 8,
													Elevation: 4,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/bids",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Flex: 1,
														BackgroundColor: "#ff0000",
														Padding: 16,
														AlignItems: "center",
														JustifyContent: "center",
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "gavel",
																Size:  24,
																Color: "#fff",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "8",
																FontSize:  28,
																FontWeight: "bold",
																Color:     "#FFFFFF",
																MarginTop: 8,
																MarginBottom: 4,
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Pending Bids",
																FontSize:  12,
																Color:     "#FFFFFF",
																TextAlign: "center",
																FontWeight: "500",
															},
														},
														// Badge
														{
															Type: "View",
															Data: bff.ViewData{
																Position: "absolute",
																Top:      12,
																Right:    12,
																BackgroundColor: "#FFF0F0",
																PaddingHorizontal: 8,
																PaddingVertical: 2,
																BorderRadius: 10,
															},
															Children: []bff.UISnippet{
																{
																	Type: "Text",
																	Data: bff.TextData{
																		Text:      "+2 new",
																		FontSize:  10,
																		Color:     "#ff0000",
																		FontWeight: "600",
																	},
																},
															},
														},
													},
												},
											},
										},
										// Stat Card 3 - Live Trips
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width:  160,
													Height: 120,
													BorderRadius: 16,
													MarginRight: 12,
													BackgroundColor: "#FFFFFF",
													BorderWidth: 1,
													BorderColor: "#F0F0F0",
													ShadowColor: "#000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 2,
													ShadowOpacity: 0.1,
													ShadowRadius: 8,
													Elevation: 4,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/trips",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Flex: 1,
														BackgroundColor: "#ff0000",
														Padding: 16,
														AlignItems: "center",
														JustifyContent: "center",
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "map-marker-path",
																Size:  24,
																Color: "#fff",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "5",
																FontSize:  28,
																FontWeight: "bold",
																Color:     "#FFFFFF",
																MarginTop: 8,
																MarginBottom: 4,
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Live Trips",
																FontSize:  12,
																Color:     "#FFFFFF",
																TextAlign: "center",
																FontWeight: "500",
															},
														},
													},
												},
											},
										},
										// Stat Card 4 - Pending Payments
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width:  160,
													Height: 120,
													BorderRadius: 16,
													MarginRight: 12,
													BackgroundColor: "#FFFFFF",
													BorderWidth: 1,
													BorderColor: "#F0F0F0",
													ShadowColor: "#000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 2,
													ShadowOpacity: 0.1,
													ShadowRadius: 8,
													Elevation: 4,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/payments",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Flex: 1,
														BackgroundColor: "#DC2616",
														Padding: 16,
														AlignItems: "center",
														JustifyContent: "center",
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "cash",
																Size:  24,
																Color: "#fff",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "3",
																FontSize:  28,
																FontWeight: "bold",
																Color:     "#FFFFFF",
																MarginTop: 8,
																MarginBottom: 4,
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Pending Payments",
																FontSize:  12,
																Color:     "#FFFFFF",
																TextAlign: "center",
																FontWeight: "500",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "₹78,300",
																FontSize:  12,
																Color:     "#ffffff",
																FontWeight: "bold",
																MarginTop: 4,
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
						// Quick Actions Section
						{
							Type: "View",
							Data: bff.ViewData{
								PaddingHorizontal: 20,
								PaddingVertical: 16,
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Quick Actions",
										FontSize:  22,
										FontWeight: "bold",
										Color:     "#1A1A1A",
										MarginBottom: 16,
									},
								},
								{
									Type: "View",
									Data: bff.ViewData{
										FlexDirection: "row",
										Gap: 12,
									},
									Children: []bff.UISnippet{
										// Add Load Button
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Flex: 1,
													BorderRadius: 16,
													Overflow: "hidden",
													ShadowColor: "#ff0000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 4,
													ShadowOpacity: 0.3,
													ShadowRadius: 8,
													Elevation: 4,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/add-load",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														FlexDirection: "row",
														AlignItems: "center",
														JustifyContent: "center",
														PaddingVertical: 18,
														BorderRadius: 16,
														Gap: 8,
														BackgroundColor: "#ff0000",
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "package-variant-plus",
																Size:  28,
																Color: "#fff",
															},
														},
														{
															Type: "Text",
															Data: bff.TextData{
																Text:      "Add Load",
																FontSize:  16,
																FontWeight: "600",
																Color:     "#FFFFFF",
															},
														},
													},
												},
											},
										},
										// Add Truck Button
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Flex: 1,
													BackgroundColor: "#FFFFFF",
													FlexDirection: "row",
													AlignItems: "center",
													JustifyContent: "center",
													PaddingVertical: 18,
													BorderRadius: 16,
													Gap: 8,
													BorderWidth: 2,
													BorderColor: "#ff0000",
													ShadowColor: "#000",
													ShadowOffsetX: 0,
													ShadowOffsetY: 2,
													ShadowOpacity: 0.1,
													ShadowRadius: 4,
													Elevation: 2,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/add-truck",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "Icon",
													Data: bff.IconData{
														Name:  "truck-plus",
														Size:  28,
														Color: "#ff0000",
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Add Truck",
														FontSize:  16,
														FontWeight: "600",
														Color:     "#ff0000",
													},
												},
											},
										},
									},
								},
							},
						},
						// Quick Tools Section
						{
							Type: "View",
							Data: bff.ViewData{
								PaddingHorizontal: 20,
								PaddingVertical: 16,
							},
							Children: []bff.UISnippet{
								{
									Type: "Text",
									Data: bff.TextData{
										Text:      "Quick Tools",
										FontSize:  22,
										FontWeight: "bold",
										Color:     "#1A1A1A",
										MarginBottom: 16,
									},
								},
								{
									Type: "View",
									Data: bff.ViewData{
										FlexDirection: "row",
										FlexWrap: "wrap",
										JustifyContent: "space-between",
										Gap: 12,
									},
									Children: []bff.UISnippet{
										// My Trucks
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/my-trucks",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#ff0000",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "truck",
																Size:  20,
																Color: "#fff",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "My Trucks",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
										// Create Load
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/create-load",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#F8F9FA",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "package-variant-plus",
																Size:  20,
																Color: "#ff0000",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Create Load",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
										// Wallet
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/wallet",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#F8F9FA",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "wallet-outline",
																Size:  20,
																Color: "#ff0000",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Wallet",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
										// Documents
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/documents",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#F8F9FA",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "file-document",
																Size:  20,
																Color: "#ff0000",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Documents",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
										// Analytics
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/analytics",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#F8F9FA",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "analytics-outline",
																Size:  20,
																Color: "#ff0000",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Analytics",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
										// Place Bid
										{
											Type: "TouchableOpacity",
											Data: bff.TouchableOpacityData{
												Style: bff.ViewData{
													Width: "23%",
													AlignItems: "center",
													MarginBottom: 16,
												},
												OnPress: bff.ActionData{
													Type: "navigate",
													To:   "/place-bid",
												},
											},
											Children: []bff.UISnippet{
												{
													Type: "View",
													Data: bff.ViewData{
														Width: 60,
														Height: 60,
														BackgroundColor: "#F8F9FA",
														BorderRadius: 16,
														AlignItems: "center",
														JustifyContent: "center",
														MarginBottom: 8,
														BorderWidth: 1,
														BorderColor: "#F0F0F0",
														ShadowColor: "#000",
														ShadowOffsetX: 0,
														ShadowOffsetY: 1,
														ShadowOpacity: 0.1,
														ShadowRadius: 2,
														Elevation: 1,
													},
													Children: []bff.UISnippet{
														{
															Type: "Icon",
															Data: bff.IconData{
																Name:  "hammer",
																Size:  20,
																Color: "#ff0000",
															},
														},
													},
												},
												{
													Type: "Text",
													Data: bff.TextData{
														Text:      "Place Bid",
														FontSize:  12,
														Color:     "#333",
														TextAlign: "center",
														FontWeight: "500",
													},
												},
											},
										},
									},
								},
							},
						},
						// Bottom Spacer
						{
							Type: "View",
							Data: bff.ViewData{
								Height: 20,
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
		Screen: "home",
		UI:     ui,
		Data: map[string]interface{}{
			"stats": map[string]interface{}{
				"activeLoads":     12,
				"pendingBids":     8,
				"liveTrips":       5,
				"pendingPayments": 3,
				"paymentAmount":   "₹78,300",
				"newBids":         "+2 new",
			},
			"quickActions": []map[string]interface{}{
				{
					"id":    1,
					"icon":  "package-variant-plus",
					"title": "Add Load",
					"color": "#ff0000",
				},
				{
					"id":    2,
					"icon":  "truck-plus",
					"title": "Add Truck",
					"color": "#ff0000",
				},
			},
			"quickTools": []map[string]interface{}{
				{
					"id":    1,
					"icon":  "truck",
					"title": "My Trucks",
					"route": "/my-trucks",
				},
				{
					"id":    2,
					"icon":  "package-variant-plus",
					"title": "Create Load",
					"route": "/create-load",
				},
				{
					"id":    3,
					"icon":  "wallet-outline",
					"title": "Wallet",
					"route": "/wallet",
				},
				{
					"id":    4,
					"icon":  "file-document",
					"title": "Documents",
					"route": "/documents",
				},
				{
					"id":    5,
					"icon":  "analytics-outline",
					"title": "Analytics",
					"route": "/analytics",
				},
				{
					"id":    6,
					"icon":  "hammer",
					"title": "Place Bid",
					"route": "/place-bid",
				},
			},
		},
	}

	c.JSON(200, response)
}