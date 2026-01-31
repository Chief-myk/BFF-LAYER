package driver

import (
	"backend/bff"
	"fmt"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

func HomeScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Fetch or generate home data
	homeData := getHomeScreenData()

	// Generate UI with new design
	ui := generateModernUI(homeData)

	response := bff.ScreenResponse{
		Status:  "success",
		Screen:  "Home",
		Data:    homeData,
		UI:      ui,
		Message: "Welcome back!",
	}

	c.JSON(200, response)
}

func getHomeScreenData() bff.HomeScreenData {
	return bff.HomeScreenData{
		IsTripStarted:   false,
		LocationSharing: false,
		TripStatus:      "not_started",
		DocumentsUploaded: map[string]bool{
			"eWayBill":      false,
			"invoice":       false,
			"vehicleRC":     true,
			"driverLicense": true,
			"insurance":     true,
			"pollutionCert": true,
		},
		ActiveTrip: map[string]string{
			"id":               "TRK789012",
			"tripNumber":       "TRIP-2024-001",
			"origin":           "Mumbai Port",
			"originCity":       "Mumbai, MH",
			"destination":      "Delhi Logistics Park",
			"destinationCity":  "Delhi, DL",
			"cargo":            "Electronics & Appliances",
			"cargoType":        "General Goods",
			"weight":           "15 Tons",
			"payment":          "₹68,500",
			"advancePaid":      "₹20,000",
			"balanceDue":       "₹48,500",
			"distance":         "1,412 km",
			"estimatedTime":    "26 hrs",
			"startTime":        "Today, 10:00 AM",
			"brokerName":       "Sharma Logistics Pvt. Ltd.",
			"brokerPhone":      "+91 9876543210",
			"brokerRating":     "4.8",
			"senderName":       "TechCorp India Ltd.",
			"senderPhone":      "+91 9876543211",
			"receiverName":     "Metro Retail Chains",
			"receiverPhone":    "+91 9876543212",
			"vehicleNumber":    "MH01AB1234",
			"driverName":       "Rajesh Kumar",
			"currentLocation":  "Pune, MH",
			"progress":         "35",
			"estimatedArrival": "Tomorrow, 12:00 PM",
			"fuelLevel":        "65",
			"vehicleHealth":    "Good",
			"tripScore":        "92",
		},
		QuickActions: []bff.QuickAction{
			{ID: 1, Icon: "play-circle", Title: "Start Trip", Color: "#4CAF50"},
			{ID: 2, Icon: "search", Title: "Find Loads", Color: "#2196F3"},
			{ID: 3, Icon: "wallet", Title: "Payments", Color: "#FF9800"},
			{ID: 4, Icon: "document-text", Title: "Docs", Color: "#9C27B0"},
			{ID: 5, Icon: "car", Title: "My Vehicle", Color: "#607D8B"},
			{ID: 6, Icon: "stats-chart", Title: "Analytics", Color: "#00BCD4"},
		},
		RecentActivities: []bff.RecentActivity{
			{ID: 1, Type: "payment", Message: "Advance received ₹20,000", Time: "2 hours ago", Icon: "checkmark-circle", Color: "#4CAF50"},
			{ID: 2, Type: "assignment", Message: "New trip assigned: Mumbai to Delhi", Time: "4 hours ago", Icon: "car", Color: "#2196F3"},
			{ID: 3, Type: "maintenance", Message: "Vehicle service due in 500 km", Time: "1 day ago", Icon: "construct", Color: "#FF9800"},
			{ID: 4, Type: "alert", Message: "Toll payment reminder", Time: "2 days ago", Icon: "alert-circle", Color: "#F44336"},
		},
	}
}

