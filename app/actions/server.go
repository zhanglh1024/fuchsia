package actions

import "net/http"

//注册各个路由
func Register() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/get_resource", GetResourceHandler)
	http.HandleFunc("/upgrade_hero", UpdateHeroHandler)
	http.HandleFunc("/section_fight", SectionFightHandler)
	http.HandleFunc("/register_fight", RegisterFightHandler)
	http.HandleFunc("/match_fight", MatchFightHandler)
}
