package Models
import (
	"fmt"
	"github.com/freshers-bootcamp/Day4/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllProducts(product *[]Product ) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product ) (err error) {
	if err = Config.DB.Model(&product).Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product , id string) (err error) {
	if err = Config.DB.Model(&product).Where("product_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func PatchProduct(product *Product , id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product , id string) (err error) {
	Config.DB.Model(&product).Where("product_id = ?", id).Delete(product)
	return nil
}