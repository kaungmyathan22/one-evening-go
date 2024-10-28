package main

func Direction(d string) string {
	switch d {
	case "N":
		return "North"
	case "E":
		return "East"
	case "W":
		return "West"
	case "S":
		return "south"
	default:
		return "unknown"
	}
}
