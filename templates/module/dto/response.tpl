package dto

type ResponseDTO struct {
{{- range $name, $prop := .Properties }}
  {{- $type := JSONTypeToGoType $prop.Type }}
  {{- if isRequired $name $.Required }}
  {{ $name | toPascal }} {{ $type }} `json:"{{ $name | toSnake }}"` 
  {{- else }}
  {{ $name | toPascal }} *{{ $type }} `json:"{{ $name | toSnake }}"` 
  {{- end }}
{{- end }}
}

func NewResponseDTO() *ResponseDTO {
	return &ResponseDTO{}
}

func (e ResponseDTO) Validate() error {
	return nil
}
