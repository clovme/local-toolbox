package persistence

import (
	"errors"
	"fmt"
	"strconv"
	"toolbox/internal/infrastructure/libs"
	"toolbox/internal/infrastructure/query"
	"toolbox/internal/models"
	"toolbox/internal/schema/dto/category"
	categoryVO "toolbox/internal/schema/vo/category"
	"toolbox/pkg/logger/log"

	"gorm.io/gorm"
)

type ApiCategoryRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *ApiCategoryRepository) GetCategoryList() ([]*categoryVO.ApiCategoryVO, error) {
	c := r.Q.Category
	a := r.Q.Article
	var results []*categoryVO.ApiCategoryVO
	err := c.Order(c.Sort.Asc()).Scan(&results)
	if err != nil {
		return nil, err
	}

	var articleGroup []categoryVO.CountResultVO
	if err = a.Select(a.CategoryID, a.CategoryID.Count().As("article_count")).Group(a.CategoryID).Scan(&articleGroup); err != nil {
		log.Error().Err(err).Msg("文章分类查询失败")
	}
	countMap := make(map[int64]int64, len(articleGroup))
	var total int64
	for _, c := range articleGroup {
		countMap[c.CategoryID] = c.ArticleCount
		total += c.ArticleCount
	}

	// 组装结果
	for _, res := range results {
		if res.Title == "全部文章" {
			res.ArticleCount = total // “全部分类” 的总数
		} else {
			id, _ := strconv.ParseInt(res.ID, 10, 64)
			res.ArticleCount = countMap[id]
		}
	}

	return results, err
}

// DeleteCategory
//
// 参数:
//   - id int46
//
// 返回:
//   - int64, error
//
// 说明:
//   - 删除当前以及子节点，需要根据当前节点查找子节点
//   - 把子节点相关的文章设置成默认分类
func (r *ApiCategoryRepository) DeleteCategory(id int64) (int64, int64, error) {
	c := r.Q.Category
	var resultList []categoryVO.ApiCategoryVO
	if err := c.Scan(&resultList); err != nil {
		return 0, 0, err
	}

	children, err := libs.GetCategoryWithChildren(resultList, id)
	if err != nil {
		return 0, 0, err
	}

	var errRowsAffected int64
	var updateRowsAffected int64

	err = r.Q.Transaction(func(tx *query.Query) error {
		// 删除分类
		info, err := c.Where(c.ID.NotIn(1, 2), c.ID.In(children...)).Delete()
		if err != nil {
			return err
		}
		categoryID, err := c.Select(c.ID).Where(c.Name.Eq("default")).First()
		if err != nil {
			return err
		}
		// 更新文章
		update, err := r.Q.Article.Where(r.Q.Article.CategoryID.In(children...)).Update(r.Q.Article.CategoryID, categoryID.ID)
		if err != nil {
			return err
		}
		errRowsAffected = info.RowsAffected
		updateRowsAffected = update.RowsAffected
		return err
	})

	return errRowsAffected, updateRowsAffected, err
}

func (r *ApiCategoryRepository) AddCategory(dto category.ApiCategoryAddDTO) error {
	c := r.Q.Category

	take, err := c.Select(c.ID).Where(c.Title.Eq(dto.Title), c.Pid.Eq(*dto.Pid)).Count()
	if err != nil {
		return err
	}

	if take > 0 {
		return errors.New("该分类已存在，请检查重试！")
	}

	cm := models.Category{
		Name:        "custom",
		Title:       dto.Title,
		Pid:         *dto.Pid,
		Description: dto.Description,
		Sort:        dto.Sort,
	}

	if err := c.Create(&cm); err != nil {
		return err
	}
	return nil
}

func (r *ApiCategoryRepository) UpdateCategory(dto category.ApiCategoryAddDTO) error {
	c := r.Q.Category

	// 1️⃣ 检查不能分配给自己或子分类
	resultList := make([]categoryVO.ApiCategoryVO, 0)
	if err := c.Scan(&resultList); err != nil {
		return err
	}

	children, err := libs.GetCategoryWithChildren(resultList, *dto.ID)
	if err != nil {
		return err
	}

	// 如果 pid 在自己或子分类列表里
	for _, id := range children {
		if id == *dto.Pid {
			return fmt.Errorf("分配错误，不能把 [%s] 分配给自己或下属分类", dto.Title)
		}
	}

	// 2️⃣ 检查同级是否存在重名分类
	take, err := c.Where(c.Pid.Eq(*dto.Pid), c.Title.Eq(dto.Title), c.ID.Neq(*dto.ID)).Count()
	if err != nil {
		return err
	}
	if take > 0 {
		return errors.New("该分类已存在，请检查重试！")
	}

	// 3️⃣ 更新分类
	info, err := c.Where(c.ID.Eq(*dto.ID)).UpdateSimple(c.Title.Value(dto.Title), c.Pid.Value(*dto.Pid), c.Description.Value(dto.Description), c.Sort.Value(dto.Sort))
	if err != nil {
		return err
	}

	return info.Error
}
