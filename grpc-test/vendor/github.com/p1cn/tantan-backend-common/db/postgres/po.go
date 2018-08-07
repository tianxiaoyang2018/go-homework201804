package postgres

import pg "gopkg.in/pg.v3"

type IdRowPO struct {
	Id string `pg:"id"`
}

type IdRowsPO struct {
	C []IdRowPO
}

var _ pg.Collection = &IdRowsPO{}

func (self *IdRowsPO) NewRecord() interface{} {
	self.C = append(self.C, IdRowPO{})
	return &self.C[len(self.C)-1]
}

func (self *IdRowsPO) Ids() []string {
	ids := make([]string, 0)
	for _, user := range self.C {
		ids = append(ids, user.Id)
	}
	return ids
}

type IdsRowPO struct {
	Ids []string `pg:"ids"`
}

type IdsRowsPO struct {
	C []IdsRowPO
}

var _ pg.Collection = &IdsRowsPO{}

func (self *IdsRowsPO) NewRecord() interface{} {
	self.C = append(self.C, IdsRowPO{})
	return &self.C[len(self.C)-1]
}

func (self *IdsRowsPO) Ids() []string {
	ids := make([]string, 0)
	for _, user := range self.C {
		ids = append(ids, user.Ids...)
	}
	return ids
}
