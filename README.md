# Alien Invasion simulation

Simple Golang program simulates Alien invasion  

## How to get
```sh
go get github.com/sysdevguru/alien-invasion
```

## How to run
```sh
make build
./alien-invasion -aliens=[number of aliens] -map=[path to map]
```
Default value of Aliens is 10 and default Map file is "maps/map.txt"  

## Test
```sh
make test
```

## Description
Alien Invasion.pdf  

## Assumptions
Technical assumptions  

### Movement assumption
- All aliens movements will be counted from 1 to 10,000  

### Map file assumption
- Map file contains all information about all cities  
For example, 5 lines for 5 cities  

### City removement assumption
- City will be destroyed if two Alien fights in the city  
- Destroyed city will be removed from the map  

### Aliens death assumption
- Aliens will not fight on the road  
- Aliens will not fight at the first start  
- Aliens will fight if two or more of them meet in same city and all of them will die  
