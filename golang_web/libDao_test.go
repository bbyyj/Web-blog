package main

import (
	"blog_web/db/dao"
	"blog_web/model"
	"testing"
)

func TestAddResource(t *testing.T) {
	lbd := dao.NewLinkDao()
	var r model.ResourceManage
	r.ID = 1000
	r.Name = "test.zip"
	r.Desc = "test.zip"
	err := lbd.AddResource(&r)
	if err != nil {
		t.Fatal("AddResource Error")
	}
	t.Log("AddResource Success")
}

func TestAddResourceCheck(t *testing.T) {
	lbd := dao.NewLinkDao()
	var r model.ResourceManage
	r.ID = 1000
	r.Name = "test.zip"
	r.Desc = "test.zip"
	err := lbd.AddResourceCheck(&r)
	if err != nil {
		t.Fatal("AddResourceCheck Error")
	}
	t.Log("AddResourceCheck Success")
}

func TestFindAllResource(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, err := lbd.FindAllResource(1, 5)
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("Find All Resource Error")
	}
	t.Log("Find All Resource Success")
}

func TestFindAllResourceByCategoryId(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, _, err := lbd.FindAllResourceByCategoryId(1, 1, 5)
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("Find All Resource By CategoryId Error")
	}
	t.Log("Find All Resource By CategoryId Success")
}

func TestFindAllResourceLikeName(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, _, err := lbd.FindAllResourceLikeName("test", 1, 5)
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("Find All Resource Like Name Error")
	}
	t.Log("Find All Resource Like Name Success")
}

func TestGetDownloadUrl(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, err := lbd.GetDownloadUrl("test.zip")
	if err != nil {
		t.Fatal("GetDownloadUrl Error")
	}
	t.Log("GetDownloadUrl Success")
}

func TestFindAllResourceByName(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, err := lbd.FindAllResourceByName("test.zip")
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("Error")
	}
	t.Log("Success")
}

func TestUpdateResourceDownloadNumByName(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.UpdateResourceDownloadNumByName("test.zip")
	if err != nil {
		t.Fatal("UpdateResourceDownloadNumByName Error")
	}
	t.Log("UpdateResourceDownloadNumByName Success")
}

func TestFindAllCategory(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, err := lbd.FindAllCategory()
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("FindAllCategory Error")
	}
	t.Log("FindAllCategory Success")
}

func TestFindLimitedResource(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, _, err := lbd.FindLimitedResource(1, 5)
	if err != nil && err.Error() != "No matching records found" {
		t.Fatal("FindLimitedResource Error")
	}
	t.Log("FindLimitedResource Success")
}

func TestFindLimitedResourceCheck(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, _, err := lbd.FindLimitedResourceCheck(1, 5)
	if err != nil {
		t.Fatal("Error")
	}
	t.Log("Success")
}

func TestUpdateResource(t *testing.T) {
	lbd := dao.NewLinkDao()
	var r model.ResourceManage
	r.ID = 1000
	r.Name = "test.zip"
	r.Desc = "test1.zip"
	err := lbd.UpdateResource(&r)
	if err != nil {
		t.Fatal("UpdateResource Error")
	}
	t.Log("UpdateResource Success")
}

func TestCheckResourceName(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.CheckResourceName(1000, "test1111.zip")
	if err != nil {
		t.Fatal("CheckResourceName Error")
	}
	t.Log("CheckResourceName Success")
}

func TestCheckResourceCheckName(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.CheckResourceCheckName(1000, "test1111.zip")
	if err != nil {
		t.Fatal("CheckResourceCheckName Error")
	}
	t.Log("CheckResourceCheckName Success")
}

func TestGetResourceLikeName(t *testing.T) {
	lbd := dao.NewLinkDao()
	_, _, err := lbd.GetResourceLikeName("test.zip", 1, 5)
	if err != nil {
		t.Fatal("GetResourceLikeName Error")
	}
	t.Log("GetResourceLikeName Success")
}

func TestReUploadUpdateTime(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.ReUploadUpdateTime("test.zip")
	if err != nil {
		t.Fatal("ReUploadUpdateTime Error")
	}
	t.Log("ReUploadUpdateTime Success")
}

func TestGetResourceUrl(t *testing.T) {
	lbd := dao.NewLinkDao()
	url := lbd.GetResourceUrl(1000)
	if url.Name != "test.zip" {
		t.Fatal("GetResourceUrl Error")
	}
	t.Log("GetResourceUrl Success")
}

func TestGetCheckResourceById(t *testing.T) {
	lbd := dao.NewLinkDao()
	r := lbd.GetCheckResourceById(1000)
	if r.ID != 1000 {
		t.Fatal("GetCheckResourceById Error")
	}
	t.Log("GetCheckResourceById Success")
}

func TestDeleteResource(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.DeleteResource(1000)
	if err != nil {
		t.Fatal("DeleteResource Error")
	}
	t.Log("DeleteResource Success")
}

func TestDeleteResourceCheck(t *testing.T) {
	lbd := dao.NewLinkDao()
	err := lbd.DeleteResourceCheck(1000)
	if err != nil {
		t.Fatal("DeleteResourceCheck Error")
	}
	t.Log("DeleteResourceCheck Success")
}
