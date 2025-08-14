package dto

import (
	"fmt"

	"github.com/library/utils"
)

type UpdateRequestDTO struct {
  {{- range .Keys }}
    {{- $name := . }}
    {{- $prop := index $.Properties $name }}
    {{- if not (shouldIgnore $prop "update") }}
      {{- $type := JSONTypeToGoType $prop.Type }}
    {{ $name | toPascal }} *{{ $type }} `json:"{{ toSnake $name }},omitempty"`
    {{- end }}
  {{- end }}
}

func (e UpdateRequestDTO) Validate() error {
{{- range .Keys }}
  {{- $name := . }}
  {{- $prop := index $.Properties $name }}
  {{- if not (shouldIgnore $prop "update") }}
    {{- $field := $name | toPascal }}
    {{- if or $prop.MinLength $prop.MaxLength $prop.Pattern $prop.Minimum $prop.Maximum $prop.ExclusiveMinimum $prop.ExclusiveMaximum $prop.MultipleOf }}
  if e.{{ $field }} != nil {
    {{- if eq $prop.Type "string" }}
      {{- if $prop.MinLength }}
    if err := utils.ValidateMinLength(*e.{{ $field }}, {{ $prop.MinLength }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.MaxLength }}
    if err := utils.ValidateMaxLength(*e.{{ $field }}, {{ $prop.MaxLength }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.Pattern }}
    if err := utils.ValidatePattern(*e.{{ $field }}, `{{ $prop.Pattern }}`); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
    {{- else if or (eq $prop.Type "number") (eq $prop.Type "integer") }}
      {{- if $prop.Minimum }}
    if err := utils.ValidateMinimum(float64(*e.{{ $field }}), {{ $prop.Minimum }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.Maximum }}
    if err := utils.ValidateMaximum(float64(*e.{{ $field }}), {{ $prop.Maximum }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.ExclusiveMinimum }}
    if err := utils.ValidateExclusiveMinimum(float64(*e.{{ $field }}), {{ $prop.ExclusiveMinimum }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.ExclusiveMaximum }}
    if err := utils.ValidateExclusiveMaximum(float64(*e.{{ $field }}), {{ $prop.ExclusiveMaximum }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
      {{- if $prop.MultipleOf }}
    if err := utils.ValidateMultipleOf(float64(*e.{{ $field }}), {{ $prop.MultipleOf }}); err != nil {
      return fmt.Errorf("field '{{ $name }}': %w", err)
    }
      {{- end }}
    {{- end }}
  }
    {{- end }}
  {{- end }}
{{- end }}

  return nil
}
