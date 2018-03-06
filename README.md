# featuremill

[![Build Status](https://travis-ci.org/dustin-decker/featuremill.svg?branch=master)](https://travis-ci.org/dustin-decker/featuremill)

Featuremill is a general-purpose fast, stateless, and deterministic feature extractor writting in golang for use in machine learning.

The text feature extractor makes heavy use of a hashing vectorizer in order to rapidly generate features with very low overhead. Hash vectorizing does not require storing any state and is deterministic so it has these advantages:

- fast - does not need to perform index lookups or transformations because hashing vectorizer is used for generating feature IDs
- low memory consumption for large amounts of data - no index information or state stored in memory
- can easily be scaled horizontally - requires no stored state, so no coordination required
- features can be added at a later date - features are determinitstically generated

The advantages do not come without a cost, though. Featuremill cannot apply transformations such as Inverse Document Frequency (IDF) to text features, which can be helpful to discount the impact of common text features. If this has significant impact on prediction with your ML algorithm or dataset, consider adding that support or use different tooling or approach. Other optimizations, such as number scaling can be done if you're able to provide some basic metrics on your number data from your datastore, such as the maximum value in the dataset.

## Output format

[libsvm format](https://stats.stackexchange.com/questions/61328/libsvm-data-format) (`<index_number>:<value>`)
A slice of them, if the feature generates multiple vectors (text, timestamp, date).
This format makes this library easy to use along with [hector](https://github.com/xlvector/hector)

## Supported data

- IP - gets converted to integer representation
- timestamp - represeneted as 3 seasonality vectors: minute of hour, hour of day, day of week
- date - represented as 2 seasonality vectors: day of week, and month of year
- text - tokenized by word using hashing vectorizer
- numerical - min/max rescaler
- boolean - represented as 0 or 1
- categorical - like text, but not tokenized at all

## Documentation

### Godocs and tests

[godocs](https://godoc.org/github.com/dustin-decker/featuremill)

Look at the [included example](example/main.go) and look at the tests for example usage, as well.

### How I use featuremill

For each incoming sample being processed, I `append` the returned string or expanded slice to a `features` slice.
To assemble the final sample in libsvm format, just join it with the sample label:

``` go
sample := "0 " + strings.Join(features, " ")
```