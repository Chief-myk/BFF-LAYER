package auth

import (
	"backend/bff"
	"fmt"
	"github.com/gin-gonic/gin"
)

func R8Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Extract mobile number from query parameter or context
	// In real implementation, you might get this from session or previous API call
	// phoneNumber := r.URL.Query().Get("phone")
	phoneNumber := c.Query("phone")
	if phoneNumber == "" {
		phoneNumber = "9876543210" // Default fallback
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R8_OTP_VERIFICATION",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
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

							// ================= BACK BUTTON =================
							{
								Type: "ICON_BUTTON",
								Data: bff.IconButtonData{
									Icon: "arrow-left",
									OnPress: bff.ActionData{
										Type: "NAVIGATE_BACK",
									},
								},
							},

							// ================= HEADER SECTION =================
							{
								Type: "VIEW",
								Data: bff.ViewData{
									AlignItems:   "center",
									MarginBottom: 40,
									MarginTop:    20,
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
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "We've sent a 6-digit OTP to",
											FontSize:  16,
											Color:     "#666666",
											TextAlign: "center",
											MarginBottom: 4,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      fmt.Sprintf("+91 %s", phoneNumber),
											FontSize:  18,
											FontWeight: "600",
											Color:     "#1A1A1A",
											TextAlign: "center",
										},
									},
									// {
									// 	Type: "TEXT_BUTTON",
									// 	Data: bff.TextButtonData{
									// 		Text: "Change number",
									// 		Color: "#FF0000",
									// 		FontSize: 14,
									// 		FontWeight: "500",
									// 		OnPress: bff.ActionData{
									// 			Type: "NAVIGATE_BACK",
									// 		},
									// 	},
									// },
								},
							},

							// ================= OTP INPUT SECTION =================
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:        "100%",
									MarginBottom: 40,
									AlignItems:   "center",
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Enter OTP Code",
											FontSize:   16,
											FontWeight: "600",
											Color:      "#1A1A1A",
											MarginBottom: 20,
											TextAlign:  "center",
										},
									},
									{
										Type: "OTP_INPUT",
										Data: bff.OtpInputData{
											Id:                    "otp",
											Length:                6,
											BoxSpacing:            12,
											BoxBorderColor:        "#CCCCCC",
											BoxBorderRadius:       8,
											BoxBackgroundColor:    "#F9F9F9",
											BoxWidth:              50,
											BoxHeight:             50,
										},
									},
								},
							},

							// ================= VERIFY BUTTON =================
							{
								Type: "BUTTON",
								Data: bff.ButtonData{
									Text: "Verify & Continue",
									Action: bff.ActionData{
										Type:            "API_CALL",
										Url:             "/api/v2/user/verify-otp",
										Method:          "POST",
										SuccessNavigate: "/r1",
										FailureNavigate: "", 
									},
									Style: bff.ViewData{
										BackgroundColor: "#FF0000",
										PaddingVertical: 18,
										PaddingHorizontal: 40,
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

							// ================= RESEND OTP SECTION =================
							{
								Type: "RESEND_OTP",
								Data: bff.ResendOtpData{
									Timer:     30,
									MarginTop: 20,
								},
							},

							// ================= SECURITY NOTE =================
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
											Text:      "🔒 Your OTP is secure and encrypted",
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
	}

	c.JSON(200, response)
}