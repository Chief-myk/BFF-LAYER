package bff

type UISnippet struct {
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	Children []UISnippet `json:"children,omitempty"`
}

type ScreenResponse struct {
	Status  string      `json:"status"`
	Screen  string      `json:"screen"`
	UI      []UISnippet `json:"ui"`
	Data    interface{} `json:"data,omitempty"`    // Add this
	Message string      `json:"message,omitempty"` // Add this
}

type ViewData struct {
	Flex             int         `json:"flex,omitempty"`
	FlexGrow         int         `json:"flexGrow,omitempty"`
	FlexDirection    string      `json:"flexDirection,omitempty"`
	AlignItems       string      `json:"alignItems,omitempty"`
	JustifyContent   string      `json:"justifyContent,omitempty"`
	BackgroundColor  string      `json:"backgroundColor,omitempty"`
	Bottom           int         `json:"bottom,omitempty"`
	PaddingTop       int         `json:"paddingTop,omitempty"`
	MarginHorizontal  int         `json:"marginHorizontal,omitempty"`
	Left             int         `json:"left,omitempty"`
	Top              int         `json:"top,omitempty"`
	Horizontal    	 bool			`json:"horizontal,omitempty"`
	MaxHeight        interface{} `json:"maxHeight,omitempty"`
	ShowsHorizontalScrollIndicator bool `json:"showsHorizontalScrollIndicator,omitempty"`
	Right            int         `json:"right,omitempty"`
	FlexWrap         string      `json:"flexWrap,omitempty"`
	Padding          int         `json:"padding,omitempty"`
	BorderTopWidth   int         `json:"borderTopWidth,omitempty"`
	Color            string      `json:"color,omitempty"`
	PaddingLeft      int         `json:"paddingLeft,omitempty"`
	TextColor        string      `json:"textColor,omitempty"`
	BorderBottomWidth int        `json:"borderBottomWidth,omitempty"`
	PaddingHorizontal int        `json:"paddingHorizontal,omitempty"`
	PaddingVertical  int         `json:"paddingVertical,omitempty"`
	MarginBottom     int         `json:"marginBottom,omitempty"`
	MarginTop        int         `json:"marginTop,omitempty"`
	MarginRight      int         `json:"marginRight,omitempty"`
	BorderRadius     int         `json:"borderRadius,omitempty"`
	PaddingBottom    int         `json:"paddingBottom,omitempty"`
	FontSize         int         `json:"fontSize,omitempty"`
	BorderWidth      int         `json:"borderWidth,omitempty"`
	BorderColor      string      `json:"borderColor,omitempty"`
	Width            interface{} `json:"width,omitempty"`
	Height           interface{} `json:"height,omitempty"`
	MarginLeft       int         `json:"marginLeft,omitempty"`
	BorderLeftWidth  int         `json:"borderLeftWidth,omitempty"`
	BorderLeftColor  string      `json:"borderLeftColor,omitempty"`
	Gap              int         `json:"gap,omitempty"`
	ShowsVerticalScrollIndicator bool        `json:"showsVerticalScrollIndicator,omitempty"`
	BorderTopLeftRadius  int         `json:"borderTopLeftRadius,omitempty"`
	BorderTopRightRadius int         `json:"borderTopRightRadius,omitempty"`
	Margin           int         `json:"margin,omitempty"`
	Overflow         string      `json:"overflow,omitempty"`
	ShadowColor      string      `json:"shadowColor,omitempty"`
	ShadowOffsetX    int         `json:"shadowOffsetX,omitempty"`
	ShadowOffsetY    int         `json:"shadowOffsetY,omitempty"`
	ShadowOpacity    float64     `json:"shadowOpacity,omitempty"`
	ShadowRadius     int         `json:"shadowRadius,omitempty"`
	BorderTopColor	string      `json:"borderTopColor,omitempty"`
	Elevation        int         `json:"elevation,omitempty"`
	Position         string      `json:"position,omitempty"`
	ZIndex           int         `json:"zIndex,omitempty"`
	AlignSelf        string      `json:"alignSelf,omitempty"`
}

