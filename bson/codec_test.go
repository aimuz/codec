package bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func Test_bsonCodec_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name    string
		c       c
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "map",
			c:    *Codec,
			args: args{
				v: map[string]interface{}{
					"int":    1,
					"string": "hello world",
				},
			},
			want: []byte{38, 0, 0, 0, 16, 105, 110, 116, 0, 1, 0, 0, 0, 2, 115, 116, 114, 105, 110, 103, 0, 12, 0, 0, 0, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 0, 0},
		},
		{
			name: "struct",
			c:    *Codec,
			args: args{
				v: struct {
					Value string
				}{
					Value: "hello world",
				},
			},
			want: func() []byte {
				b, err := bson.Marshal(struct {
					Value string
				}{
					Value: "hello world",
				})
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
		},
		{
			name: "struct-full",
			c:    *Codec,
			args: args{
				v: struct {
					Value string
					Map   map[string]interface{}
				}{
					Value: "hello world",
					Map: map[string]interface{}{
						"int":    1,
						"string": "hello world",
					},
				},
			},
			want: func() []byte {
				b, err := bson.Marshal(struct {
					Value string
					Map   map[string]interface{}
				}{
					Value: "hello world",
					Map: map[string]interface{}{
						"int":    1,
						"string": "hello world",
					},
				})
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
		},
		{
			name: "bson",
			c:    *Codec,
			args: args{
				v: bson.M{
					"int":    1,
					"string": "hello world",
				},
			},
			want: func() []byte {
				b, err := bson.Marshal(bson.M{
					"int":    1,
					"string": "hello world",
				})
				if err != nil {
					t.Fatal(err)
				}
				return b
			}(),
		},
		{
			name: "string failed",
			c:    *Codec,
			args: args{
				v: "hello world",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bsonCodec_Name(t *testing.T) {
	tests := []struct {
		name string
		c    c
		want string
	}{
		{
			name: name,
			c:    *Codec,
			want: name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bsonCodec_Unmarshal(t *testing.T) {
	type args struct {
		b []byte
		v interface{}
	}
	tests := []struct {
		name    string
		c       c
		args    args
		wantErr bool
	}{
		{
			name: "map",
			c:    *Codec,
			args: args{
				b: func() []byte {
					b, err := bson.Marshal(map[string]interface{}{
						"int": 1,
					})
					if err != nil {
						t.Fatal(err)
					}
					return b
				}(),
				v: &map[string]interface{}{},
			},
			wantErr: false,
		},
		{
			name: "map2struct",
			c:    *Codec,
			args: args{
				b: func() []byte {
					b, err := bson.Marshal(map[string]interface{}{
						"int": 1,
					})
					if err != nil {
						t.Fatal(err)
					}
					return b
				}(),
				v: &struct {
					Int int `bson:"int"`
				}{},
			},
			wantErr: false,
		},
		{
			name: "struct",
			c:    *Codec,
			args: args{
				b: func() []byte {
					b, err := bson.Marshal(struct {
						Value string
					}{
						Value: "hello world",
					})
					if err != nil {
						t.Fatal(err)
					}
					return b
				}(),
				v: &struct {
					Value string
				}{},
			},
			wantErr: false,
		},
		{
			name: "struct-full",
			c:    *Codec,
			args: args{
				b: func() []byte {
					b, err := bson.Marshal(struct {
						Value string
						Map   map[string]interface{}
					}{
						Value: "hello world",
						Map: map[string]interface{}{
							"int":    1,
							"string": "hello world",
						},
					})
					if err != nil {
						t.Fatal(err)
					}
					return b
				}(),
				v: &struct {
					Value string
					Map   map[string]interface{}
				}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Unmarshal(tt.args.b, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
