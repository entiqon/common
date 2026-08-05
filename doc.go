// Package common provides foundational building blocks shared across Entiqon.
//
// It contains reusable utilities, error handling extensions, reflection helpers,
// pointer utilities, and strongly typed parsers for primitive and structured
// values.
//
// Subpackages include:
//
//   - collection: generic typed collections with rich helper methods.
//   - errors: structured error types (CausableError, ProcessStageError).
//   - extension: strongly typed value parsers and helpers:
//   - boolean: flexible boolean parsing (true/false, yes/no, on/off…)
//   - date: date parsing and normalization.
//   - decimal: decimal parsing with precision control.
//   - float: floating-point parsing.
//   - number: integer parsing and rounding.
//   - object: reflection helpers (Exists, GetValue, SetValue).
//   - ptr: generic pointer utilities (Of, FromPtr, Deref, NilIfEmpty,
//     NilIfBlank, EmptyString).
//   - text: common string utilities such as padding, truncation, and
//     presentation helpers.
//
// The common module enables consistency and code reuse across Entiqon's
// ecosystem.
package common
