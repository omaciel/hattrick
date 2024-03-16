# Hat Trick
Sample CLI application written in Go to access the schedule for NHL games using the open NHL API as documented [here](https://gitlab.com/dword4/nhlapi/-/tree/master).

## Build the CLI application
The following command will generate a new `hattrick` executable in your current directory:

  go build .

## Running the CLI application
Get the current week's schedule:

  ./hattrick
  
    Schedule for NHL games on  2024-03-16
    BUF: 0 vs DET: 0
    OTT: 0 vs NYI: 0
    NYR: 0 vs PIT: 0
    NJD: 0 vs ARI: 0
    TBL: 0 vs FLA: 0
    PHI: 0 vs BOS: 0
    SJS: 0 vs CBJ: 0
    MTL: 0 vs CGY: 0
    CAR: 0 vs TOR: 0
    LAK: 0 vs DAL: 0
    MIN: 0 vs STL: 0
    COL: 0 vs EDM: 0
    NSH: 0 vs SEA: 0
    WSH: 0 vs VAN: 0

Get last week's schedule:

  ./hattrick  -date 2024-03-09

    Schedule for NHL games on  2024-03-09
    EDM: 2 vs BUF: 3
    NSH: 2 vs CBJ: 1
    CAR: 4 vs NJD: 2
    PIT: 1 vs BOS: 5
    CGY: 1 vs FLA: 5
    TOR: 3 vs MTL: 2
    OTT: 1 vs SJS: 2
    PHI: 0 vs TBL: 7
    CHI: 1 vs WSH: 4
    STL: 0 vs NYR: 4
    WPG: 0 vs VAN: 5
    DET: 3 vs VGK: 5
    DAL: 4 vs LAK: 1

Get last week's schedule for the New Jersey Devils:

  ./hattrick -date 2024-03-09 -team NJD
  
    Schedule for  NJD  on  2024-03-09
    2024-03-09 - CAR: 4 vs NJD: 2
    2024-03-11 - NJD: 1 vs NYR: 3
    2024-03-14 - NJD: 6 vs DAL: 2
