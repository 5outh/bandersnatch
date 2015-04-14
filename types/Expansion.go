package types

type Expansion struct {
    ObjectId int `xml:"objectid,attr"`
    Inbound bool `xml:"inbound,attr"`
    Name string `xml:",chardata"`
}
