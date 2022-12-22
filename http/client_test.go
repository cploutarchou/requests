package http

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_goHTTPClient_Get(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{

		{
			name: "Test goHTTPClient Get",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewBuilder()
			client := builder.Build()
			res, err := client.Get(tt.args.url, tt.fields.Headers.GetAll())
			got := &Response{
				statusCode: res.StatusCode(),
				status:     res.Status(),
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("goHTTPClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goHTTPClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goHTTPClient_Post(t *testing.T) {
	type User struct {
		Name     string `json:"name"`
		UserName string `json:"username"`
	}
	user := User{
		Name:     "Christos Ploutarchou",
		UserName: "christos",
	}
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Post",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url: "https://api.github.com",
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				client.DisableTimeouts()
				res, err := client.Post(tt.args.url, tt.fields.Headers.GetAll(), user)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Post() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Post() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Put(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Put",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Put(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Put() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Put() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Patch(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Patch",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Patch(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Patch() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Patch() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Delete(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Delete",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Delete(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Delete() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Delete() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_goHTTPClient_Head(t *testing.T) {
	type fields struct {
		Headers Headers
		Timeout Timeout
	}
	type args struct {
		url     string
		headers http.Header
		body    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name: "Test goHTTPClient Head",
			fields: fields{
				Headers: NewHeaders(),
				Timeout: newTimeouts(),
			},
			args: args{
				url:     "https://api.github.com",
				headers: nil,
				body:    nil,
			},
			want: &Response{
				statusCode: 200,
				status:     "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				builder := NewBuilder()
				client := builder.Build()
				res, err := client.Head(tt.args.url, tt.fields.Headers.GetAll(), tt.args.body)
				got := &Response{
					statusCode: res.StatusCode(),
					status:     res.Status(),
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("goHTTPClient.Head() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("goHTTPClient.Head() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
