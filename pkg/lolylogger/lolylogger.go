package lolylogger

import (
	"fmt"
	"log"
	"lolyshop/pkg/utils"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ILolyLogger interface {
	Print() ILolyLogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type lolyLogger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitLolyLogger(c *fiber.Ctx, res any) ILolyLogger {
	log := &lolyLogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: c.Response().StatusCode(),
	}

	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)

	return log
}

func (l *lolyLogger) Print() ILolyLogger {
	utils.Debug(l)
	return l
}

func (l *lolyLogger) Save() {
	data := utils.Output(l)
	filename := fmt.Sprintf("./assets/logs/lolylogger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	defer file.Close()
	file.WriteString(string(data) + "")
}

func (l *lolyLogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}
	l.Query = body
}

func (l *lolyLogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "sad"
	default:
		l.Body = body
	}
}

func (l *lolyLogger) SetResponse(res any) {
	l.Response = res
}
