package main

import (
    "github.com/5outh/bandersnatch"
    "fmt"
)

func main() {
    boardGames := bandersnatch.GetBoardGame(30000)

    fmt.Printf("%+v", boardGames)
}
