package foundation

import (
	"os"
	"io/ioutil"
	"fmt"
	"io"
	"bufio"
)

func read1(path string) string {
	// 打开文件
	fi, err := os.Open(path)
	// panic err 关闭文件
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	// chunks 读取 Byte型数组 每次一kb
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func read3(path string) string {
	// 打开文件
	fi, err := os.Open(path)
	// 接收错误，defer 关闭
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	// ioutil 读取
	fd, err := ioutil.ReadAll(fi)
	fmt.Println(string(fd))
	return string(fd)
}

func main() {
	read3("/root/GoglandProjects/GoLearn/ReadMe")
}
