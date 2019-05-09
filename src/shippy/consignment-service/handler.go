package main

import (
	"context"
	"gopkg.in/mgo.v2"
	"log"
	consPb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

// 微服务服务端 struct handler 必须实现 protobuf 中定义的 rpc 方法
// 实现方法的传参等可参考生成的 consignment.consPb.go
type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

// 从主会话中 Clone() 出新会话处理查询
func (h *handler) GetRepo() Repository {
	return &ConsignmentRepository{h.session.Clone()}
}

func (h *handler) CreateConsignment(ctx context.Context, req *consPb.Consignment, resp *consPb.Response) error {
	repo := h.GetRepo()
	defer repo.Close()

	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := h.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	// 货物被承运
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	//consignment, err := h.repo.Create(req)
	err = repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *consPb.GetRequest, resp *consPb.Response) error {
	repo := h.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}
