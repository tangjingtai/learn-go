package v1

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/service/tag_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

//获取多个文章标签
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	//name := c.Query("name")
	//state := -1
	//if arg := c.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//}
	//
	//tagService := tag_service.Tag{
	//	Name:     name,
	//	State:    state,
	//	PageNum:  util.GetPage(c),
	//	PageSize: setting.AppSetting.PageSize,
	//}
	//tags, err := tagService.GetAll()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	//	return
	//}
	//
	//count, err := tagService.Count()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": nil,
		"total": nil,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS{
		appG.Response(httpCode, errCode, nil)
	}

	serviceTag:=tag_service.Tag{
		Name: form.Name,
		CreatedBy: form.CreatedBy,
		State:form.State,
	}

	//name := c.Query("name")
	//state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	//createdBy := c.Query("created_by")
	//
	//valid := validation.Validation{}
	//valid.Required(name, "name").Message("名称不能为空")
	//valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	//valid.Required(createdBy, "created_by").Message("创建人不能为空")
	//valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	//valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	exists, err := serviceTag.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = serviceTag.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "ID"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//删除文章标签
// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

