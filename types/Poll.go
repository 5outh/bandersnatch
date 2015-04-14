package types

type Poll struct {
    Name string `xml:"name"`
    Title string `xml:"title"`
    TotalVotes int `xml:"totalvotes"`
    // Result []Result Not sure how to parse this yet...
}
