package plus

import (
  "flag"
  "fmt"
  "io"
  "io/ioutil"
  "strings"
  "github.com/sundy-li/html2article"
  md "github.com/JohannesKaufmann/html-to-markdown"
)

func args() string {
  var url string
  flag.StringVar(&url, "u", "url", "request url: https://www.baidu.com")
  flag.Parse()
  fmt.Println(url)
  return url
}

func Html2article(urlStr string) string {
  ext, err := html2article.NewFromUrl(urlStr)
  if err != nil {
    panic(err)
  }
  article, err := ext.ToArticle()
  if err != nil {
    panic(err)
  }
  // println("article title is =>", article.Title)
  // println("article publishtime is =>", article.Publishtime) //using UTC timezone
  // println("article content is =>", article.Content)

  //parse the article to be readability
  article.Readable(urlStr)
  // fmt.Println(article.ReadContent)
  return article.ReadContent
}

func Html2Markdown(html string) string {
  converter := md.NewConverter("", true, nil)
  markdown, err := converter.ConvertString(html)
  if err != nil {
    panic(err)
  }
  // fmt.Println(markdown)
  return markdown
}

func Hook(urlStr string) io.ReadCloser {
  // urlStr := args()
  htmlStr := Html2article(urlStr)
  // fmt.Println(htmlStr)
  MdStr := Html2Markdown(htmlStr)
  // fmt.Println(MdStr)
  return ioutil.NopCloser(strings.NewReader(MdStr))
}
