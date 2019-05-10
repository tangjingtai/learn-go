package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const (
	DB_NAME           = "vessels"
	VESSEL_COLLECTION = "vessels"
)

type Repository interface {
	FindAvailable(*vesselPb.Specification) (*vesselPb.Vessel, error)
	Create(*vesselPb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

// 接口实现
func (repo *VesselRepository) FindAvailable(spec *vesselPb.Specification) (*vesselPb.Vessel, error) {
	// 选择最近一条容量、载重都符合的货轮
	var v *vesselPb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$bte": spec.MaxWeight},
	}).One(&v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 完成与数据库交互的创建动作
func (repo *VesselRepository) Create(v *vesselPb.Vessel) error {
	return repo.collection().Insert(v)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(VESSEL_COLLECTION)
}
