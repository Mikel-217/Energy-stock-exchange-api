# About
This is a API which tells the user when to buy energy. This can be all kinds of energy, like gas or electrycity

## Endpoints:



## APIs in use:
- [Energy-Chart](https://api.energy-charts.info/) API
- [Bundesnetzargentur (SMARD)](https://smard.api.bund.dev/) API
- TODO: search for more

# Planing

## Important stuff
- Add a builder / factory pattern which builds the two or more API go-routines at the start
- Add a startup func which creates all kind of stuff (DB tables, API connections)
- Define API Endpoints -> for ui
- Add when to buy recommandation

## Future stuff
- Add basic user authentication
- Add a server console to add more api's / change some stuff
- Add a realtime chart to see what the current price is
- Add Android App
- Add Web-UI
