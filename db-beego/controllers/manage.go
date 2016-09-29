package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"db-beego/models"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

// Home() - shows the homepage controller
func (manage *ManageController) Home() {
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/home.html"
}

// Add() - adds the user detail
func (manage *ManageController) Add() {
	manage.Data["Form"] = &models.Article{}
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/add.html"
	flash := beego.ReadFromRequest(&manage.Controller)
	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["flash"] = ok
	}
	o := orm.NewOrm()
	o.Using("default")
	article := models.Article{}
	if err := manage.ParseForm(&article); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		manage.Data["Articles"] = article
		valid := validation.Validation{}
		isValid, _ := valid.Valid(article)

		if manage.Ctx.Input.Method() == "POST" {
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {
				searchArticle := models.Article{Name: article.Name}
				beego.Debug("Article name supplied:", article.Name)
				err = o.Read(&searchArticle)
				beego.Debug("Err:", err)
				flash := beego.NewFlash()

				if err == orm.ErrNoRows || err == orm.ErrMissPK {
					beego.Debug("No article found matching details supplied. Attempting to insert article: ", article)
					id, err := o.Insert(&article)
					if err == nil {
						msg := fmt.Sprintf("Article inserted with id:", id)
						beego.Debug(msg)
						flash.Notice(msg)
						flash.Store(&manage.Controller)
					} else {
						msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
						beego.Debug(msg)
						flash.Error(msg)
						flash.Store(&manage.Controller)
					}
				} else {
					beego.Debug("Article found matching details supplied. Cannot insert")
				}
			}
		}
	}
}

// Update() - updates the user detail by using their id
func (manage *ManageController) Update() {
	manage.Data["Form"] = &models.Article{}
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/edit.html"
	flash := beego.ReadFromRequest(&manage.Controller)
	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["flash"] = ok
	}
	o := orm.NewOrm()
	o.Using("default")
	// convert the string value to an int
	if articleId, err := strconv.Atoi(manage.Ctx.Input.Param(":id")); err == nil {
		article := models.Article{Id: articleId}
		// attempt to load the record from the database
		if o.Read(&article) == nil {
			if err := manage.ParseForm(&article); err != nil {
				beego.Error("Couldn't parse the form. Reason: ", err)
			} else {
				manage.Data["Articles"] = article
				valid := validation.Validation{}
				isValid, _ := valid.Valid(article)
				
				if manage.Ctx.Input.Method() == "POST" {
					if !isValid {
						manage.Data["Errors"] = valid.ErrorsMap
						beego.Error("Form didn't validate.")
					} else {
						searchArticle := models.Article{Name: article.Name}
						beego.Debug("Article name supplied:", article.Name)
						err = o.Read(&searchArticle)
						beego.Debug("Err:", err)
						flash := beego.NewFlash()
						
						if num, err := o.Update(&article); err == nil {
							beego.Debug("Attempting to insert article: ", article)
							flash.Notice("Record Was Updated.")
							flash.Store(&manage.Controller)
							beego.Info("Record Was Updated. ", num)
						} else {
							flash.Notice("Record Was NOT Updated.")
							flash.Store(&manage.Controller)
							beego.Error("Couldn't find article matching id: ", articleId)
						}
					} 
				}
			}	
		}
	}
}

// Delete() - deletes the user by their id, if that user's available.
func (manage *ManageController) Delete() {
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/home.html"
	// convert the string value to an int
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")
	article := models.Article{}
	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
		if num, err := o.Delete(&models.Article{Id: articleId}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
}

// View() - views all the user
func (manage *ManageController) View() {
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/view.html"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		manage.Data["notices"] = ok
	}
	
	o := orm.NewOrm()
	o.Using("default")

	var formss []*models.Article
	num, err := o.QueryTable("formss").All(&formss)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = formss
	}
}

// Vieww() - view the user, by their id, if that user's available.
func (manage *ManageController) Vieww() {
	manage.Layout = "basic-layout.html"
	manage.LayoutSections = make(map[string]string)
	manage.LayoutSections["Header"] = "header.html"
	manage.LayoutSections["Footer"] = "footer.html"
	manage.TplName = "manage/view.html"

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		manage.Data["notices"] = ok
	}
	articleId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")

	var formss []*models.Article
	num, err := o.QueryTable("formss").Filter("id",articleId).All(&formss)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = formss
	}
}

