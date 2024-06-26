// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsumerGroup struct {
	Name string   `json:"name"`
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
}

func (o *ConsumerGroup) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConsumerGroup) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ConsumerGroup) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
