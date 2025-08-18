package dto

type UpdateRequestDTO struct {
{{- range $name, $prop := .Properties }}
  {{- $type := JSONTypeToGoType $prop.Type }}
  {{ $name | toPascal }} *{{ $type }} `json:"{{ $name | toSnake }},omitempty"` 
{{- end }}
}

func NewUpdateRequestDTO() *UpdateRequestDTO {
	return &UpdateRequestDTO{}
}

func (e UpdateRequestDTO) Validate() error {
	return nil
}
