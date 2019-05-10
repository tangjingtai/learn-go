package main

import (
	"context"
	"gopkg.in/mgo.v2"
	vesselPb "shippy/vessel-service/proto/vessel"
)

// 实现微服务的服务端
type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

func (h *handler) FindAvailable(ctx context.Context, req *vesselPb.Specification, resp *vesselPb.Response) error {
	repo := h.GetRepo()
	defer repo.Close()
	v, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}

func (h *handler) Create(ctx context.Context, req *vesselPb.Vessel, resp *vesselPb.Response) error {
	repo := h.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}
