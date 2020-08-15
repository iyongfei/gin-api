package origin_price_model

import (
	"gin-frame/models/base"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

type OriginPrice struct {
	//gorm.Model
	Id            int `gorm:"primary_key"`
	Customer_id   int
	Province_id   int
	City_id       int
	County_id     int
	Location_id   int
	Product_id    int
	Breed_id      int
	Point_key     string
	Day_time      string
	Price_list    string
	Desc_list     string
	Status        int
	Created_time  int
	Updated_time  int
	Refuse_reason string
	Is_sync       int
}

func (OriginPrice) TableName() string {
	return "origin_price"
}

type OriginPriceModel struct {
	base.BaseModel
}

var onceOriginPriceModel sync.Once
var originPriceModel *OriginPriceModel

func NewOriginPriceModel() *OriginPriceModel {
	onceOriginPriceModel.Do(func() {
		originPriceModel = &OriginPriceModel{}
		originPriceModel.GetConn("hangqing")
		log.Printf("new origin_price_model")
	})

	return originPriceModel
}

func (instance *OriginPriceModel) GetFirst(c *gin.Context) []OriginPrice {
	originPrices := []OriginPrice{}
	orm := instance.Db.SlaveOrm()

	instance.Start(c)
	dbRes := orm.First(&originPrices)
	sql := "select * from test where id = 1"
	instance.End(sql)

	instance.CheckRes(dbRes)
	return originPrices
}
