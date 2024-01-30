package main

const MessageRange = 1000

type Shouty struct {
	locationsByShouter map[string]*Coordinate
	shoutsByShouter    map[string][]string
}

func NewShouty() *Shouty {
	return &Shouty{
		locationsByShouter: make(map[string]*Coordinate),
		shoutsByShouter:    make(map[string][]string),
	}
}

func (s *Shouty) SetLocation(person string, coordinate *Coordinate) {
	s.locationsByShouter[person] = coordinate
}

func (s *Shouty) Shout(person, shout string) {
	s.shoutsByShouter[person] = append(s.shoutsByShouter[person], shout)
}

func (s *Shouty) GetShoutsHeardBy(listener string) map[string][]string {
	shoutsHeard := make(map[string][]string)
	listenerLocation, exists := s.locationsByShouter[listener]

	if !exists {
		return shoutsHeard
	}

	for shouter, shouts := range s.shoutsByShouter {
		shouterLocation, exists := s.locationsByShouter[shouter]

		if !exists {
			continue
		}

		if shouterLocation.DistanceFrom(listenerLocation) < MessageRange {
			shoutsHeard[shouter] = shouts
		}
	}

	return shoutsHeard
}
