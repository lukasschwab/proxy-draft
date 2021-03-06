# proxy-draft

This is a tool to randomly generate realistic proxied booster packs for a Magic: the Gathering draft.

After too many years of looking down my nose at MTG, I picked it up as a hobby this fall––if you haven't played it, I'd highly recommend giving it a shot. The game (or at least the more interactive strategies) has super high replay values.

But some experiences in Magic are just too expensive for a casual player like myself. What if I want to play a sealed draft tournament with Alpha, which would cost tens of thousands of dollars in real life?

Instead of selling organs, I can print out ("proxy") cards for this kind of competitive play. This program uses the [Magic: the Gathering API](https://docs.magicthegathering.io/#api_v1booster_get) to generate realistic booster packs.

Run, print, draft, and play. Enjoy!

## Setup

You'll need `imagemagick` and `montage` installed. On a Mac, I installed both using Homebrew:

```
$ brew install imagemagick
$ brew install montage
```

To install the Go dependencies, run this from the project directory:

```
$ go get -d ./...
```

Make sure there exist subdirectories `boosters` and `cards`. `mtg-draft.go` stores card images in `./cards` and outputs booster pack PDFs to `./boosters`.

## Usage

### Cleanup

To remove downloaded card image files and generated booster pack PDFs, run:

```
$ bash cleanup.sh
```

### Draft

Right now draft parameters are hard-coded into `mtg-draft.go`. You may want to modify them to suit the draft you're running.

| Constant              | Meaning                                                                                    |
|:-----------------------|:--------------------------------------------------------------------------------------------|
| `NUM_PLAYERS`         | Number of players in the draft; used to calculate number of packs to generate. Default: 6. |
| `BOOSTERS_PER_PLAYER` | Number of booster packs to generate per player. Default: 3.                                |
| `SET_CODE`            | The abbreviated "code" designation of the set to draft; see code/draft mappings below.     |

Once you've selected the appropriate parameters, call the program from the command line.

```
$ go run mtg-draft.go
```

The resulting printable booster pack PDFs will be located in `/boosters`.

## Set codes

These are, roughly, the valid set codes for this program.

If there weren't booster packs issues for the set you select, the program will not generate booster packs. I've removed a number of these sets already––for example, the Duel Decks––but haven't done so exhaustively.

| Code | Set name |
|:----|:---------|
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
| 10E | Tenth Edition

## To do

+ Parse command line options.