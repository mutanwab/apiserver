package handlers

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"time"

	"github.com/rancher/apiserver/pkg/apierror"
	"github.com/rancher/apiserver/pkg/metrics"
	"github.com/rancher/apiserver/pkg/types"
)

func MetricsHandler(successCode string, next func(apiRequest *types.APIRequest) (types.APIObject, error)) func(apiRequest *types.APIRequest) (types.APIObject, error) {
	return func(request *types.APIRequest) (types.APIObject, error) {
		obj, err := next(request)
		if err != nil {
			if apiError, ok := err.(*apierror.APIError); ok {

				metrics.IncTotalResponses(request.Schema.ID, request.Method, strconv.Itoa(apiError.Code.Status))
			}
			return types.APIObject{}, err
		}

		metrics.IncTotalResponses(request.Schema.ID, request.Method, successCode)
		return obj, err
	}
}

func MetricsListHandler(successCode string, next func(apiRequest *types.APIRequest) (types.APIObjectList, error)) func(apiRequest *types.APIRequest) (types.APIObjectList, error) {
	return func(request *types.APIRequest) (types.APIObjectList, error) {
		logrus.Info("MetricsListHandler 1: type: %s, flag: %s, time: %s", request.Request.URL, request.Request.Header.Get("flag"), time.Now().String())
		objList, err := next(request)
		if err != nil {
			if apiError, ok := err.(*apierror.APIError); ok {
				metrics.IncTotalResponses(request.Schema.ID, request.Method, strconv.Itoa(apiError.Code.Status))
			}
			return types.APIObjectList{}, err
		}
		logrus.Info("MetricsListHandler 2: type: %s, flag: %s, time: %s", request.Request.URL, request.Request.Header.Get("flag"), time.Now().String())
		metrics.IncTotalResponses(request.Schema.ID, request.Method, successCode)
		logrus.Info("MetricsListHandler 3: type: %s, flag: %s, time: %s", request.Request.URL, request.Request.Header.Get("flag"), time.Now().String())
		return objList, err
	}
}
