// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UsageMetric - An individual metric about usage of a user.
type UsageMetric struct {
	// The user ID associated with the metric
	UserID *string `json:"UserID,omitempty"`
	// The creator ID associated with the metric
	CreatorID *string `json:"CreatorID,omitempty"`
	// Total minutes of delivery usage.
	DeliveryUsageMins *float64 `json:"DeliveryUsageMins,omitempty"`
	// Total transcoded minutes.
	TotalUsageMins *float64 `json:"TotalUsageMins,omitempty"`
	// Total minutes of storage usage.
	StorageUsageMins *float64 `json:"StorageUsageMins,omitempty"`
}

func (o *UsageMetric) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UsageMetric) GetCreatorID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *UsageMetric) GetDeliveryUsageMins() *float64 {
	if o == nil {
		return nil
	}
	return o.DeliveryUsageMins
}

func (o *UsageMetric) GetTotalUsageMins() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalUsageMins
}

func (o *UsageMetric) GetStorageUsageMins() *float64 {
	if o == nil {
		return nil
	}
	return o.StorageUsageMins
}
