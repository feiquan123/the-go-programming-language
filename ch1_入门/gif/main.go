// lissajous 产生随机利萨如图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.Gray16{0x006699}, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = iota // 第一中颜色: 白色
	blackIndex        // 下一种颜色: 黑色
	cyanIndex         // 青蓝色
	greenIndex        //  绿色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Println("Listen:", "http://localhost:8000")
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5        // 完整的x 振荡变化的个数
		res     = 0.001    // 角度分辨率
		w       = 1920 / 2 // 画布宽
		h       = 1080 / 2 // 画布高
		nframes = 64       // 动画的帧数
		delay   = 8        // 以10ms 为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*w+1, 2*h+1)
		img := image.NewPaletted(rect, palette)
		for i := range img.Pix {
			img.Pix[i] = blackIndex
		}
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(w+int(x*w+0.5), h+int(y*h+0.5), greenIndex)
		}
		phase += 0.1

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: 忽略编码错误
}
