package interfaces

import "github.com/Christheoreo/project-manager/models"

type IProjectsRepository interface {
	GetByID(ID int) (project models.Project, err error)
	GetByUser(user models.User) (projects []models.Project, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Insert(project models.NewProject, user models.User) (insertedID int, err error)
	GetOwnerID(projectID int) (userID int, err error)
}

type IProjectsService interface {
	Get(ID int, user models.User) (project models.Project, err error)
	All(user models.User) (projects []models.Project, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Create(newProjectDto models.NewProject, user models.User) (project models.Project, err error)
}
