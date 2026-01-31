package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)
func AuthScreenHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "AUTH",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							FlexGrow:          1,
							PaddingHorizontal: 24,
							PaddingVertical:   40,
							AlignItems:        "center",
							JustifyContent:    "center",
							BackgroundColor:   "#ffffff",
						},
						Children: []bff.UISnippet{
							// HEADER
							{
								Type: "VIEW",
								Data: bff.ViewData{
									AlignItems:   "center",
									MarginBottom: 30,
								},
								Children: []bff.UISnippet{
									{
										Type: "IMAGE",
										Data: bff.ImageData{
											Url:        "https://cdn.truckhai.com/logo.png",
											Width:      180,
											Height:     100,
											ResizeMode: "contain",
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "Get started with your freight journey",
											FontSize:  16,
											Color:     "#666666",
											TextAlign: "center",
											MarginTop: -10,
										},
									},
								},
							},

							// LOGIN CONTAINER
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width: "100%",
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Enter Mobile Number",
											FontSize:     18,
											FontWeight:   "600",
											Color:        "#1A1A1A",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "We'll send you an OTP to verify your number",
											FontSize:     14,
											Color:        "#666666",
											MarginBottom: 24,
										},
									},

									// PHONE INPUT ROW
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection: "row",
											AlignItems:    "center",
											MarginBottom:  30,
										},
										Children: []bff.UISnippet{
											{
												Type: "BUTTON",
												Data: bff.ButtonData{
													Text: "+91",
													Style: bff.ViewData{
														FlexDirection:     "row",
														AlignItems:        "center",
														PaddingHorizontal: 16,
														PaddingVertical:   15,
														BorderRadius:      12,
														Color:             "#1A1A1A",
														BackgroundColor:   "#F8F8F8",
														TextColor:         "#1a1a1a",
														BorderWidth:       1,
														BorderColor:       "#E5E5E5",
														MarginRight:       12,
													},
												},
											},
											{
												Type: "INPUT",
												Data: bff.InputData{
													Id:           "phone",
													Placeholder:  "Mobile Number",
													KeyboardType: "phone-pad",
													MaxLength:    10,
													Style: bff.ViewData{
														Flex:              1,
														BackgroundColor:   "#F8F8F8",
														PaddingHorizontal: 16,
														PaddingVertical:   15,
														BorderRadius:      12,
														BorderWidth:       1,
														BorderColor:       "#E5E5E5",
													},
												},
											},
										},
									},

									// CONTINUE BUTTON
									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Continue",
											Style: bff.ViewData{
												BackgroundColor: "#FF0000",
												PaddingVertical: 16,
												BorderRadius:    12,
												AlignItems:      "center",
											},
											Action: bff.ActionData{
												Type:            "API_CALL",
												// Url:             "/api/v1/user/request-otp",
												Url:             "/api/v1/user/request-otpp",
												SuccessNavigate: "/(auth)/otp",
												// FailureNavigate: "/(auth)/registration-role",
												FailureNavigate: "/(auth)/otp",
											},
										},
									},

									// TERMS
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "By continuing, you agree to our Terms of Service and Privacy Policy",
											FontSize:  12,
											Color:     "#666666",
											TextAlign: "center",
											MarginTop: 30,
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