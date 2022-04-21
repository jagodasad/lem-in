package structs

// Farm represents an adjaceny list graph
type Farm struct {
	Ants  []*Ant
	Rooms []*Room
}

// Ant
type Ant struct {
	Id          int
	Path        []*Room
	CurrentRoom *Room
	RoomsPassed int
}

// Room in farm (Vertex in the graph)
type Room struct {
	Name    string
	Ants    int
	X_pos   int
	Y_pos   int
	IsStart bool
	IsEnd   bool
	Links []*Room
}

// Data to generate future farm
type GenerationData struct {
	NumberOfAnts int
	Rooms        []string
	Links        []string
	StartIndex   int
	EndIndex     int
}

// Structurised paths by links in start room
type PathStruct struct {
	Path []*Room
}