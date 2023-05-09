package plugin

type UploadSource string

const (
	UserAvatar    UploadSource = "user_avatar"
	UserPost      UploadSource = "user_post"
	AdminBranding UploadSource = "admin_branding"
)

var (
	DefaultFileTypeCheckMapping = map[UploadSource]map[string]bool{
		UserAvatar: {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
		UserPost: {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
		AdminBranding: {
			".ico": true,
		},
	}
)

type UploadFileResponse struct {
	// FullURL is the URL that can be used to access the file
	FullURL string
	// OriginalError is the error returned by the storage plugin. It is used for debugging.
	OriginalError error
	// DisplayErrorMsg is the error message that will be displayed to the user.
	DisplayErrorMsg Translator
}

type Storage interface {
	Base

	// UploadFile uploads a file to storage.
	// The file is in the Form of the ctx and the key is "file"
	UploadFile(ctx *GinContext, source UploadSource) UploadFileResponse
}

var (
	// CallStorage is a function that calls all registered storage
	CallStorage,
	registerStorage = MakePlugin[Storage](false)
)
