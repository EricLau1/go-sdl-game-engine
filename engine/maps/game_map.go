package maps

type GameMap struct {
	MapLayers []*TileLayer
}

func NewGameMap() *GameMap {
	return &GameMap{MapLayers: []*TileLayer{}}
}

func (gm *GameMap) Render() {
	for _, tileLayer := range gm.MapLayers {
		tileLayer.Render()
	}
}
