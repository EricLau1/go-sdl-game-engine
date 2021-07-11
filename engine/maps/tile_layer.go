package maps

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/graphics"
)

type TileMap struct {
	tiles [][]*int
	rows  int
	cols  int
}

func NewTileMap(rows, cols, defaultValue int) *TileMap {
	tiles := make([][]*int, rows)

	for row := 0; row < rows; row++ {
		tiles[row] = make([]*int, 0, cols)
		for col := 0; col < cols; col++ {
			def := defaultValue
			tiles[row] = append(tiles[row], &def)
		}
	}

	return &TileMap{tiles: tiles, rows: rows, cols: cols}
}

func (tm *TileMap) Add(row []*int) {
	tm.tiles = append(tm.tiles, row)
}

func (tm *TileMap) Set(row, col, value int) {
	if tm.tiles[row] == nil {
		tm.tiles[row] = make([]*int, tm.cols)
	}
	tm.tiles[row][col] = &value
}

func (tm *TileMap) Get(i, j int) int {
	if i >= tm.rows || j >= tm.cols {
		return 0
	}
	return *tm.tiles[i][j]
}

func (tm *TileMap) Rows() int {
	return tm.rows
}

func (tm *TileMap) Columns() int {
	return tm.cols
}

func (tm *TileMap) Log() {
	for _, row := range tm.tiles {
		for _, col := range row {
			fmt.Printf("%d ", *col)
		}
		fmt.Println("")
	}
}

type TilesetList struct {
	tilesets []*Tileset
}

func NewTilesetList() *TilesetList {
	return &TilesetList{tilesets: []*Tileset{}}
}

func (tl *TilesetList) Add(tileset *Tileset) {
	tl.tilesets = append(tl.tilesets, tileset)
}

func (tl *TilesetList) Size() int {
	return len(tl.tilesets)
}

func (tl *TilesetList) List() []*Tileset {
	return tl.tilesets
}

func (tl *TilesetList) Get(i int) *Tileset {
	return tl.tilesets[i]
}

func (tl *TilesetList) Log() {
	for _, item := range tl.tilesets {
		sdl.Log("Tileset: %v", item)
	}
}

type TileLayer struct {
	Size           int
	Rows           int
	Columns        int
	Tilemap        *TileMap
	Tilesets       *TilesetList
	textureManager graphics.TextureManager
}

func NewTileLayer(size, rows, cols int, tileMap *TileMap, tilesets *TilesetList, textureManager graphics.TextureManager) *TileLayer {
	var tileLayer TileLayer

	tileLayer.Size = size
	tileLayer.Rows = rows
	tileLayer.Columns = cols
	tileLayer.Tilemap = tileMap
	tileLayer.Tilesets = tilesets
	tileLayer.textureManager = textureManager

	for _, tileset := range tilesets.List() {

		sdl.Log("load asset: Name = %s, Source=%s", tileset.Name, tileset.Image.Source)

		filePath := "assets/maps/" + tileset.Image.Source

		if !textureManager.Load(tileset.Name, filePath) {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "Failed load asset: %s", filePath)
		}

	}
	return &tileLayer
}

func (tl *TileLayer) Log() {
	sdl.Log("TileLayer: TileSize=%d, Rows=%d, Cols=%d, Tilesets=%d", tl.Size, tl.Rows, tl.Columns, tl.Tilesets.Size())
}

func (tl *TileLayer) Render() {
	for i := 0; i < tl.Rows; i++ {

		for j := 0; j < tl.Columns; j++ {

			tileID := tl.Tilemap.Get(i, j)

			if tileID == 0 {
				continue
			}

			var index int

			if tl.Tilesets.Size() > 1 {

				for k := 1; k < tl.Tilesets.Size(); k++ {

					tilesets := tl.Tilesets.List()

					if tileID >= tilesets[k].FirstID && tileID <= tilesets[k].LastID {
						tileID = tileID + tilesets[k].Count - tilesets[k].LastID
						index = k
						break
					}
				}
			}

			ts := tl.Tilesets.Get(index)
			tileRow := tileID / ts.Columns
			tileCol := tileID - tileRow*ts.Columns - 1

			if tileID%ts.Columns == 0 {
				tileRow--
				tileCol = ts.Columns - 1
			}

			tl.textureManager.DrawTile(ts.Name, int32(ts.Width), int32(j*ts.Width), int32(i*ts.Width), int32(tileRow), int32(tileCol), sdl.FLIP_NONE)
		}

	}
}
