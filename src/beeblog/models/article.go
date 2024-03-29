package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

//文章查询参数
type ArticleParams struct {
	Title      string
	Status     int
	FromTime   string
	EndTime    string
	Tag        string
	CategoryId int64
	Ids        []int64
}

//根据文章id获取文章内容
func GetArticleOne(id int64) (*Article, error) {
	article := new(Article)
	o := orm.NewOrm()
	err := o.QueryTable("article").One(article)
	return article, err
}

//分页
func GetArticlePager(params ArticleParams,pageNo,pageSize int64)(article []*Article,count int64,err error) {
	count, err = GetArticleCount(params)
	if err != nil {
		return article, 0, err
	}
	articles, err := GetArticleList(params, pageNo, PageSize)
	if err != nil {
		return article, 0, err
	}
	return articles, count, nil
}

//获取列表
func GetArticleList(params ArticleParams,pageNo,pageSize int64) (article []*Article,err error) {
		o := orm.NewOrm()
		qs := o.QueryTable("article")
		if len(params.Title) > 0 {
			qs = qs.Filter("title__contains", params.Title)
		}
		if params.Status > 0 {
			qs = qs.Filter("status", params.Status)
		}
		if params.CategoryId > 0 {
			qs = qs.Filter("category_id", params.CategoryId)
		}
		if len(params.FromTime) > 0 {
			t, err := time.Parse("2000-01-01 12:00:00", params.FromTime)
			if err == nil {
				qs = qs.Filter("created_at__gt", t)
			}
		}
		if len(params.EndTime) > 0 {
			t, err := time.Parse("2000-01-01 12:00:00", params.EndTime)
			if err == nil {
				qs = qs.Filter("created_at__lte", t)
			}
		}
		if len(params.Tag) > 0 {
			qs = qs.Filter("tags__contains", params.Tag)
		}
		if len(params.Ids) > 0 {
			qs = qs.Filter("id__in", params.Ids)
		}
		offset := (pageNo - 1) * pageSize
		_, err = qs.OrderBy("-id").Limit(pageSize, offset).All(&article)
		if err != nil {
			return nil, err
		}
		return article, nil	
}

//获取总数
func GetArticleCount(params ArticleParams) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	if len(params.Title) > 0 {
		qs = qs.Filter("title__contains", params.Title)
	}
	if params.Status > 0 {
		qs = qs.Filter("status", params.Status)
	}
	if len(params.FromTime) > 0 {
		t, err := time.Parse("2019-01-01 15:00:00", params.FromTime)
		if err == nil {
			qs = qs.Filter("created_at__gt", t)
		}
	}
	if len(params.EndTime) > 0 {
		t, err := time.Parse("2019-01-01 15:00:00", params.EndTime)
		if err == nil {
			qs = qs.Filter("created_at__lte", t)
		}
	}
	if len(params.Tag) > 0 {
		qs = qs.Filter("tags__contains", params.Tag)
	}
	cnt, err := qs.Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

//添加文章
func AddArticle(article *Article, content string) error {
	orm.Debug = true
	o := orm.NewOrm()
	o.Begin()
	tags := article.Tags
	var tagsArr []string
	if len(tags) > 0 { //标签处理
		tagsArr = strings.Split(tags, ",")
		tags = strings.Join(tagsArr, "$#")
		tags = "#" + tags + "$"
	}
	article.Tags = tags
	_, err := o.Insert(article)
	if err != nil {
		err = o.Rollback()
	}
	//添加文章内容
	cont := &Content{
		Content:   content,
		ArticleId: article.Id,
	}
	_, err = o.Insert(cont)
	if err != nil {
		err = o.Rollback()
	}
	existsTags, err := GetTagsByNames(tagsArr)
	if err != nil {
		err = o.Rollback()
	}
	mapTags := ArrayToMap(existsTags)
	for _, tag := range tagsArr {
		var tagModel *Tag
		if m, ok := mapTags[tag]; ok {
			tagModel = m
		} else { //标签不存在，添加
			t := &Tag{
				Name: tag,
			}
			_, err := o.Insert(t)
			if err != nil {
				err = o.Rollback()
			}
			tagModel = t
		}
		//添加标签和文章对应关系
		articleTag := &ArticleTag{
			ArticleId: article.Id,
			TagId:     tagModel.Id,
		}
		_, err = o.Insert(articleTag)
		if err != nil {
			err = o.Rollback()
		}
	}
	o.Commit()
	return err
}

//更新文章
func UpdateArticle(article *Article, content *Content) error {
	o := orm.NewOrm()
	o.Begin()
	tags := article.Tags
	var tagsArr []string
	if len(tags) > 0 { //标签处理
		tagsArr = strings.Split(tags, ",")
		tags = strings.Join(tagsArr, "$#")
		tags = "#" + tags + "$"
	}
	article.Tags = tags
	_, err := o.Update(article)
	if err != nil {
		err = o.Rollback()
	}
	//更新文章内容
	_, err = o.Update(content)
	if err != nil {
		err = o.Rollback()
	}
	existsTags, err := GetTagsByNames(tagsArr)
	if err != nil {
		err = o.Rollback()
	}
	mapTags := ArrayToMap(existsTags)
	//删除已存在文章标签关系
	err = DeleteArticleTags(article.Id)
	if err != nil {
		err = o.Rollback()
	}
	for _, tag := range tagsArr {
		var tagModel *Tag
		if m, ok := mapTags[tag]; ok {
			tagModel = m
		} else { //标签不存在，添加
			t := &Tag{
				Name: tag,
			}
			_, err := o.Insert(t)
			if err != nil {
				err = o.Rollback()
			}
			tagModel = t
		}
		//添加标签和文章对应关系
		articleTag := &ArticleTag{
			ArticleId: article.Id,
			TagId:     tagModel.Id,
		}
		_, err = o.Insert(articleTag)
		if err != nil {
			err = o.Rollback()
		}
	}
	o.Commit()
	return err
}

//删除文章
func DeleteArticle(article *Article) error {
	o := orm.NewOrm()
	o.Begin()
	//删除文章内容
	_, err := o.Raw("delete from content where article_id = ?", article.Id).Exec()
	if err != nil {
		err = o.Rollback()
	}
	_, err = o.Delete(article)
	if err != nil {
		err = o.Rollback()
	}
	o.Commit()
	return err
}

//数组转map
func ArrayToMap(tags []*Tag) map[string]*Tag {
	var mTags = make(map[string]*Tag)
	for _, t := range tags {
		mTags[t.Name] = t
	}
	return mTags
}

