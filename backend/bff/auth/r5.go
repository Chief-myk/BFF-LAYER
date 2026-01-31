package auth

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func R5Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R5_DOCUMENT_UPLOAD",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							Flex:            1,
							BackgroundColor: "#F8FAFC",
						},
						Children: []bff.UISnippet{
							// ================= MAIN CONTAINER =================
							{
								Type: "VIEW",
								Data: bff.ViewData{
									Flex:            1,
									BackgroundColor: "#F8FAFC",
									Padding:         24,
									PaddingBottom:   40,
								},
								Children: []bff.UISnippet{

									// ================= HEADER WITH BACK BUTTON =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection: "row",
											AlignItems:    "center",
											MarginBottom:  24,
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Position: "relative",
												},
												Children: []bff.UISnippet{
													{
														Type: "ICON_BUTTON",
														Data: bff.IconButtonData{
															Icon:  "arrow-left",
															Size:  24,
															Color: "#1E293B",
															OnPress: bff.ActionData{
																Type: "NAVIGATE_BACK",
															},
														},
													},
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Flex:       1,
													AlignItems: "center",
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "Upload Documents",
															FontSize:   24,
															FontWeight: "bold",
															Color:      "#1E293B",
														},
													},
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:  44,
													Height: 44,
												},
											},
										},
									},

									// ================= PROGRESS SECTION =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#FFFFFF",
											BorderRadius:    16,
											Padding:         20,
											MarginBottom:    24,
											ShadowColor:     "#000000",
											ShadowOffsetX:   0,
											ShadowOffsetY:   2,
											ShadowOpacity:   0.05,
											ShadowRadius:    8,
											Elevation:       2,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Complete Your Profile",
													FontSize:     18,
													FontWeight:   "600",
													Color:        "#1E293B",
													MarginBottom: 12,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													Width:        "100%",
													MarginBottom: 8,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Height:          8,
															BorderRadius:    4,
															BackgroundColor: "#E5E7EB",
															Overflow:        "hidden",
														},
														Children: []bff.UISnippet{
															{
																Type: "VIEW",
																Data: bff.ViewData{
																	Width:           "100%",
																	Height:          8,
																	BackgroundColor: "#FF0000",
																	BorderRadius:    4,
																},
															},
														},
													},
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													FlexDirection:  "row",
													JustifyContent: "space-between",
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:       "100% Complete",
															FontSize:   14,
															FontWeight: "600",
															Color:      "#FF0000",
														},
													},
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text:     "Step 6 of 6",
															FontSize: 14,
															Color:    "#6B7280",
														},
													},
												},
											},
										},
									},

									// ================= DOCUMENTS TITLE =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											MarginBottom: 16,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "Required Documents",
													FontSize:     20,
													FontWeight:   "bold",
													Color:        "#1E293B",
													MarginBottom: 8,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:     "Upload all required documents to complete your verification",
													FontSize: 14,
													Color:    "#6B7280",
												},
											},
										},
									},

									// ================= DOCUMENT GRID =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											FlexDirection:  "row",
											FlexWrap:       "wrap",
											JustifyContent: "space-between",
											MarginBottom:   32,
										},
										Children: []bff.UISnippet{
											documentCard("📄", "License", "Upload license document", "UPLOAD_LICENSE"),
											documentCard("📷", "Photo", "Upload profile photo", "UPLOAD_PHOTO"),
											documentCard("🚛", "RC", "Upload RC document", "UPLOAD_RC"),
											documentCard("🏥", "Fitness", "Upload fitness certificate", "UPLOAD_FITNESS"),
											documentCard("🌍", "National Permit", "Upload national permit", "UPLOAD_NATIONAL"),
											documentCard("🏛️", "State Permit", "Upload state permit", "UPLOAD_STATE"),
										},
									},

									// ================= INFO CARD =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#E6F2FF",
											BorderRadius:    12,
											Padding:         16,
											MarginBottom:    32,
											BorderWidth:     1,
											BorderColor:     "#B3D9FF",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:         "📋 Requirements",
													FontSize:     16,
													FontWeight:   "600",
													Color:        "#0066CC",
													MarginBottom: 8,
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text: `• All documents must be clear and readable
• Maximum file size: 5MB per document
• Supported formats: JPG, PNG, PDF
• Make sure documents are not expired`,
													FontSize:   14,
													Color:      "#0066CC",
													LineHeight: 20,
												},
											},
										},
									},

									// ================= SUBMIT BUTTON =================
									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Submit All Documents",
											Action: bff.ActionData{
												Type:            "API_CALL",
												Url:             "/api/v1/user/signup/driver",
												Method:          "POST",
												SuccessNavigate: "/r7",
												FailureNavigate: "",
											},
											Style: bff.ViewData{
												BackgroundColor: "#FF0000",
												BorderRadius:    12,
												PaddingVertical: 18,
												AlignItems:      "center",
												ShadowColor:     "#000000",
												ShadowOpacity:   0.2,
												ShadowOffsetX:   0,
												ShadowOffsetY:   4,
												ShadowRadius:    8,
												Elevation:       4,
											},
										},
									},

									// ================= FOOTER NOTE =================
									{
										Type: "VIEW",
										Data: bff.ViewData{
											AlignItems: "center",
											MarginTop:  24,
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "🔒 Your documents are securely encrypted",
													FontSize:  12,
													Color:     "#6B7280",
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
			},
		},
	}
	c.JSON(200, response)
}

// ================= HELPER =================
func documentCard(icon, title, subtitle, action string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Width:           "48%",
			BackgroundColor: "#FFFFFF",
			BorderRadius:    16,
			Padding:         16,
			MarginBottom:    16,
			ShadowColor:     "#000000",
			ShadowOffsetX:   0,
			ShadowOffsetY:   2,
			ShadowOpacity:   0.05,
			ShadowRadius:    6,
			Elevation:       1,
			AlignItems:      "center",
		},
		Children: []bff.UISnippet{
			{
				Type: "VIEW",
				Data: bff.ViewData{
					Width:           56,
					Height:          56,
					BorderRadius:    28,
					BackgroundColor: "#FFF0F0",
					JustifyContent:  "center",
					AlignItems:      "center",
					MarginBottom:    12,
				},
				Children: []bff.UISnippet{
					{
						Type: "TEXT",
						Data: bff.TextData{
							Text:     icon,
							FontSize: 24,
						},
					},
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         title,
					FontSize:     16,
					FontWeight:   "600",
					Color:        "#1E293B",
					TextAlign:    "center",
					MarginBottom: 4,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         subtitle,
					FontSize:     12,
					Color:        "#6B7280",
					TextAlign:    "center",
					MarginBottom: 16,
					LineHeight:   16,
				},
			},
			{
				Type: "BUTTON",
				Data: bff.ButtonData{
					Text: "Upload",
					Action: bff.ActionData{
						Type: "UPLOAD",
						Url:  action,
					},
					Style: bff.ViewData{
						BackgroundColor:   "#FF0000",
						BorderRadius:      8,
						PaddingVertical:   10,
						PaddingHorizontal: 16,
						Width:             "100%",
						AlignItems:        "center",
					},
				},
			},
		},
	}
}
