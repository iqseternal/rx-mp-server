package auth

import "testing"

func TestVerifyRefershToken(t *testing.T) {

	var token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAiLCJpc3MiOiJyYXBpZF9yaiIsImV4cCI6MTc0MzA0NzgwOSwibmJmIjoxNzQyOTYxNDA5LCJpYXQiOjE3NDI5NjE0MDksImp0aSI6IjEwIn0.KVDsH6d81v6qadZMu0wMEs5cS-48TxPH4mQUuPdcLHtVx6Lhxe85QZDaTOurI3vkdBcTbzs3XO_zI66lYkUWhQ"

	claims, err := VerifyRefershToken(token)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("claims:", claims)
}
