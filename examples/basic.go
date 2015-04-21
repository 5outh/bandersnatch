package main

import (
    "github.com/5outh/bandersnatch"
    "fmt"
)

func main() {
    boardGames, _ := bandersnatch.GetBoardGame(30000)

    fmt.Printf("%+v\n\n", boardGames)

    request, _ := bandersnatch.NewSearchRequest(
        "Dominion",
        bandersnatch.OptionExact(true) )

    settlersSearch, _ := bandersnatch.GetSearch(request)

    fmt.Printf("%+v", settlersSearch)
}
