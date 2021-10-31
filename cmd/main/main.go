package main

import (
	"log"

	"github.com/webview/webview"
)

func init() {

}

var versionProg string

func main() {
	log.Println("to 0")
	w := webview.New(false) // New(true)
	log.Println("to 1")
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview..HintNone)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
	// w.SetTitle("Hello")
	// w.Bind("noop", func() string {
	// 	log.Panicln("hello")
	// 	return "hello"
	// })
	// w.Bind("add", func(a, b int) int {
	// 	return a + b
	// })
	// w.Bind("quit", func() {
	// 	//	w.Terminate()
	// })
	// log.Println("to navigate")
	// w.Navigate(`data:text/html,
	// 	<!doctype html>
	// 	<html>
	// 		<body>hello</body>
	// 		<script>
	// 			window.onload = function() {
	// 				document.body.innerText = ` + "`hello, ${navigator.userAgent}`" + `;
	// 				noop().then(function(res) {
	// 					console.log('noop res', res);
	// 					add(1, 2).then(function(res) {
	// 						console.log('add res', res);
	// 						quit();
	// 					});
	// 				});
	// 			};
	// 		</script>
	// 	</html>
	// )`)
	// w.Run()
	log.Println("выходим")
	var k chan int
	<-k
}
