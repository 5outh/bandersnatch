package types

type Mechanic struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
