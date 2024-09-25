package loaddata

import (
	"encoding/json"
	"log"
	"nack/models"

	"gorm.io/gorm"
)

func LoadData(db *gorm.DB) {
    loadDataBanner(db)
	loadCategory(db) // Ensure categories are loaded first
	loadShoppingMall(db)
	loadProduct(db) // Then load products
}

func loadShoppingMall(db *gorm.DB) {
	jsonData := `[
    {
        "id": 1,
        "name": "Shopping Mall 1",
        "description": "Shopping Mall 1",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S001.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=fb900da5916325c7ae7a41f18188a6658ed86f2bd75c458536f5e9274f12a1f6"
    },
    {
        "id": 4,
        "name": "Shopping Mall 2",
        "description": "Shopping Mall 2",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S002.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=13e357c679188361751be2710ef15cc3475b6e9dc518a66bc6aa88d608c0ed6a"
    },
    {
        "id": 7,
        "name": "Shopping Mall 3",
        "description": "Shopping Mall 3",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S003.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=f3141e77cc67c360ac93ad0c84f577328e8d513881f880b4d87902573bcca4b6"
    },
    {
        "id": 10,
        "name": "Shopping Mall 4",
        "description": "Shopping Mall 4",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S004.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=4a4674bcd7935913025be163c9c1c3ef5c695e03b1563c028e56591cdf097360"
    },
    {
        "id": 13,
        "name": "Shopping Mall 5",
        "description": "Shopping Mall 5",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S005.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=5398b600db3ad6526a1c7647589dbcec3711c2bab52f89ebe9c18d9ea2ea78ee"
    },
    {
        "id": 16,
        "name": "Shopping Mall 6",
        "description": "Shopping Mall 6",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S006.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=c7573546e8d7f76e33b13bfecd8b0d002c30baeaedecd7f538f76b89680fe36f"
    },
    {
        "id": 19,
        "name": "Shopping Mall 7",
        "description": "Shopping Mall 7",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S007.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=50443a97af2f19f823e70e560d78073772900cd181909e891dda7bb9b7dc36a3"
    },
    {
        "id": 22,
        "name": "Shopping Mall 8",
        "description": "Shopping Mall 8",
        "image": "https://s3.eco-deals.com/shopping-bucket/shopping-mall/S008.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080758Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=3362f26aa2caa5fab6e80be793a43f0c7369fc44c49868da33037e412830e859"
    }
]` // Your JSON data

	var malls []model.Mall
	if err := json.Unmarshal([]byte(jsonData), &malls); err != nil {
		log.Fatalf("Error unmarshalling shopping malls JSON: %v", err)
	}

	// Insert data into PostgreSQL
	for _, mall := range malls { // Changed the loop variable to 'mall'
		if err := db.Create(&mall).Error; err != nil {
			log.Fatalf("Error inserting shopping mall record: %v", err)
		}
	}

	log.Println("Data Shopping Malls inserted successfully!")
}

