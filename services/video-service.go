package services

import "golang-gin/entities"

type VideoService interface {
	Save(entities.Video) entities.Video
	FindAll() []entities.Video
}

type videoService struct {
	videos []entities.Video
}

func New() VideoService {
	return &videoService{
		videos: []entities.Video{},
	}
}

func (service *videoService) Save(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entities.Video {
	return service.videos
}