type ImageData struct {
	Url        string `json:"url"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	ResizeMode string `json:"resizeMode"`
	Animation  string `json:"animation"`
}

type TextData struct {
	Text              string  `json:"text"`
	FontSize          int     `json:"fontSize,omitempty"`
	FontWeight        string  `json:"fontWeight,omitempty"`
	Color             string  `json:"color,omitempty"`
	Opacity		 float64 `json:"opacity,omitempty"`
	FlexWrap          string  `json:"flexWrap,omitempty"`    // Add this line
	FlexShrink        int     `json:"flexShrink,omitempty"`  // Add this line
	FontFamily        string  `json:"fontFamily,omitempty"`  // Add this line
	Width             string  `json:"width,omitempty"`       // Add this line
	TextAlign         string  `json:"textAlign,omitempty"`
	Bold              bool    `json:"bold,omitempty"`
	Size              int     `json:"size,omitempty"`
	Flex 		   int     `json:"flex,omitempty"`
	Weight            string  `json:"weight,omitempty"`
	Value             string  `json:"value,omitempty"`
	MarginLeft        int     `json:"marginLeft,omitempty"`
	MarginTop         int     `json:"marginTop,omitempty"`
	MarginBottom      int     `json:"marginBottom,omitempty"`
	MarginRight       int     `json:"marginRight,omitempty"`
	LetterSpacing     float64 `json:"letterSpacing,omitempty"`
	LineHeight        int     `json:"lineHeight,omitempty"`
	BackgroundColor   string  `json:"backgroundColor,omitempty"`
	PaddingHorizontal int     `json:"paddingHorizontal,omitempty"`
	PaddingVertical   int     `json:"paddingVertical,omitempty"`
	BorderRadius      int     `json:"borderRadius,omitempty"`
	AlignSelf         string  `json:"alignSelf,omitempty"`
}

type NavigateData struct {
	To    string `json:"to"`
	After int    `json:"after"`
}

type InputData struct {
	Id           string   `json:"id"`
	Placeholder  string   `json:"placeholder"`
	KeyboardType string   `json:"keyboardType"`
	MaxLength    int      `json:"maxLength"`
	Style        ViewData `json:"style,omitempty"`
	TextColor    string   `json:"textColor,omitempty"`
	FontSize     int      `json:"fontSize,omitempty"`
	FontWeight   string   `json:"fontWeight,omitempty"`
}

type ButtonData struct {
	Text     string     `json:"text"`
	Disabled bool       `json:"disabled"`
	Action   ActionData `json:"action"`
	Style    ViewData   `json:"style,omitempty"`
	Icon     IconData   `json:"icon,omitempty"`
}

type ActionData struct {
	Type            string `json:"type"`
	Url             string `json:"url"`
	Method          string `json:"method,omitempty"`
	SuccessNavigate string `json:"successNavigate,omitempty"`
	FailureNavigate string `json:"failureNavigate,omitempty"`
	Navigate        string `json:"navigate,omitempty"`
	Value           string `json:"value,omitempty"`
	RouteID         string `json:"routeId,omitempty"`
	Data            map[string]interface{} `json:"data,omitempty"` // Add this line
	To              string `json:"to,omitempty"`
}

type AlertData struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type OtpInputData struct {
	Id                 string `json:"id"`
	Length             int    `json:"length"`
	BoxSpacing         int    `json:"boxSpacing,omitempty"`
	BoxBorderColor     string `json:"boxBorderColor,omitempty"`
	BoxBorderRadius    int    `json:"boxBorderRadius,omitempty"`
	BoxBackgroundColor string `json:"boxBackgroundColor,omitempty"`
	BoxWidth           int    `json:"boxWidth,omitempty"`
	BoxHeight          int    `json:"boxHeight,omitempty"`
}

type ResendOtpData struct {
	Timer     int `json:"timer"`
	MarginTop int `json:"marginTop,omitempty"`
}

type StatusBarData struct {
	BackgroundColor string `json:"backgroundColor"`
	Style           string `json:"style"`
}

type IconButtonData struct {
	Icon    string     `json:"icon"`
	OnPress ActionData `json:"onPress"`
	Size    int        `json:"size,omitempty"`
	Color   string     `json:"color,omitempty"`
}

type CardData struct {
	BackgroundColor string     `json:"backgroundColor,omitempty"`
	Padding         int        `json:"padding,omitempty"`
	BorderRadius    int        `json:"borderRadius,omitempty"`
	BorderWidth     int        `json:"borderWidth,omitempty"`
	BorderColor     string     `json:"borderColor,omitempty"`
	MarginTop       int        `json:"marginTop,omitempty"`
	Shadow          bool       `json:"shadow,omitempty"`
	OnPress         ActionData `json:"onPress"`
}

type PressableCardData struct {
	CardData
	Children []UISnippet `json:"children,omitempty"`
}

type TouchableOpacityData struct {
	Style   ViewData   `json:"style"`
	OnPress ActionData `json:"onPress"`
}

type IconData struct {
	Name            string `json:"name"`
	Size            int    `json:"size"`
	Color           string `json:"color,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	ContainerSize   int    `json:"containerSize,omitempty"`
	BorderRadius    int    `json:"borderRadius,omitempty"`
	MarginRight     int    `json:"marginRight,omitempty"`
	Padding         int    `json:"padding,omitempty"`
}

type TextButtonData struct {
	Text       string     `json:"text"`
	Color      string     `json:"color,omitempty"`
	FontSize   int        `json:"fontSize,omitempty"`
	FontWeight string     `json:"fontWeight,omitempty"`
	OnPress    ActionData `json:"onPress"`
}

type IconSnippet struct {
	Type string   `json:"type"`
	Data IconData `json:"data"`
}

type HomeScreenData struct {
	IsTripStarted     bool              `json:"isTripStarted"`
	LocationSharing   bool              `json:"locationSharing"`
	TripStatus        string            `json:"tripStatus"`
	DocumentsUploaded map[string]bool   `json:"documentsUploaded"`
	ActiveTrip        map[string]string `json:"activeTrip"`
	QuickActions      []QuickAction     `json:"quickActions,omitempty"`
	RecentActivities  []RecentActivity  `json:"recentActivities,omitempty"`
}

type QuickAction struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type RecentActivity struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    string `json:"time"`
	Icon    string `json:"icon"`
	Color   string `json:"color"`
}

type ActionRequest struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type ActionResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type RouteData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}