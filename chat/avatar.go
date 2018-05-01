package main

import (
	"github.com/pkg/errors"
)

var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}
type GravatarAvatar struct{}

var UseAuthAvatar AuthAvatar
var UseGravaratAvatar GravatarAvatar

// Auth認証の場合
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

// Gravatarを使った場合
func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}

	return "", ErrNoAvatarURL
}
