//bin存放编译后的可执行文件；pkg存放编译后的包文件；src存放项目源文件。
// 一般，bin和pkg目录可以不创建，go命令会自动创建（如 go install），只需要创建src目录即可。
package foundation

import (
	"fmt"
	"os"
	"path/filepath"
	"wordcount"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	wordcounter := make(wordcount.WordCount)
	// for _, filename := range os.Args[1:] {
	//  wordcount.UpdateFreq(filename)
	// }
	wordcounter.WordFreqCounter(os.Args[1:])

	wordcounter.SortReport()
}
