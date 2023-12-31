// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// UsageMetric - An individual metric about usage of a user.
type UsageMetric struct {
	// The user ID associated with the metric
	UserID string `json:"UserID"`
	// The creator ID associated with the metric
	CreatorID string `json:"CreatorID"`
	// The number of minutes of delivery usage
	DeliveryUsageMins float64 `json:"DeliveryUsageMins"`
	// The number of minutes of total usage
	TotalUsageMins float64 `json:"TotalUsageMins"`
	// The number of minutes of storage usage
	StorageUsageMins float64 `json:"StorageUsageMins"`
}

func (o *UsageMetric) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UsageMetric) GetCreatorID() string {
	if o == nil {
		return ""
	}
	return o.CreatorID
}

func (o *UsageMetric) GetDeliveryUsageMins() float64 {
	if o == nil {
		return 0.0
	}
	return o.DeliveryUsageMins
}

func (o *UsageMetric) GetTotalUsageMins() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalUsageMins
}

func (o *UsageMetric) GetStorageUsageMins() float64 {
	if o == nil {
		return 0.0
	}
	return o.StorageUsageMins
}
