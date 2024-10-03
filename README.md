# PLB
Created for testing service http dependency. 

## Features
* Custom port(default: 8080)
* Customizable routes via json file. Check test.json for reference
* Request latency

## How to Run
### Build
```
make build
```
### Run
```
PORT=3000 ./out/plb test.json
```
