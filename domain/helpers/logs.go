package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func InitializeLogDir() {
	logFile := "/payment-gateway.log"
	logDir := "logs"

	_ = os.Mkdir(logDir, os.ModePerm)
	f, err := os.OpenFile(logDir+logFile, os.O_RDWR|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file:%v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogRequest(ctx *gin.Context) {
	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: ctx.Writer,
	}

	ctx.Writer = blw
	ctx.Next()
	statusCode := ctx.Writer.Status()
	response := "No errors found"

	level := "INFO"

	if statusCode >= 400 {
		response = blw.body.String()
		level = "ERROR"
	}
	data, err := json.Marshal(&LogStruct{
		TimeStamp:       time.Now().Format(time.RFC3339),
		Version:         "1",
		Level:           level,
		StatusCode:      strconv.Itoa(statusCode),
		Message:         http.StatusText(statusCode) + ":" + response,
		LoggerName:      "",
		AppName:         "Payment-Gateway",
		Path:            ctx.Request.URL.String(),
		Method:          ctx.Request.Method,
		CorrelationId:   uuid.New().String(),
		UserAgent:       ctx.Request.Header.Get("User-Agent"),
		ResponseTime:    time.Since(time.Now()).String(),
		ApplicationHost: ctx.Request.Host,
		RemoteIP:        ctx.ClientIP(),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)
	ctx.Next()
}

func LogEvent(level string, message interface{}) {
	data, err := json.Marshal(struct {
		TimeStamp string      `json:"time_stamp"`
		Level     string      `json:"level"`
		Message   interface{} `json:"message"`
	}{
		TimeStamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", data)
}
