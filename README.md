# About
This is a API which tells the user when to buy energy. This can be all kinds of energy, like gas or electrycity.
The project uses a generic api-client builder, which builds all clients given from the config. If you want to add your own api see below.
Also there is an generic database-read-client builder, which can build a database client for getting data.

## Endpoints:

### Recommendation endpoints
- GET /recommendation -> gets recommendations for the current date
- GET /recommendation?date=2026-04-02 -> gets recommendations for the given date
- GET /recommendation?start=2026-04-02&end=2026-06-02 -> gets recommendations for a timespan
- GET /recommendation?id=1 -> gets recommendations for a given id


### Price endpoints
- GET /price?all=2026-04-02 -> gets all prices for a date
- GET /price?start=2026-04-02&end=2026-04-02 -> gets all prices for a timespan
- GET /price?id=1 -> gets a price for the given id


## APIs in use:
- [Energy-Chart](https://api.energy-charts.info/) API
- [Bundesnetzargentur (SMARD)](https://smard.api.bund.dev/) API
- TODO: search for more

# Planing

## Important stuff
- Add when to buy recommandation -> not finished yet
- Add a server console to change some stuff like intervals or urls
- Add a realtime chart to see what the current price is
- Add Web-UI

## Future stuff
- Add basic user authentication -> maybe
- Add Android App
