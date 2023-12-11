package service

import (
	"context"
	"log"

	"github.com/felipefbs/grpc/internal/databases"
	"github.com/felipefbs/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	repo *databases.CategoryRepository
}

func NewCategoryService(repo *databases.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (svc *CategoryService) CreateCategory(ctx context.Context, request *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	log.Print("Creating category")

	category, err := svc.repo.Create(request.Name, request.Description)
	if err != nil {
		log.Println("failed to create category")

		return nil, err
	}

	response := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return response, nil
}