func generateModernUI(data bff.HomeScreenData) []bff.UISnippet {
	return []bff.UISnippet{
		// Status Bar
		{
			Type: "STATUS_BAR",
			Data: bff.StatusBarData{
				BackgroundColor: "#1a237e",
				Style:           "light",
			},
		},
		// Main Scroll View
		{
			Type: "SCROLL",
			Data: bff.ViewData{
				FlexGrow:        1,
				BackgroundColor: "#f5f7fa",
				PaddingTop:      20,
				PaddingBottom:   20,
			},
			Children: []bff.UISnippet{
				// Header Section
				headerSection(),

				// Quick Actions
				quickActionsSection(data.QuickActions),

				// Documents Status
				documentsSection(data.DocumentsUploaded, data.IsTripStarted),

				// Active Trip Card
				activeTripSection(data),

				// Recent Activities
				recentActivitiesSection(data.RecentActivities),

				// Performance Metrics (if trip started)
				performanceMetricsSection(data),
			},
		},
	}
}

func headerSection() bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
			PaddingBottom:     16,
		},
		Children: []bff.UISnippet{
			// Greeting and Profile
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "center",
					MarginBottom:   8,
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:     "Good Morning,",
									FontSize: 14,
									Color:    "#666",
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       "Rajesh Kumar",
									FontSize:   24,
									FontWeight: "bold",
									Color:      "#1a237e",
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
								Type: "ICON_BUTTON",
								Data: bff.IconButtonData{
									Icon:  "notifications-outline",
									Size:  24,
									Color: "#1a237e",
									OnPress: bff.ActionData{
										Type: "NAVIGATE",
										To:   "/notifications",
									},
								},
							},
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:           48,
									Height:          48,
									BorderRadius:    24,
									BackgroundColor: "#2196F3",
									AlignItems:      "center",
									JustifyContent:  "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "RK",
											FontSize:   18,
											FontWeight: "bold",
											Color:      "#fff",
										},
									},
								},
							},
						},
					},
				},
			},

			// Stats Overview
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:   "row",
					JustifyContent:  "space-between",
					BackgroundColor: "#fff",
					BorderRadius:    16,
					Padding:         16,
					ShadowColor:     "#000",
					ShadowOpacity:   0.1,
					ShadowRadius:    4,
					Elevation:       2,
				},
				Children: []bff.UISnippet{
					statItem("Trips", "12", "this month", "#4CAF50", "car-outline"),
					statItem("Earnings", "₹2.4L", "current month", "#FF9800", "wallet-outline"),
					statItem("Rating", "4.8", "driver score", "#2196F3", "star-outline"),
				},
			},
		},
	}
}

func statItem(label, value, subtext, color, icon string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			AlignItems: "center",
			Flex:       1,
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					AlignItems:    "center",
					Gap:           8,
					MarginBottom:  4,
				},
				Children: []bff.UISnippet{
					{
						Type: "ICON",
						Data: bff.IconData{
							Name:  icon,
							Size:  16,
							Color: color,
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       label,
							FontSize:   12,
							Color:      "#666",
							FontWeight: "500",
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       value,
					FontSize:   20,
					FontWeight: "bold",
					Color:      "#1a237e",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:     subtext,
					FontSize: 10,
					Color:    "#999",
				},
			},
		},
	}
}

func quickActionsSection(actions []bff.QuickAction) bff.UISnippet {
	var actionChildren []bff.UISnippet

	for _, action := range actions {
		actionChildren = append(actionChildren, bff.UISnippet{
			Type: "PRESSABLE_CARD",
			Data: bff.PressableCardData{
				CardData: bff.CardData{
					BackgroundColor: action.Color + "15", // 15 = 8% opacity
					Padding:         16,
					BorderRadius:    12,
					BorderWidth:     1,
					BorderColor:     action.Color + "30", // 30 = 18% opacity
					OnPress: bff.ActionData{
						Type:  "ACTION",
						Value: action.Title,
						Url:   "/bff/driver/home/action",
					},
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							AlignItems: "center",
							Gap:        8,
						},
						Children: []bff.UISnippet{
							{
								Type: "ICON",
								Data: bff.IconData{
									Name:  action.Icon,
									Size:  24,
									Color: action.Color,
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       action.Title,
									FontSize:   12,
									FontWeight: "600",
									Color:      "#1a237e",
									TextAlign:  "center",
								},
							},
						},
					},
				},
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
			MarginBottom:      20,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         "Quick Actions",
					FontSize:     18,
					FontWeight:   "bold",
					Color:        "#1a237e",
					MarginBottom: 12,
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					FlexWrap:      "wrap",
					Gap:           12,
				},
				Children: actionChildren,
			},
		},
	}
}

