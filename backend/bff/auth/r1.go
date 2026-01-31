package auth

import (
	"backend/bff"

	"github.com/gin-gonic/gin"
)

func R1Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R1_DRIVER_BASIC_DETAILS",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							FlexGrow:          1,
							PaddingHorizontal: 24,
							PaddingVertical:   20,
							BackgroundColor:   "#FFFFFF",
							JustifyContent:    "space-between",
						},
						Children: []bff.UISnippet{

							// BACK BUTTON
							{
								Type: "ICON_BUTTON",
								Data: bff.IconButtonData{
									Icon: "arrow-left",
									OnPress: bff.ActionData{
										Type: "NAVIGATE_BACK",
									},
								},
							},

							// HEADER
							{
								Type: "VIEW",
								Data: bff.ViewData{
									AlignItems:   "center",
									MarginBottom: 40,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Join TruckHai Today",
											FontSize:     28,
											FontWeight:   "bold",
											Color:        "#1A1A1A",
											TextAlign:    "center",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Get started with your freight journey",
											FontSize:     16,
											Color:        "#666666",
											TextAlign:    "center",
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
													BackgroundColor: "#F0F0F0",
													BorderRadius:    4,
													Padding:         0,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "20%",
															BackgroundColor: "#FF0000",
															BorderRadius:    4,
															PaddingVertical: 4,
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
													MarginTop: 12,
												},
											},
										},
									},
								},
							},

							// FORM
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width: "100%",
								},
								Children: []bff.UISnippet{

									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Personal Information",
											FontSize:     18,
											FontWeight:   "600",
											Color:        "#1A1A1A",
											MarginBottom: 25,
										},
									},

									// FIRST + LAST NAME
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection: "row",
											MarginBottom:  20,
										},
										Children: []bff.UISnippet{

											{
												Type: "INPUT",
												Data: bff.InputData{
													Id:          "firstName",
													Placeholder: "Enter first name",
													Style: bff.ViewData{
														Flex:            1,
														MarginRight:     8,
														BorderRadius:    16,
														BorderWidth:     2,
														BorderColor:     "#F0F0F0",
														Padding:         18,
														BackgroundColor: "#FAFAFA",
													},
												},
											},

											{
												Type: "INPUT",
												Data: bff.InputData{
													Id:          "lastName",
													Placeholder: "Enter last name",
													Style: bff.ViewData{
														Flex:            1,
														MarginLeft:      8,
														BorderRadius:    16,
														BorderWidth:     2,
														BorderColor:     "#F0F0F0",
														Padding:         18,
														BackgroundColor: "#FAFAFA",
													},
												},
											},
										},
									},

									// VEHICLE CATEGORY
									{
										Type: "INPUT",
										Data: bff.InputData{
											Id:          "vehicleCategory",
											Placeholder: "Enter your vehicle name",
											Style: bff.ViewData{
												BorderRadius:    16,
												BorderWidth:     2,
												BorderColor:     "#F0F0F0",
												Padding:         18,
												BackgroundColor: "#FAFAFA",
												MarginBottom:    20,
											},
										},
									},

									// VEHICLE NUMBER
									{
										Type: "INPUT",
										Data: bff.InputData{
											Id:          "vehicleNumber",
											Placeholder: "Vehicle Number",
											MaxLength:   15,
											Style: bff.ViewData{
												BorderRadius:    16,
												BorderWidth:     2,
												BorderColor:     "#F0F0F0",
												Padding:         18,
												BackgroundColor: "#FAFAFA",
											},
										},
									},

									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "Format: MH12AB1234",
											FontSize:  12,
											Color:     "#666666",
											MarginTop: 10,
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
										PaddingVertical: 18,
										BorderRadius:    16,
										AlignItems:      "center",
										MarginTop:       20,
									},
									Action: bff.ActionData{
										Type:     "NAVIGATE",
										Navigate: "/r3",
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
