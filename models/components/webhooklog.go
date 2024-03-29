// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Request struct {
	// URL used for the request
	URL *string `json:"url,omitempty"`
	// HTTP request method
	Method *string `json:"method,omitempty"`
	// HTTP request headers
	Headers map[string]string `json:"headers,omitempty"`
	// request body
	Body *string `json:"body,omitempty"`
}

func (o *Request) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Request) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *Request) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *Request) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

type Response struct {
	// response body
	Body *string `json:"body,omitempty"`
	// HTTP status code
	Status *float64 `json:"status,omitempty"`
	// response status text
	StatusText *string `json:"statusText,omitempty"`
}

func (o *Response) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Response) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Response) GetStatusText() *string {
	if o == nil {
		return nil
	}
	return o.StatusText
}

type WebhookLog struct {
	ID string `json:"id"`
	// ID of the webhook this request was made for
	WebhookID string `json:"webhookId"`
	// The event type that triggered the webhook request
	Event *string `json:"event,omitempty"`
	// Timestamp (in milliseconds) at which webhook request object was
	// created
	//
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// The time taken (in seconds) to make the webhook request
	Duration *float64 `json:"duration,omitempty"`
	// Whether the webhook request was successful
	Success  *bool     `json:"success,omitempty"`
	Request  *Request  `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
}

func (o *WebhookLog) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WebhookLog) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}

func (o *WebhookLog) GetEvent() *string {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *WebhookLog) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WebhookLog) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *WebhookLog) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *WebhookLog) GetRequest() *Request {
	if o == nil {
		return nil
	}
	return o.Request
}

func (o *WebhookLog) GetResponse() *Response {
	if o == nil {
		return nil
	}
	return o.Response
}
