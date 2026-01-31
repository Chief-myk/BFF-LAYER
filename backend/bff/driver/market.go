package driver

import (
	"backend/bff"
	"github.com/gin-gonic/gin"
)

func MarketScreen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	ui := []bff.UISnippet{
		{
			Type: "SAFE_AREA",
			Children: []bff.UISnippet{

				// STATUS BAR
				{
					Type: "STATUS_BAR",
					Data: bff.StatusBarData{
						BackgroundColor: "#ffffff",
						Style:           "dark",
					},
				},

				// MAIN SCROLL
				{
					Type: "SCROLL",
					Data: bff.ViewData{
						FlexGrow:        1,
						BackgroundColor: "#f8f9fa",
					},
					Children: []bff.UISnippet{

						// HEADER (SEARCH + FILTER)
						{
							Type: "VIEW",
							Data: bff.ViewData{
								FlexDirection:     "row",
								AlignItems:        "center",
								Padding:           16,
								BackgroundColor:   "#ffffff",
								BorderBottomWidth: 1,
								BorderColor:       "#f0f0f0",
								Gap:               12,
							},
							Children: []bff.UISnippet{
								searchBox(),
								filterButton(),
							},
						},

						// TABS
						filterTabs(),

						// LOAD LIST
						loadCard(
							"Mumbai → Delhi",
							"₹28,000 – ₹32,000",
							"2024-09-20 • 10:00 AM",
							"Container Truck",
							"Electronics • 8 Tons",
							"1,416 km",
							"4 hrs left",
							95,
						),

						loadCard(
							"Chennai → Bangalore",
							"₹35,000 – ₹40,000",
							"2024-09-19 • 02:00 PM",
							"Trailer Truck",
							"FMCG Goods • 12 Tons",
							"346 km",
							"2 hrs left",
							88,
						),
					},
				},
			},
		},
	}

	response := bff.ScreenResponse{
		Status: "success",
		Screen: "market",
		UI:     ui,
	}

	c.JSON(200, response)
}

func searchBox() bff.UISnippet {
	return bff.UISnippet{
		Type: "INPUT",
		Data: bff.InputData{
			Id:          "search",
			Placeholder: "Search by city or route",
			Style: bff.ViewData{
				Flex:            1,
				BackgroundColor: "#f5f5f5",
				BorderRadius:    10,
				Padding:         12,
			},
			FontSize: 16,
		},
	}
}

func filterButton() bff.UISnippet {
	return bff.UISnippet{
		Type: "ICON_BUTTON",
		Data: bff.IconButtonData{
			Icon: "options-outline",
			OnPress: bff.ActionData{
				Type: "ACTION",
			},
		},
	}
}

func filterTabs() bff.UISnippet {
	tabs := []string{
		"All Loads",
		"Recommended",
		"Nearby",
		"High Paying",
		"Urgent",
		"My Bids",
	}

	children := []bff.UISnippet{}

	for _, tab := range tabs {
		children = append(children, bff.UISnippet{
			Type: "TEXT",
			Data: bff.TextData{
				Text:              tab,
				FontSize:          14,
				FontWeight:        "500",
				Color:             "#666",
				PaddingHorizontal: 16,
				PaddingVertical:   8,
				BackgroundColor:   "#f5f5f5",
				BorderRadius:      20,
			},
		})
	}

	return bff.UISnippet{
		Type: "SCROLL",
		Data: bff.ViewData{
			FlexDirection: "row",
			Padding:       16,
			Gap:           8,
		},
		Children: children,
	}
}

func loadCard(
	route string,
	budget string,
	time string,
	vehicle string,
	cargo string,
	distance string,
	timeLeft string,
	match int,
) bff.UISnippet {

	return bff.UISnippet{
		Type: "VIEW",
		Data: bff.ViewData{
			BackgroundColor: "#ffffff",
			BorderRadius:    12,
			Padding:         16,
			ShadowColor:     "#000",
			ShadowOpacity:   0.1,
			ShadowRadius:    4,
			Elevation:       3,
			MarginBottom:    12,
		},
		Children: []bff.UISnippet{

			// Match Row
			{
				Type: "VIEW",
				Data: rowBetween(),
				Children: []bff.UISnippet{
					badgeText(match),
					subtitleRed(timeLeft),
				},
			},

			// Route + Budget
			{
				Type: "VIEW",
				Data: rowBetween(),
				Children: []bff.UISnippet{
					title(route),
					price(budget),
				},
			},

			subtitle(time),
			subtitle(vehicle),
			subtitle(cargo),
			subtitle(distance),

			// CTA
			actionLink("View Load Details"),
		},
	}
}

func badgeText(match int) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:              string(rune(match)) + "% Match",
			FontSize:          12,
			Color:             "#ffffff",
			BackgroundColor:   "#4CAF50",
			PaddingHorizontal: 8,
			PaddingVertical:   4,
			BorderRadius:      6,
		},
	}
}

func subtitleRed(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:       text,
			Color:      "#FF3B30",
			FontSize:   12,
			FontWeight: "500",
		},
	}
}

func price(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT",
		Data: bff.TextData{
			Text:       text,
			FontSize:   16,
			FontWeight: "700",
			Color:      "#FF3B30",
		},
	}
}

func actionLink(text string) bff.UISnippet {
	return bff.UISnippet{
		Type: "TEXT_BUTTON",
		Data: bff.TextButtonData{
			Text:       text,
			Color:      "#FF3B30",
			FontSize:   14,
			FontWeight: "600",
			OnPress: bff.ActionData{
				Type: "NAVIGATE",
				To:   "loadDetails",
			},
		},
	}
}
