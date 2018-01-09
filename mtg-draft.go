package main

import (
    "github.com/MagicTheGathering/mtg-sdk-go"
    "github.com/briandowns/spinner"
    "fmt"
    "net/http"
    "io"
    "os"
    "os/exec"
    "strconv"
    "time"
    "sync"
)

const NUM_PLAYERS = 2
const BOOSTERS_PER_PLAYER = 1
const SET_CODE = "KTK"

func downloadArt (c *mtg.Card, artFileName string, wg *sync.WaitGroup) (err error) {
    defer wg.Done()
    res, err := http.Get(c.ImageUrl)
    if err != nil { return err }
    defer res.Body.Close()
    file, err := os.Create(artFileName)
    if err != nil { return err }
    defer file.Close()
    _, err = io.Copy(file, res.Body)
    if err != nil { return err }
    return nil
}

func generateBooster (i int, wg *sync.WaitGroup) (err error) {
    defer wg.Done()
    command := exec.Command("montage")
    cards, err := mtg.SetCode(SET_CODE).GenerateBooster()
    if err != nil { return err }
    var subWg sync.WaitGroup
    subWg.Add(len(cards))
    for _, c := range cards {
        artFileName := "./cards/" + string(c.Id) + ".png"
        downloadArt(c, artFileName, &subWg)
        command.Args = append(command.Args, artFileName)
    }
    subWg.Wait()
    command.Args = append(command.Args, 
                         "-geometry", "744x1039", 
                         "-density", "300", 
                         "-tile", "3x3")
    outputFileName := "./boosters/booster" + strconv.Itoa(i) + ".pdf"
    command.Args = append(command.Args, outputFileName)
    err = command.Start()
    if err != nil { return err }
    return nil
}

func main () {
    numBoosters := NUM_PLAYERS * BOOSTERS_PER_PLAYER
    fmt.Printf("Generating %d %s booster packs. ", numBoosters, SET_CODE)
    spin := spinner.New(spinner.CharSets[14], 50*time.Millisecond)
    spin.Start()
    // Concurrently generate boosters.
    var wg sync.WaitGroup
    wg.Add(numBoosters)
    for i := 0; i < numBoosters; i++ {
        go generateBooster(i, &wg)
    }
    wg.Wait()
    spin.Stop()
}
