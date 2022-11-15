package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// レシピデータの構造体
type recipeData struct {
  Name string `json:name`
  Img string `json:img`
}

// MYレシピデータの構造体
type myRecipeData struct {
	Monday []recipeData
	Tuesday []recipeData
	Wednesday []recipeData
	Thursday []recipeData
	Friday []recipeData
	Saturday []recipeData
	Sunday []recipeData
}

func RecipeDataGet(c *gin.Context) {
  // レシピデータ
  recipeData := []recipeData{
    {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
    {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
    {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
    {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
    {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
    {Name: "卵かけご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
    {Name: "美味しい煮物", Img: "../../assets/images/nimono.jpg"},
	}
  c.JSON(http.StatusOK, recipeData)
}

func MyRecipeDataSave(c *gin.Context) {
  c.String(http.StatusOK, "my献立保存API")
}

func MyRecipeDataGet(c *gin.Context) {
	myRecipeData := myRecipeData{
    Monday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Tuesday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵かけご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Wednesday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Thursday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Friday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Saturday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
    Sunday: []recipeData{
      {Name: "唐揚げ", Img: "../../assets/images/karaage.jpg"},
      {Name: "サラダ", Img: "../../assets/images/sarada.jpg"},
      {Name: "煮物", Img: "../../assets/images/nimono.jpg"},
      {Name: "卵をかけたご飯", Img: "../../assets/images/tamagokakegohan.jpg"},
      {Name: "味噌汁", Img: "../../assets/images/misoshiru.jpg"},
		},
  }
  c.JSON(http.StatusOK, myRecipeData)
}

func RecipeCreateSettingGet(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件取得API")
}

func RecipeCreateSettingSave(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件保存API")
}