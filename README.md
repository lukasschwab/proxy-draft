# proxy-draft

This is a tool to 

## Setup

You'll need `imagemagick` and `montage` installed. On a Mac, I installed both using Homebrew:

```
$ brew install imagemagick
$ brew install montage
```

To install the go dependencies, run this from the project directory:

```
$ go get -d ./...
```

Make sure there exist subdirectories `boosters` and `cards`. `mtg.go` stores card images in `./cards` and outputs booster pack PDFs to `./boosters`.

## Usage

### Cleanup

To remove downloaded card image files and generated booster pack PDFs, run:

```
$ bash cleanup.sh
```

### Draft

Right now draft parameters are hard-coded into `mtg.go`.

## Set codes

| Code | Set name |
|-----|----------|
| UST | Unstable |
| UNH | Unhinged |
| UGL | Unglued |
| PD3 | Premium Deck Series: Graveborn |
| PD2 | Premium Deck Series: Fire and Lightning |
| H09 | Premium Deck Series: Slivers |
| PTK | Portal Three Kingdoms |
| POR | Portal |
| PO2 | Portal Second Age |
| PCA | Planechase Anthology |
| PC2 | Planechase 2012 Edition |
| HOP | Planechase |
| VMA | Vintage Masters |
| MMA | Modern Masters |
| MM3 | Modern Masters 2017 Edition |
| MM2 | Modern Masters 2015 Edition |
| MED | Masters Edition |
| ME4 | Masters Edition IV |
| ME3 | Masters Edition III |
| ME2 | Masters Edition II |
| IMA | Iconic Masters |
| EMA | Eternal Masters |
| MPS_AKH | Masterpiece Series: Amonkhet Invocations |
| MPS | Masterpiece Series: Kaladesh Inventions |
| EXP | Zendikar Expeditions |
| E02 | Explorers of Ixalan |
| V16 | From the Vault: Lore |
| V15 | From the Vault: Angels |
| V14 | From the Vault: Annihilation (2014) |
| V13 | From the Vault: Twenty |
| V12 | From the Vault: Realms |
| V11 | From the Vault: Legends |
| V10 | From the Vault: Relics |
| V09 | From the Vault: Exiled |
| DRB | From the Vault: Dragons |
| EVG | Duel Decks: Elves vs. Goblins |
| DDT | Duel Decks: Merfolk vs. Goblins |
| DDS | Duel Decks: Mind vs. Might |
| DDR | Duel Decks: Nissa vs. Ob Nixilis |
| DDQ | Duel Decks: Blessed vs. Cursed |
| DDP | Duel Decks: Zendikar vs. Eldrazi |
| DDO | Duel Decks: Elspeth vs. Kiora |
| DDN | Duel Decks: Speed vs. Cunning |
| DDM | Duel Decks: Jace vs. Vraska |
| DDL | Duel Decks: Heroes vs. Monsters |
| DDK | Duel Decks: Sorin vs. Tibalt |
| DDJ | Duel Decks: Izzet vs. Golgari |
| DDI | Duel Decks: Venser vs. Koth |
| DDH | Duel Decks: Ajani vs. Nicol Bolas |
| DDG | Duel Decks: Knights vs. Dragons |
| DDF | Duel Decks: Elspeth vs. Tezzeret |
| DDE | Duel Decks: Phyrexia vs. the Coalition |
| DDD | Duel Decks: Garruk vs. Liliana |
| DDC | Duel Decks: Divine vs. Demonic |
| DD3_JVC | Duel Decks Anthology, Jace vs. Chandra |
| DD3_GVL | Duel Decks Anthology, Garruk vs. Liliana |
| DD3_EVG | Duel Decks Anthology, Elves vs. Goblins |
| DD3_DVD | Duel Decks Anthology, Divine vs. Demonic |
| DD2 | Duel Decks: Jace vs. Chandra |
| CNS | Magic: The Gathering—Conspiracy |
| CN2 | Conspiracy: Take the Crown |
| CMD | Magic: The Gathering-Commander |
| CMA | Commander Anthology |
| CM1 | Commander's Arsenal |
| C17 | Commander 2017 |
| C16 | Commander 2016 |
| C15 | Commander 2015 |
| C14 | Commander 2014 |
| C13 | Commander 2013 Edition |
| CEI | International Collector's Edition |
| CED | Collector's Edition |
| E01 | Archenemy: Nicol Bolas |
| ARC | Archenemy |
| ZEN | Zendikar |
| XLN | Ixalan |
| WWK | Worldwake |
| WTH | Weatherlight |
| W17 | Welcome Deck 2017 |
| W16 | Welcome Deck 2016 |
| VIS | Visions |
| VAN | Vanguard |
| USG | Urza's Saga |
| ULG | Urza's Legacy |
| UDS | Urza's Destiny |
| TSP | Time Spiral |
| TSB | Time Spiral "Timeshifted" |
| TPR | Tempest Remastered |
| TOR | Torment |
| TMP | Tempest |
| THS | Theros |
| STH | Stronghold |
| SOM | Scars of Mirrodin |
| SOK | Saviors of Kamigawa |
| SOI | Shadows over Innistrad |
| SHM | Shadowmoor |
| SCG | Scourge |
| BTD | Beatdown Box Set |
| S99 | Starter 1999 |
| S00 | Starter 2000 |
| RTR | Return to Ravnica |
| ROE | Rise of the Eldrazi |
| RAV | Ravnica: City of Guilds |
| PLS | Planeshift |
| PLC | Planar Chaos |
| PCY | Prophecy |
| ORI | Magic Origins |
| ONS | Onslaught |
| OGW | Oath of the Gatewatch |
| ODY | Odyssey |
| NPH | New Phyrexia |
| NMS | Nemesis |
| MRD | Mirrodin |
| MOR | Morningtide |
| MMQ | Mercadian Masques |
| MIR | Mirage |
| MGB | Multiverse Gift Box |
| MD1 | Modern Event Deck 2014 |
| MBS | Mirrodin Besieged |
| M15 | Magic 2015 Core Set |
| M14 | Magic 2014 Core Set |
| M13 | Magic 2013 |
| M12 | Magic 2012 |
| M11 | Magic 2011 |
| M10 | Magic 2010 |
| LRW | Lorwyn |
| LGN | Legions |
| LEG | Legends |
| LEB | Limited Edition Beta |
| LEA | Limited Edition Alpha |
| KTK | Khans of Tarkir |
| KLD | Kaladesh |
| JUD | Judgment |
| JOU | Journey into Nyx |
| ISD | Innistrad |
| INV | Invasion |
| ICE | Ice Age |
| HOU | Hour of Devastation |
| HML | Homelands |
| GTC | Gatecrash |
| GPT | Guildpact |
| FUT | Future Sight |
| FRF_UGIN | Ugin's Fate promos |
| FRF | Fate Reforged |
| FEM | Fallen Empires |
| EXO | Exodus |
| EVE | Eventide |
| EMN | Eldritch Moon |
| BRB | Battle Royale Box Set |
| DTK | Dragons of Tarkir |
| DST | Darksteel |
| DRK | The Dark |
| DKA | Dark Ascension |
| DIS | Dissension |
| DGM | Dragon's Maze |
| CST | Coldsnap Theme Decks |
| CSP | Coldsnap |
| CP3 | Magic Origins Clash Pack |
| CP2 | Fate Reforged Clash Pack |
| CP1 | Magic 2015 Clash Pack |
| CON | Conflux |
| CHR | Chronicles |
| CHK | Champions of Kamigawa |
| BOK | Betrayers of Kamigawa |
| BNG | Born of the Gods |
| BFZ | Battle for Zendikar |
| AVR | Avacyn Restored |
| ATQ | Antiquities |
| ARN | Arabian Nights |
| ARB | Alara Reborn |
| APC | Apocalypse |
| ALL | Alliances |
| ALA | Shards of Alara |
| AKH | Amonkhet |
| AER | Aether Revolt |
| 9ED | Ninth Edition |
| 8ED | Eighth Edition |
| 7ED | Seventh Edition |
| 6ED | Classic Sixth Edition |
| 5ED | Fifth Edition |
| 5DN | Fifth Dawn |
| 4ED | Fourth Edition |
| 3ED | Revised Edition |
| 2ED | Unlimited Edition |
10E | Tenth Edition

## To do

