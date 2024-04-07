package handler

import (
	"encoding/base64"
	"io"
	"lgtm-gen/pkg/lgtmgen"
	"lgtm-gen/svc/pkg/application/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LGTMHandler struct {
}

func NewLGTMHandler() *LGTMHandler {
	return &LGTMHandler{}
}

func (l *LGTMHandler) CreateLGTM() gin.HandlerFunc {
	return func(c *gin.Context) {
		imgData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		// LGTM画像を作成するメソッドを実行
		lgtm, err := lgtmgen.Generate(imgData)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// NOTE:本当はGCSかどこかに保存
		lgtmBase64 := base64.StdEncoding.EncodeToString(lgtm)

		res := response.CreateLGTMResponse{
			ImageUrl: lgtmBase64,
		}

		c.JSON(http.StatusCreated, res)
	}
}