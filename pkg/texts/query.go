package texts

import (
	"slices"
	"strconv"
	"strings"
)

const (
	QueryWordDelimiter  = " "
	QueryParamDelimiter = "&"
	QueryValueDelimiter = "="
	QuerySliceDelimiter = ","
	QueryParamPrefix    = "/"
)

type Query struct {
	Texts  []string
	Params map[string][]string
}

func QueryFromParams(params map[string][]string) *Query {
	return &Query{
		Texts:  nil,
		Params: params,
	}
}

func (q *Query) GetTextsSlice() []string {
	return q.Texts
}

func (q *Query) GetTextsInto(into *string) {
	if len(q.Texts) == 0 {
		return
	}

	*into = strings.Join(q.Texts, QueryWordDelimiter)
}

func (q *Query) GetInt64(key string) (int64, bool) {
	values, ok := q.Params[key]
	if !ok {
		return 0, false
	}

	if len(values) == 0 {
		return 0, true
	}

	res, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return 0, false
	}

	return res, true
}

func (q *Query) GetInt64Slice(key string) ([]int64, bool) {
	values, ok := q.Params[key]
	if !ok {
		return nil, false
	}

	res := make([]int64, 0, len(values))

	for _, value := range values {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, false
		}

		res = append(res, val)
	}

	return res, true
}

func (q *Query) GetInt64Into(key string, into *int64) bool {
	value, ok := q.GetInt64(key)
	if !ok {
		return false
	}

	*into = value

	return true
}

func (q *Query) GetString(key string) (string, bool) {
	values, ok := q.Params[key]
	if !ok {
		return "", false
	}

	if len(values) == 0 {
		return "", true
	}

	return values[0], true
}

func (q *Query) GetStringInto(key string, into *string) bool {
	value, ok := q.GetString(key)
	if !ok {
		return false
	}

	*into = value

	return true
}

func (q *Query) GetStringSlice(key string) ([]string, bool) {
	values, ok := q.Params[key]
	if !ok {
		return nil, false
	}

	return values, true
}

func (q *Query) Encode() string {
	parts := make([]string, 0, len(q.Texts)+len(q.Params))

	params := make([]string, 0, len(q.Params))

	for key, values := range q.Params {
		if len(values) == 0 {
			params = append(params, key)

			continue
		}

		params = append(params, key+QueryValueDelimiter+strings.Join(values, QuerySliceDelimiter))
	}

	if len(params) > 0 {
		slices.Sort(params)

		parts = append(parts, QueryParamPrefix+strings.Join(params, QueryParamDelimiter))
	}

	parts = append(parts, q.Texts...)

	return strings.Join(parts, QueryWordDelimiter)
}

func DecodeQuery(query string) Query {
	words := strings.Split(query, QueryWordDelimiter)

	res := Query{
		Texts:  nil,
		Params: make(map[string][]string, len(words)),
	}

	nonParamCurrentWords := make([]string, 0, len(words))

	for _, word := range words {
		if word == "" {
			continue
		}

		if !strings.HasPrefix(word, QueryParamPrefix) {
			nonParamCurrentWords = append(nonParamCurrentWords, word)

			continue
		}

		if len(nonParamCurrentWords) > 0 {
			res.Texts = append(res.Texts, strings.Join(nonParamCurrentWords, QueryWordDelimiter))
			nonParamCurrentWords = make([]string, 0, len(words))
		}

		paramsList := strings.Split(word[1:], QueryParamDelimiter)

		for _, param := range paramsList {
			parts := strings.SplitN(param, QueryValueDelimiter, 2)

			var key string

			var values []string

			if len(parts) >= 1 {
				key = parts[0]
			}

			if len(parts) >= 2 {
				values = strings.Split(parts[1], QuerySliceDelimiter)
			}

			res.Params[key] = append(res.Params[key], values...)
		}
	}

	if len(nonParamCurrentWords) > 0 {
		res.Texts = append(res.Texts, strings.Join(nonParamCurrentWords, QueryWordDelimiter))
	}

	return res
}
