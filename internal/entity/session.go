package entity

type Session struct {
	UserUid   string `db:"user_uid"`
	Token     string `db:"token"`
	ExpiresAt int64  `db:"expires_at"`
	IssuedAt  int64  `db:"issued_at"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQwNjgxMTcsImlhdCI6MTcxNDA2NzUxNywicGhvbmVudW1iZXIiOiI5MDA5MDA5MDkwIiwidWlkIjoiMTBiNDNmOTItZWRhMC00ZWFhLTkwNTktNjVkMWFkYzdhNzM3In0.9rntx_CFGxCMByOlZa0zFLhIW4LNgQpruKX6wNSKXUo"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQwNzExMTcsImlhdCI6MTcxNDA2NzUxNywidWlkIjoiMTBiNDNmOTItZWRhMC00ZWFhLTkwNTktNjVkMWFkYzdhNzM3In0.2UVtNZ-lKub9ZrlDiwf0tD9Wx2e73QhdKeRxbtFwTE4"`
}
