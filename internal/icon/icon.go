package icon

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"io"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// SVGToPNG renders SVG bytes to PNG at the given size (width and height).
// Uses white for currentColor so the icon works as a template on light/dark backgrounds.
func SVGToPNG(svgBytes []byte, size int) ([]byte, error) {
	icon, err := oksvg.ReadReplacingCurrentColor(bytes.NewReader(svgBytes), "#ffffff", oksvg.WarnErrorMode)
	if err != nil {
		return nil, err
	}
	icon.SetTarget(0, 0, float64(size), float64(size))

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	scanner := rasterx.NewScannerGV(size, size, img, img.Bounds())
	raster := rasterx.NewDasher(size, size, scanner)
	icon.Draw(raster, 1.0)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// SVGToPNGFromReader renders SVG from r to PNG at the given size.
func SVGToPNGFromReader(r io.Reader, size int) ([]byte, error) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	return SVGToPNG(buf.Bytes(), size)
}

// TransparentIcon returns ICO-format bytes for a fully transparent 16x16 icon.
// Used when only speed text should be shown in the tray (no visible app icon).
func TransparentIcon() ([]byte, error) {
	const size = 16
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	// RGBA zero value is transparent
	var pngBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		return nil, err
	}
	pngBytes := pngBuf.Bytes()
	// Wrap in ICO container (Vista+ PNG-in-ICO)
	ico := bytes.NewBuffer(nil)
	ico.Write([]byte{0, 0, 1, 0, 1, 0}) // ICONDIR
	// ICONDIRENTRY: width(1), height(1), colors(1), reserved(1), planes(2), bpp(2), size(4), offset(4)
	entry := [16]byte{}
	entry[0] = byte(size)
	entry[1] = byte(size)
	entry[2] = 0
	entry[3] = 0
	binary.LittleEndian.PutUint16(entry[4:6], 1)
	binary.LittleEndian.PutUint16(entry[6:8], 32)
	binary.LittleEndian.PutUint32(entry[8:12], uint32(len(pngBytes)))
	binary.LittleEndian.PutUint32(entry[12:16], 22) // offset to image data
	ico.Write(entry[:])
	ico.Write(pngBytes)
	return ico.Bytes(), nil
}
