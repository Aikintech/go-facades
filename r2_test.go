package gofacades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestR2(t *testing.T) {
	LoadEnv()

	assert := assert.New(t)
	accountId := GetEnv().GetString("CLOUDFLARE_ACCOUNT_ID")
	url := GetEnv().GetString("CLOUDFLARE_CDN")

	config := S3Config{
		Bucket:          GetEnv().GetString("CLOUDFLARE_BUCKET"),
		AccessKeyId:     GetEnv().GetString("CLOUDFLARE_ACCESS_KEY_ID"),
		AccessKeySecret: GetEnv().GetString("CLOUDFLARE_ACCESS_KEY_SECRET"),
		AccountId:       &accountId,
		Region:          GetEnv().GetString("CLOUDFLARE_REGION"),
		CdnURL:          &url,
	}

	r2 := NewR2(config)

	t.Run("delete file", func(t *testing.T) {
		err := r2.Delete("avatar/OIcQuNKcNhHFApygXERxQdFWCAeuWTasfBagXFPK.png")

		assert.NoError(err)
	})

	t.Run("file exits", func(t *testing.T) {
		result := r2.Exist("avatar/yNlLvdcJCcQMvUuwjIywpOKZcBtjhyeoZwqTUcVJ.png")

		assert.Equal(true, result)
	})

	t.Run("get bytes", func(t *testing.T) {
		result, err := r2.GetBytes("avatar/yNlLvdcJCcQMvUuwjIywpOKZcBtjhyeoZwqTUcVJ.png")

		assert.NoError(err)
		assert.NotEmpty(result)
	})

	t.Run("mime type", func(t *testing.T) {
		result, err := r2.MimeType("avatar/yNlLvdcJCcQMvUuwjIywpOKZcBtjhyeoZwqTUcVJ.png")

		assert.NoError(err)
		assert.NotEmpty(result)
		assert.Equal("application/octet-stream", result)
	})

	t.Run("size", func(t *testing.T) {
		result, err := r2.Size("avatar/yNlLvdcJCcQMvUuwjIywpOKZcBtjhyeoZwqTUcVJ.png")

		assert.NoError(err)
		assert.NotEmpty(result)
	})

	t.Run("url", func(t *testing.T) {
		result := r2.Url("avatar/yNlLvdcJCcQMvUuwjIywpOKZcBtjhyeoZwqTUcVJ.png")

		assert.NotEmpty(result)
	})

	t.Run("put or put file", func(t *testing.T) {

	})
}
