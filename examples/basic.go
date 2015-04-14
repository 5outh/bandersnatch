package main

import (
    "bandersnatch"
    "fmt"
)

func main() {
    boardGames := bandersnatch.GetBoardGame(30000)

    fmt.Printf("%+v", boardGames)
}
