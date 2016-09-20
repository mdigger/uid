# Generator of globally unique identifiers

[![GoDoc](https://godoc.org/github.com/mdigger/uid?status.svg)](https://godoc.org/github.com/mdigger/uid)
[![Build Status](https://travis-ci.org/mdigger/uid.svg)](https://travis-ci.org/mdigger/uid)
[![Coverage Status](https://coveralls.io/repos/github/mdigger/uid/badge.svg?branch=master)](https://coveralls.io/github/mdigger/uid?branch=master)

The algorithm used to generate globally unique identifiers based on the same principle that is used to generate unique IDs in MongoDB. The unique ID is a 12 byte sequence consisting of the time of generation, computer ID, and process, as well as counter. 

The main difference is that the identifier is represented as a string using base64-encoding. It also supports a function to quickly parse this string, which returns information about all the values, which are assembled from the identifier.

	V-B9WTRe2V45jQUU, V-B9WTRe2V45jQUV, V-B9WTRe2V45jQUW