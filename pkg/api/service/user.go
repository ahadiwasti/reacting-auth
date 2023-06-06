package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"reacting-auth/pkg/api/dao"
	"reacting-auth/pkg/api/domain/account"
	"reacting-auth/pkg/api/domain/perm"
	"reacting-auth/pkg/api/domain/user"
	"reacting-auth/pkg/api/dto"
	"reacting-auth/pkg/api/model"
	"reacting-auth/pkg/api/utils"
	"reacting-auth/pkg/api/utils/log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const pwHashBytes = 64

var (
	/* Dao layer */
	userDao           = dao.User{}
	userOauthDao      = dao.UserOAuthDao{}
	logDao            = dao.LoginLogDao{}
	errInvalidAccount = errors.New("Invalid Account")
	errInvalidCode    = errors.New("INVALID:999")
	errAccountLocked  = errors.New("Account Locked")
	RootDomainCode    = "root"
)

const (
	UserStatusNotPub = iota
	UserStatusNormal
	UserStatusLock
)

type UserService struct {
	//oauthdao *dao.UserOAuthDao
}

func (UserService) InfoOfId(dto dto.GeneralGetDto) model.User {
	return userDao.Get(dto.Id, true)
}

// List - users list with pagination
func (UserService) List(ctx context.Context, gdto dto.GeneralListDto) ([]model.User, int64) {
	cols := "*"
	gdto.Q, cols = dataPermService.DataPermFilter(ctx, "users", gdto)
	return userDao.List(gdto, cols)
}

// Create - create a new account
func (us UserService) Create(userDto dto.UserCreateDto) (*model.User, error) {
	//if username is exits,it can't create this user
	userModel := userDao.GetByUserName(userDto.Username)
	if userModel.Username == userDto.Username {
		return nil, errors.New("username is exits")
	}
	salt, _ := account.MakeSalt()
	pwd, _ := account.HashPassword(userDto.Password, salt)
	newUser := &model.User{
		Username:     userDto.Username,
		Mobile:       userDto.Mobile,
		Password:     pwd,
		DepartmentId: userDto.DepartmentId,
		Salt:         salt,
		Sex:          userDto.Sex,
		Email:        userDto.Email,
		Title:        userDto.Title,
		Realname:     userDto.Realname,
		Status:       userDto.Status,
	}
	c := userDao.Create(newUser)
	if c.Error != nil {
		log.Error(c.Error.Error())
		return nil, errors.New("create user failed")
	}

	if userDto.Roles != "" {
		us.AssignRole(strconv.Itoa(newUser.Id), strings.Split(userDto.Roles, ","))
	}

	return &userModel, nil
}

// Update - update user's information
func (us UserService) Update(userDto dto.UserEditDto) int64 {
	userModel := model.User{
		Id: userDto.Id,
	}
	c := userDao.Update(&userModel, map[string]interface{}{
		"mobile":        userDto.Mobile,
		"department_id": userDto.DepartmentId,
		"status":        userDto.Status,
		"title":         userDto.Title,
		"realname":      userDto.Realname,
		"sex":           userDto.Sex,
		"email":         userDto.Email,
	})
	us.AssignRole(strconv.Itoa(userDto.Id), strings.Split(userDto.Roles, ","))
	return c.RowsAffected
}

// UpdateStatus - update user's status only
func (UserService) UpdateStatus(dto dto.UserEditStatusDto) int64 {
	u := userDao.Get(dto.Id, false)
	//u.Status = dto.Status
	c := userDao.Update(&u, map[string]interface{}{
		"status": dto.Status,
	})
	return c.RowsAffected
}

// UpdatePassword - update password only
func (UserService) UpdatePassword(dto dto.UserEditPasswordDto) int64 {
	salt, _ := account.MakeSalt()
	pwd, _ := account.HashPassword(dto.Password, salt)
	u := userDao.Get(dto.Id, false)
	//u.Password = pwd
	//u.Salt = salt
	c := userDao.Update(&u, map[string]interface{}{
		"password": pwd,
		"salt":     salt,
	})
	return c.RowsAffected
}

