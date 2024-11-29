package main

import "os"

func main() {
	// 其解决方案就是在该Go程中引入新的变量（注意对 := 的使用）：
	//...
	//_, err := f1.Write(data)
	//...
	//_, err := f2.Write(data)
	//...
}

// ParallelWrite 将数据写入 file1 和 file2 中，并返回一个错误。
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// 此处的 err 是与主Go程共享的，
			// 因此该写入操作就会与下面的写入操作产生竞争。
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // 第二个冲突的对 err 的写入。
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
