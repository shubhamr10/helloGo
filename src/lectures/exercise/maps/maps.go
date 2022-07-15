//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayserver(servers_2 map[string]int) {
	fmt.Println("__", " MAPS SERVER ", "__")
	totalServers := 0
	var online, offline, maintainenece, retired int = 0, 0, 0, 0
	for _, value := range servers_2 {
		totalServers += totalServers
		switch value {
		case Online:
			online += 1
		case Offline:
			offline += 1
		case Maintenance:
			maintainenece += 1
		case Retired:
			retired += 1
		}
	}

	fmt.Println("Total servers: ", totalServers)
	fmt.Println("Online Servers:", online)
	fmt.Println("Offline servers:", offline)
	fmt.Println("Maintainence servers:", maintainenece)
	fmt.Println("Retired servers:", retired)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	myServer := make(map[string]int)
	// Set all of the server statuses to `Online` when creating the map
	for _, server := range servers {
		myServer[server] = Online
	}
	fmt.Println("Server created with online status", myServer)

	//  - call display server info function
	displayserver(myServer)
	//  - change server status of `darkstar` to `Retired`
	myServer["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	myServer["aiur"] = Offline
	//  - call display server info function
	displayserver(myServer)
	//  - change server status of all servers to `Maintenance`
	for key, _ := range myServer {
		myServer[key] = Maintenance
	}
	//  - call display server info function
	displayserver(myServer)
}