// ResetPassword - automatically make a password
func (UserService) ResetPassword(gDto dto.GeneralGetDto) string {
	salt, _ := account.MakeSalt()
	//pwd, _ := account.HashPassword(dto.Password, salt)
	u := userDao.Get(gDto.Id, false)
	autoPwd := utils.RandomPwd(10)
	pwd, _ := account.HashPassword(autoPwd, salt)
	//u.Password = pwd
	//u.Salt = salt
	userDao.Update(&u, map[string]interface{}{
		"password": pwd,
		"salt":     salt,
	})
	return autoPwd
}

// Delete - delete user
func (UserService) Delete(dto dto.GeneralDelDto) int64 {
	userModel := model.User{
		Id: dto.Id,
	}
	c := userDao.Delete(&userModel)
	if c.RowsAffected > 0 {
		user.DeleteUser(strconv.Itoa(dto.Id))
	}
	return c.RowsAffected
}

// UpdateLoginTime update last login time after user successfully sign in
func (UserService) UpdateLoginTime(user dto.UserEditDto) int64 {
	u := userDao.Get(user.Id, false)
	//u.Status = dto.Status
	c := userDao.Update(&u, map[string]interface{}{
		"last_login_time": time.Now(),
	})
	return c.RowsAffected
}

// AssignRoleByRoleIds - assign roles to specific user
func (UserService) AssignRoleByRoleIds(userId string, roles string) {
	// update roles
	rs := roleDao.GetRolesByIds(roles)
	var groups [][]string
	for _, role := range rs {
		groups = append(groups, []string{userId, role.RoleName})
	}
	user.OverwriteRoles(userId, groups)
}

// AssignRole - assign roles to specific user
func (UserService) AssignRole(userId string, roleNames []string) {
	var roles [][]string
	for _, role := range roleNames {
		if userId == "" || role == "" {
			continue
		}
		roles = append(roles, []string{userId, role})
	}
	user.OverwriteRoles(userId, roles)
}

// GetRelatedDomains - get related domains
func (UserService) GetRelatedDomains(uid string, skipRoot bool) []model.Domain {
	var domains []model.Domain
	var single = map[string]bool{}
	//1.get roles by user
	roles := perm.GetGroupsByUser(uid)
	//2.get domains by roles
	for _, rn := range roles {
		role := roleDao.GetByName(rn[1])
		if skipRoot {
			if role.Domain.Code == RootDomainCode {
				continue
			}
		}
		if _, ok := single[role.Domain.Code]; !ok {
			single[role.Domain.Code] = true
			domains = append(domains, role.Domain)
		}
	}
	return domains
}

// GetAllRoles would return all roles of a user
func (UserService) GetAllRoles(uid string) []string {
	groups := perm.GetGroupsByUser(uid)
	roles := []string{}
	for _, group := range groups {
		roles = append(roles, group[1])
	}
	return roles
}

// GetAllPermissions - get all permission by specific user
func (UserService) GetAllPermissions(uid string) []string {
	perms := []string{}
	var path = map[string]bool{}
	for _, perm := range perm.GetAllPermsByUser(uid) {
		prefix := strings.Split(perm[1], ":")
		seg := strings.Split(prefix[0], "/")
		ss := ""
		for _, s := range seg[1:] {
			ss += "/" + s
			if ok := path[ss]; !ok {
				path[ss] = true
				perms = append(perms, ss)
			}
		}
		perms = append(perms, perm[1])
	}
	return perms
}

// GetPermissionsOfDomain - Get pure permission list  in specific domain(another backend system)
func (UserService) GetPermissionsOfDomain(uid string, domain string) []string {
	perms := perm.GetAllPermsByUser(uid)
	var polices []string
	for _, p := range perms {
		if p[3] == domain {
			polices = append(polices, p[1])
		}
	}
	return polices
}

