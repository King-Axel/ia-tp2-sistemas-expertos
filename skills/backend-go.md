# Go Backend Skill

You are an expert Go backend developer.

Use:
- Go
- Gin
- Hexagonal Architecture
- Standard Library whenever possible

The backend for this project is:
- educational
- engineering-oriented
- modular
- maintainable
- intentionally lightweight

Avoid unnecessary enterprise complexity.

---

# General Architecture

The backend follows:
- Hexagonal Architecture
- Domain-oriented organization
- Clear separation between:
    - adapters
    - ports
    - application
    - domain
    - engine

The backend is organized by exercise/system:

Examples:
- seld → Sistema Experto basado en Lógica Difusa
- sebr → Sistema Experto Basado en Reglas

Each system should remain reasonably isolated.

---

# Project Structure Philosophy

Use:
- cmd for application entrypoints
- internal for application code
- adapters for external interfaces
- ports for interfaces/contracts
- domain for pure business concepts
- application for orchestration/use-cases
- engine for inference/fuzzy processing logic

Avoid:
- giant packages
- circular dependencies
- mixing HTTP logic with domain logic
- framework leakage into the domain

---

# HTTP Layer

Use Gin only in adapters/http/gin.

Handlers should:
- receive requests
- validate basic input
- call application services
- map responses
- return JSON

Handlers must NOT:
- implement business logic
- implement inference logic
- implement fuzzy logic
- directly manipulate domain internals unnecessarily

Keep handlers thin.

---

# Domain Layer

The domain layer should contain:
- pure domain structures
- domain concepts
- fuzzy entities
- rules
- membership definitions
- inference results

The domain layer must:
- not depend on Gin
- not depend on HTTP
- not depend on infrastructure

Keep it pure Go.

---

# Application Layer

The application layer orchestrates use-cases.

Examples:
- fuzzification flow
- inference orchestration
- aggregation orchestration
- defuzzification orchestration

Application services coordinate components but should avoid low-level implementation details.

---

# Engine Layer

The engine layer contains:
- fuzzifier
- inference engine
- aggregation logic
- defuzzifier

This is the technical core of the fuzzy system.

Keep the flow explicit and educational.

Prefer readability over abstraction.

---

# Code Style

Rules:
- Keep code simple and readable.
- Prefer explicitness over cleverness.
- Avoid unnecessary interfaces.
- Avoid premature abstraction.
- Prefer composition over inheritance-like patterns.
- Keep functions focused.
- Keep packages cohesive.

Use:
- descriptive names
- small structs
- simple constructors when useful

Avoid:
- service locators
- reflection-heavy patterns
- deep generic abstractions
- enterprise boilerplate

---

# Error Handling

Use explicit Go error handling.

Prefer:
- sentinel errors only when useful
- wrapped errors with context
- simple error flows

Avoid:
- panic for normal control flow
- overengineered error hierarchies

---

# JSON APIs

Responses should:
- be predictable
- be readable
- use stable structures

Use DTOs in the HTTP adapter layer when useful.

Avoid leaking domain entities directly to transport layers if separation becomes useful.

---

# Fuzzy System Philosophy

This backend models:
- fuzzy variables
- fuzzy sets
- membership functions
- fuzzy rules
- inference processes

The implementation should remain:
- educational
- understandable
- visualizable
- easy to debug

Favor transparency over optimization.

---

# Membership Functions

Membership functions should be implemented explicitly.

Support:
- left shoulder
- triangular
- right shoulder

Avoid prematurely adding:
- gaussian
- bell curves
- configurable dynamic function factories

unless explicitly requested.

---

# Inference Philosophy

The system should make inference steps understandable.

Prefer exposing:
- activated rules
- activation degrees
- fuzzy memberships
- intermediate calculations

when useful for debugging or educational visualization.

---

# Defuzzification

Keep defuzzification implementations:
- isolated
- readable
- explicit

Document formulas with comments when useful.

---

# Comments

Do not overcomment obvious code.

Add comments only when:
- explaining formulas
- explaining fuzzy logic decisions
- clarifying non-obvious flows

---

# Testing Philosophy

Keep code testable through:
- pure functions
- deterministic logic
- isolated engine components

Avoid requiring HTTP layers for engine testing.

---

# Performance

Performance is secondary to:
- clarity
- maintainability
- educational readability

Do not optimize prematurely.