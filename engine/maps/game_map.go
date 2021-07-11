package maps

type GameMap struct {
	MapLayers []*TileLayer
}

func NewGameMap() *GameMap {
	return &GameMap{MapLayers: []*TileLayer{}}
}

func (gm *GameMap) Front() *TileLayer {
	return gm.MapLayers[0]
}

func (gm *GameMap) Back() *TileLayer {
	return gm.MapLayers[len(gm.MapLayers)-1]
}

func (gm *GameMap) Render() {
	for _, tileLayer := range gm.MapLayers {
		tileLayer.Render()
	}
}
