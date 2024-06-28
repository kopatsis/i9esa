package datatypes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Parent       string             `bson:"parent"`
	MinLevel     float32            `bson:"minlevel"`
	MaxLevel     float32            `bson:"maxlevel"`
	MinReps      int                `bson:"minreps"`
	PlyoRating   int                `bson:"plyorating"`
	StartQuality float32            `bson:"startquality"`
	BodyParts    []int              `bson:"bodyparts"`
	RepVars      [3]float32         `bson:"repvars"`
	InSplits     bool               `bson:"insplits"`
	InPairs      bool               `bson:"inpairs"`
	UnderCombos  bool               `bson:"undercombos"`
	CardioRating float32            `bson:"cardiorating"`
	PushupType   string             `bson:"pushuptype"`
	GeneralType  []string           `bson:"generaltype"`
	SinglesGroup int                `bson:"singlesgroup"`
}

type Stretch struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	MinLevel     float32            `bson:"minlevel"`
	Status       string             `bson:"status"`
	BodyParts    []int              `bson:"bodyparts"`
	InPairs      bool               `bson:"inpairs"`
	DynamicPairs []string           `bson:"dynamicpairs"`
	Weight       float32            `bson:"weight"`
	ReqGroup     int                `bson:"reqgroup"`
}

type TypeMatrix struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Matrix [11][11]float32    `bson:"matrix"`
}
