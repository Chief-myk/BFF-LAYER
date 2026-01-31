package auth

import (
	// "fmt"
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func R7Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	nextScreen := "(footbar)/home"


	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R7_SUCCESS",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							Flex:            1,
							BackgroundColor: "#FFFFFF",
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex:            1,
									BackgroundColor: "#FFFFFF",
									Padding:         32,
									PaddingTop:      40,
									PaddingBottom:   60,
									AlignItems:      "center",
								},
								Children: []bff.UISnippet{

									// ================= SUCCESS ICON =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:           120,
											Height:          120,
											BorderRadius:    60,
											BackgroundColor: "#F0F9FF",
											JustifyContent:  "center",
											AlignItems:      "center",
											MarginBottom:    32,
											ShadowColor:     "#0EA5E9",
											ShadowOffsetX:   0,
											ShadowOffsetY:   8,
											ShadowOpacity:   0.15,
											ShadowRadius:    20,
											Elevation:       10,
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:           80,
													Height:          80,
													BorderRadius:    40,
													BackgroundColor: "#0EA5E9",
													JustifyContent:  "center",
													AlignItems:      "center",
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "✓",
															FontSize:   40,
															FontWeight: "bold",
															Color:      "#FFFFFF",
														},
													},
												},
											},
										},
									},

									// ================= MAIN HEADING =================
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Welcome Aboard! 🎉",
											FontSize:   32,
											FontWeight: "bold",
											Color:      "#1E293B",
											TextAlign:  "center",
											MarginBottom: 12,
										},
									},

									// ================= SUBTITLE =================
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Your TruckHai Journey Begins",
											FontSize:   18,
											Color:      "#64748B",
											TextAlign:  "center",
											MarginBottom: 40,
										},
									},

									// ================= PROGRESS COMPLETION =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#F8FAFC",
											BorderRadius:    16,
											Padding:         20,
											Width:           "100%",
											MarginBottom:    40,
											ShadowColor:     "#000000",
											ShadowOffsetX:   0,
											ShadowOffsetY:   2,
											ShadowOpacity:   0.05,
											ShadowRadius:    8,
											Elevation:       2,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Registration Complete",
													FontSize:   18,
													FontWeight: "600",
													Color:      "#1E293B",
													MarginBottom: 16,
													TextAlign:  "center",
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:        "100%",
													MarginBottom: 12,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Height:          10,
															BorderRadius:    5,
															BackgroundColor: "#E2E8F0",
															Overflow:        "hidden",
														},
														Children: []bff.UISnippet{
															{
																Type: "VIEW",
																Data: bff.ViewData{
																	Width:           "100%",
																	Height:          10,
																	BackgroundColor: "#10B981",
																	BorderRadius:    5,
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
													JustifyContent: "space-between",
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "100% Complete",
															FontSize:   14,
															FontWeight: "600",
															Color:      "#10B981",
														},
													},
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "Step 6 of 6",
															FontSize:   14,
															Color:      "#64748B",
														},
													},
												},
											},
										},
									},

									// ================= FEATURES TITLE =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:        "100%",
											MarginBottom: 24,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "What's Next?",
													FontSize:   22,
													FontWeight: "bold",
													Color:      "#1E293B",
													MarginBottom: 8,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "Explore amazing features waiting for you",
													FontSize:  16,
													Color:     "#64748B",
												},
											},
										},
									},

									// ================= FEATURES GRID =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:          "100%",
											MarginBottom:   48,
											FlexDirection:  "row",
											FlexWrap:       "wrap",
											JustifyContent: "space-between",
										},
										Children: []bff.UISnippet{
											featureCard("🚀", "Instant Access", "Start booking loads immediately"),
											featureCard("✅", "Verified Profile", "Build trust with verified badge"),
											featureCard("📈", "Business Growth", "Access premium freight opportunities"),
											featureCard("🛡️", "Secure Platform", "Your data is encrypted & safe"),
											featureCard("💬", "24/7 Support", "Always here to help you"),
											featureCard("💰", "Earn More", "Competitive rates & bonuses"),
										},
									},

									// ================= GET STARTED BUTTON =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:        "100%",
											MarginBottom: 32,
										},
										Children: []bff.UISnippet{
											{
												Type: "BUTTON",
												Data: bff.ButtonData{
													Text: "Launch Your Dashboard",
													Action: bff.ActionData{
														Type:     "NAVIGATE",
														Navigate: nextScreen,
													},
													Style: bff.ViewData{
														BackgroundColor: "#FF0000",
														BorderRadius:    14,
														PaddingVertical: 20,
														AlignItems:      "center",
														ShadowColor:     "#000000",
														ShadowOpacity:   0.2,
														ShadowOffsetX:   0,
														ShadowOffsetY:   6,
														ShadowRadius:    12,
														Elevation:       5,
													},
												},
											},
										},
									},

									// ================= WELCOME MESSAGE =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#FFFBEB",
											BorderRadius:    16,
											Padding:         24,
											Width:           "100%",
											BorderLeftWidth: 6,
											BorderColor:     "#F59E0B",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "✨ Welcome to the Family!",
													FontSize:   18,
													FontWeight: "600",
													Color:      "#92400E",
													MarginBottom: 12,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "You've joined India's fastest-growing freight community. We're excited to help you grow your business with smart logistics solutions.",
													FontSize:  15,
													Color:     "#92400E",
													LineHeight: 22,
												},
											},
										},
									},

									// ================= FOOTER =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											MarginTop: 40,
											AlignItems: "center",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "Need help? Contact us at support@truckhai.com",
													FontSize:  14,
													Color:     "#64748B",
													TextAlign: "center",
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

	c.JSON(200, response)
}

// ================= FEATURE CARD HELPER =================
func featureCard(icon, title, desc string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Width:           "48%",
			BackgroundColor: "#FFFFFF",
			BorderRadius:    16,
			Padding:         20,
			MarginBottom:    16,
			ShadowColor:     "#000000",
			ShadowOffsetX:   0,
			ShadowOffsetY:   2,
			ShadowOpacity:   0.05,
			ShadowRadius:    8,
			Elevation:       1,
			AlignItems:      "center",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Width:           56,
					Height:          56,
					BorderRadius:    28,
					BackgroundColor: "#F0F9FF",
					JustifyContent:  "center",
					AlignItems:      "center",
					MarginBottom:    16,
				},
				Children: []bff.UISnippet{
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:      icon,
							FontSize:  24,
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       title,
					FontSize:   16,
					FontWeight: "600",
					Color:      "#1E293B",
					TextAlign:  "center",
					MarginBottom: 8,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:      desc,
					FontSize:  13,
					Color:     "#64748B",
					TextAlign: "center",
					LineHeight: 18,
				},
			},
		},
	}
}