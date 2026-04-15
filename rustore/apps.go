package rustore

import "fmt"

var defaultDeviceConfig = DownloadLinkRequest{
	FirstInstall:         true,
	MobileServices:       []string{},
	SupportedAbis:        []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	ScreenDensity:        420,
	SupportedLocales:     []string{"ru_RU"},
	SdkVersion:           34,
	WithoutSplits:        true,
	SignatureFingerprint: nil,
}

func GetAppInfo(packageName string) (AppInfo, error) {
	return apiGet[AppInfo](fmt.Sprintf("/applicationData/overallInfo/%s", packageName))
}

func GetDownloadLink(appId int) (DownloadLinkResponse, error) {
	req := defaultDeviceConfig
	req.AppId = appId
	return apiPost[DownloadLinkResponse, DownloadLinkRequest]("/applicationData/v2/download-link", req)
}

func GetAppSummary(packageName string) (AppSummary, error) {
	appInfo, err := GetAppInfo(packageName)
	if err != nil {
		return AppSummary{}, err
	}

	summary := AppSummary{
		AppId:       appInfo.AppId,
		PackageName: appInfo.PackageName,
		AppName:     appInfo.AppName,
		CompanyName: appInfo.CompanyName,
		Version:     appInfo.VersionName,
		VersionCode: appInfo.VersionCode,
		Downloads:   appInfo.RoundedDownloadsText,
		Rating:      appInfo.Rating.Average,
		RatingVotes: appInfo.Rating.Votes,
		FileSize:    FormatFileSize(appInfo.FileSize),
		IconUrl:     appInfo.IconUrl,
	}

	downloadInfo, err := GetDownloadLink(appInfo.AppId)
	if err == nil && len(downloadInfo.DownloadUrls) > 0 {
		url := downloadInfo.DownloadUrls[0].Url
		size := downloadInfo.DownloadUrls[0].Size
		summary.DownloadUrl = &url
		summary.DownloadSize = &size
	}

	return summary, nil
}