func loadProduct(db *gorm.DB) {
	jsonData := `[
    {
        "id": 4,
        "code": "P001",
        "name": "Bite Of Wild",
        "description": "อาหารแมว 1KG Grain Free โปรตีน 42% 3 เนื้อฟรีซดราย ไก่ปลาแซลมอนเนื้อปลา สำหรับแมวทุกช่ว",
        "price": 1000.5,
        "rating": 5,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P001.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=3c52c93cb8f17d344f4a1837758e3d5b08c4302b0e8f4e1518eb2e3d0ecc8b7e",
        "category_id": 1,
        "mall_id": 1
    },
    {
        "id": 7,
        "code": "P002",
        "name": "Pramy Supreme",
        "description": "อาหารเม็ดแมว Mother/Kitten/Adult/Indoor/Skin&Coat สำหรับแมวทุกช่วงวัย",
        "price": 100.5,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P002.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=4f4d6467bb416f771277293d22f96b8c92ce3e359186c91d91855143466c9451",
        "category_id": 4,
        "mall_id": 4
    },
    {
        "id": 10,
        "code": "P003",
        "name": "WHISKAS",
        "description": "วิสกัส อาหารแมว ชนิดแห้ง แบบเม็ด – อาหารแมว สูตรแมวโต, 7 กก. สำหรับแมวโตอายุ 1 ปีขึ้นไป",
        "price": 1001,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P003.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=4a0d168f51390ff2de8b4eea17d27233d406903162a96c645d20efaf196b3acf",
        "category_id": 7,
        "mall_id": 7
    },
    {
        "id": 13,
        "code": "P004",
        "name": "Petheria",
        "description": "อาหารแมว เพ็ทเทอเรีย ครบทุกสูตร ลดการเกินก้อนขน ไม่เค็ม ขนาด 1.5 กิโล",
        "price": 100,
        "rating": 3,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P004.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=f661b51b60e7830ae6a7a98a7edac504eb22443e436ef1c0c9b78e5d44a892e0",
        "category_id": 10,
        "mall_id": 10
    },
    {
        "id": 16,
        "code": "P005",
        "name": "Whiskas Junior",
        "description": "อาหารเม็ดสำหรับแมวอายุ 2-12 เดือน รส Ocean Fish (ปลาทะเล) 1.1Kg.",
        "price": 199,
        "rating": 5,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P005.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=a078b367935b154227a7199798a5e0ebf943bfd3a245a0a904913277d7bee117",
        "category_id": 1,
        "mall_id": 13
    },
    {
        "id": 19,
        "code": "P006",
        "name": "Purina One",
        "description": "อาหารแมวพรีเมี่ยม เพียวริน่า วัน ขนาด 1.2 kg",
        "price": 299,
        "rating": 5,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P006.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=28dacc77b04f9713a50e0e42eb87323ee982bbe45f70e80a098753afdd7308df",
        "category_id": 4,
        "mall_id": 16
    },
    {
        "id": 22,
        "code": "P007",
        "name": "CatHoliday",
        "description": "เพ็ทเทอเรีย petheria อาหารแมวแบบเม็ด ขนาด 1.5 กิโล",
        "price": 499.5,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P007.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=1f77dccad01362d71398491482cc4b7b913c7fe887a1512965924bdc18334a29",
        "category_id": 7,
        "mall_id": 19
    },
    {
        "id": 25,
        "code": "P008",
        "name": "Royal Canin",
        "description": "Hair & Skin Care รอยัลคานิน อาหารแมวโต สูตร ดูแลผิวหนัง เส้นขน อายุ 1 ปีขึ้นไป 2kg",
        "price": 399,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P008.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=f72e6da85b1976d973b308dfc8f2015f19cc26fcd20b72ecfba9c21f097cdace",
        "category_id": 10,
        "mall_id": 22
    },
    {
        "id": 28,
        "code": "P009",
        "name": "PURINA ONE",
        "description": "เพียวริน่าวัน อาหารแมว ถุงขนาด 1.2kg ทุกสูตร เกรดพรีเมี่ยม อุดมด้วยสารอาหารที่เหมาะสมกับแมวในแต่ละช่วงวัย",
        "price": 999,
        "rating": 5,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P009.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=629759aacde4ca45cd162a96478e53b81df20f6c4f1910db94ec4d1e7811c0dc",
        "category_id": 7,
        "mall_id": 22
    },
    {
        "id": 31,
        "code": "P010",
        "name": "Pramy",
        "description": "พรามี่ อาหารแมวเปียก สูตรสำหรับลูกแมว แมวโต แมวสูงวัย ขนาด 70g.ไม่คละรส",
        "price": 1999,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P010.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=213f36bc72ad7f9b857bd4e223e24ebda1a14be4c61d38f5741599e4b7ce9395",
        "category_id": 10,
        "mall_id": 19
    },
    {
        "id": 34,
        "code": "P011",
        "name": "BUZZ BEYOND",
        "description": "บัซซ์ อาหารแมว พรีเมียม สูตรเกรนฟรี อาหารชนิดเม็ด อาหารเม็ดแมว 1/1.2 Kg.",
        "price": 789.5,
        "rating": 5,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P011.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=52f2094a399169eeb3f1c495b5ee857be01d95a50776ed82fa0d405f630e9af2",
        "category_id": 1,
        "mall_id": 19
    },
    {
        "id": 37,
        "code": "P012",
        "name": "Kaniva",
        "description": "คานิว่า​ อาหารเม็ดสำหรับเเมว ทานยาก​ ไม่เค็ม​ อึไม่เหม็น ขนาด 1.3 - 1.5 กก.",
        "price": 1189.5,
        "rating": 4,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P012.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=297980c5c5776a19ff94c5b9a7ab7e78897ee1005b310b1c6a08bc19bb791891",
        "category_id": 10,
        "mall_id": 22
    },
    {
        "id": 40,
        "code": "P013",
        "name": "Buzz Netura",
        "description": "อาหารแมว Holistic บัซซ์ สูตร โฮลิสติก เกรนฟรี แมวโต ลูกแมว ทุกช่วงวัย ไก่ / แซลมอน / Protein X ขนาด 1 กิโล",
        "price": 1189.5,
        "rating": 3,
        "image": "https://s3.eco-deals.com/shopping-bucket/products/P013.webp?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080813Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=e51aad884b65ef5318058f2a14e80b6bad03040a6de8e9f4332ffde8923247c9",
        "category_id": 4,
        "mall_id": 22
    }
]` // Your JSON data

	var products []model.Product
	if err := json.Unmarshal([]byte(jsonData), &products); err != nil {
		log.Fatalf("Error unmarshalling products JSON: %v", err)
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("Error inserting product record: %v", err)
		}
	}

	log.Println("Data Products inserted successfully!")
}

