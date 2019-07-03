package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main()  {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))  // 静态文件代理
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}

	TILESDB = tilesDB()  // 初始化数据库并赋值给全局变量
	fmt.Println("马赛克服务启动。。。")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request)  {
	t0 := time.Now()
	r.ParseMultipartForm(10485760) // 表单的参数 
	file, _, _ := r.FormFile("image")  // 获取用户上传的目标图片
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size")) // 获取传入的瓷砖图片大小

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	fmt.Println("此处打个标签0")

	// newImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	
	db := cloneTilesDB()  // 复制瓷砖数据库

	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)

	c := combine(bounds, c1, c2, c3, c4) // 内部 goroutine 对 c 放入值

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil) // 将原始图片编码成 jpeg 格式
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	// t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic": <- c,  // 阻塞直到取到对应的值
		// "duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}
	t1 := time.Now()
	images["duration"] = fmt.Sprintf("%v ", t1.Sub(t0))
	// fmt.Println("返回结果打印", images)
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}

func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <- chan image.Image {

	c := make(chan image.Image)
	// 此处会运行一个并发函数，需要注意如果存在竞争资源，则需要
	sp := image.Point{0, 0} // 为每张瓷砖图片设置起始点

	go func() {
		newImage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))

		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA() // 获取当前遍历位置的原始图片的左上角像素
				color := [3]float64{float64(r), float64(g), float64(b)}
	
				nearest := db.nearest(color)  // 获取最近的瓷砖图片名
				file, err := os.Open(nearest)
	
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := resize(img, tileSize) // 对瓷砖图片重新指定大小
						tile := t.SubImage(t.Bounds()) // 此函数的作用 
						tileBounds := image.Rect(x, y, x + tileSize, y + tileSize)
						draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error:", err, nearest)
					}
				} else {
					fmt.Println("error:", err, nearest)
				}
				file.Close()
			}
		}
		c <- newImage.SubImage(newImage.Rect) // 此函数的意义？
	}()
	return c
}

func combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	c := make(chan string)

	go func() {
		var wg sync.WaitGroup
		img := image.NewRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done()  // 每复制完一张子图片，就对计数器执行一次减一操作
		}
		wg.Add(4) // 将等待组计数器设置为 4
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool

		for {
			select {
			case s1, ok1 = <-c1:
				go copy(img, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(img, s2.Bounds(), s2, image.Point{r.Max.X/2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(img, s3.Bounds(), s3, image.Point{r.Min.X, r.Max.Y/2})
			case s4, ok4 = <-c4:
				go copy(img, s4.Bounds(), s4, image.Point{r.Max.X/2, r.Max.Y/2})
			}
			if ok1 && ok2 && ok3 && ok4 {
				break // 当所有通道都被读取了值--即都已关闭，跳出循环
			}
		}
		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())

	}()
	return c
}