package driver

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func MyTripScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		{
			Type: "SAFE_AREA",
			Data: bff.ViewData{
				Flex:            1,
				BackgroundColor: "#F8FAFC",
			},
			Children: []bff.UISnippet{

				// STATUS BAR
				{
					Type: "STATUS_BAR",
					Data: bff.StatusBarData{
						BackgroundColor: "#ff0000",
						Style:           "light",
					},
				},

				// TAB BAR
				tabBar(),

				// TRIP LIST
				{
					Type: "SCROLL",
					Data: bff.ViewData{
						Flex:    1,
						Padding: 16,
					},
					Children: []bff.UISnippet{
						tripCard(currentTrip()),
						tripCard(upcomingTrip()),
						tripCard(completedTrip()),
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "myTrip",
		UI:     ui,
	}

	c.JSON(200, response)
}
func tabBar() bff.UISnippet {
	tabs := []string{"Current", "Upcoming", "Completed"}
	children := []bff.UISnippet{}

	for i, tab := range tabs {
		active := i == 0

		children = append(children, bff.UISnippet{
			Type: "PRESSABLE_CARD",
			Data: bff.ViewData{
				Flex:            1,
				AlignItems:      "center",
				PaddingVertical: 12,
			},
			Children: []bff.UISnippet{
				{
					Type: "TEXT",
					Data: bff.TextData{
						Text:       tab,
						FontSize:   16,
						FontWeight: "600",
						Color:      ternary(active, "#ff0000", "#6B7280"),
					},
				},
				{
					Type: "VIEW",
					Data: bff.ViewData{
						Height:          ternaryInt(active, 3, 0),
						Width:           "60%",
						BackgroundColor: "#ff0000",
						BorderRadius:    2,
						MarginTop:       8,
					},
				},
			},
		})
	}

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			FlexDirection:     "row",
			BackgroundColor:   "#FFFFFF",
			PaddingHorizontal: 20,
			BorderBottomWidth: 1,
			BorderColor:       "#F1F5F9",
		},
		Children: children,
	}
}
func statusPill(text, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			PaddingHorizontal: 12,
			PaddingVertical:   6,
			BorderRadius:      12,
			BackgroundColor:   color + "15",
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       text,
					FontSize:   12,
					FontWeight: "600",
					Color:      color,
				},
			},
		},
	}
}

func tripCard(trip map[string]string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#FFFFFF",
			BorderRadius:    16,
			Padding:         20,
			MarginBottom:    16,
			ShadowColor:     "#000",
			ShadowOpacity:   0.1,
			ShadowRadius:    3,
			Elevation:       4,
		},
		Children: []bff.UISnippet{

			{
				Type: "VIEW",
				Data: rowBetween(),
				Children: []bff.UISnippet{
					{
						Type: "VIEW",
						Children: []bff.UISnippet{
							title(trip["route"]),
							subtitle(trip["distance"] + " • " + trip["time"]),
						},
					},
					statusPill(trip["status"], trip["statusColor"]),
				},
			},

			spacer(16),
			iconRow("calendar-outline", trip["pickup"]),
			spacer(8),
			iconRowGreen("cash-outline", trip["payment"]),
			spacer(16),
			viewDetailsButton(),
		},
	}
}

func spacer(height int) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			Height: height,
		},
	}
}

func iconRow(icon, text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			FlexDirection: "row",
			AlignItems:    "center",
			Gap:           6,
		},
		Children: []bff.UISnippet{
			iconAtom(icon, "#6B7280"),
			subtitle(text),
		},
	}
}

func iconRowGreen(icon, text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			FlexDirection: "row",
			AlignItems:    "center",
			Gap:           6,
		},
		Children: []bff.UISnippet{
			iconAtom(icon, "#6B7280"),
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       text,
					FontSize:   14,
					FontWeight: "600",
					Color:      "#16A34A",
				},
			},
		},
	}
}
func viewDetailsButton() bff.UISnippet {
	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			FlexDirection:   "row",
			JustifyContent:  "center",
			AlignItems:      "center",
			PaddingVertical: 12,
			BorderTopWidth:  1,
			BorderColor:     "#F1F5F9",
		},
		Children: []bff.UISnippet{
			{
				Type: "TEXT",
				Data: bff.TextData{
					Text:       "View Trip Details",
					FontSize:   16,
					FontWeight: "600",
					Color:      "#ff0000",
				},
			},
			iconAtom("chevron-forward", "#ff0000"),
		},
	}
}
func currentTrip() map[string]string {
	return map[string]string{
		"route":       "Mumbai → Delhi",
		"pickup":      "15 Dec 2024, 08:00 AM",
		"payment":     "₹45,000",
		"status":      "In Transit",
		"statusColor": "#3B82F6",
		"distance":    "1,450 km",
		"time":        "28 hrs",
	}
}

func upcomingTrip() map[string]string {
	return map[string]string{
		"route":       "Chennai → Hyderabad",
		"pickup":      "18 Dec 2024, 09:00 AM",
		"payment":     "₹52,000",
		"status":      "Assigned",
		"statusColor": "#EA580C",
		"distance":    "1,200 km",
		"time":        "24 hrs",
	}
}

func completedTrip() map[string]string {
	return map[string]string{
		"route":       "Ahmedabad → Mumbai",
		"pickup":      "12 Dec 2024, 07:00 AM",
		"payment":     "₹35,000",
		"status":      "Completed",
		"statusColor": "#6B7280",
		"distance":    "530 km",
		"time":        "12 hrs",
	}
}

func ternary(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func ternaryInt(condition bool, a int, b int) int {
	if condition {
		return a
	}
	return b
}

func iconAtom(name string, color string) bff.UISnippet {
	return bff.UISnippet{
		Type: "ICON",
		Data: bff.IconData{
			Name:  name,
			Size:  20,
			Color: color,
		},
	}
}