// GetDataPermissionsOfDomain - Get data permission list  in specific domain(another backend system)
func (us UserService) GetDataPermissionsOfDomain(uid, domain string) []map[string]string {
	gs := perm.GetGroupsByUser(uid)
	var (
		polices []map[string]string
		roles   []string
	)
	for _, p := range gs {
		roles = append(roles, p[1])
	}
	dmHash := map[int]bool{}
	for _, dm := range us.GetRelatedDomains(uid, false) {
		if dm.Code != domain {
			continue
		}
		dmHash[dm.Id] = true
	}
	for _, r := range roleDao.GetRolesByNames(roles) {
		for _, dp := range r.DataPerm {
			if dp.PermsType == 2 {
				if _, ok := dmHash[dp.DomainId]; ok {
					polices = append(polices, map[string]string{
						"perm":   dp.Perms,
						"rule":   dp.PermsRule,
						"weight": strconv.Itoa(dp.OrderNum),
					})
				}
			}
		}
	}
	return polices
}

// GetMenusOfDomain - get menus in specific domain
func (UserService) GetMenusOfDomain(uid string, domain string) []model.Role {
	roles := perm.GetGroupsByUser(uid)
	var roleNames []string
	for _, r := range roles {
		roleNames = append(roleNames, r[1])
	}
	return roleDao.GetRolesByNames(roleNames)
}

// MoveToAnotherDepartment - move users to another department
func (UserService) MoveToAnotherDepartment(uids []string, target int) error {
	return userDao.UpdateDepartment(uids, target)
}

// GetDomainMenu -  get specific user's menus of specific domain
func (us UserService) GetDomainMenu(uid string, domain string) []model.Menu {
	roles := us.GetAllRoles(uid)
	mids := []string{}
	for _, r := range roleDao.GetRolesByNames(roles) {
		if r.Domain.Code == domain {
			mids = append(mids, r.MenuIds)
		}
	}
	return menuDao.GetMenusByIds(strings.Join(mids, ","))
}

// CheckPermission - check user's permission in specific domain with specific policy
func (us UserService) CheckPermission(uid string, domain string, policy string) (bool, error) {
	//Could it be an alias?
	domainModel := domainDao.GetByCode(domain)
	row := menuPermAliasDao.GetByAlias(policy, domainModel.Id)
	if row.Id > 0 {
		policy = row.Perms
	}
	log.Info(fmt.Sprintf("check permission for : %#v,%#v,%#v", uid, policy, domain))
	return perm.Enforce(uid, policy, "*", domain)
}

// InsertLoginLog insert login log
func (us UserService) InsertLoginLog(loginLogDto *dto.LoginLogDto) error {
	return logDao.Create(loginLogDto)
}

// GetLastPwdChangeDaySinceNow check when user changed pwd
// if account did not change pwd until 90 days later,should use user created time instead
func (us UserService) GetLastPwdChangeDaySinceNow(uDto dto.GeneralGetDto) (ok bool, days int) {
	if viper.GetInt("security.level") == 0 {
		return false, -1
	}
	oLog := operationLogDao.GetLatestPwdLogOfAccount(uDto.Id)
	// Got action log
	var lastTime time.Time
	if oLog.Id > 0 {
		lastTime = oLog.CreateTime
	} else {
		userInfo := userDao.Get(uDto.Id, false)
		lastTime = userInfo.CreateTime
	}
	deadLineTime := lastTime.Add(time.Second * 24 * 3600 * viper.GetDuration("login.pwdExpiredDays"))
	diff := deadLineTime.Sub(time.Now()).Seconds()
	//less then 7 days,should warn login user to change his/her password
	if diff <= 24*3600*viper.GetFloat64("login.pwdWarningDays") {
		return true, int(math.Ceil(diff / (24 * 3600)))
	}
	return false, -1
}
