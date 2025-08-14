package dto

import (
	"fmt"

	"github.com/library/utils"
	"github.com/google/uuid"
)

type ResponseDTO struct {
{{- range .Keys }}
  {{- $name := . }}
  {{- $prop := index $.Properties $name }}
  {{- if not (shouldIgnore $prop "response") }}
    {{- $type := JSONTypeToGoType $prop.Type }}
    {{- if not (isRequired $name $.Required) }}
  {{ $name | toPascal }} *{{ $type }} `json:"{{ toSnake $name }}"` 
    {{- else }}
  {{ $name | toPascal }} {{ $type }} `json:"{{ toSnake $name }}"` 
    {{- end }}
  {{- end }}
{{- end }}
}

func (e ResponseDTO) Validate() error {
{{- range .Keys }}
  {{- $name := . }}
  {{- $prop := index $.Properties $name }}
  {{- if not (shouldIgnore $prop "response") }}
    {{- $field := $name | toPascal }}

    {{- if isRequired $name $.Required }}
  if e.{{ $field }} == {{ zeroValue $prop.Type }} {
    return fmt.Errorf("field '{{ $name }}' is required")
  }
      {{- if eq $prop.Type "string" }}
        {{- if $prop.MinLength }}
  if err := utils.ValidateMinLength(e.{{ $field }}, {{ $prop.MinLength }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.MaxLength }}
  if err := utils.ValidateMaxLength(e.{{ $field }}, {{ $prop.MaxLength }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.Pattern }}
  if err := utils.ValidatePattern(e.{{ $field }}, `{{ $prop.Pattern }}`); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
      {{- else if or (eq $prop.Type "number") (eq $prop.Type "integer") }}
        {{- if $prop.Minimum }}
  if err := utils.ValidateMinimum(float64(e.{{ $field }}), {{ $prop.Minimum }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.Maximum }}
  if err := utils.ValidateMaximum(float64(e.{{ $field }}), {{ $prop.Maximum }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.ExclusiveMinimum }}
  if err := utils.ValidateExclusiveMinimum(float64(e.{{ $field }}), {{ $prop.ExclusiveMinimum }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.ExclusiveMaximum }}
  if err := utils.ValidateExclusiveMaximum(float64(e.{{ $field }}), {{ $prop.ExclusiveMaximum }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
        {{- if $prop.MultipleOf }}
  if err := utils.ValidateMultipleOf(float64(e.{{ $field }}), {{ $prop.MultipleOf }}); err != nil {
    return fmt.Errorf("field '{{ $name }}': %w", err)
  }
        {{- end }}
      {{- end }}
    {{- end }}

    {{- if not (isRequired $name $.Required) }}
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
{{- end }}

  return nil
}
