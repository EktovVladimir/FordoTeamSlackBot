package db

type UniqId = int64

type HasUniqId interface {
	GetId() UniqId
	SetId(UniqId)
}

type UniqFields struct {
	Id UniqId `json:"id" bson:"_id" bun:",pk,autoincrement"`
}

func (u *UniqFields) GetId() UniqId {
	return u.Id
}

func (u *UniqFields) SetId(id UniqId) {
	u.Id = id
}

func WithId(id UniqId) UniqFields {
	return UniqFields{id}
}
