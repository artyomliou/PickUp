package models

type CartItem struct {
	ID        uint `form:"id" binding:"optional" valid:"int"`
	StoreId   uint `form:"storeId" binding:"required" valid:"int"`
	ProductId uint `form:"productId" binding:"required" valid:"int"`
	Amount    uint `form:"amount" binding:"required" valid:"int"`
	Selects   []SelectAnswer
	Customs   []CustomAnswer
}
