package main

// 服务器交互测试
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tttlkkkl/lark"
)

var router *gin.Engine
var chat *lark.Lark

func init() {
	router = gin.Default()
	var err error
	chat, err = lark.NewLark(
		"xx",
		"xx",
		lark.SetReceiveMessageAPI("xx", "xx"),
	)
	if err != nil {
		log.Fatalln("初始化失败", err)
	}
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "", []byte("hello lark!"))
	})
	// 事件处理
	router.Any("/lark/event", func(c *gin.Context) {
		m, err := chat.ReceiveMessage.Handle(c.Request, c.Writer)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusInternalServerError, "", []byte("fail"))
		}
		fmt.Println("--------->>>>>", m, err)
		switch m.GetMessageType() {
		case "":
			// 回调事件
		}
	})
	// card 事件处理
	router.Any("/lark/card/event", func(c *gin.Context) {
		m, err := chat.ReceiveMessage.HandleCard(c.Request, c.Writer)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusInternalServerError, "", []byte("fail"))
		}
		fmt.Println("--------->>>>>", m, err)
	})
}
func main() {
	svr := http.Server{
		Addr:    ":80",
		Handler: router,
	}
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
