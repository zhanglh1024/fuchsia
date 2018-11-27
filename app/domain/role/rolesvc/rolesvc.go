package rolesvc

import (
    "github.com/cxr29/log"
    "beginner-server/app"
    "beginner-server/app/domain/battle/battleobj"
    "beginner-server/app/domain/hero/heroobj"
    "beginner-server/app/domain/resource/resourceobj"
    "beginner-server/app/domain/role/roleobj"
    "beginner-server/app/model"
    "beginner-server/app/model/dyn"
    "bytes"
    "crypto/des"
    "encoding/hex"
    "errors"
    "strings"
)

var key = []byte("2fa6c1e9")


//获取用户信息
func GetRoleInfo(roleId string) (dyn.RoleInfoDyn, error) {
    roleObj:=roleobj.RoleObj{}
    info, err := roleObj.GetRoleInfoByRoleId(roleId)
    return info, err
}


//roleId的用户是否存在
func IsExitRoleInfo(roleId string) bool {
    roleObj := roleobj.RoleObj{}
    return roleObj.IsIncludeRoleInfoRoleId(roleId)
}

//注册操作
func Register(userId, name, password string) app.CodeErrorType{
	roleObj := roleobj.RoleObj{}

    includeRoleId := roleObj.IsIncludeRoleInfoRoleId(userId)
    if !includeRoleId{
        return app.RoleIdAlreadyExists
    }

    includeName := roleObj.IsIncludeRoleInfoName(name)
    if !includeName{
        return app.RoleNameAlreadyExists
    }

    passwordEncrypted, err := Encrypt(password, key)
    if err != nil {
        log.Fatal(err)
    }

	Code1 := roleObj.InsertRoleInfo(userId,name, passwordEncrypted)
	if Code1 != 0{
	    return app.InsertDataError
    }
    flag := initHero(userId)
    if flag != nil{
        return app.InsertDataError
    }
    err1 := initResource(userId)
    if err1 != nil{
        return app.InsertDataError
    }

    err = InitResourceFresh(userId)
    if err != nil{
        return app.InsertDataError
    }

    err = initCurrentLevel(userId)
    if err != nil{
        return 5
    }

    err = initMatchBattle(userId)
    if err != nil{
        return app.InsertDataError
    }

	return 0

}

//用户密码验证
func LoginVerification(roleId, password string)(bool,error){

    roleObj := roleobj.RoleObj{}
    Info , err := roleObj.GetRoleInfoByRoleId(roleId)
    if err != nil{
        log.Debugf("获取账号信息错误%s", err)
        return false, err
    }

    //密码信息解码
    strDecrypted, err := Decrypt(Info.Password, key)
    if err != nil {
        log.Fatal(err)
        return false, errors.New("解密失败")
    }

    correct := strings.Compare(strDecrypted, password)
    if correct == 0 {
        return true, nil
    }

    return false, errors.New("密码错误")
}

//初始化英雄数据
func initHero(roleId string) error {
    hero := model.ProfConfig.ConfHero[0]
    herodyn := dyn.HeroInfoDyn{}
    herodyn.No = int(hero.HeroNum)
    herodyn.Lv = 1
    herodyn.RoleId = roleId
    herodyn.Wise = hero.InitWise
    herodyn.Loyal = hero.InitLoyalty
    herodyn.Diligent = hero.InitDiligent
    herodyn.Heroic = hero.InitHeroic
    err := heroobj.InsertHeroValue(herodyn)

    return err
}

//初始化用户资源数据
func initResource(roleId string)error{
    resource := dyn.ResourceInfoDyn{}
    resource.RoleId = roleId
    resource.Soldier = 1000
    resource.Coin = 1000
    resource.Food = 1000

    err := resourceobj.InsertResource(resource)
    return err
}

//初始化关卡数据
func initCurrentLevel(roleId string) error {
    info := model.ProfConfig.ConfLever[0]
    CurrentLevel := dyn.LevelBattle{}
    CurrentLevel.RoleId = roleId
    CurrentLevel.RoundId = int(info.No)
    CurrentLevel.SoldierLeft = int(info.SoldierNum)
    CurrentLevel.HeroFight = int(info.HeroFight)
    return battleobj.InsertLevelBattleValue(CurrentLevel)
}

//初始化战斗匹配数据
func initMatchBattle(roleId string) error {
    info := dyn.MatchBattle{}
    info.RoleId = roleId
    info.Integral = 0
    info.Register = 0
    info.HeroNo = 0
    info.Soldier = 0
    return battleobj.InsertDateIntoMatchBattle(info)
}

//初始化资源更新时间数据
func InitResourceFresh(roleId string)error{
    refreshResource := resourceobj.ResourceRefreshObj{}
    for _,v := range app.ResourceTypeList{
        err := refreshResource.InsertResourceRefreshTime(roleId,v,0)
        if err != nil{
            log.Errorf("初始化資源表錯誤")
            return err
        }
    }
    return nil
}


func ZeroPadding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{0}, padding)
    return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
    return bytes.TrimFunc(origData,
        func(r rune) bool {
            return r == rune(0)
        })
}


//加密数据
func Encrypt(text string, key []byte) (string, error) {
    src := []byte(text)
    block, err := des.NewCipher(key)
    if err != nil {
        return "", err
    }
    bs := block.BlockSize()
    src = ZeroPadding(src, bs)
    if len(src)%bs != 0 {
        return "", errors.New("Need a multiple of the blocksize")
    }
    out := make([]byte, len(src))
    dst := out
    for len(src) > 0 {
        block.Encrypt(dst, src[:bs])
        src = src[bs:]
        dst = dst[bs:]
    }
    return hex.EncodeToString(out), nil
}

//数据解密

func Decrypt(decrypted string , key []byte) (string, error) {
    src, err := hex.DecodeString(decrypted)
    if err != nil {
        return "", err
    }
    block, err := des.NewCipher(key)
    if err != nil {
        return "", err
    }
    out := make([]byte, len(src))
    dst := out
    bs := block.BlockSize()
    if len(src)%bs != 0 {
        return "", errors.New("crypto/cipher: input not full blocks")
    }
    for len(src) > 0 {
        block.Decrypt(dst, src[:bs])
        src = src[bs:]
        dst = dst[bs:]
    }
    out = ZeroUnPadding(out)
    return string(out), nil
}