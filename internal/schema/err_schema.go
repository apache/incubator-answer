package schema

type ErrTypeData struct {
	ErrType string `json:"err_type"`
}

var ErrTypeModal = ErrTypeData{ErrType: "modal"}

var ErrTypeToast = ErrTypeData{ErrType: "toast"}

var ErrTypeAlert = ErrTypeData{ErrType: "alert"}
