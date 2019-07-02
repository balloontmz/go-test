package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	// "path/filepath"
	"sync"
)

var TILESDB map[string][3]float64 // 全局只读数据库

type DB struct {
	mutex *sync.Mutex // 互斥锁
	store map[string][3]float64
}

//cloneTilesDB 拷贝数据库副本，防止重复生成数据库耗时。初运行程序时需要初始化一次数据库
func cloneTilesDB() DB {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	tiles := DB{
		store: db,
		mutex: &sync.Mutex{},
	}
	return tiles
}

//计算图片的平均颜色
/*
averageColor 函数会将给定图片的每个像素中的红绿蓝三种颜色加起来，并将这些颜色的总和除以图片的像素数量，并将最后的除法的计算结果记录在一个新创建的三元组里面。
*/
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()  // 此函数的作用？
	r, g, b := 0.0, 0.0, 0.0  // 浮点数的原因可能是怕溢出？
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X;x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA() // 此函数的作用？
			r, g, b = r + float64(r1), g + float64(g1), b + float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y) // 总像素
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

//resize 将图片缩放至指定宽度? -- 按理说应该能
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx()/newWidth   // 设定缩放 比例
	// 此函数的用处？
	// 此处第二个参数到底是 x 还是 y
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.Y/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y + ratio, j + 1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x + ratio, i + 1 {
			r, g, b, a := in.At(x, y).RGBA()
			// 此处函数的作用是？
			out.SetNRGBA(i, j, color.NRGBA{uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8),})
		}
	}
	return *out
}

//tilesDB 通过扫描瓷砖图片所在的目录来创建一个瓷砖图片数据库
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles") // 读取 tiles 文件夹的文件
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		// 此处的错误处理为什么采用嵌套的。
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("解码瓷砖图片失败", err, name)
			}
		} else {
			fmt.Println("无法打开瓷砖图片", name, err)
		}
		file.Close()
	}
	fmt.Println("完成填充数据库初始化")
	return db
}

//nearest 比对目标图片和数据库的数据，找到和目标图片平均颜色最为接近的图片，并返回文件名，并删除数据库中这条记录
func (db *DB) nearest(target [3]float64) string {
	var filename string
	// 因为在数据库移除被选中的图片之前，还是存在多个 goroutine 将同张图片匹配为最佳结果的情况，所以需要对整个匹配过程加锁
	db.mutex.Lock()
	smallest := 1000000.0 // 浮点数
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename) // 删除最后一次的数据库中文件名键，即选定的文件。
	db.mutex.Unlock()
	return filename
}

//distance 计算两个点的欧几里得距离。三维点的距离
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p1[0] - p2[0]) + sq(p1[1] - p2[1]) + sq(p1[2] - p2[2]))
}

func sq(n float64) float64 {
	return n * n 
}

