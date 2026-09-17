package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)


type RoleController struct{
	RoleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController{
	return &RoleController{
		RoleService: _roleService,
	}
}

func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request){
	roleId := chi.URLParam(r, "id") 

	if roleId == ""{
		utils.WriteJsonErrorsResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	id ,err := strconv.ParseInt(roleId,10,64)
	if err != nil{
		utils.WriteJsonErrorsResponse(w,http.StatusBadRequest,"Invalid role id",err)
		return
	}

	role,err := rc.RoleService.GetRoleById(id)

	if err != nil{
		utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"Failed to fetch role",err)
		return
	}


	if role == nil{
		utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"Role not found",fmt.Errorf("role with ID %d not found", roleId))
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request){
	roles,err := rc.RoleService.GetAllRoles()

	if err != nil{
		utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"Failed to fetch roles",err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
} 