package rustore

type ApiResponse[T any] struct {
	Code      string  `json:"code"`
	Message   *string `json:"message"`
	Body      T       `json:"body"`
	Timestamp string  `json:"timestamp"`
}

type AgeRestriction struct {
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

type FileUrl struct {
	FileUrl     string `json:"fileUrl"`
	Ordinal     int    `json:"ordinal"`
	Type        string `json:"type"`
	Orientation string `json:"orientation"`
}

type Rating struct {
	Average float64 `json:"average"`
	Votes   int     `json:"votes"`
}

type AppInfo struct {
	AppId                int            `json:"appId"`
	PackageName          string         `json:"packageName"`
	AppName              string         `json:"appName"`
	Categories           []string       `json:"categories"`
	CompanyName          string         `json:"companyName"`
	ShortDescription     string         `json:"shortDescription"`
	FullDescription      string         `json:"fullDescription"`
	FileSize             int64          `json:"fileSize"`
	VersionName          string         `json:"versionName"`
	VersionCode          int            `json:"versionCode"`
	Downloads            int64          `json:"downloads"`
	Price                float64        `json:"price"`
	IconUrl              string         `json:"iconUrl"`
	FileUrls             []FileUrl      `json:"fileUrls"`
	WhatsNew             string         `json:"whatsNew"`
	AgeRestriction       AgeRestriction `json:"ageRestriction"`
	Rating               Rating         `json:"rating"`
	MinSdkVersion        int            `json:"minSdkVersion"`
	TargetSdkVersion     int            `json:"targetSdkVersion"`
	MinAndroidVersion    string         `json:"minAndroidVersion"`
	DeviceTypes          []string       `json:"deviceTypes"`
	RoundedDownloadsText string         `json:"roundedDownloadsText"`
	FirstPublishedAt     string         `json:"firstPublishedAt"`
	AppVerUpdatedAt      string         `json:"appVerUpdatedAt"`
	VersionId            int            `json:"versionId"`
}

type DownloadLinkRequest struct {
	AppId                int      `json:"appId"`
	FirstInstall         bool     `json:"firstInstall"`
	MobileServices       []string `json:"mobileServices"`
	SupportedAbis        []string `json:"supportedAbis"`
	ScreenDensity        int      `json:"screenDensity"`
	SupportedLocales     []string `json:"supportedLocales"`
	SdkVersion           int      `json:"sdkVersion"`
	WithoutSplits        bool     `json:"withoutSplits"`
	SignatureFingerprint *string  `json:"signatureFingerprint"`
}

type DownloadUrl struct {
	Url  string `json:"url"`
	Size int64  `json:"size"`
	Hash string `json:"hash"`
}

type PreApproveInfo struct {
	Label  string `json:"label"`
	Locale string `json:"locale"`
}

type DownloadLinkResponse struct {
	AppId          int            `json:"appId"`
	VersionCode    int            `json:"versionCode"`
	VersionId      int            `json:"versionId"`
	DownloadUrls   []DownloadUrl  `json:"downloadUrls"`
	PreApproveInfo PreApproveInfo `json:"preApproveInfo"`
	Signature      string         `json:"signature"`
}

type AppSummary struct {
	AppId        int
	PackageName  string
	AppName      string
	CompanyName  string
	Version      string
	VersionCode  int
	Downloads    string
	Rating       float64
	RatingVotes  int
	FileSize     string
	IconUrl      string
	DownloadUrl  *string
	DownloadSize *int64
}

type SearchSuggest struct {
	Text        string  `json:"text"`
	TextMarked  string  `json:"textMarked"`
	PackageName *string `json:"packageName"`
	IconUrl     *string `json:"iconUrl"`
}

type SearchSuggestResponse struct {
	Suggests       []SearchSuggest `json:"suggests"`
	History        []string        `json:"history"`
	SuggestQueryId string          `json:"suggestQueryId"`
}

type SearchAppItem struct {
	AppId                int     `json:"appId"`
	PackageName          string  `json:"packageName"`
	AppName              string  `json:"appName"`
	ShortDescription     string  `json:"shortDescription"`
	IconUrl              string  `json:"iconUrl"`
	AppType              string  `json:"appType"`
	VersionCode          int     `json:"versionCode"`
	MinSdkVersion        int     `json:"minSdkVersion"`
	Price                float64 `json:"price"`
	AverageUserRating    float64 `json:"averageUserRating"`
	TotalRatings         int     `json:"totalRatings"`
	LabelIds             []int   `json:"labelIds"`
	IsNavSuggest         bool    `json:"isNavSuggest"`
	RoundedDownloadsText string  `json:"roundedDownloadsText"`
}

type Spellchecker struct {
	OriginQuery   string `json:"originQuery"`
	ModifiedQuery string `json:"modifiedQuery"`
}

type SearchResponse struct {
	Content       []SearchAppItem `json:"content"`
	PageNumber    int             `json:"pageNumber"`
	PageSize      int             `json:"pageSize"`
	TotalElements int             `json:"totalElements"`
	TotalPages    int             `json:"totalPages"`
	SearchQueryId string          `json:"searchQueryId"`
	Spellchecker  Spellchecker    `json:"spellchecker"`
}
