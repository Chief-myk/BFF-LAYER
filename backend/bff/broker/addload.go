package broker

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func AddLoadScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Options for chips and dropdowns
	truckTypes := []string{"14 Tyre", "22 Tyre", "Trailer", "Container Truck", "Tanker"}
	bodyTypes := []string{"Open", "Closed", "Container", "Flatbed"}
	ownershipTypes := []string{"Owned", "Partner", "Attached"}

	// UI layout
	ui := []bff.UISnippet{
		// Header
		row(
			iconAtom("arrow-left", "#333"),
			text("Add Truck", 20, true, "#333"),
			spacerAtom(32),
		),

		// Truck Information Section
		sectionHeader("Truck Information", ""),
		formInput("Truck Number", "*", "Enter truck number"),
		formDropdown("Truck Type", "Select truck type", truckTypes),
		formInput("Capacity (tons)", "", "Enter capacity"),
		formInput("Make / Model", "", "Enter make/model"),
		formChipSelector("Body Type", bodyTypes),
		formChipSelector("Ownership Type", ownershipTypes),

		// Document Upload Section
		sectionHeader("Upload Truck Documents", ""),
		row(
			uploadCard("RC Certificate", "rc_certificate.pdf"),
			uploadCard("Insurance", "insurance.pdf"),
		),
		row(
			uploadCard("Fitness Certificate", "fitness_cert.pdf"),
			uploadCard("PUC Certificate", "puc_certificate.pdf"),
		),

		// Add Truck Button
		row(
			button("Add Truck", "#ff0000", "#FFFFFF"),
		),

		// Success Modal (hidden by default)
		modal("Truck Added Successfully!", "Your truck details have been saved to the database.",
			button("View My Trucks", "#ff0000", "#FFFFFF"),
			button("Add Another Truck", "#FFFFFF", "#ff0000"),
		),
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "addLoad",
		UI:     ui,
	}

	c.JSON(200, response)
}

// --- Helper functions to mimic React Native components in BFF ---

func row(children ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type:     "ROW",
		Children: children,
	}
}

func column(children ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type:     "COLUMN",
		Children: children,
	}
}

func text(value string, size int, bold bool, color string) bff.UISnippet {
	weight := "regular"
	if bold {
		weight = "bold"
	}
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Value:  value,
			Size:   size,
			Weight: weight,
			Color:  color,
		},
	}
}

func spacerAtom(width int) bff.UISnippet {
	return bff.UISnippet{
		Type: "SPACER",
		Data: map[string]int{
			"width": width,
		},
	}
}

func iconAtom(name string, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "ICON",
		Data: map[string]string{
			"name":  name,
			"color": color,
		},
	}
}

func formInput(label string, required string, placeholder string) bff.UISnippet {
	return column(
		text(label+" "+required, 14, true, "#333"),
		bff.UISnippet{
			Type: "INPUT",
			Data: map[string]string{
				"placeholder": placeholder,
				"color":       "#333",
			},
		},
	)
}

func formDropdown(label string, placeholder string, options []string) bff.UISnippet {
	return column(
		text(label, 14, true, "#333"),
		bff.UISnippet{
			Type: "DROPDOWN",
			Data: map[string]interface{}{
				"placeholder": placeholder,
				"options":     options,
			},
		},
	)
}

func formChipSelector(label string, options []string) bff.UISnippet {
	chips := []bff.UISnippet{}
	for _, option := range options {
		chips = append(chips, bff.UISnippet{
			Type: "CHIP",
			Data: map[string]string{
				"value": option,
				"color": "#666",
			},
		})
	}
	return column(
		text(label, 14, true, "#333"),
		row(chips...),
	)
}

func button(label, bgColor, textColor string) bff.UISnippet {
	return bff.UISnippet{
		Type: "BUTTON",
		Data: map[string]string{
			"label":    label,
			"bgColor":  bgColor,
			"textColor": textColor,
		},
	}
}

func modal(title, subtitle string, buttons ...bff.UISnippet) bff.UISnippet {
	return bff.UISnippet{
		Type: "MODAL",
		Data: map[string]string{
			"title":    title,
			"subtitle": subtitle,
		},
		Children: buttons,
	}
}

func uploadCard(title, fileName string) bff.UISnippet {
	return column(
		text(title, 14, true, "#333"),
		bff.UISnippet{
			Type: "UPLOAD",
			Data: map[string]string{
				"fileName": fileName,
			},
		},
	)
}

func sectionHeader(title, action string) bff.UISnippet {
	return row(
		text(title, 18, true, "#333"),
		text(action, 14, true, "#ff0000"),
	)
}
