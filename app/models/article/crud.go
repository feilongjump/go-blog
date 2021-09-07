package article

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/model"
	"go-blog/pkg/pagination"
	"go-blog/pkg/route"
	"go-blog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToInt(idStr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 获取视图数据
	viewData := _pager.Paging()

	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

func (article *Article) Create() (err error) {

	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func (article Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (article Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func GetByUserID(uid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 初始化分页实例
	db := model.DB.Model(Article{}).Where("user_id = ?", uid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("users.show", "id", uid), perPage)

	// 获取视图数据
	viewData := _pager.Paging()

	// 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 初始化分页实例
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", cid), perPage)

	// 获取视图数据
	viewData := _pager.Paging()

	// 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}