func documentsSection(docs map[string]bool, isTripStarted bool) bff.UISnippet {
	if isTripStarted {
		return bff.UISnippet{}
	}

	var requiredDocs []string
	var uploadedDocs []string

	for doc, uploaded := range docs {
		if !uploaded {
			requiredDocs = append(requiredDocs, doc)
		} else {
			uploadedDocs = append(uploadedDocs, doc)
		}
	}

	if len(requiredDocs) == 0 {
		return bff.UISnippet{}
	}

	var docItems []bff.UISnippet

	for _, doc := range requiredDocs {
		docItems = append(docItems, bff.UISnippet{
			Type: "VIEW",
			Data: bff.ViewData{
				FlexDirection:   "row",
				AlignItems:      "center",
				JustifyContent:  "space-between",
				Padding:         12,
				BackgroundColor: "#fff",
				BorderRadius:    8,
				BorderWidth:     1,
				BorderColor:     "#ffebee",
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
							Type: "ICON",
							Data: bff.IconData{
								Name:            getDocumentIcon(doc),
								Size:            20,
								Color:           "#F44336",
								ContainerSize:   36,
								BorderRadius:    18,
								BackgroundColor: "#ffebee",
							},
						},
						{
							Type: "VIEW",
							Data: bff.ViewData{},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:       getDocumentName(doc),
										FontSize:   14,
										FontWeight: "600",
										Color:      "#1a237e",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text:     "Required for trip start",
										FontSize: 12,
										Color:    "#F44336",
									},
								},
							},
						},
					},
				},
				{
					Type: "BUTTON",
					Data: bff.ButtonData{
						Text: "Upload",
						Style: bff.ViewData{
							PaddingHorizontal: 16,
							PaddingVertical:   8,
							BorderRadius:      6,
							BackgroundColor:   "#F44336",
						},
						Action: bff.ActionData{
							Type:  "ACTION",
							Value: "UPLOAD_DOCUMENT",
							Url:   "/bff/driver/home/action",
							Data:  map[string]interface{}{"documentType": doc},
						},
					},
				},
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
			MarginBottom:      20,
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					JustifyContent: "space-between",
					AlignItems:     "center",
					MarginBottom:   12,
				},
				Children: []bff.UISnippet{
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       "Required Documents",
							FontSize:   18,
							FontWeight: "bold",
							Color:      "#1a237e",
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:     fmt.Sprintf("%d/%d uploaded", len(uploadedDocs), len(docs)),
							FontSize: 14,
							Color:    "#666",
						},
					},
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Gap: 8,
				},
				Children: docItems,
			},
			{
				Type: "BUTTON",
				Data: bff.ButtonData{
					Text:     "Start Trip After Upload",
					Disabled: len(requiredDocs) > 0,
					Style: bff.ViewData{
						MarginTop:       12,
						PaddingVertical: 16,
						BorderRadius:    12,
						BackgroundColor: "#4CAF50",
					},
					Action: bff.ActionData{
						Type:  "ACTION",
						Value: "START_TRIP",
						Url:   "/bff/driver/home/action",
					},
				},
			},
		},
	}
}

