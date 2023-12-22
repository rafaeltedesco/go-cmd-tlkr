package services

import (
	"encoding/json"
	"errors"
	"talker/manager/api/src/models"
	talkerutils "talker/manager/api/src/utils/io"
)

const FILEPATH = "./src/files/people.json"

var TalkerNotFound = errors.New("Talker Not Found")

var talkers []models.Talker = []models.Talker{}

type crud struct {
	ioOperator *talkerutils.IoUtils
}

func New(ioOperator *talkerutils.IoUtils) *crud {
	return &crud{ioOperator: ioOperator}
}

func (c *crud) GetTalkers() []models.Talker {
	c.ioOperator.ReadFile(FILEPATH, &talkers)
	return talkers
}

func (c *crud) GetTalkerById(id int) (models.Talker, error) {
	c.ioOperator.ReadFile(FILEPATH, &talkers)
	for _, talker := range talkers {
		if talker.Id == id {
			return talker, nil
		}
	}
	return models.Talker{}, TalkerNotFound
}

func (c *crud) CreateTalker(data string, talkerData *models.Talker) {
	json.Unmarshal([]byte(data), talkerData)
	talkers := c.GetTalkers()

	talkerData.Id = talkers[len(talkers)-1].Id + 1
	talkers = append(talkers, *talkerData)
	if updated, err := json.MarshalIndent(talkers, "", "    "); err != nil {
		panic(err.Error())
	} else {
		c.ioOperator.WriteFile(FILEPATH, updated)
	}
}

func (c *crud) DeleteTalker(id int) bool {
	talker, err := c.GetTalkerById(id)
	if err != nil {
		panic("Talker Not found!")
	}
	talkers := c.GetTalkers()
	var newTalkersList []models.Talker
	for _, tlkr := range talkers {
		if tlkr.Id != talker.Id {
			newTalkersList = append(newTalkersList, tlkr)
		}
	}
	if jsonData, err := json.MarshalIndent(newTalkersList, "", "   "); err != nil {
		panic(err.Error())
	} else {
		c.ioOperator.WriteFile(FILEPATH, jsonData)
	}
	return true
}
