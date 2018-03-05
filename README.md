# feature extractor

general-purpose fast, stateless, and deterministic feature extractor writting in golang for use in machine learning

## Supported data

- IP (gets converted to integer representation)
- text (tokenized by word)

## TODO supported data

- time (broken out into seasonality buckets - 3 vectors)
- boolean (0/1)
- numbers (gets scaled)