func activeTripSection(data bff.HomeScreenData) bff.UISnippet {
	trip := data.ActiveTrip

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
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
							Text:       "Active Trip",
							FontSize:   20,
							FontWeight: "bold",
							Color:      "#1a237e",
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
								Type: "TEXT",
								Data: bff.TextData{
									Text:              trip["id"],
									FontSize:          12,
									FontWeight:        "600",
									PaddingHorizontal: 10,
									PaddingVertical:   4,
									BackgroundColor:   "#e8eaf6",
									Color:             "#3949ab",
									BorderRadius:      12,
								},
							},
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:              data.TripStatus,
									FontSize:          12,
									FontWeight:        "600",
									PaddingHorizontal: 10,
									PaddingVertical:   4,
									BackgroundColor:   getStatusColor(data.TripStatus) + "20",
									Color:             getStatusColor(data.TripStatus),
									BorderRadius:      12,
								},
							},
						},
					},
				},
			},

			// Trip Card
			{
				Type: "VIEW",
				Data: bff.ViewData{
					BackgroundColor: "#fff",
					BorderRadius:    20,
					Padding:         20,
					ShadowColor:     "#000",
					ShadowOpacity:   0.08,
					ShadowRadius:    12,
					Elevation:       3,
				},
				Children: []bff.UISnippet{
					// Route with Visual Progress
					{
						Type: "VIEW",
						Data: bff.ViewData{
							MarginBottom: 20,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									FlexDirection: "row",
									AlignItems:    "center",
									MarginBottom:  16,
								},
								Children: []bff.UISnippet{
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:           40,
											Height:          40,
											BorderRadius:    20,
											BackgroundColor: "#4CAF50",
											AlignItems:      "center",
											JustifyContent:  "center",
											MarginRight:     12,
										},
										Children: []bff.UISnippet{
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:  "location",
													Size:  20,
													Color: "#fff",
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
													Text:       trip["origin"],
													FontSize:   16,
													FontWeight: "600",
													Color:      "#1a237e",
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:     trip["originCity"],
													FontSize: 14,
													Color:    "#666",
												},
											},
										},
									},
								},
							},

							// Progress Line
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Height:          60,
									Width:           4,
									BackgroundColor: "#e0e0e0",
									MarginLeft:      18,
									Position:        "relative",
								},
								Children: []bff.UISnippet{
									// Progress Fill
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Height:          trip["progress"] + "%",
											Width:           4,
											BackgroundColor: "#4CAF50",
											Position:        "absolute",
											Top:             0,
											Left:            0,
										},
									},
									// Current Location Indicator
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:           16,
											Height:          16,
											BorderRadius:    8,
											BackgroundColor: "#4CAF50",
											Position:        "absolute",
											Top:             0,
											Left:            -6,
											BorderWidth:     3,
											BorderColor:     "#fff",
										},
									},
								},
							},

							{
								Type: "VIEW",
								Data: bff.ViewData{
									FlexDirection: "row",
									AlignItems:    "center",
									MarginTop:     16,
								},
								Children: []bff.UISnippet{
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:           40,
											Height:          40,
											BorderRadius:    20,
											BackgroundColor: "#F44336",
											AlignItems:      "center",
											JustifyContent:  "center",
											MarginRight:     12,
										},
										Children: []bff.UISnippet{
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:  "flag",
													Size:  20,
													Color: "#fff",
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
													Text:       trip["destination"],
													FontSize:   16,
													FontWeight: "600",
													Color:      "#1a237e",
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:     trip["destinationCity"],
													FontSize: 14,
													Color:    "#666",
												},
											},
										},
									},
								},
							},
						},
					},

					// Progress and ETA
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection:  "row",
							JustifyContent: "space-between",
							MarginBottom:   20,
						},
						Children: []bff.UISnippet{
							progressInfo("Progress", trip["progress"]+"%", "#4CAF50"),
							progressInfo("Distance", trip["distance"], "#2196F3"),
							progressInfo("ETA", trip["estimatedArrival"], "#FF9800"),
						},
					},

					// Trip Details Grid
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							FlexWrap:      "wrap",
							Gap:           12,
							MarginBottom:  20,
						},
						Children: []bff.UISnippet{
							tripDetailItem("Cargo", trip["cargo"], "cube-outline", "#9C27B0"),
							tripDetailItem("Weight", trip["weight"], "scale-outline", "#607D8B"),
							tripDetailItem("Payment", trip["payment"], "cash-outline", "#4CAF50"),
							tripDetailItem("Vehicle", trip["vehicleNumber"], "car-outline", "#2196F3"),
						},
					},

					// Contact Section
					{
						Type: "VIEW",
						Data: bff.ViewData{
							BackgroundColor: "#f8f9fa",
							BorderRadius:    12,
							Padding:         16,
							MarginBottom:    20,
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									FlexDirection:  "row",
									JustifyContent: "space-between",
									MarginBottom:   12,
								},
								Children: []bff.UISnippet{
									contactPerson("Sender", trip["senderName"], trip["senderPhone"], "person-outline", "#4CAF50"),
									contactPerson("Receiver", trip["receiverName"], trip["receiverPhone"], "person-outline", "#F44336"),
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
											Name:  "business",
											Size:  16,
											Color: "#FF9800",
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       trip["brokerName"],
											FontSize:   14,
											FontWeight: "600",
											Color:      "#1a237e",
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:     trip["brokerPhone"],
											FontSize: 14,
											Color:    "#666",
										},
									},
								},
							},
						},
					},

					// Action Buttons
					{
						Type: "VIEW",
						Data: bff.ViewData{
							FlexDirection: "row",
							Gap:           12,
						},
						Children: []bff.UISnippet{
							{
								Type: "BUTTON",
								Data: bff.ButtonData{
									Text: func() string {
										if data.IsTripStarted {
											return "Trip in Progress"
										}
										return "Start Trip"
									}(),
									Disabled: data.IsTripStarted || len(getRequiredDocs(data.DocumentsUploaded)) > 0,
									Style: bff.ViewData{
										Flex:            1,
										PaddingVertical: 16,
										BorderRadius:    12,
										BackgroundColor: "#4CAF50",
									},
									Action: bff.ActionData{
										Type:  "ACTION",
										Value: "START_TRIP",
										Url:   "/bff/driver/home/action",
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

// Add these helper functions at the bottom of home.go file

func rowBetween() bff.ViewData {
	return bff.ViewData{
		FlexDirection:  "row",
		JustifyContent: "space-between",
		AlignItems:     "center",
	}
}

func title(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:       text,
			FontSize:   18,
			FontWeight: "bold",
			Color:      "#1a237e",
		},
	}
}

func subtitle(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:     text,
			FontSize: 14,
			Color:    "#666",
		},
	}
}

func recentActivitiesSection(activities []bff.RecentActivity) bff.UISnippet {
	var activityItems []bff.UISnippet

	for _, activity := range activities {
		activityItems = append(activityItems, bff.UISnippet{
			Type: "VIEW",
			Data: bff.ViewData{
				FlexDirection:   "row",
				AlignItems:      "center",
				Padding:         12,
				BackgroundColor: "#fff",
				BorderRadius:    12,
				BorderWidth:     1,
				BorderColor:     "#f0f0f0",
			},
			Children: []bff.UISnippet{
				{
					Type: "VIEW",
					Data: bff.ViewData{
						Width:           40,
						Height:          40,
						BorderRadius:    20,
						BackgroundColor: activity.Color + "20",
						AlignItems:      "center",
						JustifyContent:  "center",
						MarginRight:     12,
					},
					Children: []bff.UISnippet{
						{
							Type: "ICON",
							Data: bff.IconData{
								Name:  activity.Icon,
								Size:  20,
								Color: activity.Color,
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
								Text:       activity.Message,
								FontSize:   14,
								FontWeight: "500",
								Color:      "#1a237e",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:     activity.Time,
								FontSize: 12,
								Color:    "#999",
							},
						},
					},
				},
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
			MarginBottom:      20,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         "Recent Activities",
					FontSize:     18,
					FontWeight:   "bold",
					Color:        "#1a237e",
					MarginBottom: 12,
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Gap: 8,
				},
				Children: activityItems,
			},
		},
	}
}

func performanceMetricsSection(data bff.HomeScreenData) bff.UISnippet {
	if !data.IsTripStarted {
		return bff.UISnippet{}
	}

	trip := data.ActiveTrip

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 20,
			MarginBottom:      20,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         "Trip Performance",
					FontSize:     18,
					FontWeight:   "bold",
					Color:        "#1a237e",
					MarginBottom: 16,
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					Gap:           12,
				},
				Children: []bff.UISnippet{
					metricCard("Fuel Level", trip["fuelLevel"]+"%", "speedometer-outline", "#FF9800", trip["fuelLevel"]),
					metricCard("Trip Score", trip["tripScore"], "trophy-outline", "#4CAF50", trip["tripScore"]),
					metricCard("Vehicle Health", trip["vehicleHealth"], "checkmark-circle-outline", "#2196F3", "85"),
				},
			},
		},
	}
}

