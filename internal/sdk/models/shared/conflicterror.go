// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConflictError - standard error
type ConflictError struct {
	Status   interface{} `json:"status"`
	Title    interface{} `json:"title"`
	Type     interface{} `json:"type,omitempty"`
	Instance interface{} `json:"instance"`
	Detail   interface{} `json:"detail"`
}

func (o *ConflictError) GetStatus() interface{} {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ConflictError) GetTitle() interface{} {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ConflictError) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConflictError) GetInstance() interface{} {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ConflictError) GetDetail() interface{} {
	if o == nil {
		return nil
	}
	return o.Detail
}
