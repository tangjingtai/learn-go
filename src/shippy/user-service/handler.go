package main

import (
	"context"
	userPb "shippy/user-service/proto/user"
)

type handler struct {
	repo Repository
}

func (h *handler) Create(ctx context.Context, req *userPb.User, resp *userPb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return nil
	}
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *userPb.User, resp *userPb.Response) error {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *userPb.Request, resp *userPb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *userPb.User, resp *userPb.Token) error {
	_, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	resp.Token = "`x_2nam"
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *userPb.Token, resp *userPb.Token) error {
	return nil
}
