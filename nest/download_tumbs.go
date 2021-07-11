package main
 
import (
  "bufio"
  "fmt"
  "io"
  "net/http"
  "os"
  "path"
  "sync"
)
 
func main() {
  wg := sync.WaitGroup{}
  wg.Add(3)
  var imgs = [...]string{
    "8caac792d5567da81e6846dbda833a57.png",
    "4f90905fd77c1c9456bd5dfe1ceddc34.png",
    "deeaf9d51fc3f13f11f8e1a65553061a.png",
  }

  imgPath := "/Users/xuewenlong/Study/0xfaket.github.io/nest/test/"
  imgUrl := "https://www.michaelfogleman.com/static/nes/"

  for i, c := range imgs {
    fmt.Printf("downloading: %d\n", i)
    go download(imgPath, imgUrl + c, &wg)
  }
  wg.Wait()
}

func download(imgPath string, imgUrl string, wg *sync.WaitGroup) { 
  fileName := path.Base(imgUrl)
 
  res, err := http.Get(imgUrl)
  if err != nil {
    fmt.Println("A error occurred!")
    return
  }
  defer res.Body.Close()
  // 获得get请求响应的reader对象
  reader := bufio.NewReaderSize(res.Body, 32 * 1024)
 
 
  file, err := os.Create(imgPath + fileName)
  if err != nil {
    panic(err)
  }
  // 获得文件的writer对象
  writer := bufio.NewWriter(file)
 
  written, _ := io.Copy(writer, reader)
  fmt.Printf("Total length: %d\n", written)

  wg.Done()
}