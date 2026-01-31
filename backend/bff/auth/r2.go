package auth

import (
	"backend/bff"

	"github.com/gin-gonic/gin"
)

func R2Screen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "R2_KYC_VERIFICATION",
		UI: []bff.UISnippet{
			{
				Type: "SAFE_AREA",
				Children: []bff.UISnippet{
					{
						Type: "SCROLL",
						Data: bff.ViewData{
							FlexGrow:          1,
							PaddingHorizontal: 20,
							PaddingVertical:   30,
							BackgroundColor:   "#F8FAFC",
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
									MarginBottom: 25,
								},
								Children: []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "KYC Verification",
											FontSize:     26,
											FontWeight:   "bold",
											Color:        "#1E293B",
											MarginBottom: 8,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:         "Complete your identity verification",
											FontSize:     14,
											Color:        "#6B7280",
											MarginBottom: 20,
										},
									},

									// PROGRESS
									{
										Type: "VIEW",
										Data: bff.ViewData{
											Width: "100%",
										},
										Children: []bff.UISnippet{
											{
												Type: "VIEW",
												Data: bff.ViewData{
													BackgroundColor: "#E5E7EB",
													BorderRadius:    3,
													PaddingVertical: 3,
												},
												Children: []bff.UISnippet{
													{
														Type: "VIEW",
														Data: bff.ViewData{
															Width:           "66%",
															BackgroundColor: "#FF0000",
															BorderRadius:    3,
														},
													},
												},
											},
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text:      "66% Complete",
													FontSize:  14,
													Color:     "#6B7280",
													MarginTop: 8,
												},
											},
										},
									},
								},
							},

							// PAN CARD
							kycSection(
								"PAN Card Details",
								"card-account-details",
								"panNumber",
								"Enter 10-digit PAN Number",
								"Format: ABCDE1234F",
								"/api/v1/kyc/verify-pan",
							),

							// AADHAAR
							kycSection(
								"Aadhaar Card Details",
								"fingerprint",
								"aadhaarNumber",
								"Enter 12-digit Aadhaar Number",
								"12-digit unique identification number",
								"/api/v1/kyc/verify-aadhaar",
							),

							// BANK DETAILS
							{
								Type: "VIEW",
								Data: cardStyle(),
								Children: []bff.UISnippet{

									sectionHeader("Bank Account Details", "bank"),

									labeledInput("Account Holder Name", "accountHolderName", "Enter full name as per bank"),
									labeledInput("Account Number", "accountNumber", "Enter bank account number"),
									labeledInput("IFSC Code", "ifscCode", "Enter 11-digit IFSC code"),

									{
										Type: "BUTTON",
										Data: bff.ButtonData{
											Text: "Verify",
											Style: bff.ViewData{
												BackgroundColor: "#FF0000",
												PaddingVertical: 12,
												BorderRadius:    8,
												AlignItems:      "center",
												MarginTop:       10,
											},
											Action: bff.ActionData{
												Type: "API_CALL",
												Url:  "/api/v1/kyc/verify-bank",
											},
										},
									},
								},
							},

							// CONTINUE
							{
								Type: "BUTTON",
								Data: bff.ButtonData{
									Text: "Continue to Next Step",
									Style: bff.ViewData{
										BackgroundColor: "#FF0000",
										PaddingVertical: 18,
										BorderRadius:    12,
										AlignItems:      "center",
										MarginTop:       25,
									},
									Action: bff.ActionData{
										Type:     "NAVIGATE",
										Navigate: "/r5",
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

// ---------- HELPERS ----------

func cardStyle() bff.ViewData {
	return bff.ViewData{
		BackgroundColor: "#FFFFFF",
		BorderRadius:    14,
		Padding:         20,
		MarginBottom:    20,
	}
}

func sectionHeader(title, icon string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			FlexDirection: "row",
			AlignItems:    "center",
			MarginBottom:  20,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       title,
					FontSize:   18,
					FontWeight: "700",
					Color:      "#1E293B",
				},
			},
		},
	}
}

func labeledInput(label, id, placeholder string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:         label,
					FontSize:     14,
					Color:        "#4B5563",
					MarginBottom: 6,
				},
			},
			{
				Type: "INPUT",
				Data: bff.InputData{
					Id:          id,
					Placeholder: placeholder,
					Style: bff.ViewData{
						BackgroundColor: "#F9FAFB",
						BorderRadius:    10,
						BorderWidth:     1,
						BorderColor:     "#E5E7EB",
						Padding:         14,
					},
				},
			},
		},
	}
}

func kycSection(title, icon, inputId, placeholder, hint, api string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: cardStyle(),
		Children: []bff.UISnippet{
			sectionHeader(title, icon),

			{
				Type: "INPUT",
				Data: bff.InputData{
					Id:          inputId,
					Placeholder: placeholder,
					Style: bff.ViewData{
						BackgroundColor: "#F9FAFB",
						BorderRadius:    10,
						BorderWidth:     1,
						BorderColor:     "#E5E7EB",
						Padding:         14,
					},
				},
			},

			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:      hint,
					FontSize:  12,
					Color:     "#9CA3AF",
					MarginTop: 6,
				},
			},

			{
				Type: "BUTTON",
				Data: bff.ButtonData{
					Text: "Verify",
					Style: bff.ViewData{
						BackgroundColor: "#FF0000",
						PaddingVertical: 12,
						BorderRadius:    8,
						AlignItems:      "center",
						MarginTop:       10,
					},
					Action: bff.ActionData{
						Type: "API_CALL",
						Url:  api,
					},
				},
			},
		},
	}
}
