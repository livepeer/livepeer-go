// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// ViewershipMetric - An individual metric about viewership of an asset. Necessarily, at least
// 1 of playbackId and dStorageUrl will be present, depending on the query.
type ViewershipMetric struct {
	// The browser used by the viewer.
	Browser *string `json:"browser,omitempty"`
	// The browser engine used by the viewer's browser.
	BrowserEngine *string `json:"browserEngine,omitempty"`
	// The continent where the viewer is located.
	Continent *string `json:"continent,omitempty"`
	// The country where the viewer is located.
	Country *string `json:"country,omitempty"`
	// The CPU used by the viewer's device.
	CPU *string `json:"cpu,omitempty"`
	// The ID of the creator associated with the metric.
	CreatorID *string `json:"creatorId,omitempty"`
	// The URL of the distributed storage used for the asset
	DStorageURL *string `json:"dStorageUrl,omitempty"`
	// The device used by the viewer.
	Device *string `json:"device,omitempty"`
	// The type of the device used by the viewer.
	DeviceType *string `json:"deviceType,omitempty"`
	// The error rate for the asset.
	ErrorRate *float64 `json:"errorRate,omitempty"`
	// The percentage of sessions that existed before the asset started
	// playing.
	//
	ExitsBeforeStart *float64 `json:"exitsBeforeStart,omitempty"`
	// Geographic encoding of the viewers location. Accurate to 3 digits.
	Geohash *string `json:"geohash,omitempty"`
	// The operating system used by the viewer.
	Os *string `json:"os,omitempty"`
	// The playback ID associated with the metric.
	PlaybackID *string `json:"playbackId,omitempty"`
	// The total playtime in minutes for the asset.
	PlaytimeMins float64 `json:"playtimeMins"`
	// The rebuffering ratio for the asset.
	RebufferRatio *float64 `json:"rebufferRatio,omitempty"`
	// The subdivision (e.g., state or province) where the viewer is
	// located.
	//
	Subdivision *string `json:"subdivision,omitempty"`
	// Timestamp (in milliseconds) when the metric was recorded. If the
	// query contains a time step, this timestamp will point to the
	// beginning of the time step period.
	//
	Timestamp *float64 `json:"timestamp,omitempty"`
	// The timezone where the viewer is located.
	Timezone *string `json:"timezone,omitempty"`
	// The time-to-first-frame (TTFF) in milliseconds.
	TtffMs *float64 `json:"ttffMs,omitempty"`
	// The number of views for the asset.
	ViewCount int64 `json:"viewCount"`
	// The ID of the viewer associated with the metric.
	ViewerID *string `json:"viewerId,omitempty"`
}

func (o *ViewershipMetric) GetBrowser() *string {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *ViewershipMetric) GetBrowserEngine() *string {
	if o == nil {
		return nil
	}
	return o.BrowserEngine
}

func (o *ViewershipMetric) GetContinent() *string {
	if o == nil {
		return nil
	}
	return o.Continent
}

func (o *ViewershipMetric) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *ViewershipMetric) GetCPU() *string {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *ViewershipMetric) GetCreatorID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *ViewershipMetric) GetDStorageURL() *string {
	if o == nil {
		return nil
	}
	return o.DStorageURL
}

func (o *ViewershipMetric) GetDevice() *string {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *ViewershipMetric) GetDeviceType() *string {
	if o == nil {
		return nil
	}
	return o.DeviceType
}

func (o *ViewershipMetric) GetErrorRate() *float64 {
	if o == nil {
		return nil
	}
	return o.ErrorRate
}

func (o *ViewershipMetric) GetExitsBeforeStart() *float64 {
	if o == nil {
		return nil
	}
	return o.ExitsBeforeStart
}

func (o *ViewershipMetric) GetGeohash() *string {
	if o == nil {
		return nil
	}
	return o.Geohash
}

func (o *ViewershipMetric) GetOs() *string {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *ViewershipMetric) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *ViewershipMetric) GetPlaytimeMins() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaytimeMins
}

func (o *ViewershipMetric) GetRebufferRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.RebufferRatio
}

func (o *ViewershipMetric) GetSubdivision() *string {
	if o == nil {
		return nil
	}
	return o.Subdivision
}

func (o *ViewershipMetric) GetTimestamp() *float64 {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

func (o *ViewershipMetric) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *ViewershipMetric) GetTtffMs() *float64 {
	if o == nil {
		return nil
	}
	return o.TtffMs
}

func (o *ViewershipMetric) GetViewCount() int64 {
	if o == nil {
		return 0
	}
	return o.ViewCount
}

func (o *ViewershipMetric) GetViewerID() *string {
	if o == nil {
		return nil
	}
	return o.ViewerID
}