func loadCategory(db *gorm.DB) {
	jsonData := `[
    {
        "id": 1,
        "name": "Animal food",
        "description": "Animal food",
        "image": "https://s3.eco-deals.com/shopping-bucket/category/Animal_food.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080825Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=e04a60464f2daed42b3610821db05a65bbef579a54cc244dd3564ca383cd5d31"
    },
    {
        "id": 4,
        "name": "Pet supplies",
        "description": "Pet supplies",
        "image": "https://s3.eco-deals.com/shopping-bucket/category/pet_supplies.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080825Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=b6de45ef7906df6c0aa6b2004202a6f6c4f2089d6313e1c146ba4981f9874bea"
    },
    {
        "id": 7,
        "name": "Clothes and accessories",
        "description": "Clothes and accessories",
        "image": "https://s3.eco-deals.com/shopping-bucket/category/Clothes_and_accessories.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080825Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=2301119d1326278bfd3f6c9ee6cb1b59a0ad59ad0bfd702506c1639e197e3a62"
    },
    {
        "id": 10,
        "name": "Cleaning equipment",
        "description": "Cleaning equipment",
        "image": "https://s3.eco-deals.com/shopping-bucket/category/soap-bottle.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080825Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=46c7b04c63c1b2ced69aba5203fc3a87ac1f4e261f1ba906a0146bfd4feeb357"
    }
]` // Your JSON data

	var categories []model.Category
	if err := json.Unmarshal([]byte(jsonData), &categories); err != nil {
		log.Fatalf("Error unmarshalling categories JSON: %v", err)
	}

	// Insert data into PostgreSQL
	for _, category := range categories { // Changed the loop variable to 'category'
		if err := db.Create(&category).Error; err != nil {
			log.Fatalf("Error inserting category record: %v", err)
		}
	}

	log.Println("Data Categories inserted successfully!")
}

func loadDataBanner(db *gorm.DB) {
	jsonData := `[
    {
        "id": 1,
        "name": "Banner1",
        "description": "Banner1",
        "image": "https://s3.eco-deals.com/shopping-bucket/banner/B001.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080837Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=8af2ed0fda808335e2a8d591384f5f813a90c0873fd612ff6437beae68d05ca0"
    },
    {
        "id": 4,
        "name": "Banner2",
        "description": "Banner2",
        "image": "https://s3.eco-deals.com/shopping-bucket/banner/B002.avif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=9E9QBO149ACWGBRO6VELIWF9%2F20240925%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240925T080837Z&X-Amz-Expires=10000&X-Amz-SignedHeaders=host&response-content-disposition=inline%3B&X-Amz-Signature=c1ebbba7b2deca81bc86b28ce77f43f0ebde66221b86b0fbb752798d58afd406"
    }
]` // Your JSON data

	var banners []model.Banner
	if err := json.Unmarshal([]byte(jsonData), &banners); err != nil {
		log.Fatalf("Error unmarshalling banners JSON: %v", err)
	}

	// Insert data into PostgreSQL
	for _, banner := range banners { // Changed the loop variable to 'banner'
		if err := db.Create(&banner).Error; err != nil {
			log.Fatalf("Error inserting banner record: %v", err)
		}
	}

	log.Println("Data Banners inserted successfully!")
}
