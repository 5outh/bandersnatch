package types

type Name struct {
    Primary string `xml:"primary,attr"`
    SortIndex string `xml:"sortindex,attr"`
    Name string `xml:",chardata"`
}
