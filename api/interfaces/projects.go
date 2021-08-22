package interfaces

import "github.com/Christheoreo/project-manager/dtos"

type IProjectsRepository interface {
	GetByID(ID int) (project dtos.ProjectDto, err error)
	GetByUser(user dtos.UserDto) (projects []dtos.ProjectDto, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Insert(project dtos.NewProjectDto, user dtos.UserDto) (insertedID int, err error)
	GetOwnerID(projectID int) (userID int, err error)
}

type IProjectsService interface {
	Get(ID int, user dtos.UserDto) (project dtos.ProjectDto, err error)
	All(user dtos.UserDto) (projects []dtos.ProjectDto, err error)
	TitleTaken(title string, userID int) (isTaken bool, err error)
	Create(newProjectDto dtos.NewProjectDto, user dtos.UserDto) (project dtos.ProjectDto, err error)
}
