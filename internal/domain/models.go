package domain

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Img   string `json:"img"`
	Count string `json:"count"`
	IsAll bool   `json:"isAll,omitempty"`
}

type Item struct {
	ID              string   `json:"id"`
	OwnerID         string   `json:"ownerId"`
	OwnerName       string   `json:"ownerName"`
	Title           string   `json:"title"`
	Category        string   `json:"category"`
	Model           string   `json:"model"`
	Manufacturer    string   `json:"manufacturer"`
	Series          string   `json:"series"`
	Scale           string   `json:"scale"`
	Completeness    string   `json:"completeness"`
	Condition       string   `json:"condition"`
	PrivateNote     string   `json:"privateNote"`
	Visibility      string   `json:"visibility"`
	ModelCode       string   `json:"modelCode"`
	CustomPhoto     string   `json:"customPhoto"`
	Price           string   `json:"price"`
	Location        string   `json:"location"`
	Date            string   `json:"date"`
	LinkedListingID string   `json:"linkedListingId"`
	Likes           []string `json:"likes"`
}

type WishProposal struct {
	ListingID   string `json:"listingId"`
	ItemTitle   string `json:"itemTitle"`
	OfferedBy   string `json:"offeredBy"`
	OfferPrice  string `json:"offerPrice"`
	Status      string `json:"status"`
	StatusLabel string `json:"statusLabel"`
	Comment     string `json:"comment"`
}

type Wish struct {
	ID            string         `json:"id"`
	AuthorID      string         `json:"authorId"`
	AuthorName    string         `json:"authorName"`
	Title         string         `json:"title"`
	Version       string         `json:"version"`
	ConditionDiff string         `json:"conditionDiff"`
	Budget        string         `json:"budget"`
	Visibility    string         `json:"visibility"`
	Status        string         `json:"status"`
	ModelCode     string         `json:"modelCode"`
	Proposals     []WishProposal `json:"proposals"`
}

type Clarification struct {
	ClarifID  string `json:"clarifId"`
	AuthorID  string `json:"authorId"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

type Reply struct {
	ID               string          `json:"id"`
	VersionID        string          `json:"versionId"`
	AuthorID         string          `json:"authorId"`
	AuthorName       string          `json:"authorName"`
	Text             string          `json:"text"`
	SourceURL        string          `json:"sourceUrl"`
	SellerDisclosure string          `json:"sellerDisclosure"`
	CreatedAt        string          `json:"createdAt"`
	HelpfulByAuthor  bool            `json:"helpfulByAuthor"`
	Clarifications   []Clarification `json:"clarifications"`
}

type Question struct {
	ID         string  `json:"id"`
	AuthorID   string  `json:"authorId"`
	AuthorName string  `json:"authorName"`
	Title      string  `json:"title"`
	Context    string  `json:"context"`
	LinkedType string  `json:"linkedType"`
	LinkedID   string  `json:"linkedId"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
	Replies    []Reply `json:"replies"`
}

type Recognition struct {
	RewardID             string   `json:"rewardId"`
	QuestionID           string   `json:"questionId"`
	QuestionTitle        string   `json:"questionTitle"`
	RecipientID          string   `json:"recipientId"`
	RecipientName        string   `json:"recipientName"`
	MarkedByAuthorName   string   `json:"markedByAuthorName"`
	SupportingReplyIDs   []string `json:"supportingReplyIds"`
	SupportingVersionIDs []string `json:"supportingVersionIds"`
	Status               string   `json:"status"`
	Archived             bool     `json:"archived"`
	ConsentVersion       string   `json:"consentVersion"`
	DecisionID           string   `json:"decisionId"`
	RevokeReason         string   `json:"revokeReason"`
	AppealStatus         string   `json:"appealStatus"`
	UpdatedAt            string   `json:"updatedAt"`
}

type AppState struct {
	Categories   []Category    `json:"categories"`
	Items        []Item        `json:"items"`
	Wishes       []Wish        `json:"wishes"`
	Questions    []Question    `json:"questions"`
	Recognitions []Recognition `json:"recognitions"`
}
