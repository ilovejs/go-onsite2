package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"onsite/models"
	"time"
)

type ProjectStore struct {
	db *sql.DB
}

func NewProjectStore(db *sql.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (ps *ProjectStore) Create(p *models.Project) error {
	err := p.Insert(ps.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Create Project")
	return nil
}

func (ps *ProjectStore) Read(id int) (*models.Project, error) {
	p, err := models.Projects(Where("id = ? and is_deleted = ?", id, false)).One(ps.db)
	if p == nil {
		return nil, errors.New("project not found, might be deleted")
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *ProjectStore) Update(p *models.Project) error {
	p.Updated.SetValid(time.Now())
	updatedP, err := p.Update(ps.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("Project Updated")
	spew.Dump(updatedP)
	return nil
}

func (ps *ProjectStore) Delete(id int) error {
	p, err := ps.Read(id)
	if err != nil {
		//no record
		return err
	}
	p.IsDeleted = true
	rowAffected, err := p.Update(ps.db, boil.Whitelist("is_deleted"))
	if err != nil {
		return err
	}
	fmt.Println("Delete project row affect: ", rowAffected)
	return nil
}

//GetAll: Active Contractors
func (ps *ProjectStore) Contractors(pid int) ([]models.User, error) {
	panic("no")
	//var users []models.User

	//err := ps.db.Debug().models(&models.User{}).Find(&users).Error
	//if err != nil {
	//	return nil, err
	//}
	//return users, err
}
