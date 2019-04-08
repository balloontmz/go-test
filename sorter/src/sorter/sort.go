package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var inputfile = flag.String("i", "input.bat", "File contains values for sorting")
var outputfile = flag.String("o", "output.bat", "File to receive sorted values")
var algorithm = flag.String("a", "algorithm", "Sort algorithm")

// 读取文件中的值到数组中
func readValues(inputfile string) (values []int, err error) {
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("Fail to open the input file", inputfile)
		return
	}

	// 保证文件关闭
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0) // 初始化结果数组

	// 死循环，读取完全部数据
	for {
		line, isPrefix, err1 := br.ReadLine()

		// 如果到了文件尾部或者发生错误，则返回。
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		// 如果没有读取到全文件，则打印并返回
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		// 字节转换为字符串
		str := string(line)

		// 字符串转数字
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outputfile string) error {
	file, err := os.Create(outputfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outputfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value) // 数字转字符串
		file.WriteString(str + "\n")
		fmt.Println("当前遍历的结果为: ", value)
	}
	return nil
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 打印命令行参数
	if inputfile != nil {
		fmt.Println("inputfile = ", *inputfile, "outputfile = ", *outputfile, "algorithm = ", *algorithm)
	}

	values, err := readValues(*inputfile)
	// 没有错误的情况下
	if err == nil {
		fmt.Println("Read values:", values)
		writeValues(values, *outputfile)
	} else {
		fmt.Println(err)
	}
}
