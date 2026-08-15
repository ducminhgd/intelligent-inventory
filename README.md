# Intelligent Inventory Dashboard

This repository is the submission for one of four challenges for Keyloop's Engineering Team Lead.

## Detail requirements

1. Domain: Supply
2. Task: Build an Intelligent Inventory Dashboard to give dealership managers a real-time overview of their vehicle stock.
3. Core Requirements:
   1. Inventory Visualization: Display a filterable list of all vehicles in a dealership's inventory (e.g., filter by make, model, age).
   2. Aging Stock Identification: Automatically identify and prominently display "aging stock" (vehicles in inventory for >90 days).
   3. Actionable Insights: Allow a manager to log and persist a status or proposed action for each aging vehicle (e.g., "Price Reduction Planned").

## AI Usages

Client: Claude Code
LLM: Deepseek

I list all my steps here, both implementations by me and implementations by AI:
1. **I** created ![architectural design](./docs/00-architectural-design.md) and ![database design document](./docs/01-database-design.md)
2. **I** created ![AI Stack](./.claude)
3. **I** init the project with Go
4. **AI** verifies the docs with this init prompt:

   ```
   This project is just a POC, as a @"principal-swe (agent)", you read "## Detail requirements" in @README.md and verify what I designed in @docs/00-architectural-design.md and @docs/01-database-design.md, use /grill-me to discuss with me. After all, store the conversations in @docs/ai-logs/00-verify-docs.md
   ```
5. **I** implement the project layout, the full flow for APIs of manufacturer.
6. **AI** helps to review the code:

   ```
   As a @.claude/agents/principal-swe.md , /code-review the existing code. Remember that the scope for this project is just an POC.
   ```