// Helper functions
func getDocumentIcon(docType string) string {
	switch docType {
	case "eWayBill":
		return "document-text"
	case "invoice":
		return "receipt"
	case "vehicleRC":
		return "car"
	case "driverLicense":
		return "id-card"
	case "insurance":
		return "shield-checkmark"
	case "pollutionCert":
		return "leaf"
	default:
		return "document"
	}
}

func getDocumentName(docType string) string {
	switch docType {
	case "eWayBill":
		return "E-way Bill"
	case "invoice":
		return "Invoice"
	case "vehicleRC":
		return "Vehicle RC"
	case "driverLicense":
		return "Driver License"
	case "insurance":
		return "Insurance"
	case "pollutionCert":
		return "Pollution Certificate"
	default:
		return docType
	}
}

func getStatusColor(status string) string {
	switch status {
	case "not_started":
		return "#9E9E9E"
	case "reached_origin":
		return "#4CAF50"
	case "in_transit":
		return "#2196F3"
	case "reached_destination":
		return "#FF9800"
	case "completed":
		return "#9C27B0"
	default:
		return "#666"
	}
}

func getRequiredDocs(docs map[string]bool) []string {
	var required []string
	for doc, uploaded := range docs {
		if !uploaded {
			required = append(required, doc)
		}
	}
	return required
}

