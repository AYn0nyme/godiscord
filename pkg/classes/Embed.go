package classes

import "strconv"

type Embed struct {
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Thumbnail   thumnbail `json:"thumbnail"`
	Image       image     `json:"image"`
	Timestamp   string    `json:"timestamp"` // ISO8601 timestamp
	Color       int       `json:"color"`
	Footer      footer    `json:"footer"`
	Fields      []field   `json:"fields"`
}
type thumnbail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func (e Embed) SetTitle(Title string) Embed {
	e.Title = Title
	return e
}
func (e Embed) SetDescription(Description string) Embed {
	e.Description = Description
	return e
}
func (e Embed) SetURL(URL string) Embed {
	e.URL = URL
	return e
}
func (e Embed) SetTimestamp(Timestamp string) Embed {
	e.Timestamp = Timestamp
	return e
}

// Set the color of the embed from Hex Code. Don't include the #
func (e Embed) SetColor(HexColor string) Embed {
	color_in_int, err := strconv.ParseInt(HexColor, 16, 64)
	if err != nil {
		panic(err)
	}
	e.Color = int(color_in_int)
	return e
}
func (e Embed) SetFooter(Text, IconURL string) Embed {
	e.Footer.IconURL = IconURL
	e.Footer.Text = Text
	return e
}
func (e Embed) AddField(Name, Value string, Inline bool) Embed {
	e.Fields = append(e.Fields, field{Name: Name, Value: Value, Inline: Inline})
	return e
}
func (e Embed) SetThumbnail(URL string, Height, Width int) Embed {
	e.Thumbnail.URL = URL
	e.Thumbnail.Height = Height
	e.Thumbnail.Width = Width
	return e
}

func (e Embed) SetImage(URL string, Height, Width int) Embed {
	e.Image.URL = URL
	e.Image.Height = Height
	e.Image.Width = Width
	return e
}
