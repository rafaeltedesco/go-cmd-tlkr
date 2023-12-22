package presenter

import (
	"fmt"
	"talker/manager/api/src/models"
	"talker/manager/api/src/services"
)

type Presenter struct {
	crudService services.Crud
}

func New(crudService services.Crud) *Presenter {
	return &Presenter{
		crudService: crudService,
	}
}

func (p *Presenter) DisplayTalkers() {
	talkers := p.crudService.GetTalkers()
	fmt.Println(talkers)
}

func (p *Presenter) DisplayTalkerById(id int) {
	if talker, err := p.crudService.GetTalkerById(id); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(talker)
	}
}

func (p *Presenter) CreateAndDisplay(jsonData string) {
	var talker models.Talker
	p.crudService.CreateTalker(jsonData, &talker)
	fmt.Printf("New talker created %v", talker)
}

func (p *Presenter) DeleteAndDisplay(id int) {
	if success := p.crudService.DeleteTalker(id); success {
		fmt.Printf("Talker with id %d delete successfully", id)
	}
}
