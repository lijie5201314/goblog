package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BlogController struct {
	beego.Controller
}
/**
首页
*/
func (c *BlogController) Home()  {
	o := orm.NewOrm()
	var typemaps []orm.Params
	o.Raw("select category.name,count(*) as count from tb_category as category left join tb_post as post on category.id=post.id group by category.name").Values(&typemaps)
	c.Data["types"] = typemaps

	c.TplName= "blog/home.html"
}
/**
博客列表
*/
func (c *BlogController) Bloglist()  {
	o := orm.NewOrm()

	var typemaps []orm.Params
	o.Raw("select category.name,count(*) as count from tb_category as category left join tb_post as post on category.id=post.id group by category.name").Values(&typemaps)
	c.Data["types"] = typemaps

	//日期分类导航
	var datemaps []orm.Params
	o.Raw("select count(*) as count, DATE_FORMAT(created,'%Y-%m') as date  from tb_post group by DATE_FORMAT(created,'%Y-%m')").Values(&datemaps)
	c.Data["dates"] = datemaps

	var blogmaps []orm.Params
	o.Raw("select post.id,post.title,post.created,post.content,post.author,category.name from tb_post as post inner join tb_category as category on category.id=post.id").Values(&blogmaps)
	c.Data["blogs"] = blogmaps

	c.TplName= "blog/bloglist.html"
}

func (c *BlogController) Blogtypelist()  {
	o := orm.NewOrm()
	blogtype := c.Ctx.Input.Param(":hi")
	//分类博客列表
	var blogtypemaps []orm.Params
	o.Raw("select post.id,post.title,post.created,post.content,post.author,category.name from tb_post as post inner join tb_category as category on category.id=post.id where name=?",blogtype).Values(&blogtypemaps)
	c.Data["blogs"] = blogtypemaps

	var typemaps []orm.Params
	o.Raw("select category.name,count(*) as count from tb_category as category left join tb_post as post on category.id=post.id group by category.name").Values(&typemaps)
	c.Data["types"] = typemaps

	var datemaps []orm.Params
	o.Raw("select count(*) as count, DATE_FORMAT(created,'%Y-%m') as date  from tb_post group by DATE_FORMAT(created,'%Y-%m')").Values(&datemaps)
	c.Data["dates"] = datemaps

	c.TplName= "blog/bloglist.html"
}

func (c *BlogController) Datetypelist()  {
	o := orm.NewOrm()

	var datemaps []orm.Params
	o.Raw("select count(*) as count, DATE_FORMAT(created,'%Y-%m') as date  from tb_post group by DATE_FORMAT(created,'%Y-%m')").Values(&datemaps)
	c.Data["dates"] = datemaps

	var typemaps []orm.Params
	o.Raw("select category.name,count(*) as count from tb_category as category left join tb_post as post on category.id=post.id group by category.name").Values(&typemaps)
	c.Data["types"] = typemaps

	//日期博客列表
	var datetypemaps []orm.Params
	o.Raw("select * from tb_post where date_format( created, '%Y%m' ) = date_format(curdate( ) , '%Y%m' )").Values(&datetypemaps)
	c.Data["blogs"] = datetypemaps

	c.TplName= "blog/bloglist.html"
}

func (c *BlogController) Blogdetail()  {
	o := orm.NewOrm()
	blogid := c.Ctx.Input.Param(":id")
	var blogdetailmaps []orm.Params
	o.Raw("select post.id,post.title,post.created,post.content,post.author,category.name from tb_post as post inner join tb_category as category on category.id=post.id where post.id=?", blogid).Values(&blogdetailmaps)
	c.Data["blogsdetail"] = blogdetailmaps

	var randommaps []orm.Params
	o.Raw("select id,title from tb_post limit 0,7").Values(&randommaps)
	c.Data["randomblog"] = randommaps

	//上一篇博客
	var premmaps []orm.Params
	o.Raw("select id,title from tb_post where id<(select id from tb_post where id=?) order by id desc limit 1",blogid).Values(&premmaps)
	c.Data["preblog"] = premmaps

	//下一篇博客
	var nextmaps []orm.Params
	o.Raw("select id,title from tb_post where id>(select id from tb_post where id=?) order by id asc limit 1",blogid).Values(&nextmaps)
	c.Data["nextblog"] = nextmaps

	c.TplName= "blog/blogdetail.html"
}