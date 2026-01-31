package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func G2Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "G2",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							FlexGrow:        1,
							PaddingHorizontal: 24,
							PaddingVertical: 40,
							BackgroundColor: "#FFFFFF",
						},
						Children: []bff.UISnippet{

							// HEADER
							{
								Type: "VIEW",
								Data: bff.ViewData{
									FlexDirection: "row",
									AlignItems:    "flex-start",
									MarginBottom:  30,
								},
								Children: []bff.UISnippet{
									// {
									// 	Type: "BUTTON",
									// 	Data: bff.ButtonData{
									// 		Text: "",
									// 		Icon: bff.IconData{
									// 			Name:  "arrow-left",
									// 			Size:  24,
									// 			Color: "#1A1A1A",
									// 		},
									// 		Action: bff.ActionData{
									// 			Type: "NAVIGATE_BACK",
									// 		},
									// 		Style: bff.ViewData{
									// 			Padding:         8,
									// 			MarginRight:     16,
									// 			BackgroundColor: "#F8F8F8",
									// 			BorderRadius:    8,
									// 		},
									// 	},
									// },
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Flex: 1,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "KYC Verification",
													FontSize:   24,
													FontWeight: "700",
													Color:      "#1A1A1A",
													MarginBottom: 8,
													LetterSpacing: -0.5,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "Upload Required Documents",
													FontSize:  16,
													Color:     "#666666",
													LineHeight: 22,
												},
											},
										},
									},
								},
							},

							// PROGRESS BAR
							{
								Type: "VIEW",
								Data: bff.ViewData{
									MarginBottom: 40,
								},
								Children: []bff.UISnippet{
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:           "100%",
											Height:          6,
											BackgroundColor: "#F0F0F0",
											BorderRadius:    3,
											MarginBottom:    12,
											Overflow:        "hidden",
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:           "50%",
													Height:          6,
													BackgroundColor: "#FF0000",
													BorderRadius:    3,
												},
											},
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Step 2 of 4",
											FontSize:   14,
											Color:      "#666666",
											FontWeight: "600",
											TextAlign:  "center",
										},
									},
								},
							},

							// DOCUMENT UPLOAD CARDS
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Gap: 20,
									MarginBottom: 30,
								},
								Children: []bff.UISnippet{
									DocumentCard("Upload PAN Card", "panCard"),
									DocumentCard("Upload Aadhar Card", "aadharCard"),
									DocumentCard("Upload Broker License", "brokerLicense"),
									DocumentCard("Upload Company Documents", "companyDocuments"),
								},
							},

							// INFO TEXT
							{
								Type: "VIEW",
								Data: bff.ViewData{
									BackgroundColor: "#F8F8F8",
									Padding:         20,
									BorderRadius:    12,
									MarginBottom:    30,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text: `• Ensure documents are clear and readable
• File size should be less than 5MB
• Supported formats: JPG, PNG, PDF`,
											FontSize: 14,
											Color:    "#666666",
											LineHeight: 20,
										},
									},
								},
							},

							// CONTINUE BUTTON
							{
								Type: "BUTTON",
								Data: bff.ButtonData{
									Text: "Continue",
									Action: bff.ActionData{
										Type:            "NAVIGATE",
										Navigate: "/g5",
									},
									Style: bff.ViewData{
										BackgroundColor: "#FF0000",
										FlexDirection:   "row",
										AlignItems:      "center",
										JustifyContent:  "center",
										Gap:             10,
										PaddingVertical: 18,
										BorderRadius:    16,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	c.JSON(200, response)

}

func DocumentCard(title, id string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#FFFFFF",
			BorderWidth:     2,
			BorderColor:     "#F0F0F0",
			BorderRadius:    16,
			Padding:         24,
			AlignItems:      "center",
			ShadowColor:     "#000000",
			ShadowOffsetX:   0,
			ShadowOffsetY:   2,
			ShadowOpacity:   0.05,
			ShadowRadius:    8,
			Elevation:       2,
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					MarginBottom: 16,
				},
				Children: []bff.UISnippet{
					{
						// Using TEXT instead of ICON since ICON type isn't fully implemented
						Type: "TEXT",
						Data: bff.TextData{
							Text: "➕", // Plus icon as text
							FontSize: 28,
							Color: "#FF0000",
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       title,
					FontSize:   18,
					FontWeight: "600",
					Color:      "#1A1A1A",
					MarginBottom: 8,
					TextAlign:  "center",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:      "Upload JPG, PNG, or PDF",
					FontSize:  14,
					Color:     "#666666",
					TextAlign: "center",
				},
			},
		},
	}
}