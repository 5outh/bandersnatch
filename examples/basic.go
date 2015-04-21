package main

import (
    "github.com/5outh/bandersnatch"
    "fmt"
)

func main() {
    boardGames, _ := bandersnatch.GetBoardGame(30000)

    fmt.Printf("%+v\n\n", boardGames)

    settlersSearch, _ := bandersnatch.GetSearch("Settlers of Catan")

    fmt.Printf("%+v", settlersSearch)
}
