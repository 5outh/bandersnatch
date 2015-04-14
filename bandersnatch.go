package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/xml"
)

type BoardGames struct {
    TermsOfUse string `xml:"termsofuse,attr"`
    BoardGame []BoardGame `xml:"boardgame"`
}

type Name struct {
    Primary string `xml:"primary,attr"`
    SortIndex string `xml:"sortindex,attr"`
    Name string `xml:",chardata"`
}

type Publisher struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Category struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Designer struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Artist struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Expansion struct {
    ObjectId int `xml:"objectid,attr"`
    Inbound bool `xml:"inbound,attr"`
    Name string `xml:",chardata"`
}

type Subdomain struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Version struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Family struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Mechanic struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}

type Poll struct {
    Name string `xml:"name"`
    Title string `xml:"title"`
    TotalVotes int `xml:"totalvotes"`
    // Result []Result Not sure how to parse this yet...
}

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

func main() {
    resp, _ := http.Get("http://www.boardgamegeek.com/xmlapi/boardgame/35424")

    contents, _ := ioutil.ReadAll(resp.Body)

    // xml := string(contents[:])

    var boardGames BoardGames

	xml.Unmarshal(contents, &boardGames)

    fmt.Printf("%+v", boardGames)
}
