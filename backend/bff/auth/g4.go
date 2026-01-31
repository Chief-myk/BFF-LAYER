package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func G4Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	popularRoutes := []bff.RouteData{
		{ID: "mumbai-delhi", Name: "Mumbai — Delhi", Description: "Western Corridor"},
		{ID: "delhi-chennai", Name: "Delhi — Chennai", Description: "North-South Route"},
		{ID: "bangalore-delhi", Name: "Bangalore — Delhi", Description: "Tech Corridor"},
		{ID: "pune-hyderabad", Name: "Pune — Hyderabad", Description: "Deccan Route"},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "G4",
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
									// BACK BUTTON - Use ICON_BUTTON instead
									// {
									// 	Type: "ICON_BUTTON",
									// 	Data: bff.IconButtonData{
									// 		Icon: "arrow-left",
									// 		OnPress: bff.ActionData{Type: "NAVIGATE_BACK"},
									// 	},
									// },

									// HEADER & PROGRESS
									{
										Type: "VIEW",
										Data: bff.ViewData{
											AlignItems:   "center",
											MarginBottom: 30,
											MarginTop:    10, // Reduced margin
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:       "Routes of Operation",
													FontSize:   24,
													FontWeight: "bold",
													Color:      "#1A1A1A",
													MarginBottom: 8,
													TextAlign:  "center",
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:     "Select routes you typically operate on",
													FontSize: 16,
													Color:    "#666666",
													MarginBottom: 8,
													TextAlign: "center",
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width: "100%",
													AlignItems: "center",
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "100%",
															Height:          6,
															BackgroundColor: "#F0F0F0",
															BorderRadius:    3,
															MarginBottom:    8,
															Overflow:        "hidden",
														},
														Children: []bff.UISnippet{
															{
																Type: "VIEW",
																Data: bff.ViewData{
																	Width:           "60%", // 60% progress
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
															Text:       "60% Complete",
															FontSize:   12,
															Color:      "#666666",
															FontWeight: "500",
														},
													},
												},
											},
										},
									},

									// STEP INDICATOR
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#FF0000",
											PaddingHorizontal: 12,
											PaddingVertical:   6,
											BorderRadius:      16,
											MarginBottom:      20,
											AlignSelf:         "flex-start",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:          "Step 3 of 5",
													FontSize:      12,
													FontWeight:    "600",
													Color:         "#FFFFFF",
													TextAlign:     "center",
												},
											},
										},
									},

									// POPULAR ROUTES TITLE
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Popular Routes",
											FontSize:   20,
											FontWeight: "bold",
											Color:      "#1A1A1A",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Choose from frequently used routes or add your own",
											FontSize:   16,
											Color:      "#666666",
											MarginBottom: 12,
										},
									},

									// Routes Grid - UPDATED with proper text styling
									RoutesGrid(popularRoutes),

									// Custom Route Card
									CustomRouteCard(),

									// BUTTONS - SIMPLIFIED without icons
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection: "row",
											JustifyContent: "space-between",
											Gap: 12,
											MarginTop: 20,
										},
										Children: []bff.UISnippet{
											{
												Type: "BUTTON",
												Data: bff.ButtonData{
													Text: "Back",
													Action: bff.ActionData{Type: "NAVIGATE_BACK"},
													Style: bff.ViewData{
														Flex: 1,
														AlignItems: "center",
														JustifyContent: "center",
														PaddingVertical: 16,
														TextColor: "#1A1A1A",
														BorderRadius: 12,
														BorderWidth: 2,
														BorderColor: "#E5E5E5",
														BackgroundColor: "#FFFFFF",
													},
												},
											},
											{
												Type: "BUTTON",
												Data: bff.ButtonData{
													Text: "Continue",
													Action: bff.ActionData{
														Type:     "NAVIGATE",
														Navigate: "/g2",
													},
													Style: bff.ViewData{
														Flex: 2,
														AlignItems: "center",
														JustifyContent: "center",
														PaddingVertical: 16,
														BorderRadius: 12,
														BackgroundColor: "#FF0000",
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
		},
	}

	c.JSON(200, response)
}

// Helper: Popular Routes Grid - UPDATED with text color
func RoutesGrid(routes []bff.RouteData) bff.UISnippet {
	children := []bff.UISnippet{}
	for _, r := range routes {
		children = append(children, bff.UISnippet{
			Type: "BUTTON",
			Data: bff.ButtonData{
				Text: r.Name + "\n" + r.Description,
				Action: bff.ActionData{
					Type:    "SELECT_ROUTE",
					RouteID: r.ID,
				},
				Style: bff.ViewData{
					BackgroundColor: "#FAFAFA",
					Padding: 16,
					BorderRadius: 12,
					BorderWidth: 2,
					BorderColor: "#F0F0F0",
					MarginBottom: 12,
					TextColor: "#1A1A1A", // ADD THIS LINE
				},
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			MarginBottom: 20,
		},
		Children: children,
	}
}


// Helper: Custom Route Card
func CustomRouteCard() bff.UISnippet {
	return bff.UISnippet{
		Type: "INPUT",
		Data: bff.InputData{
			Id:          "customRoute",
			Placeholder: "Add Custom Route (e.g., Jaipur — Indore)",
			TextColor:   "#1A1A1A",
			FontSize:    16,
			FontWeight:  "500",
			Style: bff.ViewData{
				BackgroundColor: "#FAFAFA",
				Padding: 16,
				BorderRadius: 12,
				BorderWidth: 2,
				BorderColor: "#F0F0F0",
			},
		},
	}
}