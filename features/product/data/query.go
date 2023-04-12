package data

import (
	"errors"
	"kstyleAPI/features/product"
	"log"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productQuery{
		db: db,
	}
}

func (pq *productQuery) checkDuplicate(productName string) bool {
	p := Product{}
	pq.db.Where("name_product = ?", productName).First(&p)

	return p.ID != 0
}

func (pq *productQuery) checkProduct(id uint) bool {
	p := Product{}
	pq.db.Where("id = ?", id).First(&p)

	return p.ID != 0
}

func (pq *productQuery) Insert(newProduct product.Core) (product.Core, error) {
	if pq.checkDuplicate(newProduct.NameProduct) {
		log.Println("duplicated product")
		return product.Core{}, errors.New("duplicated")
	}

	cnv := CoreToData(newProduct)
	if err := pq.db.Create(&cnv).Error; err != nil {
		log.Println("error insert new product: ", err)
		return product.Core{}, err
	}

	newProduct.ID = cnv.ID

	return newProduct, nil
}

func (pq *productQuery) Update(updProduct product.Core) (product.Core, error) {
	if !pq.checkProduct(updProduct.ID) {
		log.Println("product not found")
		return product.Core{}, errors.New("product not found")
	}

	cnv := CoreToData(updProduct)
	if err := pq.db.Updates(&cnv).Error; err != nil {
		log.Println("error update product: ", err)
		return product.Core{}, err
	}

	updProduct.ID = cnv.ID

	return updProduct, nil
}

func (pq *productQuery) Delete(id uint) error {
	if !pq.checkProduct(id) {
		log.Println("product not found")
		return errors.New("product not found")
	}

	prod := Product{}
	err := pq.db.Model(&prod).Delete("id_product = ?", id).Error
	if err != nil {
		log.Println("error delete product: ", err)
		return err
	}
	return nil
}

func (pq *productQuery) GetProducts() ([]product.Core, error) {
	products := []product.Core{}
	err := pq.db.Raw("SELECT * FROM products WHERE deleted_at IS NULL").Scan(&products).Error
	if err != nil {
		log.Println("error get products: ", err)
	}

	return products, nil
}

func (pq *productQuery) GetProductById(id uint) (product.ProductRes, error) {
	p := Product{}
	pq.db.First(&p, id)

	res := product.ProductRes{
		ID:          p.ID,
		NameProduct: p.NameProduct,
		Price:       p.Price,
	}

	reviews := []product.ReviewRes{}
	pq.db.Raw("SELECT username, gender, skintype skin_type, skincolor skin_color, desc_review, COUNT(lr.id_member) likes FROM review_products rp LEFT JOIN members m ON rp.id_member = m.id LEFT JOIN like_reviews lr ON lr.id_review = rp.id WHERE rp.id_product = ? GROUP BY rp.id;", p.ID).Scan(&reviews)

	res.Reviews = reviews

	return res, nil
}
