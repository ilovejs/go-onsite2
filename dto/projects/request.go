package projects

import (
	"github.com/labstack/echo"
	m "onsite/models"
	"time"
)

/* create */
type CreateProjectRequest struct {
	Project struct {
		ManagerID int    `json:"manager_id" validate:"required"`
		CreatorID int    `json:"creator_id" validate:"required"`
		Name      string `json:"name,omitempty" validate:"required"`
		SerialNo  string `json:"serial_no,omitempty"`
		Address   string `json:"address,omitempty"`
		// Value                float64 `json:"total_trade_value,omitempty"`
		// ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue float64 `json:"total_contract_value,omitempty"`
		QuantitySurveyor   string  `json:"quantity_surveyor,omitempty"`
		Notes              string  `json:"notes,omitempty"`
	} `json:"project"`
}

func (r *CreateProjectRequest) Bind(c echo.Context, p *m.Project) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	p.ManagerID = r.Project.ManagerID
	p.CreatorID = r.Project.CreatorID
	// optional
	p.Name.SetValid(r.Project.Name)
	p.SerialNo.SetValid(r.Project.SerialNo)
	p.Address.SetValid(r.Project.Address)
	p.TotalContractValue.SetValid(r.Project.TotalContractValue)
	p.QuantitySurveyor.SetValid(r.Project.QuantitySurveyor)
	p.Notes.SetValid(r.Project.Notes)
	return nil
}

/* update */
type UpdateProjectRequest struct {
	Project struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		// value required by db fk
		ManagerID            int     `json:"manager_id" validate:"required"`
		CreatorID            int     `json:"creator_id" validate:"required"`
		SerialNo             string  `json:"serial_no,omitempty"`
		Address              string  `json:"address,omitempty"`
		TotalTradeValue      float64 `json:"total_trade_value,omitempty"`
		ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue   float64 `json:"total_contract_value,omitempty"`
		QuantitySurveyor     string  `json:"quantity_surveyor,omitempty"`
		Notes                string  `json:"notes,omitempty"`
	} `json:"project"`
}

func (r *UpdateProjectRequest) Bind(c echo.Context, p *m.Project) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	p.Updated.SetValid(time.Now())
	p.Name.SetValid(r.Project.Name)
	p.ManagerID = r.Project.ManagerID
	p.CreatorID = r.Project.CreatorID
	p.SerialNo.SetValid(r.Project.SerialNo)
	p.Address.SetValid(r.Project.Address)
	p.TotalTradeValue.SetValid(r.Project.TotalTradeValue)
	p.ContractorTotalClaim.SetValid(r.Project.ContractorTotalClaim)
	p.TotalContractValue.SetValid(r.Project.TotalContractValue)
	p.QuantitySurveyor.SetValid(r.Project.QuantitySurveyor)
	p.Notes.SetValid(r.Project.Notes)
	return nil
}
