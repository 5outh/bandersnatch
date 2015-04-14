package types

type Publisher struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
