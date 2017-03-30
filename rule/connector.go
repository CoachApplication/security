package rule

type Connector interface {
	Get(string) (Rule, error)
	Order() []string
}