func progressInfo(label, value, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			AlignItems: "center",
			Flex:       1,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         label,
					FontSize:     12,
					Color:        "#666",
					MarginBottom: 4,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       value,
					FontSize:   16,
					FontWeight: "bold",
					Color:      color,
				},
			},
		},
	}
}

func tripDetailItem(label, value, icon, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Width:           "48%",
			FlexDirection:   "row",
			AlignItems:      "center",
			Gap:             8,
			Padding:         12,
			BackgroundColor: "#f8f9fa",
			BorderRadius:    8,
		},
		Children: []bff.UISnippet{
			{
				Type: "ICON",
				Data: bff.IconData{
					Name:  icon,
					Size:  16,
					Color: color,
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
							Text:     label,
							FontSize: 12,
							Color:    "#666",
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:          value,
							FontSize:      14,
							FontWeight:    "600",
							Color:         "#1a237e",
							FlexWrap:      "wrap",
							FlexShrink:    1,
							LineHeight:    18,
							FontFamily:    "monospace",
							LetterSpacing: 0.5,
							TextAlign:     "left",
							Width:         "100%",
						},
					},
				},
			},
		},
	}
}

func contactPerson(role, name, phone, icon, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			AlignItems: "flex-start",
			Flex:       1,
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					AlignItems:    "center",
					Gap:           8,
					MarginBottom:  4,
				},
				Children: []bff.UISnippet{
					{
						Type: "ICON",
						Data: bff.IconData{
							Name:  icon,
							Size:  14,
							Color: color,
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       role,
							FontSize:   12,
							FontWeight: "600",
							Color:      color,
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       name,
					FontSize:   14,
					FontWeight: "600",
					Color:      "#1a237e",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:     phone,
					FontSize: 13,
					Color:    "#666",
				},
			},
		},
	}
}

func metricCard(label, value, icon, color, progress string) bff.UISnippet {
	progressInt := 0
	if p, err := strconv.Atoi(progress); err == nil {
		progressInt = p
	}

	progressWidth := progressInt * 100 / 100

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Flex:            1,
			BackgroundColor: "#fff",
			BorderRadius:    12,
			Padding:         16,
			ShadowColor:     "#000",
			ShadowOpacity:   0.05,
			ShadowRadius:    4,
			Elevation:       1,
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection:  "row",
					AlignItems:     "center",
					JustifyContent: "space-between",
					MarginBottom:   8,
				},
				Children: []bff.UISnippet{
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:       label,
							FontSize:   14,
							FontWeight: "600",
							Color:      "#1a237e",
						},
					},
					{
						Type: "ICON",
						Data: bff.IconData{
							Name:  icon,
							Size:  18,
							Color: color,
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         value,
					FontSize:     24,
					FontWeight:   "bold",
					Color:        color,
					MarginBottom: 12,
				},
			},
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Height:          4,
					BackgroundColor: "#e0e0e0",
					BorderRadius:    2,
					Overflow:        "hidden",
				},
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							Width:           fmt.Sprintf("%d%%", progressWidth),
							Height:          4,
							BackgroundColor: color,
						},
					},
				},
			},
		},
	}
}

