# Dharma Engine

A **vector-based cognitive engine** inspired by **Buddhist dependent origination** (pratītyasamutpāda), implemented in Go.

This repository contains a working engine that represents mental conditions as continuous vectors evolving over time and selects/updates skillful responses (processes) based on gradient alignment, decay, and expectation checks.

> **Uncertainty is not an error. It is a valid state.**

---

## 🚀 Overview

Dharma Engine is **not**:

- a chatbot
- a therapy tool
- a rule engine
- a psychology simulation

Instead, it is a **mathematical model** where:

- mental aggregates (conditions) are **vectors**
- skillful methods (upāya) are **directional perturbations**
- everything is continuous, decays, and must prove effectiveness

The system enforces **core Buddhist-aligned architectural invariants** rather than pre-defined behaviors or scripts.

---

## 📐 Core Principles (Architecture)

These are **non-negotiable invariants** of the engine:

1. **Everything is a vector** — states and influences live in one vector space.
2. **No forced narratives** — uncertainty is tolerated, not resolved prematurely.
3. **Processes must prove effectiveness** — if a method doesn’t reduce suffering, it is abandoned.
4. **All influence decays** — impermanence is encoded into math.
5. **Observation ends responsibility** — seeing doesn’t imply interpreting.
6. **Silence is valid** — inaction can be the correct response.

Violation of any invariant is a design bug.

---

## 🎯 Key Concepts

### Vectors

All internal states live in a unified vector space.

A vector captures intensity along axes such as:

- **Affliction** (greed, aversion, delusion)
- **Manifestation** (thought, speech, action)
- **Temporal** (arising, abiding, fading)
- **Capacity** (readiness / openness)
- **Relation** (self, other, authority)
- **Method** (insight, silence, restraint)

Vectors support:

- addition (`Add`)
- scaling (`Scaled`)
- bounded value interpretation

---

### Condition

A `Condition` is a transient set of current mental aggregates.

```go
type Condition struct {
    Vector domain.Vector
}
