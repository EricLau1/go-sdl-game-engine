package collisions

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/maps"
)

type CollisionHandler struct {
	collisionTileMap *maps.TileMap
	collisionLayer   *maps.TileLayer
}

func NewCollisionHandler(gameMap *maps.GameMap) *CollisionHandler {
	layer := gameMap.Back()

	layer.Log()

	return &CollisionHandler{
		collisionTileMap: layer.Tilemap,
		collisionLayer:   layer,
	}
}

func (h *CollisionHandler) Check(a, b sdl.Rect) bool {
	xOverlaps := (a.X < b.X+b.W) && (a.X+a.W > b.X)
	yOverlaps := (a.Y < b.Y+b.H) && (a.Y+a.H > b.Y)
	return xOverlaps && yOverlaps
}

func (h *CollisionHandler) Map(a sdl.Rect) bool {

	rows := int32(h.collisionLayer.Rows)
	cols := int32(h.collisionLayer.Columns)
	tileSize := int32(h.collisionLayer.Size)

	//sdl.Log("TileSize=%d, Rows=%d, Cols=%d", tileSize, rows, cols)

	leftTile := a.X / tileSize
	rightTile := (a.X + a.W) / tileSize

	topTile := a.Y / tileSize
	bottomTile := (a.Y + a.H) / tileSize

	if leftTile < 0 {
		leftTile = 0
	}

	if rightTile > cols {
		rightTile = cols
	}

	if topTile < 0 {
		topTile = 0
	}
	if bottomTile > rows {
		bottomTile = rows
	}

	//sdl.Log("Map Collision: Player=%v, LeftTile=%v, RightTile=%v, TopTile=%d, BottomTile=%d", a,leftTile, rightTile, topTile, bottomTile)

	for column := leftTile; column <= rightTile; column++ {
		for row := topTile; row <= bottomTile; row++ {

			v := h.collisionTileMap.Get(int(row), int(column))
			//sdl.Log("TileMap: Row=%d, Column=%d, Value=%d", row, column, v)
			if v > 0 {
				return true
			}
		}
	}

	return false
}
