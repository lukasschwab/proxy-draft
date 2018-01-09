package main

import (
    "github.com/MagicTheGathering/mtg-sdk-go"
    "fmt"
)

func main () {
    sets, _ := mtg.NewSetQuery().All()
    for _, set := range sets {
        fmt.Println(set.SetCode, set.Name)
    }
}

// Is there a way to automatically print proxies?
// Can use https://github.com/jung-kurt/gofpdf to automatically generate PDFs to print... 
// as long as the cards being drafted have ImageUrl(). Not all cards do.

// "Sets not on Gatherer are: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter code 
// that starts with a lowercase ‘p’."

