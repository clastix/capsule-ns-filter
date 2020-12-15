package request

type Request interface {
	GetUserAndGroups() (string, []string, error)
	IsNamespaceListing() (ok bool)
	IsNodeListing() (ok bool)
}
