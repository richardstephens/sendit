package routes

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sendit/internal/context"
	"sendit/internal/utils"
)

func HandleUpload(w http.ResponseWriter, r *http.Request, app *context.App) {
	fmt.Println("File Upload Endpoint Hit")

	mpf, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 internal server error getting multipart reader"))
		return
	}

	for {
		part, err := mpf.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 internal server error getting form part"))
			return
		} else if part.FormName() == "file" {
			f, err := os.Create(path.Join(app.Config.UploadPath, utils.SecureFilename(part.FileName())))
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(bufio.NewWriter(f), part)
			if err != nil {
				panic(err)
			}
			err = f.Close()
			if err != nil {
				panic(err)
			}
		}
	}

	w.Write([]byte("OK"))
}
