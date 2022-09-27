package comment

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/repo/unique"
	unique2 "github.com/segmentfault/answer/internal/service/unique"
)

var (
	dataSource *data.Data
	log        log.log
)

func init() {
	s, _ := os.LookupEnv("TESTDATA-DB-CONNECTION")
	cache, _, _ := data.NewCache(log.Getlog(), &data.CacheConf{})
	dataSource, _, _ = data.NewData(log.Getlog(), data.NewDB(true, &data.Database{
		Connection: s,
	}), cache)
	log = log.Getlog()
}

func Test_commentRepo_AddComment(t *testing.T) {
	type fields struct {
		log          log.log
		data         *data.Data
		uniqueIDRepo unique2.UniqueIDRepo
	}
	type args struct {
		ctx     context.Context
		comment *entity.Comment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "AddComment",
			fields: fields{

				data:         dataSource,
				uniqueIDRepo: unique.NewUniqueIDRepo(log, dataSource),
			},
			args: args{
				ctx: nil,
				comment: &entity.Comment{
					UserID:       "123",
					ObjectID:     "555",
					VoteCount:    0,
					Status:       1,
					OriginalText: "12312312",
					ParsedText:   "123123123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &commentRepo{
				log:          tt.fields.log,
				data:         tt.fields.data,
				uniqueIDRepo: tt.fields.uniqueIDRepo,
			}
			if err := cr.AddComment(tt.args.ctx, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commentRepo_GetComment(t *testing.T) {
	type fields struct {
		log          log.log
		data         *data.Data
		uniqueIDRepo unique2.UniqueIDRepo
	}
	type args struct {
		ctx       context.Context
		commentID string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantComment *entity.Comment
		wantExist   bool
		wantErr     bool
	}{
		{
			name: "test",
			fields: fields{

				data:         dataSource,
				uniqueIDRepo: unique.NewUniqueIDRepo(log, dataSource),
			},
			args: args{
				ctx:       nil,
				commentID: "10070000000000236",
			},
			wantComment: nil,
			wantExist:   false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &commentRepo{
				log:          tt.fields.log,
				data:         tt.fields.data,
				uniqueIDRepo: tt.fields.uniqueIDRepo,
			}
			gotComment, gotExist, err := cr.GetComment(tt.args.ctx, tt.args.commentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotComment, tt.wantComment) {
				t.Errorf("GetComment() gotComment = %v, want %v", gotComment, tt.wantComment)
			}
			if gotExist != tt.wantExist {
				t.Errorf("GetComment() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
