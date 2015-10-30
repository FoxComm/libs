package configs

import "github.com/gorilla/sessions"

var CookieStore = sessions.NewCookieStore([]byte("super-secret-donkey"))
