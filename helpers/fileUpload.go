package helpers

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	// Upload the image on the cloud, get the image url and remove from server
	defer func() {
		os.Remove("assets/uploads/" + file.Filename)
	}()

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	var ctx = context.Background()
	resp, err := cld.Upload.Upload(ctx, "assets/uploads/"+file.Filename, uploader.UploadParams{PublicID: "my_avatar" + "-" + file.Filename + "-" + GenerateUid()})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return resp.SecureURL, nil
}
