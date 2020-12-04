package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//新建空文件
	test1()
	//新建空文件夹
	test1_1()
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。
	test2()
	//查看文件信息
	test3()
	//移动和重命名文件
	test4()
	//删除文件
	test5()
	//打开和关闭文件
	test6()
	//检查文件是否存在
	test7()
	//检查文件读写权限
	test8()
	//改变权限、拥有者、时间戳
	test9()
	//硬连接和软连接
	test10()
	//复制文件
	test11()
	//跳转到文件指定位置(Seek)
	test12()
	//写文件
	test13()
	//快写文件
	test14()
	//缓存写
	test15()
	//读取最多N个字节
	test16()
	//正好读取N个字节
	test17()
	//读取至少N个字节
	test18()
	//读取全部字节
	test19()
	//快读到内存
	test20()
	//缓存读
	test21()
	//使用 scanner
	test22()
	//打包(zip) 文件
	test23()
	//抽取(unzip) 文件
	test24()
	//gzip压缩文件
	test25()
	//gzip解压缩文件
	test26()
	//临时文件和目录
	test27()
}

func test1() {
	newFile, err := os.Create("test.txt") //该文件创建在项目的根目录中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newFile)
	_ = newFile.Close()
}

func test1_1() {
	_ = createFile("D:/gopro/src/aa/bb/cc")
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func test2() {
	err := os.Truncate("test.txt", 100)
	if err != nil {
		fmt.Println(err)
	}
}

func test3() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func test4() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		fmt.Println(err)
	}
}

func test5() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func test6() {
	// 简单地以只读的方式打开，其路径是根路径，且不能带/
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	_ = file.Close()
	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	_ = file.Close()
}

func test7() {
	//err不为空则表示文件不存在
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exists.")
		}
	}
	fmt.Println("File does exist. File information:")
	fmt.Println(fileInfo)
}

func test8() {
	//写权限
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Error: Write permission denied.")
		}
	}
	_ = file.Close()
	//读权限
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Error: Read permission denied.")
		}
	}
	_ = file.Close()
}

func test9() {
	// 使用Linux风格改变文件权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		fmt.Println(err)
	}
	// 改变文件所有者
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		fmt.Println(err)
	}
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		fmt.Println(err)
	}
}

func test10() {
	// 创建一个硬链接。
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	// 删除和重命名不会影响另一个。
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("creating sym")
	// Create a symlink
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		fmt.Println(err)
	}
	// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	// Symlink在Windows中不工作。
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)
	//改变软链接的拥有者不会影响原始文件。
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		fmt.Println(err)
	}
}

func test11() {
	// 打开原始文件
	originalFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Copied %d bytes.", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

func test12() {
	file, _ := os.Open("test.txt")
	defer file.Close()
	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5
	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Just moved to 5:", newPosition)
	// 从当前位置回退两个字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Just moved back two:", newPosition)
	// 使用下面的技巧得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)
	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)
}

func test13() {
	// 可写方式打开文件
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %d bytes.\n", bytesWritten)
}

func test14() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func test15() {
	// 打开文件，只写
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)
	// 写字节到buffer
	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)
	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)
	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	fmt.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	_ = bufferedWriter.Flush()
	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)
	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)
	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Available buffer: %d\n", bytesAvailable)
}

func test16() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了
	// 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes read: %d\n", bytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
}

func test17() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
}

func test18() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	byteSlice := make([]byte, 512)
	minBytes := 8
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
}

func test19() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

func test20() {
	// 读取文件到byte slice中
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data read: %s\n", data)
}

func test21() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	bufferedReader := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)
	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)
	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)
	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Read string: %s\n", dataString)
	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}

func test22() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanWords)
	// scan下一个token.
	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err == nil {
			fmt.Println("Scan completed and reached EOF")
		} else {
			fmt.Println(err)
		}
	}
	// 得到数据，Bytes() 或者 Text()
	fmt.Println("First word found:", scanner.Text())
	// 再次调用scanner.Scan()发现下一个token
}

func test23() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()
	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)
	// 往打包文件中写文件。
	// 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}
	// 下面将要打包的内容写入到打包文件中，依次写入。
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}
	// 清理
	err = zipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func test24() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer zipReader.Close()
	// 遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer zippedFile.Close()
		// 指定抽取的文件名。
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
		// 我们这个例子使用打包文件中相同的文件名。
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			fmt.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//抽取正常的文件
			fmt.Println("Extracting file:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				fmt.Println(err)
			}
			defer outputFile.Close()
			// 通过io.Copy简洁地复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func test25() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Println(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	// 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
	// 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Compressed data written to file.")
}

func test26() {
	// 打开一个gzip文件。
	// 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
	// 它的内容不是一个文件，而是一个内存流
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		fmt.Println(err)
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println(err)
	}
	defer gzipReader.Close()
	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer outfileWriter.Close()
	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		fmt.Println(err)
	}
}

func test27() {
	// 在系统临时文件夹中创建一个临时文件夹
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Temp dir created:", tempDirPath)
	// 在临时文件夹中创建临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())
	// ... 做一些操作 ...
	// 关闭文件
	err = tempFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 删除我们创建的资源
	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		fmt.Println(err)
	}
}
