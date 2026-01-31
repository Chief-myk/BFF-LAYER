package auth

import (
	"backend/bff"

	"github.com/gin-gonic/gin"
)

func RegistrationRoleHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "REGISTRATION_ROLE",
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
						},
						Children: []bff.UISnippet{

							// LOGO
							{
								Type: "VIEW",
								Data: bff.ViewData{
									AlignItems: "center",
									MarginTop:  20,
								},
								Children: []bff.UISnippet{
									{
										Type: "IMAGE",
										Data: bff.ImageData{
											Url:        "@/assets/images/logo.png",
											Width:      180,
											Height:     100,
											ResizeMode: "contain",
										},
									},
								},
							},

							// HEADER
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
											Text:         "Choose Your Role",
											FontSize:     24,
											FontWeight:   "bold",
											Color:        "#1A1A1A",
											TextAlign:    "center",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:      "Select how you want to use TruckHai to continue registration",
											FontSize:  16,
											Color:     "#666666",
											TextAlign: "center",
										},
									},
								},
							},

							// ROLES CONTAINER
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Width:        "100%",
									MarginBottom: 40,
								},
								Children: []bff.UISnippet{

									// BROKER CARD
									{
										Type: "PRESSABLE_CARD",
										Data: bff.CardData{
											BackgroundColor: "#EFF6FF",
											Padding:         20,
											BorderRadius:    16,
											BorderWidth:     1,
											BorderColor:     "#E5E5E5",
											Shadow:          true,
											OnPress: bff.ActionData{
												Type:     "SET_USER_ROLE",
												Value:    "broker",
												Navigate: "/g1",
											},
										},
										Children: []bff.UISnippet{
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:            "account-tie",
													Size:            32,
													Color:           "#155DFC",
													BackgroundColor: "rgba(19, 29, 222, 0.1)",
													ContainerSize:   50,
													BorderRadius:    25,
													MarginRight:     16,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Flex: 1,
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:         "Broker",
															FontSize:     18,
															FontWeight:   "600",
															Color:        "#1A1A1A",
															MarginBottom: 4,
														},
													},
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:     "Manages freight and transport operations.",
															FontSize: 14,
															Color:    "#666666",
														},
													},
												},
											},
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:  "chevron-right",
													Size:  24,
													Color: "#131DDE",
												},
											},
										},
									},

									// DRIVER CARD
									{
										Type: "PRESSABLE_CARD",
										Data: bff.CardData{
											BackgroundColor: "#FAF5FF",
											Padding:         20,
											BorderRadius:    16,
											BorderWidth:     1,
											BorderColor:     "#E5E5E5",
											Shadow:          true,
											MarginTop:       16,
											OnPress: bff.ActionData{
												Type:     "SET_USER_ROLE",
												Value:    "driver",
												Navigate: "/r6",
											},
										},
										Children: []bff.UISnippet{
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:            "truck-fast",
													Size:            32,
													Color:           "#9810FA",
													BackgroundColor: "rgba(19, 29, 222, 0.1)",
													ContainerSize:   50,
													BorderRadius:    25,
													MarginRight:     16,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Flex: 1,
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:         "Driver",
															FontSize:     18,
															FontWeight:   "600",
															Color:        "#1A1A1A",
															MarginBottom: 4,
														},
													},
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:     "Transporting goods safely and on time.",
															FontSize: 14,
															Color:    "#666666",
														},
													},
												},
											},
											{
												Type: "ICON",
												Data: bff.IconData{
													Name:  "chevron-right",
													Size:  24,
													Color: "#131DDE",
												},
											},
										},
									},
								},
							},

							// BACK LINK
							{
								Type: "TEXT_BUTTON",
								Data: bff.TextButtonData{
									Text:       "Back to mobile number",
									Color:      "#FF0000",
									FontSize:   14,
									FontWeight: "500",
									OnPress: bff.ActionData{
										Type: "NAVIGATE_BACK",
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
