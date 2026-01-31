package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func R6Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R6_MOBILE_INPUT",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Data: bff.ViewData{
							Flex:            1,
							Position:        "relative",
							BackgroundColor: "#FFFFFF",
						},
						Children: []bff.UISnippet{
							// ================= BACK BUTTON (Absolute positioned) =================
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Position: "absolute",
									Top:      20,
									Left:     16,
									ZIndex:   10,
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON_BUTTON",
										Data: bff.IconButtonData{
											Icon: "arrow-left",
											Size: 28,
											// Style: bff.ViewData{
											// 	Width: 48,
											// 	Height: 48,
											// 	BorderRadius: 24,
											// 	BackgroundColor: "#F5F5F5",
											// 	JustifyContent: "center",
											// 	AlignItems: "center",
											// },
											OnPress: bff.ActionData{
												Type: "NAVIGATE_BACK",
											},
										},
									},
								},
							},

							// ================= MAIN CONTENT =================
							{
								Type: "SCROLL",
								Data: bff.ViewData{
									FlexGrow:          1,
									BackgroundColor:   "#FFFFFF",
									PaddingHorizontal: 24,
									PaddingVertical:   40,
									JustifyContent:    "center",
								},
								Children: []bff.UISnippet{
									// ================= HEADER SECTION =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											AlignItems:   "center",
											MarginBottom: 40,
											MarginTop:    60, // Increased for back button space
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Enter Mobile Number",
													FontSize:     28,
													FontWeight:   "bold",
													Color:        "#1A1A1A",
													TextAlign:    "center",
													MarginBottom: 12,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "We'll send you a 6-digit OTP to verify your number",
													FontSize:  16,
													Color:     "#666666",
													TextAlign: "center",
												},
											},
										},
									},

									// ================= MOBILE INPUT SECTION =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width:        "100%",
											MarginBottom: 40,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Mobile Number *",
													FontSize:     16,
													FontWeight:   "600",
													Color:        "#1A1A1A",
													MarginBottom: 12,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													FlexDirection:   "row",
													AlignItems:      "center",
													BorderWidth:     2,
													BorderColor:     "#F0F0F0",
													BorderRadius:    16,
													Padding:         16,
													BackgroundColor: "#FAFAFA",
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															FlexDirection:     "row",
															AlignItems:        "center",
															BackgroundColor:   "#F0F0F0",
															PaddingHorizontal: 12,
															PaddingVertical:   8,
															BorderRadius:      8,
															MarginRight:       12,
														},
														Children: []bff.UISnippet{
															{
																Type: "TEXT",
																Data: bff.TextData{
																	Text:        "+91",
																	FontSize:    16,
																	FontWeight:  "600",
																	Color:       "#1A1A1A",
																	MarginRight: 6,
																},
															},
															{
																Type: "ICON",
																Data: bff.IconData{
																	Name:  "chevron-down",
																	Size:  16,
																	Color: "#666666",
																},
															},
														},
													},
													{
														Type: "INPUT",
														Data: bff.InputData{
															Id:           "phone", // Changed from mobileNumber to phone
															Placeholder:  "10-digit mobile number",
															KeyboardType: "phone-pad",
															MaxLength:    10,
															TextColor:    "#1A1A1A",
															FontSize:     16,
															FontWeight:   "500",
															Style: bff.ViewData{
																Flex: 1,
															},
														},
													},
												},
											},
										},
									},

									// ================= SEND OTP BUTTON =================
									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Send OTP",
											Action: bff.ActionData{
												Type:            "API_CALL",
												Url:             "/api/v2/user/request-otp",
												Method:          "POST",
												SuccessNavigate: "/r8",
												FailureNavigate: "",
											},
											Style: bff.ViewData{
												BackgroundColor: "#FF0000",
												PaddingVertical: 18,
												BorderRadius:    16,
												AlignItems:      "center",
												ShadowColor:     "#000000",
												ShadowOpacity:   0.2,
												ShadowOffsetX:   0,
												ShadowOffsetY:   4,
												ShadowRadius:    6,
												Elevation:       3,
											},
										},
									},

									// ================= TERMS NOTE =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											MarginTop:  30,
											AlignItems: "center",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "By continuing, you agree to our Terms & Privacy Policy",
													FontSize:  12,
													Color:     "#666666",
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
