package bff

import (
	"github.com/gin-gonic/gin"
)

func SplashScreenHandler(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")

	response := ScreenResponse{
		Status: "success",
		Screen: "SPLASH",
		UI: []UISnippet{
			{
				Type: "VIEW",
				Data: ViewData{
					Flex:            1,
					JustifyContent:  "center",
					AlignItems:      "center",
					BackgroundColor: "#ff0000",
					Padding:         24,
				},
				Children: []UISnippet{
					{
						Type: "VIEW",
						Data: ViewData{
							FlexDirection:  "row",
							AlignItems:     "center",
							JustifyContent: "center",
						},
						Children: []UISnippet{
							{
								Type: "IMAGE",
								Data: ImageData{
									Url:        "https://cdn.truckhai.com/rr.gif",
									Width:      120,
									Height:     120,
									ResizeMode: "contain",
									Animation:  "pulse",
								},
							},
							{
								Type: "TEXT",
								Data: TextData{
									Text:       "TruckHai",
									FontSize:   40,
									FontWeight: "bold",
									Color:      "#ffffff",
									MarginLeft: -10,
									MarginTop:  8,
								},
							},
						},
					},
				},
			},
			{
				Type: "NAVIGATE",
				Data: NavigateData{
					To:    "/(auth)/auth",
					After: 9000,
				},
			},
		},
	}

	c.JSON(200, response)
}
