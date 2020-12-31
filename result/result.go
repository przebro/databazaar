package result

type BazaarResult struct {
	ID       string `json:"id" bson:"_id"`
	Revision string `json:"rev" bson:"rev"`
}
