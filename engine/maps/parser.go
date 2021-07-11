package maps

import (
	"encoding/xml"
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/graphics"
	"io"
	"os"
	"strconv"
	"strings"
)

type MapParser struct {
	gameMaps map[string]*GameMap
}

func NewMapParser(id string, source string, textureManager graphics.TextureManager) *MapParser {
	var mapParser MapParser
	mapParser.gameMaps = make(map[string]*GameMap)

	sdl.Log("loading map...")

	m, err := Load(source)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return nil
	}

	rows, cols, tileSize := m.Height, m.Width, m.TileWidth

	tilesetList := NewTilesetList()

	for index, _ := range m.Tilesets {
		ts := ParseTileset(&m.Tilesets[index])
		tilesetList.Add(ts)
	}

	tilesetList.Log()

	gameMap := NewGameMap()

	for _, layer := range m.Layers {
		tileLayer := ParseTileLayer(&layer, tilesetList, tileSize, rows, cols, textureManager)
		gameMap.MapLayers = append(gameMap.MapLayers, tileLayer)
	}

	mapParser.gameMaps[id] = gameMap

	return &mapParser
}

func (mp *MapParser) GetMap(id string) *GameMap {
	return mp.gameMaps[id]
}

func (mp *MapParser) Clean() {
	for key := range mp.gameMaps {
		mp.gameMaps[key] = nil
		delete(mp.gameMaps, key)
	}
}

func ParseTileset(tileset *Tileset) *Tileset {
	tileset.LastID = (tileset.FirstID + tileset.Count) - 1
	tileset.Rows = tileset.Count / tileset.Columns
	return tileset
}

func ParseTileLayer(layer *Layer, tilesets *TilesetList, tileSize, rowCount, colCount int, textureManager graphics.TextureManager) *TileLayer {

	tileMap := NewTileMap(rowCount, colCount, 0)

	data := ParseData(layer)

	for i, row := range data {
		for j, value := range row {
			tileMap.Set(i, j, value)
		}
	}

	tileMap.Log()

	return NewTileLayer(tileSize, rowCount, colCount, tileMap, tilesets, textureManager)
}

func ParseData(layer *Layer) [][]int {
	data := make([][]int, layer.Height)

	rows := strings.Split(strings.TrimSpace(layer.Data.Value), "\n")

	for index, row := range rows {
		data[index] = make([]int, 0, layer.Width)

		for _, value := range strings.Split(strings.TrimSpace(row), ",") {
			v, err := strconv.Atoi(value)
			if err == nil {
				data[index] = append(data[index], v)
			}
		}

	}

	return data
}

func Load(source string) (*Map, error) {
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var m Map

	err = xml.Unmarshal(b, &m)

	return &m, err
}
