package types

type BoardGame struct {
    ObjectId int `xml:"objectid,attr"`
    YearPublished int `xml:"yearpublished"`
    MinPlayers int `xml:"minplayers"`
    MaxPlayers int `xml:"maxplayers"`
    PlayingTime int `xml:"playingtime"`
    MinPlayTime int `xml:"minplaytime"`
    MaxPlayTime int `xml:"maxplaytime"`
    Age int `xml:"age"`
    Name Name `xml:"name"`
    Description string `xml:"description"`
    ThumbnailUrl string `xml:"thumbnail"`
    ImageUrl string `xml:"image"`
    Publisher []Publisher `xml:"boardgamepublisher"`
    Category []Category `xml:"boardgamecategory"`
    Designer []Designer `xml:"boardgamedesigner"`
    Artist []Artist `xml:"boardgameartist"`
    Expansion []Expansion `xml:"boardgameexpansion"`
    Poll []Poll `xml:"poll"`
}