// HandleHomeAction - Updated with better action handling
// func HandleHomeAction(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	var req bff.ActionRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(bff.ActionResponse{
// 			Status:  "error",
// 			Message: "Invalid request format",
// 		})
// 		return
// 	}

// 	response := handleModernAction(req)
// 	json.NewEncoder(w).Encode(response)
// }

func HandleHomeAction(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")

	var req bff.ActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, bff.ActionResponse{
			Status:  "error",
			Message: "Invalid request format",
		})
		return
	}

	response := handleModernAction(req)
	c.JSON(200, response)
}


func handleModernAction(req bff.ActionRequest) bff.ActionResponse {
	switch req.Action {
	case "START_TRIP":
		// Check if documents are uploaded
		docsUploaded := req.Data["documentsUploaded"].(map[string]interface{})
		allUploaded := true
		var missingDocs []string

		for doc, uploaded := range docsUploaded {
			if !uploaded.(bool) {
				allUploaded = false
				missingDocs = append(missingDocs, doc)
			}
		}

		if !allUploaded {
			return bff.ActionResponse{
				Status:  "error",
				Message: "Please upload all required documents first",
				Data: map[string]interface{}{
					"missingDocuments": missingDocs,
				},
			}
		}

		// Start trip
		return bff.ActionResponse{
			Status:  "success",
			Message: "Trip started successfully! Location sharing enabled.",
			Data: map[string]interface{}{
				"isTripStarted":   true,
				"locationSharing": true,
				"tripStatus":      "in_transit",
				"startTime":       time.Now().Format(time.RFC3339),
			},
		}

	case "UPLOAD_DOCUMENT":
		docType := req.Data["documentType"].(string)
		// Simulate document upload
		return bff.ActionResponse{
			Status:  "success",
			Message: fmt.Sprintf("%s uploaded successfully", getDocumentName(docType)),
			Data: map[string]interface{}{
				"documentType": docType,
				"uploaded":     true,
				"uploadedAt":   time.Now().Format(time.RFC3339),
			},
		}

	case "UPDATE_STATUS":
		status := req.Data["status"].(string)
		tripID := req.Data["tripId"].(string)

		return bff.ActionResponse{
			Status:  "success",
			Message: fmt.Sprintf("Trip status updated to: %s", status),
			Data: map[string]interface{}{
				"tripId":     tripID,
				"tripStatus": status,
				"updatedAt":  time.Now().Format(time.RFC3339),
			},
		}

	case "VIEW_TRIP_DETAILS":
		tripID := req.Data["tripId"].(string)
		return bff.ActionResponse{
			Status: "success",
			Data: map[string]interface{}{
				"navigateTo": "/trip-details/" + tripID,
			},
		}

	case "CHAT_WITH_BROKER":
		brokerID := req.Data["brokerId"].(string)
		return bff.ActionResponse{
			Status: "success",
			Data: map[string]interface{}{
				"navigateTo": "/chat/" + brokerID,
			},
		}

	case "CALL_CONTACT":
		contactType := req.Data["contactType"].(string)
		phoneNumber := req.Data["phoneNumber"].(string)

		return bff.ActionResponse{
			Status:  "success",
			Message: fmt.Sprintf("Calling %s...", contactType),
			Data: map[string]interface{}{
				"action":      "call",
				"phoneNumber": phoneNumber,
				"contactType": contactType,
			},
		}
	}

	return bff.ActionResponse{
		Status:  "error",
		Message: "Unknown action",
	}
}
