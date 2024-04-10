package handler

import (
	"fmt"
	"io"
	"lgtm-gen/pkg/config"
	"lgtm-gen/pkg/lgtmgen"
	"lgtm-gen/svc/pkg/application/response"
	"lgtm-gen/svc/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LGTMHandler struct {
	lgtmTableRepo  domain.ILGTMTableRepository
	lgtmBucketRepo domain.ILGTMBucketRepository
}

func NewLGTMHandler(lgtmTableRepo domain.ILGTMTableRepository, lgtmBucketRepo domain.ILGTMBucketRepository) *LGTMHandler {
	return &LGTMHandler{
		lgtmTableRepo:  lgtmTableRepo,
		lgtmBucketRepo: lgtmBucketRepo,
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

		// GCSに保存
		id := uuid.New().String()
		err = l.lgtmBucketRepo.Create(id, lgtm)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// FireStoreにデータを保存
		conf := config.Get()
		url := fmt.Sprintf("https://storage.googleapis.com/%v/%v", conf.Application.GCS.BucketName, id)
		err = l.lgtmTableRepo.Create(id, url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := response.CreateLGTMResponse{
			ImageUrl: url,
		}

		c.JSON(http.StatusCreated, res)
	}
}
