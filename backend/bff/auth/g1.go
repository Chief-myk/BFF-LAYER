package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func G1Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "G1",
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
									AlignItems: "center",
									MarginBottom: 40,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Join TruckHai Today",
											FontSize:   28,
											FontWeight: "bold",
											Color:      "#1A1A1A",
											TextAlign:  "center",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "Get started with your freight journey",
											FontSize:  16,
											Color:     "#666666",
											TextAlign: "center",
											MarginBottom: 30,
										},
									},

									// PROGRESS BAR
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width: "100%",
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:           "100%",
													Height:          8,
													BackgroundColor: "#F0F0F0",
													BorderRadius:    4,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "20%",
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
													Text:      "20% Complete",
													FontSize:  14,
													Color:     "#666666",
													TextAlign: "center",
													MarginTop: 12,
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
											Text:       "Personal Information",
											FontSize:   18,
											FontWeight: "600",
											Color:      "#1A1A1A",
											MarginBottom: 25,
										},
									},

									// NAME ROW
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection: "row",
											MarginBottom: 20,
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Flex: 1,
													MarginRight: 16,
												},
												Children: []bff.UISnippet{
													InputField("firstName", "First Name *", "Enter first name", "text"),
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Flex: 1,
													MarginLeft: 16,
												},
												Children: []bff.UISnippet{
													InputField("lastName", "Last Name *", "Enter last name", "text"),
												},
											},
										},
									},

									// EMAIL
									InputField("email", "Email *", "Enter email address", "email"),

									// MOBILE NUMBER
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BorderWidth:     2,
											BorderColor:     "#F0F0F0",
											BorderRadius:    16,
											Padding:         18,
											BackgroundColor: "#FAFAFA",
											MarginBottom:    20,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Mobile Number *",
													FontSize:   15,
													FontWeight: "600",
													Color:      "#1A1A1A",
													MarginBottom: 10,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													FlexDirection: "row",
													AlignItems:    "center",
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															FlexDirection:     "row",
															AlignItems:        "center",
															BackgroundColor:   "#F0F0F0",
															PaddingHorizontal: 14,
															PaddingVertical:   10,
															BorderRadius:      10,
															MarginRight:       14,
														},
														Children: []bff.UISnippet{
															{
																Type: "TEXT",
																Data: bff.TextData{
																	Text:       "+91",
																	FontSize:   15,
																	FontWeight: "600",
																	Color:      "#1A1A1A",
																	MarginRight: 6,
																},
															},
															{
																Type: "TEXT",
																Data: bff.TextData{
																	Text: "▼", // Using text instead of Icon type
																	FontSize: 12,
																	Color: "#666666",
																},
															},
														},
													},
													{
														Type: "INPUT",
														Data: bff.InputData{
															Id:           "mobileNumber",
															Placeholder:  "10-digit mobile number",
															KeyboardType: "phone-pad",
															MaxLength:    10,
														},
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
        Action: bff.ActionData{
            Type:     "NAVIGATE",  // Changed type
            Navigate: "/g3",       // Changed field name
        },
        Style: bff.ViewData{
            BackgroundColor: "#FF0000",
            PaddingVertical: 18,
            BorderRadius:    16,
            AlignItems:      "center",
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

func InputField(id, label, placeholder, inputType string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BorderWidth:     2,
			BorderColor:     "#F0F0F0",
			BorderRadius:    16,
			Padding:         18,
			BackgroundColor: "#FAFAFA",
			MarginBottom:    20,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       label,
					FontSize:   15,
					FontWeight: "600",
					Color:      "#1A1A1A",
					MarginBottom: 10,
				},
			},
			{
				Type: "INPUT",
				Data: bff.InputData{
					Id:          id,
					Placeholder: placeholder,
					KeyboardType: inputType,
				},
			},
		},
	}
}