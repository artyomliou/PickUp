package command

import (
	_ "pick-up/backend/configs"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/models"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SeedCommand struct{}

func (cmd SeedCommand) Run() error {

	_, err := models.SeedUser()
	if err != nil {
		return err
	}

	store, err := models.NewStore(&models.Store{
		Name:     "60藍",
		Pic:      "https://picsum.photos/id/292/3852/2556",
		Status:   models.StoreStatus(models.StoreStatusOpened),
		OpenedAt: "09:00",
		ClosedAt: "18:00",
	})
	if err != nil {
		return err
	}

	iceQuestion, _, err := models.NewSelectQuestion(&models.SelectQuestion{
		StoreId:    store.ID,
		Name:       "冰量",
		IsRequired: true,
		AtMost:     1,
	}, []string{
		"正常冰",
		"少冰",
		"微冰",
		"去冰",
		"完全去冰",
	})
	if err != nil {
		return err
	}

	sugarQuestion, _, err := models.NewSelectQuestion(&models.SelectQuestion{
		StoreId:    store.ID,
		Name:       "糖量",
		IsRequired: true,
		AtMost:     1,
	}, []string{
		"全糖",
		"少糖",
		"半糖",
		"微糖",
		"一分糖",
		"無糖",
	})
	if err != nil {
		return err
	}

	customQuestion, err := models.NewCustomQuestion(&models.CustomQuestion{
		StoreId:     store.ID,
		Name:        "有額外需求嗎？",
		Placeholder: "請在這邊寫您的需求",
	})
	if err != nil {
		return err
	}

	teaCategory, err := models.NewCategory(store.ID, "純茶")
	if err != nil {
		return err
	}

	fruitTeaCategory, err := models.NewCategory(store.ID, "水果茶")
	if err != nil {
		return err
	}

	milkTeaCategory, err := models.NewCategory(store.ID, "鮮奶茶")
	if err != nil {
		return err
	}

	t1 := &models.Product{
		StoreId: store.ID,
		Name:    "金色烏龍",
		Price:   30,
		Intro:   "金色茶湯、醇香烏龍",
	}
	if _, err := models.NewProduct(t1); err != nil {
		return err
	}
	err = AssociateProductWithCategoriesAndQuestions(
		t1,
		[]*models.Category{teaCategory},
		[]*models.SelectQuestion{iceQuestion, sugarQuestion},
		[]*models.CustomQuestion{customQuestion},
	)
	if err != nil {
		return err
	}

	t2 := &models.Product{
		StoreId: store.ID,
		Name:    "金色烏龍鮮奶茶",
		Price:   60,
		Intro:   "金色茶湯、醇香烏龍、頂級牛奶",
	}
	if _, err := models.NewProduct(t2); err != nil {
		return err
	}
	err = AssociateProductWithCategoriesAndQuestions(
		t2,
		[]*models.Category{milkTeaCategory},
		[]*models.SelectQuestion{iceQuestion, sugarQuestion},
		[]*models.CustomQuestion{customQuestion},
	)
	if err != nil {
		return err
	}

	t3 := &models.Product{
		StoreId: store.ID,
		Name:    "水果烏龍茶",
		Price:   60,
		Intro:   "金色茶湯、醇香烏龍、台灣水果",
	}
	if _, err := models.NewProduct(t3); err != nil {
		return err
	}
	err = AssociateProductWithCategoriesAndQuestions(
		t3,
		[]*models.Category{fruitTeaCategory},
		[]*models.SelectQuestion{iceQuestion, sugarQuestion},
		[]*models.CustomQuestion{customQuestion},
	)
	if err != nil {
		return err
	}

	return nil
}

func AssociateProductWithCategoriesAndQuestions(p *models.Product, categories []*models.Category, selectQuestions []*models.SelectQuestion, customQuestions []*models.CustomQuestion) error {
	if err := db.Conn().Model(p).Association("Categories").Append(categories); err != nil {
		return err
	}
	if err := db.Conn().Model(p).Association("SelectQuestions").Append(selectQuestions); err != nil {
		return err
	}
	if err := db.Conn().Model(p).Association("CustomQuestions").Append(customQuestions); err != nil {
		return err
	}
	return nil
}
