package libs

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"path/filepath"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func CompressImage(img graphql.Upload) ([]byte, error) {
	var err error
	img.Filename, err = preprocessImage(img.Filename)
	if err != nil {
		return nil, err
	}

	srcImg, _, err := image.Decode(img.File)
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(1080, 0, srcImg, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, resizedImg, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func preprocessImage(imgFileName string) (string, error) {
	re := regexp.MustCompile(`\.(png|jpe?g)$`)

	if !re.MatchString(imgFileName) {
		return "", fmt.Errorf("unsupported image format")
	}

	ext := filepath.Ext(imgFileName)
	imageUUID := uuid.New()

	newFileName := fmt.Sprintf("%s%s", imageUUID, ext)

	return newFileName, nil
}
