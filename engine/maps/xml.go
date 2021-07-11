package maps

import "encoding/xml"

// <map version="1.2" tiledversion="1.3.2" orientation="orthogonal" renderorder="right-down" compressionlevel="0" width="60" height="20" tilewidth="32" tileheight="32" infinite="0" nextlayerid="3" nextobjectid="1">
type Map struct {
	XMLName          xml.Name  `xml:"map"`
	Version          string    `xml:"version,attr"`
	TiledVersion     string    `xml:"tiledversion,attr"`
	Orientation      string    `xml:"orientation,attr"`
	RenderOrder      string    `xml:"renderorder,attr"`
	CompressionLevel int       `xml:"compressionlevel,attr"`
	Width            int       `xml:"width,attr"`
	Height           int       `xml:"height,attr"`
	TileWidth        int       `xml:"tilewidth,attr"`
	TileHeight       int       `xml:"tileheight,attr"`
	Infinite         int       `xml:"infinite,attr"`
	NextLayerID      int       `xml:"nextlayerid,attr"`
	NextObjectID     int       `xml:"nextobjectid,attr"`
	Tilesets         []Tileset `xml:"tileset"`
	Layers           []Layer   `xml:"layer"`
}

// <tileset firstgid="421" name="objects" tilewidth="32" tileheight="32" tilecount="91" columns="13">
type Tileset struct {
	XMLName xml.Name `xml:"tileset"`
	FirstID int      `xml:"firstgid,attr"`
	Name    string   `xml:"name,attr"`
	Width   int      `xml:"tilewidth,attr"`
	Height  int      `xml:"tileheight,attr"`
	Count   int      `xml:"tilecount,attr"`
	Columns int      `xml:"columns,attr"`
	Image   Image    `xml:"image"`
	LastID  int
	Rows    int
}

// <image source="objects.PNG" width="416" height="224"/>
type Image struct {
	XMLName xml.Name `xml:"image"`
	Source  string   `xml:"source,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

// <layer id="2" name="B2" width="60" height="20">
type Layer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Data    Data     `xml:"data"`
}

type Data struct {
	XMLName  xml.Name `xml:"data"`
	Encoding string   `xml:"encoding,attr"`
	Value    string   `xml:",chardata"`
}
