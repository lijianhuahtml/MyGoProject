package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func f1() {
	inputFile, inputError := os.Open("io/input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {

		}
	}(inputFile)

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

// 按列读取文件中的数据
func f3() {
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

// 每行的第一个字段为标题，第二个字段为价格，第三个字段为数量。
// 内容的格式基本与 示例 12.3c 的相同，除了分隔符改成了分号。
// 请读取出文件的内容，创建一个结构用于存取一行的数据，然后使用结构的切片，并把数据打印出来。

type Book struct {
	title    string
	price    float64
	quantity int
}

func f4() {
	bks := make([]Book, 1)
	file, err := os.Open("io/products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		readErr := err
		// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
		line = string(line[:len(line)-2])
		//fmt.Printf("The input was: -%s-", line)

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
		if readErr == io.EOF {
			break
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
