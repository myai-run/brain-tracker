package repository

import (
	"airun/brain-tracker/server/src/configs"
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/types"
	"airun/brain-tracker/server/src/utils"
)

func GetSignalData(params types.GetSingalDataParamsDTO) ([]models.SignalData, error) {
	var result []models.SignalData
	validatedParams := utils.GetValidatedSignalDataParams(params)
	tx := configs.DB.Offset(*validatedParams.Offset).Limit(*validatedParams.Limit).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}

func CreateSignalData(data *models.SignalData) error {
	result := configs.DB.Create(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteSignalDataByID(id int) error {
	result := configs.DB.Where("id = ?", id).Delete(&models.SignalData{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
