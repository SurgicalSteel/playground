package main

import (
	"net/url"
)

type URLWrapper struct {
	URL     url.URL
	FullURL string
	//IsValid bool
	err error
}

type URLParser interface {
	GetHostName() (string, error)
	GetPath() (string, error)
	GetQueryMap() (map[string]string, error)
	GetPathWithoutQuery() (string, error)
}

func (u *URLWrapper) GetHostName() (string, error) {
	if u.err != nil {
		return "", u.err
	}
	return u.URL.Hostname(), nil
}

func (u *URLWrapper) GetPath() (string, error) {
	if u.err != nil {
		return "", u.err
	}
	return u.URL.Path, nil
}

func (u *URLWrapper) GetQueryMap() (map[string]string, error) {
	if u.err != nil {
		return map[string]string{}, u.err
	}
	m := make(map[string]string)
	for k, _ := range u.URL.Query() {
		m[k] = u.URL.Query().Get(k)
	}
	return m, nil
}

func (u *URLWrapper) GetPathWithoutQuery() (string, error) {
	if u.err != nil {
		return "", u.err
	}
	return "", nil
}

func Initialize(rawUrl string) URLParser {
	urlWrapper := URLWrapper{}
	urlObject, err := url.Parse(rawUrl)
	if err != nil {
		urlWrapper.err = err
		urlWrapper.FullURL = ""
		return &urlWrapper
	}

	urlWrapper.FullURL = rawUrl
	urlWrapper.URL = *urlObject
	return &urlWrapper
}
