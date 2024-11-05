package httptest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       int
	Username string
}

type Server struct {
	db        map[int]*User
	userCache map[int]*User

	dbhit int
}

func NewServer() *Server {
	db := make(map[int]*User)

	for i := 1; i <= 100; i++ {
		db[i] = &User{
			ID:       i,
			Username: fmt.Sprintf("user_%d", i),
		}
	}

	return &Server{
		db: db,

		// init cache
		userCache: make(map[int]*User),
	}
}

func (s *Server) tryCache(id int) (*User, bool) {
	user, ok := s.userCache[id]

	return user, ok
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idString)

	user, ok := s.tryCache(id)
	if ok {
		json.NewEncoder(w).Encode(user)
		return
	}

	// hit the database
	user, ok = s.db[id]
	if !ok {
		panic("user not found")
	}
	s.dbhit++

	// insert into user cache
	s.userCache[id] = user

	json.NewEncoder(w).Encode(user)

	return
}

func main() {}
