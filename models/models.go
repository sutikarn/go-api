package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Code     string `gorm:"type:varchar(255);not null;unique" json:"code"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

type Profile struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);not null" json:"firstname"`
	LastName  string `gorm:"type:varchar(255);not null" json:"lastname"`
	Mobile    string `gorm:"type:varchar(10)" json:"mobile"`
	Sex       string `gorm:"type:varchar(10); default:m" json:"sex"`
	Status    string `gorm:"type:varchar(10); default:a" json:"status" `
	Image     string `gorm:"type:text" json:"image"`
	UserID    uint   `gorm:"not null" json:"userid"` // FK referencing User.ID
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Address struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);not null" json:"firstname"`
	LastName  string `gorm:"type:varchar(255);not null" json:"lastname"`
	Mobile    string `gorm:"type:varchar(10)" json:"mobile"`
	Address   string `gorm:"type:varchar(1024)" json:"address"`
	Type      string `gorm:"type:text; default:1" json:"type"`
	UserID    uint   `gorm:"not null" json:"userid"` // FK referencing User.ID
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique" json:"name"`
	Description string `gorm:"type:varchar(1024)" json:"description"`
	Image       string `gorm:"type:text" json:"image"`
}

type Mall struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:varchar(1024)" json:"description"`
	Image       string `gorm:"type:text" json:"image"`
}

type Product struct {
	gorm.Model
	Code        string   `gorm:"type:varchar(255);not null;unique" json:"code"`
	Name        string   `gorm:"type:varchar(255);not null;unique" json:"name"`
	Description string   `gorm:"type:varchar(1024);not null" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	Rating      int      `gorm:"not null" json:"rating"`
	Image       string   `gorm:"type:text" json:"image"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`                                          // FK referencing Category.ID
	Category    Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Corrected to use Category
	MallID      uint     `gorm:"not null" json:"mall_id"`                                              // FK referencing Mall.ID
	Mall        Mall     `gorm:"foreignKey:MallID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`     // Corrected to use Mall
}

type Order struct {
	gorm.Model
	Price     float64 `gorm:"not null" json:"price"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	ProductID uint    `gorm:"not null" json:"productID"`                                           // FK referencing Product.ID
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Corrected to use Product
	UserID    uint    `gorm:"not null" json:"userid"`                                              // FK referencing User.ID
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Banner struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:varchar(1024)" json:"description"`
	Image       string `gorm:"type:text" json:"image"`
}

type Cart struct {
	gorm.Model
	Price     float64 `gorm:"not null" json:"price"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	ProductID uint    `gorm:"not null" json:"productID"`                                           // FK referencing Product.ID
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Corrected to use Product
	UserID    uint    `gorm:"not null" json:"userid"`                                              // FK referencing User.ID
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Favorite struct {
	gorm.Model
	Price     float64 `gorm:"not null" json:"price"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	ProductID uint    `gorm:"not null" json:"productID"`                                           // FK referencing Product.ID
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Corrected to use Product
	UserID    uint    `gorm:"not null" json:"userid"`                                              // FK referencing User.ID
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
