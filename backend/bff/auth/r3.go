package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func R3Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R3",
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
							// ROOT VIEW
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex:            1,
									BackgroundColor: "#FFFFFF",
									Padding:         24,
									PaddingBottom:   40,
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
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Routes of Operation",
													FontSize:   24,
													FontWeight: "bold",
													Color:      "#1A1A1A",
													TextAlign:  "center",
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Select routes you typically operate on",
													FontSize:     16,
													Color:        "#666666",
													TextAlign:    "center",
													MarginTop:    8,
													MarginBottom: 20,
												},
											},

											// PROGRESS BAR
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:           "100%",
													Height:          6,
													BackgroundColor: "#F0F0F0",
													BorderRadius:    3,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "60%",
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
													Text:      "60% Complete",
													FontSize:  12,
													Color:     "#666666",
													TextAlign: "center",
													MarginTop: 8,
												},
											},
										},
									},

									// STEP INDICATOR
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor:   "#FF0000",
											PaddingHorizontal: 12,
											PaddingVertical:   6,
											BorderRadius:      16,
											MarginBottom:      20,
											Width:             "40%",
											AlignSelf:         "center",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Step 3 of 5",
													FontSize:   12,
													FontWeight: "600",
													Color:      "#FFFFFF",
													TextAlign:  "center",
												},
											},
										},
									},

									// POPULAR ROUTES TITLE
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Popular Routes",
											FontSize:     20,
											FontWeight:   "bold",
											Color:        "#1A1A1A",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Choose from frequently used routes or add your own",
											FontSize:     16,
											Color:        "#666666",
											MarginBottom: 25,
										},
									},

									// ROUTE CARD — MUMBAI DELHI
									RouteCard("Mumbai — Delhi", "Western Corridor"),

									// ROUTE CARD — DELHI CHENNAI
									RouteCard("Delhi — Chennai", "North-South Route"),

									// ROUTE CARD — BANGALORE DELHI
									RouteCard("Bangalore — Delhi", "Tech Corridor"),

									// ROUTE CARD — PUNE HYDERABAD
									RouteCard("Pune — Hyderabad", "Deccan Route"),

									// CUSTOM ROUTE
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BorderWidth:  2,
											BorderColor:  "#F0F0F0",
											BorderRadius: 12,
											Padding:      16,
											MarginBottom: 20,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Add Custom Route",
													FontSize:     16,
													FontWeight:   "600",
													Color:        "#666666",
													MarginBottom: 12,
												},
											},
											{
												Type: "INPUT",
												Data: bff.InputData{
													Id:          "custom_route",
													Placeholder: "e.g., Jaipur — Indore",
													Style: bff.ViewData{
														BorderWidth:     1,
														BorderColor:     "#E0E0E0",
														BorderRadius:    8,
														Padding:         12,
														BackgroundColor: "#F9F9F9",
														FontSize:        16,
														Color:           "#1A1A1A",
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
												Type:     "NAVIGATE",
												Navigate: "/r2",
											},
											Style: bff.ViewData{
												BackgroundColor:   "#FF0000",
												PaddingVertical:   16,
												PaddingHorizontal: 24,
												BorderRadius:      12,
												AlignItems:        "center",
												MarginTop:         20,
												ShadowColor:       "#000000",
												ShadowOpacity:     0.1,
												ShadowOffsetX:     0,
												ShadowOffsetY:     2,
												ShadowRadius:      4,
												Elevation:         2,
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

func RouteCard(title, subtitle string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BorderWidth:  2,
			BorderColor:  "#F0F0F0",
			BorderRadius: 12,
			Padding:      16,
			MarginBottom: 12,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       title,
					FontSize:   16,
					FontWeight: "600",
					Color:      "#1A1A1A",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:      subtitle,
					FontSize:  14,
					Color:     "#666666",
					MarginTop: 4,
				},
			},
		},
	}
}
