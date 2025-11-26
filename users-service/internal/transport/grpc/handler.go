package grpc

import (
	"context"

	userpb "github.com/PBGlory/project-protos/proto/user"
	"github.com/PBGlory/users-service/internal/user"
)

// Handler реализует gRPC сервер
type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

// NewHandler создаёт новый gRPC Handler
func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

// mapUser конвертирует internal/user.User в protobuf User
func mapUser(u *user.User) *userpb.User {
	return &userpb.User{
		Id:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

// CreateUser gRPC метод
func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.Create(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: mapUser(u),
	}, nil
}

// GetUser gRPC метод
func (h *Handler) GetUser(ctx context.Context, req *userpb.User) (*userpb.User, error) {
	u, err := h.svc.Get(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return mapUser(u), nil
}

// ListUsers gRPC метод
func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.List()
	if err != nil {
		return nil, err
	}

	result := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		result = append(result, mapUser(&u))
	}

	return &userpb.ListUsersResponse{User: result}, nil
}

// UpdateUser gRPC метод
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u, err := h.svc.Update(uint(req.Id), req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: mapUser(u),
	}, nil
}

// DeleteUser gRPC метод
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.Delete(uint(req.Id))
	return &userpb.DeleteUserResponse{}, err
}
