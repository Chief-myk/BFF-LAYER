package driver

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func ProfileScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Mock driver data coming from backend
	driver := map[string]interface{}{
		"name":        "Rajesh Kumar",
		"phone":       "+919123456789",
		"truckNumber": "MH01AB1234",
		"brokerName":  "Kumar Logistics",
		"avatar":      "https://i.pravatar.cc/150?img=12",
		"kycStatus":   "verified",
		"wallet":      "₹12,500",
		"documents": map[string]string{
			"aadhar":         "uploaded",
			"pan":            "uploaded",
			"drivingLicense": "pending",
			"rcInsurance":    "uploaded",
		},
	}

	// Build UI snippets
	ui := []bff.UISnippet{
		{
			Type: "VIEW",
			Data: bff.ViewData{
				FlexDirection:     "row",
				JustifyContent:    "space-between",
				PaddingHorizontal: 20,
				PaddingVertical:   20,
				BackgroundColor:   "#FFFFFF",
			},
			Children: []bff.UISnippet{
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:       "Profile",
						FontSize:   24,
						FontWeight: "700",
						Color:      "#1A1A1A",
					},
				},
				{
					Type: "ICON_BUTTON",
					Data: bff.IconButtonData{
						Icon: "pencil-outline",
						OnPress: bff.ActionData{
							Type: "NAVIGATE",
							To:   "EditProfile",
						},
					},
				},
			},
		},
		{
			Type: "CARD",
			Data: bff.CardData{
				BackgroundColor: "#FFFFFF",
				Padding:         20,
				BorderRadius:    12,
				Shadow:          true,
			},
			Children: []bff.UISnippet{
				{
					Type: "VIEW",
					Data: bff.ViewData{
						FlexDirection: "row",
						AlignItems:    "center",
						MarginBottom:  20,
					},
					Children: []bff.UISnippet{
						{
							Type: "IMAGE",
							Data: bff.ImageData{
								Url:    driver["avatar"].(string),
								Width:  70,
								Height: 70,
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
										Text:       driver["name"].(string),
										FontSize:   20,
										FontWeight: "600",
										Color:      "#1A1A1A",
									},
								},
								{
									Type: "VIEW",
									Data: bff.ViewData{
										FlexDirection:     "row",
										AlignItems:        "center",
										PaddingHorizontal: 12,
										PaddingVertical:   6,
										BorderRadius:      16,
										BackgroundColor:   "#F8F8F8",
										AlignSelf:         "flex-start",
									},
									Children: []bff.UISnippet{
										{
											Type: "ICON",
											Data: bff.IconData{
												Name:  "checkmark-circle",
												Size:  16,
												Color: "#4CAF50",
											},
										},
										{
											Type: "TEXT",
											Data: bff.TextData{
												Text:       "Verified",
												FontSize:   12,
												FontWeight: "600",
												Color:      "#4CAF50",
											},
										},
									},
								},
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
							Type: "TEXT",
							Data: bff.TextData{
								Text:     "+919123456789",
								FontSize: 14,
								Color:    "#666",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:     "MH01AB1234",
								FontSize: 14,
								Color:    "#666",
							},
						},
						{
							Type: "TEXT",
							Data: bff.TextData{
								Text:     "Kumar Logistics",
								FontSize: 14,
								Color:    "#666",
							},
						},
					},
				},
			},
		},
		{
			Type: "CARD",
			Data: bff.CardData{
				BackgroundColor: "#FFFFFF",
				Padding:         20,
				BorderRadius:    12,
				Shadow:          true,
			},
			Children: []bff.UISnippet{
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:       "Wallet Balance",
						FontSize:   18,
						FontWeight: "600",
						Color:      "#1A1A1A",
					},
				},
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:       driver["wallet"].(string),
						FontSize:   28,
						FontWeight: "700",
						Color:      "#1A1A1A",
					},
				},
			},
		},
		{
			Type: "BUTTON",
			Data: bff.ButtonData{
				Text: "Logout",
				Style: bff.ViewData{
					BackgroundColor: "#ff0000",
					Padding:         16,
					BorderRadius:    12,
				},
				Action: bff.ActionData{
					Type: "LOGOUT",
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "Profile",
		UI:     ui,
	}

	c.JSON(200, response)
}
