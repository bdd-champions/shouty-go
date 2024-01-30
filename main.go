package main

func main() {
	// Example usage of Shouty
	shouty := NewShouty()
	shouty.SetLocation("Alice", &Coordinate{X: 800, Y: 0})
	shouty.SetLocation("Cap", &Coordinate{X: 600, Y: 0})
	shouty.SetLocation("Bob", &Coordinate{X: 0, Y: 0})
	shouty.Shout("Alice", "Hello, world!")
	shouty.Shout("Cap", "Hello, world!")
	shoutsHeardByBob := shouty.GetShoutsHeardBy("Bob")

	// Print shouts heard by Bob
	for shouter, shouts := range shoutsHeardByBob {
		for _, shout := range shouts {
			println("Bob heard", shouter, "shout:", shout)
		}
	}
}
