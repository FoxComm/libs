package configs

import "github.com/FoxComm/libs/Godeps/_workspace/src/github.com/gorilla/sessions"

var CookieStore = sessions.NewCookieStore([]byte("super-secret-donkey"))
