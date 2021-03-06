package caching

import (
	"github.com/google/uuid"
	"github.com/project-midgard/midgarts/pkg/graphic"
)

type CachedTextureProvider map[uuid.UUID]*graphic.Texture

func NewCachedTextureProvider() CachedTextureProvider {
	return make(map[uuid.UUID]*graphic.Texture)
}

func (t CachedTextureProvider) NewTextureFromRGBA(rgba *graphic.UniqueRGBA) (*graphic.Texture, error) {
	if txt, ok := t[rgba.ID]; ok {
		return txt, nil
	}

	tex, err := graphic.NewTextureFromRGBA(rgba)
	if err != nil {
		return nil, err
	}

	t[rgba.ID] = tex
	return tex, nil
}
