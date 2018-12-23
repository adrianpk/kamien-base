package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleResource() *models.Resource {
	u := createResource("name", "")
	return u
}

func createSampleResource2() *models.Resource {
	u := createResource("name", "2")
	return u
}

func createResource(name, sufix string) *models.Resource {
	tx := tenv.GetDBxTx()
	resourceRepo := repo.MakeResourceRepoTx(tx)
	resource := models.MakeResource()
	ns := fmt.Sprintf("Name%s", sufix)
	resource.Name = mdl.ToNullsString(ns)
	resource.GenerateID()
	resource.GenerateTag()
	resourceRepo.Create(resource)
	err := tx.Commit()
	checkErr(err, "Cannot create sample resource.")
	return resource
}

func getResource(id uuid.UUID) *models.Resource {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourceRepoTx(tx)
	resource, err := repo.Get(id)
	checkErr(err, "Cannot get resource.")
	err = tx.Commit()
	checkErr(err, "Cannot get resource.")
	return resource
}

func getResourceByName(name string) *models.Resource {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourceRepoTx(tx)
	resource, err := repo.GetByName(name)
	checkErr(err, "Cannot get resource.")
	err = tx.Commit()
	checkErr(err, "Cannot get resource.")
	return resource
}

func checkResource(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourceRepoTx(tx)
	resource, err := repo.Get(id)
	logErr(err, "Cannot get resource.")
	err = tx.Commit()
	logErr(err, "Cannot get resource.")
	return resource != nil && err == nil
}

func checkResourceByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourceRepoTx(tx)
	resource, err := repo.GetByName(name)
	logErr(err, "Cannot get resource.")
	err = tx.Commit()
	logErr(err, "Cannot get resource.")
	return resource != nil && err == nil
}

func resourcesMatch(resource, tc *models.Resource) bool {
	return resource.Match(tc)
}
