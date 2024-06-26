// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ForbiddenError - standard error
type ForbiddenError struct {
	Status   interface{} `json:"status"`
	Title    interface{} `json:"title"`
	Type     interface{} `json:"type,omitempty"`
	Instance interface{} `json:"instance"`
	Detail   interface{} `json:"detail"`
}

func (o *ForbiddenError) GetStatus() interface{} {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ForbiddenError) GetTitle() interface{} {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ForbiddenError) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ForbiddenError) GetInstance() interface{} {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ForbiddenError) GetDetail() interface{} {
	if o == nil {
		return nil
	}
	return o.Detail
}
