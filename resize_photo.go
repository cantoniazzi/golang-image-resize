package main


import (
	"os"
	"net/http"
	"io"
	"github.com/nfnt/resize"
	"log"
	"image/jpeg"
	"sync"
	"fmt"
)

type filePhoto struct {
	fileName string
	urlFile string
	main bool
	updated bool
}

func downloadFile(fileName string, url string) (err error) {
	filePath := "files/"

	// get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer resp.Body.Close()

	// create empty file
	out, err := os.Create(filePath + fileName)
	if err != nil  {
		fmt.Print(err)
		return err
	}
	defer out.Close()

	// writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		fmt.Print(err)
		return err
	}

	return nil
}

func resizePhoto(fileName string) {
	sizes := [][]uint{}

	widths := []uint{540, 65, 80, 311, 250, 372, 45, 225, 250, 158, 80, 113, 113}
	heights := []uint{900, 110, 100, 415, 300, 620, 60, 300, 300, 158, 100, 115, 150}

	sizes = append(sizes, widths)
	sizes = append(sizes, heights)

	var wg sync.WaitGroup
	wg.Add(len(widths))

	for i := 0; i < len(widths); i++ {

		go func(i int) {
			filePath := "files/"

			file, err := os.Open(filePath + fileName)
			if err != nil {
				log.Fatal(err)
			}

			// decode jpeg into image.Image
			img, err := jpeg.Decode(file)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			m := resize.Resize(widths[i], heights[i], img, resize.Lanczos3)

			filePhotoResized := filePath + fmt.Sprint(widths[i]) + "X" + fmt.Sprint(heights[i])
			filePhotoResized += "_" + fileName

			out, err := os.Create(filePhotoResized)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// write new image to file
			jpeg.Encode(out, m, nil)

			defer wg.Done()
		}(i)

	}

	//wg.Wait()
}

func main() {
	var photos = [10]filePhoto{}

	var wg sync.WaitGroup
	wg.Add(len(photos))

	photos[0] = filePhoto{
		"d8df12a9-f696-409b-893e-7b3e388171a7.jpg",
		"https://thumbor.softimob.com.br/d8df12a9-f696-409b-893e-7b3e388171a7.jpg",
		false,
		true,
	}
	photos[1] = filePhoto{
		"f13af8cd-b766-42f1-9fe1-9391fa5c8287.jpg",
		"https://thumbor.softimob.com.br/f13af8cd-b766-42f1-9fe1-9391fa5c8287.jpg",
		false,
		true,
	}
	photos[2] = filePhoto{
		"820241df-6504-41c6-9a8a-c19037f395ad.jpg",
		"https://thumbor.softimob.com.br/820241df-6504-41c6-9a8a-c19037f395ad.jpg",
		false,
		true,
	}
	photos[3] = filePhoto{
		"118aab97-9d8d-4f2d-a723-f2e2f5246b99.jpg",
		"https://thumbor.softimob.com.br/118aab97-9d8d-4f2d-a723-f2e2f5246b99.jpg",
		false,
		true,
	}
	photos[4] = filePhoto{
		"854634e1-0afc-40cf-b050-24318e02141c.jpg",
		"https://thumbor.softimob.com.br/854634e1-0afc-40cf-b050-24318e02141c.jpg",
		false,
		true,
	}
	photos[5] = filePhoto{
		"ac0ae905-d265-4bfa-aade-3fd7ec27af12.jpg",
		"https://thumbor.softimob.com.br/ac0ae905-d265-4bfa-aade-3fd7ec27af12.jpg",
		false,
		true,
	}
	photos[6] = filePhoto{
		"eb3cb6f8-3615-4c0c-aaa3-b9dad1728a3f.jpg",
		"https://thumbor.softimob.com.br/eb3cb6f8-3615-4c0c-aaa3-b9dad1728a3f.jpg",
		false,
		true,
	}
	photos[7] = filePhoto{
		"fa594d4b-505c-4a56-89c6-e4bc8abb9468.jpg",
		"https://thumbor.softimob.com.br/fa594d4b-505c-4a56-89c6-e4bc8abb9468.jpg",
		false,
		true,
	}
	photos[8] = filePhoto{
		"f7ad88f8a-3588-458b-aed3-e31ef8be5e12.jpg",
		"https://thumbor.softimob.com.br/7ad88f8a-3588-458b-aed3-e31ef8be5e12.jpg",
		false,
		true,
	}
	photos[9] = filePhoto{
		"7ec9516f-b60e-488f-9c4d-225f19004689.jpg",
		"https://thumbor.softimob.com.br/7ec9516f-b60e-488f-9c4d-225f19004689.jpg",
		false,
		true,
	}

	for i := 0; i < len(photos); i++ {

		go func(i int) {
			downloadFile(photos[i].fileName, photos[i].urlFile)
			resizePhoto(photos[i].fileName)

			defer wg.Done()
		}(i)

	}

	wg.Wait()
}
