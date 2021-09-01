package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/testdata"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/metadata"
	"reflect"
	"testing"
)

func Test_authorizer_claims(t *testing.T) {
	if err := testdata.LoadDotEnv(); err != nil {
		t.Fatal(err)
	}
	idTok, err := testdata.IdToken()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		ctx context.Context
		tok *oauth2.Token
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountInfo
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "gcloud auth print-identity-token",
			args: args{
				ctx: context.Background(),
				tok: &oauth2.Token{
					AccessToken: idTok,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got map[string]interface{}
			auth := &authorizer{
				env: environment.DefaultEnvironment,
			}
			gotClaims, err := auth.claims(tt.args.ctx, tt.args.tok)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := json.Marshal(gotClaims)
			if err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatal(err)
			}
			for k, v := range got {
				if v == nil {
					t.Errorf("value for key: %s found empty", k)
				}
			}
		})
	}
}

func Test_authorizer_tokenFromContext(t *testing.T) {
	md := metadata.New(map[string]string{
		authorizationKey: fmt.Sprintf("%s%s", bearerPrefix, "notasecret"),
	})
	ctx := metadata.NewIncomingContext(context.Background(), md)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *oauth2.Token
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "notempty",
			args: args{
				ctx: ctx,
			},
			want: &oauth2.Token{
				AccessToken: "notasecret",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := &authorizer{
				env: environment.DefaultEnvironment,
			}
			got, err := auth.tokenFromContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenFromContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenFromContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}
