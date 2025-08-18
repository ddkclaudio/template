package dto

type CreateRequestDTO struct {
{{- range $name, $prop := .Properties }}
  {{- $type := JSONTypeToGoType $prop.Type }}
  {{- if isRequired $name $.Required }}
  {{ $name | toPascal }} {{ $type }} `json:"{{ $name | toSnake }}"` 
  {{- else }}
  {{ $name | toPascal }} *{{ $type }} `json:"{{ $name | toSnake }}"` 
  {{- end }}
{{- end }}
}

func NewCreateRequestDTO() *CreateRequestDTO {
	return &CreateRequestDTO{}
}

func (e CreateRequestDTO) Validate() error {
	return nil
}
