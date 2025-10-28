package persistence

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"toolbox/internal/core"
	"toolbox/internal/infrastructure/query"
	"toolbox/internal/models"
	articleDTO "toolbox/internal/schema/dto/article"
	articleVO "toolbox/internal/schema/vo/article"
	"toolbox/pkg/config"
	"toolbox/pkg/constants"
	"toolbox/pkg/utils"
	"toolbox/pkg/utils/file"

	"gorm.io/gorm"
)

type ApiArticleRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *ApiArticleRepository) DeleteArticle(dto articleDTO.ApiArticleUpdateDTO) error {
	a := r.Q.Article
	info, err := a.Where(a.ID.Eq(*dto.ID), a.Title.Eq(dto.Title), a.CategoryID.Eq(*dto.CategoryID), a.Tags.Eq(dto.Tags), a.Summary.Eq(dto.Summary), a.CreatedAt.Eq(*dto.CreatedAt), a.UpdatedAt.Eq(*dto.UpdatedAt)).Delete()
	if err != nil {
		return errors.New(fmt.Sprintf("删除《%s》失败，%v", dto.Title, err))
	}

	return info.Error
}

func (r *ApiArticleRepository) GetArticle(c *core.Context) (data *articleVO.ApiArticleVO, err error) {
	a := r.Q.Article
	err = a.Select(a.ID, a.Title, a.Tags, a.Summary, a.Content, a.CreatedAt, a.UpdatedAt, a.CategoryID, r.Q.Category.Title.As("category_title"), r.Q.Category.Name.As("category_name")).
		Where(a.ID.Eq(c.QueryInt64("id")), a.CategoryID.Eq(c.QueryInt64("cid")), r.Q.Category.Name.Eq(c.Query("type"))).LeftJoin(r.Q.Category, r.Q.Category.ID.EqCol(a.CategoryID)).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ApiArticleRepository) GetArticleList(c *core.Context) (data []articleVO.ApiArticleVO, err error) {
	a := r.Q.Article

	w := a.Select(a.ID, a.Title, a.Tags, a.Summary, a.CreatedAt, a.UpdatedAt, a.CategoryID, r.Q.Category.Title.As("category_title"), r.Q.Category.Name.As("category_name"))
	if c.Query("type") != "all" {
		w = w.Where(a.CategoryID.Eq(c.QueryInt64("cid")), r.Q.Category.Name.Eq(c.Query("type")))
	}
	if kw := c.Query("kw"); len(strings.TrimSpace(kw)) > 0 {
		kw = fmt.Sprintf("%%%s%%", kw)
		w = w.Where(a.Or(a.Title.Like(kw)).Or(a.Content.Like(kw)).Or(a.Tags.Like(kw)).Or(a.Summary.Like(kw)))
	}
	w = w.LeftJoin(r.Q.Category, r.Q.Category.ID.EqCol(a.CategoryID))
	if sort := c.Query("sort"); sort != "" && sort != "updatedAt" {
		if sort == "title" {
			w = w.Order(a.Title.Asc(), a.ID.Desc())
		} else {
			w = w.Order(a.CreatedAt.Desc(), a.ID.Desc())
		}
	} else {
		w = w.Order(a.UpdatedAt.Desc(), a.ID.Desc())
	}
	if err = w.Scan(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ApiArticleRepository) AddArticle(dto articleDTO.ApiArticleDTO) (int64, error) {
	a := r.Q.Article

	// 查询分类
	if dto.CategoryName == "all" {
		category, err := r.Q.Category.Select(r.Q.Category.ID).Where(r.Q.Category.Name.Eq("default")).First()
		if err != nil {
			return 0, errors.New(fmt.Sprintf("分类查询失败，%v", err))
		}
		dto.CategoryID = &category.ID
	}

	count, err := a.Where(a.Title.Eq(dto.Title), a.CategoryID.Eq(*dto.CategoryID)).Count()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("[%s]查询失败，%v", dto.Title, err))
	}
	if count > 0 {
		return 0, errors.New(fmt.Sprintf("该分类，该文章[%s]已存在", dto.Title))
	}

	// 按逗号、顿号、斜杠、竖线拆分、空格
	splitChars := []string{",", "，", "/", "|", " "}
	for _, s := range splitChars {
		dto.Tags = strings.ReplaceAll(dto.Tags, s, ",") // 先统一为逗号
	}

	var tags []string
	for _, tag := range strings.Split(dto.Tags, ",") {
		if tag == "" {
			continue
		}
		tags = append(tags, strings.TrimSpace(tag))
	}

	dto.Tags = strings.Join(tags, "、")

	m := &models.Article{
		Title:      dto.Title,
		CategoryID: *dto.CategoryID,
		Tags:       strings.ReplaceAll(dto.Tags, " ", ""),
		Summary:    dto.Summary,
		Content:    dto.Content,
	}

	if err := a.Create(m); err != nil {
		return 0, errors.New(fmt.Sprintf("文章创建失败，%v", err))
	}
	return m.ID, nil
}

