package pkgredis

type (
	SetRequest struct {
		Key   string
		Value interface{}
		Exp   int64
	}

	SetResponse struct{}

	GetRequest struct {
		Key string
	}

	GetResponse struct {
		Result string
	}

	DelRequest struct {
		Key string
	}

	DelResponse struct {
		Result string
	}

	IncrRequest struct {
		Key string
	}

	IncrResponse struct {
		Result int64
	}

	ExpireRequest struct {
		Key     string
		Seconds int64
	}

	ExpireResponse struct {
		Success bool
	}
)
