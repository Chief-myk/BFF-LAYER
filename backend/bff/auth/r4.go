package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func R4Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R4",
		UI: []bff.UISnippet{

			// ROOT CONTAINER
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Flex:            1,
					BackgroundColor: "#FFFFFF",
					Padding:         24,
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

							// TITLE
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       "Vehicle Types",
									FontSize:   24,
									FontWeight: "bold",
									Color:      "#1A1A1A",
									TextAlign:  "center",
								},
							},

							// SUBTITLE
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:         "Which vehicles are you interested in?",
									FontSize:     16,
									Color:        "#666666",
									TextAlign:    "center",
									MarginTop:    8,
									MarginBottom: 25,
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
											Width:           "80%",
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
									Text:      "80% Complete",
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
							BackgroundColor:   "#FFF0F0",
							PaddingHorizontal: 12,
							PaddingVertical:   6,
							BorderRadius:      16,
							MarginBottom:      20,
							Width:             "40%",
						},
						Children: []bff.UISnippet{
							{
								Type: "TEXT",
								Data: bff.TextData{
									Text:       "Step 4 of 5",
									FontSize:   12,
									FontWeight: "600",
									Color:      "#FF0000",
									TextAlign:  "center",
								},
							},
						},
					},

					// SECTION TITLE
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:         "Available Vehicles",
							FontSize:     20,
							FontWeight:   "bold",
							Color:        "#1A1A1A",
							MarginBottom: 8,
						},
					},
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:         "Choose the types of vehicles you work with",
							FontSize:     16,
							Color:        "#666666",
							MarginBottom: 25,
						},
					},

					// VEHICLE CARDS
					VehicleCard("Open Truck", "General goods transport", "9-16 tons"),
					VehicleCard("Container", "Sealed cargo transport", "20-40 ft"),
					VehicleCard("Trailer", "Heavy load transport", "20-40 tons"),
					VehicleCard("Mini Truck", "Small cargo transport", "1-2 tons"),
					VehicleCard("Tempo", "Local deliveries", "0.5-1 ton"),
					VehicleCard("Auto", "Last mile delivery", "Up to 500 kg"),

					// CUSTOM VEHICLE
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
									Text:       "Other Vehicle Type",
									FontSize:   16,
									FontWeight: "600",
									Color:      "#666666",
								},
							},

							{
								Type: "INPUT",
								Data: bff.InputData{
									Id:          "custom_vehicle",
									Placeholder: "e.g., Refrigerated Truck, Tanker",
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
								Type:            "navigate",
								SuccessNavigate: "/auth/r5",
							},
						},
					},
				},
			},
		},
	}

	c.JSON(200, response)
}

func VehicleCard(title, description, capacity string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Width:        "48%",
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
					Text:         description,
					FontSize:     12,
					Color:        "#666666",
					MarginBottom: 8,
				},
			},

			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:     capacity,
					FontSize: 10,
					Color:    "#374151",
				},
			},
		},
	}
}
