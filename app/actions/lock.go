package actions

import (
	"github.com/cxr29/log"
	"sync"
)


type MatchLock struct {
	lock sync.Mutex
	isMatchedRole []string
}


func (l *MatchLock) Lock(roleId string) (isSuccess bool){
	l.lock.Lock()
	log.Info("Locking, roleID: ", roleId)
	defer l.lock.Unlock()
	if containString(roleId, l.isMatchedRole) {
		log.Info("Lock failed, roleID: ", roleId)
		return false
	}

	l.isMatchedRole = append(l.isMatchedRole, roleId)
	return true
}

func (l *MatchLock) Unlock(roleId string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	log.Info("Unlocking roleID: ", roleId)
	matchedRole := l.isMatchedRole
	log.Info("lock Id",l.isMatchedRole)
	for i, s := range matchedRole {
		if s == roleId{
			//pop role
			l.isMatchedRole = append(l.isMatchedRole[:i], l.isMatchedRole[i+1:]...)
			return
		}
	}
	return
}

func containString(s string, list []string) bool{
	for _, str := range list {
		if s == str {
			return true
		}
	}
	return false
}




