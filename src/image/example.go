package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

// image实现了基本的2D图片库
// 基本接口叫作Image。图片的色彩定义在image/color包。
// Image接口可以通过调用如NewRGBA和NewPaletted函数等获得；也可以通过调用Decode函数解码包含GIF、JPEG或PNG格式图像数据的输入流获得
// 解码任何具体图像类型之前都必须注册对应类型的解码函数
// 注册过程一般是作为包初始化的副作用，放在包的init函数里。因此，要解码PNG图像，只需在程序的main包里嵌入  import _ "image/png"
func main() {

	// png图片生成
	examplePngEncode()
	// png图片解析
	examplePngDecode()

	exampleJpegEncode()
}

const (
	PngFilePath = "testdata/rgb.png"
	JpegFilePath = "testdata/rgb.jpeg"
	GifFilePath = "testdata/rgb.gif"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy = c.X-x, c.Y-y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func createRGBAImage() *image.RGBA {

	// 指定长宽
	var w, h = 280, 240

	var hw, hh = float64(w/2), float64(h/2)
	r := 40.0

	// 三分之二的圆周率
	p := 2 * math.Pi / 3
	// 红色圆
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	// 绿色圆
	cg := &Circle{hw - r*math.Sin(p), hh - r*math.Cos(p), 60}
	// 蓝色圆
	cb := &Circle{hw - r*math.Sin(-p), hh - r*math.Cos(-p), 60}


	// 指定两个坐标点生成矩形
	// Rect是Rectangle {Pt（x0，y0），Pt（x1，y1）}的简写
	// 返回的矩形具有必要时交换的最小和最大坐标，以便格式良好
	rect := image.Rect(0, 0, w, h)

	// 初始化一张NRGBA图片
	m := image.NewRGBA(rect)

	// 为图片上每个点设置颜色
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				R: cr.Brightness(float64(x), float64(y)),
				G: cg.Brightness(float64(x), float64(y)),
				B: cb.Brightness(float64(x), float64(y)),
				A: 255,
			}
			m.Set(x, y, c)
		}
	}

	return m
}

func examplePngEncode() {

	// 生成RGBA图片内容
	m := createRGBAImage()

	// 创建并打开一个只能写入的png文件
	f, err := os.OpenFile(PngFilePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 将RGBA内容编码成png写入文件
	if err := png.Encode(f, m); err != nil {
		log.Fatal(err)
	}
}

func examplePngDecode() {

	// 打开png文件
	f, err := os.Open(PngFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// png解码
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	// Bounds返回图片域，边界不一定包含点（0，0）
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {

			// at返回像素在（x，y）处的颜色
			// at（bounds（）.min.x，bounds（）.min.y）返回网格的左上角像素
			// at（bounds（）.max.x-1，bounds（）.max.y-1）返回右下角的像素
			pointColor := img.At(x, y)

			// 使用灰度色覆盖
			c := color.GrayModel.Convert(pointColor).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

func exampleJpegEncode() {

	// 生成RGBA图片内容
	m := createRGBAImage()

	// 创建并打开一个只能写入的jpeg文件
	f, err := os.OpenFile(JpegFilePath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 将RGBA内容编码成jpeg写入文件
	// Options图像质量值为100，是最好的图像显示
	if err := jpeg.Encode(f, m, &jpeg.Options{100}); err != nil {
		log.Fatal(err)
	}
}