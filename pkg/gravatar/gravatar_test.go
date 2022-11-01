package gravatar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvatarURL(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "answer@answer.com",
			args: args{email: "answer@answer.com"},
			want: "https://www.gravatar.com/avatar/b2be4e4438f08a5e885be8de5f41fdd7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetAvatarURL(tt.args.email))
		})
	}
}

func TestResize(t *testing.T) {
	type args struct {
		originalAvatarURL string
		sizePixel         int
	}
	tests := []struct {
		name                 string
		args                 args
		wantResizedAvatarURL string
	}{
		{
			name: "original url",
			args: args{
				originalAvatarURL: "https://www.gravatar.com/avatar/b2be4e4438f08a5e885be8de5f41fdd7",
				sizePixel:         128,
			},
			wantResizedAvatarURL: "https://www.gravatar.com/avatar/b2be4e4438f08a5e885be8de5f41fdd7?s=128",
		},
		{
			name: "already resized url",
			args: args{
				originalAvatarURL: "https://www.gravatar.com/avatar/b2be4e4438f08a5e885be8de5f41fdd7?s=128",
				sizePixel:         64,
			},
			wantResizedAvatarURL: "https://www.gravatar.com/avatar/b2be4e4438f08a5e885be8de5f41fdd7?s=64",
		},
		{
			name: "empty url",
			args: args{
				originalAvatarURL: "",
				sizePixel:         64,
			},
			wantResizedAvatarURL: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantResizedAvatarURL, Resize(tt.args.originalAvatarURL, tt.args.sizePixel), "Resize(%v, %v)", tt.args.originalAvatarURL, tt.args.sizePixel)
		})
	}
}