func (r *ApiArticleRepository) UpdateArticle(dto articleDTO.ApiArticleDTO) error {
	a := r.Q.Article

	// 查询分类
	if dto.CategoryName == "all" {
		category, err := r.Q.Category.Select(r.Q.Category.ID).Where(r.Q.Category.Name.Eq("default")).First()
		if err != nil {
			return errors.New(fmt.Sprintf("分类查询失败，%v", err))
		}
		dto.CategoryID = &category.ID
	}

	// 查询当前分类下的
	count, err := a.Where(a.ID.Neq(*dto.ID), a.Title.Eq(dto.Title), a.CategoryID.Eq(*dto.CategoryID)).Count()
	if err != nil {
		return errors.New(fmt.Sprintf("[%s]查询失败，%v", dto.Title, err))
	}
	if count > 0 {
		return errors.New(fmt.Sprintf("该分类，该文章[%s]已存在", dto.Title))
	}

	for _, s := range []string{",", "，", "/", "|"} {
		dto.Tags = strings.ReplaceAll(dto.Tags, s, "、")
	}
	info, err := a.Where(a.ID.Eq(*dto.ID)).UpdateSimple(
		a.Title.Value(dto.Title),
		a.Tags.Value(dto.Tags),
		a.Summary.Value(dto.Summary),
		a.Content.Value(dto.Content),
		a.CategoryID.Value(*dto.CategoryID),
	)
	if err != nil {
		return err
	}
	return info.Error
}

func (r *ApiArticleRepository) UploadImages(c *core.Context) (*articleVO.ApiMdImagesVO, error) {
	_file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("文件读取失败: %v", err)
	}
	defer _file.Close()

	cfg := config.GetConfig()
	if header.Size > cfg.Server.UploadSize {
		return nil, fmt.Errorf("文件太大，不能超过 %dMB", cfg.Server.UploadSize>>20)
	}

	buff := make([]byte, 512)
	if _, err := _file.Read(buff); err != nil {
		return nil, fmt.Errorf("文件读取失败: %v", err)
	}
	filetype := http.DetectContentType(buff)
	_file.Seek(0, io.SeekStart)
	if !strings.HasPrefix(filetype, "image/") {
		return nil, errors.New("仅支持图片文件上传")
	}

	// 计算哈希
	hasher := sha256.New()
	if _, err := io.Copy(hasher, _file); err != nil {
		return nil, fmt.Errorf("计算文件哈希失败: %v", err)
	}
	hashValue := hex.EncodeToString(hasher.Sum(nil))
	_file.Seek(0, io.SeekStart)

	// 查询数据库是否已存在
	if record, err := r.Q.FileRecord.Where(r.Q.FileRecord.Hash.Eq(hashValue)).First(); err == nil {
		return &articleVO.ApiMdImagesVO{
			Url:  record.Url,
			Name: record.Name,
			Type: record.Type,
			Size: record.Size,
			Hash: record.Hash,
		}, nil
	}

	// 创建目录
	now := time.Now()
	saveDir := filepath.Join(constants.UploadPath, "images", fmt.Sprintf("%d/%02d/%02d", now.Year(), now.Month(), now.Day()))
	if err := file.CreateDir(saveDir); err != nil {
		return nil, fmt.Errorf("创建目录失败: %v", err)
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", utils.GenerateID(), ext)
	savePath := filepath.ToSlash(filepath.Join(saveDir, filename))

	out, err := os.Create(savePath)
	if err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, _file); err != nil {
		return nil, fmt.Errorf("写入文件失败: %v", err)
	}

	fileURL := fmt.Sprintf("/uploads%s", strings.TrimPrefix(savePath, constants.UploadPath))

	newFile := models.FileRecord{
		Name: header.Filename,
		Url:  fileURL,
		Type: filetype,
		Size: header.Size,
		Hash: hashValue,
	}

	if err := r.Q.FileRecord.Create(&newFile); err != nil {
		_ = os.Remove(savePath)
		return nil, fmt.Errorf("存储文件信息失败: %v", err)
	}

	// 成功返回
	return &articleVO.ApiMdImagesVO{
		Name: newFile.Name,
		Url:  newFile.Url,
		Type: newFile.Type,
		Size: newFile.Size,
		Hash: newFile.Hash,
	}, nil
}
