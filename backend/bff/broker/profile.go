package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func ProfileScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Mock broker data
	broker := map[string]interface{}{
		"name":        "Rajesh Kumar",
		"companyName": "Kumar Logistics Solutions",
		"mobile":      "+919123456789",
		"email":       "rajesh@kumarlogistics.com",
		"avatar":      "https://i.pravatar.cc/150?img=12",
		"kycStatus":   "verified",
		"kycDocuments": map[string]string{
			"pan":             "uploaded",
			"aadhar":          "uploaded",
			"brokerLicense":   "pending",
			"companyDocuments":"uploaded",
		},
	}

	// Helper function for document status color
	getStatusColor := func(status string) string {
		if status == "uploaded" {
			return "#4CAF50"
		}
		return "#FF0000"
	}

	// Build UI
	ui := []bff.UISnippet{
		{
			Type: "SAFE_AREA",
			Children: []bff.UISnippet{
				{
					Type: "SCROLL",
					Data: bff.ViewData{
						FlexGrow:        1,
						PaddingHorizontal: 20,
						PaddingVertical:  20,
						BackgroundColor: "#FFFFFF",
					},
					Children: []bff.UISnippet{
						// Header
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection: "row",
								JustifyContent: "space-between",
								AlignItems: "center",
								MarginBottom: 20,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text: "Profile",
										FontSize: 24,
										FontWeight: "700",
										Color: "#1A1A1A",
									},
								},
								{
									Type: "BUTTON",
									Data: bff.ButtonData{
										Icon: bff.IconData{
											Name: "pencil-outline",
											Color: "#FF0000",
											Size: 22,
										},
									},
								},
							},
						},

						// Profile Card
						{
							Type: "VIEW",
							Data: bff.ViewData{
								BackgroundColor: "#FFFFFF",
								BorderRadius: 12,
								Padding: 20,
								ShadowColor: "#000",
								ShadowOpacity: 0.05,
								Elevation: 2,
							},
							Children: []bff.UISnippet{
								{
									Type: "IMAGE",
									Data: bff.ImageData{
										Url: broker["avatar"].(string),
										Width: 70,
										Height: 70,
										ResizeMode: "cover",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text: broker["name"].(string),
										FontSize: 20,
										Color: "#1A1A1A",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text: broker["companyName"].(string),
										FontSize: 14,
										Color: "#666666",
									},
								},
								{
									Type: "VIEW",
									Data: bff.ViewData{
										FlexDirection: "row",
										AlignItems: "center",
										BackgroundColor: "#E8F5E8",
										PaddingHorizontal: 12,
										PaddingVertical: 6,
										BorderRadius: 16,
										AlignSelf: "flex-start",
									},
									Children: []bff.UISnippet{
										{
											Type: "ICON",
											Data: bff.IconData{
												Name: "checkmark-circle",
												Size: 16,
												Color: "#4CAF50",
											},
										},
										{
											Type: "TEXT",
											Data: bff.TextData{
												Text: "Verified",
												Color: "#4CAF50",
												FontSize: 12,
												FontWeight: "600",
											},
										},
									},
								},
							},
						},

						// Contact Info
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection: "column",
								Gap: 12,
							},
							Children: []bff.UISnippet{
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text: "Mobile: " + broker["mobile"].(string),
										FontSize: 14,
										Color: "#666666",
									},
								},
								{
									Type: "TEXT",
									Data: bff.TextData{
										Text: "Email: " + broker["email"].(string),
										FontSize: 14,
										Color: "#666666",
									},
								},
							},
						},

						// Documents Section
						{
							Type: "VIEW",
							Data: bff.ViewData{
								BackgroundColor: "#FFFFFF",
								BorderRadius: 12,
								Padding: 16,
								MarginTop: 16,
							},
							Children: func() []bff.UISnippet {
								docs := []map[string]string{
									{"label": "PAN Card", "type": "pan"},
									{"label": "Aadhar Card", "type": "aadhar"},
									{"label": "Broker License", "type": "brokerLicense"},
									{"label": "Company Documents", "type": "companyDocuments"},
								}
								result := []bff.UISnippet{
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text: "KYC & Verification",
											FontSize: 18,
											Color: "#1A1A1A",
										},
									},
								}
								for _, d := range docs {
									status := broker["kycDocuments"].(map[string]string)[d["type"]]
									result = append(result, bff.UISnippet{
										Type: "VIEW",
										Data: bff.ViewData{
											BackgroundColor: "#FAFAFA",
											BorderRadius: 8,
											Padding: 12,
											MarginBottom: 12,
											AlignItems: "center",
										},
										Children: []bff.UISnippet{
											{
												Type: "TEXT",
												Data: bff.TextData{
													Text: d["label"],
													FontSize: 12,
													Color: "#333333",
													FontWeight: "500",
													MarginBottom: 6,
												},
											},
											{
												Type: "VIEW",
												Data: bff.ViewData{
													BackgroundColor: "#E8F5E8",
													PaddingHorizontal: 8,
													PaddingVertical: 2,
													BorderRadius: 10,
													MarginBottom: 8,
												},
												Children: []bff.UISnippet{
													{
														Type: "TEXT",
														Data: bff.TextData{
															Text: status,
															Color: getStatusColor(status),
															FontSize: 10,
															FontWeight: "600",
														},
													},
												},
											},
										},
									})
								}
								return result
							}(),
						},
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "Profile",
		UI:     ui,
	}

	c.JSON(200, response)
}
