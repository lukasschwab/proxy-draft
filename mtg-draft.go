package main

import (
    "github.com/MagicTheGathering/mtg-sdk-go"
    "fmt"
    "net/http"
    "io"
    "os"
    "os/exec"
    "strconv"
)

const NUM_PLAYERS = 6
const BOOSTERS_PER_PLAYER = 3
const SET_CODE = "KTK"

// func downloadCardArtwork

// func generateBooster

func main () {
    for i := 0; i < NUM_PLAYERS * BOOSTERS_PER_PLAYER; i++ {
        command := exec.Command("montage")
        fmt.Println("BOOSTER", i)
        cards, err := mtg.SetCode(SET_CODE).GenerateBooster()
        if err != nil {
            fmt.Println(err)
        }
        for _, c := range cards {
            fmt.Println(c.Rarity, c.Name)
            // Download the image
            res, err := http.Get(c.ImageUrl)
            if err != nil {
                fmt.Println(err)
                return
            }
            defer res.Body.Close()
            file, err := os.Create("./cards/" + string(c.Id) + ".png")
            if err != nil {
                fmt.Println(err)
            }
            defer file.Close()
            _, err = io.Copy(file, res.Body)
            if err != nil {
                fmt.Println(err)
                return
            }
            // Add the image to the command string
            // command += "\"./cards/" + string(c.Id) + ".png\" "
            command.Args = append(command.Args, "./cards/" + string(c.Id) + ".png")
        }
        command.Args = append(command.Args, "-geometry", "744x1039", "-density", "300", "-tile", "3x3")
        command.Args = append(command.Args, "./boosters/booster" + strconv.Itoa(i) + ".pdf")
        err = command.Start()
    }
}

// Is there a way to automatically print proxies?
// Can use https://github.com/jung-kurt/gofpdf to automatically generate PDFs to print... 
// as long as the cards being drafted have ImageUrl(). Not all cards do.

// "Sets not on Gatherer are: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter code 
// that starts with a lowercase ‘p’."

