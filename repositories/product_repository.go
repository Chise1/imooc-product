package repositories

import (
	"Chise1/imooc-product/datamodels"
	"Chise1/imooc-product/db"
	"gorm.io/gorm"
)

//第一步，先开发对应的接口
//第二步，实现定义的接口
type IProduct interface {
	//连接数据
	Insert(product *datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]datamodels.Product, error)
}
type ProductManager struct {
	mysqlConn *gorm.DB
}

func NewProductManager() IProduct {
	res := &ProductManager{}
	var err error
	res.mysqlConn, err = db.GetDB()
	if err == nil {
		res.mysqlConn.AutoMigrate(&datamodels.Product{})
	} else {
		panic(err)
	}
	return res
}

//插入
func (p *ProductManager) Insert(product *datamodels.Product) (productId int64, err error) {
	//1.判断连接是否存在
	tx := p.mysqlConn.Create(&product)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return product.Id, nil
}

//商品的删除
func (p *ProductManager) Delete(productId int64) bool {
	tx := p.mysqlConn.Delete(&datamodels.Product{}, productId)
	if tx.Error != nil {
		return false
	}
	return true
}

//商品的更新
func (p *ProductManager) Update(product *datamodels.Product) error {
	//1.判断连接是否存在
	//tx := p.mysqlConn.Model(&product).Select("productName", "productImage", "productUrl").Updates(product)
	tx := p.mysqlConn.Model(&product).Select("ProductName", "ProductNum", "ProductImage", "ProductUrl").Updates(product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

//根据商品ID查询商品
func (p *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	//1.判断连接是否存在
	var pd *datamodels.Product
	tx := p.mysqlConn.First(&pd, productID)
	return pd, tx.Error
}

//获取所有商品
func (p *ProductManager) SelectAll() (productArray []datamodels.Product, errProduct error) {
	var products []datamodels.Product
	tx := p.mysqlConn.Find(&products)
	return products, tx.Error
}
