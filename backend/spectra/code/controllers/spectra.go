package controllers

import (
	"log"
	"net/http"
	"spectra/interfaces"
	"spectra/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SpectraController struct{}

func (s *SpectraController) CreateSpectra(ctx *gin.Context) {
	spectraServices := services.SpectraServices{}
	input := services.CreateSpectraInput{}

	if err := ctx.MustBindWith(&input, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	spectraFile, err := ctx.FormFile("spectra_file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userOwnerEmail := ctx.GetString("user_owner_email")

	if userOwnerEmail == "" {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Data:   "Undefined email of owner spectra",
			Status: http.StatusInternalServerError,
		})
		return
	}

	input.EmailOwner = userOwnerEmail
	hexId, errorCreate := spectraServices.CreateSpectraService(input, spectraFile)

	if errorCreate.Message != "" {
		ctx.JSON(errorCreate.StatusCode(), interfaces.ErrorResponse{
			Data:   errorCreate.Error(),
			Status: errorCreate.StatusCode(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.SpectraCreatedResponse{
		Data:   interfaces.SpectraIdResponse{Id: hexId},
		Status: http.StatusCreated,
	})
}

func (s *SpectraController) ListByOwner(ctx *gin.Context) {
	spectraServices := services.SpectraServices{}

	userOwnerEmail := ctx.GetString("user_owner_email")

	if userOwnerEmail == "" {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Data:   "Undefined email of owner spectra",
			Status: http.StatusInternalServerError,
		})
		return
	}

	records, errorCreate := spectraServices.ListByOwner(userOwnerEmail)

	if errorCreate.Message != "" {
		ctx.JSON(errorCreate.StatusCode(), interfaces.ErrorResponse{
			Data:   errorCreate.Error(),
			Status: errorCreate.StatusCode(),
		})
		return
	}

	if len(records) == 0 {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, records)
}

func (s *SpectraController) GetById(ctx *gin.Context) {
	spectraServices := services.SpectraServices{}
	id := ctx.Param("id")
	record, errorCreate := spectraServices.GetById(id)

	if errorCreate.Message != "" {
		ctx.JSON(errorCreate.StatusCode(), interfaces.ErrorResponse{
			Data:   errorCreate.Error(),
			Status: errorCreate.StatusCode(),
		})
		return
	}
	ctx.JSON(http.StatusOK, record)
}

func (s *SpectraController) UpdatePrediction(ctx *gin.Context) {
	spectraServices := services.SpectraServices{}
	input := services.UpdatePredictionSpectra{}
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, interfaces.ErrorResponse{
			Data:   "Id not provided",
			Status: http.StatusBadRequest,
		})
		return
	}

	if err := ctx.ShouldBind(&input); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println("Sending data for service")
	record, errorUpdatePrediction := spectraServices.UpdatePrediction(id, input)

	if errorUpdatePrediction.Message != "" {
		log.Println(errorUpdatePrediction.Message)
		ctx.JSON(errorUpdatePrediction.StatusCode(), interfaces.ErrorResponse{
			Data:   errorUpdatePrediction.Error(),
			Status: errorUpdatePrediction.StatusCode(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, record)
}
