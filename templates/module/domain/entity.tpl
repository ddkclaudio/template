package domain

type Entity struct {
{{- range $name, $prop := .Properties }}
  {{- $type := JSONTypeToGoType $prop.Type }}
  {{- if not (isRequired $name $.Required) }}
  {{ $name | toPascal }} *{{ $type }} `json:"{{ toSnake $name }}"` 
  {{- else }}
  {{ $name | toPascal }} {{ $type }} `json:"{{ toSnake $name }}"` 
  {{- end }}
{{- end }}
}

func (e *Entity) Validate() error {
	return nil
}

func (*Entity) TableName() string {
	return "{{ pluralize (toSnake .Title) }}"
}
