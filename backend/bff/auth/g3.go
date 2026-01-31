package auth

import (
	"backend/bff"

	"github.com/gin-gonic/gin"
)

func G3Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "G3",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							FlexGrow:        1,
							BackgroundColor: "#FFFFFF",
						},
						Children: []bff.UISnippet{
							{
								Type: "VIEW",
								Data: bff.ViewData{
									PaddingHorizontal: 24,
									PaddingVertical:   20,
								},
								Children: []bff.UISnippet{
									// BACK BUTTON - Use ICON_BUTTON instead of BUTTON with icon
									{
										Type: "ICON_BUTTON",
										Data: bff.IconButtonData{
											Icon: "arrow-left",
											OnPress: bff.ActionData{
												Type: "NAVIGATE_BACK",
											},
										},
									},

									// HEADER AND PROGRESS
									{
										Type: "VIEW",
										Data: bff.ViewData{
											AlignItems:   "center",
											MarginBottom: 40,
											MarginTop:    20, // Reduced margin since no back button overlap
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:          "Step 2 Of 4",
													FontSize:      28,
													FontWeight:    "bold",
													Color:         "#1A1A1A",
													MarginBottom:  30,
													TextAlign:     "center",
													LetterSpacing: -0.5,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:      "100%",
													AlignItems: "center",
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "100%",
															Height:          8,
															BackgroundColor: "#F0F0F0",
															BorderRadius:    4,
															MarginBottom:    12,
															Overflow:        "hidden",
														},
														Children: []bff.UISnippet{
															{
																Type: "VIEW",
																Data: bff.ViewData{
																	Width:           "40%", // 40% progress
																	Height:          8,
																	BackgroundColor: "#FF0000",
																	BorderRadius:    4,
																},
															},
														},
													},
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "40% Complete",
															FontSize:   14,
															Color:      "#666666",
															FontWeight: "600",
														},
													},
												},
											},
										},
									},

									// FORM SECTION
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width: "100%",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:          "Business Information",
													FontSize:      18,
													Color:         "#1A1A1A",
													MarginBottom:  25,
													FontWeight:    "600",
													LetterSpacing: -0.3,
												},
											},

											BusinessInputField("Company Name *", "companyName", "Enter company name"),
											BusinessInputField("GST Number", "gstNumber", "Enter GST number"),
											BusinessInputField("City", "city", "Enter your city"),
										},
									},

									// CONTINUE BUTTON - SIMPLIFIED VERSION (no icon)
									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Continue",
											Action: bff.ActionData{
												Type:     "NAVIGATE",
												Navigate: "/g4", // Fixed property name
											},
											Style: bff.ViewData{
												BackgroundColor: "#FF0000",
												AlignItems:      "center",
												JustifyContent:  "center",
												PaddingVertical: 18,
												BorderRadius:    16,
												MarginTop:       20,
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

// InputField helper function
func BusinessInputField(label, id, placeholder string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			MarginBottom:    20,
			BorderWidth:     2,
			BorderColor:     "#F0F0F0",
			BorderRadius:    16,
			Padding:         18,
			BackgroundColor: "#FAFAFA",
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:          label,
					FontSize:      15,
					FontWeight:    "600",
					Color:         "#1A1A1A",
					MarginBottom:  10,
					LetterSpacing: -0.2,
				},
			},
			{
				Type: "INPUT",
				Data: bff.InputData{
					Id:          id,
					Placeholder: placeholder,
					TextColor:   "#1A1A1A",
					FontSize:    16,
					FontWeight:  "500",
				},
			},
		},
	}
}
