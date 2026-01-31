package auth

import (
	"backend/bff"
	"fmt"
	"github.com/gin-gonic/gin"
)

func OtpScreenHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	phoneNumber := c.Query("phone")
	if phoneNumber == "" {
		phoneNumber = "your number" // fallback text
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "OTP",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "VIEW", // Wrap SCROLL in a parent VIEW
						Data: bff.ViewData{
							Flex:     1,
							Position: "relative", // Use relative positioning for absolute children
						},
						Children: []bff.UISnippet{
							// BACK BUTTON - Absolute positioned at top-left
							// {
							// 	Type: "VIEW", // Container for absolute positioning
							// 	Data: bff.ViewData{
							// 		Position: "absolute",
							// 		Top: 20,
							// 		Left: 16,
							// 		ZIndex: 10,
							// 	},
							// 	Children: []bff.UISnippet{
							// 		{
							// 			Type: "ICON_BUTTON",
							// 			Data: bff.IconButtonData{
							// 				Icon: "arrow-left",
							// 				Size: 28, // Slightly larger
							// 				Style: bff.ViewData{
							// 					Width: 48,
							// 					Height: 48,
							// 					BorderRadius: 24,
							// 					BackgroundColor: "#F5F5F5",
							// 					JustifyContent: "center",
							// 					AlignItems: "center",
							// 				},
							// 				OnPress: bff.ActionData{
							// 					Type: "NAVIGATE_BACK",
							// 				},
							// 			},
							// 		},
							// 	},
							// },

							// SCROLL VIEW for main content
							{
								Type: "SCROLL",
								Data: bff.ViewData{
									FlexGrow:          1,
									PaddingHorizontal: 24,
									PaddingVertical:   20,
									BackgroundColor:   "#FFFFFF",
									JustifyContent:    "center", // vertical centering
									AlignItems:        "center", // horizontal centering
								},
								Children: []bff.UISnippet{
									// HEADER
									{
										Type: "VIEW",
										Data: bff.ViewData{
											AlignItems:   "center",
											MarginBottom: 40,
											MarginTop:    40,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Verify Mobile Number",
													FontSize:   24,
													FontWeight: "bold",
													Color:      "#1A1A1A",
													TextAlign:  "center",
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      fmt.Sprintf("You'll receive a call soon at %s", phoneNumber),
													FontSize:  16,
													Color:     "#666666",
													TextAlign: "center",
													MarginTop: 12,
												},
											},
										},
									},

									// OTP INPUT
									{
										Type: "OTP_INPUT",
										Data: bff.OtpInputData{
											Id:                 "otp",
											Length:             6,
											BoxSpacing:         12, // Added spacing between boxes
											BoxBorderColor:     "#CCCCCC",
											BoxBorderRadius:    8,
											BoxBackgroundColor: "#F9F9F9",
											BoxWidth:           50,
											BoxHeight:          50,
										},
									},

									// VERIFY BUTTON
									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Verify & Continue",
											Style: bff.ViewData{
												BackgroundColor:   "#FF4D4D",
												PaddingVertical:   16,
												PaddingHorizontal: 40, // extra horizontal padding
												BorderRadius:      16,
												AlignItems:        "center",
												MarginTop:         30,
												ShadowColor:       "#000000",
												ShadowOpacity:     0.2,
												ShadowOffsetX:     0,
												ShadowOffsetY:     4,
												ShadowRadius:      6,
											},
											Action: bff.ActionData{
												Type: "API_CALL",
												Url:  "/api/v1/user/verify-otp",
												// SuccessNavigate: "(footbar)/home",
												SuccessNavigate: "(tabs)/",
												// FailureNavigate: "/(footbar)/home",
												FailureNavigate: "(tabs)/",
											},
										},
									},

									// RESEND
									{
										Type: "RESEND_OTP",
										Data: bff.ResendOtpData{
											Timer:     30,
											MarginTop: 20,
										},
									},

									// SECURITY NOTE
									{
										Type: "TEXT", // Changed from FOOTER_NOTE to TEXT
										Data: bff.TextData{
											Text:      "🔒 Your OTP is secure and encrypted",
											FontSize:  12,
											Color:     "#666666",
											TextAlign: "center",
											MarginTop: 40,
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
