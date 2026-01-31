package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func AddTruckScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "AddTruck",
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

							// ---------------- HEADER ----------------
							{
								Type: "VIEW",
								Data: bff.ViewData{
									FlexDirection:   "row",
									AlignItems:      "center",
									JustifyContent:  "space-between",
									PaddingHorizontal: 20,
									PaddingVertical: 16,
									BorderBottomWidth: 1,
									BorderColor:     "#E0E0E0",
								},
								Children: []bff.UISnippet{
									{
										Type: "ICON",
										Data: bff.IconData{
											Name: "arrow-left",
											Size: 24,
										},
									},
									{
										Type: "TEXT",
										Data: bff.TextData{
											Text:       "Add Truck",
											FontSize:   20,
											FontWeight: "600",
											Color:      "#333333",
										},
									},
									{
										Type: "VIEW",
										Data: bff.ViewData{Width: 32},
									},
								},
							},

							// ---------------- FORM SECTION ----------------
							FormSection(),

							// ---------------- DOCUMENT UPLOAD ----------------
							DocumentSection(),

							// Spacer
							{
								Type: "VIEW",
								Data: bff.ViewData{Height: 120},
							},
						},
					},

					// ---------------- ADD BUTTON FIXED ----------------
					{
						Type: "VIEW",
						Data: bff.ViewData{
							Position:        "absolute",
							Bottom:          0,
							Left:            0,
							Right:           0,
							Padding:         20,
							BorderTopWidth:  1,
							BorderColor:     "#E0E0E0",
							BackgroundColor: "#FFFFFF",
						},
						Children: []bff.UISnippet{
							{
								Type: "BUTTON",
								Data: bff.ButtonData{
									Text: "Add Truck",
									Action: bff.ActionData{
										Type: "MODAL",
										Value: "SUCCESS_MODAL",
									},
									Style: bff.ViewData{
										BackgroundColor: "#ff0000",
										PaddingVertical: 16,
										BorderRadius:    12,
										AlignItems:      "center",
										ShadowColor:     "#ff0000",
										ShadowOpacity:   0.2,
										Elevation:       4,
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
func FormSection() bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{Padding: 20},
		Children: []bff.UISnippet{

			SectionTitle("Truck Information"),

			InputCard("truckNumber", "Truck Number *", "Enter truck number", "default"),
			SelectCard("truckType", "Truck Type", "Select truck type"),
			InputCard("capacity", "Capacity (tons)", "Enter capacity", "numeric"),
			InputCard("model", "Make / Model", "Enter make/model", "default"),

			ChipSelector("Body Type", []string{"Open", "Closed", "Container", "Flatbed"}),
			ChipSelector("Ownership Type", []string{"Owned", "Partner", "Attached"}),
		},
	}
}
func DocumentSection() bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{Padding: 20},
		Children: []bff.UISnippet{

			SectionTitle("Upload Truck Documents"),

			{
				Type: "VIEW",
				Data: bff.ViewData{
					FlexDirection: "row",
					FlexWrap:      "wrap",
					JustifyContent: "space-between",
				},
				Children: []bff.UISnippet{
					UploadCard("RC Certificate", "add-circle"),
					UploadCard("Insurance", "file-contract"),
					UploadCard("Fitness Certificate", "clipboard-check"),
					UploadCard("PUC Certificate", "smog"),
				},
			},
		},
	}
}
func SectionTitle(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:       text,
			FontSize:   18,
			FontWeight: "600",
			Color:      "#333333",
			MarginBottom: 20,
		},
	}
}

func InputCard(id, label, placeholder, keyboard string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BorderWidth:   1,
			BorderColor:   "#E0E0E0",
			BorderRadius:  12,
			Padding:       16,
			MarginBottom:  16,
			BackgroundColor: "#FFFFFF",
			Elevation:     2,
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text: label,
					FontSize: 14,
					FontWeight: "500",
					Color: "#333333",
					MarginBottom: 8,
				},
			},
			{
				Type: "INPUT",
				Data: bff.InputData{
					Id: id,
					Placeholder: placeholder,
					KeyboardType: keyboard,
				},
			},
		},
	}
}

func SelectCard(id, label, placeholder string) bff.UISnippet {
	return InputCard(id, label, placeholder, "default")
}

func ChipSelector(title string, options []string) bff.UISnippet {
	chips := []bff.UISnippet{}
	for _, opt := range options {
		chips = append(chips, bff.UISnippet{
			Type: "TEXT",
			Data: bff.TextData{
				Text: opt,
				PaddingHorizontal: 16,
				PaddingVertical: 8,
				BorderRadius: 20,
				BackgroundColor: "#F8F8F8",
				Color: "#666666",
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{MarginBottom: 16},
		Children: append([]bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text: title,
					FontSize: 14,
					FontWeight: "500",
					Color: "#333333",
					MarginBottom: 8,
				},
			},
		}, chips...),
	}
}

func UploadCard(label, icon string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Width: "48%",
			BorderWidth: 1,
			BorderColor: "#E0E0E0",
			BorderRadius: 12,
			Padding: 16,
			MarginBottom: 16,
			AlignItems: "center",
		},
		Children: []bff.UISnippet{
			{
				Type: "ICON",
				Data: bff.IconData{
					Name: icon,
					Size: 32,
					Color: "#ff0000",
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text: label,
					FontSize: 14,
					FontWeight: "500",
					Color: "#333333",
					MarginTop: 8,
				},
			},
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text: "JPG, PNG, PDF supported",
					FontSize: 12,
					Color: "#999999",
					MarginTop: 4,
				},
			},
		},
	}
}
