package lgtmgen

import (
	"errors"
	"math"

	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	maxSideLength float64 = 425
	font          string  = "assets/fonts/Archivo_Black/ArchivoBlack-Regular.ttf"
)

func Generate(src []byte) ([]byte, error) {
	imagick.Initialize()
	defer imagick.Terminate()

	srcmw := imagick.NewMagickWand()

	// 画像の読み込み
	if err := srcmw.ReadImageBlob(src); err != nil {
		return nil, err
	}

	// 横
	w := srcmw.GetImageWidth()
	// 縦
	h := srcmw.GetImageHeight()

	imgW, imgH := calcImageSize(float64(w), float64(h))
	titleFontSize, textFontSize := calcFontSize(imgW, imgH)

	titleDw := imagick.NewDrawingWand()
	defer titleDw.Destroy()
	textDw := imagick.NewDrawingWand()
	defer textDw.Destroy()

	if err := titleDw.SetFont(font); err != nil {
		return nil, err
	}
	if err := textDw.SetFont(font); err != nil {
		return nil, err
	}

	fillPw := imagick.NewPixelWand()
	if ok := fillPw.SetColor("#FFFFFF"); !ok {
		return nil, errors.New("invalid color")
	}
	strokePw := imagick.NewPixelWand()
	if ok := strokePw.SetColor("#000000"); !ok {
		return nil, errors.New("invalid color")
	}

	// 文字の塗りつぶしの色
	titleDw.SetFillColor(fillPw)
	textDw.SetFillColor(fillPw)
	// 文字の枠線の色
	titleDw.SetStrokeColor(strokePw)
	textDw.SetStrokeColor(strokePw)
	// 文字の太さ
	titleDw.SetStrokeWidth(1)
	textDw.SetStrokeWidth(0.8)
	// フォントサイズ
	titleDw.SetFontSize(titleFontSize)
	textDw.SetFontSize(textFontSize)
	// 中央寄せ
	titleDw.SetGravity(imagick.GRAVITY_CENTER)
	textDw.SetGravity(imagick.GRAVITY_CENTER)
	// 文字追加
	titleDw.Annotation(0, 0, "L G T M")
	textDw.Annotation(0, titleFontSize/1.5, "L o o k s   G o o d   T o   M e")

	cimw := srcmw.CoalesceImages()
	defer cimw.Destroy()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	_ = mw.SetImageDelay(srcmw.GetImageDelay())

	for i := 0; i < int(cimw.GetNumberImages()); i++ {
		if ok := cimw.SetIteratorIndex(i); !ok {
			return nil, errors.New("invalid index")
		}

		img := cimw.GetImage()
		defer img.Destroy()

		if err := img.AdaptiveResizeImage(uint(imgW), uint(imgH)); err != nil {
			return nil, err
		}
		if err := img.DrawImage(titleDw); err != nil {
			return nil, err
		}
		if err := img.DrawImage(textDw); err != nil {
			return nil, err
		}
		if err := mw.AddImage(img); err != nil {
			return nil, err
		}
	}

	return mw.GetImagesBlob(), nil
}

func calcImageSize(w, h float64) (float64, float64) {
	if w > h {
		return maxSideLength, maxSideLength * h / w
	} else {
		return maxSideLength * w / h, maxSideLength
	}
}

func calcFontSize(w, h float64) (float64, float64) {
	return math.Min(h/2, w/6), math.Min(h/9, w/27)
}
