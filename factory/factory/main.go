package factory

import (
	"errors"

	"github.com/amine-amaach/design-pattern/factory/contracts"
	"github.com/amine-amaach/design-pattern/factory/guns"
)

func NewGun(gunType string) (contracts.IGun, error) {
	switch gunType {
	case "ak47":
		ak47 := &guns.AK47{
			Name:  "AK47 gun",
			Power: 4,
		}
		return ak47, nil
	case "musket":
		musket := &guns.Musket{
			Name:  "Musket gun",
			Power: 1,
		}
		return musket, nil
	default:
		return nil, errors.New("wrong gun type passed")
	}
}
