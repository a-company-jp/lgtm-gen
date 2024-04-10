package handler

import (
	"encoding/base64"
	"io"
	"lgtm-gen/pkg/lgtmgen"
	"lgtm-gen/svc/pkg/application/response"
	"lgtm-gen/svc/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LGTMHandler struct {
	lgtmRepo domain.ILGTMRepository
}

func NewLGTMHandler(lgtmRepo domain.ILGTMRepository) *LGTMHandler {
	return &LGTMHandler{
		lgtmRepo: lgtmRepo,
	}
}

func (l LGTMHandler) CreateLGTM() gin.HandlerFunc {
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

		// FireStoreにデータを保存
		id := uuid.New().String()
		err = l.lgtmRepo.Create(id)
		if err != nil {
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
