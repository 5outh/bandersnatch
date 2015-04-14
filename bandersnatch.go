package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/xml"
    "bandersnatch/types"
)

func main() {
    resp, _ := http.Get("http://www.boardgamegeek.com/xmlapi/boardgame/35424")

    contents, _ := ioutil.ReadAll(resp.Body)

    // xml := string(contents[:])

    var boardGames types.BoardGames

	xml.Unmarshal(contents, &boardGames)

    fmt.Printf("%+v", boardGames)
}
