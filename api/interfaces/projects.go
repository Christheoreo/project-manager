package interfaces

import "github.com/Christheoreo/project-manager/models"

type IProjectsRepository interface {
	GetByID(ID int) (project models.ProjectDto, err error)
	GetByUser(user models.UserDto) (projects []models.ProjectDto, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Insert(project models.NewProjectDto, user models.UserDto) (insertedID int, err error)
	GetOwnerID(projectID int) (userID int, err error)
}

type IProjectsService interface {
	Get(ID int, user models.UserDto) (project models.ProjectDto, err error)
	All(user models.UserDto) (projects []models.ProjectDto, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Create(newProjectDto models.NewProjectDto, user models.UserDto) (project models.ProjectDto, err error)
}
