package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	gproto "matsu.dev/matsuri-contest-log/proto"
)

var (
	guiClient         gproto.GuiClient
	radioClient       gproto.RadioClient
	realTimeGuiClient gproto.RealtimeGuiClient
)

// We should notice that the performance of reflect, however,
// as we are expecting relative low frequency of these APIs,
// the performance should be acceptable.

func handleApiCall(class interface{}, api string, c *gin.Context) (interface{}, error) {
	var (
		err error
		req interface{}
	)

	fun := reflect.ValueOf(class).MethodByName(api)
	if !fun.IsValid() {
		return nil, fmt.Errorf("method %s not found", api)
	}

	var callArgs []reflect.Value
	callType := fun.Type()
	argType := callType.In(1).Elem()
	req = reflect.New(argType).Interface()

	if err = c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		return nil, fmt.Errorf("failed to parse request: %v", err)
	}

	callArgs = []reflect.Value{
		reflect.ValueOf(c.Request.Context()),
		reflect.ValueOf(req),
	}

	response := fun.Call(callArgs)
	if len(response) != 2 {
		panic("unexpected number of response")
	}

	gotError := !response[1].IsNil()
	if gotError {
		err, ok := response[1].Interface().(error)

		if ok {
			return nil, err
		}
		return nil, fmt.Errorf("unknown error: %v", err)
	}

	return response[0].Interface(), nil
}

func handleHttpApiCall(c *gin.Context) {
	api := struct {
		Class string `uri:"class" binding:"required"`
		Api   string `uri:"api" binding:"required"`
	}{}

	err := c.BindUri(&api)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = fmt.Errorf("no such component: %v", api.Class)

	var res interface{}
	switch api.Class {
	case "Gui":
		res, err = handleApiCall(guiClient, api.Api, c)
	case "Radio":
		res, err = handleApiCall(radioClient, api.Api, c)
	case "RealtimeGui":
		res, err = handleApiCall(realTimeGuiClient, api.Api, c)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    res,
	})
}
