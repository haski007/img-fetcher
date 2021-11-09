package fetcher

type Resource interface {
	GetTitle() string
	GetImages() []string
}
