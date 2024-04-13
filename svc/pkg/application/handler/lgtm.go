package handler

import (
	"fmt"
	"io"
	"lgtm-gen/pkg/config"
	"lgtm-gen/pkg/lgtmgen"
	"lgtm-gen/pkg/snowflake"
	"lgtm-gen/svc/pkg/application/response"
	"lgtm-gen/svc/pkg/domain"
	"lgtm-gen/svc/pkg/domain/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LGTMHandler struct {
	lgtmTableRepo  domain.ILGTMTableRepository
	lgtmBucketRepo domain.ILGTMBucketRepository
	safeSearchRepo domain.ISafeSearchRepository
	idGen          snowflake.Snowflake
}

func NewLGTMHandler(lgtmTableRepo domain.ILGTMTableRepository, lgtmBucketRepo domain.ILGTMBucketRepository, safeSearchRepo domain.ISafeSearchRepository) *LGTMHandler {
	return &LGTMHandler{
		lgtmTableRepo:  lgtmTableRepo,
		lgtmBucketRepo: lgtmBucketRepo,
		safeSearchRepo: safeSearchRepo,
		idGen:          snowflake.NewSnowflake(),
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

		// セーフサーチ
		annotations, err := l.safeSearchRepo.Detect(imgData)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(annotations) > 0 {
			log.Printf("detected: %v", annotations)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Inappropriate image")})
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
		id := l.idGen.String()
		err = l.lgtmBucketRepo.Create(id, lgtm)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// FireStoreにデータを保存
		conf := config.Get()
		url := fmt.Sprintf("https://storage.googleapis.com/%v/%v", conf.Application.GCS.BucketName, id)
		err = l.lgtmTableRepo.Create(model.LGTM{
			ID:  id,
			Url: url,
		})
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

func (l LGTMHandler) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		lgtms, err := l.lgtmTableRepo.List()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := response.NewGetListResponse(lgtms)

		c.JSON(http.StatusOK, res)
	}
}
