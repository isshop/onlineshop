package models

import (
	"github.com/astaxie/beego/orm"

	//"time"
)
type User struct {
	Id          int
	Name        string

	//设置表关系
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	//设置表关系
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string

	//设置表关系
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

////购物车
//type ShopCart struct {
//	Id int `orm:"pk;auto"`
//	Name int
//	Goods []*GoodDetail `orm:"null;rel(m2m)"`
//	User *User `orm:"null;rel(one);on_delete(cascade)"`
//}

////收藏历史
//type History struct {
//	Id int `orm:"pk;auto"`
//	Goods []*Good `orm:"rel(m2m)"`
//	User *User `orm:"rel(one);on_delete(cascade)"`
//}

////商品
//type Good struct {
//	Id int `orm:"pk;auto"`
//	Name string
//	GoodDetail *GoodDetail `orm:"reverse(one)"`
//	Description string
//	// 这个价格是不准确的， 具体应该以GoodDetail里的价格为准
//	Price float32
//	//关注数量
//	Follows int64 `orm:default(0)`
//	//分类
//	TypeCode int32
//	//库存剩余
//	Stock int64
//	//评论
//	Comments []*Comment `orm:"rel(m2m)"`
//	//型号s,m,
//	version int
//
//}
//
//
//type GoodDetail struct {
//	Id int `orm:"pk;auto"`
//	Good *Good `orm:"rel(one)"`
//
//	// 颜色分类
//	ColorType string `orm:"null"`
//	// 尺码大小
//	Size string `orm:"null"`
//	//扩展字段
//	Other string `orm:"null"`
//	// 同一个商品下不同颜色和尺码属于不同的version
//	Version int
//	Price float32
//	//图片
//	Image string
//	// 库存
//	Stock int64
//	//订单
//	Orders []*Orders `orm:"reverse(many)"`
//	// 关注数量
//	Follows int64 `orm:"default(0)"`
//	// 分类
//	TypeCode string
//}
//
////评论
//type Comment struct {
//	Id int `orm:"pk;auto"`
//	User *User `orm:"rel(fk)"`
//	Content string
//	//是直接评论（1）还是回复(2)
//	Type_ int `orm:"default(1)"`
//	//创建时间
//	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
//	//上一条评论
//	Comment *Comment `orm:"null;rel(one)"`
//}
//
////秒杀活动
//type SecKillGoods struct {
//	Id int `orm:"pk;auto"`
//	Good *GoodDetail `orm:"rel(one)"`
//	// 秒杀活动开始的时间
//	CreatedTime time.Time `orm:"type(datetime)"`
//	// 截止时间
//	EndTime time.Time `orm:"type(datetime)"`
//	// 优惠之后的价格
//	DiscountPrice float32
//	// 总的抢购数量
//	TotalNum int64
//}
//
///*
//订单
//支付功能 省略， 提供接口，可以自行扩展实现
//订单取消功能 省略，可自行扩展实现
//*/
//type Orders struct {
//	Id int `orm:"pk;auto"`
//	User *User `orm:"rel(fk)"`
//	Good []*GoodDetail `orm:"reverse(many)"`
//	// 0刚下单但是未支付， 1表示已支付， 2订单完成， 其他更多状态自行扩展
//	Status int `orm:"default(0)"`
//	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
//}

func init(){
	//orm.RegisterModel(new(User), new(ShopCart), new(History), new(Good), new(Comment), new(GoodDetail), new(Orders),
	//	new(SecKillGoods))
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
}