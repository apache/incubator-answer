package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/base/validator"
	myErrors "github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// HandleResponse Handle response body
func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	lang := GetLang(ctx)
	// no error
	if err == nil {
		ctx.JSON(http.StatusOK, NewRespBodyData(http.StatusOK, reason.Success, data).TrMsg(lang))
		return
	}

	var myErr *myErrors.Error
	// unknown error
	if !errors.As(err, &myErr) {
		log.Error(err, "\n", myErrors.LogStack(2, 5))
		ctx.JSON(http.StatusInternalServerError, NewRespBody(
			http.StatusInternalServerError, reason.UnknownError).TrMsg(lang))
		return
	}

	// log internal server error
	if myErrors.IsInternalServer(myErr) {
		log.Error(myErr)
	}

	respBody := NewRespBodyFromError(myErr).TrMsg(lang)
	if data != nil {
		respBody.Data = data
	}
	ctx.JSON(myErr.Code, respBody)
	return
}

// BindAndCheck bind request and check
func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	if err := ctx.ShouldBind(data); err != nil {
		log.Errorf("http_handle BindAndCheck fail, %s", err.Error())
		HandleResponse(ctx, myErrors.New(http.StatusBadRequest, reason.RequestFormatError), nil)
		return true
	}

	errField, err := validator.GetValidatorByLang(GetLang(ctx).Abbr()).Check(data)
	if err != nil {
		HandleResponse(ctx, myErrors.New(http.StatusBadRequest, reason.RequestFormatError).WithMsg(err.Error()), errField)
		return true
	}
	return false
}
