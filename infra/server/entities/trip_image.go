package entities

import "mime/multipart"

type TripImage struct {
	Filename string
	File     multipart.File
}
