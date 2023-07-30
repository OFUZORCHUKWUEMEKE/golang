package main

import "fmt"

const (
	Online      = 0
	Offlne      = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(server map[string]int) {
	fmt.Println("\nThere are", len(server), "servers")
	stats := make(map[int]int)

	for _, status := range server {
		switch status {
		case Online:
			stats[Online] += 1
		case Offlne:
			stats[Offlne] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offlne], "Servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintenance")
	fmt.Println(stats[Retired], "servers are retired")

}

func main() {
	server := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)

	for _, server := range server {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offlne

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}
	fmt.Println("Thank You ver Much")
	printServerStatus(serverStatus)
}
